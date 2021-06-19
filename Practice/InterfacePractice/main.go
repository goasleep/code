package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	email string
}

func (u *user) notify() {
	fmt.Printf("user interface name:%s email:%s\n", u.name, u.email)
}

func (u *admin) notify() {
	fmt.Printf("Admin interface name:%s email:%s\n", u.name, u.email)
}

func main() {
	// u := &user{"smith", "smith@email.com"}
	// sendNotification(u)

	user1 := user{"smith", "smith@email.com"}
	ad := admin{
		user:  user1,
		email: "super@test.com",
	}
	fmt.Println(ad.email)      //smith
	fmt.Println(ad.user.email) //smith
	ad.notify()
	// sendNotification(&ad) //smithsmith@email.com

}

func sendNotification(n notifier) {
	n.notify()
}
