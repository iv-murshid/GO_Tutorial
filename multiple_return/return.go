package multiple_return

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func Multiple_return() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}