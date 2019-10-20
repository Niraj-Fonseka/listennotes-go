//GetBasicMetadataEpisodes

package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))
	fmt.Println(client.BasicMetadataEpisodes("c577d55b2b2b483c969fae3ceb58e362", "0f34a9099579490993eec9e8c8cebb82"))

}
