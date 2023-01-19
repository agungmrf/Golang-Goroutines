package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*Membuat Channel*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Agung Ma'ruf" //Untuk mengirim data
		fmt.Println("Selesai mengirim Data ke Channel")
	}()

	data := <-channel //untuk menerima data
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

/*Channel Sebagai Parameter*/

func GiveMeResponse(channel chan string) { //tidak butuh pointer
	time.Sleep(2 * time.Second)
	channel <- "Agung Ma'ruf"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel //untuk menerima data
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

/*Channel In dan Out*/

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Agung Ma'ruf"
}

func OnlyOut(channel <-chan string) {
	data := <-channel //untuk menerima data
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

/*Buffered Channel*/

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Agung"
	// channel <- "Ma'ruf"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	go func() {
		channel <- "Agung"
		channel <- "Ma'ruf"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

/*Note: Sedikit berbeda dengan channel biasa, jadi ketika channel biasa saat memasukkan data-nya kedalam channel
karena dia tidak memiliki buffered maka dia diminta untuk menunggu sampai ada yang mengambil. Jika memakai buffered maka otomatis masuk kedalam buffered,
tidak perlu menunggu lagi kecuali bufferednya sudah habis maka diminta menunggu*/

/*Range Channel*/

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

/*Select Channel*/

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

/*Default Channel*/

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}

/*---------------------------------------------------------------------------------------------------------------*/

/*
Cara Menjalankan Test:
go test -v -run=(function name)
contoh:

golang-goroutines> go test -v -run=TestRangeChannel
*/
