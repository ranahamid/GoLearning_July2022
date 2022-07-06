package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferWriter := bufio.NewWriter(file)
	bs := []byte("I LOVE GO Programming")
	bytesWritten, err := bufferWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)
	bytesAvailable := bufferWriter.Available()
	log.Printf("bytesAvailable: %d\n", bytesAvailable)

	bytesWritten, err = bufferWriter.WriteString("Now its writing")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	bufferWriter.Flush()

	//READ
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data is: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
	data, err = ioutil.ReadFile("main.go")
	fmt.Printf("Data is: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}
