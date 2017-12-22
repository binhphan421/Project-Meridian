package CrawlerTypes

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Source struct {
	sourceName string
	sourceLink string
}

type ESPN struct {
	Source
}

func NewESPN() ESPN {
	result := ESPN{}
	result.sourceName = "ESPN"
	result.sourceLink = "http://www.espn.com/espn/rss/nba/news"
	return result
}

func (s *Source) RetrieveSearchContent(client *http.Client) string {
	response, errorDescription := client.Get(s.sourceLink)

	if !CheckErrors(errorDescription) {
		fmt.Println("ERROR AT RESPONSE MESSAGE")
	}

	defer response.Body.Close()

	result, errorDescription := ioutil.ReadAll(response.Body)

	if !CheckErrors(errorDescription) {
		fmt.Println("ERROR AT READING RESPONSE BODY")
	}

	return string(result)
}

func CheckErrors(err error) bool {
	result := false

	if err == nil {
		result = true
	}

	return result
}

/*****************************OLD ESPN CODE***************************************************
// package ESPN

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os/exec"
// )

// /*Article - Struct to hold article related values
// * param - articleURL - string - the url to the article
// * param - articleTitle - string - title of the article
// * param - articleDescription - string - short blurb of article
// * param - articleData - string - date of publishing
// * param - GUID - guid of item
//  */
// type Article struct {
// 	ArticleTitle       string `xml:"title"`
// 	ArticleDescription string `xml:"description"`
// 	ArticleURL         string `xml:"link"`
// 	ArticleDate        string `xml:"pubDate"`
// 	GUID               string `xml:"guid"`
// }

// /*Image - Struct to hold image related values
// * param - URL - string - the url to the image
// * param - Title - string - title of the image
// * param - Link - string - hyperlink of image
// * param - Width - int - width of image
// * param - Height - int - height of image
//  */
// type Image struct {
// 	URL    string `xml:"url"`
// 	Title  string `xml:"title"`
// 	Link   string `xml:"link"`
// 	Width  int    `xml:"width"`
// 	Height int    `xml:"height"`
// }

// /*Channel - Struct to hold rss channel related values
// * param - Title - string - the title of the rss channel
// * param - Desc - string - description of rss channel
// * param - Link - string - link to rss channel
// * param - TTL - string - TTL settings (controls how long a feed lives)
// * param - Language - string - language of feed
// * param - Generator - string - Generator?
// * param - Copyright - string - Copyright information of feed
// * param - LastBuildDate - string - date feed was last built/refreshed
// * param - Images - Image - image/logo of feed
// * param - Articles - []Article - array of articles with article related information
//  */
// type Channel struct {
// 	Title         string    `xml:"title"`
// 	Desc          string    `xml:"description"`
// 	Link          string    `xml:"link"`
// 	TTL           string    `xml:"ttl"`
// 	Language      string    `xml:"language"`
// 	Generator     string    `xml:"generator"`
// 	Copyright     string    `xml:"copyright"`
// 	LastBuildDate string    `xml:"lastbuilddate"`
// 	Images        Image     `xml:"image"`
// 	Articles      []Article `xml:"item"`
// }

// /*Rss - Overarching struct for rss feed
// * param - Channel - struct that holds channel related information
//  */
// type Rss struct {
// 	Channel Channel `xml:"channel"`
// }

// /*OpenURLInBrowser - Opens Article URL in Browser based on Article it is called from
//  */
// func (a Article) OpenURLInBrowser() {
// 	var err error
// 	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", a.ArticleURL).Start()

// 	if err != nil {
// 		fmt.Println("ERROR OPENING LINK IN BROWSER")
// 	}
// }

// /*GrabArticlesFromData - Generates article array
//  * Return - array of type Articles
//  * TODO: Break down function, (returing array content vs printing content)
//  */
// func GrabArticlesFromData() (articleArray []Article) {
// 	fmt.Println("Project Meridian: Pulling Content from ESPN NBA RSS FEED")
// 	response, error := http.Get("http://www.espn.com/espn/rss/nba/news")

// 	if error != nil {
// 		fmt.Println("Project Meridian: ERROR At GET CALL of ESPN NBA FEED")
// 	}

// 	defer response.Body.Close() //Close response after content has been processed

// 	body, error := ioutil.ReadAll(response.Body)

// 	if error != nil {
// 		fmt.Println("Project Meridian: ERROR At READALL of ESPN RSS FEED CALL")
// 	}

// 	var espnFeed Rss

// 	fmt.Println("Project Meridian: UNMARSHALING XML CONTENT INTO ARTICLE STRUCT")
// 	fmt.Println(string(body))
// 	xml.Unmarshal(body, &espnFeed)

// 	articleArray = espnFeed.Channel.Articles
// 	for i := 0; i < len(articleArray); i++ {
// 		fmt.Println("-----------------------------------------")
// 		fmt.Println("ARTICLE Title: " + articleArray[i].ArticleTitle)
// 		fmt.Println("ARTICLE Description: " + articleArray[i].ArticleDescription)
// 		fmt.Println("ARTICLE Date: " + articleArray[i].ArticleDate)
// 		fmt.Println("ARTICLE Url: " + articleArray[i].ArticleURL)
// 		fmt.Println("-----------------------------------------")
// 	}

// 	return articleArray
// }
