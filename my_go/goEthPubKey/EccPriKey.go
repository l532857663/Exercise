package main

import (
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"crypto/elliptic"
	"crypto/rand"
)

func main() {
	src := []byte(string("少壮不努力,活该你单身23333"))
	//这里是我实验,使用不同消息hash摘要进行验证,不用就好
	//src1 := []byte(string("少壮不努力,老大徒伤悲3344"))
	mysrc := myHash(src)
	//mysrc1 := myHash(src1)

	prk, puk, _ := genePriPubKey()
	fmt.Println("prk:", prk)
	mystring := sign(prk, mysrc)

	r, s := getSign(mystring)

	result := verifySign(&r, &s, mysrc, puk)
	fmt.Print(result)
}

//生成公/私钥
func genePriPubKey() (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {

	var err error
	var pubkey ecdsa.PublicKey
	var prikey *ecdsa.PrivateKey
	var curve elliptic.Curve

	curve = elliptic.P384() //使用的是P384椭圆曲线,可以尝试其他曲线，结果一样
	prikey, err = ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return prikey, pubkey, err
	}
	pubkey = prikey.PublicKey

	return prikey, pubkey, err
}

//将消息进行hash摘要;我这里使用的是md5,大家也可以尝试其他hash算法
func myHash(src []byte) []byte {
	myhash := md5.New()
	myhash.Write(src)
	return myhash.Sum(nil)
}

//使用私钥,hash摘要进行签名
func sign(key *ecdsa.PrivateKey, myhash []byte) string {
	r, s, _ := ecdsa.Sign(rand.Reader, key, myhash) //返回一对大整数
	rm, _ := r.MarshalText()                        //需要将大整数r,s转换成[]byte,才能被string
	sm, _ := s.MarshalText()

	return hex.EncodeToString([]byte(string(rm) + "+" + string(sm))) //编译成hex
}

//获取签名之后的密文,进行解密
func getSign(hexrs string) (rint, sint big.Int) {
	st, _ := hex.DecodeString(hexrs)
	str := strings.Split(string(st), "+") //解码之后的st,由r,s2部分拼接,需要对其拆分,然后逐一解码成大整数
	_ = rint.UnmarshalText([]byte(str[0]))
	//假如rint是指针:error: invalid memory address or nil pointer dereference
	_ = sint.UnmarshalText([]byte(str[1]))

	return
}

//对解密之后的明文,hash,公钥验证
func verifySign(rint, sint *big.Int, myhash []byte, pubkey ecdsa.PublicKey) bool {
	//使用ecdsa.Verify进行验证即可
	result := ecdsa.Verify(&pubkey, myhash, rint, sint)
	return result
}
