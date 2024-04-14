package rules

import (
	"encoding/json"
	"io"
	"kuma-rules-gitops/tools"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
)

type kumaRule interface {
	FromYaml([]byte) error
	WriteTo(io.Writer, string) error
}

func write(w io.Writer, rule kumaRule, password string) error {
	if password == "" {
		jsDoc, err := json.Marshal(rule)
		if err != nil {
			return err
		}
		_, err = w.Write(jsDoc)
		return err
	} else {
		bsDoc, err := bson.Marshal(rule)
		if err != nil {
			return err
		}
		encrypted, err := tools.Enrypt(bsDoc, []byte("qwe"))
		if err != nil {
			panic(err)
		}
		_, err = w.Write(encrypted)
		return err
	}
}

func initFromYaml(yamlBs []byte, rule kumaRule) error {
	err := yaml.Unmarshal(yamlBs, rule)
	if err != nil {
		return err
	}
	return nil
}
