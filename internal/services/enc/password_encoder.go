package enc

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

var (
	Sha384 = sha512.New384
	Sha1   = sha1.New
)

type (
	PasswordEncoder interface {
		Encode(password string) (string, error)
	}

	passwordEncoder struct{}
)

func NewPasswordEncoder() PasswordEncoder {
	return &passwordEncoder{}
}

func (pe *passwordEncoder) hash(h hash.Hash, str string) (string, error) {
	_, err := h.Write([]byte(str))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func (pe *passwordEncoder) Encode(password string) (string, error) {
	hash, err := pe.hash(Sha384(), "#@@a23%%4V3%%4"+password+"_#%@%$5X3%")
	if err != nil {
		return "", err
	}

	salt, err := pe.hash(Sha1(), password)
	if err != nil {
		return "", err
	}

	return hash + ":" + salt[0:5], nil
}
