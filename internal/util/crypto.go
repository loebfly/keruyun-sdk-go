package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string) string {
	// data 进行 sha256 计算
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(data))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
