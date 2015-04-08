package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rakyll/pb"
)

var (
	skus = []string{"064cffff790f", "0ea6c810bae1", "11e8e54ec344", "12777436360a", "13d99a1c5a10", "16cdeead9d18", "1ba2c9316eb6",
		"1c72c1fed137", "203bc5cd1961", "21313ed99f3f", "236aabadbf01", "23acbbe70644", "255228890172", "2681c7e225e3",
		"28c5ae44f6ce", "2c7c395d39ea", "2cb7fd4f4158", "30abffbf46f4", "33dd1a9416d9", "348feeab63dc", "349878099841",
		"375174d3c855", "37de10a42308", "38fe41d7641e", "3c6fd8b56ea2", "3ce173fd93a3", "450b9d4f015a", "464294e147c6",
		"4a90ca691194", "4acf690cdd5c", "4fa80381ce10", "5115c9ab36b8", "52ecc75317e4", "53757fd7d132", "54f42df9d1af",
		"55a6e0d0084c", "55f0196d5602", "5948a7e2a5f4", "5effe4083fe7", "644235071e0d", "658da9646f79", "66443ce091b1",
		"6a1f03de38a6", "6aa330dcb7b4", "6dd0e529d859", "6e76d82047e7", "72439090a281", "76a57126b913", "795497bb97df", "79be94825a05"}
	skuPrice    = []float64{2709.29, 3807.43, 1619.53, 9968.41, 9947.60, 9086.92, 6113.40, 8858.29, 1779.91, 7142.04, 8383.05, 9365.44, 729.28, 5225.00, 8010.68, 5338.60, 8327.74, 7543.48, 581.25, 1247.40, 9516.05, 3537.26, 6350.70, 6318.25, 314.23, 7257.00, 4245.51, 1855.51, 798.45, 1691.63, 521.44, 654.73, 9970.34, 5890.25, 565.37, 9153.11, 933.74, 2271.48, 2167.67, 5292.71, 1262.41, 4023.65, 9114.24, 2466.03, 1827.41, 2990.55, 7608.13, 1950.14, 1587.29, 8627.01}
	skuCount    = 50
	startDate   = time.Date(2012, time.January, 1, 0, 0, 0, 0, time.Local)
	endDate     = time.Date(2014, time.December, 31, 0, 0, 0, 0, time.Local)
	priceGrowth = 5.0
	dateLayout  = "2006-01-02"
	writer      *bufio.Writer
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var (
		count              int
		volumePerYear      int
		volumePerNormalDay int
		volumePerHoliday   int
		volumePerDay       int
		holidayMarkdown    int
		sku                string
		price              float64
		r                  int
	)
	flag.IntVar(&count, "count", 40000000, "Number of records to generate")
	flag.Parse()
	// Progress bar
	bar := pb.StartNew(count)

	// Open & defer close the data file
	fileName := time.Now().Format("mockdata.20060102.150405.csv")
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to open file.", err)
		os.Exit(1)
	}
	defer func() {
		if err = outFile.Close(); err != nil {
			fmt.Println("Failed to close  file.", err)
			os.Exit(1)
		}
	}()
	// Init bufio writer
	writer = bufio.NewWriter(outFile)

	// Initial value counting 10% Volume growth YOY
	volumePerYear = int(float64(count) * 0.27)

	checkYear := 0
	for dt := startDate; endDate.Sub(dt) > 0; dt = dt.AddDate(0, 0, 1) {

		// Increment Every year
		if dt.Year() != checkYear {
			volumePerYear = volumePerYear + volumePerYear*10/100
			// Assumption: 20 days of holiday sales account for 10 percentage of total sales volume
			holidayMarkdown = volumePerYear * 10 / 100
			// Normal sales day volume
			volumePerNormalDay = (volumePerYear - holidayMarkdown) / 345
			// Holiday volume
			volumePerHoliday = volumePerYear / 200
			// Update price by "priceGrowth" percentage
			for key, val := range skuPrice {
				skuPrice[key] = val + val*priceGrowth/100
			}
			checkYear = dt.Year()
		}
		// Check if current day is Holiday
		if (dt.Month() == time.November && dt.Day() > 15 && dt.Day() < 30) || (dt.Month() == time.December && dt.Day() > 15 && dt.Day() < 30) {
			volumePerDay = volumePerHoliday
		} else {
			volumePerDay = volumePerNormalDay
		}

		for i := 0; i < volumePerDay; i++ {
			r = rand.Intn(49)
			sku = skus[r]
			price = skuPrice[r]
			fmt.Fprintf(writer, "%s,%s,%.2f\n", sku, dt.Format(dateLayout), price)
			bar.Increment()
		}
	}
	writer.Flush()
	bar.FinishPrint("Data Generated!")
}
