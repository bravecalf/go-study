# 行为型模式
行为模式负责对象间的高效沟通和职责委派。这些设计模式特别关注对象之间的通信。

## 1. 责任链模式
责任链是一种行为设计模式， 允许你将请求沿着处理者链进行发送， 直至其中一个处理者对其进行处理。
该模式允许多个对象来对请求进行处理， 而无需让发送者类与具体接收者类相耦合。 链可在运行时由遵循标准处理者接口的任意处理者动态生成。

[chain of responsibility design pattern](../behavioral/chain_of_responsibility)

## 2. 命令模式
命令模式（Command Pattern）是一种数据驱动的设计模式，它属于行为型模式。请求以命令的形式包裹在对象中，并传给调用对象。调用对象寻找可以处理该命令的合适的对象，并把该命令传给相应的对象，该对象执行命令。
该转换让你能根据不同的请求将方法参数化、 延迟请求执行或将其放入队列中， 且能实现可撤销操作。

[command design pattern](../behavioral/command)

## 3. 解释器模式
解释器模式（Interpreter Pattern）提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。

[interpreter design pattern](../behavioral/interpreter)

## 4. 迭代器模式
迭代器模式（Iterator Pattern）是 Java 和 .Net 编程环境中非常常用的设计模式。这种模式用于顺序访问集合对象的元素，不需要知道集合对象的底层表示。
迭代器模式是一种行为设计模式， 让你能在不暴露集合底层表现形式 （列表、 栈和树等） 的情况下遍历集合中所有的元素。

[iterator design pattern](../behavioral/iterator)

## 5. 中介者模式
中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。

[mediator design pattern](../behavioral/mediator)

## 6. 备忘录模式
忘录模式（Memento Pattern）保存一个对象的某个状态，以便在适当的时候恢复对象。备忘录模式让我们可以保存对象状态的快照。 你可使用这些快照来将对象恢复到之前的状态。 这在需要在对象上实现撤销-重做操作时非常实用。

[memento design pattern](../behavioral/memento)

## 7. 观察者模式
观察者模式（Observer Pattern）是一种行为设计模式， 允许你定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象。

[observer design pattern](../behavioral/observer)

## 8. 状态模式
在状态模式（State Pattern）中，类的行为是基于它的状态改变的。这种类型的设计模式属于行为型模式。 在状态模式中，我们创建表示各种状态的对象和一个行为随着状态对象改变而改变的 context 对象。
状态模式是一种行为设计模式， 让你能在一个对象的内部状态变化时改变其行为。

[state design pattern](../behavioral/state)