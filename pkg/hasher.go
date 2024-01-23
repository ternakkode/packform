package pkg

import "crypto/sha256"

func HashString(str string) []byte {
	hashTool := sha256.New()
	hashTool.Write([]byte(str))
	return hashTool.Sum(nil)
}
