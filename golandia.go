package golandia

import (
	"fmt"
)

var url string
var version int64

type JasonAccount string
type JasonAccountId string

func Create(s JasonAccount) error {
	fmt.Printf("%s %s\n", url, s)
	return nil
}

func Fetch(ps *JasonAccount) error {
	*ps = "JASON VIVE"
	fmt.Printf("%s %s\n", url, *ps)
	return nil
}

func Delete(id JasonAccountId) error {
	fmt.Printf("%s %s %d\n", url, id, version)
	return nil
}
