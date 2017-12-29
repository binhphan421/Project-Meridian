**Project Meridian**

* Progress: 
    * Seperation of content into actual article structs
* ToDo: 
    * Add comments for all structs and functions
    * Allow server to receive packets from client (for user requested updates)
    * Server will recognize when time limit (1 - 2 hrs) have been reached and execute crawl for content
        * This will be expanded when additional sources come in and we have source specific recrawl times
    * Set global (will be source specific in the future) recrawl time limit (30 minutes)
* Goal: 
    * Create a crawler that will reside on a server
        * 30 minute crawl limit on request
        * Recrawl every 1-2 hrs
        * Save article meta data to Google DB Platform (Ethereal)
        * Robust crawling features: Should be able to imitate as closely to a browser as possible. (Will be important when
            hitting sites with no api)
            * ex. Adjustable Headers
        * DESIGN. Solving the problem is good! But it is better when you solve the problem smartly! Think out what you want
            now and what you might want in the future!
    * Create Client (Zen) to Pull Articles from DB (Ethereal)
        * Maybe use DB to store sourceRegexes, sourceLinks, to allow for updates to retrieval/parsing without rebuilding
        * Crawler should just be a crawler!
    * GoLangUI
        * Research, if and what are GoLang UI libraries
    * Squirrels.