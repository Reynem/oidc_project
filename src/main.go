package main

import (
	"fmt"

	_ "github.com/coreos/go-oidc/v3/oidc"
	_ "golang.org/x/oauth2"
)

func main() {
	fmt.Println("hello world")

	router := mux.NewRouter
}
