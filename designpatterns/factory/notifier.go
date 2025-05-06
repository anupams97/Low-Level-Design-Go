package main

import "fmt"

type notifier interface {
	send()
}

type push struct {
}

func (p *push) send() {
	fmt.Println("push notification sent")
}

type email struct {
}

func (e *email) send() {
	fmt.Println("email notification sent")
}

type sms struct {
}

func (s *sms) send() {
	fmt.Println("email notification sent")
}
func getNotifier(t string) notifier {
	switch t {
	case "sms":
		return &sms{}
	case "email":
		return &email{}
	case "push":
		return &push{}
	default:
		return nil
	}

}

func main() {
	n := getNotifier("sms")
	n.send()
	n = getNotifier("push")
	n.send()
}
