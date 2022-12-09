import subprocess
from pathlib import Path
import os
import pandas as pd
from dotenv import load_dotenv, find_dotenv
import sys
import json
import yaml
from yaml.loader import SafeLoader


stdlog = ''

targetdirname = 'target'
basedir = Path(__file__).resolve().parent
targetpath = os.path.join(basedir, targetdirname)

img_prefix = 'target_'

load_dotenv(find_dotenv())

REG_URL = os.environ.get("REG_URL")
DOCKER_ID = os.environ.get("DOCKER_ID")
DOCKER_PW = os.environ.get("DOCKER_PW")
ANSWER = os.environ.get("ANSWER")




record_id = int(sys.argv[1])

try :
    
    dockercomposeup = subprocess.run(['docker-compose','up','-d'], cwd=targetpath)
    stdlog += 'CMD : docker-compose up@\n'+str(dockercomposeup.returncode)+'\n'

    dockercomposedown = subprocess.run(['docker-compose','down'], cwd=targetpath)
    stdlog += 'CMD : docker-compose down@\n'+str(dockercomposedown.returncode)+'\n'

    print('new image(s) have been built successfully \n')


    dockerlogout = subprocess.run(['docker','logout'],cwd=targetpath)
    stdlog += 'CMD : docker login@\n'+str(dockerlogout.returncode)+'\n'

    dockerlogin = subprocess.run(['docker','login','-u',DOCKER_ID,'-p',DOCKER_PW],cwd=targetpath)
    stdlog += 'CMD : docker login@\n'+str(dockerlogin.returncode)+'\n'


    compose_file_path = os.path.join(targetpath,'docker-compose.yaml')

    readf = open(compose_file_path,'r',encoding='utf8')

    loaded_yaml = yaml.load_all(readf,Loader=SafeLoader)

    list_yaml = list(loaded_yaml)

    doc_yaml = list_yaml[0]

    img_list = []

    img_list = list(doc_yaml["services"].keys())

    for i in range(len(img_list)):

        img_list[i] = img_prefix + img_list[i]


    for i in range(len(img_list)):
        dockertag = subprocess.run(['docker','tag', img_list[i] , REG_URL + '/'+img_list[i] ],cwd=targetpath)
        stdlog += 'CMD : docker tag@\n'+str(dockertag.returncode)+'\n'
        
    for i in range(len(img_list)):
        dockerpush = subprocess.run(['docker','push', REG_URL + '/'+ img_list[i]],cwd=targetpath)
        stdlog += 'CMD : docker push@\n'+str(dockerpush.returncode)+'\n'


    print('new image(s) have been pushed successfully \n')



    dockerprune = subprocess.run(['docker','system','prune','-a'],input=ANSWER.encode(),cwd=targetpath)
    stdlog += 'CMD : docker system prune@\n'+str(dockerprune.returncode)+'\n'

    print('new image(s) have been cleared successfully \n')


except:

    auth_table = pd.read_csv('auth_table.csv',sep=',',skipinitialspace=True)

    auth_table.at[record_id,'MSG'] = 'BLDERR'

    auth_table.to_csv('auth_table.csv',index=False)


else:


    auth_table = pd.read_csv('auth_table.csv',sep=',',skipinitialspace=True)

    auth_table.at[record_id,'MSG'] = 'VRFBLD'

    auth_table.to_csv('auth_table.csv',index=False)