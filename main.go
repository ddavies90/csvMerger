package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type item struct {
	id           string
	size         string
	label        string
	availability string
}

func main() {
	items := map[string]item{}
	file1, err := os.Open("./fileSet1/file1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	csvReader := csv.NewReader(file1)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(data); i++ {
		itemData := item{id: data[i][0], size: data[i][1], label: data[i][2]}
		items[itemData.id] = itemData
	}

	file2, err := os.Open("./fileSet1/file2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	csvReader = csv.NewReader(file2)
	data, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(data); i++ {
		itemData := item{id: data[i][0], availability: data[i][1]}
		if val, ok := items[itemData.id]; ok {
			fmt.Println(val)
			val.availability = itemData.availability
		} else {
			items[itemData.id] = itemData
		}
	}

	//set testItem to the result of searching the map for the value passed in to the [] which is a unique ID - should
	//return a single record. Append availability property on to this record.
	//testItem := items["fff94c24-2700-4332-af22-f423091100a9"]
	//testItem.availability = "Available"

	//Search items map to find the record that has the ID that testItem has and set it to the updated testItem, as we
	//just amended it by adding availability.
	//items[testItem.id] = testItem

	csvfile, err := os.Create("MergedCSV.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvwriter := csv.NewWriter(csvfile)

	//fmt.Println(len(items))

}
