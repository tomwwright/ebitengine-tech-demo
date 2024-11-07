package yaml

import "gopkg.in/yaml.v3"

type Interactions struct {
	Interactions map[string][]Interaction `yaml:"interactions"`
}

type Interaction struct {
	Debug    *Debug    `yaml:"debug"`
	Despawn  *Despawn  `yaml:"despawn"`
	Dialogue *Dialogue `yaml:"dialogue"`
	State    *State    `yaml:"state"`
}

type Debug struct {
	Text string `yaml:"text"`
}

type Despawn struct {
	Name string `yaml:"name"`
}

type Dialogue struct {
	Text string `yaml:"text"`
}

type State struct {
	Key    string `yaml:"key"`
	Value  int    `yaml:"value"`
	Action string `yaml:"action"`
}

func UnmarshallInteractions(content []byte) (interactions *Interactions, err error) {
	err = yaml.Unmarshal(content, &interactions)
	return interactions, err
}
