package typing

import (
	"fmt"
	"reflect"
)

type Numeric interface{
	~int  | ~int8  | ~int16  | ~int32  | ~int64
	~uint | ~uint8 | ~uint16 | ~uint32 
}

type Textual interface{
	~string 
}

type Base interface {
	~int  | ~int8  | ~int16  | ~int32  | ~int64
	~uint | ~uint8 | ~uint16 | ~uint32 | ~string 
	~complex64 | ~complex128 | ~bool
}


func Zeroed[ T any ] ( T ) T {
	return *new(T) ; 
}

func dynamic_cast[T any, U any] ( a U ) U {
	if r, ok :=  any(a).(U) ; ok {
		return r ; 
	}
	panic("Runtime Conversion Failed") ; 
}


func decltype[T any] ( a T ) reflect.Type {
	return reflect.TypeOf(a) ; 
}



func main(){
	var y int = 1 ; 
	x := dynamic_cast[float64,int](y) ; 
	fmt.Println(x) ; 
}


