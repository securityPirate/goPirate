package local

import (
	"gopirate.com/crypter"
	"io/ioutil"
	"os"
)

//OpenFile function loading file
//to close the file you need to call assingee.close()
func OpenFile(path string) (*os.File, error) {
	// If the file doesn't exist, create it, or append to the file
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

//SaveFile func
func SaveFile(path string, bytes []byte, mode os.FileMode) error {
	return ioutil.WriteFile(path, bytes, mode)
}

//ReadFromFile func
func ReadFromFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}

//MoveFile func
func MoveFile(old, new string) error {
	return os.Rename(old, new)
}

//DeleteFile func
func DeleteFile(path string) error {
	return os.Remove(path)
}

//SaveToFileEncrypted func
func SaveToFileEncrypted(path string, data []byte, sym crypter.Symmetric) {
	SaveFile(path, sym.Encrypt(data), 0600)
}

//ReadEncryptedFile func
func ReadEncryptedFile(path string, sym crypter.Symmetric) []byte {
	file, _ := ioutil.ReadFile(path)
	return sym.Decrypt([]byte(file))
}
