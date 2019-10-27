//BasicPodcastMetaDataByIDRequest

package main

import (
	"fmt"
	listennotes "listennotes-go"
	"os"
)

func main() {
	client := listennotes.NewListenNotesClient(os.Getenv("API_KEY"))

	client.Params["rsses"] = "https://rss.art19.com/recode-decode,https://rss.art19.com/the-daily,https://www.npr.org/rss/podcast.php?id=510331,https://www.npr.org/rss/podcast.php?id=510331"
	client.Params["itunes_ids"] = "1457514703,1386234384,659155419"

	fmt.Println(client.BasicPodcastMetadata("3302bc71139541baa46ecb27dbf6071a", "68faf62be97149c280ebcc25178aa731", "37589a3e121e40debe4cef3d9638932a", "9cf19c590ff0484d97b18b329fed0c6a"))

}
