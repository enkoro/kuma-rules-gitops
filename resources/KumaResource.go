package resources

import (
	"kuma-rules-gitops/tools"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
)

type iResource interface {
	NewResource() *Resource
}

type Resource struct {
	Resource iResource
}

func (r *Resource) FromYaml(yamlBs []byte) error {
	err := yaml.Unmarshal(yamlBs, r.Resource)
	if err != nil {
		return err
	}
	return nil
}

func (r *Resource) Encode(password string) ([]byte, error) {
	bsDoc, err := bson.Marshal(r.Resource)
	if err != nil {
		return nil, err
	}
	if password != "" {
		encrypted, err := tools.Enrypt(bsDoc, []byte(password))
		if err != nil {
			return nil, err
		}
		return encrypted, nil
	}
	return bsDoc, nil
}
