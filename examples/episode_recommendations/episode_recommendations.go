package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	fmt.Println(client.EpisodeRecommendations("914a9deafa5340eeaa2859c77f275799"))

}
