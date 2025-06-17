# bsc-attack-experiment

## Introduction
| **Output Results (TXT)**               | **Chain Data (ZIP)**                              | **Description**                                                                 |
|----------------------------------------|-------------------------------------------------|---------------------------------------------------------------------------------|
| `attack-1-reward.txt`                  | `node-attack-1-reward.zip`                      | Reward details for Attack 1.                                                   |
| `attack-1.txt`                         | `node-attack-1.zip`                             | Details and logs for Attack 1.                                                  |
| `attack-2-reward.txt`                  | `node-attack-2-reward.zip`                      | Reward details for Attack 2.                                                   |
| `attack-2.txt`                         | `node-attack-2.zip`                             | Details and logs for Attack 2.                                                  |
| `attack-3-bootnode-25.txt`             | `node-attack-3-bootnode-25.zip`                 | Details and logs for Attack 3: using bootnode-based connection with a 25ms delay.                  |
| `attack-3-bootnode-50-reward.txt`      | `node-attack-3-bootnode-50-reward.zip`          | Reward details for Attack 3: using bootnode-based connection with a 50ms delay.                   |
| `attack-3-bootnode-50.txt`             | `node-attack-3-bootnode-50.zip`                 | Details and logs for Attack 3: using bootnode-based connection with a 50ms delay.                  |
| `attack-3-bootnode-75.txt`             | `node-attack-3-bootnode-75.zip`                 | Details and logs for Attack 3: using bootnode-based connection with a 75ms delay.                  |
| `attack-3-staticnode-25.txt`           | `node-attack-3-staticnode-25.zip`               | Details and logs for Attack 3: using full connection with a 25ms delay.                |
| `attack-3-staticnode-50-reward.txt`    | `node-attack-3-staticnode-50-reward.zip`        | Reward details for Attack 3: using full connection with a 50ms delay.                 |
| `attack-3-staticnode-50.txt`           | `node-attack-3-staticnode-50.zip`               | Details and logs for Attack 3: using full connection with a 50ms delay.                |
| `attack-3-staticnode-75.txt`           | `node-attack-3-staticnode-75.zip`               | Details and logs for Attack 3: using full connection with a 75ms delay.                |
| `normal-reward.txt`                    | `node-normal-reward.zip`                        | Reward details for benchmark.                                           |


## Attack Code Repository

  Bsc repository (based on v1.4.16): https://github.com/bnb-chain/bsc/tree/v1.4.16

  - attack 1 code : ./code/attack-1-code.zip
  - attack 2 code : ./code/attack-2-code.zip
  - attack 3 code : ./code/attack-3-bootnode-code.zip
  - attack 3 code : ./code/attack-3-staicnode-code.zip


  Nodes deployment script: https://github.com/bnb-chain/node-deploy.git

## Installation
Before proceeding to the next step, please ensure that the following packages and software are well installed in your local computer: 
- Ubuntu 20.04/22.04
- nodejs: 18.20.2 
- npm: 6.14.6
- go: 1.18+
- python3: 3.12+
- docker: 27.5.1
- foundry: 1.1.0
- poety: 2.0.0
- jq: 1.7

## Start

1. Unzip and enter the project directory
```bash
unzip node-deploy.zip
cd node-deploy
```
2. Create and activate a virtual environment
```
# Create the virtual environment (if the venv package is not installed)
python3 -m venv path/to/venv

# Create virtual environments
apt install python3.12-venv

# Activate the virtual environment
source path/to/venv/bin/activate
```
3. Install dependencies
```
pip3 install -r requirements.txt
chmod +x install-dev.sh
sudo ./install-dev.sh
```

4. compile the geth binary, and place it in the node-deploy/bin/ folder
```bash
unzip ./code/attack-1-code.zip 
cd attack-1-code && make geth
unzip node-deploy.zip
mv attack-1-code/build/bin/geth node-deploy/bin/geth
```

## Launch attack simulation

Description of attack types:

- attack 1: Basic network attack simulation
- attack 2: Enhanced network attack simulation
- attack 3: The latency-based network attack simulation supports both bootnode and full connection modes

1. Start attack

```bash
# attack 1 and attack 2
bash -x ./bsc_cluster.sh reset # will reset the cluster and start

# attack 3
# Set delay to 25ms (DELAY_INTERVAL_MS can be adjusted to 25, 50, 75)
export DELAY_INTERVAL_MS=25 && bash -x ./bsc_cluster.sh reset
```

2. Start the monitoring script

```bash
cd query && go run main.go --node=21
```
> Notice: The monitoring data will be exported to the query/21.txt file.

## Start by Docker

Preconditions:
- Installed Docker Engine

### attack 1  
```bash
touch 1.txt && docker run -it --rm \
  -v ./1.txt:/app/query/21.txt \
  erick785/bsc-attack-1:latest
```
### attack 2
```bash
touch 2.txt && docker run -it --rm \
  -v ./2.txt:/app/query/21.txt \
  erick785/bsc-attack-2:latest
```

### attack 3 (Bootnode connection mode)
```bash
touch 3.txt && docker run -it --rm \
  -v ./3.txt:/app/query/21.txt \
  -e DELAY_INTERVAL_MS=25 \
  erick785/bsc-attack-3-bootnode:latest
```

### attack 3 (Full connection mode)
```bash
touch 3.txt && docker run -it --rm \
  -v ./3.txt:/app/query/21.txt \
  -e DELAY_INTERVAL_MS=25 \
  erick785/bsc-attack-3-staicnode:latest
```

## Description of indicators of experimental success

Data field definitions
| **Field Order**               | **Field Name**                              | **Description**                                                                 |
|----------------------------------------|-------------------------------------------------|---------------------------------------------------------------------------------|
| 1                  | The latest block height                      | The number of the latest block that has been generated by the current node (`LatestBlock`).    |
| 2                  | Finalized block height                      | Block number that has been finalized (`FinalizedBlock`).    |
| 3                  | Attestation of block header                      | `true` means that the block received a vote attestation, `false` means that it did not.    |

 > \* **All attacks start at block height 250** *

## Criteria for successful implementation of each type of attack
### attack 1

Success Conditions:
- When the block height ≥ 250, the FinalizedBlock height continues to grow, indicating that the _Fast Finality_(FF) mechanism advances normally despite the attack.log

logs:
```
...
247,245,true
248,246,true
249,247,true
250,248,true # Latest block height is 250. Starting attack.
251,248,false # Finalize the block to stop growing
252,248,false
253,248,false
...
266,248,false
267,248,false
268,248,true
269,267,true
270,268,true # Finalized blocks catch up to 268, indicating a gradual recovery of consensus 
271,269,true
272,269,false
273,269,false
274,269,false
...
289,270,false
290,288,true
291,289,true
292,290,true
293,290,false
...
```

### attack 2

Success Conditions:
- When the block height ≥ 250, the FinalizedBlock always stops at the pre-attack level (e.g., 248), indicating a complete failure of the FF mechanism.

logs:
```
...
248,246,true
249,247,true
250,248,true # The last normal finalized block height before the attack was 248
251,248,false # After the attack started, the finalized block was briefly boosted to 248
252,248,false # Subsequent blocks finalize block height stagnation at 248
253,248,false
254,248,false
255,248,false
256,248,false
257,248,true
...
274,248,false
275,248,false
276,248,false
278,248,false
279,248,false
280,248,false
281,248,false
282,248,false
283,248,false
284,248,false
285,248,false
286,248,false
287,248,false
...

```

### attack-3

Success Conditions:
- The delayed growth of the FinalizedBlock height when the block height is ≥ 250 indicates that the attack affects the consensus through the delay strategy but does not completely destroy it.

logs:
```
...
438,429,false
439,429,false
440,429,false
441,429,false
442,429,false
443,429,false
444,429,false
445,429,false
446,429,false
447,429,false
448,429,false
449,429,false
450,429,false # At block 450, the finalized block height remains 429 (21 blocks delayed)  
451,429,false
452,429,false
453,429,false
454,429,false
455,429,false
456,429,false
457,429,false
458,429,false
459,429,false
460,429,false
461,429,false
462,429,false
463,429,false
464,429,false
465,429,false
466,429,false
467,429,false
468,429,false
469,429,false
470,429,false
471,429,false
472,429,false
473,429,true
474,429,false
475,429,true
476,474,true # After a delay, the `finalized block` starts catching up (`finalized block` height is the height of the `latest block` - 2) 
477,474,false
478,474,true
479,477,true
480,478,true
481,479,true
482,480,true
483,480,false
484,480,true
485,483,true
486,483,false
487,483,true
488,486,true
489,486,false
490,486,true
...
```
## Contribution
- For help, please submit an issue to the project repository.
