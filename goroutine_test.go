package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

/*---------------------------------------------------------------------------------------------------------------*/

/*

golang-goroutines> go test -v -run=TestManyGoroutine

Note : Jika hasil TestManyGoroutine tidak berurutan Biarkan saja, Karena jika running menggunakan Goroutine dia running-nya asynchronous,
walaupun dia concurrent jika menggunakan perangkat yang multicore dia tidak hanya merunning secara concurrent dia juga merunning secara pararel.
Maka dari itu angkanya bisa berubah-ubah, Yang penting adalah semuanya kelar artinya 10 ribu goroutine semuanya sukses dieksekusi, jadi tidak ada error apapun, tidak ada memory overflow.
*/
