package main

import (
	"encoding/json"
	"os"
)

type Layer struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type Tilemap struct {
	Tiles []Layer `json:"layers"`
}

func LoadTilemap(filepath string) (*Tilemap, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	tilemap := &Tilemap{}
	err = json.Unmarshal(contents, tilemap)
	if err != nil {
		return nil, err
	}
	return tilemap, nil
}
