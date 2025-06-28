package typing

import (
	"fmt"
	"reflect"
)


var zero_values map[reflect.Kind]interface{} = map[reflect.Kind]interface{}{
	reflect.Uint   : 0 , 
	reflect.Uint8  : 0 , 
	reflect.Uint16 : 0 , 
	reflect.Uint32 : 0 , 
	reflect.Int    : 0 , 
	reflect.Int8   : 0 , 
	reflect.Int16  : 0 , 
	reflect.Int32  : 0 , 
	reflect.Int64  : 0 , 
	reflect.String : "",
	reflect.Bool   : false ,
}


type Numeric interface{
	~int  | ~int8  | ~int16  | ~int32  | ~int64
	~uint | ~uint8 | ~uint16 | ~uint32 
}

type Textual interface{
	~string 
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


