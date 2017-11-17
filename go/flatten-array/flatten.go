package flatten

//import "fmt"
import "reflect"


func Flatten(list interface{}) []interface{} {
	result := new([]interface{})
	
	flatten(list, &result)

	if len(*result) == 0 { //null elements => return empty list instead of nil
		return make([]interface{}, 0)
	}

	return *result
}

func flatten(list interface{}, result **[]interface{}){
	if list == nil {
		return
	}

	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:

		l, _ := list.([]interface{})

		for _, i := range l {
			flatten(i, result)
		}

	default: //not a slice => append
		r := append( *(*result) , list)
		*result = &r //point pointer to the location of the new list, hence the double pointer
	}
}