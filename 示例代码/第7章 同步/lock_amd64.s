//第7章/lock_amd64.s
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
    PAUSE
    JMP   again
ok:
    RET
