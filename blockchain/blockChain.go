package blockchain

import "github.com/boltdb/bolt-master"

var BUCKET_NAME ="blocks"
var LAST_KEY ="lasthash"

/*
定义BlockChain结构体的目的
区块链结构体实例定义：用于表示代表一条区块链
	功能：
		①将新产生的区块与已有的区块链接起来，并保存
		②可以查询某个区块的信息
		③可以将所有区块遍历，输出区块信息
 */

type BlockChain struct {
	LastHash []byte  //最新区块的hash
	BlotDb *bolt.DB

}
/*
	用于创建一条区块链，并返回该区块链实例
		解释：要创建一条区块链先得创建区块，作为区块链的创始区块
 */
func NewBlockChain() BlockChain {
	//1创建区块
	genesis :=CreateGenesisBlock()//创建区块的函数（创始区块）
	//2创建存储区块数据的文件
	db,err:=bolt.Open("chain.db",0600,nil)
	if err!=nil {
		panic(err.Error())
	}
	bl:=BlockChain{
		LastHash:genesis.Hash,
		BlotDb: db,
	}
	//3将创建的区块存入chain.db中的一个桶中
	db.Update(func(tx *bolt.Tx) error {
		bucket,err:=tx.CreateBucket([]byte(BUCKET_NAME))
		if err!=nil {
			panic(err.Error())
		}
		//将创世区块保存到桶中
		serialBlock,err:=genesis.Serialize()
		if err!=nil {
			panic(err.Error())
		}
		bucket.Put(genesis.Hash,serialBlock)
		bucket.Put([]byte(LAST_KEY),genesis.Hash)
		return nil
	})
	return bl

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
}

/*
将生产的新区块保存到文件中
 */

func (bc BlockChain) SaveData(data []byte) {
	db:=bc.BlotDb
	var lastBlock *Block
	//查询最新的区块的hash
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte(BUCKET_NAME))
		if bucket==nil {
			panic("boltdb没有创建")
		}
		lasthash:=bucket.Get([]byte(LAST_KEY))
		lastBlockBytes:=bucket.Get(lasthash)
		//反序列化
		lastBlock,_=DeSerialize(lastBlockBytes)
		return nil
	})
	//先生成一个区块 把data存入到新生成的区块中
	newBlock:=NewBlock(lastBlock.Height+1,data,lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		return nil
	})
	})
	return
}
