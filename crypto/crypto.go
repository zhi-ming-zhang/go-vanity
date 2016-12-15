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

func SeckeyToWIF(seckey []byte) string {
	var varsionNum []byte = append([]byte{0x80}, seckey...)
	checksumHash := Sha256(Sha256(varsionNum))

	return base58.Encode(append(varsionNum, checksumHash[:4]...))
}

func PubkeyToAddr(pubkey65 []byte) string {
	hash160 := append([]byte{0x00}, Ripemd160(Sha256(pubkey65))...)
	hash2s := Sha256(Sha256(hash160))

	return base58.Encode(append(hash160, hash2s[:4]...))
}
