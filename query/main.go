package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

var startNumber int
var endNumber int
var nodeId int
var nodes int

var validators map[string]int
var noderpcs map[int]string
var sortValidatorAddrs []common.Address

func init() {
	// startNumber
	flag.IntVar(&startNumber, "start", 0, "The start block number")
	// endNumber
	flag.IntVar(&endNumber, "end", 0, "The end block number")
	// nodeId
	flag.IntVar(&nodeId, "node", -1, "The node id")
	// nodes
	flag.IntVar(&nodes, "num", 21, "The number of nodes")
	flag.Parse()
	if nodes == 7 {
		validators = map[string]int{
			"0x3ad55d1d552cc55dee90c0faf0335383b2e6c5ce": 1, //8548
			"0x5e2a531a825d8b61bcc305a35a7433e9a8920f0f": 2, //8547 2
			"0x5fda3ff6ea581ea7a5a9c2cb310b13c2126b4e8b": 3, //8551
			"0xbbd1acc20bd8304309d31d8fd235210d0efc049d": 4, //8546
			"0xbcdd0d2cda5f6423e57b6a4dcd75decbe31aecf0": 5, //8545 5
			"0xf7698afa5461438ff438c2322d6d29a5f7abdffd": 6, //8550
			"0xfe02c8ff2374583c47b1d62fdf3e1b72c20ebe29": 7, //8549

		}
	} else {
		// 1 ：1 4 7 10 13 16
		// node1: 13 14 16 10 9 11

		// 2 ：2 3 5 6 8 9 11 12
		// node2: 20 3 17 8 2 6 19 15

		// 3 ：14 15 17 18 19 20 21
		// node3: 1 0 12 7 18 5 4

		validators = map[string]int{
			"0x20be3a44b2ae6be29acf84ed63afe60b09179cdc": 1,  //8558  1 13
			"0x297e5ebba75bbb67de013eb3d319dd0a2a9861e9": 2,  //8565  2 20
			"0x3ad55d1d552cc55dee90c0faf0335383b2e6c5ce": 3,  //8548  3 3
			"0x50b947c8643c7694037b29545fbc423951e28442": 4,  //8559  4 14
			"0x511aa4d222618f8698feaab811023ca4e8bebfe5": 5,  //8562  5 17
			"0x51cb3d0f6b77ef8317b31f4aaeaa75e4cff3cca7": 6,  //8553  6 8
			"0x5a7ae634876fb264f97eacc24a9261005e9bc39a": 7,  //8561  7 16
			"0x5e2a531a825d8b61bcc305a35a7433e9a8920f0f": 8,  //8547  8	2
			"0x5fda3ff6ea581ea7a5a9c2cb310b13c2126b4e8b": 9,  //8551  9 6
			"0x6c73f4f3295f83ce342e4a82e8a50d218442451b": 10, //8555  10 10
			"0x9b50a300da0cd7e036ec2cc12418756ec07004bd": 11, //8564  11 19
			"0xa8938f397823afcaa252bb7df137d39396456983": 12, //8560  12 15
			"0xabb28e397ae478366271806b4851d81a678e404b": 13, //8554  13 9
			"0xbbd1acc20bd8304309d31d8fd235210d0efc049d": 14, //8546  14 1
			"0xbcdd0d2cda5f6423e57b6a4dcd75decbe31aecf0": 15, //8545  15 0
			"0xc12cf70a667d541a33bd51c623f8a7024ed8c2fe": 16, //8556  16 11
			"0xd2d3139575c2824d793d1664c2e1aaeecade11c0": 17, //8557  17 12
			"0xd30d79639bc9c4ed71031bce28216862b80f4b6b": 18, //8552  18 7
			"0xe9693a85e563485da999b7d378d60483e89caa0e": 19, //8563  19 18
			"0xf7698afa5461438ff438c2322d6d29a5f7abdffd": 20, //8550  20 5
			"0xfe02c8ff2374583c47b1d62fdf3e1b72c20ebe29": 21, //8549  21 4
		}

	}

	if nodes == 7 {

		noderpcs = map[int]string{
			1: "http://0.0.0.0:8548",
			2: "http://0.0.0.0:8547",
			3: "http://0.0.0.0:8551",
			4: "http://0.0.0.0:8546",
			5: "http://0.0.0.0:8545",
			6: "http://0.0.0.0:8550",
			7: "http://0.0.0.0:8549",
		}
	} else {

		noderpcs = map[int]string{

			1:  "http://0.0.0.0:8558",
			2:  "http://0.0.0.0:8565",
			3:  "http://0.0.0.0:8548",
			4:  "http://0.0.0.0:8559",
			5:  "http://0.0.0.0:8562",
			6:  "http://0.0.0.0:8553",
			7:  "http://0.0.0.0:8561",
			8:  "http://0.0.0.0:8547",
			9:  "http://0.0.0.0:8551",
			10: "http://0.0.0.0:8555",
			11: "http://0.0.0.0:8564",
			12: "http://0.0.0.0:8560",
			13: "http://0.0.0.0:8554",
			14: "http://0.0.0.0:8546",
			15: "http://0.0.0.0:8545",
			16: "http://0.0.0.0:8556",
			17: "http://0.0.0.0:8557",
			18: "http://0.0.0.0:8552",
			19: "http://0.0.0.0:8563",
			20: "http://0.0.0.0:8550",
			21: "http://0.0.0.0:8549",
		}
	}

	sortValidatorAddrs = make([]common.Address, 0, len(validators))
	for k := range validators {
		sortValidatorAddrs = append(sortValidatorAddrs, common.HexToAddress(k))
	}

	sort.Sort(validatorsAscending(sortValidatorAddrs))

	for k, v := range sortValidatorAddrs {
		validators[strings.ToLower(v.Hex())] = k + 1
	}

	s, _ := json.MarshalIndent(validators, " ", "")
	fmt.Println(string(s))

	// for _, v := range validators {
	// 	noderpcs[v] = "http://0.0.0.0:" + strconv.Itoa(8545+(v-4))
	// }

}

type validatorsAscending []common.Address

func (s validatorsAscending) Len() int           { return len(s) }
func (s validatorsAscending) Less(i, j int) bool { return bytes.Compare(s[i][:], s[j][:]) < 0 }
func (s validatorsAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {

	if nodeId != -1 {
		nodeID := strconv.Itoa(nodeId)

		validatorAddr := common.Address{}

		for k, v := range validators {
			if v == nodeId {
				validatorAddr = common.HexToAddress(k)
			}
		}

		// sort validators
		fmt.Println(nodeID, validatorAddr, noderpcs[nodeId])

		scanner := NewBlockScanner(nodeID, validatorAddr, noderpcs[nodeId])
		go scanner.ScanLoop()
	}

	select {}

}

type BlockScanner struct {
	nodeID    string
	validator common.Address
	client    *rpc.Client
}

func NewBlockScanner(nodeID string, validator common.Address, url string) *BlockScanner {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return &BlockScanner{
		nodeID:    nodeID,
		validator: validator,
		client:    client,
	}
}

type Block struct {
	Number     hexutil.Big   `json:"number"`
	Hash       string        `json:"hash"`
	ParentHash string        `json:"parentHash"`
	Miner      string        `json:"miner"`
	Timestamp  hexutil.Big   `json:"timestamp"`
	Difficulty hexutil.Big   `json:"difficulty"`
	ExtraData  hexutil.Bytes `json:"extraData"`
}

type Resp struct {
	LatestBlock    *Block
	SafeBlock      *Block
	FinalizedBlock *Block
}

var attackCounts, total int64

func (r *Resp) String() string {
	return fmt.Sprintf("hash: %v,parenthash: %v, Latest: %v  miner: %v,should: %v, diff: %v, Finalized: %v, Attestation %v \n",

		r.LatestBlock.Hash, r.LatestBlock.ParentHash,
		r.LatestBlock.Number.ToInt(), validators[strings.ToLower(r.LatestBlock.Miner)],
		validators[strings.ToLower(inturnValidator(r.LatestBlock.Number.ToInt().Uint64()).Hex())], r.LatestBlock.Difficulty.ToInt(),
		r.FinalizedBlock.Number.ToInt(), getVoteAttestationFromHeader(r.LatestBlock.ExtraData, r.LatestBlock.Number.ToInt().Uint64(), 20))
}

func (r *Resp) WString() string {
	return fmt.Sprintf("%v,%v,%v,%v",
		r.LatestBlock.Number.ToInt(), r.FinalizedBlock.Number.ToInt(),
		r.LatestBlock.Number.ToInt(), getVoteAttestationFromHeader(r.LatestBlock.ExtraData, r.LatestBlock.Number.ToInt().Uint64(), 20))
}

const (
	extraVanity          = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal            = 65 // Fixed number of extra-data suffix bytes reserved for signer seal
	validatorNumberSize  = 1  // Fixed number of extra prefix bytes reserved for validator number after Luban
	validatorBytesLength = common.AddressLength + 48
	turnLengthSize       = 1 // Fixed number of extra-data suffix bytes reserved for turnLength

)

func getVoteAttestationFromHeader(extra []byte, number uint64, epoch uint64) bool {
	if len(extra) <= extraVanity+extraSeal {
		return false
	}

	var attestationBytes []byte
	if number%epoch != 0 {
		attestationBytes = extra[extraVanity : len(extra)-extraSeal]
	} else {
		num := int(extra[extraVanity])
		start := extraVanity + validatorNumberSize + num*validatorBytesLength
		start += turnLengthSize
		end := len(extra) - extraSeal
		if end <= start {
			return false
		}
		attestationBytes = extra[start:end]
	}

	return len(attestationBytes) > 0

}

func (b *BlockScanner) ScanLoop() {
	currentBlockNum := ""

	for {
		time.Sleep(1500 * time.Millisecond)
		latestblock := &Block{}
		if err := getBlock(b.client, latestblock, "latest"); err != nil {
			fmt.Println("getBlock error: ", err)
			continue
		}
		if currentBlockNum != latestblock.Number.String() {
			currentBlockNum = latestblock.Number.String()

			safeBlock := &Block{}
			getBlock(b.client, safeBlock, "safe")

			finalized := &Block{}
			getBlock(b.client, finalized, "finalized")

			resp := &Resp{
				LatestBlock:    latestblock,
				SafeBlock:      safeBlock,
				FinalizedBlock: finalized,
			}

			b.WriteBlockToFile(resp)
		}
	}

}

// inturnValidator returns the validator for the following block height.
func inturnValidator(number uint64) common.Address {
	offset := (number) / uint64(1) % uint64(len(sortValidatorAddrs))
	return sortValidatorAddrs[offset]
}

func (b *BlockScanner) WriteBlockToFile(resp *Resp) {
	// write to file

	fmt.Println(resp.String())

	fileName := fmt.Sprintf("%s.txt", b.nodeID)

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open or create file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(resp.WString() + "\n"); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func getBlock(client *rpc.Client, block *Block, blockParam string) error {
	err := client.CallContext(context.Background(), block, "eth_getBlockByNumber", blockParam, true)
	if err != nil {
		return err
	}
	return nil
}
