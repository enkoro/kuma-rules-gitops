package rules

import "io"

type Enrichment struct {
	APIVersion string `yaml:"apiVersion" bson:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" bson:"kind" json:"kind"`
	Metadata   struct {
		Name      string `yaml:"name" bson:"name" json:"name"`
		Namespace string `yaml:"namespace" bson:"namespace" json:"namespace"`
		Labels    struct {
			RouterDeisIoRoutable string `yaml:"router.deis.io/routable" bson:"router.deis.io/routable" json:"router.deis.io/routable"`
		} `yaml:"labels" bson:"labels" json:"labels"`
		Annotations struct {
			RouterDeisIoDomains string `yaml:"router.deis.io/domains" bson:"router.deis.io/domains" json:"router.deis.io/domains"`
		} `yaml:"annotations" bson:"annotations" json:"annotations"`
	} `yaml:"metadata" bson:"metadata" json:"metadata"`
	Spec struct {
		Type     string `yaml:"type" bson:"type" json:"type"`
		Selector struct {
			App string `yaml:"app" bson:"app" json:"app"`
		} `yaml:"selector" bson:"selector" json:"selector"`
		Ports []struct {
			Name       string `yaml:"name" bson:"name" json:"name"`
			Port       int    `yaml:"port" bson:"port" json:"port"`
			TargetPort int    `yaml:"targetPort" bson:"targetPort" json:"targetPort"`
		} `yaml:"ports" bson:"ports" json:"ports"`
	} `yaml:"spec" bson:"spec" json:"spec"`
}

func (r *Enrichment) FromYaml(yamlBs []byte) error {
	err := initFromYaml(yamlBs, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Enrichment) WriteTo(w io.Writer, password string) error {
	err := write(w, r, password)
	if err != nil {
		return err
	}
	return nil
}
