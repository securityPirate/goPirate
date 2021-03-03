package crypter

//generating AES-256-CTR

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

// Symmetric ...
type Symmetric struct {
	Key, IV, iv8 []byte
}

// Generate AES Key , IV
func (sym *Symmetric) Generate(pin string){
	x := md5.Sum([]byte(pin))
	sym.Key = []byte(hex.EncodeToString(x[:]))
	sym.IV = make([]byte, aes.BlockSize)
	sym.iv8 = make([]byte, aes.BlockSize)
	rand.Read(sym.IV)
	rand.Read(sym.iv8)
}

// Encrypt ...
func (sym Symmetric) Encrypt(plain []byte) []byte {
	block, _ := aes.NewCipher(sym.Key)
	stream := cipher.NewCTR(block, sym.IV)
	ciphered := make([]byte, (2*aes.BlockSize)+len(plain))
	copy(ciphered[aes.BlockSize/2:], sym.IV)
	copy(ciphered[0:aes.BlockSize/2], sym.iv8[0:aes.BlockSize/2])
	copy(ciphered[aes.BlockSize/2+aes.BlockSize:], sym.iv8[aes.BlockSize/2:])
	stream.XORKeyStream(ciphered[2*aes.BlockSize:], plain)
	return ciphered
}

// Decrypt ...
func (sym Symmetric) Decrypt(ciphered []byte) []byte {
	block, _ := aes.NewCipher(sym.Key)
	stream := cipher.NewCTR(block, sym.IV)
	plain := make([]byte, len(ciphered)-2*aes.BlockSize)
	stream.XORKeyStream(plain, ciphered[2*aes.BlockSize:])
	return plain
}
