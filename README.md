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

  ## Start by Code


1. Set up the environment
```bash
unzip node-deploy.zip
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

3. Start the script

```bash

# attack-1 and attack-2
bash -x ./bsc_cluster.sh reset # will reset the cluster and start


# attack-3
export DELAY_INTERVAL_MS=25 && bash -x ./bsc_cluster.sh reset # DELAY_INTERVAL_MS can 25，50.75


```

4. Start the monitoring script

```bash
cd query && go run main.go --node=21
```

## Start by Docker

need docker 

```bash

# attack-1  

touch 1.txt && docker run -it --rm \
  -v ./1.txt:/app/query/21.txt \
  erick785/bsc-attack-1:latest

# attack-2
touch 2.txt && docker run -it --rm \
  -v ./2.txt:/app/query/21.txt \
  erick785/bsc-attack-2:latest

# attack-bootnode-3 ,DELAY_INTERVAL_MS can 25,50,75
touch 3.txt && docker run -it --rm \
  -v ./3.txt:/app/query/21.txt \
  -e DELAY_INTERVAL_MS=25 \
  erick785/bsc-attack-3-bootnode:latest

# attack-staicnode-3 ,DELAY_INTERVAL_MS can 25,50,75
touch 3.txt && docker run -it --rm \
  -v ./3.txt:/app/query/21.txt \
  -e DELAY_INTERVAL_MS=25 \
  erick785/bsc-attack-3-staicnode:latest

```


## Indicators of Experiment Success

After running the monitoring script, the following logs indicate a successful startup:
```
0,0,0,false
1,0,1,false
2,0,2,false
3,0,3,false
4,0,4,false
5,0,5,false
6,0,6,false
7,0,7,false
8,0,8,false
9,0,9,false
10,0,10,false
11,0,11,false
12,0,12,false
13,0,13,false
...
```
The first and third numbers represent the latest block height, the second represents the FinalizedBlock height, and the fourth indicates whether there is a vote attestation in the block header.
All experiments start at block height 250.

1. attack-1

After the block height reaches 250, if the FinalizedBlock height continues to increase, the experiment is considered successful.

```
...
247,245,247,true
248,246,248,true
249,247,249,true
250,248,250,true
251,249,251,true
252,249,252,false
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
270,268,270,true
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

2. attack-2

When the block height is above 250, if the FinalizedBlock height remains constant, the experiment is considered successful.
```
...
248,246,248,true
249,247,249,true
250,248,250,true
251,249,251,true
252,249,252,false
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

3. attack-3

After the block height reaches 250, if the FinalizedBlock height continues to increase, the experiment is considered successful。

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
450,429,450,false
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
476,474,476,true
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



