package coupons

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

var LoadCoupons = func(gzFile string) string {
	file, err := os.Open(gzFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("loading coupons from", gzFile)

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer gzReader.Close()

	content, err := io.ReadAll(gzReader)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
