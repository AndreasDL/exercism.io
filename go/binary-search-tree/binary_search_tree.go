package binarysearchtree

type SearchTreeData struct{
	left, right *SearchTreeData
	data int
}

func Bst(v int) *SearchTreeData{
	return &SearchTreeData{nil,nil,v}
}

func (bst *SearchTreeData) Insert(v int){

	//walk trough tree

	currTree := bst

	done := false
	for !done {

		if v <= currTree.data {
			
			if currTree.left == nil { 
				currTree.left = Bst(v) 
				done = true
			}
			currTree = currTree.left

		} else {

			if currTree.right == nil { 
				currTree.right = Bst(v)
				done = true
			}
			currTree = currTree.right

		}

	}
}



func (bst *SearchTreeData) MapString(f func(int)string) []string{
	if bst == nil { return []string{} }

	res := []string{}
	res  = append(res, bst.left.MapString(f)...)
	res  = append(res, f(bst.data))
	res  = append(res, bst.right.MapString(f)...)

	return res
}

func (bst *SearchTreeData) MapInt(f func(int)int) []int{
	if bst == nil { return []int{} }

	res := []int{}
	res  = append(res, bst.left.MapInt(f)...)
	res  = append(res, f(bst.data))
	res  = append(res, bst.right.MapInt(f)...)

	return res
}



type stack []*SearchTreeData

func (s *stack) pop() *SearchTreeData{
	l := len(*s)
	if l == 0 {
		return nil
	} else { 
		c := (*s)[l-1]
		*s = (*s)[:l-1] 
		return c
	}
}

func (s *stack) push(c *SearchTreeData) {
	*s = append(*s, c)
}
