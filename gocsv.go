package main

import (
	// "encoding/csv"
	"fmt"
	"reflect"
	"strings"

	"github.com/a-poor/tags"
)

// type CSV struct {
// 	r io.Reader
// }

type User struct {
	ID       int    `csv:"user_id"`
	Name     string `csv:"name"`
	IsCool   bool   `csv:"is_cool"`
	Nickname string `csv:"nickname,omitempty"`
}

const TestData = `user_id,name,is_cool,nickname
1,"Tom",true,"Tommy"
2,"Sue",true,
3,"Ron",false,"Ronny"
`

func getFieldNames(d interface{}) []string {
	var names []string
	t := reflect.TypeOf(d)
	for i := 0; i < t.NumField(); i++ {
		names = append(names, t.Field(i).Name)
	}
	return names
}

func formatHeader(fields []string) string {
	return strings.Join(fields, ",")
}

func main() {
	// Create a user
	u := User{
		ID:       1,
		Name:     "Tom",
		IsCool:   true,
		Nickname: "",
	}

	fn := getFieldNames(u)
	h := formatHeader(fn)
	fmt.Printf("Field names: %s\n", fn)
	fmt.Printf("CSV header: %q\n", h)

	// Get the user's struct tags
	ust := tags.ParseStructTags("csv", u)
	fmt.Println(ust)
}
