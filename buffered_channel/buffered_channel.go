package buffered_channel

import "fmt"

func Buffered_channel() {

    messages := make(chan string, 2)

	go func(){
		messages <- "buffered"
		messages <- "channel"
	}()

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}