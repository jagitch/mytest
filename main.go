package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jagitch/mytest/unit"
)

func main() {
	fmt.Println("nihao")
	uuid := uuid.New().String()
	fmt.Println("result:%d", uuid)
	fmt.Println("result:%d", unit.Add(1, 2))
}
