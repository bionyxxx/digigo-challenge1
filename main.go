package main

import "fmt"

func main() {
	//DUTTA FACHREZY

	//Panduan pengerjaan
	//https://cheatography.com/gpascual/cheat-sheets/golang-fmt-printing/

	//menampilkan nilai i : 21
	i := 21
	fmt.Printf("%v \n", i)
	//menampilkan tipe data variabel 1
	fmt.Printf("%T \n", i)

	//menampilkan tanda %
	fmt.Printf("%% \n")

	//menampilkan nilai boolean j : true
	j := true
	fmt.Printf("%v \n", j)
	//menampilkan nilai boolean j : false
	fmt.Printf("%t \n", !j)

	//menampilkan unicode russia Я (ya)
	fmt.Printf("%c \n", 1071)

	//menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", 21)

	//menampilkan nilai base 8 :25
	fmt.Printf("%o \n", 21)

	//menampilkan nilai base 16 : f
	fmt.Printf("%x \n", 15)

	//menampilkan nilai base 16 : F 13
	fmt.Printf("%X \n", 15)

	//menampilkan unicode karakter Я : U+042F var k ﬂoat64 = 123.456;
	fmt.Printf("%U \n", 1071)

	//menampilkan ﬂoat : 123.456000
	var k float64 = 123.456
	fmt.Printf("%f \n", k)

	//menampilkan ﬂoat scientiﬁc : 1.234560E+02
	fmt.Printf("%e \n", k)
}
