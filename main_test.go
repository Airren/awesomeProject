package awesomeProject

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGENE(t *testing.T) {
	fmt.Println(">>>>>>>>>> start the work <<<<<<<<<<<<<")
	startAt := time.Now()
	file, err := os.Open("./result.gtf")
	res, err := os.Create(fmt.Sprintf("res_%v_.txt", time.Now().Unix()))
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		err = res.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	beforeStr := ""
	beforeWrite := false
	tempList := make([]string, 0)
	for scanner.Scan() {
		elemStr := scanner.Text()
		elemType, elemKey := getKeyType(elemStr)
		beforeType, beforeKey := getKeyType(beforeStr)
		if elemKey == beforeKey {
			if len(tempList) == 0 {
				tempList = append(tempList, beforeStr)
			}
			tempList = append(tempList, elemStr)
			if elemType != beforeType {
				beforeWrite = true
			}
		} else {
			if beforeWrite {
				writeArray(tempList, res)
				beforeWrite = false
			}
			tempList = make([]string, 0)
		}

		beforeStr = elemStr
	}

	if beforeWrite {
		writeArray(tempList, res)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	timeCost := time.Now().Sub(startAt).String()
	log.Printf("It takes >>>>> %s <<<<<\n", timeCost)
}

func getKeyType(elemStr string) (string, string) {
	elems := strings.Split(elemStr, "\t")
	if len(elems) > 8 {
		return elems[2], elems[8]
	}
	return "", ""
}

func writeArray(arr []string, f *os.File) {
	for _, str := range arr {
		_, err := f.WriteString(fmt.Sprintf("%s\n", str))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
