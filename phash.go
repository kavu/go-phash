package phash

/*
#cgo LDFLAGS: -lstdc++
#cgo pkg-config: pHash

#include <stdlib.h>

#if defined( _MSC_VER) || defined(_BORLANDC_)
  typedef unsigned _uint64 ulong64;
  typedef signed _int64 long64;
#else
  typedef unsigned long long ulong64;
  typedef signed long long long64;
#endif

extern ulong64 pc_dct_imagehash_Wrapper(const char *file);
*/
import "C"

import "unsafe"

func ImageHash(file string) (uint64, error) {
	cs := C.CString(file)

	h, err := C.pc_dct_imagehash_Wrapper(cs)
	C.free(unsafe.Pointer(cs))

	return uint64(h), err
}
