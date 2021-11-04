package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/xuri/excelize/v2"
)

func getBasicInfo(basicInfo *BasicInfo, basicInfoFile string, wg *sync.WaitGroup) {
	xlsx, err := excelize.OpenFile(basicInfoFile)
	if err != nil {
		panic(err)
	}

	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i < 3 {
			continue
		}

		var item BasicInfoItem

		item.KodeProduk = row[0]
		item.SkuInduk = row[1]
		item.NamaProduk = row[2]
		item.DeskripsiProduk = row[3]

		basicInfo.AddItem(item)
	}
	wg.Done()
}

func getSalesInfo(salesInfo *SalesInfo, salesInfoFile string, wg *sync.WaitGroup) {
	xlsx, err := excelize.OpenFile(salesInfoFile)
	if err != nil {
		panic(err)
	}

	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i < 4 {
			continue
		}

		var item SalesInfoItem

		item.KodeProduk = row[0]
		item.NamaProduk = row[1]
		item.KodeVariasi = row[2]
		item.NamaVariasi = row[3]
		item.SkuInduk = row[4]
		item.Sku = row[5]
		item.Harga = row[6]
		item.Stok = row[7]

		salesInfo.AddItem(item)
	}
	wg.Done()
}

func getShippingInfo(shippingInfo *ShippingInfo, shippingInfoFile string, wg *sync.WaitGroup) {
	xlsx, err := excelize.OpenFile(shippingInfoFile)
	if err != nil {
		panic(err)
	}

	rows, err := xlsx.GetRows("the sku info")
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i < 5 {
			continue
		}

		var item ShippingInfoItem

		item.KodeProduk = row[0]
		item.SkuInduk = row[1]
		item.NamaProduk = row[2]
		item.BeratProduk = row[3]
		item.Panjang = row[4]
		item.Lebar = row[5]
		item.Tinggi = row[6]

		shippingInfo.AddItem(item)
	}
	wg.Done()
}

func getMediaInfo(mediaInfo *MediaInfo, mediaInfoFile string, wg *sync.WaitGroup) {
	xlsx, err := excelize.OpenFile(mediaInfoFile)
	if err != nil {
		panic(err)
	}

	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i < 5 {
			continue
		}

		var item MediaInfoItem

		item.KodeProduk = row[0]
		item.SkuInduk = row[1]
		item.NamaProduk = row[2]
		item.Kategori = row[3]
		item.FotoSampul = row[4]

		for j, col := range row {
			if j == 14 {
				item.NamaVariasi = col
			}

			if j >= 5 && j <= 12 {
				item.FotoProduk = append(item.FotoProduk, col)
			} else if j >= 15 && j <= 54 {
				if j%2 == 1 {
					item.Variasi = append(item.Variasi, col)
				} else if j%2 == 0 {
					item.FotoVariasi = append(item.FotoVariasi, col)
				}
			} else {
				continue
			}
		}

		mediaInfo.AddItem(item)
	}
	wg.Done()
}

func MergeInfo(basicInfo BasicInfo, salesInfo SalesInfo, shippingInfo ShippingInfo, mediaInfo MediaInfo, massUploadFile string) {
	xlsx, err := excelize.OpenFile(massUploadFile)
	if err != nil {
		panic(err)
	}

	temp := ""
	sheet := "Template"
	kodeIntegrasiVariasi := 1
	v := 0
	for i, sales := range salesInfo.Items {
		xlsx.SetCellValue(sheet, fmt.Sprintf("A%d", i+6), "")
		xlsx.SetCellValue(sheet, fmt.Sprintf("B%d", i+6), sales.NamaProduk)
		for _, basic := range basicInfo.Items {
			if sales.KodeProduk == basic.KodeProduk {
				xlsx.SetCellValue(sheet, fmt.Sprintf("C%d", i+6), basic.DeskripsiProduk)
			}
		}
		xlsx.SetCellValue(sheet, fmt.Sprintf("D%d", i+6), sales.SkuInduk)
		xlsx.SetCellValue(sheet, fmt.Sprintf("E%d", i+6), "")

		if temp != sales.KodeProduk && sales.NamaVariasi != "" {
			kodeIntegrasiVariasi += rand.Intn(999)
			v = 0
		}

		if sales.NamaVariasi != "" {
			temp = sales.KodeProduk
			xlsx.SetCellValue(sheet, fmt.Sprintf("F%d", i+6), kodeIntegrasiVariasi)
			xlsx.SetCellValue(sheet, fmt.Sprintf("G%d", i+6), "Varian")
		}

		xlsx.SetCellValue(sheet, fmt.Sprintf("H%d", i+6), sales.NamaVariasi)

		for _, media := range mediaInfo.Items {
			if sales.KodeProduk == media.KodeProduk {
				xlsx.SetCellValue(sheet, fmt.Sprintf("O%d", i+6), media.FotoSampul)
				for j, foto := range media.FotoProduk {
					xlsx.SetCellValue(sheet, fmt.Sprintf("%c%d", j+80, i+6), foto)
				}
				if media.FotoVariasi != nil {
					if media.FotoVariasi[v] == "" {
						xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", i+6), media.FotoSampul)
					} else {
						xlsx.SetCellValue(sheet, fmt.Sprintf("I%d", i+6), media.FotoVariasi[v])
					}

					if v == len(media.FotoVariasi)-1 {
						v = len(media.FotoVariasi) - 1
					} else {
						v++
					}

				}
			}
		}

		xlsx.SetCellValue(sheet, fmt.Sprintf("L%d", i+6), sales.Harga)
		xlsx.SetCellValue(sheet, fmt.Sprintf("M%d", i+6), sales.Stok)
		for _, shipping := range shippingInfo.Items {
			if shipping.KodeProduk == sales.KodeProduk {
				xlsx.SetCellValue(sheet, fmt.Sprintf("X%d", i+6), shipping.BeratProduk)
				/* xlsx.SetCellValue(sheet, fmt.Sprintf("Y%d", i+6), shipping.Panjang)
				xlsx.SetCellValue(sheet, fmt.Sprintf("Z%d", i+6), shipping.Lebar)
				xlsx.SetCellValue(sheet, fmt.Sprintf("AA%d", i+6), shipping.Tinggi) */
			}
		}

		for j := 0; true; j++ {
			test, err := xlsx.GetCellValue(sheet, fmt.Sprintf("A%c%d", j+68, 2))
			if err != nil {
				panic(err)
			}
			xlsx.SetCellValue(sheet, fmt.Sprintf("A%c%d", j+66, i+6), "Aktif")
			if test == "" {
				break
			}
		}
	}

	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}
}
