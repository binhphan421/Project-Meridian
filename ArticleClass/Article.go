package ArticleClass

import (
	"fmt"
	"os/exec"
)

/* Article - Struct to hold article related values
* param - articleURL - string - the url to the article
* param - articleTitle - string - title of the article
* param - articleDescription - string - short blurb of article
* param - articleData - string - date of publishing
* param - GUID - guid of item
 */
type Article struct {
	ArticleTitle       string `xml:"title"`
	ArticleDescription string `xml:"description"`
	ArticleURL         string `xml:"link"`
	ArticleDate        string `xml:"pubDate"`
	GUID               string `xml:"guid"`
}

func (a Article) OpenURLInBrowser() {
	var err error
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", a.ArticleURL).Start()

	if err != nil {
		fmt.Println("ERROR OPENING LINK IN BROWSER")
	}
}
