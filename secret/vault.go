package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/vitojph/gophercises/secret/cipher"
)

// Vault stores an encoding key
type Vault struct {
	filepath    string
	encodingKey string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// File creates a file vault
func File(filepath, encodingKey string) (*Vault, error) {
	return &Vault{
		filepath:    filepath,
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}, nil
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r) // -> decryptReader -> file
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w) // encryptWriter -> file
	return enc.Encode(&v.keyValues)
}

// Get the value of a set key
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}

// Set a new key:value pair
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.save()
	return err
}

// ListSecrets returns the list of secrets stored
func (v *Vault) ListSecrets() ([]string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return nil, err
	}

	secrets := make([]string, 0, len(v.keyValues))
	for k := range v.keyValues {
		secrets = append(secrets, k)
	}
	return secrets, nil
}

// Remove deletes a secret
func (v *Vault) Remove(k string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}

	_, ok := v.keyValues[k]
	if ok {
		delete(v.keyValues, k)
		err = v.save()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("💥 %s was successfully deleted 💥", k), nil
	}
	return "", errors.New("secret: no value for that key")
}
