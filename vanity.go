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

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/zhi-ming-zhang/go-vanity/crypto"
	"github.com/zhi-ming-zhang/go-vanity/crypto/secp256k1"
)

const (
	// alphabet is the modified base58 alphabet used by Bitcoin.
	alphabet     = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	alphabetIdx0 = '1'
)

// main, executes addrGen ad-infinitum, until a match is found
//
func main() {
	runtime.GOMAXPROCS(8)

	var toMatch string
	if len(os.Args) == 1 {
		errNoArg()
	} else {
		toMatch = os.Args[1]
	}

	for i := len(toMatch) - 1; i >= 0; i-- {
		if toMatch[0] != alphabetIdx0 {
			errPreix()
		} else if !strings.ContainsAny(alphabet, string(toMatch[i])) {
			errChar(string(toMatch[i]))
		}
	}

	println("\nSearching address...\n")

	for true {
		go addrGen(toMatch)
		time.Sleep(1 * time.Millisecond)
	}
}

func addrGen(toMatch string) {
	pubkey65, seckey := secp256k1.GenerateKeyPair()
	addrStr := crypto.PubkeyToAddr(pubkey65)
	addrMatch(addrStr, toMatch, seckey)
}

// tries to match the address with the string provided by the user, exits if successful
//
func addrMatch(addrStr string, toMatch string, seckey []byte) {
	addrStrMatch := strings.TrimPrefix(addrStr, toMatch)
	found := addrStrMatch != addrStr
	if found {
		keyStr := crypto.SeckeyToWIF(seckey)
		addrFound(addrStr, keyStr)
		os.Exit(0) // here the program exits when it found a match
	}
}

// non-interesting functions follow...

func addrFound(addrStr string, keyStr string) {
	println("Address found:")
	fmt.Printf("Address: %s\n", addrStr)
	fmt.Printf("Private: %s\n", keyStr)
	println("\nexiting...")
}

func errChar(char string) {
	fmt.Printf("\nPlease do not use \"%s\".", char)
	println("\nYou need to pass a vanity match, retry with the base58 alphabet:")
	fmt.Printf("\"%s\"", alphabet)
	println("\nexample: go run vanity.go 123")
	println("\nexiting...")
	os.Exit(1)
}

func errPreix() {
	println("\nYou need to pass a vanity match, retry with prefix '1' like: 123")
	println("\nexample: go run vanity.go 123")
	println("\nexiting...")
	os.Exit(1)
}

func errNoArg() {
	println("\nYou need to pass a vanity match, retry with an extra agrument like: 123")
	println("\nexample: go run vanity.go 123")
	println("\nexiting...")
	os.Exit(1)
}
