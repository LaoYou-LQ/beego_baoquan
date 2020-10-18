package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func Md5Hash(data string) string {
	md5hash :=md5.New()
	md5hash.Write([]byte(data))
	bytes :=md5hash.Sum(nil)
	return  hex.EncodeToString(bytes)
}

func Md5HashReader(reader io.Reader) (string,error) {
	bytes, err := ioutil.ReadAll(reader)
	if err!=nil {
		return "",err
	}
	md5HashReader :=md5.New()
	md5HashReader.Write([]byte(bytes))
	md5HashBytes := md5HashReader.Sum(nil)
	return hex.EncodeToString(md5HashBytes),nil
}

