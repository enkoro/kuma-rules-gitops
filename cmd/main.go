package main

import (
	"kuma-rules-gitops/rules"
	"os"
)

func main() {

	var testRule rules.Enrichment
	testInput, err := os.ReadFile("test_input.yaml")
	if err != nil {
		panic(err)
	}
	err = testRule.FromYaml(testInput)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("out", os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	testRule.WriteTo(file, "qwe")
	testRule.WriteTo(os.Stdout, "")
}
