package main

import (
	//"math/big"
	//"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
	"math/big"
)


func main() {
	// Instantiation of a trusted entity that
	// will generate master keys and FE key
	x, _ := readMatFromFile("test/txt.txt")
	y, _ := readMatFromFile("test/pattern.txt")
	l := 128 // length of input vectors
	//y := data.NewVector([]*big.Int{big.NewInt(1), big.NewInt(2)})
	//x := data.NewVector([]*big.Int{big.NewInt(3), big.NewInt(4)})

	boundX := big.NewInt(10000)
	boundY := big.NewInt(10000)
	trustedEntFHIPE, _ := fullysec.NewFHIPE(l, boundX, boundY)

	var mskFHIPE *fullysec.FHIPESecKey
	mskFHIPE, _ = trustedEntFHIPE.GenerateMasterKey()


	//Create key/cipher the pattern matrix
	feKeyFHIPEs := make([]*fullysec.FHIPEDerivedKey, y.Rows())

	for i := 0; i < y.Rows(); i++ {
		feKeyFHIPE, _ := trustedEntFHIPE.DeriveKey(y[i], mskFHIPE)
		feKeyFHIPEs[i] = feKeyFHIPE
	}

	//cipher the txt matrix
	cipherFHIPEs := make([]*fullysec.FHIPECipher, x.Rows())

 	for i := 0; i < x.Rows(); i++ {
		cipherFHIPE, _ := trustedEntFHIPE.Encrypt(x[i], mskFHIPE)
		cipherFHIPEs[i] = cipherFHIPE
	}

	decFHIPE := fullysec.NewFHIPEFromParams(trustedEntFHIPE.Params)

	var state int64
	state = 0
	//M := new(big.Int).Set(big.NewInt(int64(y.Rows())))

	for i := 0; i < x.Rows(); i++ {

		state_x, _  := decFHIPE.Decrypt(cipherFHIPEs[i], feKeyFHIPEs[state])
 		state = state_x.Int64()

  		if  state == int64(y.Rows()-1)    {
			print("Pattern found")
		}
	}







}