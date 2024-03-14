# About Go

Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency. It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

So, in comparison to C++, Go can be classified as a higher-level programming language, although it's not as high-level as languages like Python or JavaScript.

# Go is a statically typed

When we say that Go is a statically typed language, it means that variables must be explicitly declared along with their data types at compile time. Once a variable is declared with a certain data type, it cannot change to a different data type later in the program.

# Go is a compiled language

Go is a compiled language and JavaScript is an interpreted language.

### Run code Compiled languages vs Interpreted language

So, we need to get the source code like Go or JavaScript converted into machine code somehow before it can run on the computer.

### There are two main ways of doing this:

what's called compiling the source code and what's called interpreting the source code.

So, let's have a simple scenario: let's say it's just you and me. You have your computer, and I have my computer, and you're going to write a program that you want me to run.

### Compiled language

Now, with a compiled language, what happens is you write your source code, and then you have a program called a compiler that will go through that source code and create a separate file that contains the machine code, and you just give me that file.

This end result is sometimes referred to as an executable or an executive file because I can directly execute it. I can now just run your program; you keep your source code, and I never see it.

### Interpreted language

Now, with an interpreted language, on the other hand, you don't compile your source code beforehand.

You just give me a copy your source code, so I'll need my machine to interpret it whenever I want to run your program.

Now, an interpreter is different from a compiler; it does this on the fly. We can think of it as going through your source code line by line and processing it on the spot.

It does not save it as a separate machine code file.

Whenever you've looked at a webpage with JavaScript, the JavaScript has been sent to you over the web as source code onto your machine, and your web browser has just interpreted that JavaScript so it can run that code.

### So, which one's best?

### Compiled language

Benefits of compiled code: once it's compiled, it's immediately ready to run, and you could send it to a hundred or a thousand or a hundred thousand different people; it's ready to go. It can be optimized for a CPU, so it can actually be faster, and you don't have to send your source code to everybody, which might be a good thing.

However, the downsides are if I compile it on a PC, that executable file won't work on a Mac.
In fact, it often needs to be compiled separately for different kinds of CPUs.
And when you're writing code to compile, it's an extra step that you have to take every time you want to test your program.

### Interpreted language

Now, with interpreted code, the big benefits are I don't really care what kind of machine is on the other end because we don't provide machine code; we just send the source code, and we let the other side take care of it, so it can be more portable and more flexible across platforms.

It's also a little easier when testing because you just write your source code and then run it, letting the interpreter take care of converting it; there is no in-between compile step.

And it can be easier to debug when things go wrong because you always have access to all the source code.

However, it has its downsides too because everyone who needs to run that program on their machine has to have an interpreter for that language on their machine.

It also can be slower because you have to interpret it every time the program is run, and the source code is effectively public because you're sending it to everyone who needs to run that program.

### Other languages

Now, languages like PHP and JavaScript, indeed most languages with the word 'script' at the end, are usually interpreted, and languages like Java, C#, VB.NET, Python use this intermediate hybrid approach.

Now, whether a language is compiled or interpreted or somewhere in between is rarely a reason by itself to choose a language, but it can be something that you take into account.

If one main priority of your program is absolute maximum speed running on one single platform, you'll probably look at a compiled language.

If you're more interested in easily moving your code across multiple platforms, you're probably more interested in an interpreted one.

But more usually, you're driven more by what you need to do.

Do you need to build iPhone apps or Windows desktop apps or dynamic websites, or in our case, just learn the fundamentals of programming, and you let that decision drive the language choice, and the language choice will determine whether you're compiled, interpreted, or somewhere in the middle.

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
