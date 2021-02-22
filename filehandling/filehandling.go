package filehandling

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var stringArray []string

func Main() {
	readTestAbsolut()
	readTestFlag()
	readTestChunks()
	readTestLinebyLine()
	printArray()
}

func readTestAbsolut() {
	data, err := ioutil.ReadFile("filehandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readTestFlag() {
	fptr := flag.String("fpath2", "filehandling/test.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readTestChunks() {
	fptr := flag.String("fpath", "filehandling/test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

func readTestLinebyLine() []string {
	fptr := flag.String("fpath3", "filehandling/test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
		stringArray = append(stringArray, s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return stringArray
}
func printArray() {
	// Print Array of Strings
	fmt.Println(stringArray)
	// Print Array of Strings Linewise
	fmt.Println((strings.Join(stringArray, "\n")))
}
func GiveFileAsStringArray() []string {
	return stringArray
}
