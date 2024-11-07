package yaml

import "gopkg.in/yaml.v3"

type Interactions struct {
	Interactions map[string][]Interaction `yaml:"interactions"`
}

type Interaction struct {
	Debug    *Debug    `yaml:"debug"`
	Dialogue *Dialogue `yaml:"dialogue"`
}

type Debug struct {
	Text string `yaml:"text"`
}

type Dialogue struct {
	Text string `yaml:"text"`
}

func UnmarshallInteractions(content []byte) (interactions *Interactions, err error) {
	err = yaml.Unmarshal(content, &interactions)
	return interactions, err
}
