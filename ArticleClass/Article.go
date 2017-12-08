package ArticleClass

/* Article - Struct to hold article related values
* param - articleURL - string - the url to the article
* param - articleTitle - string - title of the article
* param - articleDescription - string - short blurb of article
 */
type Article struct {
	ArticleTitle string `xml:"title"`
}

type Articles struct {
	ArticleArray []Article `xml:"item"`
	// ArticleDescription string `xml:"description"`
	// ArticleURL         string `xml:"link"`
	// ArticleDate        string `xml:"pubDate"`
	// GUID               string `xml:"guid"`
}

// func (a Article) openURLInBrowser() {
// 	var err error
// 	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", a.ArticleURL).Start()

// 	if err != nil {
// 		fmt.Println("ERROR OPENING LINK IN BROWSER")
// 	}
// }
