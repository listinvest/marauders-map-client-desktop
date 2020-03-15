package ostools

import (
	"io"
	"log"
	"os"
)

func copy(from string, to string) {
	ffrom, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer ffrom.Close()

	fto, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fto.Close()

	_, err = io.Copy(fto, ffrom)
	if err != nil {
		log.Fatal(err)
	}
}
