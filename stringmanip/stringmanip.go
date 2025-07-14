package stringmanip

import (
	"strings"
)

var  ASCII_UPPERCASE_LETTER_RANGE  = struct{ LowerBound, UpperBound byte }{ 
	LowerBound : 65, // A
	UpperBound : 90, // Z
}


var  ASCII_LOWERCASE_LETTER_RANGE  = struct{ LowerBound, UpperBound byte }{ 
	LowerBound : 97,  // a
	UpperBound : 122, // z
}

/*
   @purpose:      Convert a string in pascal case to Screaming/Upper Snake case

   @param:        [in]                     form    The string to convert 

                                       return

   @code:         string                   String converted from Pascal case to Screaming Snake Case

   @Example:      if "BorderStroke" is passed in as an argument it is converted to BORDER_STROKE.

*/

func PascalToUSC( from string ) string {
	if len(from) == 0 {
		return ""  ; 
	}
	split := SplitStringByUppercase(from) ; 
	accumulator := strings.ToUpper(split[0]) ; 
	for idx, key := range split {
		if idx != 0 {
			accumulator +=  "_" + strings.ToUpper(key) ; 
		}
	}
	return accumulator ; 
}



/*
   @purpose:      Splits a string by uppercase characters. 

   @param:        [in]                     s    The string to split

                                       return

   @code:         []string                String splice containing the original portion of the string split by 
	 																				its uppercase characters. 

   @Example:      if "BorderStroke" is passed in as an argument, ["Border", "Stroke"] is returned. 

   @Assumes:      This routine assumes that s is pascal cased. Passing a non pascal cased value is UB. 

*/

func SplitStringByUppercase(s string) []string {
	var rval []string
	accumulator := string(s[0])
	for i := 1; i < len(s); i++ {
		ch := byte(s[i])
		if ch <= ASCII_UPPERCASE_LETTER_RANGE.UpperBound && ch >= ASCII_UPPERCASE_LETTER_RANGE.LowerBound {
			rval = append(rval, accumulator)
			accumulator = ""
		}
		accumulator += string(s[i])
	}
	rval = append(rval, accumulator)
	return rval
}



