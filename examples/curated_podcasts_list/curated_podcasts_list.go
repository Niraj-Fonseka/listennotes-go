package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	client.Params["page"] = "2"
	fmt.Println(client.CuratedPodcasts())

}
