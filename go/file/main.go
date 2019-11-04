package main

import (
	"bufio"
	"bytes"
	"fmt"
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

	buf := make([]byte, 10)
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
	r := bufio.NewReaderSize(fi, 32)

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

	buf := make([]byte, 39)
	for {
		// read a chunk
		num, err := r.Read(buf)
		fmt.Println(num)
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

func readLine() {
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
	r := bufio.NewReaderSize(fi, 10)
	for {
		var buffer bytes.Buffer
		var (
			line     []byte
			isPrefix bool
		)
		for {
			line, isPrefix, err = r.ReadLine()
			buffer.Write(line)
			if !isPrefix {
				break
			}
			if err != nil {
				break
			}
		}
		if err == io.EOF {
			break
		}
		fmt.Println(buffer.String())
	}
}

func readString() {
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
	r := bufio.NewReaderSize(fi, 10)

	var line string
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	if err != io.EOF {
		fmt.Println(err)
	}

}

func scanner() {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)
	scanner.Buffer(make([]byte, 2), bufio.MaxScanTokenSize)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	//go_1_compatible()
	//use_bufio()
	//use_ioutil()
	readLine()
	//readString()
	//scanner()
}
