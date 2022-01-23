package infinitechannel

type InfiniteChannel struct {
	InChan  chan interface{}
	OutChan chan interface{}
	data    []interface{}
	isOpen  bool
}

func (ch *InfiniteChannel) In(val interface{}) {
	ch.InChan <- val
}
func (ch *InfiniteChannel) Out() interface{} {
	return <-ch.OutChan
}

func (ch *InfiniteChannel) Close() {
	close(ch.InChan)
}

func (ch *InfiniteChannel) background() {
	for {
		select {
		case newVal, isOpen := <-ch.InChan:
			// fmt.Println("rece sig ch.InChan(isOpen:)", isOpen)
			if isOpen {
				ch.data = append(ch.data, newVal)
			} else {
				ch.isOpen = false
			}
		case ch.outChanWrapper() <- ch.pop():
			// fmt.Println("rece sig ch.pop")
			// default:
			// 	return
		}
	}
}

func NewInfiniteChannel() *InfiniteChannel {
	cinstance := &InfiniteChannel{
		InChan:  make(chan interface{}),
		OutChan: make(chan interface{}),
		isOpen:  true,
	}
	go cinstance.background()
	return cinstance
}

//
func (ch *InfiniteChannel) pop() interface{} {
	if len(ch.data) == 0 {
		return nil
	}
	val := ch.data[0]
	ch.data = ch.data[1:]
	return val
}

func (ch *InfiniteChannel) outChanWrapper() chan interface{} {
	if ch.isOpen && len(ch.data) == 0 {
		return nil
	}
	return ch.OutChan
}
