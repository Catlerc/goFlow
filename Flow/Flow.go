package Flow

type Flow[O any] interface {
	GetChannel() chan O
	Close() error
}

type CloseChannel[T any] struct {
	Channel chan T
}

func (cc *CloseChannel[T]) Close() error {
	close(cc.Channel)
	return nil
}
