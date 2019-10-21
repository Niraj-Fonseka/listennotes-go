package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	fmt.Println(client.CuratedPodcastsListByID("SDFKduyJ47r"))

}
