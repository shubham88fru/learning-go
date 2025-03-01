- Every go file needs to be part of a package,and so, it needs a package directive up top.

- If we need our code to be an executable, it must be part of `package main` and must has a `main` function. While if we are simply building a library that has APIs/utility functions that other part of code or developers can use, it doesn't have to be part of the `package main`, nor does it need the `main` method.

- One go module consists of multiple packages.

- To initialize a module `go mod init <module_name>`. This will spit out a `go.mod` file.

- Once a module is initialized, we can build the go project (assuming it has main).

- To build run the command - `go build -o bin/` this command will build and redirect the binaries to be
  generated inside a bin folder. Specifying the output path is important because typically we wouldn't like
  to commit binaries to a git repo. So, we can entirely ignore the bin directory in gitignore.
