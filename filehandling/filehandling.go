package filehandling

import (
	"fmt"
	"io/ioutil"
)

func Main() {
	data, err := ioutil.ReadFile("F:/ASH_Programmierung/20210215/datavis/filehandling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
