package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var imageFile string
var text string
var size float64
var scale int

func main() {
	flag.StringVar(&imageFile, "image", "", "图片")
	flag.StringVar(&text, "text", "", "水印文本")
	flag.Float64Var(&size, "size", 12.0, "水印文本大小")
	flag.IntVar(&scale, "scale", 2, "控制水印之间的间距")
	flag.Parse()

	imageBlob, err := ioutil.ReadFile(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	markedImage, err := WaterMark(imageBlob, text, size, scale)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("watermark_%s", filepath.Base(imageFile)), markedImage, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
