package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	"fmt"

	_ "DataCertProject/routers"

	"github.com/astaxie/beego"
)


func main() {
	//生成第一个区块
	/*
		1.程序执行时，对象存在电脑内存中
		2.长时间保存对象：将数据持久化的存储在硬盘上/数据库中，称之为持久化存储
		3.内存中的对象数据不能直接进行持久化存储，也不能进行直接传输
		4.序列化、反序列化：
				json是一种序列化和反序列化的格式
				xml也是一种序列化和反序列化的格式
		5.json格式：
			{
				"Id":1,
				"Phone":"13",
				"Password":"113411",
			}
		xml格式：
			<User>
				<Id>1<Id>
				<Phone>13<Phone>
				<Password>113411<Password>
	*/

	/*
	序列化和反序列化
	user1 := models.User{
		Id:       1,
		Phone:    "13",
		Password: "113611",
	}
	fmt.Println("内存中的数据user1：",user1)
	jsonByte, _ := json.Marshal(user1)序列化
	fmt.Println(string(jsonByte))
	var user2 models.User
	json.Unmarshal(jsonByte,&user2)
	fmt.Println("反序列化的user",user2)
*/

	//生成第一个区块
/*
	block:=blockchain.CreateGenesisBlock()
	fmt.Println(block)
	fmt.Printf("区块的hash值:%x", block.Hash)
	/*
	//Open：
	//Open在给定的路径上创建并打开一个数据库。
	////如果该文件不存在，那么它将自动创建。
	////传入nil选项将导致Bolt用默认选项打开数据库。
	db,err:=bolt.Open(CHAINDB,0600,nil)
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
		tong =tx.Bucket([]byte(BUCKET_NAME))//赋值给桶
		//判断是否有桶
		if tong==nil {
			//桶不存在，创建一个桶creatBucket
			tong,err=tx.CreateBucket([]byte(BUCKET_NAME))
			if err!=nil {
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
			lastBlock:=tong.Get([]byte("lasthash"))
			//区块序列化Serialize
			blockHash,err:=block.Serialize()
			if err!=nil {
				return nil
			}
			if lastBlock ==nil {//未获取到最新区块的hash
				//Put设置存储桶中键的值。
				//如果键存在，那么它之前的值将被覆盖
				tong.Put(block.Hash,blockHash)
				tong.Put([]byte("lasthash"),blockHash)
			}

			return nil
		}
	})


	 */

	//先准备一条区块链
	blockchain.NewBlockChain()
	fmt.Printf("111")
	db_mysql.ConDB()

	beego.SetStaticPath("/js", ".static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
