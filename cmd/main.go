package main

import (
	"fmt"
	kuma "kuma-rules-gitops/resources"
	"os"
)

func main() {
	enrichment := kuma.NewRuleEnrichment()
	chtototam := kuma.NewRuleChtototam()
	inputYaml, err := os.ReadFile("test_input.yaml")
	if err != nil {
		panic(err)
	}
	if err = enrichment.FromYaml(inputYaml); err != nil {
		panic(err)
	}
	if err = chtototam.FromYaml(inputYaml); err != nil {
		panic(err)
	}
	encodedChtototam, err := chtototam.Encode("")
	if err != nil {
		panic(err)
	}
	encodedEnrichment, err := enrichment.Encode("")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encodedChtototam))
	fmt.Println(string(encodedEnrichment))
}
