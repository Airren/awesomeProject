package awesomeProject

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestGENEVerify(t *testing.T) {
	fmt.Println(">>>>>>>>>> start the work <<<<<<<<<<<<<")
	startAt := time.Now()
	file, err := os.Open("./result.gtf")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	elemMap := make(map[string]string)
	resMap := make(map[string]struct{})
	countMap := make(map[string]int)
	for scanner.Scan() {
		elemStr := scanner.Text()
		elemType, elemKey := getKeyType(elemStr)

		if t, ok := elemMap[elemKey]; !ok {
			elemMap[elemKey] = elemType
			countMap[elemKey] = 1
		} else {
			countMap[elemKey] += 1
			if t != elemType {
				resMap[elemKey] = struct{}{}
			}
		}
	}

	count := 0
	for k, _ := range resMap {
		count += countMap[k]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

	timeCost := time.Now().Sub(startAt).String()
	log.Printf("It takes >>>>> %s <<<<<\n", timeCost)
}
