package typing

import (
	"fmt"
	"reflect"
)


/*
   @purpose:      Returns the zero type of any concrete type. 

                                       return

   @return:        [out]                     r            T 
	 																					 Instace of a zero valued Concrete type.

   @notes:        This function only works with interface values who's dynamic type 
	 								is well defined ( concrete ).
*/

func Zeroed[ T any ] (  ) T {
	return *new(T) ; 
}


/*
   @purpose:      Casts type a to concrete/interface type T without causing a panic. Had a panic
	 								ensued from the result of the type assertion, the function returns the zero 
									type of T. 

   @param:        [in]                       a             T
	 																					 Concrete type we wish to cast

                                       return

   @return:        [out]                     r            T 
	 																					 Interface value with dynamic type T and dynamic value of
																						 a. 

   @notes:        This function has only been tested with type conversions to 
	 								concrete base types. It should work with any compounds types
									as *new(T) returns the zero type of any compound type ( although
									it escapes ). 
*/

func dynamic_cast[T any] ( a any ) T {
	if r, ok :=  any(a).(T) ; ok {
		return r ; 
	}
	return Zeroed[T]() ; 
}


func decltype[T any] ( a T ) reflect.Type {
	return reflect.TypeOf(a) ; 
}

