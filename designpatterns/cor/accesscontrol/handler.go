package accesscontrol

import "fmt"

type handler interface {
	Handle(*User)
	SetNext(next handler)
}
type User struct {
	Name            string
	Role            string
	IsAuthenticated bool
	IsBanned        bool
}
type AuthHandler struct {
	next handler
}

func (a *AuthHandler) Handle(user *User) {
	if user.IsAuthenticated {
		fmt.Println("Auth Handler: " + "Authenticated")
		a.next.Handle(user)
		return
	}
	fmt.Println("Auth Handler: " + "Not Authenticated")
}

func (a *AuthHandler) SetNext(next handler) {
	a.next = next
}

type RoleHandler struct {
	next handler
}

func (r *RoleHandler) Handle(user *User) {
	if user.Role == "Admin" {
		fmt.Println("Role Handler: " + "Allowed")
		r.next.Handle(user)
		return
	}
	fmt.Println("Role Handler: " + "Not Allowed")
}

func (r *RoleHandler) SetNext(next handler) {
	r.next = next
}

type BanHandler struct {
	next handler
}

func (b *BanHandler) Handle(user *User) {
	if !user.IsBanned {
		fmt.Println("Ban Handler: " + "Allowed")
		//ad nil check everywhere
		if b.next != nil {
			b.next.Handle(user)
		}

		return
	}
	fmt.Println("Ban Handler: " + "Not Allowed")
}

func (b *BanHandler) SetNext(next handler) {
	b.next = next
}
