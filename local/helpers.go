package local

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
)

//Check func if the default directory installed
func Check() (bool, string) {
	path, err := os.UserHomeDir()
	if err != nil {
		return false, err.Error()
	}
	path += "/.goPirate"
	return exists(path), path
}

//Readline func
func Readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

//Answer func
func Answer() bool {
	if Readline() == "y" {
		return true
	}
	return false
}

//RandomString func
func RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)

	return string(b)
}
