package sequences

import (
	"errors"
)

type RunnableManager struct {
	currentRunnable Runnable
	OnStart         func()
	OnFinish        func()
}

var ErrRunnableExists = errors.New("Runnable already running")

func (rm *RunnableManager) Start(r Runnable) error {
	if rm.currentRunnable != nil {
		return ErrRunnableExists
	}

	rm.currentRunnable = r
	rm.onRunnableStart()
	return nil
}

func (rm *RunnableManager) onRunnableStart() {
	if rm.OnStart != nil {
		rm.OnStart()
	}
	rm.currentRunnable.Run(rm.onRunnableFinish)
}

func (rm *RunnableManager) onRunnableFinish() {
	if rm.OnFinish != nil {
		rm.OnFinish()
	}
	rm.currentRunnable = nil
}
