package main

import (
	"bufio"
	"fmt"
//	"io"
	"io/ioutil"
	"os"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

//	p := fmt.Println
	p1 := fmt.Printf

/* 	dat, err := ioutil.ReadFile("/defer.txt")
	check(err)
	p(string(dat))

	f, err := os.Open("/defer.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	p1("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	p1("%d bytes @ %d: ", n2, o2)
	p1("%v\n", string(b2[:n2]))


	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	p1("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	p("5 bytes: %s\n", string(b4))

	f.Close() */



	//Writing Files

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	p1("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	p1("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	p1("wrote %d bytes\n", n4)

	w.Flush()

	//Writing Files
}