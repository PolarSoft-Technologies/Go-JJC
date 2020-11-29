package auth

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/PolarSoft-Technologies/Go-JJC/src/config"
	"github.com/dgrijalva/jwt-go"
)

//get pollbase signing key
var pbSigningKey = []byte(config.Config("SECRETE_SIGNING_KEY"))

//GenerateToken for user
func GenerateToken(userid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() //JWT token would expires in 30 minutes time

	tokenString, err := token.SignedString(pbSigningKey)

	if err != nil {

		fmt.Errorf("Token Generation Failed: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

//PWHashStore use moderate profile to pack hashed password into PWHashStr.
func PWHashStore(pw string) PWHashStr {
	s := make([]C.char, cryptoPWHashStrBytes)
	pwc := C.CString(pw)
	defer C.free(unsafe.Pointer(pwc))

	if int(C.crypto_pwhash_str(
		&s[0],
		pwc,
		(C.ulonglong)(len(pw)),
		(C.ulonglong)(CryptoPWHashOpsLimitModerate),
		(C.size_t)(CryptoPWHashMemLimitModerate))) != 0 {
		panic("see libsodium")
	}
	return PWHashStr{C.GoStringN(&s[0], C.int(cryptoPWHashStrBytes))}
}
