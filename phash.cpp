// Copyright (C) 2013 Max Riveiro
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

#include <pHash.h>
#include <iostream>

#ifdef __cplusplus
extern "C" {
#endif

ulong64 pc_dct_imagehash_Wrapper(const char *file) {
    cimg::exception_mode(0);
    ulong64 hash;

    if (ph_dct_imagehash(file, hash) == 0)
            errno = 0;

    return hash;
}

ulong64* pc_dct_videohash_Wrapper(const char *file, int *length) {
    cimg::exception_mode(0);
    
    printf("length before: %i\n", length);

    ulong64* h = ph_dct_videohash(file, length); 
    errno = 0;
    
    for (int i = 0; i < length; i++) 
        printf("pointer: %llu\n", *(h+i));
    printf("length after: %i\n", length);
    
    return h;
}

#ifdef __cplusplus
}
#endif
