package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// This package is used to read and write data to and from buffers.
	// It provides buffered I/O operations, which can be more efficient than
	// unbuffered I/O operations.
	//
	// The bufio package is commonly used for reading and writing text files,
	// as well as for reading and writing binary data.

	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	//read bytes from the reader
	data := make([]byte, 20)
	n, err := reader.Read(data)

	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	//read a line from the reader
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading line: ", err)
		return
	}
	fmt.Printf("Read line: %s", line)

	writer := bufio.NewWriter(os.Stdout)

	// writing byte slice
	writeData := []byte("Hello, bufio package!\n")
	bw, err := writer.Write(writeData)
	if err != nil {
		fmt.Println("Error writing data: ", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", bw)
	err = writer.Flush() // Ensure all buffered data is written to the output
	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}

	// writing a string
	writeString := "This is a string written using bufio package.\n"
	bw, err = writer.WriteString(writeString)
	if err != nil {
		fmt.Println("Error writing string: ", err)
		return
	}
	fmt.Printf("Wrote string: %s", writeString)
	err = writer.Flush() // Ensure all buffered data is written to the output
	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}
}
