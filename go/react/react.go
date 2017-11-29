package react

//import "fmt"

type cell struct{
	value int
	callbacks []*func(int)
	sheet *reactor
	srcCell Cell
	updatefunc func(int) int
}

func (c *cell) Value() int{
	return c.value
}

func (c *cell) SetValue(v int){
	c.value = v
	
	if c.srcCell == nil { //no srcCell => input cell
		c.sheet.update()	
	}
}

func (c *cell) update(){
	c.SetValue( c.updatefunc( c.srcCell.Value()) )
}

type canceler struct{
	f *func(int)
	cell *cell
}

func (c *cell) AddCallback(f func(int)) Canceler{
	c.callbacks = append(c.callbacks, &f)
	return canceler{
		cell: c,
		f: &f,
	}
}

func (c canceler) Cancel(){

	cb := make([]*func(int), len(c.cell.callbacks)-1)
	a := 0
	for i, f := range c.cell.callbacks {

		
		if f == c.f { 
			a = 1
		} else {
			cb[i - a] = f	
		}
	}

	c.cell.callbacks = cb
}

type reactor struct{
	cells *[]*cell
}

func New() Reactor {
	list := make([]*cell, 0)

	return &reactor{
		cells: &list,
	}
}

func (s *reactor) update() {
	for _, c := range *(s.cells) {
		c.update()
	}
}

func (s *reactor) CreateInput(v int) InputCell {
	return &cell{ 
		value: v,
		sheet: s,
	}
}

func (s *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell{
	cellAddr := &cell{
		srcCell: c,
		value: f(c.Value()),
		updatefunc: f,
		sheet: s,
	}

	*(s.cells) = append(*(s.cells), cellAddr)

	return cellAddr
}

func (s reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell{
	return &cell{}
}