package services

import n "github.com/matoous/go-nanoid/v2"

func Gid() string {
	id, _ := n.New(21)
	return id
}
