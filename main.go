package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func parse_url(url string) (*goquery.Document, error) {
	// Getリクエスト
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// 読み取り
	buf, _ := io.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	fmt.Println(detRslt.Charset)

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func getTitle(doc *goquery.Document) string {
	return doc.Find("title").Text()
}

func main() {
	url := os.Args[1]

	doc, err := parse_url(url)
	if err != nil {
		log.Fatal(err)
	}

	title := getTitle(doc)

	fmt.Println(title)
}
