package taskpool

import (
	"bytes"
	"github.com/panjf2000/ants/v2"
	"runtime"
	"strconv"
	"sync"
)

type TaskPool struct {
	Pool *ants.PoolWithFunc
	//存放ID
	workerSize int
	queueSize  int
	queueChan  chan int64
	handle     func(interface{})
}

var instance *TaskPool
var once sync.Once

func GetInstance() *TaskPool {
	once.Do(func() {
		instance = NewPool()
	})
	return instance
}

func NewPool() *TaskPool {
	task_pool := new(TaskPool)
	task_pool.queueSize = 10000
	task_pool.workerSize = 100
	return task_pool
}

func (p *TaskPool) SetQueueSize(size int) *TaskPool {
	p.queueSize = size
	return p
}

func (p *TaskPool) SetWorkerSize(size int) *TaskPool {
	p.workerSize = size
	return p
}

func (p *TaskPool) SetHandle(handle func(interface{})) *TaskPool {
	p.handle = handle
	return p
}

func (p *TaskPool) Start() {
	p.queueChan = make(chan int64, p.queueSize)
	p.Pool, _ = ants.NewPoolWithFunc(p.workerSize, p.handle, ants.WithNonblocking(false))
	go func() {
		for {
			select {
			case id := <-p.queueChan:
				p.Pool.Invoke(id)
			}
		}
	}()
}

func (p *TaskPool) Release() {
	p.Pool.Release()
}

func (p *TaskPool) GetPool() *ants.PoolWithFunc {
	return p.Pool
}

func (p *TaskPool) Do(id int64) {
	p.queueChan <- id
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
