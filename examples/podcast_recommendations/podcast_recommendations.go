package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	fmt.Println(client.PodcastRecommendations("25212ac3c53240a880dd5032e547047b"))

}
