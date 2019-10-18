package main

import (
	"fmt"
	listennotes "listennotes-go"
)

func main() {
	client := listennotes.NewListenNotesClient("API_KEY")

	fmt.Println(client)
}
