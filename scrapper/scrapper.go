package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const FileName = "jobs.csv"

type jobInfo struct {
	id         string
	badge      string
	title      string
	corpName   string
	jobDate    string
	conditions []string
	jobSector  string
}

func Run(searchWord string) {
	fmt.Println("Scraping... ", searchWord)
	url := fmt.Sprintf("https://www.saramin.co.kr/zf_user/search?searchType=search&searchword=%s", searchWord)

	pageCount := getPageCount(url)
	if pageCount == nil {
		return
	}

	jobs := getPage(url, 1)
	makeCsv(jobs)
}

func makeCsv(jobs []jobInfo) {
	file, err := os.Create(FileName)

	if err != nil {
		log.Fatalln(err)
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "badge", "title", "corpName", "jobDate", "conditions", "jobSector", "apply Link"}
	err = writer.Write(headers)

	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, job := range jobs {
		applyLink := "https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id
		writer.Write([]string{job.id, job.badge, job.title, job.corpName, job.jobDate, strings.Join(job.conditions, ","), job.jobSector, applyLink})
	}

	fmt.Println("CSV 파일이 성공적으로 생성되었습니다.")
}

func getPage(baseUrl string, pageIndex int) []jobInfo {
	pageUrl := baseUrl + "&recruitPage=" + strconv.Itoa(pageIndex) + "&recruitPageCount=40"
	response, err := http.Get(pageUrl)

	if !checkRequestError(err) || !checkResponseCode(response) {
		return []jobInfo{}
	}

	doc, docError := goquery.NewDocumentFromReader(response.Body)

	if !checkRequestError(docError) {
		return []jobInfo{}
	}

	channel := make(chan jobInfo)
	jobs := []jobInfo{}

	items := doc.Find(".item_recruit")
	items.Each(func(i int, s *goquery.Selection) {
		go extractJobInfo(s, channel)
	})

	for i := 0; i < items.Length(); i++ {
		jobInfo := <-channel
		jobs = append(jobs, jobInfo)
	}

	return jobs
}

func extractJobInfo(s *goquery.Selection, channel chan<- jobInfo) {
	id := s.AttrOr("value", "")
	badge := s.Find(".badge").Text()
	title, _ := s.Find(".job_tit").Find("a").Attr("title")
	corpName := s.Find(".corp_name").Text()
	jobDate := s.Find(".job_date > .date").Text()

	conditions := []string{}
	s.Find(".job_condition").Find("span").Each(func(i int, s *goquery.Selection) {
		conditions = append(conditions, CleanString(s.Text()))
	})

	jobSector := s.Find(".job_sector").Text()

	jobInfo := jobInfo{
		id:         id,
		badge:      CleanString(badge),
		title:      CleanString(title),
		corpName:   CleanString(corpName),
		jobDate:    CleanString(jobDate),
		conditions: conditions,
		jobSector:  CleanString(jobSector),
	}
	channel <- jobInfo
}

func CleanString(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\f", "")
	str = strings.ReplaceAll(str, "\v", "")
	str = strings.ReplaceAll(str, "\b", "")
	str = strings.ReplaceAll(str, "\a", "")
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func getPageCount(url string) *int {
	pageCount := 0
	response, err := http.Get(url)

	if !checkRequestError(err) || !checkResponseCode(response) {
		return nil
	}

	defer response.Body.Close()

	doc, docError := goquery.NewDocumentFromReader(response.Body)

	if !checkRequestError(docError) {
		return nil
	}

	pagination := doc.Find(".pagination")
	pageCount = pagination.Find("a").Length()

	return &pageCount
}

func checkRequestError(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func checkResponseCode(response *http.Response) bool {
	if response.StatusCode != 200 {
		log.Fatalln("Request failed with status code : ", response.StatusCode)
		return false
	}

	return true
}
