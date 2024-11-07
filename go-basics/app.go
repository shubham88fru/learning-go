/*
Code that has be built as an executable,
must belong to `package main` and have a
main method.

Files that are part of package main, are sort of the
entrypoint when executing
*/
package main

import "fmt"

/*
	Entrypoint for this file. `main` method
	in a file is invoked when the file is executed.

	It's important to note that, although we can spread our
	code in multiple files, where each file could be part of
	package main, but there must be just one main function of
	all the files that are part of package main.
*/
func main() {
	fmt.Print("Hello World")
}