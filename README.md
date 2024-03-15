# About Go

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

So, in comparison to C++, Go can be classified as a higher-level programming language, although it's not as high-level as languages like Python or JavaScript.

# Go is a statically typed

When we say that Go is a statically typed language, it means that variables must be explicitly declared along with their data types at compile time. Once a variable is declared with a certain data type, it cannot change to a different data type later in the program.

# Go is a compiled language

Go is a compiled language and JavaScript is an interpreted language.

To run a compiled file called `description` you might just type `./description` To run an interpreted Python program, you might have to type `./python ./description.py` That's because the first one is a compiled binary, consisting of just ones and zeros that your machine knows how to execute. In the second example, 'text editor.py' is just a file of Python source code, so it needs access to the Python interpreter.

### Compiled languages (Go, Java, Rust, C# or C++)

Compiled programs do not rely on any other programs to be able to run. They're usually written in a language like C, Rust, or Go, but are then compiled down to binary, the raw language of your computer's hardware.

Now, with a compiled language, what happens is you write your source code, and then you have a program called a compiler that will go through that source code and create a separate file that contains the machine code, and you just give me that file.

Compilation, Go source code is typically compiled into machine code. The Go compiler (go command) translates Go source code into machine code, producing an executable binary file that can be run directly by the target machine's processor.

In Go, the source code is translated into machine code, specific to the target platform (like Windows, Linux, or macOS), before execution.

This translation is done by the Go compiler (go command) into an executable binary file.

This binary file contains instructions that the computer's processor can directly execute, making it generally faster than interpreted languages.

Compilation typically happens before the program is executed, and the resulting binary can be run independently without needing the source code or the compiler.

### Interpreted languages (JavaScript, Ruby, Python or PHP)

Now, with an interpreted language, on the other hand, you don't compile your source code beforehand.

Whenever you've looked at a webpage with JavaScript, the JavaScript source code has been sent to you over the web.

Interpreted programs, on the other hand, do require other software to be executed. Python, Ruby, JavaScript, or shell scripts rely on an interpreter to read the code as they run when you execute a program on your machine.

In JavaScript, the source code is read and executed line by line by an interpreter within a runtime environment, typically a web browser or a server-side JavaScript engine like Node.js.
The JavaScript engine translates the JavaScript code into machine code or bytecode on-the-fly during execution.

There is no separate compilation step, and the code is executed directly from the source files each time it runs.

### Benefits Compiled language (Go)

Benefits of compiled code, once it's compiled, it's immediately ready to run, and you could send it to a hundred or a thousand or a hundred thousand different people; it's ready to go. It can be optimized for a CPU, so it can actually be faster, and you don't have to send your source code to everybody, which might be a good thing.

However, the downsides are if I compile it on a PC, that executable file won't work on a Mac.
In fact, it often needs to be compiled separately for different kinds of CPUs.
And when you're writing code to compile, it's an extra step that you have to take every time you want to test your program.

### Benefits Interpreted language (JavaScript or PHP)

Now, with interpreted code, the big benefits are I don't really care what kind of machine is on the other end because we don't provide machine code; we just send the source code, and we let the other side take care of it, so it can be more portable and more flexible across platforms.

It's also a little easier when testing because you just write your source code and then run it, letting the interpreter take care of converting it; there is no in-between compile step.

And it can be easier to debug when things go wrong because you always have access to all the source code.

However, it has its downsides too because everyone who needs to run that program on their machine has to have an interpreter for that language on their machine.

It also can be slower because you have to interpret it every time the program is run, and the source code is effectively public because you're sending it to everyone who needs to run that program.

### Performance

Compiled languages tend to have better performance since the translation to machine code happens beforehand, and there's no need for interpretation during runtime.

### Portability

Interpreted languages can be more portable because the source code can run on any platform with the appropriate interpreter or runtime environment. Compiled languages may need recompilation for different platforms.

### Debugging

Debugging compiled languages might be more challenging because the compiled binary might not contain all the debugging information available in the source code. Interpreted languages typically offer easier debugging since you can inspect the source code directly during runtime.
