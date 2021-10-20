package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type encrypt struct{}

var Handler = new(encrypt)

func (e *encrypt) Run() {
	md5Algo()
	bcryptAlgo("test")
}

func md5Algo() {
	str := "rootpass"

	// method 1
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	fmt.Println(md5str1)

	// method 2
	w := md5.New()
	io.WriteString(w, str)
	// w.Sum(nil) 将w的hash转成[]byte格式
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))

	fmt.Println(md5str2)

	// method 3

	hasher := md5.New()
	hasher.Write([]byte(str))
	cipherStr := hasher.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	// hex.DecodeString(encryptedData)

	fmt.Println([]byte(str), cipherStr, encryptedData)
}

// Salt and hash the password using the bcrypt algorithm
func bcryptAlgo(str string) {
	// The second paramerter is 'cost', the number of cost implies is the power-of-two number of rounds of hashing
	hashedkey, _ := bcrypt.GenerateFromPassword([]byte(str), 8)
	fmt.Println("bcrypt", hashedkey)
	fmt.Println("bcrypt", string(hashedkey), len(string(hashedkey)))

	// The example of comparing the stored hashed password
	another := "$2a$08$IYjGRLt6.n6p7Jjw4Z8e/.haqAr0iJEM2ZZ/kjTGCvDrzjHmmc422"
	err := bcrypt.CompareHashAndPassword([]byte(another), []byte("test"))
	if err != nil {
		fmt.Println("bcrypt", err)
	}
}
