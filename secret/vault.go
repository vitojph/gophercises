package secret

import (
	"errors"

	"github.com/vitojph/gophercises/secret/encrypt"
)

// Memory creates a in memory vault
func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

// Vault stores an encoding key
type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

// Get the value of a set key
func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	ret, err := encrypt.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}
	return ret, nil
}

// Set a new key:value pair
func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	v.keyValues[key] = encryptedValue
	return nil
}
