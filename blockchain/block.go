package blockchain

import (
	"bytes"
	"encoding/gob"
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
//Genesis 起源
func CreateGenesisBlock() Block {
	block := NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return block
}

func NewBlock(height int64, data []byte, prevHash []byte) Block {
	//1构建一个block实例，用于生产区块
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}

	//2为新生成的block，寻找nonce值
	pow := NePow(block)
	blockHash, nonce := pow.Run()

	//3将block的nonce设置为找到的合适的nonce数
	block.Nonce = nonce
	block.Hash = blockHash
	/*
		heightBytes, _ := util.IntToBytes(block.Height)
		timeBytes, _ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		//nonce值
		nonceBytes,_:=util.IntToBytes(block.Nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			data,
			prevHash,
			versionBytes,
			nonceBytes,
		}, []byte{})


	//设置第七个字段
	block.Hash = util.SHA256Hash(blockBytes)
	 */
	return block
}
/*
区块的序列化
 */
func (bl Block) Serialize() ([]byte,error) {
	buff:=new(bytes.Buffer)
	err:=gob.NewEncoder(buff).Encode(bl)
	if err!=nil {
		return nil, err
	}
	return buff.Bytes(),nil
}
/*
区块的反序列化

 */
func DeSerialize(data []byte) (*Block,error) {
	var block Block
	err:=gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err!=nil {
		return nil,err
	}
	return &block,nil
}