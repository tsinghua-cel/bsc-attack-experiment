# bsc-attack-experiment

# bsc-experiment-program

## Introduction
| **Output Results (TXT)**               | **Chain Data (ZIP)**                              | **Description**                                                                 |
|----------------------------------------|-------------------------------------------------|---------------------------------------------------------------------------------|
| `attack-1-reward.txt`                  | `node-attack-1-reward.zip`                      | Reward details for Attack 1.                                                   |
| `attack-1.txt`                         | `node-attack-1.zip`                             | Logs or details for Attack 1.                                                  |
| `attack-2-reward.txt`                  | `node-attack-2-reward.zip`                      | Reward details for Attack 2.                                                   |
| `attack-2.txt`                         | `node-attack-2.zip`                             | Logs or details for Attack 2.                                                  |
| `attack-3-bootnode-25.txt`             | `node-attack-3-bootnode-25.zip`                 | Logs or details for Attack 3 with delay 25ms bootnode participation.                  |
| `attack-3-bootnode-50-reward.txt`      | `node-attack-3-bootnode-50-reward.zip`          | Reward details for Attack 3 with delay 50ms bootnode participation.                   |
| `attack-3-bootnode-50.txt`             | `node-attack-3-bootnode-50.zip`                 | Logs or details for Attack 3 with delay 50ms bootnode participation.                  |
| `attack-3-bootnode-75.txt`             | `node-attack-3-bootnode-75.zip`                 | Logs or details for Attack 3 with delay 75ms bootnode participation.                  |
| `attack-3-staticnode-25.txt`           | `node-attack-3-staticnode-25.zip`               | Logs or details for Attack 3 with delay 25ms staticnode participation.                |
| `attack-3-staticnode-50-reward.txt`    | `node-attack-3-staticnode-50-reward.zip`        | Reward details for Attack 3 with delay 50ms staticnode participation.                 |
| `attack-3-staticnode-50.txt`           | `node-attack-3-staticnode-50.zip`               | Logs or details for Attack 3 with delay 50ms staticnode participation.                |
| `attack-3-staticnode-75.txt`           | `node-attack-3-staticnode-75.zip`               | Logs or details for Attack 3 delay 75ms staticnode participation.                |
| `normal-reward.txt`                    | `node-normal-reward.zip`                        | Reward details for normal operation.                                           |


## Attack Code Repository

  bsc repository (based on v1.4.16):https://github.com/bnb-chain/bsc.git

  - attck 1 code : ./code/attack-1-code.zip
  - attck 2 code : ./code/attack-2-code.zip
  - attck 3 code : ./code/attack-3-bootnode-code.zip
  - attck 3 code : ./code/attack-3-staicnode-code.zip


  Nodes deployment scrip https://github.com/bnb-chain/node-deploy.git

  ## Installation
Before proceeding to the next steps, please ensure that the following packages and softwares are well installed in your local machine: 
- nodejs: 18.20.2 
- npm: 6.14.6
- go: 1.18+
- foundry
- python3 3.12+
- poetry
- jq

  ## Quick Start


1. Set up the environment
```bash
git clone https://github.com/bnb-chain/node-deploy.git
cd node-deploy
python3 -m venv path/to/venv
apt install python3.12-venv
source path/to/venv/bin/activate
pip3 install -r requirements.txt
```

2. compile the geth binary, and place it in the node-deploy/bin/ folder
```bash
unzip ./code/attack-1-code.zip 
cd attack-1-code && make geth
```

4. Configure the cluster
```
  You can configure the cluster by modifying the following files:
   - `config.toml`
   - `genesis/genesis-template.json`
   - `.env`
```

5. Start the script

```bash
bash -x ./bsc_cluster.sh reset # will reset the cluster and start
bash -x ./bsc_cluster.sh stop [vidx] # Stops the cluster
bash -x ./bsc_cluster.sh start [vidx] # only start the cluster
bash -x ./bsc_cluster.sh restart [vidx] # start the cluster after stopping it
```

6. Start the monitoring script

```bash
cd query && go run main.go --node=21
```