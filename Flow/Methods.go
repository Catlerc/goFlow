package Flow

import "io"

func Start[O any](out chan O, dependency io.Closer) *FlowStart[O] {
	return MakeFlowStart(out, dependency)
}
func To[I, O any](in Flow[I], logic func(data I, output chan O)) *FlowTo[O] {
	return MakeFlowTo(in.GetChannel(), logic, in)
}

func Join[O any](ins ...Flow[O]) *FlowJoin[O] {
	inputChannels := make([]chan O, len(ins))
	closers := make([]io.Closer, len(ins))
	for i, input := range ins {
		inputChannels[i] = input.GetChannel()
		closers[i] = input
	}
	return MakeFlowJoin[O](inputChannels, closers)
}

func End[I any](in Flow[I], logic func(data I)) *FlowEnd {
	return MakeFlowEnd(in.GetChannel(), logic, in)
}
