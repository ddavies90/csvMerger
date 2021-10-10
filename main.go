package main

import (
	"encoding/csv"
	"log"
	"os"
)

type item struct {
	id           string
	size         string
	label        string
	availability string
	color        string
}

func main() {
	items := make(map[string]item)
	file1, err := os.Open("./fileSet1/file1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	csvReader := csv.NewReader(file1)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	file1.Close()

	for i := 1; i < len(records); i++ {
		itemData := item{id: records[i][0], size: records[i][1], label: records[i][2]}
		items[itemData.id] = itemData
	}

	file2, err := os.Open("./fileSet1/file2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	csvReader = csv.NewReader(file2)
	records, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	file2.Close()

	for i := 1; i < len(records); i++ {
		itemData := item{id: records[i][0], availability: records[i][1]}
		if val, ok := items[itemData.id]; ok {
			val.availability = itemData.availability
			items[itemData.id] = val
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

	file3, err := os.Open("./fileSet1/file3.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	csvReader = csv.NewReader(file3)
	records, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	file3.Close()

	for i := 1; i < len(records); i++ {
		itemData := item{id: records[i][0], color: records[i][1]}
		if val, ok := items[itemData.id]; ok {
			val.color = itemData.color
			items[itemData.id] = val
		} else {
			items[itemData.id] = itemData
		}
	}

	csvFile, err := os.Create("MergedCSV.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(csvFile)

	var data [][]string
	data = append(data, []string{"id", "size", "label", "availability", "color"})
	for _, itemRecord := range items {
		row := []string{itemRecord.id, itemRecord.size, itemRecord.label, itemRecord.availability, itemRecord.color}
		data = append(data, row)
	}
	err = csvWriter.WriteAll(data)
	if err != nil {
		log.Fatal("Unable to write to CSV due to following error: ", err)
	}

	//fmt.Println(len(items))

}
