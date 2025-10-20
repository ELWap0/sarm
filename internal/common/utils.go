package common

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func FileExits(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}else {
		return true 
	}
}

func GenHash(path string) (string, error) {
	hasher := sha256.New()
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
