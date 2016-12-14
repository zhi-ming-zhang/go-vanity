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

/*
Package base58 provides an API for working with modified base58 and Base58Check
encodings.

Modified Base58 Encoding

Standard base58 encoding is similar to standard base64 encoding except, as the
name implies, it uses a 58 character alphabet which results in an alphanumeric
string and allows some characters which are problematic for humans to be
excluded.  Due to this, there can be various base58 alphabets.

The modified base58 alphabet used by Bitcoin, and hence this package, omits the
0, O, I, and l characters that look the same in many fonts and are therefore
hard to humans to distinguish.

Base58Check Encoding Scheme

The Base58Check encoding scheme is primarily used for Bitcoin addresses at the
time of this writing, however it can be used to generically encode arbitrary
byte arrays into human-readable strings along with a version byte that can be
used to differentiate the same payload.  For Bitcoin addresses, the extra
version is used to differentiate the network of otherwise identical public keys
which helps prevent using an address intended for one network on another.
*/
package base58
