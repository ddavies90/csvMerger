package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	a := map[string]item{}
	//i := item{id: "2wqerqjo1-21141", size: "XL", label: "Extra Large"}
	f1, err := os.Open("./fileSet1/file1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	csvReader := csv.NewReader(f1)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(data); i++ {
		itemData := item{id: data[i][0], size: data[i][1], label: data[i][2]}
		a[itemData.id] = itemData
	}

	fmt.Println(a)

}

type item struct {
	id           string
	size         string
	label        string
	availability string
}
