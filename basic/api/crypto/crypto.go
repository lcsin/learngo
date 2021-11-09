package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"log"
)

func main() {

	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	fmt.Println(string(hash))

	err = bcrypt.CompareHashAndPassword(hash, []byte("123456"))
	if err != nil {
		fmt.Println(false)
		return
	}
	fmt.Println(true)
}

// Md5SaltPassword MD5盐值加密
func Md5SaltPassword(password, salt string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	hash.Write([]byte(salt))
	st := hash.Sum(nil)
	return hex.EncodeToString(st)
}

// ScryptPassword scrypt加密
func ScryptPassword(password, salt string) string {
	saltByte := []byte(salt)
	dk, err := scrypt.Key([]byte(password), saltByte, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}

func BcryptPassword(password, salt string) string {
	return ""
}
