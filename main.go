package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func main() {
	var (
		basicInfo     BasicInfo
		basicInfoFile string

		salesInfo     SalesInfo
		salesInfoFile string

		shippingInfo     ShippingInfo
		shippingInfoFile string

		mediaInfo     MediaInfo
		mediaInfoFile string

		massUploadFile string
	)

	var wg sync.WaitGroup

	source, err := ioutil.ReadDir("source")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range source {
		if strings.Contains(file.Name(), ".xlsx") {
			if strings.Contains(file.Name(), "basic_info") {
				basicInfoFile = "source/" + file.Name()
			}
			if strings.Contains(file.Name(), "sales_info") {
				salesInfoFile = "source/" + file.Name()
			}
			if strings.Contains(file.Name(), "shipping_info") {
				shippingInfoFile = "source/" + file.Name()
			}
			if strings.Contains(file.Name(), "media_info") {
				mediaInfoFile = "source/" + file.Name()
			}
		}
	}

	wg.Add(5)

	go getBasicInfo(&basicInfo, basicInfoFile, &wg)
	go getSalesInfo(&salesInfo, salesInfoFile, &wg)
	go getShippingInfo(&shippingInfo, shippingInfoFile, &wg)
	go getMediaInfo(&mediaInfo, mediaInfoFile, &wg)

	go func() {
		template, err := ioutil.ReadDir("templates")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range template {
			if file.Name()[0] != '.' {
				if strings.Contains(file.Name(), "mass_upload") {
					massUploadFile = "templates/" + file.Name()
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()

	MergeInfo(basicInfo, salesInfo, shippingInfo, mediaInfo, massUploadFile)
}
