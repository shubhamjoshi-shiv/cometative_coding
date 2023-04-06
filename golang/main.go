package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uid := uuid.New().String()
	fmt.Println(uid)

}
