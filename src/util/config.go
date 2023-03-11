package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func createIfNotExists(file string) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		file, err := os.Create(file)
		if err != nil {
			fmt.Println("Error creating file: ", err, " at ", file)
			fmt.Println("You may need to run this command as root (sudo).")
			os.Exit(1)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)
	}
}

func SaveToTemp(key string, value string) {
	// saves the value to a temporary .yml file
	ymlFile := GetTempFile()
	// create the file if it doesn't exist
	createIfNotExists(ymlFile)

	// read in the file
	data, err := os.ReadFile(ymlFile)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}

	// Update the map with the new key-value pair
	m[key] = value

	yamlData, err := yaml.Marshal(&m)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(ymlFile, yamlData, 0644)
	if err != nil {
		panic(err)
	}

}

func GetFromTemp(key string) string {
	// gets the value from a temporary .yml file
	ymlFile := GetTempFile()

	// read the file
	data, err := os.ReadFile(ymlFile)
	if err != nil {
		panic(err)
	}

	// unmarshal the file
	var m map[string]interface{}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	} else {
		return m[key].(string)
	}
}
