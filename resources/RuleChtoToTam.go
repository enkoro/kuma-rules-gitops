package resources

type RuleChtoToTamFields struct {
	Metadata struct {
		Name      string `yaml:"name" bson:"name"`
		Namespace string `yaml:"namespace" bson:"namespace"`
		Labels    struct {
			RouterDeisIoRoutable string `yaml:"router.deis.io/routable" bson:"router.deis.io/routable"`
		} `yaml:"labels" bson:"labels"`
		Annotations struct {
			RouterDeisIoDomains string `yaml:"router.deis.io/domains" bson:"router.deis.io/domains"`
		} `yaml:"annotations" bson:"annotations"`
	} `yaml:"metadata" bson:"metadata"`
	Spec struct {
		Type     string `yaml:"type" bson:"type"`
		Selector struct {
			App string `yaml:"app" bson:"app"`
		} `yaml:"selector" bson:"selector"`
		Ports []struct {
			Name       string `yaml:"name" bson:"name"`
			Port       int    `yaml:"port" bson:"port"`
			TargetPort int    `yaml:"targetPort" bson:"targetPort"`
		} `yaml:"ports" bson:"ports"`
	} `yaml:"spec" bson:"spec"`
}

type RuleChtoToTam struct {
	RuleGeneralFields   `yaml:",inline" bson:",inline"`
	RuleChtoToTamFields `yaml:",inline" bson:",inline"`
}

func (r *RuleChtoToTam) NewResource() *Resource {
	return &Resource{Resource: r}
}

func NewRuleChtototam() *Resource {
	return (&RuleChtoToTam{}).NewResource()
}
