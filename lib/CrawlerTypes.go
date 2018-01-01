package CrawlerTypes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*Article - holds all article meta data that will be sent to DB for storage
 *param -
 */
type Article struct {
	articleName    string
	articleDate    string   //TODO: Research golang Time and replace string with that if possible
	articleAuthor  []string //Array to check for potential of multiple authors on an article
	articleLink    string
	articleSummary string
	articleSource  string
}

type Source struct {
	sourceName    string
	sourceLinks   []string
	sourceRegexes []string
}

type ESPN struct {
	Source
	apiKey string
}

func NewESPN() ESPN {
	//Set Basic Seller Varibles
	result := ESPN{}
	result.sourceName = "ESPN"
	result.apiKey = "26714e095d154b4fabf953a5aefc684f" //newsApi

	//Set SourceLinks - Add any links to pull content from in this array
	result.sourceLinks = append(result.sourceLinks, "http://www.espn.com/espn/rss/nba/news")
	result.sourceLinks = append(result.sourceLinks,
		"https://newsapi.org/v2/everything?q=nba&sources=espn&sortBy=publishedAt&apiKey="+result.apiKey)

	//Set SourceRegexes - Add any regexes to run on content here
	result.sourceRegexes = append(result.sourceRegexes, "http://www.espn.com/espn/rss/nba/news")
	result.sourceRegexes = append(result.sourceRegexes, "http://www.espn.com/espn/rss/nba/news")

	return result
}

/* JS
var url = 'https://newsapi.org/v2/everything?' +
          'q=Apple&' +
          'from=2017-12-25&' +
          'sortBy=popularity&' +
          'apiKey=26714e095d154b4fabf953a5aefc684f';
*/
func (s *Source) RetrieveSearchContent(client *http.Client) string {
	result := "****************************CONTENT**************************\n"

	for i := 0; i < len(s.sourceLinks); i++ {
		response, errorDescription := client.Get(s.sourceLinks[i])

		CheckError(errorDescription, "ERROR AT RESPONSE MESSAGE")

		defer response.Body.Close()

		curContent, errorDescription := ioutil.ReadAll(response.Body)

		CheckError(errorDescription, "ERROR AT READING RESPONSE BODY")

		result += string(curContent) + "\n************************************************************"
	}

	return result
}

//is it possible to make this more efficient? 2 for loops - O(2n) - O(n) - running time, not bad, but can it be
//BETTER?
func (s *Source) RunSellerRegexes(content string) []Article {
	var matchedContent []string
	var result []Article

	for i := 0; i < len(s.sourceRegexes); i++ {
		regex, errorDescription := regexp.Compile(s.sourceRegexes[i])
		CheckError(errorDescription, "REGEX")

		matchedContent = append(matchedContent, regex.FindAllString(content, -1)...)
	}

	var temp []Article
	for i := 0; i < len(matchedContent); i++ {
		//Parse Out Data in matched Content here
		//TODO: Make a RegexType class to show what kind of value is the regex trying to parse, eg. Title, etc.
		//Once done, replace sourceRegex with new Regex Type

	}

	return result
}

func CheckError(err error, message string) bool {
	result := false

	if err != nil {
		fmt.Println(message)
		fmt.Print(err.Error())
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
