package main

import (
	"os"

	listennotes "github.com/Niraj-Fonseka/listennotes-go"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("ListenNotesToken"))

	client.BestPodcasts("", "", 0, 0)
}
