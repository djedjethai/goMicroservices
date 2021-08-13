package main

import (
	"fmt"
	// "reflect"
	// "regexp"
)

type perimeter interface {
	per()
}

type rect struct {
	long  float64
	large float64
}

func (r rect) per() {
	fmt.Printf("pre rect: %v\n", r.long*r.large)
}
func (r rect) surf() {
	fmt.Println("surff rect")
}

type tria struct {
	base  float64
	hight float64
}

type shape interface {
	perimeter()
}

type rect struct {
	long  int
	large int
}

func (r rect) perimeter() {
	fmt.Printf("rect: %v\n", r.long*r.large)
}
func myrect(r rect) {
	fmt.Printf("i am a rect: %v\n", r.long)
}

type triangle struct {
	base  int
	hight int
}

func (t triangle) perimeter() {
	fmt.Printf("trian: %v\n", (t.base*t.hight)/2)
}
func mytri(t triangle) {
	fmt.Printf("i am a triangle: %v\n", t.base)
}

// special switch for type assertion(sh will get any type)
func yard(s shape) {
	switch sh := s.(type) {
	case rect:
		s.perimeter()
		myrect(sh)
	case triangle:
		s.perimeter()
		mytri(sh)
	}
}

// heavy way
// func yard(s shape) {
// 	rec, ok := s.(rect)
// 	if ok {
// 		s.perimeter()
// 		myrect(rec)
// 	}
//
// 	tri, ok := s.(triangle)
// 	if ok {
// 		s.perimeter()
// 		mytri(tri)
// 	}
// }

func main() {
	re := rect{23, 45}
	tr := triangle{12, 67}

	yard(re)
	yard(tr)
}

//
// type speacker interface {
// 	say()
// }
//
// type dog struct {
// 	name string
// }
//
// func (d *dog) say() {
// 	if d == nil {
// 		fmt.Println("erererere")
// 	} else {
// 		fmt.Println(d.name)
// 	}
// }
//
// func main() {
// 	var s speacker
// 	var d *dog
// 	s = d
// 	s.say()
// 	// fmt.Println(reflect.TypeOf(s))
// }
//
// type I interface {
// 	M()
// }
//
// type I2 interface {
// 	I
// 	N()
// }
//
// type T struct {
// 	name string
// }
//
// func (T) M() {}
// func (T) N() {}
//
// func main() {
// 	var v1 I = T{"foo"}
// 	var v2 I2
// 	v2, ok := v1.(I2)
// 	fmt.Printf("%T %v\n", v1, v1)
// 	fmt.Printf("%T %v %v\n", v2, v2, ok)
// }

// type I interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// type T1 struct{}
//
// func (T1) M() {}
//
// func main() {
// 	var v I = T{}
// 	v2, ok := v.(T1)
// 	if !ok {
// 		fmt.Printf("ok: %v\n", ok)
// 		fmt.Printf("%v, %T\n", v2, v2)
// 	}
// } // ok

// type I interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// type T1 struct{}
//
// func (T1) M() {}
//
// func main() {
// 	var v I = T{}
// 	v2 := v.(T1)
// 	fmt.Printf("%T\n", v2)
// } // panic: interface conversion: main.I is main.T, not main.T1

// type I interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// func main() {
// 	var v I = T{}
// 	v2 := v.(T)
// 	fmt.Printf("%T\n", v2)
// } // ok

// type I interface {
// 	M()
// }
//
// type I2 interface {
// 	M()
// 	N()
// }
//
// func main() {
// 	var v I
// 	fmt.Println(I2(v))
// } // I does not implement I2 (missing N method)

// type I interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// func main() {
// 	var v I = T{}
// 	fmt.Println(T(v))
// } // cannot convert v (type I) to type T: need type assertion

// type I interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// func main() {
// 	var v I = T{}
// 	var v1 T = v
// 	_ = v1
// } // ERR cannot use v (type I) as type T in assignment: need type assertion

// type I interface{
// 	M()
// }
//
// type T struct{}
//
// func(T)M(){}
//
// func main() {
// 	var v I = T{}
// 	_ = v
// } // okkk

// // Value of type I2 implements I1 since its method set is a subset of methods from I1. If this is not the case then compiler will react accordingly
// type I1 interface {
// 	M1()
// }
//
// type I2 interface {
// 	M1()
// 	M2()
// }
//
// type T struct{}
//
// func (T) M1() {}
// func (T) M2() {}
//
// func main() {
// 	var v1 I1 = T{}
// 	var v2 I2 = v1
// 	_ = v2		// ERR: cannot use v1 (type I1) as type I2 in assignment:
// 			// I1 does not implement I2 (missing M2 method)
// }

// // Even if I2 has other interface embedded but I1 does not then these interfaces still implement each other.
// // Order of methods also doesn’t matter. It’s worth to remember that methods sets don’t have to be equal
// type I1 interface {
// 	M1()
// 	M2()
// }
//
// type I2 interface {
// 	M1()
// }
//
// type T struct{}
//
// func (T) M1() {}
// func (T) M2() {}
//
// func main() {
// 	var v1 I1 = T{}
// 	var v2 I2 = v1
// 	_ = v2
// }

// // T is an interface type and x implements T
// // because v1’s type implements I2 interface. It doesn’t matter how these types are structured
// type I1 interface {
// 	M1()
// 	M2()
// }
//
// type I2 interface {
// 	M1()
// 	I3
// }
//
// type I3 interface {
// 	M2()
// }
//
// type T struct{}
//
// func (T) M1() {}
// func (T) M2() {}
//
// func main() {
// 	var v1 I1 = T{}
// 	var v2 I2 = v1
// 	_ = v2
// }

// // we’ve two interface type variables and we want to assign one to another
// type I1 interface {
// 	M()
// }
//
// type I2 interface {
// 	M()
// }
//
// type T struct{}
//
// func (T) M() {}
//
// func main() {
// 	var v1 I1 = T{}
// 	var v2 I2 = v1
// 	fmt.Printf("%v - %T\n", v1, v1)
// 	fmt.Printf("%v - %T", v2, v2)
// }

// type T1 struct {
// 	name string
// }
//
// type T2 struct {
// 	name string
// }
//
// func main() {
// 	vs := []interface{}{T2(T1{"foo"}), string("322"), []byte("qwerty")}
// 	for _, d := range vs {
// 		fmt.Printf("%v %T\n", d, d)
// 	}
// }
