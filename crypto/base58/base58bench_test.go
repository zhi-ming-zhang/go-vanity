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

// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58_test

import (
	"bytes"
	"testing"

	"github.com/zhi-ming-zhang/go-vanity/crypto/base58"
)

func BenchmarkBase58Encode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		base58.Encode(data)
	}
}

func BenchmarkBase58Decode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	encoded := base58.Encode(data)
	b.SetBytes(int64(len(encoded)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		base58.Decode(encoded)
	}
}
