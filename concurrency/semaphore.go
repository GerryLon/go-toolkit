package concurrency

type Semaphore struct {
	bufSize int
	channel chan struct{}
}

func New(concurrency int) *Semaphore {
	return &Semaphore{channel: make(chan struct{}, concurrency), bufSize: concurrency}
}

func (p *Semaphore) TryAcquire() bool {
	select {
	case p.channel <- struct{}{}:
		return true
	default:
		return false
	}
}

func (p *Semaphore) Acquire() {
	p.channel <- struct{}{}
}

func (p *Semaphore) Release() {
	<-p.channel
}

func (p *Semaphore) Available() int {
	return p.bufSize - len(p.channel)
}
