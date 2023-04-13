package utils

import (
	_ "crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

func Md5Encryption(pwd string, str string) (encryptpassword string) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha256.New}
	salt, encodedPwd := password.Encode(pwd, options)
	newPassword := fmt.Sprintf(str+"$%s$%s", salt, encodedPwd)
	return newPassword
}

func CheckPwd(pwd, EncryptedPassWord string) (check bool, err error) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha256.New}
	passWordInfo := strings.Split(EncryptedPassWord, "$")
	check = password.Verify(pwd, passWordInfo[1], passWordInfo[2], options)
	return check, nil
}
