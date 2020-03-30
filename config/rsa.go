package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// RsaEncrypt RSA加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	// 解密pem格式的公钥
	publicKey, err := ioutil.ReadFile("./files/public.pem")
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt RSA解密
func RsaDecrypt(cipherText []byte) ([]byte, error) {
	// 解密
	privateKey, err := ioutil.ReadFile("./files/private.pem")
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
