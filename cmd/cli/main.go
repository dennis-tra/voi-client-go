package main

import (
	voi "dennis-tra/voi-client-go/pkg"
)

func main() {
	client := voi.NewClient()

	err := client.Authenticate("https://link.voiapp.io/J9zmmFqaG0")
	if err != nil {
		panic(err)
	}
}
