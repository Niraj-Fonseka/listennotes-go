package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	fmt.Println(client)

	fmt.Println(client.Params)

	client.Params["genre_id"] = 93
	client.Params["page"] = 2
	client.Params["region"] = "us"
	client.Params["safe_mode"] = 1
	fmt.Println(client.BestPodcasts())
	fmt.Println(client.Params)
}
