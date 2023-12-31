# Introduction to GO by Caleb Doxsey

## Chapter 1: Getting Started

### Key Takeaways
- Installing Go...using **Ubuntu 22.04 on WSL** to go through this book and although an install manual isn't available for WSL, it's basically the same as installing for Linux...durrhh🙄

- Either install by downloading the setup [here](https://go.dev/dl/) or just using the apt package manager - `sudo apt install golang`.  
I used the installer linked above cos why not??

- One key thing is: As of this document, the latest Go version is version 1.20.xx while latest on the apt package manager is version 1.18.xx.  
Not sure is there are any real differences that may affect my understanding of this book.  
If I find any, I'll list them here as I go:
  - In the book, `godoc` is used to get help in Go, but in 1.20, `godoc` isn't available, but `go doc` is...to get `godoc`, you'd have to install x-tools which essentially installs version 1.18.

- As a rule of thumb, generally start Go code or files with the like `package main`.

- Comments in Go can be;
  - Single line  - Starting with `//` until the end of the line.

  - Multi line - Starts with `/*` until it's terminated with `*/`.

- A Go Program is compiled and run by executing `go run <program name>` in the terminal from the directory the program is in or using the full/relative path to the program from another directory.

- The `import` statement is used to import external packages for use in your own Go program. i.e. `import "fmt"` which is pretty useful for basic io operations in Go. fmt is short for format.