# About Go

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

So, in comparison to C++, Go can be classified as a higher-level programming language, although it's not as high-level as languages like Python or JavaScript.

# Go is a statically typed

When we say that Go is a statically typed language, it means that variables must be explicitly declared along with their data types at compile time. Once a variable is declared with a certain data type, it cannot change to a different data type later in the program.

# Go is a compiled language

Go is a compiled language and JavaScript is an interpreted language.

### Compiled language (Go)

Compilation, Go source code is typically compiled into machine code. The Go compiler (go command) translates Go source code into machine code, producing an executable binary file that can be run directly by the target machine's processor.

In Go, the source code is translated into machine code, specific to the target platform (like Windows, Linux, or macOS), before execution. This translation is done by the Go compiler (go command) into an executable binary file.
This binary file contains instructions that the computer's processor can directly execute, making it generally faster than interpreted languages.
Compilation typically happens before the program is executed, and the resulting binary can be run independently without needing the source code or the compiler.

### Interpreted language (JavaScript)

First, the code is transpiled using babel or any other web pack.

This form of code is given to the Engine which converts it to AST(Abstract Syntax Tree).

This AST is then converted to the byte code which is understood by the machine. This is an Intermediate Representation(IR) which is further optimized by the JIT compiler.

After the optimization, the JS Virtual Machine Executes the code.

Hence we can conclude that JS code is executed in three phases.

In JavaScript, the source code is read and executed line by line by an interpreter within a runtime environment, typically a web browser or a server-side JavaScript engine like Node.js.
The JavaScript engine translates the JavaScript code into machine code or bytecode on-the-fly during execution.
There is no separate compilation step, and the code is executed directly from the source files each time it runs.

### The key differences between these approaches are

### Performance

Compiled languages tend to have better performance since the translation to machine code happens beforehand, and there's no need for interpretation during runtime.

### Portability

Interpreted languages can be more portable because the source code can run on any platform with the appropriate interpreter or runtime environment. Compiled languages may need recompilation for different platforms.

### Debugging

Debugging compiled languages might be more challenging because the compiled binary might not contain all the debugging information available in the source code. Interpreted languages typically offer easier debugging since you can inspect the source code directly during runtime.

# Why Go
