package main

import (
	. "Meridian/lib"
	"fmt"
	"net/http"
	"time"
)

/**
* PROGRESS: Began Conversion of Meridan to Server Crawler
			Renaming Of Files
			Create base abstract class named Source
			Might not need to create seller class, but instead
			more NewSource functions to generate sources
*/

/**
* TODO:	Convert Meridian Into Server Based Crawler
		Set up seperation of content into Article struct
		Crawl on request, with 30 minute restriction
		Crawl automatically every one-two hours
		Store Information

		Project Meridian - Server Side Crawler
		Project Zen - Client Side GUI
		Project Ethereal - Google Cloud Platform Database

		Create Client to Handle Information
		UI to display new struct
		Support for other news sites, maybe handler (espn handler, bleacher report handler, etc)
*/

func main() {
	//Initialize Global client, We want to reuse the http client rather than constantly recreate it
	globalHttpClient := &http.Client{
		Timeout: time.Second * 20,
	}

	fmt.Println("Project Meridian (Server Crawler): V 0.20 Alpha")
	text := ""

	//TODO: Make AUTORUN Mode and Learn how to accept request from sockets/packet
	for text != "Q" {
		fmt.Print("Enter Source ID or Q to Quit: ")
		fmt.Scanln(&text)

		GrabArticlesFromSourceId(text, globalHttpClient)
	}
}

//Obsolete for now
func GrabArticlesFromSourceId(sourceType string, globalHttpClient *http.Client) {
	espnSource := NewESPN() //Might figure out how to change this in the future
	//Possiblity make espnRetrieveWebContent function static
	content := ""

	if sourceType == "1" {
		fmt.Println("Retrieving ESPN Data")
		content = espnSource.RetrieveSearchContent(globalHttpClient)

		//Will seperate this into source specific parsing soon
	}

	espnSource.RunSellerRegexes(content)
}
