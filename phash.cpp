#include <pHash.h>

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

#ifdef __cplusplus
}
#endif
