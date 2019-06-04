package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/quxiaolong/account/pkg/config"
	"golang.org/x/crypto/argon2"
)

/*
  Argon2( Argon2(sha512(password) + uid, strength 10) + global_salt, strength 10)
 */
func MakePassword(password string, userSlat string) string {
	// sha512(password)
	sha := sha512.New512_256()
	sha.Write([]byte(password))
	pwd := sha.Sum(nil)

	// Argon2
	pwd = argon2.IDKey(pwd, []byte(userSlat), 1, 64*1024, 4, 64)

	// AES TODO: 换个 AES 算法吧
	pwd = argon2.IDKey(pwd, []byte(config.Conf.Auth.SecretKey), 1, 64*1024, 4, 64)

	return base64.StdEncoding.EncodeToString(pwd)
}
