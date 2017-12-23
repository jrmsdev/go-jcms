// version info
package version

import "fmt"

const (
    MAJOR = 0
    MINOR = 0
    PATCH = 171223
)

func String () string {
    v := fmt.Sprintf ("%d.%d", MAJOR, MINOR)
    if PATCH > 0 {
        v = fmt.Sprintf ("%s.%d", v, PATCH)
    }
    return v
}
