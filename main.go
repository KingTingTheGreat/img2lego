package main

import (
	"log"

	"github.com/kingtingthegreat/img2lego/server"
)

func main() {
	// // open image
	// f, err := os.Open("LightsaberCSWE.webp")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	//
	// img, _, err := image.Decode(f)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // scale image with converter
	// img = resize.Resize(100, 0, img, resize.Lanczos3)
	// kittyimg.Fprintln(os.Stdout, img)
	//
	// // convert image
	// output := converter.ConvertImage(img)
	//
	// // print out converted image
	// output.PrintRGBGrid(output)

	server := server.Server()
	log.Println("Starting server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
