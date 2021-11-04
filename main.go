package main

import (
	"io/ioutil"
	"log"
	"strings"
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

	source, err := ioutil.ReadDir("source")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range source {
		if file.Name()[0] != '.' {
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

	getBasicInfo(&basicInfo, basicInfoFile)
	getSalesInfo(&salesInfo, salesInfoFile)
	getShippingInfo(&shippingInfo, shippingInfoFile)
	getMediaInfo(&mediaInfo, mediaInfoFile)

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

	// pretty.Println(salesInfo.Items[0])

	MergeInfo(basicInfo, salesInfo, shippingInfo, mediaInfo, massUploadFile)
}
