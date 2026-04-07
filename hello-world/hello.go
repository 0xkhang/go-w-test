package main

import (
    "strings"
    "fmt"
)

const EngPrefix = "Hello, "
const ChinesePrefix = "Cao ni ma, "
const EspanolPrefix = "Ola, "
const VietPrefix = "Chao, "

func Hello(name, lang string) string {
    lang = strings.ToLower(lang)
    
    if name == "" {
        name = "World"
    }

    var prefix string

    switch lang {
        case "english":
            prefix = EngPrefix
        case "chinese":
            prefix = ChinesePrefix
        case "espanol":
            prefix = EspanolPrefix
        case "viet":
            prefix = VietPrefix
        default:
            fmt.Println("Unknown language")
    }

    return prefix + name
}
