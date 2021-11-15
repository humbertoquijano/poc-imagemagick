package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	imgkcaller "mti.com/go-imgk/imgk-caller"
)

func main() {
	var wg sync.WaitGroup
	batDir := os.Args[1]

	start := time.Now()

	files, err := ioutil.ReadDir(batDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".bat" {
			wg.Add(1)
			go imgkcaller.MakeCall(batDir+file.Name(), &wg)
		}
	}

	wg.Wait()

	end := time.Now()
	result := end.Sub(start)
	log.Println(result)
}
