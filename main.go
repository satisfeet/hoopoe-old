package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := struct {
		Created time.Time `json:"created"`
	}{time.Now()}

	b, err := json.Marshal(&s)

	if err != nil {
		panic(err)
	}

	fmt.Printf("json: %s\n", string(b))
	fmt.Printf("manuel: %s\n", `{"created":"`+s.Created.Format(time.RFC3339Nano)+`"}`)
}
