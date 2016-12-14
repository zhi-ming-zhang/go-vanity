<!--
Copyright 2016 Zhiming Zhang <zhiming@protonmail.com> All rights reserved.
This file is part of go-vanity.

go-vanity is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

go-vanity is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with go-vanity.  If not, see <http://www.gnu.org/licenses/>.
-->

base58
==========

[![Build Status](http://img.shields.io/travis/btcsuite/btcutil.svg)]
(https://travis-ci.org/btcsuite/btcutil) [![ISC License]
(http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://godoc.org/github.com/btcsuite/btcutil/base58?status.png)]
(http://godoc.org/github.com/btcsuite/btcutil/base58)

Package base58 provides an API for encoding and decoding to and from the
modified base58 encoding.  It also provides an API to do Base58Check encoding,
as described [here](https://en.bitcoin.it/wiki/Base58Check_encoding).

A comprehensive suite of tests is provided to ensure proper functionality.

## Installation and Updating

```bash
$ go get -u github.com/btcsuite/btcutil/base58
```

## Examples

* [Decode Example]
  (http://godoc.org/github.com/btcsuite/btcutil/base58#example-Decode)  
  Demonstrates how to decode modified base58 encoded data.
* [Encode Example]
  (http://godoc.org/github.com/btcsuite/btcutil/base58#example-Encode)  
  Demonstrates how to encode data using the modified base58 encoding scheme.
* [CheckDecode Example]
  (http://godoc.org/github.com/btcsuite/btcutil/base58#example-CheckDecode)  
  Demonstrates how to decode Base58Check encoded data.
* [CheckEncode Example]
  (http://godoc.org/github.com/btcsuite/btcutil/base58#example-CheckEncode)  
  Demonstrates how to encode data using the Base58Check encoding scheme.

## License

Package base58 is licensed under the [copyfree](http://copyfree.org) ISC
License.
