// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	external_3524d37 "github.com/go-generalize/multipart-util"
	"time"
)

type GetRequest struct {
	Enum  Enum
	Param string
	Time  time.Time
}

type GetResponse struct {
	Data string
}

type PostCreateTableRequest struct {
	File  *external_3524d37.MultipartPayload   `form:"file" json:"-"`
	Files []*external_3524d37.MultipartPayload `form:"files" json:"-"`
	Flag  int
	ID    string
	Map   map[int]int `json:"map"`
	Text  string
}

type PostCreateTableResponse struct {
	ID          string
	RequestTime time.Time
}

type PostCreateUserRequest struct {
	Birthday time.Time
	Gender   int
	ID       string `param:"id"`
	Password string
	Roles    []*Role
}

type PostCreateUserResponse struct {
	CreatedType CreatedType
	Message     string
	RequestedAt time.Time
	Status      bool
}

type Role struct {
	ID             int
	Name           string
	RecursionRoles []Role
}

type CreatedType int

const (
	CreatedTypeGuest  int = 2
	CreatedTypeMember int = 1
	CreatedTypeOwner  int = 0
)

type Enum string

const (
	EnumA string = "A"
	EnumB string = "B"
	EnumC string = "C"
)
