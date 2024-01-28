// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetID() string
}

type AuthPayload struct {
	Token *string `json:"token,omitempty"`
}

type Shop struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (Shop) IsNode()            {}
func (this Shop) GetID() string { return this.ID }

type ShopsSearchInput struct {
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }

type UserCreateInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSearchInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
