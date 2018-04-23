package pool

type Pool struct {
	QueueSize int
	Workers   int
	Feed      chan *Work
}

func (p *Pool) open() {
	n := p.Workers

	for n != 0 {
		go p.worker()
		n--
	}
}

func (p *Pool) worker() {
	for work := range p.Feed {
		work.Executor(work.Context)
	}
}

func (p *Pool) Close() {
	close(p.Feed)
}

func New(workers int, queueSize int) *Pool {
	pool := &Pool{
		QueueSize: queueSize,
		Workers:   workers,
		Feed:      make(chan *Work, queueSize),
	}

	pool.open()

	return pool
}

func NewFromConfig(c Config) *Pool {
	return New(
		c.Workers,
		c.QueueSize,
	)
}
