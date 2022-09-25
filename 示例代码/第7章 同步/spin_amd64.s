//第7章/spin_amd64.s
#include "textflag.h"

// func lock(ptr *int32, old, new int32)
TEXT ·lock(SB), NOSPLIT, $0-16
    MOVQ  ptr+0(FP), BX
    MOVL  old+8(FP), DX
    MOVL  new+12(FP), CX
again:
    MOVL  DX, AX
    LOCK
    CMPXCHGL  CX, 0(BX)
    JE    ok
    JMP   again
ok:
    RET

// func unlock(ptr *int32, val int32)
TEXT ·unlock(SB), NOSPLIT, $0-12
    MOVQ  ptr+0(FP), BX
    MOVL  val+8(FP), AX
    XCHGL AX, 0(BX)
    RET
