package react

//import "fmt"

type cell struct{
	value int
	callbacks []*func(int)
	sheet *reactor
	srcCells []Cell
	updatefunc1 func(int) int
	updatefunc2 func(int,int) int //TODO clean up
}

func (c *cell) Value() int{
	return c.value
}

func (c *cell) SetValue(v int){
	doSomeCalls := c.value != v //changes ?

	c.value = v
	

	if c.srcCells == nil || len(c.srcCells) == 0 { //no srcCell => input cell
		c.sheet.update()	
	}

	if doSomeCalls {
		for _, f := range c.callbacks {
			(*f)(v)
		}
	}

}

func (c *cell) update(){

	if c.updatefunc1 != nil { 
		c.SetValue( c.updatefunc1( c.srcCells[0].Value()) ) 
	} else if c.updatefunc2 != nil { 
		c.SetValue(c.updatefunc2( c.srcCells[0].Value(), c.srcCells[1].Value() )) 
	}
	
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
		srcCells: []Cell{c},
		value: f(c.Value()),
		updatefunc1: f,
		sheet: s,
	}

	*(s.cells) = append(*(s.cells), cellAddr)

	return cellAddr
}

func (s *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell{
	cellAddr := &cell{
		srcCells: []Cell{c1, c2},
		value: f(c1.Value(), c2.Value()),
		updatefunc2: f,
		sheet: s,
	}

	*(s.cells) = append(*(s.cells), cellAddr)

	return cellAddr
}