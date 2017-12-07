package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* PROGRESS: Basic get call to grab content from RSS feed

 */

/**
* TODO: Create article struct to hold (url, title, description)
		List of new struct
		UI to display new struct
		Support for other news sites, maybe handler (espn handler, bleacher report handler, etc)
*/

func main() {
	fmt.Println("Project Meridian: V 0.1 Alpha")
	fmt.Println("Project Meridian: Pulling Content from ESPN NBA RSS FEED")
	response, error := http.Get("http://www.espn.com/espn/rss/nba/news")

	if error != nil {
		fmt.Println("Project Meridian: ERROR At GET CALL of ESPN NBA FEED")
	}

	defer response.Body.Close() //Close response after content has been processed

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		fmt.Println("Project Meridian: ERROR At READALL of ESPN RSS FEED CALL")
	}

	fmt.Println("\nRSS CONTENT\n\n\n" + string(body))
}
