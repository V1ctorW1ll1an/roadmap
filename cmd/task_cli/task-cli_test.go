package main

import "testing"

func TestReadJsonFile(t *testing.T) {
	path := "../../tasks.json"
	ReadJson(path)
}
