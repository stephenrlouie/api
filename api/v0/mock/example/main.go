package main

import (
	"fmt"

	"wwwin-github.cisco.com/edge/optikon/api/v0/mock"
)

type test struct {
	Name string
	Age  int
}

func main() {
	testing := test{}
	statusCode, err := mock.GetMock("./mock.json", &testing)
	if err != nil {
		fmt.Printf("ERR: %v", err)
	}
	fmt.Printf("testing: %v\n", testing)
	fmt.Printf("Status: %v\n", statusCode)
}
