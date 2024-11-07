package sequences_test

import (
	"techdemo/sequences"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRunnable struct {
	mock.Mock
}

func (m *MockRunnable) Run(done sequences.Done) {
	m.Called(done)
	done()
}

func MockSequence() ([]*MockRunnable, sequences.Sequence) {

	mocks := []*MockRunnable{
		{},
		{},
		{},
	}

	for _, m := range mocks {
		m.On("Run", mock.Anything).Run(func(args mock.Arguments) {
			done := args.Get(1).(func())
			done()
		})
	}

	return mocks, sequences.Sequence{
		Steps: lo.Map(mocks, func(m *MockRunnable, i int) sequences.Runnable { return sequences.Runnable(m) }),
	}
}

func TestInteractionSequence(t *testing.T) {
	t.Run("calls interactions in sequence", func(t *testing.T) {
		mocks, sequence := MockSequence()
		sequence.Run(func() {})

		for _, m := range mocks {
			m.AssertCalled(t, "Run", mock.Anything)
		}
	})

	t.Run("calls callback when done", func(t *testing.T) {
		_, sequence := MockSequence()
		wasCalled := false
		sequence.Run(func() {
			wasCalled = true
		})

		assert.True(t, wasCalled)
	})
}
