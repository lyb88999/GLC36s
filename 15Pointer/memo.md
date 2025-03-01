# 关于指针的有限操作

## 复习

> 我们可以通过指针值无缝地访问到基本值包含的任何字段，以及调用与之关联的任何方法。这就是我们在编写Go程序过程中用得最频繁的“指针”了。
>
> 从传统意义上来说，指针是一个指向某个确切的内存地址的值。这个内存地址可以是任何数据或代码的起始地址。比如，某个变量、某个字段或某个函数。
>
> 刚才提到的只是其中的一种情况，在Go语言中还有其他几样东西可以代表“指针”。其中最贴近传统意义的就是`uintptr`类型了。该类型实际上是一个数值类型，也是Go语言内建的数据类型之一。
>
> 根据当前计算机的计算架构的不同，它可以存储32位或64位无符号整数，可以代表任何指针的位（bit）模式，也就是原始的内存地址。
>
> 再看Go语言标准库中的`unsafe`包，其中有一个类型叫`Pointer`，也代表了“指针”。
>
> `unsafe.Pointer`可以表示任何指向可寻址的值的指针，同时它也是前面提到的指针值和`uintptr`值之间的桥梁。也就是说，通过它我们可以在这两种值上进行双向的转换。这里有一个很关键的词——可寻址的，首先我们需要搞清楚这个词的确切含义。

## Q：列举一下Go语言中哪些值是不可寻址的？

## A：以下列表中的值都是不可寻址的

> 字面量--没有出现变量名，直接出现了值
>
> * 常量的值
> * 基本类型值的字面量
> * 算术操作的结果值
> * 对各种字面量的索引表达式和切片表达式的结果值。（例外：对切片字面量的索引结果值是可以寻址的）
> * 对字符串变量的索引表达式和切片表达式的结果值
> * 对字典变量的索引表达式的结果值
> * 函数字面量和方法字面量，以及对它们的调用表达式的结果值
> * 结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值
> * 类型表达式转换的结果值
> * 类型断言表达式的结果值
> * 接收表达式的结果值

## 问题解析

> 初看答案中的这些不可寻址的值好像并没有什么规律，不过别急，我们来梳理一下。
>
> 常量的值总是会被存储到一个确切的内存区域中，并且这种值肯定是不可变的。基本类型值的字面量也是一样，其实它们本就可以被视为常量，只不过没有任何标识符可以代表它们罢了。
>
> **第一个关键字：不可变的。** 由于Go语言中的字符串值也是不可变的，所以对于一个字符串类型的变量来说，基于它的索引或切片的结果值也都是不可寻址的，因为即使拿到了这种值的内存地址也改变不了什么。
>
> 算术操作的结果值属于一种**临时结果**。在我们把这种结果赋值给任何变量或常量之前，即使能拿到它的内存地址也是没有任何意义的。
>
> **第二个关键词：临时结果。**这个关键词能被用来解释很多现象，我们可以把各种对值字面量施加的表达式的求值结果都看作是临时结果。
>
> Go语言中的表达式有很多种，其中常用的包括以下几种：
>
> * 用于获得某个元素的索引表达式
> * 用于获得某个切片（片段）的切片表达式
> * 用于访问某个字段的选择表达式
> * 用于调用某个函数或方法的调用表达式
> * 用于转换值的类型的类型转换表达式
> * 用于判断值的类型的类型断言表达式
> * 向通道发送元素值或从通道那里接收元素值的接收表达式
>
> 我们把以上这些表达式施加在某个值字面量上一般都会得到一个临时结果，这种都是不可寻址的。
>
> 一个需要注意的例外是，对切片字面量的索引结果值是可寻址的。因为不论怎样，每个切片值都会持有一个底层数组，而这个底层数组中的每个元素都是有一个确切的内存地址的。
>
> 那么对切片字面量的切片结果值为什么却是不可寻址的呢？这是因为切片表达式总会返回一个新的切片值，而这个新的切片值在赋给变量之前属于临时结果。
>
> 上面一直说的是针对数组值、切片值或字典值的字面量的表达式会产生临时结果，如果针对的是数组类型或切片类型的变量，那么索引或切片的结果值就都不属于临时结果了，是可寻址的。
>
> 这因为变量的值本身就不是“临时的”。对比而言，值字面量还没有与任何变量绑定之前是没有落脚点的，我们无法以任何方式引用到它们，这样的值就是“临时”的。
>
> 再说一个例外。我们通过对字典类型的变量施加索引表达式，得到的结果值不属于临时结果，可是，这样的值却是不可寻址的。原因是，字典中的每个键-元素对的存储位置都可能会变化，而且这种变化是外界无法感知的。
>
> 我们都知道，字典中总会有若干个哈希桶用于均匀地存储键-元素对。当满足一定条件时，字典可能会改变哈希桶的数量，并适时地把其中的键-元素对搬运到对应的新的哈希桶中。
>
> 在这种情况下，获取字典中任何元素值的指针都是无意义的，也是**不安全**的。
>
> **第三个关键词：不安全的。**“不安全的”操作很可能会破坏程序的一致性，引发不可预知的错误，从而严重影响程序的功能和稳定性。
>
> 再来看函数。函数在Go语言中是一等公民，所以我们可以把代表函数或方法的字面量或标识符赋给某个变量、传给某个函数或从某个函数传出。但是，这样的函数和方法都是不可寻址的。一个原因是函数就是代码，是不可变的。
>
> 另一个原因是，拿到指向一段代码的指针是不安全的。此外，对函数或方法的调用结果值也是不可寻址的，这是因为它们都属于临时结果。
>
> 最后，如果我们把临时结果赋给一个变量，那么它就是可寻址的了。这样一来，取得的指针指向的就是这个变量持有的那个值了。

## Q：不可寻址的值在使用上有哪些限制？

## A：详情如下：

> 首当其冲的就是无法使用取地址符`&`获取它们的指针了。不过，对不可寻址的值加取地址都会使编译器报错，这个不太需要担心。
>
> 接下来看下面这个小问题，仍然以结构体类型`Dog`为例：

```go
func New(name string) Dog {
  return Dog{name}
}
```

> 这个函数会接受一个名为`name`的`string`类型的参数，并会用这个参数初始化一个`Dog`类型的值，最后返回该值。如果调用该函数，并直接以链式的手法调用其结果值的指针方法`SetName`，可以达到预期的效果吗？

```go
New("little pig").SetName("monster")
```

> 根据前面的内容，我们知道`New`函数所得到的结果属于临时结果，是不可寻址的。
>
> 可是，别忘了，我们在讲结构体类型及其方法的时候还说过，我们可以在一个基本类型的值上调用它的指针方法，因为Go语言会自动帮我们转译。
>
> 所以`dog.SetName("monster")`会被自动转译为`(&dog).SetName("monster")`，即先取`dog`的指针值，然后再在该指针值上调用`SetName`方法。
>
> 发现问题了吗？由于`New`函数的调用结果是不可寻址的，所以无法对它进行取址操作。因此，上面这行链式调用会让编译器报告两个错误，一个是果，即：不能在`New("little pig")`的结果值上调用指针方法；一个是因，即：不能取得`New("little pig")`的地址。
>
> 除此之外，我们都知道，Go语言中的`++`和`--`并不属于操作符，而分别是自增语句和自减语句的重要组成部分。
>
> 虽然Go语言规范中的语法定义是，只要在`++`或`--`的左边添加一个表达式，就可以组成一个自增或者自减语句，但是还有一个很重要的限制，那就是这个表达式的结果值必须是可寻址的。
>
> 不过这有一个例外，虽然对字典字面量和字典变量索引值的结果值都是不可寻址的，但是这样的表达式却可以被用在自增和自减语句中。
>
> 与之类似的规则还有两个。一个是，在赋值语句中，赋值操作符左边的表达式的结果值必须可寻址的，但是对字典的索引结果值也是可以的。
>
> 另一个是，在带有`range`子句的`for`语句中，在`range`关键字左边的表达式的结果值也都必须是可寻址的，不过对字典的索引结果值同样可以被用在这里。**上面这三条规则合并起来记忆就可以了。**

## Q：怎样通过`unsafe.Pointer`操纵可寻址的值？

## A：详情如下：

> 刚才说过，`unsafe.Pointer`是像`*Dog`类型的值这样的指针值和`uintptr`值之间的桥梁，那么我们怎样利用`unsafe.Pointer`的中转和`uintptr`的底层操作来操纵像`dog`这样的值呢？
>
> 首先说明，这是一项黑科技。它可以绕过Go语言的编译器和其他工具的重重检查，并达到潜入内存修改数据的目的。这并不是一种正常的编程手段，使用它会很危险，很有可能造成安全隐患。
>
> 我们总是应该优先使用常规代码包中提供的API去编写程序，当然也可以把像`reflect`以及`go/ast`这样的代码包作为备选项。作为上层应用的开发者，请谨慎地使用`unsafe`包中的任何程序实体。
>
> 首先看下面的代码：

```go
dog := Dog{"little dog"}
dogP := &dog
dogPtr := uintptr(unsafe.Pointer(dogP))
```

> 首先声明了一个`Dog`类型的变量`dog`，然后使用取地址操作符`&`，取出了它的指针值，并把它赋给了变量`dogP`。
>
> 最后使用了两个类型转换，先把`dogP`转换成了一个`unsafe.Pointer`类型的值，然后紧接着又把后者转换成了一个`uintptr`的值，然后把它赋给了变量`dogPtr`。这背后隐藏着一些转换规则，如下：
>
> 1. 一个指针值可以被转换为一个`unsafe.Pointer`类型的值，反之亦然。
> 2. 一个`uintptr`类型的值也可以被转换为一个`unsafe.Pointer`类型的值，反之亦然。
> 3. 一个指针值无法直接被转换为一个`uintptr`类型的值，反之亦然。
>
> 所以，对于指针值和`uintptr`类型值之间的转换，必须使用`unsafe.Pointer`类型的值作为中转。那么，我们把指针值转换为`uintptr`类型的值有什么意义呢？

```go
namePtr := dogPtr + unsafe.Offsetof(dogP.name)
nameP := (*string)(unsafe.Pointer(namePtr))
```

这里需要与`unsafe.Offsetof`函数搭配使用才能看出端倪。`unsafe.Offsetof`函数用于获取两个值在内存中的起始存储地址之间的偏移量，以字节为单位。

这两个值一个是某个字段的值，另一个是该字段所属的那个结构体值。我们在调用这个函数的时候，需要把针对字段的选择表达式传给它，例如`dogP.name`。

有了这个偏移量，又有了结构体值在内存中的起始存储位置，把它们相加我们就可以得到`dogP`的`name`字段值的起始存储地址了。这个地址由变量`namePtr`代表。

此后，我们可以再通过两次类型转换把`namePtr`的值转换成一个`*string`类型的值，这样就得到了指向`dogP`的`name`字段值的指针值。

你可能会问，我直接用取址表达式`&(dogP.name)`不就能拿到这个指针值了吗？为什么要绕这么一大圈呢？可以想象一下，如果我们根本就不知道这个结构体类型是什么，也拿不到`dogP`这个变量，那么还能去访问它的`name`字段吗？

答案是，只要有`namePtr`就可以，它就是一个无符号整数，但同时也是一个指向了程序内部数据的内存地址。它可能会给我们带来一些好处，比如可以直接修改埋藏得很深的内部数据。

但是，一旦我们有意或无意地把这个内存地址泄露出去，那么其他人就能够肆意地改动`dogP.name`的值，以及周围的内存地址上存储的任何数据了。



