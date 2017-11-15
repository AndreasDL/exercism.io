package erratum


func Use(o ResourceOpener, input string) (err error){

	res, err := o()
	for {

		if err == nil {
			break
		} else if _, ok := err.(TransientError) ; !ok {
			return err
		}

		res, err = o()
	}
	if err != nil {
		return err
	}
	defer res.Close()


	defer func(){
		if r := recover(); r != nil {
			if obj, ok := r.(FrobError) ; ok {
				res.Defrob(obj.defrobTag)
			} 

			if obj, ok := r.(error); ok {
				err = obj //overwrite error value ! => we need named return
			}
		}
	}()
	res.Frob(input)


	return
}