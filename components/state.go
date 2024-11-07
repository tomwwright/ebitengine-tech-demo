package components

import (
	"github.com/yohamta/donburi"
)

type StateData struct {
	state map[string]int
}

var State = donburi.NewComponentType[StateData]()

func NewState() *StateData {
	return &StateData{
		state: map[string]int{},
	}
}

func (s *StateData) Get(key string) int {
	return s.state[key]
}

func (s *StateData) Set(key string, value int) int {
	s.state[key] = value
	return s.state[key]
}

func (s *StateData) SetTrue(key string) {
	s.Set(key, 1)
}

func (s *StateData) IsTrue(key string) bool {
	return s.state[key] == 1
}

func (s *StateData) Increment(key string) int {
	return s.Set(key, s.Get(key)+1)
}
