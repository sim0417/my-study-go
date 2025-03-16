package main

import (
	"errors"
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.google.com",
	"https://www.naver.com",
	"https://www.daum.net",
	"https://www.youtube.com",
	"https://www.facebook.com",
	"https://www.twitter.com",
	"https://www.instagram.com",
	"https://www.amazon.com",
}

type Response struct {
	URL        string
	StatusCode string
}

var errorRequestFailed = errors.New("request failed")

func main() {
	// normalProcess()
	concurrentProcess()
}

func normalProcess() {
	var results = map[string]string{}

	for _, url := range urls {
		result := "PASS"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	fmt.Println(results)
}

func hitUrl(url string) error {
	fmt.Println("hitUrl start:", url)
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		return errorRequestFailed
	}

	return nil
}

// go routine 을 사용하여 동시에 여러 개의 요청을 보내고 결과를 받아온다.
// 채널은 고 루린과 실행한 함수를 연결하는 파이프 역할을 한다.
func concurrentProcess() {
	results := make(map[int]Response)
	channel := make(chan Response)
	// 수신용 채널은 아래처럼 선언할 수 있다.
	// channel := make(<-chan Response)
	// 송신용 채널은 아래처럼 선언할 수 있다.
	// channel := make(chan<- Response)

	for _, url := range urls {
		// 함수 앞에 go 키워드를 붙이면 함수를 동시에 실행할 수 있다.
		go hitUrlRoutine(url, channel)
	}

	// 채널을 통해서 결과를 받아온다.
	// 고 루틴은 동시에 실행되기 때문에 순서대로 결과를 받아오지 않는다.
	for index := range urls {
		results[index] = <-channel
	}
	fmt.Println(results)
}

func hitUrlRoutine(url string, channel chan<- Response) {
	response, err := http.Get(url)
	var status string

	if err != nil || response.StatusCode >= 400 {
		status = fmt.Sprintf("FAILED, %d", response.StatusCode)
	} else {
		status = fmt.Sprintf("PASS, %d", response.StatusCode)
	}

	channel <- Response{url, status}
}
