package sequences

type Sequence struct {
	Steps []Runnable
	done  func()
}

func (s *Sequence) Run(done Done) {
	s.done = done
	s.nextStep()
}

func (s *Sequence) nextStep() {
	if len(s.Steps) == 0 {
		s.done()
		return
	}

	step := s.Steps[0]
	s.Steps = s.Steps[1:]

	step.Run(s.onStepDone)
}

func (s *Sequence) onStepDone() {
	s.nextStep()
}
