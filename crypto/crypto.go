// Copyright 2016 Zhiming Zhang <zhiming@protonmail.com> All rights reserved.
// This file is part of go-vanity.
//
// go-vanity is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-vanity is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-vanity.  If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"

	"github.com/zhi-ming-zhang/go-vanity/crypto/base58"
	"github.com/zhi-ming-zhang/go-vanity/crypto/secp256k1"

	"golang.org/x/crypto/ripemd160"
)

func Sha256(data []byte) []byte {
	hash := sha256.Sum256(data)

	return hash[:]
}

func Ripemd160(data []byte) []byte {
	ripemd := ripemd160.New()
	ripemd.Write(data)

	return ripemd.Sum(nil)
}

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
}

func SeckeyToStr(seckey []byte) string {
	var varsionNum []byte = append([]byte{0x80}, seckey...)
	checksumHash := Sha256(Sha256(varsionNum))
	var all []byte = append(varsionNum, checksumHash[:4]...)

	return base58.Encode(all)
}

func PubkeyToAddress(publicKey []byte) string {
	// 0 - Private ECDSA Key
	//fmt.Printf("0 - Private ECDSA Key: \n\t%X\n", privateKey)

	// 1 - Public ECDSA Key
	//fmt.Printf("1 - Public ECDSA Key: \n\t%X\n", publicKey)

	// 2 - SHA-256 hash of 1
	hashPubKey := Sha256(publicKey)
	//fmt.Printf("2 - SHA-256 hash of 1: \n\t%X\n", hashPubKey)

	// 3 - RIPEMD-160 Hash of 2

	var ripemd160hash []byte = Ripemd160(hashPubKey)
	//fmt.Printf("3 - RIPEMD-160 Hash of 2: \n\t%X\n", ripemd160hash)

	// 4 - Adding network bytes to 3
	var networkByteAndRipemd160hash []byte = append([]byte{0x00}, ripemd160hash...)
	//fmt.Printf("4 - Adding network bytes to 3: \n\t%X\n", networkByteAndRipemd160hash)

	// 5 - SHA-256 hash of 4
	// 445C7A8007A93D8733188288BB320A8FE2DEBD2AE1B47F0F50BC10BAE845C094

	hashOfStepFour := Sha256(networkByteAndRipemd160hash)
	//fmt.Printf("5 - SHA-256 hash of 4: \n\t%X\n", hashOfStepFour)

	// 6 - SHA-256 hash of 5
	hashOfStep5 := sha256.Sum256(hashOfStepFour)
	hashOfStepFive := hashOfStep5[:]
	//fmt.Printf("6 - SHA-256 hash of 5: \n\t%X\n", hashOfStepFive)

	// 7 - First four bytes of 6
	var first4ofStep6 []byte = hashOfStepFive[:4]
	//fmt.Printf("7 - First four bytes of 6: \n\t%X\n", first4ofStep6)

	// 8 - Adding 7 at the end of 4
	var fourAddSeven []byte = append(networkByteAndRipemd160hash, first4ofStep6...)
	//fmt.Printf("8 - Adding 7 at the end of 4: \n\t%X\n", fourAddSeven)

	// 9 - Base58 encoding of 8
	base58of8 := base58.Encode(fourAddSeven)
	//fmt.Printf("9 - Base58 encoding of 8: \n\t%s\n", base58of8)

	return base58of8
}
