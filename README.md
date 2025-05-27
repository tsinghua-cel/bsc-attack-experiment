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

  - attck 1 code : ./code/attack-1-code.zip
  - attck 2 code : ./code/attack-2-code.zip
  - attck 3 code : ./code/attack-3-bootnode-code.zip
  - attck 3 code : ./code/attack-3-staicnode-code.zip


  Nodes deployment script: https://github.com/bnb-chain/node-deploy.git

## Installation
Before proceeding to the next steps, please ensure that the following packages and softwares are well installed in your local machine: 
- Ubuntu 20.04/22.04
- nodejs: 18.20.2 
- npm: 6.14.6
- go: 1.18+
- foundry
- python3 3.12+
- poetry
- jq
- docker

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
```

4. compile the geth binary, and place it in the node-deploy/bin/ folder
```bash
unzip ./code/attack-1-code.zip 
cd attack-1-code && make geth
```

## Launch attack simulation

Description of attack types:

- Attack 1: Basic Network Attack Simulation
- Attack 2: Enhanced Network Attack Simulation
- Attack 3: Latency-based network attack simulation (supports both bootnode and full connetion modes)

1. Start attack

```bash
# attack 1 and attack 2
bash -x ./bsc_cluster.sh reset # will reset the cluster and start

# attack-3
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

### attack 3 (Bootnode mode)
```bash
touch 3.txt && docker run -it --rm \
  -v ./3.txt:/app/query/21.txt \
  -e DELAY_INTERVAL_MS=25 \
  erick785/bsc-attack-3-bootnode:latest
```

### attack 3 (Full connetion mode)
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
| 3                  | Same as 1                      | Same as 1.    |
| 4                  | Attestation of block header                      | `true` means that the block received a vote attestation, `false` means that it did not.    |

 > \* **All attacks start at block height 250** *

## Criteria for determining success for each type of attack
### attack 1

Success Conditions:
- When the block height ≥ 250, the FinalizedBlock height continues to grow, indicating that the _Fast Finality_(FF) mechanism advances normally despite the attack.log

logs:
```
...
247,245,247,true
248,246,248,true
249,247,249,true
250,248,250,true # Latest block height is 250. Starting attack.
251,249,251,true # Finalized block height growth to 249
252,249,252,false # Finalize the block to stop growing
253,249,253,false
254,249,254,true
255,249,255,false
256,249,256,false
257,249,257,true
258,249,258,false
259,249,259,false
260,249,260,true
261,249,261,false
262,249,262,false
263,249,263,false
264,249,264,false
265,249,265,false
266,249,266,false
267,249,267,false
268,249,268,false
269,249,269,true
270,268,270,true # Finalized blocks catch up to 268, indicating a gradual recovery of consensus 
271,269,271,true
272,270,272,true
273,270,273,false
274,270,274,false
275,270,275,true
276,270,276,false
277,270,277,false
278,270,278,true
279,270,279,false
280,270,280,false
281,270,281,true
282,270,282,false
283,270,283,false
284,270,284,true
286,270,286,false
287,270,287,true
289,270,289,false
290,270,290,true
...
```

### attack 2

Success Conditions:
- When the block height ≥ 250, the FinalizedBlock always stops at the pre-attack level (e.g., 249), indicating a complete failure of the FF mechanism.

logs:
```
...
248,246,248,true
249,247,249,true
250,248,250,true # The last normal finalized block height before the attack was 248
251,249,251,true # After the attack started, the finalized block was briefly boosted to 249
252,249,252,false # Subsequent blocks finalize block height stagnation at 249
253,249,253,false
254,249,254,false
255,249,255,false
256,249,256,false
257,249,257,true
258,249,258,false
259,249,259,false
260,249,260,false
261,249,261,false
262,249,262,false
263,249,263,false
264,249,264,false
265,249,265,false
266,249,266,false
267,249,267,false
268,249,268,false
269,249,269,false
270,249,270,false
271,249,271,false
272,249,272,false
273,249,273,false
274,249,274,false
275,249,275,false
276,249,276,false
278,249,278,false
279,249,279,false
280,249,280,false
281,249,281,false
282,249,282,false
283,249,283,false
284,249,284,false
285,249,285,false
286,249,286,false
287,249,287,false
...

```

### attack-3

Success Conditions:
- The delayed growth of the FinalizedBlock height when the block height is ≥ 250 indicates that the attack affects the consensus through the delay strategy but does not completely destroy it.

logs:
```
...
432,429,432,false
433,429,433,true
434,429,434,false
435,429,435,false
436,429,436,false
437,429,437,false
438,429,438,false
439,429,439,false
440,429,440,false
441,429,441,false
442,429,442,false
443,429,443,false
444,429,444,false
445,429,445,false
446,429,446,false
447,429,447,false
448,429,448,false
449,429,449,false
450,429,450,false # At block 450, the finalized block height remains 429 (21 blocks delayed)  
451,429,451,false
452,429,452,false
453,429,453,false
454,429,454,false
455,429,455,false
456,429,456,false
457,429,457,false
458,429,458,false
459,429,459,false
460,429,460,false
461,429,461,false
462,429,462,false
463,429,463,false
464,429,464,false
465,429,465,false
466,429,466,false
467,429,467,false
468,429,468,false
469,429,469,false
470,429,470,false
471,429,471,false
472,429,472,false
473,429,473,true
474,429,474,false
475,429,475,true
476,474,476,true # After a delay, the `finalized block` starts catching up (`finalized block` height is the height of the `latest block` - 2) 
477,474,477,false
478,474,478,true
479,477,479,true
480,478,480,true
481,479,481,true
482,480,482,true
483,480,483,false
484,480,484,true
485,483,485,true
486,483,486,false
487,483,487,true
488,486,488,true
489,486,489,false
490,486,490,true
...
```
## Contribution
- For help, please submit an issue to the project repository.
