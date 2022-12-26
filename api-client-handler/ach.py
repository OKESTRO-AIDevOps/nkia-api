import requests
import json


SERVER_URL = 'http://localhost:7331'
SERVER_SUB_URL = 'http://localhost:7331/build'


def interact(acd):

    post_dict = {"ACD": "testpass", "SID": "N", "RDP": "N",
            "MNT": "Y", "MSG": "N", "CMD": "N", "CNT": "N"}

    trm = 0

    ns = 'default'

    post_dict["ACD"] = acd

    print('You\'re in Interact Stage, which includes deployment, QoS, lifecycle, auto-scaling management and resource monitoring.')
    print('You can do those things by typing a corresponding command into the prompt')
    print('Type [ list ] to check available commands.')
    print('Type [ trm ] to exit.')

    yn = input('Interact Start?  [ y | n ] : ')

    if yn.lower() not in ['no','n','yes','y'] :

        print('Option Not Available, bye.')

        return
    
    elif yn.lower() == 'no' or yn.lower() == 'n':

        print('Exit, bye')

        return    


    res = requests.post(SERVER_URL, json=post_dict, timeout=300)

    print('RET--------------------')
    print(res.text)
    print('-----------------------\n')

    post_dict = json.loads(res.text)


    if post_dict['ACD'] == 'N' :

        print('Auth Failed, check if you\'ve typed a correct authcode.')
        print('Exit, bye')
        return


    while trm == 0:

        print('Current Namespace : '+ns)

        command = input('CMD : ')

        if command == 'list':

            print('*COMMAND LIST*\n')
            print('a1. GET:PDS,[ namespace ]')
            print('----> gets pods in a namespace')
            print('a2. GET:SVC,[ namespace ]')
            print('----> gets services in a namespace')
            print('a3. GET:DPL,[ namespace ]')
            print('----> gets deployments in a namespace')
            print('a4. GET:NDS,all')
            print('----> gets all nodes of the target cluster')
            print('a5. GET:EVN,[ namespace ]')
            print('----> gets all events in a namespace')
            print('a6. GET:RSC,[ namespace ]')
            print('----> gets all resources in a namespace')
            print('a7. GET:NSC,all')
            print('----> gets all namespaces available of the target cluster')
            print('b1. APPLY:APP,[ namespace ]')
            print('----> deploys the apps predefined in a user\'s source code repository')
            print('b2. APPLY:QOS,[ namespace ],DPL,[ name ],high')
            print('----> modifies a deployment\'s QoS policy in a namespace to Guaranteed')
            print('b3. APPLY:QOS,[ namespace ],DPL,[ name ],middle')
            print('----> modifies a deployment\'s QoS policy in a namespace to Burstable')
            print('b4. APPLY:QOS,[ namespace ],DPL,[ name ],low')
            print('----> modifies a deployment\'s QoS policy in a namespace to Best-Effort')
            print('b5. APPLY:CRTNS,[ namespace ]')
            print('----> creates a namespace')
            print('b6. APPLY:USENS,[ namespace ]')
            print('----> uses a namespace')
            print('c1. LIFE:RESTART,[ namespace ],DPL,[ name ]')
            print('----> restarts a deployment in a namespace')
            print('c2. LIFE:ROLLBACK,[ namespace ],DPL,[ name ]')
            print('----> reverts a deployment in a namespace to a previous status')
            print('c3. LIFE:HISTORY,[ namespace ],DPL,[ name ]')
            print('----> gets revision history of a deployment in a namespace')
            print('c4. LIFE:KILL,[ namespace ],DPL,[ name ]')
            print('----> deletes a deployment in a namespace and a corresponding service')
            print('d1. PA:HORIZONTAL,[ namespace ],DPL,[ name ],max')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with maximum capacity allowed')
            print('d2. PA:HORIZONTAL,[ namespace ],DPL,[ name ],optimal')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with optimal capacity allowed')
            print('d3. PA:HORIZONTAL,[ namespace ],DPL,[ name ],min')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with minimum capacity allowed')
            print('d4. PA:VERTICAL,[ namespace ],DPL,[ name ],max')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with maximum capacity allowed')
            print('d5. PA:VERTICAL,[ namespace ],DPL,[ name ],optimal')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with optimal capacity allowed')
            print('d6. PA:VERTICAL,[ namespace ],DPL,[ name ],min')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with minimum capacity allowed')
            print('list')
            print('----> lists all available commands')
            print('trm')
            print('----> terminates session')

            print('\nUSAGE')

            print('CMD : [ code ]')
            print('ex) CMD : a7')
            print('----> this command will get all namespaces available without any following prompt')
            print('ex) CMD : c1')
            print('----> this command will restart a deployment in a namespace with following prompts for user-typed deployment name')
            print('----> hence a user must know beforehand the deployment name to target')

            continue

        elif command == 'trm':

            post_dict["CMD"] = 'TRM:all'

            res = requests.post(SERVER_URL, json=post_dict, timeout=300)

            print('RET--------------------')
            print(res.text)
            print('-----------------------\n')

            post_dict = json.loads(res.text)

            print('Session terminated.')
            print('Exit, bye')
            
            trm = 1

            continue


        if command == 'a1':

            print('a1. GET:PDS,[ namespace ]')
            print('----> gets pods in a namespace')



            cmd = 'GET:PDS,' + ns

            post_dict["CMD"] = cmd




        elif command == 'a2':   
             
            print('a2. GET:SVC,[ namespace ]')
            print('----> gets services in a namespace')

            


            cmd = 'GET:SVC,' + ns

            post_dict["CMD"] = cmd

        elif command == 'a3':    
            print('a3. GET:DPL,[ namespace ]')
            print('----> gets deployments in a namespace')



            cmd = 'GET:DPL,' + ns

            post_dict["CMD"] = cmd


        elif command == 'a4':    

            print('a4. GET:NDS,all')
            print('----> gets all nodes of the target cluster')

            cmd = 'GET:NDS,all'

            post_dict["CMD"] = cmd


        elif command == 'a5':    

            print('a5. GET:EVN,[ namespace ]')
            print('----> gets all events in a namespace')


            cmd = 'GET:EVN,' + ns

            post_dict["CMD"] = cmd
        
        elif command == 'a6':    

            print('a6. GET:RSC,[ namespace ]')
            print('----> gets all resources in a namespace')


            cmd = 'GET:RSC,' + ns

            post_dict["CMD"] = cmd

        elif command == 'a7':    

            print('a7. GET:NSC,all')
            print('----> gets all namespaces available of the target cluster')

            cmd = 'GET:NSC,all'

            post_dict["CMD"] = cmd


        elif command == 'b1':    

            print('b1. APPLY:APP,[ namespace ]')
            print('----> deploys the apps predefined in a user\'s source code repository')


            cmd = 'APPLY:APP,' + ns

            post_dict["CMD"] = cmd


        elif command == 'b2':    

            print('b2. APPLY:QOS,[ namespace ],DPL,[ name ],high')
            print('----> modifies a deployment\'s QoS policy in a namespace to Guaranteed')


            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'APPLY:QOS,'+ns+',DPL,'+nm+',high'

            post_dict["CMD"] = cmd


        elif command == 'b3':   

            print('b3. APPLY:QOS,[ namespace ],DPL,[ name ],middle')
            print('----> modifies a deployment\'s QoS policy in a namespace to Burtable') 

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'APPLY:QOS,'+ns+',DPL,'+nm+',middle'

            post_dict["CMD"] = cmd


        elif command == 'b4':    

            print('b4. APPLY:QOS,[ namespace ],DPL,[ name ],low')
            print('----> modifies a deployment\'s QoS policy in a namespace to Best-Effort')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'APPLY:QOS,'+ns+',DPL,'+nm+',low'

            post_dict["CMD"] = cmd

        elif command == 'b5':

            print('b5. APPLY:CRTNS,[ namespace ]')
            print('----> creates a namespace')

            nm  = input('namespace you want to create : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'APPLY:CRTNS,'+nm

            post_dict["CMD"] = cmd


        elif command == 'b6':

            print('b6. APPLY:USENS,[ namespace ]')
            print('----> uses a namespace')

            nm  = input('namespace you want to switch to : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'APPLY:USENS,'+nm

            post_dict["CMD"] = cmd

            ns = nm

            continue


        elif command == 'c1':    

            print('c1. LIFE:RESTART,[ namespace ],DPL,[ name ]')
            print('----> restarts a deployment in a namespace')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'LIFE:RESTART,'+ns+',DPL,'+nm

            post_dict["CMD"] = cmd



        elif command == 'c2':    
            print('c2. LIFE:ROLLBACK,[ namespace ],DPL,[ name ]')
            print('----> reverts a deployment in a namespace to a previous status')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'LIFE:ROLLBACK,'+ns+',DPL,'+nm

            post_dict["CMD"] = cmd


        elif command == 'c3':    

            print('c3. LIFE:HISTORY,[ namespace ],DPL,[ name ]')
            print('----> gets revision history of a deployment in a namespace')
            
            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'LIFE:HISTORY,'+ns+',DPL,'+nm

            post_dict["CMD"] = cmd

        elif command == 'c4':    

            print('c4. LIFE:KILL,[ namespace ],DPL,[ name ]')
            print('----> deletes a deployment in a namespace and a corresponding service')


            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'LIFE:KILL,'+ns+',DPL,'+nm

            post_dict["CMD"] = cmd


        elif command == 'd1':    

            print('d1. PA:HORIZONTAL,[ namespace ],DPL,[ name ],max')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with maximum capacity allowed')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:HORIZONTAL,'+ns+',DPL,'+nm+',max'

            post_dict["CMD"] = cmd

        elif command == 'd2':    

            print('d2. PA:HORIZONTAL,[ namespace ],DPL,[ name ],optimal')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with optimal capacity allowed')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:HORIZONTAL,'+ns+',DPL,'+nm+',optimal'

            post_dict["CMD"] = cmd

        elif command == 'd3':    

            print('d3. PA:HORIZONTAL,[ namespace ],DPL,[ name ],min')
            print('----> deploys HorizontalPodAutoscaler of a deployment in a namespace with minimum capacity allowed')

            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:HORIZONTAL,'+ns+',DPL,'+nm+',min'

            post_dict["CMD"] = cmd

        elif command == 'd4':    

            print('d4. PA:VERTICAL,[ namespace ],DPL,[ name ],max')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with maximum capacity allowed')
            
            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:VERTICAL,'+ns+',DPL,'+nm+',max'

            post_dict["CMD"] = cmd

        elif command == 'd5':    

            print('d5. PA:VERTICAL,[ namespace ],DPL,[ name ],optimal')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with optimal capacity allowed')


            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:VERTICAL,'+ns+',DPL,'+nm+',optimal'

            post_dict["CMD"] = cmd

        elif command == 'd6':    

            print('d6. PA:VERTICAL,[ namespace ],DPL,[ name ],min')
            print('----> deploys VerticalPodAutoscaler of a deployment in a namespace with minimum capacity allowed')


            nm  = input('deployment name : ')

            nm = str(nm)

            nm = nm.replace(' ','')

            cmd = 'PA:VERTICAL,'+ns+',DPL,'+nm+',min'

            post_dict["CMD"] = cmd

        else:

            print('Option not available')

            continue

        


        res = requests.post(SERVER_URL, json=post_dict, timeout=300)

        print('RET--------------------')
        print(res.text)
        print('-----------------------\n')

        post_dict = json.loads(res.text)


    return




def build(acd):

    post_dict = {"ACD": "testpass", "SID": "N", "RDP": "Y",
             "MNT": "N", "MSG": "N", "CMD": "N", "CNT": "N"}


    post_dict["ACD"] = acd

    vrfbld = 0


    res = requests.post(SERVER_URL, json=post_dict, timeout=300)

    print('RET--------------------')
    print(res.text)
    print('-----------------------\n')

    post_dict = json.loads(res.text)

    if post_dict['ACD'] == 'N' :

        print('Auth Failed, check if you\'ve typed a correct authcode.')
        print('Exit, bye')
        return
    

    print('You\'re in Build Stage, which includes Interact Stage with the cluster after building process is finished.')
    print('Hitting [ y ] means that you CANNOT abort the building and pushing process until it\'s finished.')
    print('You can choose whether you would start the deployment of the newly built apps or do other things after the images are pushed\n')
    yn = input('Build Start?  [ y | n ] : ')

    if yn.lower() not in ['no','n','yes','y'] :

        print('Option Not Available, bye.')

        return
    
    elif yn.lower() == 'no' or yn.lower() == 'n':

        print('Exit, bye')

        return    



    print('Build Started...')



    res = requests.post(SERVER_SUB_URL, json=post_dict, timeout=300)

    print('RET--------------------')
    print(res.text)
    print('-----------------------\n')

    post_dict = json.loads(res.text)


    print('Building in Process...')

    print('Now you can query the build plugin to check how it\'s going')


    while vrfbld == 0:


        y = input('Query build plugin? [ Enter (or whatever you want to hit) ] : ')

        res = requests.post(SERVER_SUB_URL, json=post_dict, timeout=300)

        print('RET--------------------')
        print(res.text)
        print('-----------------------\n')

        post_dict = json.loads(res.text)

        if post_dict['MSG'] == 'VRFBLD':

            print('RET--------------------')
            print(res.text)
            print('-----------------------\n')

            print('Build Successful!\n')

            vrfbld = 1

        elif post_dict['MSG'] == 'BLDERR':

            print('RET--------------------')
            print(res.text)
            print('-----------------------\n')

            print('Build Error Occured!\n')

            print('You can restart the process from the beginning')

            print('Exit, bye.')

            return


        


    print('Build Stage is completed. You can either deploy them and then interact with the cluster. \nOr exit this process')
    print('When you exit, you can still deploy the apps in Interact Stage.')
    yn = input('Do you want to move on to Interact Stage? [ y | n ] :')


    if yn.lower() not in ['no','n','yes','y'] :

        print('Option Not Available, bye.')

        return
    
    elif yn.lower() == 'no' or yn.lower() == 'n':

        print('Exit, bye')

        return    

    res = requests.post(SERVER_URL, json=post_dict, timeout=300)

    print('RET--------------------')
    print(res.text)
    print('-----------------------\n')

    post_dict = json.loads(res.text)



    interact(acd)








mnt_or_rdp = input('1. Interact \n2. Build \nType [ 1 | 2 ] : ')

mnt_or_rdp = str(mnt_or_rdp)


acd = input('Authcode : ')

acd = str(acd)

if mnt_or_rdp == '1':

    interact(acd)

elif mnt_or_rdp == '2':

    build(acd)

else:

    print('Option Not Available')
