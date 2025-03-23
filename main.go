package main

import (
	"my-study-go/scrapper"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func mainPage(c echo.Context) error {
	return c.File("main.html")
}

func postScrape(c echo.Context) error {
	searchWord := c.FormValue("searchWord")
	searchWord = scrapper.CleanString(searchWord)

	if searchWord == "" {
		return c.String(http.StatusBadRequest, "Search word is required")
	}

	scrapper.Run(searchWord)

	if _, err := os.Stat(scrapper.FileName); os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "File not found")
	}

	defer os.Remove(scrapper.FileName)
	return c.Attachment(scrapper.FileName, scrapper.FileName)
}

func main() {
	e := echo.New()
	e.GET("/", mainPage)
	e.POST("/scrape", postScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
