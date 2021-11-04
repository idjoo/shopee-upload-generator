package main

import (
	"fmt"
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

	fmt.Println("Fetching basic info...")
	go getBasicInfo(&basicInfo, basicInfoFile, &wg)
	fmt.Println("Fetching sales info...")
	go getSalesInfo(&salesInfo, salesInfoFile, &wg)
	fmt.Println("Fetching shipping info...")
	go getShippingInfo(&shippingInfo, shippingInfoFile, &wg)
	fmt.Println("Fetching media info...")
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
	fmt.Println("Fetching info done!")

	fmt.Println("Writing to template file...")
	MergeInfo(basicInfo, salesInfo, shippingInfo, mediaInfo, massUploadFile)
	fmt.Println("Finished!")
}
