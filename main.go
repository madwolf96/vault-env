package main

import "vault-env/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
