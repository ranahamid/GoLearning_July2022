package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	newFile, err = os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}
	err = os.Truncate("a.txt", 0)
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}
	newFile.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
	p("File Name", fileInfo.Name())
	p("File Size", fileInfo.Size())
	p("File ModTime", fileInfo.ModTime())
	p("File IsDir", fileInfo.IsDir())
	p("File Mode", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			//log.Fatal(err)
			//log.Fatal("File is not exists")
			fmt.Println("File is not exists")
		}
	}

	oldPath := "aaa.txt"
	newPath := "aa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	} else {
		fmt.Println("success")
	}

	err = os.Remove("a.txt")
}
