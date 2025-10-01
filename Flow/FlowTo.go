package Flow

import (
	"io"
	"sync"
)

type FlowTo[O any] struct {
	dependency io.Closer
	output     chan O
	*sync.WaitGroup
}

func MakeFlowTo[I, O any](input chan I, logic func(data I, output chan O), dependency io.Closer) *FlowTo[O] {
	output := make(chan O)
	waitGroup := new(sync.WaitGroup)
	waitGroup.Go(func() {
		for {
			select {
			case data, ok := <-input:
				if ok {
					logic(data, output)
				} else {
					close(output)
					return
				}
			}
		}
	})
	return &FlowTo[O]{
		output:     output,
		WaitGroup:  waitGroup,
		dependency: dependency,
	}
}

func (flowTo *FlowTo[O]) Close() error {
	err := flowTo.dependency.Close()
	if err != nil {
		return err
	}
	flowTo.WaitGroup.Wait()
	return nil
}

func (flowTo *FlowTo[O]) GetChannel() chan O {
	return flowTo.output
}
