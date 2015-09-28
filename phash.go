// Copyright (C) 2013 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package phash is a simple pHash wrapper library for the Go programming language.
package phash

/*
#cgo pkg-config: pHash

#include <stdlib.h>

typedef unsigned long long ulong64;

extern ulong64 pc_dct_imagehash_Wrapper(const char *file);
extern ulong64* pc_dct_videohash_Wrapper(const char *file);
extern int ph_hamming_distance(ulong64 hasha, ulong64 hashb);
*/
import "C"

import "unsafe"

// ImageHash returns a DCT pHash for image with a given path.
func ImageHashDCT(file string) (uint64, error) {
	cs := C.CString(file)

	h, err := C.pc_dct_imagehash_Wrapper(cs)
	C.free(unsafe.Pointer(cs))

	return uint64(h), err
}

func VideoHashDCT(file string) (error) {
	cs := C.CString(file)
	len := C.uint(0)

	h, err := C.pc_dct_videohash_Wrapper(cs, len)
	C.free(unsafe.Pointer(cs))
	println(*h, len)

	return err
}

// HammingDistanceForHashes returns a Hamming Distance between two images' DCT pHashes.
func HammingDistanceForHashes(hasha uint64, hashb uint64) (int, error) {
	d, err := C.ph_hamming_distance(C.ulong64(hasha), C.ulong64(hashb))

	return int(d), err
}

// HammingDistanceForFiles returns a Hamming Distance between two images with a given paths.
func HammingDistanceForFiles(filea string, fileb string) (interface{}, error) {
	hasha, err := ImageHashDCT(filea)
	if err != nil {
		return nil, err
	}

	hashb, err := ImageHashDCT(fileb)
	if err != nil {
		return nil, err
	}

	return HammingDistanceForHashes(hasha, hashb)
}
