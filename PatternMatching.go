package main

import (
	"fmt"
	"github.com/fentec-project/gofe/data"
	"log"
	"os"
	"strconv"

	//"math/big"
	//"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
	"math/big"
	"time"
)

func main() {
	var mod = os.Args[1:] //Get the text and the pattern respectively from test/txt.txt and test/pattern.txt"
	fmt.Println(string(mod[0]))
	fmt.Println(string(mod[1]))
	var start string
	start = string(mod[0])
	x, _ := readMatFromFile("test/txt.txt")
	var (
		y data.Matrix
	)
	if string(mod[1]) != "encr_batch" {
		y, _ = readMatFromFile("test/pattern.txt")
	}
	l := 128 // length of input vectors
	boundX := big.NewInt(10000)
	boundY := big.NewInt(10000)
	trustedEntFHIPE, _ := fullysec.NewFHIPE(l, boundX, boundY)

	var mskFHIPE *fullysec.FHIPESecKey
	mskFHIPE, _ = trustedEntFHIPE.GenerateMasterKey()

	//Create key/cipher the pattern matrix
	var feKeyFHIPEs []*fullysec.FHIPEDerivedKey

	if string(mod[1]) != "encr_batch"{
		feKeyFHIPEs = deriveKeyPattern(y, trustedEntFHIPE, mskFHIPE)
	}
	//cipher the txt matrix
	cipherFHIPEs:= encryptText(x, trustedEntFHIPE, mskFHIPE, start)
	// Decrypt data end find the pattern:
	if string(mod[1]) != "encr_batch"{
		decryptionFindPattern(trustedEntFHIPE, x, cipherFHIPEs, feKeyFHIPEs, y, start)
	}
}

/*

 */
func encryptText(x data.Matrix, trustedEntFHIPE *fullysec.FHIPE, mskFHIPE *fullysec.FHIPESecKey, start string) []*fullysec.FHIPECipher {
	startEncryption := time.Now()
	cipherFHIPEs := make([]*fullysec.FHIPECipher, x.Rows())

	for i := 0; i < x.Rows(); i++ {
		cipherFHIPE, _ := trustedEntFHIPE.Encrypt(x[i], mskFHIPE)
		cipherFHIPEs[i] = cipherFHIPE
	}
	durationEncryption := time.Since(startEncryption)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println("Encryption time: ", durationEncryption)

	//Write duration on a file
	writeTimeOnFile(durationEncryption, "test/EncrTime.txt", start)
	return cipherFHIPEs
}

/*

 */
func writeTimeOnFile(durationEncryption time.Duration, fileName string, start string) {

	var (
		f   *os.File
		err error
	)
	if start == "0" {
		f, err = os.OpenFile(fileName,
			os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	} else{
		f, err = os.OpenFile(fileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err := f.WriteString(strconv.FormatFloat(durationEncryption.Seconds(), 'f', -1, 32) + "\n"); err != nil {
		log.Println(err)
	}
}

func decryptionFindPattern(trustedEntFHIPE *fullysec.FHIPE, x data.Matrix, cipherFHIPEs []*fullysec.FHIPECipher, feKeyFHIPEs []*fullysec.FHIPEDerivedKey, y data.Matrix, start string) {
	decFHIPE := fullysec.NewFHIPEFromParams(trustedEntFHIPE.Params)
	var state int64
	state = 0
	startFindAllPatterns := time.Now()

	for i := 0; i < x.Rows(); i++ {
		decryptedState, _ := decFHIPE.Decrypt(cipherFHIPEs[i], feKeyFHIPEs[state])
		state = decryptedState.Int64()

		if state == int64(y.Rows()-1) {
			fmt.Println("Pattern found")
		}
	}
	durationFindAllPatterns := time.Since(startFindAllPatterns)
	fmt.Println("Found all patterns time: ", durationFindAllPatterns)

	//Write duration to a file
	writeTimeOnFile(durationFindAllPatterns, "test/DecrTime.txt", start)
}

func deriveKeyPattern(y data.Matrix, trustedEntFHIPE *fullysec.FHIPE, mskFHIPE *fullysec.FHIPESecKey) []*fullysec.FHIPEDerivedKey {
	startDeriveKey := time.Now()
	feKeyFHIPEs := make([]*fullysec.FHIPEDerivedKey, y.Rows())

	for i := 0; i < y.Rows(); i++ {
		feKeyFHIPE, _ := trustedEntFHIPE.DeriveKey(y[i], mskFHIPE)
		feKeyFHIPEs[i] = feKeyFHIPE
	}
	durationDeriveKey := time.Since(startDeriveKey)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println("Derive key time: ", durationDeriveKey)
	return feKeyFHIPEs
}