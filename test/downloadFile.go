package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	fileName string

	fullUrlPath string
)

func download() {
	fullUrlPath = "http://f.video.weibocdn.com/u0/WAAZHCt0gx07UeZzNjbW010412010xkn0E010.mp4?label=mp4_ld&template=638x360.25.0&trans_finger=40a32e8439c5409a63ccf853562a60ef&media_id=4743676350627926&tp=8x8A3El:YTkl0eM8&us=0&ori=1&bf=2&ot=h&lp=00001r4HGF&ps=mZ6WB&uid=6LuS7o&ab=7397-g1,6377-g0,1192-g0,1191-g0,1258-g0,3601-g19&Expires=1659162572&ssig=gzcx9fkND4&KID=unistore,video"
	fileName = "test6.mp4"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}
	resp, err := client.Get(fullUrlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
