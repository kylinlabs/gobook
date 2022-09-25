//第7章/fence_amd64.s
#include "textflag.h"

// func mfence()
TEXT ·mfence(SB), NOSPLIT, $0-0
    MFENCE
    RET
