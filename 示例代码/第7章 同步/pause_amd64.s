//第7章/pause_amd64.s
#include "textflag.h"

// func pause()
TEXT pause(SB), NOSPLIT, $0-0
    PAUSE
    RET
