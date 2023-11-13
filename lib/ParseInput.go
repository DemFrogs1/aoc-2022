package ParseInput

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Parse(FileName string){
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}