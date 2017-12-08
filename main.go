package main

import (
	. "Meridian/ArticleClass"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* PROGRESS: Basic Version of Article Class created

 */

/**
* TODO: Update Article Class to Support XML Marshalling
		List of new struct
		UI to display new struct
		Support for other news sites, maybe handler (espn handler, bleacher report handler, etc)
*/

func main() {
	fmt.Println("Project Meridian: V 0.11 Alpha")
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

	//fmt.Println("\nRSS CONTENT\n\n\n" + string(body))
	var articles Article

	fmt.Println("Project Meridian: UNMARSHALING XML CONTENT INTO ARTICLE STRUCT")
	fmt.Println(string(body))
	xml.Unmarshal(body, &articles)

	fmt.Println("ARTICLE TITLE" + articles.ArticleTitle)
}
