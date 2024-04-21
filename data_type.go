package main

// in go there is 4 data type basically like other programming languages
//int
// float
// boolean
// string

//?boolean store true and false value 
// it use to determine the program decision

//?string  ""
//data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:""


//?int 
//used to store a whole number without decimals  it can be positive or negative
//? Signed int ==> can store both positive and negative
// Unsigned int  ==> only store positive


//? Signed int
//int depend on system if 32bit system  -2147483648 to 2147483647 
//if 64system  -9223372036854775808 to 9223372036854775807

//int8 ==> 8 bits/1 byte	-128 to 127

//int16 ==> 16 bits/2 byte	-32768 to 32767

//int32 ==>	32 bits/4 byte	-2147483648 to 2147483647

//int64 ==> 64 bits/8 byte	-9223372036854775808 to 9223372036854775807


//?Unsigned Int
//uint depend on system if 32bit system  0 to 4294967295
//if 64system  0 to 18446744073709551615

//uint8 ==> 8 bits/1 byte 0 to 255

//uint16 ==> 16 bits/2 byte	0 to 65535

//uint32 ==>	32 bits/4 byte	0 to 4294967295

//uint64 ==> 64 bits/8 byte	0 to 18446744073709551615

//?float to store positive and negative numbers with a decimal point
//float32
//float64 (default) can store more than float 32

//? Float vs Double 
//float: Typically a single-precision floating-point number, which means it provides about 6-9 significant decimal digits of precision.
//double:a double-precision floating-point number, offering about 15-17 significant decimal digits of precision.

//normal float or float32 require 4bytes 32 bits memory 
// double require 8bytes 64bits memory 
// double == float64