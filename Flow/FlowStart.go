package Flow

import (
	"io"
)

type FlowStart[O any] struct {
	dependency io.Closer
	output     chan O
}

func MakeFlowStart[O any](output chan O, dependency io.Closer) *FlowStart[O] {
	return &FlowStart[O]{
		output:     output,
		dependency: dependency,
	}
}

func (flowStart *FlowStart[O]) Close() error {
	if flowStart.dependency != nil {
		err := flowStart.dependency.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (flowStart *FlowStart[O]) GetChannel() chan O {
	return flowStart.output
}
