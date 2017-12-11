package main

import (
	//. "Meridian/ArticleClass"
	espn "Meridian/ESPNRSSFeedClass"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* PROGRESS: Update Article Class to Support XML Marshalling
			ESPN Feed class created to handle xml marshaling of RSS feed

*/

/**
* TODO:	List of new struct
		UI to display new struct
		Support for other news sites, maybe handler (espn handler, bleacher report handler, etc)
*/

func main() {
	fmt.Println("Project Meridian: V 0.12 Alpha")
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

	var feed espn.Rss

	fmt.Println("Project Meridian: UNMARSHALING XML CONTENT INTO ARTICLE STRUCT")
	fmt.Println(string(body))
	xml.Unmarshal(body, &feed)

	var articles = feed.Channel.Articles
	for i := 0; i < len(articles); i++ {
		fmt.Println("-----------------------------------------\n\n\n\n")
		fmt.Println("ARTICLE Title: " + articles[i].ArticleTitle)
		fmt.Println("ARTICLE Description: " + articles[i].ArticleDescription)
		fmt.Println("ARTICLE Date: " + articles[i].ArticleDate)
		fmt.Println("ARTICLE Url: " + articles[i].ArticleURL)
		articles[i].OpenURLInBrowser()
		fmt.Println("\n\n\n\n-----------------------------------------")
	}
}
