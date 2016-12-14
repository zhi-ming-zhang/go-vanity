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

// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58_test

import (
	"fmt"

	"github.com/zhi-ming-zhang/go-vanity/crypto/base58"
)

// This example demonstrates how to decode modified base58 encoded data.
func ExampleDecode() {
	// Decode example modified base58 encoded data.
	encoded := "25JnwSn7XKfNQ"
	decoded := base58.Decode(encoded)

	// Show the decoded data.
	fmt.Println("Decoded Data:", string(decoded))

	// Output:
	// Decoded Data: Test data
}

// This example demonstrates how to encode data using the modified base58
// encoding scheme.
func ExampleEncode() {
	// Encode example data with the modified base58 encoding scheme.
	data := []byte("Test data")
	encoded := base58.Encode(data)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: 25JnwSn7XKfNQ
}

// This example demonstrates how to decode Base58Check encoded data.
func ExampleCheckDecode() {
	// Decode an example Base58Check encoded data.
	encoded := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	decoded, version, err := base58.CheckDecode(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show the decoded data.
	fmt.Printf("Decoded data: %x\n", decoded)
	fmt.Println("Version Byte:", version)

	// Output:
	// Decoded data: 62e907b15cbf27d5425399ebf6f0fb50ebb88f18
	// Version Byte: 0
}

// This example demonstrates how to encode data using the Base58Check encoding
// scheme.
func ExampleCheckEncode() {
	// Encode example data with the Base58Check encoding scheme.
	data := []byte("Test data")
	encoded := base58.CheckEncode(data, 0)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: 182iP79GRURMp7oMHDU
}
