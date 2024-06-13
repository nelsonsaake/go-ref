package main

import (
    "fmt"
    "log"

    "github.com/eiannone/keyboard"
)

func main() {
    if err := keyboard.Open(); err != nil {
        log.Fatal(err)
    }
    defer keyboard.Close()

    fmt.Println("Press any key to see its ASCII value or ESC to quit")

    for {
        char, key, err := keyboard.GetKey()
        if err != nil {
            log.Fatal(err)
        }

        if key == keyboard.KeyEsc {
            break
        }

        fmt.Printf("You pressed: %c, ASCII: %d\n", char, char)
    }
}
