package main

import (
        "fmt"
        "os"
        "strconv"
)

func fibonacci(n int) int {
        if n == 0 {
                return 0
        }
        a, b := 0, 1
        for i := 2; i <= n; i++ {
                a, b = b, a+b
        }
        return b
}

func main() {
        input, e := strconv.Atoi(os.Args[1])
        if e != nil {
                panic(e)
        }
        r := 0
        for i := 0; i < input; i++ {
                r += fibonacci(i)
        }
        fmt.Println(r)
}