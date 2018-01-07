/*
package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}
func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n",u.name, u.email)
}
type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}
*/
package main

import (
	"fmt"
)

type USB interface {
	Name() string
	connecter
}
type connecter interface {
	Connect()
}

type phoneConnecter struct {
	name string
}

func (pc phoneConnecter) Name() string{
	return pc.name
}
func (pc phoneConnecter) Connect() {
	fmt.Println("Connected:",pc.name)
}
func main()  {
	a := phoneConnecter{"phoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb interface{}){
	/*if pc,ok := usb.(phoneConnecter);ok{
		fmt.Println("Disconnected:" ,pc.name)
		return
	}
	fmt.Println("Unknown device.")*/
	switch v := usb.(type) {
	case phoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unkonwn device.")
	}

}