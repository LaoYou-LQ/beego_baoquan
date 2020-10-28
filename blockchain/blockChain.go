package blockchain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt-master"
	"math/big"
)

//桶的名称，该桶用于装区块信息
var BUCKET_NAME = "blocks"

//表示最新的区块的key名
var LAST_KEY = "lasthash"

//存储区块数据的文件
var CHAINDB = "chain.db"

/*
定义BlockChain结构体的目的
区块链结构体实例定义：用于表示代表一条区块链
	功能：
		①将新产生的区块与已有的区块链接起来，并保存
		②可以查询某个区块的信息
		③可以将所有区块遍历，输出区块信息
*/

type BlockChain struct {
	LastHash []byte //最新区块的hash
	BlotDb   *bolt.DB
}
/**
 * 查询所有的区块信息，并返回。将所有的区块放入到切片中
 */
func (bc BlockChain) QueryAllBlocks() []*Block {
	blocks := make([]*Block, 0)
	db := bc.BlotDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据错误")
		}
		eachkey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig := big.NewInt(0) //0的大整数
		for {
			eachBlockBytes := bucket.Get(eachkey)
			//反序列化以后得到的每一个区块
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到每一个区块结构体指针放入到[]byte容器中
			blocks = append(blocks, eachBlock)
			preHashBig.SetBytes(eachBlock.PrevHash)
			if preHashBig.Cmp(zeroBig)==0 {//通过if条件语句判断区块链遍历是否已到创世区块，如果到创世区块，跳出循环
				break
			}
			//否则，继续向前遍历
			eachkey = eachBlock.PrevHash
		}
		return nil
	})
	return blocks
}
func (bc BlockChain) QueryBlockByHeight(height int64)  {

}
/*
	用于创建一条区块链，并返回该区块链实例
		解释：要创建一条区块链先得创建区块，作为区块链的创始区块
	首先打开存储数据的文件（blot.Open)
	查看是否用桶存在，没有就创建
	有桶后查看桶内是否有数据
 	没有就调用函数CreateGenesisBlock（）创建创世区块



*/
func NewBlockChain() BlockChain {
	//0、打开存储区块数据的chain.db文件
	db, err := bolt.Open("chain.db", 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	//先从区块链中都看是否创世区块已经存在
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 { //没有创世区块
			//1、创建创世区块
			genesis := CreateGenesisBlock()
			//2、创建一个存储区块数据的文件
			fmt.Printf("genesis的Hash值:%x\n", genesis.Hash)
			//创建存储区块的文件（BlockChain ），结构体类型
			bl = BlockChain{
				LastHash: genesis.Hash,
				BlotDb:   db,
			}
			genesisBytes, _ := genesis.Serialize()
			bucket.Put(genesis.Hash, genesisBytes)
			bucket.Put([]byte(LAST_KEY), genesis.Hash)
		} else {
			//获取key值
			lastHash := bucket.Get([]byte(LAST_KEY))

			//通过key值获取value
			lastBlockBytes := bucket.Get(lastHash)
			//反序列化
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取数据失败")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BlotDb:   db,
			}
		}
		return nil
	})
	return bl
}

//1创建区块
//创建区块的函数（创始区块）
//2创建存储区块数据的文件

/*
	db:=bc.BlotDb

	db,err:=bolt.Open("chain.db",0600,nil)
	if err!=nil {
		panic(err.Error())
	}
	defer db.Close()
	//操作文件
	db.Update(func(tx *bolt.Tx) error {
		var tong *bolt.Bucket
		//bucket 桶
		//Bucket按名称检索一个Bucket。
		//如果桶不存在，返回nil。
		tong = tx.Bucket([]byte(BUCKET_NAME)) //赋值给桶
		//判断是否有桶
		if tong == nil {
			//桶不存在，创建一个桶creatBucket
			tong, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				return err
			}


			/*
				不能直接存东西，第一次装东西后里面有东西了，
				下次存东西应该先查看里面有没有东西
				如果里面有要存的东西了那就不能存
				所以先读（Get）
*/

//先查看获取桶中是否已包含要保存的区块
/*
		lastBlock := tong.Get([]byte("lasthash"))
		//区块序列化Serialize
		blockHash, err := block.Serialize()
		if err != nil {
			return nil
		}
		if lastBlock == nil { //未获取到最新区块的hash
			//Put设置存储桶中键的值。
			//如果键存在，那么它之前的值将被覆盖
			tong.Put(block.Hash, blockHash)
			tong.Put([]byte("lasthash"), blockHash)
		}


	}

*/

/*
将生产的新区块保存到文件中
*/

func (bc BlockChain) SaveData(data []byte) (Block, error) {
	db := bc.BlotDb
	var e error
	var lastBlock *Block
	//查询最新的区块的hash
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e = errors.New("boltdb未创建")
			return e
		}
		lastBlockBytes := bucket.Get([]byte(bc.LastHash))
		//反序列化
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})
	//先生成一个区块 把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化
		newBlockBytes, _ := newBlock.Serialize()
		//把区块信息保存到boltdb中
		bucket.Put(newBlock.Hash, newBlockBytes)
		//更新代表最后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})
	return newBlock, e
}
