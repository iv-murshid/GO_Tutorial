package Switch

// import (
//     "fmt"
//     "time"
// )

func Switch(t int) (string , error){

    // i := 2
    // fmt.Print("Write ", i, " as ")
    // switch i {
    // case 1:
    //     fmt.Println("one")
    // case 2:
    //     fmt.Println("two")
    // case 3:
    //     fmt.Println("three")
    // }

    // switch time.Now().Weekday() {
    // case time.Saturday, time.Sunday:
    //     fmt.Println("It's the weekend")
    // default:
    //     fmt.Println("It's a weekday")
    // }

    switch {
    case t < 12:
        return "It's before noon", nil
    case t > 12:
        return "It's after noon", nil
    default:
        return "It's noon", nil
    }

    // whatAmI := func(i interface{}) {
    //     switch t := i.(type) {
    //     case bool:
    //         fmt.Println("I'm a bool")
    //     case int:
    //         fmt.Println("I'm an int")
    //     default:
    //         fmt.Printf("Don't know type %T\n", t)
    //     }
    // }
    // whatAmI(true)
    // whatAmI(1)
    // whatAmI("hey")
}