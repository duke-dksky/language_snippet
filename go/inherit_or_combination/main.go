package main

import "fmt"

type Notifier interface {
	Notify()
}

func SendNotification(notify Notifier) {
	notify.Notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() {
	fmt.Printf("send to %v(%v)\n", u.Name, u.Email)
}

type Admin struct {
	User
	Level string
}

func (a *Admin) Notify() {
	fmt.Printf("send Admin to %v(%v)\n", a.User.Name, a.User.Email)
}

func main() {
	bill := User{"Bill", "bill@email.com"}
	bill.Notify()

	jill := &User{"Jill", "jill@email.com"}
	jill.Notify()

	admin := &Admin{
		User: User{
			Name:  "john",
			Email: "john@email.com",
		},
		Level: "super",
	}
	admin.User.Notify()

	SendNotification(admin)
}
