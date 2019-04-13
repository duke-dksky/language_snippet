package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func go_1_compatible() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		// read a chunk
		num, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if num == 0 {
			break
		}
		if _, err = fo.Write(buf[:num]); err != nil {
			panic(err)
		}
	}

}

func use_bufio() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	buf := make([]byte, 1024)
	for {
		// read a chunk
		num, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if num == 0 {
			break
		}
		if _, err = w.Write(buf[:num]); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}
}

func use_ioutil() {
	context, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("output.txt", context, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	//go_1_compatible()
	//use_bufio()
	use_ioutil()
}
