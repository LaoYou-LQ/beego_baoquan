package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"time"
)

type Block struct {
	Height    int64  //区块高度
	TimeStamp int64  //时间戳
	Hash      []byte //区块hash
	Data      []byte //数据
	PrevHash  []byte //上个区块的Hash
	Version   string //版本号
	Nonce     int64
}

func NewBlock(height int64, data []byte, prevHash []byte) Block {
	block := Block{
		Height:    height + 1,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
	heightBytes, _ := util.IntToBytes(block.Height)
	timeBytes, _ := util.IntToBytes(block.TimeStamp)
	versionBytes := util.StringToBytes(block.Version)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeBytes,
		data,
		prevHash,
		versionBytes,
	}, []byte{})
	block.Hash = util.SHA256Hash(blockBytes)
	return block
}
