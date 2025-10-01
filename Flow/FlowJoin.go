package Flow

import (
	"io"
	"sync"
)

type FlowJoin[O any] struct {
	dependencies []io.Closer
	output       chan O
	*sync.WaitGroup
}

func MakeFlowJoin[O any](inputs []chan O, dependencies []io.Closer) *FlowJoin[O] {
	output := make(chan O)
	waitGroup := new(sync.WaitGroup)
	for _, input := range inputs {
		waitGroup.Go(func() {
			for {
				select {
				case data, ok := <-input:
					if ok {
						output <- data
					} else {
						return
					}
				}
			}
		})
	}

	go func() {
		waitGroup.Wait()
		close(output)
	}()

	return &FlowJoin[O]{
		output:       output,
		WaitGroup:    waitGroup,
		dependencies: dependencies,
	}
}

func (flowJoin *FlowJoin[O]) Close() error {
	for _, dependency := range flowJoin.dependencies {
		err := dependency.Close()
		if err != nil {
			return err
		}
	}
	flowJoin.WaitGroup.Wait()
	return nil
}

func (flowJoin *FlowJoin[O]) GetChannel() chan O {
	return flowJoin.output
}
