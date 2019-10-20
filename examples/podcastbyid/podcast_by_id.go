package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	id := "4d3fe717742d4963a85562e9f84d8c79"

	fmt.Println(client.PodcastsByID(id))

}
