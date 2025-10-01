package Flow

import (
	"io"
	"sync"
)

type FlowEnd struct {
	dependency io.Closer
	*sync.WaitGroup
}

func MakeFlowEnd[I any](input chan I, logic func(data I), dependency io.Closer) *FlowEnd {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Go(func() {
		for {
			select {
			case data, ok := <-input:
				if ok {
					logic(data)
				} else {
					return
				}
			}
		}
	})

	return &FlowEnd{
		WaitGroup:  waitGroup,
		dependency: dependency,
	}
}

func (flowEnd *FlowEnd) Close() error {
	err := flowEnd.dependency.Close()
	if err != nil {
		return err
	}
	flowEnd.WaitGroup.Wait()
	return nil
}

func (flowEnd *FlowEnd) GetChannel() chan any {
	return nil
}
