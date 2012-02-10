// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

package main

import (
	"bufio"
	"flag"
	"github.com/cznic/mathutil"
	"log"
	"math"
	"os"
)

/* 

$ # Usage e.g.:
$ go run example.go -max 1024 > mathutil.dat # generate 1kB of "random" data

*/
func main() {
	if r, err := mathutil.NewFC32(math.MinInt32, math.MaxInt32, true); err != nil {
		log.Fatal(err)
	} else {
		var mflag uint64
		flag.Uint64Var(&mflag, "max", 0, "limit output to max bytes")
		flag.Parse()
		stdout := bufio.NewWriter(os.Stdout)
		if mflag != 0 {
			for i := uint64(0); i < mflag; i++ {
				if err := stdout.WriteByte(byte(r.Next())); err != nil {
					log.Fatal(err)
				}
			}
			stdout.Flush()
		} else {
			for stdout.WriteByte(byte(r.Next())) == nil {
			}
		}
	}
}