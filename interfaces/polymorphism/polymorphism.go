package polymorphism
/*
	With different dynamic values of different dynamic types boxed into the interface value, 
	the interface value behaves differently. This is called polymorphism.
	
	For non-blank interface types, the definition like the following one is used.

	type _interface struct {
		dynamicTypeInfo *struct {
			dynamicType *_type       // the dynamic type
			methods     []*_function // method table
		}
		dynamicValue unsafe.Pointer // the dynamic value
	}

	When method i.m is called, the method table in the implementation information stored in i 
	will be looked up to find and call the corresponding method t.m. 
	The method table is a slice and the lookup is just a slice element indexing, so this is quick.

	- https://go101.org/article/interface.html#polymorphism
*/

