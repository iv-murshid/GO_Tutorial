package if_else

func If_else(num int) (string,error) {

    if num%2 == 0 {
        return "even",nil
    } else {
        return "odd",nil
    }

    // if 8%4 == 0 {
    //     fmt.Println("8 is divisible by 4")
    // }

    // if 8%2 == 0 || 7%2 == 0 {
    //     fmt.Println("either 8 or 7 are even")
    // }

    // if num := 9; num < 0 {
    //     fmt.Println(num, "is negative")
    // } else if num < 10 {
    //     fmt.Println(num, "has 1 digit")
    // } else {
    //     fmt.Println(num, "has multiple digits")
    // }
}