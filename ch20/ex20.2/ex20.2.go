package main

import (
	"goprojects/ch20/fedex"
	"goprojects/ch20/koreapost"
)

// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

type Sender interface {
	Send(parcel string)
}

// func SendBook(name string, sender *koreapost.PostSender) {
// 	sender.Send(name)
// }

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender1 := &fedex.FedexSender{}
	sender2 := &koreapost.PostSender{}

	SendBook("어린 왕자", sender1)
	SendBook("그리스인 조르바", sender1)

	SendBook("어린 왕자", sender2)
	SendBook("그리스인 조르바", sender2)
}
