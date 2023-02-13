# CloudAgentCLI-NPIA


## How to run

You need

- go 1.19+

- ADM/origin.json

### ADM/origin.json 


```json
  {
    "KCFG_PATH":"",
    "MAIN_NS":"",
    "RECORDS":[
      {
        "NS":"",
        "REPO_ADDR":"",
        "REG_ADDR":""
      }
    ],
    "REPOS": [
      {
        "REPO_ADDR": "",
        "REPO_ID": "",
        "REPO_PW": ""
      }
    ],
    "REGS": [
      {
        "REG_ADDR": "",
        "REG_ID": "",
        "REG_PW": ""
      }
    ]
  }

```

### Build

At nopainctl_source directory, run below

```bash
go build

mv nopainctl ../

```

then 

```bash

cd ../

sudo ./nopainctl run  #for running nopainctl main process
sudo ./nopainctl origin  #for setting up ADM/origin.json

```

## Important Logs

- [2023-02-09] Created main CLI : nopainctl
