package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

type Flags struct {
	Help    bool
	KeySize int
	IVSize  int
	Base64  bool
}

var flags Flags

func init() {
	flag.BoolVar(&flags.Help, "h", false, "Prints help")
	flag.IntVar(&flags.KeySize, "k", 24, "Key size in bytes")
	flag.IntVar(&flags.IVSize, "i", 8, "IV size in bytes")
	flag.BoolVar(&flags.Base64, "b", true, "Convert keys and iv to base64")
}

func main() {
	flag.Parse()

	if flags.Help {
		flag.Usage()
		fmt.Println("\n  * If you don't want a base64 string you have to set the base64 flag to false like this '-b=false'")
		return
	}

	key := make([]byte, flags.KeySize)
	iv := make([]byte, flags.IVSize)

	if _, err := rand.Read(key); err != nil {
		log.Fatal("Unable to generate key")
	}

	if _, err := rand.Read(iv); err != nil {
		log.Fatal("Unable to generate iv")
	}

	if flags.Base64 {
		fmt.Println("Key:", base64.StdEncoding.EncodeToString(key))
		fmt.Println("IV:", base64.StdEncoding.EncodeToString(iv))
		return
	}

	fmt.Println("Key:", key)
	fmt.Println("IV:", iv)
}
