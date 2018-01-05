package doctype

import (
    "log"
    "fmt"
)

var engMap map[string]Engine

func init () {
    engMap = make (map[string]Engine)
}

func Register (name string, eng Engine) {
    _, exists := engMap[name]
    if exists {
        panic ("doctype engine already registered: " + name)
    }
    log.Println ("doctype register engine:", name)
    engMap[name] = eng
}

func GetEngine (name string) (Engine, error) {
    eng, ok := engMap[name]
    if !ok {
        return nil, fmt.Errorf ("invalid doctype engine: %s", name)
    }
    return eng, nil
}
