import yaml
from yaml.loader import SafeLoader
import json
import sys
import subprocess
import os
from pathlib import Path


flag = sys.argv[1]


BASE_DIR = Path(__file__).resolve().parent


def push(reg_url, reg_id, reg_pw):

    reg_prefix = '/target_'

    img_prefix = 'target_'

    docker_compose_yaml = os.path.join(BASE_DIR,'docker-compose.yaml')

    in_file = open(docker_compose_yaml, 'r')
    dc_file = yaml.safe_load(in_file)
    in_file.close()

    REG_URL =  reg_url
    REG_ID = reg_id
    REG_PW = reg_pw

    REG_URL_AUTH = reg_url.split('/',2)[0]

    subprocess.run(["docker","logout"])

    subprocess.run(['docker','login',REG_URL_AUTH,'-u',REG_ID,'-p',REG_PW])

    key_list = list(dc_file["services"].keys())


    for key in key_list:

        container_name = dc_file["services"][key]["container_name"]

        container_name = container_name.replace('_','-')
        
        dest = REG_URL + reg_prefix + container_name

        source = img_prefix + key

        subprocess.run(['docker','tag',source,dest])

        subprocess.run(['docker','push',dest])


    target = os.path.join(BASE_DIR,'target')
    
    subprocess.run(['docker-compose','down'],cwd=target)



def deploy(reg_url, reg_id, reg_pw):

        
    kompose = os.path.join(BASE_DIR,'kompose')

    ops_src_yaml = os.path.join(BASE_DIR,'ops_src.yaml')

    docker_compose_yaml = os.path.join(BASE_DIR,'docker-compose.yaml')

    subprocess.run([kompose,'-f',docker_compose_yaml,'-o',ops_src_yaml,'convert'])

    in_file = open(ops_src_yaml, 'r')
    yaml_file = yaml.safe_load(in_file)
    in_file.close()


    reg_prefix = '/target_'

    reg_prefix = reg_url + reg_prefix

    for item in yaml_file["items"]:

        if item["kind"] == 'Deployment':

            image_pull_secrets = [{"name": 'docker-secret'}]

            item["spec"]["template"]["spec"]["imagePullSecrets"] = image_pull_secrets

            containers = item["spec"]["template"]["spec"]["containers"]

            for container in containers:

                image_name = reg_prefix + container["name"]

                container["image"] = image_name

                container["imagePullPolicy"] = 'Always'


    out_file = open(ops_src_yaml, 'w')
    yaml.safe_dump(yaml_file, out_file)
    out_file.close()




def kill(obj_rsc, obj_nm):


    ops_src_yaml = os.path.join(BASE_DIR,'ops_src.yaml')

    readf = open(ops_src_yaml,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    kill_doc = []

    objs = list_yaml[0]["items"]

    for obj in objs:
        if obj["metadata"]["name"] == obj_nm:
        

            kill_doc.append(obj)

            

    readf.close()

    delete_ops_src_yaml = os.path.join(BASE_DIR,'delete_ops_src.yaml')

    writef = open(delete_ops_src_yaml,'w')

    yaml.safe_dump_all(kill_doc,writef)

    writef.close()





def hpa(obj_rsc, obj_nm, kcfg_path):

    ops_src_yaml = os.path.join(BASE_DIR,'ops_src.yaml')

    res = subprocess.run(['kubectl','--kubeconfig',kcfg_path,'get','nodes','-o','yaml'],capture_output=True,text=True)


    loaded_yaml = yaml.load_all(res.stdout,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    doc_yaml = list_yaml[0]


    prev_top = 0

    pods = 0

    for i in range(len(doc_yaml["items"])):

        pods = int(doc_yaml["items"][i]["status"]["allocatable"]["pods"])

        if pods > prev_top:
            prev_top = pods

    pods = prev_top



    optimal_upper_bound = int(pods * 0.2)
    optimal_target = int(optimal_upper_bound * 0.25)




    if obj_rsc == 'deployment':

        resource_type = 'Deployment'

    minRepl = optimal_target
    maxRepl = optimal_upper_bound




    head_metadataName = 'hpa-deployment-'+obj_nm
    apiVersion = ''
    kind = resource_type
    metadata_name = obj_nm


    readf = open(ops_src_yaml,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    objs = list_yaml[0]["items"]


    for obj in objs:

        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:

            apiVersion = obj["apiVersion"] 
            kind = obj["kind"]
            metadata_name = obj["metadata"]["name"]


    readf.close()

    hpa_tmpl_yaml = os.path.join(BASE_DIR,'hpa-tmpl.yaml')

    readf = open(hpa_tmpl_yaml,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)   

    doc_yaml = list_yaml[0]

    doc_yaml["metadata"]["name"] = head_metadataName

    doc_yaml["spec"]["scaleTargetRef"]["apiVersion"] = apiVersion

    doc_yaml["spec"]["scaleTargetRef"]["kind"] = kind

    doc_yaml["spec"]["scaleTargetRef"]["name"] = metadata_name 

    doc_yaml["spec"]["minReplicas"] = minRepl

    doc_yaml["spec"]["maxReplicas"] = maxRepl

    readf.close()

    hpa_yaml = os.path.join(BASE_DIR,'hpa.yaml')

    writef = open(hpa_yaml,'w')

    yaml.safe_dump(doc_yaml,writef)

    writef.close()





def qos(obj_rsc, obj_nm, code, kcfg_path):


    # node spec allocatable and extracted values
    res = subprocess.run(['kubectl','--kubeconfig',kcfg_path,'get','nodes','-o','yaml'],capture_output=True,text=True)


    loaded_yaml = yaml.load_all(res.stdout,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    doc_yaml = list_yaml[0]

    polled_node_idx = -1

    prev_top = 0

    pods = 0

    for i in range(len(doc_yaml["items"])):

        pods = int(doc_yaml["items"][i]["status"]["allocatable"]["pods"])

        if pods > prev_top:
            prev_top = pods
            polled_node_idx = i

    polled_cpu = doc_yaml["items"][polled_node_idx]["status"]["allocatable"]["cpu"]
    polled_mem = doc_yaml["items"][polled_node_idx]["status"]["allocatable"]["memory"]

    polled_cpu = float(polled_cpu) * 1000.0
    polled_mem = float(polled_mem.replace('Ki',''))

    cpu_limit_per_node = (polled_cpu / pods) * 8 

    mem_limit_per_node = (polled_mem / pods) * 16



    # qos
    qos_cpu_high = str(int(cpu_limit_per_node * 0.8)) + 'm'

    qos_mem_high = str(int(mem_limit_per_node * 0.8)) + 'Ki'

    qos_cpu_middle = str(int(cpu_limit_per_node * 0.5)) + 'm'

    qos_mem_middle = str(int(mem_limit_per_node * 0.5)) + 'Ki'



    if obj_rsc == 'deployment':

        resource_type = 'Deployment'

    


    if code == 'highest':

        cpu_limits = qos_cpu_middle

        mem_limits = qos_mem_middle 

        cpu_requests = qos_cpu_middle

        mem_requests = qos_mem_middle 

    elif code == 'higher':

        cpu_limits = qos_cpu_high

        mem_limits = qos_mem_high

        cpu_requests = qos_cpu_middle

        mem_requests = qos_mem_middle 


    ops_src_yaml = os.path.join(BASE_DIR,'ops_src.yaml')

    readf = open(ops_src_yaml,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    qos_doc = []

    objs = list_yaml[0]["items"]

    for obj in objs:
        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:
        
            containers = obj["spec"]["template"]["spec"]["containers"]
            
            for cont in containers :
                rsc = {"limits":{"cpu":cpu_limits,"memory":mem_limits},"requests":{"cpu":cpu_requests,"memory":mem_requests}}
                cont["resources"] = rsc

            qos_doc = obj

            

    readf.close()

    qos_yaml = os.path.join(BASE_DIR,'qos.yaml')

    writef = open(qos_yaml,'w')

    yaml.safe_dump(qos_doc,writef)

    writef.close()



def qosundo(obj_rsc, obj_nm):




    if obj_rsc == 'deployment':

        resource_type = 'Deployment'
 

    ops_src_yaml = os.path.join(BASE_DIR,'ops_src.yaml')

    readf = open(ops_src_yaml,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    qos_doc = []

    objs = list_yaml[0]["items"]

    for obj in objs:
        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:
        

            qos_doc = obj

            

    readf.close()

    qos_yaml = os.path.join(BASE_DIR,'qos.yaml')

    writef = open(qos_yaml,'w')

    yaml.safe_dump(qos_doc,writef)

    writef.close()





def ingr(ns,host_nm,svc_nm,kcfg_path):

    res = subprocess.run(['kubectl',"--kubeconfig",kcfg_path,'get', '-n',ns,'service',svc_nm,'-o','yaml'],capture_output=True,text=True)


    loaded_yaml = yaml.load_all(res.stdout,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    doc_yaml = list_yaml[0]

    port_number = doc_yaml["spec"]["ports"][0]["port"]

    ingr_tmpl_yaml = os.path.join(BASE_DIR,'ingr-tmpl.yaml')

    in_file = open(ingr_tmpl_yaml, 'r')
    ingr_file = yaml.safe_load(in_file)
    in_file.close()


    ingr_file["metadata"]["name"] = 'ingress-'+ns

    ingr_file["spec"]["rules"][0]["host"] = host_nm

    ingr_file["spec"]["rules"][0]["http"]["paths"][0]["backend"]["service"]["name"] = svc_nm

    ingr_file["spec"]["rules"][0]["http"]["paths"][0]["backend"]["service"]["port"]["number"] = port_number

    ingr_yaml = os.path.join(BASE_DIR,'ingr.yaml')

    out_file = open(ingr_yaml, 'w')
    yaml.safe_dump(ingr_file, out_file)
    out_file.close()




if flag == 'push':
    
    reg_url = sys.argv[2]

    reg_id = sys.argv[3]

    reg_pw = sys.argv[4]

    push(reg_url, reg_id, reg_pw)

elif flag == 'deploy':

    reg_url = sys.argv[2]

    reg_id = sys.argv[3]

    reg_pw = sys.argv[4]

    deploy(reg_url, reg_id, reg_pw)

elif flag == 'kill':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    kill(rsc, rsc_nm)

elif flag == 'hpa':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    kcfg_path = sys.argv[4]

    hpa(rsc,rsc_nm, kcfg_path)

elif flag == 'qos':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    code = sys.argv[4]

    kcfg_path = sys.argv[5]

    qos(rsc,rsc_nm, code, kcfg_path)

elif flag == 'qosundo':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    qosundo(rsc,rsc_nm)

elif flag == 'ingr':

    ns = sys.argv[2]

    host_nm = sys.argv[3]

    svc_nm = sys.argv[4]

    kcfg_path = sys.argv[5]

    ingr(ns,host_nm,svc_nm,kcfg_path)



