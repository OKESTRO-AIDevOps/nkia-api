
from flask import Flask, request as request_flask
import subprocess
import json
from pathlib import Path
import os
import pandas as pd
import secrets
import requests
from dotenv import load_dotenv,find_dotenv
import yaml
from yaml.loader import SafeLoader



app = Flask(__name__)
port = 7331
basedir = Path(__file__).resolve().parent
targetdirname = 'target'
targetpath = os.path.join(basedir, targetdirname)
load_dotenv(find_dotenv())


REPO_URL =os.environ.get('REPO_URL')
REPO_ID = os.environ.get('REPO_ID')
REPO_PW =os.environ.get('REPO_PW')
DOCKER_ID = os.environ.get('DOCKER_ID')
DOCKER_PW = os.environ.get('DOCKER_PW')
ANSWER = os.environ.get('ANSWER')

REPO_URL = REPO_URL%(REPO_ID,REPO_PW)



def auth(post_dict):

    auth_table = pd.read_csv('auth_table.csv',sep=',',skipinitialspace=True)
    sessionhit = 0
    authhit = 0
    token = ''
    record_id = 0
    retval = []

    for i in range(len(auth_table)):
        if auth_table.at[i,'SID'] != 'N' and auth_table.at[i,'SID'] == post_dict['SID']:
            sessionhit = 1
            break
    
    
    if sessionhit == 0 :
        for i in range(len(auth_table)):
            if auth_table.at[i,'ACD'] == post_dict['ACD']:
                record_id = i
                token  =  secrets.token_hex(128)
                post_dict['SID'] = token
                auth_table.at[record_id,'SID'] = token
                auth_table.at[record_id,'RDP'] = 'N'
                auth_table.at[record_id,'MNT'] = 'N'
                authhit = 1
    
    post_dict['ACD'] = 'N'


    if sessionhit == 0 and authhit == 0:
        post_dict['SID'] = 'N'
        post_dict['MSG'] = 'ACD'
        retval.append('ACD')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval
    
    post_dict['ACD'] = 'Y'


    if post_dict['RDP'] == 'Y' and auth_table.at[record_id,'RDP'] == 'N' and auth_table.at[record_id,'MNT'] == 'N':
        token = secrets.token_hex(128)
        auth_table.at[record_id,'RDP'] = token
        auth_table.at[record_id,'MSG'] = 'VRF'
        post_dict['RDP'] = token
        post_dict['MSG'] = 'VRF'
        retval.append('RDP')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval
        

    elif post_dict['MNT'] == 'Y' and auth_table.at[record_id,'MNT'] == 'N' and auth_table.at[record_id,'RDP'] == 'N':
        token = secrets.token_hex(128)
        auth_table.at[record_id,'MNT'] = token
        post_dict['MNT'] = token
        post_dict['MSG'] = 'MNT'
        retval.append('MNT')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval



    elif post_dict['RDP'] == auth_table.at[record_id,'RDP'] and auth_table.at[record_id,'RDP'] != 'N' and auth_table.at[record_id,'MSG'] == 'VRFBLD':
        auth_table.at[record_id,'MNT'] = 'N'
        auth_table.at[record_id,'RDP'] = 'N'
        auth_table.at[record_id,'MSG'] = 'N'
        post_dict['MNT'] = 'N'
        retval.append('RDPMNT')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval
    

    elif post_dict['MNT'] == auth_table.at[record_id,'MNT'] and auth_table.at[record_id,'MNT'] != 'N':
        
     
        if post_dict['CMD'] == 'TRM':
            auth_table.at[record_id,'SID'] = 'N'
            auth_table.at[record_id,'MNT'] = 'N'
            post_dict['ACD'] = 'N'
            post_dict['SID'] = 'N'
            post_dict['RDP'] = 'N'
            post_dict['MNT'] = 'N'
            post_dict['MSG'] = 'TRM'
            retval.append('TRM')
            retval.append(post_dict) 
            auth_table.to_csv('auth_table.csv',index=False)
            return retval
        
        
        token = secrets.token_hex(128)
        auth_table.at[record_id,'MNT'] = token
        post_dict['MNT'] = token
        post_dict['MSG'] = 'MNTMNT'
        retval.append('MNTMNT')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval


    else :
        post_dict['MSG'] = 'ERR'
        retval.append('ERR')
        retval.append(post_dict)
        auth_table.to_csv('auth_table.csv',index=False)
        return retval



def rdp(post_dict) :

    check_path = os.path.join(targetpath,'.git')

    path_exist = os.path.exists(check_path)

    if path_exist == True:

        res = subprocess.run(['rm','-r',check_path])

    res = subprocess.run(['git','init'],cwd=targetpath)

    res = subprocess.run(['git','pull',REPO_URL],cwd=targetpath)


    return post_dict




def rdpmnt(post_dict) :


    #res = subprocess.run(['rm',targetpath])

    print('build done')

    

    return post_dict


def mnt(post_dict):

    check_path = os.path.join(targetpath,'.git')

    path_exist = os.path.exists(check_path)

    if path_exist == True:

        res = subprocess.run(['rm','-r',check_path])

    res = subprocess.run(['git','init'],cwd=targetpath)

    res = subprocess.run(['git','pull',REPO_URL],cwd=targetpath)
  


    return post_dict



def mntmnt(post_dict) :


    cmd_blocs = post_dict["CMD"].split(':')

    action = cmd_blocs[0]

    object_chain = cmd_blocs[1]

    if action == 'GET':

       post_dict["CNT"] = actionGET(object_chain)
    
    elif action == 'APPLY':

       post_dict["CNT"] =  actionAPPLY(object_chain)
    
    elif action  == 'LIFE':

       post_dict["CNT"] = actionLIFE(object_chain)
    
    elif action == 'PA':

       post_dict["CNT"] =  actionPA(object_chain)


    return post_dict


def actionGET(object_chain):

    object_blocs = object_chain.split(',')

    obj_head = object_blocs[0]

    obj_ns = object_blocs[1]

    if obj_head == 'PDS':

        res = subprocess.run(['kubectl','get','pods','-n',obj_ns],capture_output=True,text=True)


    elif obj_head == 'SVC':

        res = subprocess.run(['kubectl','get','services','-n',obj_ns],capture_output=True,text=True)


    elif obj_head == 'DPL':

        res = subprocess.run(['kubectl','get','deployments','-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'NDS':

        res = subprocess.run(['kubectl','get','nodes'],capture_output=True,text=True)


    elif obj_head == 'EVN':

        res = subprocess.run(['kubectl','get','events','-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'RSC':

        res = subprocess.run(['kubectl','get','resources','-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'NSC': 

        res = subprocess.run(['kubectl','get','namespaces'],capture_output=True,text=True)

    return res.stdout

def actionAPPLY(object_chain):

    object_blocs = object_chain.split(',')

    obj_head = object_blocs[0]

    obj_ns = object_blocs[1]

    if obj_head == 'APP':

        res = subprocess.run(['kubectl','apply','-f','app.yaml','-n',obj_ns],capture_output=True,text=True,cwd=targetpath)

    elif obj_head == 'QOS': 

        obj_rsc = object_blocs[2]

        obj_nm = object_blocs[3]

        obj_status = object_blocs[4]

        qosFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status)

        res = subprocess.run(['kubectl','apply','-f','qos.yaml','-n',obj_ns],capture_output=True,text=True,cwd=targetpath)

    elif obj_head == 'CRTNS': 

        res = subprocess.run(['kubectl','create','namespace',obj_ns],capture_output=True,text=True)

    #elif obj_head == 'USENS': 

    #    res = subprocess.run(['kubectl','get','namespaces'],capture_output=True)

    return res.stdout


def qosFileGenerator(obj_ns,obj_rsc,obj_nm,obj_status):

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


    if obj_rsc == 'DPL':

        resource_type = 'Deployment'

    
    if obj_status == 'high':

        cpu_limits = qos_cpu_high

        mem_limits = qos_mem_high

        cpu_requests = qos_cpu_high

        mem_requests = qos_mem_high 

    elif obj_status == 'middle':

        cpu_limits = qos_cpu_high

        mem_limits = qos_mem_high

        cpu_requests = qos_cpu_middle

        mem_requests = qos_mem_middle 

    
    elif obj_status == 'low':

        cpu_limits = 0.0

        mem_limits = 0.0

        cpu_requests = 0.0

        mem_requests = 0.0 



    app_file_path = os.path.join(targetpath,'app.yaml')

    readf = open(app_file_path,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)



    for obj in list_yaml:
        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:
        
            containers = obj["spec"]["template"]["spec"]["containers"]
            
            for cont in containers :
                rsc = {"limits":{"cpu":cpu_limits,"memory":mem_limits},"requests":{"cpu":cpu_requests,"memory":mem_requests}}
                cont["resources"] = rsc

                if obj_status == 'low':

                    del cont["resources"]
            

    readf.close()

    writef = open(app_file_path,'w')

    yaml.dump_all(list_yaml,writef)

    writef.close()





    


def actionLIFE(object_chain):

    object_blocs = object_chain.split(',')

    obj_head = object_blocs[0]

    obj_ns = object_blocs[1]

    obj_rsc = object_blocs[2]

    obj_nm = object_blocs[3]

    kube_rscnm = obj_rsc + '/'+obj_nm

    if obj_head == 'RESTART':

        res = subprocess.run(['kubectl','rollout','restart',kube_rscnm,'-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'ROLLBACK': 

        res = subprocess.run(['kubectl','rollout','undo',kube_rscnm,'-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'HISTORY': 

        res = subprocess.run(['kubectl','rollout','history',kube_rscnm,'-n',obj_ns],capture_output=True,text=True)

    elif obj_head == 'KILL': 

        res = subprocess.run(['kubectl','delete',kube_rscnm,'-n',obj_ns],capture_output=True,text=True)

    return res.stdout


def actionPA(object_chain):

    object_blocs = object_chain.split(',')

    obj_head = object_blocs[0]

    obj_ns = object_blocs[1]

    obj_rsc = object_blocs[2]

    obj_nm = object_blocs[3]

    obj_status = object_blocs[4]

    if obj_head == 'HORIZONTAL':

        hpaFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status)

        res = subprocess.run(['kubectl','apply','-f','hpa.yaml','-n',obj_ns],capture_output=True,text=True,cwd=targetpath)

    elif obj_head == 'VERTICAL': 

        #vpaFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status)

        #res = subprocess.run(['kubectl','apply','-f','vpa.yaml','-n',obj_ns],capture_output=True,cwd=targetpath)

        hpaFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status)

        res = subprocess.run(['kubectl','apply','-f','hpa.yaml','-n',obj_ns],capture_output=True,text=True,cwd=targetpath)


    return res.stdout

def hpaFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status):

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


    max_upper_bound = int(pods * 0.4)
    max_target = int(max_upper_bound * 0.5)

    optimal_upper_bound = int(pods * 0.2)
    optimal_target = int(optimal_upper_bound * 0.25)

    min_upper_bound = int(pods * 0.1)
    min_target = int(1)

    max_upper_bound = str(max_upper_bound)
    max_target = str(max_target)

    optimal_upper_bound = str(optimal_upper_bound)
    optimal_target = str(optimal_target)

    min_upper_bound = str(min_upper_bound)
    min_target = str(min_target)



    if obj_rsc == 'DPL':

        resource_type = 'Deployment'

    
    if obj_status == 'max':

        minRepl = max_target
        maxRepl = max_upper_bound

    elif obj_status == 'optimal':

        minRepl = optimal_target
        maxRepl = optimal_upper_bound

    
    elif obj_status == 'min':

        minRepl = min_target
        maxRepl = min_upper_bound

    head_metadataName = 'hpa-deployment-'+obj_nm
    apiVersion = ''
    kind = resource_type
    metadata_name = obj_nm


    app_file_path = os.path.join(targetpath,'app.yaml')

    readf = open(app_file_path,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)


    for obj in list_yaml:

        if obj["kind"] == resource_type and obj["metadata"]["name"] == obj_nm:

            apiVersion = obj["apiVersion"] 
            kind = obj["kind"]
            metadata_name = obj["metadata"]["name"]


    readf.close()

    app_file_path = os.path.join(targetpath,'hpa-tmpl.yaml')

    readf = open(app_file_path,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)   

    doc_yaml = list_yaml[0]

    doc_yaml["metadata"]["name"] = head_metadataName

    doc_yaml["spec"]["scaleTargetRef"]["apiVersion"] = apiVersion

    doc_yaml["spec"]["scaleTargetRef"]["kind"] = kind

    doc_yaml["spec"]["scaleTargetRef"]["name"] = metadata_name 

    doc_yaml["spec"]["minReplicas"] = minRepl

    doc_yaml["spec"]["maxReplicas"] = maxRepl

    app_file_path = os.path.join(targetpath,'hpa.yaml')

    writef = open(app_file_path,'w')

    yaml.dump_all(doc_yaml,writef)

    writef.close()







def vpaFileGenerator(obj_ns, obj_rsc, obj_nm, obj_status):

    print(0)



@app.route('/', methods=['POST'])
def npiaremote():

    retval = []

    print('NPIA_remote event detected ----- \n')

    post_dict = request_flask.get_json()

    retval = auth(post_dict)

    if retval[0] == 'ACD' :

        
        return retval[1]

    elif retval[0] == 'RDP' :

        retval[1] = rdp(retval[1])
        
        return retval[1]

    elif retval[0] == 'MNT' :

        retval[1] = mnt(retval[1])
        
        return retval[1]

    elif retval[0] == 'RDPMNT' :
        
        retval[1] = rdpmnt(retval[1])
        
        return retval[1]

    elif retval[0] == 'MNTMNT' :
        
        retval[1] = mntmnt(retval[1])
        
        return retval[1]

    elif retval[0] == 'TRM' :

        
        return retval[1]

    elif retval[0] == 'ERR':
        
        return retval[1]


@app.route('/build', methods=['POST'])
def npiabuild():
    

   
    authhit = 0
    record_id = 0
    basedir = Path(__file__).resolve().parent
    targetpath = os.path.join(basedir, targetdirname)
    

    auth_table = pd.read_csv('auth_table.csv',sep=',',skipinitialspace=True)
    post_dict = request_flask.get_json()

    for i in range(len(auth_table)) :

        if auth_table.at[i,'RDP'] == post_dict['RDP'] and auth_table.at[i,'MSG'] == 'VRF' :
            
            record_id = i

            authhit = 1

        elif auth_table.at[i,'RDP'] == post_dict['RDP'] and auth_table.at[i,'MSG'] == 'ASK' :

            record_id = i

            authhit = 2

        elif auth_table.at[i,'RDP'] == post_dict['RDP'] and auth_table.at[i,'MSG'] == 'VRFBLD' :

            record_id = i

            authhit = 3

        elif auth_table.at[i,'RDP'] == post_dict['RDP'] and auth_table.at[i,'MSG'] == 'BLDERR' :

            record_id = i

            authhit = 4



    if authhit == 0 :

        post_dict['MSG'] = 'VRFBFB'
        

    elif authhit == 1:

        print('NPIA_build event detected, auth success, build initiated....')

        subprocess.Popen(['python3','builder.py',str(record_id)])

        post_dict['MSG'] = 'ASK'

        auth_table.at[record_id,'MSG'] = 'ASK'

        auth_table.to_csv('auth_table.csv',index=False)

        

    elif authhit == 2:

        post_dict['CNT'] = 'STDLOG'
        

    elif authhit == 3:

        post_dict['MSG'] = 'VRFBLD'

        post_dict['CNT'] = 'STDLOG DONE'


    elif authhit == 4:

        auth_table.at[i,'SID'] == 'N'
        auth_table.at[i,'RDP'] == 'N'
        auth_table.at[i,'MNT'] == 'N'
        auth_table.at[i,'MSG'] == 'N'



        post_dict['MSG'] = 'BLDERR'

        post_dict['CNT'] = 'STDLOG DONE'


    return post_dict


print('\n NPIA_remote is listening on port '+str(port)+'-----\n')
app.run(host='0.0.0.0',port=port)