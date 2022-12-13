package main

import "github.com/madwolf96/vault-env/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
