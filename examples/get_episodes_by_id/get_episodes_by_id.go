package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	id := "02f0123246c944e289ee2bb90804e41b"

	fmt.Println(client.DetailedMetadataEpisode(id))

}
