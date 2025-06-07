package yaml

import "gopkg.in/yaml.v3"

type Interactions struct {
	Interactions map[string][]Step `yaml:"interactions"`
	Triggers     map[string][]Step `yaml:"triggers"`
}

type Step struct {
	Debug    *Debug    `yaml:"debug"`
	Despawn  *Despawn  `yaml:"despawn"`
	Dialogue *Dialogue `yaml:"dialogue"`
	State    *State    `yaml:"state"`
	When     *When     `yaml:"when"`
	Teleport *Teleport `yaml:"teleport"`
	Filter   *Filter   `yaml:"filter"`
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

type Teleport struct {
	To string `yaml:"to"`
}

type Filter struct {
	Name string `yaml:"name"`
}

type When struct {
	Conditions []Condition `yaml:"conditions"`
	Steps      []Step      `yaml:"steps"`
	Else       []Step      `yaml:"else"`
}

type Condition struct {
	Or    []Condition     `yaml:"or"`
	State *StateCondition `yaml:"state"`
}

type StateCondition struct {
	Key   string `yaml:"key"`
	Value int    `yaml:"value"`
}

func UnmarshallInteractions(content []byte) (interactions *Interactions, err error) {
	err = yaml.Unmarshal(content, &interactions)
	return interactions, err
}
