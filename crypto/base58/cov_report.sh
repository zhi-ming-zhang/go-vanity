#!/bin/sh

# Copyright 2016 Zhiming Zhang <zhiming@protonmail.com> All rights reserved.
# This file is part of go-vanity.
#
# go-vanity is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# go-vanity is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with go-vanity.  If not, see <http://www.gnu.org/licenses/>.

# This script uses gocov to generate a test coverage report.
# The gocov tool my be obtained with the following command:
#   go get github.com/axw/gocov/gocov
#
# It will be installed to $GOPATH/bin, so ensure that location is in your $PATH.

# Check for gocov.
type gocov >/dev/null 2>&1
if [ $? -ne 0 ]; then
	echo >&2 "This script requires the gocov tool."
	echo >&2 "You may obtain it with the following command:"
	echo >&2 "go get github.com/axw/gocov/gocov"
	exit 1
fi
gocov test | gocov report
