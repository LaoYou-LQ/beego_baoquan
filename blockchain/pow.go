package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const DIFFFICULTY = 12


type ProofOfwork struct {
	Target *big.Int
	Block  Block
}

//实例化一个pow算法实例
func NePow(block Block) ProofOfwork {

	target := big.NewInt(1) //初始值
	target.Lsh(target, 256-DIFFFICULTY)
	pow := ProofOfwork{
		Target: target,
		Block:  block,
	}
	return pow

}

//pow算法计算，寻找符合条件的nonce值

func (p ProofOfwork) Run() ([]byte ,int64) {
	var count =0
	var nonce int64
	var block256Hash []byte
	bigBlock := new(big.Int)
	for {//循环获取缓冲区的数据，将获取到的数据加密
		block := p.Block
		heightBytes, _ := util.IntToBytes(block.Height)
		timeBytes, _ := util.IntToBytes(block.TimeStamp)
		version := util.StringToBytes(block.Version)
		nonceBytes, _ := util.IntToBytes(nonce)
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			version,
			nonceBytes,
		}, []byte{})
		//将拼接的byte字节sha256加密
		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash = sha256Hash.Sum(nil)

		//因为目标值是big.int数值，int类型无法和big.int比较，所以将加密的数据变成big.int类型
		bigBlock = bigBlock.SetBytes(block256Hash)

		if p.Target.Cmp(bigBlock) ==1 {
			break //满足条件退出循环
		}

		nonce++ //不满足条件，nonce值+1，继续循环
		count++
	}
	fmt.Println(count)
	return block256Hash, nonce
}
