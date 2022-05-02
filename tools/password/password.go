package password

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/coreservice-io/UUtils/rand_util"
)

const (
	crypto    = "shiro1"
	algorithm = "SHA-256"
	iteration = 500000
	saltLen   = 16
)

var enc = base64.StdEncoding

// hash the password with rand salt
func HashAndSalt(password string) string {
	salt := rand_util.GenRandStr(saltLen)
	hashed := computeHash([]byte(salt), []byte(password))
	return fmt.Sprintf("$%s$%s$%d$%s$%s", crypto, algorithm, iteration, encodeBase64([]byte(salt)), encodeBase64(hashed[:]))
}

// check if the password submitted matches the psssword saved
func PasswordMatch(password, saved string) (bool, error) {
	strs := strings.Split(saved, "$")
	salt, err := enc.DecodeString(strs[4])
	if err != nil {
		return false, err
	}
	hashed := computeHash(salt, []byte(password))
	if enc.EncodeToString(hashed[:]) == strs[5] {
		return true, nil
	}
	return false, nil
}

func encodeBase64(h []byte) []byte {
	dst := make([]byte, enc.EncodedLen(len(h)))
	enc.Encode(dst, h)
	return dst
}

// hash password with salt
func computeHash(salt, password []byte) [32]byte {
	hash := sha256.New()
	hash.Write(salt)
	hash.Write([]byte(password))
	hash32 := hash.Sum(nil)

	hashed := sha256.Sum256(hash32)
	for i := 0; i < iteration-2; i++ {
		hashed = sha256.Sum256(hashed[:])
	}

	return hashed
}
