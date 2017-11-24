package circular

import (
	"errors"
)

type Buffer struct{
	size, usedSlots int
	rpos, wpos int
	data []byte
}


func NewBuffer(size int) *Buffer {
	if size <= 0 { return nil }

	return &Buffer{
		size      : size,
		usedSlots : 0, 
		rpos      : 0,
		wpos      : 0,
		data      : make([]byte, size),
	}

}

func (b *Buffer) ReadByte() (byte, error){

	if b.usedSlots == 0 {
		return byte(0), errors.New("Buffer is empty")
	}

	res := b.data[b.rpos]
	b.rpos ++
	b.rpos %= b.size
	b.usedSlots--
	
	return res, nil
}


func (b *Buffer) WriteByte(c byte) error{

	if b.size == b.usedSlots { return errors.New("Buffer is Full")}
	b.usedSlots++

	b.data[b.wpos] = c
	b.wpos++
	b.wpos %= b.size

	return nil
}

func (b *Buffer) Overwrite(c byte){

	if b.usedSlots == b.size { //buffer full
		b.data[b.wpos] = c
		b.wpos++
		b.wpos %= b.size

		b.rpos++
		b.rpos %= b.size
	
	} else { // do normal write
		b.WriteByte(c)	
	}
	
}

func (b *Buffer) Reset(){
	b.usedSlots = 0
	b.wpos = 0
	b.rpos = 0
}