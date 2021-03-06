package users

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"time"
)

type User struct {
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Email     string    `json:"-" bson:"email"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"-" bson:"password,omitempty"`
	Addresses []Address `json:"-,omitempty" bson:"-"`
	Cards     []Card    `json:"-,omitempty" bson:"-"`
	UserID    string    `json:"id" bson:"-"`
	Links     Links     `json:"_links"`
	Salt      string    `json:"-" bson:"salt"`
}

func New() User {
	u := User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
	u.NewSalt()
	return u
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return fmt.Errorf("Error missing %v", "FirstName")
	}
	if u.LastName == "" {
		return fmt.Errorf("Error missing %v", "LastName")
	}
	if u.Username == "" {
		return fmt.Errorf("Error missing %v", "Username")
	}
	if u.Password == "" {
		return fmt.Errorf("Error missing %v", "Password")
	}
	return nil
}

func (u *User) MaskCCs() {
	for i, c := range u.Cards {
		c.MaskCC()
		u.Cards[i] = c
	}
}

func (u *User) AddLinks() {
	u.Links.AddCustomer(u.UserID)
}

func (u *User) NewSalt() {
	h := sha1.New()
	io.WriteString(h, strconv.Itoa(int(time.Now().UnixNano())))
	u.Salt = fmt.Sprintf("%x", h.Sum(nil))
}
