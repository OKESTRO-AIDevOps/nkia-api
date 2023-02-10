import yaml
from yaml.loader import SafeLoader
import json
import sys
import subprocess

flag = sys.argv[1]





def push(reg_url, reg_id, reg_pw):

    reg_prefix = '/target_'

    img_prefix = 'target_'

    in_file = open('docker-compose.yaml', 'r')
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
    
    subprocess.run(['docker-compose','down'],cwd='./target')



def deploy(reg_url, reg_id, reg_pw):


    in_file = open('ops_src.yaml', 'r')
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


    out_file = open('ops_src.yaml', 'w')
    yaml.safe_dump(yaml_file, out_file)
    out_file.close()




def kill(obj_rsc, obj_nm):




    readf = open('ops_src.yaml','r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    kill_doc = []

    objs = list_yaml[0]["items"]

    for obj in objs:
        if obj["metadata"]["name"] == obj_nm:
        

            kill_doc.append(obj)

            

    readf.close()

    writef = open('delete_ops_src.yaml','w')

    yaml.safe_dump_all(kill_doc,writef)

    writef.close()





def hpa(obj_rsc, obj_nm):



    res = subprocess.run(['kubectl','get','nodes','-o','yaml'],capture_output=True,text=True)


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


    readf = open('./ops_src.yaml','r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    objs = list_yaml[0]["items"]


    for obj in objs:

        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:

            apiVersion = obj["apiVersion"] 
            kind = obj["kind"]
            metadata_name = obj["metadata"]["name"]


    readf.close()

    readf = open('./hpa-tmpl.yaml','r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)   

    doc_yaml = list_yaml[0]

    doc_yaml["metadata"]["name"] = head_metadataName

    doc_yaml["spec"]["scaleTargetRef"]["apiVersion"] = apiVersion

    doc_yaml["spec"]["scaleTargetRef"]["kind"] = kind

    doc_yaml["spec"]["scaleTargetRef"]["name"] = metadata_name 

    doc_yaml["spec"]["minReplicas"] = minRepl

    doc_yaml["spec"]["maxReplicas"] = maxRepl


    writef = open('hpa.yaml','w')

    yaml.safe_dump(doc_yaml,writef)

    writef.close()





def qos(obj_rsc, obj_nm):


    # node spec allocatable and extracted values
    res = subprocess.run(['kubectl','get','nodes','-o','yaml'],capture_output=True,text=True)


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

    



    cpu_limits = qos_cpu_high

    mem_limits = qos_mem_high

    cpu_requests = qos_cpu_middle

    mem_requests = qos_mem_middle 

 


    readf = open('ops_src.yaml','r',encoding='utf8')

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

    writef = open('qos.yaml','w')

    yaml.safe_dump(qos_doc,writef)

    writef.close()



def qosundo(obj_rsc, obj_nm):




    if obj_rsc == 'deployment':

        resource_type = 'Deployment'
 


    readf = open('ops_src.yaml','r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    qos_doc = []

    objs = list_yaml[0]["items"]

    for obj in objs:
        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:
        

            qos_doc = obj

            

    readf.close()

    writef = open('qos.yaml','w')

    yaml.safe_dump(qos_doc,writef)

    writef.close()





def ingr(ns,host_nm,svc_nm):

    res = subprocess.run(['kubectl','get', '-n',ns,'service',svc_nm,'-o','yaml'],capture_output=True,text=True)


    loaded_yaml = yaml.load_all(res.stdout,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    doc_yaml = list_yaml[0]

    port_number = doc_yaml["spec"]["ports"][0]["port"]

    in_file = open('ingr-tmpl.yaml', 'r')
    ingr_file = yaml.safe_load(in_file)
    in_file.close()


    ingr_file["metadata"]["name"] = 'ingress-'+ns

    ingr_file["spec"]["rules"][0]["host"] = host_nm

    ingr_file["spec"]["rules"][0]["http"]["paths"][0]["backend"]["service"]["name"] = svc_nm

    ingr_file["spec"]["rules"][0]["http"]["paths"][0]["backend"]["service"]["port"]["number"] = port_number

    out_file = open('ingr.yaml', 'w')
    yaml.safe_dump(ingr_file, out_file)
    out_file.close()



def secretCheck(ns):

    ret_sign = 'NTOK'

    outs = subprocess.run(['kubectl','get','-n',ns,'secret','docker-secret'])

    so = str(outs.stdout,'utf-8')

    se = str(outs.stderr,'utf-8')

    if so != '' and se == '':

        ret_sign = 'OK'


    print(ret_sign)


def dcFormatCheck():

    endl = '*** 5. RESULT  ***\n'

    fata = '*** 3. FATAL ERROR ***\n'

    warn = '*** 2. POTENTIAL ERROR ***\n'

    etc = '*** 4. UNHANDLED EVENT  ***\n'

    info = '*** 1. MAIN INFORMATION ***\n'

    flag = 'OK:REP----------\n\n'

    fata_sig = 0

    warn_sig = 0

    etc_sig = 0

    info_sig = 0

    outs = subprocess.run(['./kompose','-o','ops_src.yaml','convert'],capture_output=True)

    so = str(outs.stdout,'utf-8')

    se = str(outs.stderr,'utf-8')


    if so == '' and se != '' :

        so = se

    so_list = so.splitlines()

    for elm in so_list :

        elm_li = elm.split('\x1b[0m')

        if 'WARN' in elm_li[0]:

            warn_sig = 1

            warn += '- '+ elm_li[1] + '\n'

        elif 'INFO' in elm_li[0]:

            info_sig = 1

            info += '- ' + elm_li[1] + '\n'


        elif 'FATA' in elm_li[0]:

            fata_sig = 1

            flag = 'NTOK:REP----------\n\n'

            fata += '- ' + elm_li[1] + '\n'

        else :

            etc_sig = 1

            etc += '- ' + elm_li[1] + '\n'


    if fata_sig == 0:

        fata += '- No Fatal Event to Be Reported \n'


    if warn_sig == 0:


        warn += '- No Potential Error Cause to Be Reported \n'

    if etc_sig == 0:

        etc += '- No Unhandled Event to Be Reported \n'

    if info_sig == 0:

        info += '- No Main Information to Be Reported \n'


    report = flag + info + warn + fata + etc + endl


    print(report)



def projectProbe(ns):


    project = '*** 1. Project ***\n'

    component = '*** 2. Component ***\n'

    cause = '*** 3. Root Cause ***\n'

    probe = '*** 4. Probe ***\n'

    log = '*** 5. In-Component Logs ***\n'

    report = ''

    outs = subprocess.run(['./probe.sh',ns],capture_output=True)

    so = str(outs.stdout,'utf-8')

    if so == '\n':

        report = 'Nothing To Probe\n'

        print(report)

    else :

        number = 1

        so_list = so.splitlines()

        rep_list = []

        for el in so_list:

            rep = ''

            tmp_el = el.split(' ')

            el = [ x for x in tmp_el if x != '']

            pr_nm = el[0]

            co_nm = el[1]

            ca_nm = el[2]

            prj_tmp = project + '- '+pr_nm +'\n'

            com_tmp = component + '- '+co_nm + '\n'

            cau_tmp = cause + '- ' + ca_nm + '\n'

            el_out = subprocess.run(['kubectl','get','event','-n',ns,'--field-selector','involvedObject.name='+co_nm],capture_output=True)

            el_so =  str(el_out.stdout,'utf-8')

            el_se =  str(el_out.stderr,'utf-8')

            if el_se != '':

                el_so = el_se

            prb_tmp = probe + el_so + '\n'

            el_out = subprocess.run(['kubectl','logs','-n',ns,'--tail=100','--all-containers=true',co_nm],capture_output=True)

            el_so =  str(el_out.stdout,'utf-8')

            el_se =  str(el_out.stderr,'utf-8')

            if el_se != '':

                el_so = el_se

            log_tmp = log + el_so + '\n'

            rep = str(number) + '\n'

            rep = rep + prj_tmp + com_tmp + cau_tmp + prb_tmp + log_tmp

            rep_list.append(rep)

            number += 1
        
        for i in range(len(rep_list)):

            report += rep_list[i]

        print(report)





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

    hpa(rsc,rsc_nm)

elif flag == 'qos':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    qos(rsc,rsc_nm)

elif flag == 'qosundo':

    rsc = sys.argv[2]

    rsc_nm = sys.argv[3]

    qosundo(rsc,rsc_nm)

elif flag == 'ingr':

    ns = sys.argv[2]

    host_nm = sys.argv[3]

    svc_nm = sys.argv[4]

    ingr(ns,host_nm,svc_nm)

elif flag == 'sc':

    ns = sys.argv[2]

    secretCheck(ns)

elif flag == 'dcfc':

    dcFormatCheck()

elif flag == 'pp':

    ns = sys.argv[2]

    projectProbe(ns)

