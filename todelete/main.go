package main

import (
	"bytes"
	"fmt"
	// "reflect"
	// "unicode/utf8"
	// "io/ioutil"
	// "bufio"
	// "encoding/hex"
	// "io"
	// "log"
	// "os"
)

func main() {
	var buf bytes.Buffer

	buf.Write([]byte("an old"))
	buf.WriteByte(32)
	buf.WriteString("cactus")
	buf.WriteByte(32)
	buf.WriteByte(32)
	buf.WriteRune(' ')

	fmt.Println(buf)
	fmt.Println(buf.String())

}

// func main() {
// 	data := [][]byte{[]byte("an"), []byte("old"), []byte("wolf")}
// 	joined := bytes.Join(data, []byte(" - "))
//
// 	fmt.Println(string(joined))
//
// 	fmt.Println("=======================")
//
// 	data2 := []byte{102, 97, 108, 99, 111, 110, 32}
// 	fmt.Println(string(data2))
//
// 	rpt := bytes.Repeat(data2, 3)
// 	fmt.Println(string(rpt))
//
// 	fmt.Println("========================")
//
// 	data3 := []byte{32, 32, 102, 97, 108, 99, 111, 110, 32, 32, 32}
// 	fmt.Println(data3)
// 	fmt.Println(string(data3))
//
// 	fmt.Println(bytes.Trim(data3, " "))
// }

// func main() {
//
// 	data1 := []byte{102, 97, 108, 99, 111, 110}
// 	data2 := []byte{111, 110}
//
// 	if bytes.Contains(data1, data2) {
// 		fmt.Println("containes")
// 	} else {
// 		fmt.Println("do not contains")
// 	}
//
// 	if bytes.Equal([]byte("falcon"), []byte("owl")) {
// 		fmt.Println("equal")
// 	} else {
// 		fmt.Println("not equal")
// 	}
//
// 	data3 := []byte{111, 119, 108, 9, 99, 97, 116, 32, 32, 32, 32, 100, 111, 103, 32, 112, 105, 103, 32, 32, 32, 32, 98, 101, 97, 114}
//
// 	fields := bytes.Fields(data3)
// 	fmt.Println(fields)
//
// 	for _, e := range fields {
// 		fmt.Printf("%s\n", string(e))
// 	}
//
// 	fmt.Println(string(data3))
// }

// func main() {
// 	content, err := ioutil.ReadFile("interface.go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println(content)
// 	fmt.Println("================")
// 	fmt.Println(string(content))
// }

// func main() {
//
// 	var w = []byte("je suis cool")
// 	z := []byte{106, 101, 32, 115, 117, 105, 115, 32, 99, 111, 111, 108}
//
// 	fmt.Println(w)
//
// 	// We convert a slice of bytes to a string with the string function
// 	fmt.Println(string(w))
// 	fmt.Println(string(z))
// }
