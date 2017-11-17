package paasio

import (
	"io"
	"sync"
)


type counter struct {
	opCount int
	rCount, wCount int64
	rf func(p []byte) (int, error)
	wf func(p []byte) (int, error)
	sync.Mutex
}

func (c counter) ReadCount() (int64, int){
	return c.rCount, c.opCount
}

func (c counter) WriteCount() (int64, int){
	return c.wCount, c.opCount
}

func (c *counter) Read(p []byte) (int, error){

	n, err := c.rf(p)
	
	c.Lock()
	defer c.Unlock()

	if err == nil {
		c.opCount++
		c.rCount += int64(n)	
	}
	
	return n, err
}

func (c *counter) Write(p []byte) (int, error){
	
	n, err := c.wf(p)
	
	c.Lock()
	defer c.Unlock()

	if err == nil {
		c.opCount++
		c.wCount += int64(n)
	}
	
	return n, err
}

func NewReadCounter(r io.Reader) ReadCounter{
	return &counter{
		opCount : 0,
		rCount  : 0,
		wCount  : 0,
		rf      : r.Read,
		wf      : nil,
	}
}

func NewWriteCounter(w io.Writer) WriteCounter{
	return &counter{
		opCount : 0,
		rCount  : 0,
		wCount  : 0,
		rf      : nil,
		wf      : w.Write,
	}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter{
	return &counter{
		opCount : 0,
		rCount  : 0,
		wCount  : 0,
		rf		: rw.Read,
		wf      : rw.Write,
	}
}