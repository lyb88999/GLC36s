# Go语句及其执行规则（上）

## 前导内容：进程与线程

> 进程，描述的就是程序的执行过程，是运行着的程序的代表。换句话说，一个进程其实就是某个程序运行的一个产物，如果说静静地躺在那里的代码就是程序的话，那么奔跑着的、正在发挥着既有功能的代码就可以被称为进程。
>
> 再来说说线程。首先，线程总是在进程之内的，它可以被视为进程中运行着的控制流。
>
> 一个进程至少会包含一个线程。如果一个进程只包含了一个线程，那么它里面的所有代码都只会被串行地执行。每个进程的第一个线程都会随着进程的启动而创建，它们可以被称为其所属进程的主线程。
>
> 相对应的，如果一个进程中包含了多个线程，那么其中的代码就可以被并发地执行。除了进程的第一个线程之外，其他的线程都是由进程中已经存在的线程创建出来的。
>
> 也就是说，主线程之外的其他线程都只能由代码显式地创建和销毁。这需要我们在编写程序的时候进行手动控制，操作系统以及进程本身并不会帮我们下达这样的命令，它们只会忠实地执行我们的指令。
>
> 不过，在Go程序当中，Go语言的运行时（runtime）系统会帮助我们自动地创建和销毁系统级的线程。这里的系统级的线程指的就是我们刚才说过的操作系统提供的线程。
>
> 而对应的用户级线程指的是架设在系统级线程之上的，由用户（或者说我们编写的程序）完全控制的代码执行流程。用户级线程的创建、销毁、调度、状态变更以及其中的代码和数据都完全需要我们的程序自己实现和处理。
>
> 这带来了很多优势，比如，因为它们的创建和销毁并不用操作系统去做，所以速度很快，又比如，由于不用等着操作系统去调度它们的运行，所以往往会很容易控制并且很灵活。
>
> 但是，劣势也是有的，最明显的也最重要的一个劣势就是复杂，如果我们只用了系统级线程，那么我们只需要指明需要新线程执行的代码段，并且下达创建或销毁线程的指令就好了，其他的一切具体实现都会由操作系统代劳。
>
> 但是，如果使用用户级线程，我们就不得不既是指令下达者，又是指令执行者。我们必须负责与用户级线程有关的所有具体实现。
>
> 操作系统不但不会帮忙，还会要求我们的具体实现必须与它正确地对接，否则用户级线程就无法被并发地，甚至正确地运行。毕竟我们编写的所有代码最终都需要通过操作系统才能在计算机上执行。这听起来就很麻烦，不是吗？
>
> **不过别担心，Go语言不担有着独特的并发编程模型，以及用户级线程goroutine，还拥有强大的用于调度goroutine，对接系统级线程的调度器。**
>
> 这个调度器是Go语言运行时系统的重要组成部分，它主要负责统筹调配Go并发编程模型中的三个主要元素，即：G（goroutine的缩写）、P（processor的缩写）和M（machine的缩写）。
>
> 其中的M指的就是系统级线程，而P指的是一种可以承载若干个G，且能够使这些G适时地与M进行对接，并得到真正运行的中介。
>
> 从宏观上说，G和M由于P的存在可以呈现出多对多的关系。当一个正在与某个M对接并运行着的G，需要因某个事件（比如等待I/O或锁的解除）而暂停运行的时候，调度器总会及时地发现，并把这个G与那个M分离开，以释放计算资源供那些等待运行的G使用。
>
> 而当一个G需要恢复运行的时候，调度器又会尽快地为它寻找空闲的计算资源（包括M）并安排运行。另外，当M不够用时，调度器会帮我们向操作系统申请新的系统级线程，而当某个M已无用时，调度器又会负责把它及时销毁掉。
>
> 正因为调度器帮助我们做了很多事，所以我们的Go程序才总是能高效地利用操作系统和计算机资源。程序中的所有goroutine也都会被充分地调度，其中的代码也都会被并发地运行，即使这样的goroutine有数以十万计，也仍然可以如此。
>
> 下图是简化版的M、P、G之间的关系
>
> ![img](https://static001.geekbang.org/resource/image/9e/7d/9ea14f68ffbcde373ddb61e186695d7d.png?wh=1589*820)

## Q：什么是主goroutine，它与我们启用的其他goroutine有什么不同？

## A：详情如下：

```go
package main

import "fmt"

func main(){
  for i:=0; i<10; i++ {
    go func(){
      fmt.Println(i)
    }()
  }
}
```

> 在上面的程序中，只在`main`函数中写了一条`for`语句，这条`for`语句中的代码会迭代运行10次，并有一个局部变量`i`代表着当次迭代的序号。
>
> 在这条`for`语句中仅有一条`go`语句，这条`go`语句中也仅有一条语句。这条最里面的语句调用了`fmt.Println`函数并想要打印出变量`i`的值。
>
> 那么，这个命令源码文件被执行后会打印出什么内容？
>
> 这道题的典型回答是：不会有任何内容被打印出来。

## 问题解析：

与一个进程总有一个主线程类似，每一个独立的Go程序在运行时也总会有一个主goroutine。这个主goroutine会在Go程序的运行准备工作完成后被自动地启用，并不需要我们做任何手动的操作。

每条`go`语句一般都会携带一个函数调用，这个被调用的函数常常被称为`go`函数。而主goroutine的`go`函数就是那个作为程序入口的`main`函数。

一定要注意，`go`函数真正被执行的时间，总会与其所属的`go`语句被执行的时间不同。当程序执行到一条`go`语句的时候，Go语言的运行时系统，会先试图从某个存放空闲的G的队列中获取一个G（也就是goroutine），它只有在找不到空闲的G的情况下才会去创建一个新的G。

这就是为什么上面会说“启用”一个goroutine，而不说“创建”一个goroutine。已经存在的goroutine总是会被优先复用。

然而，创建G的成本也是非常低的。创建一个G并不会像新建一个进程或者一个系统级线程那样，必须通过操作系统的系统调用来完成，在Go语言的运行时系统内部就可以完全做到了。

在拿到了一个空闲的G之后，Go语言运行时系统会用这个G去包装当前的那个`go`函数，然后再把这个G追加到某个存放可运行的G的队列中。

这类队列中的G总是会按照先入先出的顺序，很快地由运行时系统内部的调度器安排运行。虽然这会很快，但是由于上面所说的那些准备工作还是不可避免的，所以耗时还是存在的。

因此，`go`函数的执行时间总是会明显滞后于它所属的`go`语句的执行时间。当然了，这里所说的“明显滞后”是对于计算机的CPU时钟和Go程序来说的。

在说明了原理之后，我们再来看这种原理下的表象。请记住，只要`go`语句执行完毕，Go程序完全不会等待`go`函数的执行，它会立刻去执行后面的语句，这就是所谓的异步并发地执行。

上面的`for`循环会以很快的速度执行完毕，当它执行完毕时，那10个包装了`go`函数的goroutine往往还没有获得运行的机会。

请注意，`go`函数中的那个对`fmt.Println`函数的调用是以`for`语句中的变量`i`作为参数的。可以想象一下，如果当`for`语句执行完毕的时候，这些`go`函数还没有执行，那么它们引用的变量`i`的值将会是什么？

它们都会是`10`，对吗？这个题的答案是打印出10个`10`，是这样吗？

在确定最终的答案之前，你还需要知道一个与主goroutine有关的重要特性，即：一旦主goroutine中的代码执行完毕，当前的Go程序就会结束运行。

如此一来，如果在Go程序结束的那一刻，还有goroutine未得到运行机会，那么它们就真的没有运行机会了，它们中的代码也不会被执行了。

严谨来讲，Go语言并不会去保证这些goroutine会以怎样的顺序运行。由于主goroutine与我们手动启用的其他goroutine一起接受调度，又因为调度器很有可能会在goroutine中的代码执行了一部分的时候暂停，以期所有的goroutine有更公平的运行机会。

所以哪个goroutine先执行完，哪个goroutine后执行完往往是不可预知的，除非我们使用了某种Go语言提供的方式进行了人为干预。然而，在这段代码中，我们并没有进行人为干预。