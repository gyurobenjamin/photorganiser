package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gyurobenjamin/photorganiser/confirm"
)

var (
	srcDir  *string
	destDir string
	cursor  int
)

func init() {
	srcDir = flag.String("dir", "", "Location...")
	cursor = 0
}

func main() {
	flag.Parse()

	if *srcDir == "" {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		srcDir = &dir
	}

	destDir = *srcDir + "_photorganiser"

	fmt.Println()
	fmt.Println("We're going to process your files in this directory:", *srcDir)
	fmt.Println("New location:", destDir)
	fmt.Println("Is this the correct location?")
	fmt.Println()

	if confirm.AskForConfirmation() == false {
		fmt.Println("Bye...")
		os.Exit(1)
	}
	fmt.Println("Process is running...")

	processDir("")
}

func processDir(dir string) {
	src := *srcDir + "/" + dir
	dest := *srcDir + "_photograniser/" + dir

	fmt.Println("")
	fmt.Println("Process directory:", src)
	fmt.Println("Creating dest directory: " + dest)

	os.Mkdir(dest, 0777)

	files, err := ioutil.ReadDir(*srcDir + dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() == true {
			processDir(dir + "/" + f.Name())
		} else {
			processFile(dir, f.Name())
		}
	}
}

func processFile(dir string, file string) {
	if cursor > 10 {
		os.Exit(1)
	}
	src := *srcDir + dir + "/" + file

	fmt.Println()
	fmt.Println(cursor, "Process file:", src)

	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("Error ...", err)
	}

	info, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error ...", err)
	}

	dest := *srcDir + "_photograniser/" + dir + "/" + info.ModTime().Format("2006-01-02") + "_" + file
	fmt.Println("New name:", dest)

	err = ioutil.WriteFile(dest, data, 0644)
	if err != nil {
		fmt.Println("Error ...", err)
	}

	err = os.Chtimes(dest, info.ModTime(), info.ModTime())
	if err != nil {
		fmt.Println("Error ...", err)
	}

	cursor++
}

func processImage() {

}
