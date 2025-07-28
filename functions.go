package main

import (
	"encoding/json"
	"os"
)

func createDecodedMap(path string) (map[string]chapter, error) {
	var decodedStory map[string]chapter
	body, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	err = json.NewDecoder(body).Decode(&decodedStory)
	if err != nil {
		return nil, err
	}
	return decodedStory, nil
}