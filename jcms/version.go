package jcms

import "fmt"

const (
    VMAJOR = 0
    VMINOR = 0
    VPATCH = 171221
)

func Version () string {
    v := fmt.Sprintf ("%d.%d", VMAJOR, VMINOR)
    if VPATCH > 0 {
        v = fmt.Sprintf ("%s.%d", v, VPATCH)
    }
    return v
}
