package typing

import (
	"fmt"
	"reflect"
)



/*
   @purpose:      Returns whether interface value is truly nil ( dynamic value/type 
	 								== nil. ).

                                       return

   @return:        [out]                     bool 
	 																					 Boolean representing Whether interface 
																						 is truly nil.

   @notes:        This function only works with interface values who's dynamic type 
	 								is well defined ( concrete ).
*/

func isNilInterface(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

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
   @purpose:      Returns dynamic value of from to concrete/interface type T without causing a panic. 
	 								Had a panic ensued from the result of the type assertion, the function returns the zero 
									type of T. 

   @param:        [in]                       From             T
	 																					 Concrete type we wish to cast

                                       return

   @return:        [out]                     r            T 
	 																					 Interface value with dynamic type T and dynamic value of
																						 a. 

   @notes:        This function has only been tested with type conversions to 
	 								concrete base types. It should work with any compounds types
									as *new(T) returns the zero type of any compound type ( although
									it escapes ). Has no side effects on the argument passed in. 
*/

func dynamic_cast[T any] ( from any ) T {
	if from, ok :=  any(from).(T) ; ok {
		return from ; 
	}
	return Zeroed[T]() ; 
}

/*
   @purpose:      Returns dynamic value of from to concrete/interface type T without causing panic and
									without mutating from. Had a panic ensued from the result of the type assertion, the 
									function returns the zero type of T. 

   @param:        [in]                       from             T
	 																					 Concrete type we wish to cast

                                       return

   @return:        [out]                     r            T 
	 																					 Interface value with dynamic type T and dynamic value of
																						 a. 

   @notes:        This function has only been tested with type conversions to 
	 								concrete base types. It should work with any compounds types
									as *new(T) returns the zero type of any compound type ( although
									it escapes ). Has no side effects on the argument passed in. 
*/

func NMutCast[T any]( from any ) T {
	if !isNilInterface(from) {
		if from, ok := any(from).(T) ; ok {
			return from ; 
		}
		return Zeroed[T]() ; 
	}
	return Zeroed[T]() ; 
}

/*
   @purpose:      Mutates dynamic value of from to concrete/interface type T without causing panic. 

   @param:        [in]                       from             T
	 																					 Concrete type we wish to cast

                                       return

   @return:        [out]                     r            T 
	 																					 Interface value with dynamic type T and dynamic value of
																						 a. 

   @notes:        This function has only been tested with type conversions to 
	 								concrete base types. It should work with any compounds types
									as *new(T) returns the zero type of any compound type ( although
									it escapes ). Has no side effects on the argument passed in. 
*/

func MutCast[T any] ( from *any ) {
	if nv , ok :=  any(from).(T) ; ok {
		*from = nv ; 
	}
	*from = Zeroed[T]() ; 
}

func decltype[T any] ( a T ) reflect.Type {
	return reflect.TypeOf(a) ; 
}

