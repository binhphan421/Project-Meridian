package main

import (
	espn "Meridian/ESPNRSSFeedClass"
	"fmt"
)

/**
* PROGRESS: Compartmentalize code

 */

/**
* TODO:	List of new struct
		UI to display new struct
		Support for other news sites, maybe handler (espn handler, bleacher report handler, etc)
*/

func main() {
	fmt.Println("Project Meridian: V 0.13 Alpha")
	text := ""

	for text != "Q" {
		fmt.Print("Enter Source ID or Q to Quit: ")
		fmt.Scanln(&text)

		GrabArticlesFromSourceId(text)
	}

}

//For now 1 = ESPN
func GrabArticlesFromSourceId(sourceType string) {
	if sourceType == "1" {
		fmt.Println("Parsing ESPN Data")
		espn.GrabArticlesFromData()
	}
}
