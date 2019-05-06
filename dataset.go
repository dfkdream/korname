package korname

import (
	"encoding/json"
	"io/ioutil"
)

//https://github.com/agemor/korean-name-generator/blob/master/src/trained-data.js

type dataset struct {
	LastName        []int
	LastNameFreq    []int
	MaleFirstName   [][][]int
	FemaleFirstName [][][]int
}

func loadDataset(filepath string) (dataset, error) {
	var set dataset
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return dataset{}, err
	}
	json.Unmarshal(data, &set)
	return set, nil
}
