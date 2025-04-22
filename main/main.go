package main

import (
	"fmt"
	"gostart/startpackage"
	"os"
	"strings"
	"sync"
	"time"
)

// Definizione di un'interfaccia
type Quacker interface {
	Quack()
}

// Tipo che implementa il metodo richiesto dall'interfaccia
type Duck struct{}

func (d Duck) Quack() {
	fmt.Println("Quack!")
}

// Un altro tipo che implementa lo stesso metodo
type Person struct{}

func (p Person) Quack() {
	fmt.Println("I'm pretending to be a duck!")
}

func (p Person) Quacked() {
	fmt.Println("I want to be a duck!")
}
func duck() {
	fmt.Println("---Duck-------------.")
	var q Quacker
	q = Duck{}
	q.Quack() // Output: Quack!
	var p = Person{}
	p.Quacked()
	p.Quack()
	q = p
	q.Quack() // Output: I'm pretending to be a duck!
	p = q.(Person)
	p.Quacked()
	fmt.Println("**+++++")
}
func print_output() {
	fmt.Println("_Print Output__")
	var i int = 21
	var j bool = true
	num := 15
	// h := strconv.FormatInt(num, 16)
	//h := fmt.Sprintf("%X", i)
	var k float64 = 123.456
	r := 'Я'
	fmt.Println(i)
	fmt.Printf("%T\n%%\n%v\n", i, j)
	// fmt.Println(j)
	fmt.Printf("\n%b\n%c\n%d\n%o\n%x\n%X\nU+%04X\n", i, r, i, i, num, num, r)
	// fmt.Println(h)
	//fmt.Printf("U+%04X\n", r)
	fmt.Printf("\n%f\n%.6E\n", k, k)
	fmt.Println("_________")
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}
	for j := 0; j <= 10; j++ {
		if j <= 4 {
			fmt.Printf("Nilai j = %d\n", j)
		} else {
			const russia = "САШАРВО"
			for index, runeValue := range russia {
				fmt.Printf("character %#U start at byte position %d\n", runeValue, index)
			}
			break
		}
	}
	for j := 0; j <= 12; j++ {
		if j > 5 {
			fmt.Printf("Nilai j = %d\n", j)
		} else if j == 10 {
			break
		}
	}
	fmt.Println("**+++++")
}
func stringa() {

	fmt.Println("-------------stringa-->import strings<--")
	word := "selamat malam"

	result := strings.ReplaceAll(word, "m", "M")
	fmt.Println(word, "=>", result)
	counts := make(map[string]int)

	for _, letter := range result {
		counts[string(letter)]++
	}

	for _, letter := range result {
		fmt.Println(string(letter))
		// if letter == " " {
		// fmt.Println()
		// }
	}

	fmt.Println(counts)
}

func comando() {
	fmt.Println("---RIGA comando----")
	// Retrieve command-line arguments, including the program name
	args := os.Args[0:]
	fmt.Println(args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func struttura() {
	type Person struct {
		Name      string
		Alamat    string
		Pekerjaan string
		Alasan    string
	}
	people := []Person{
		{"Airell", "Rumah", "Curriculum Developer", "Alasan"},
		{"Ananda", "Rumah", "Finance", "Alasan"},
		{"Mailo", "Rumah", "Markerting", "Alasan"},
	}
	println("---Struttura---", len(os.Args), "----")
	for len(os.Args) < 2 {
		return
	}

	index := 0
	fmt.Sscanln(os.Args[1], &index)
	println("index", index)
	if index < len(people) {
		person := people[index]
		fmt.Println("Nama: ", person.Name)
		fmt.Println("Alamat: ", person.Alamat)
		fmt.Println("Pekerjaan: ", person.Pekerjaan)
		fmt.Println("Alasan: ", person.Alasan)
	} else if index >= len(people) {
		fmt.Println("Index out of range")
		return
	}
}

func process1(data any, n int) {
	start := time.Now()
	for i := 1; i < n; i++ {
		//fmt.Println(data, i)
		fmt.Print(i, "a1 ")
	}
	elapsed := time.Since(start)
	fmt.Printf("A1-Execution time: %s\n", elapsed)
}
func process2(data any, n int) {
	start := time.Now()
	for i := 1; i < n; i++ {
		//fmt.Println(data, i)
		fmt.Print(i, "a2 ")
	}
	elapsed := time.Since(start)
	fmt.Printf("A2-Execution time: %s\n", elapsed)

}
func tempi() {

	println("---Tempi--->import time<---")
	// Get the current time
	currentTime := time.Now()
	// Format the time as a string
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	// Print the formatted time
	fmt.Println("Current time:", formattedTime)
	data1 := "[bisa1 bisa2 bisa3]"
	data2 := "[coba1 coba2 coba3]"
	data3 := "[aoba1 aoba2 aoba3]"
	index := 0
	if len(os.Args) > 1 {
		fmt.Sscanln(os.Args[1], &index)
	}

	for range 1 {
		go process1(data1, index)
		go process2(data2, index)
		go process1(data3, index)
	}

	time.Sleep(2 * time.Second)

}
func process1Mutex(data any, n int, mutex *sync.Mutex) {
	startwait := time.Now()
	mutex.Lock()
	start := time.Now()
	for i := range n {
		fmt.Print(i, "m1 ")
		//	time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("\n")
	mutex.Unlock()
	elapsed := time.Since(start)
	fmt.Printf("M1-Execution time: %s\n", elapsed)
	fmt.Printf("M1-Elapsed time: %s\n", start.Sub(startwait))
}
func process2Mutex(data any, n int, mutex *sync.Mutex) {
	startwait := time.Now()
	mutex.Lock()
	start := time.Now()
	for i := range n {
		fmt.Print(i, "m2 ")
		//	time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("\n")
	mutex.Unlock()
	elapsed := time.Since(start)
	fmt.Printf("M2-Execution time: %s\n", elapsed)
	fmt.Printf("M2-Elapsed time: %s\n", start.Sub(startwait))
}
func tempiMutex() {
	println("---TempiMutex--->import sync<---")
	data1 := "[bisa1 bisa2 bisa3]"
	data2 := "[coba1 coba2 coba3]"
	data3 := "[aoba1 aoba2 aoba3]"

	mutex := &sync.Mutex{}
	index := 0
	if len(os.Args) > 1 {
		fmt.Sscanln(os.Args[1], &index)
	}

	for range 1 {
		go process1Mutex(data3, index, mutex)
		go process2Mutex(data2, index, mutex)
		go process2Mutex(data1, index, mutex)
	}

	time.Sleep(5 * time.Second)
}

func main() {

	fmt.Println("Ciao, mondo! Questo programma è stato compilato con go.")

	//duck()
	//print_output()
	//stringa()
	//comando()
	//struttura()
	tempi()
	tempiMutex()
	startpackage.IntrospezioneMain()
	fmt.Println("Fine del programma.")

}
