# 结构型模式
这些设计模式关注类和对象的组合。继承的概念被用来组合接口和定义组合对象获得新功能的方式。
结构型模式介绍如何将对象和类组装成较大的结构， 并同时保持结构的灵活和高效。

## 1. 适配器模式
适配器是一种结构型设计模式， 它能使不兼容的对象能够相互合作。 适配器可担任两个对象间的封装器， 它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口。

[adapter design pattern](../structural/adapter)

## 2. 桥接模式
用于把抽象化与实现化解耦,使得实体类的功能独立于接口实现类,可将业务逻辑或一个大类拆分为不同的层次结构， 从而能独立地进行开发。

[bridge design pattern](../structural/bridge)

## 3. 过滤器模式
过滤器模式（Filter Pattern）或标准模式（Criteria Pattern）是一种设计模式，这种模式允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来。

[filter design pattern](../structural/filter)

## 4. 组合模式
组合模式（Composite Pattern）创建了一个包含自己对象组的类。该类提供了修改相同对象组的方式。可以使用它将对象组合成树状结构， 并且能像使用独立对象一样使用它们。

[composite design pattern](../structural/composite)

## 5. 装饰器模式
装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。装饰器模式一般用在基类功能封装不错，但使用的时候需要对功能进行一些加强，而这些加强版的功能也会被其它加强版需要，这种就比较适合。

[decorator design pattern](../structural/decorator)

## 6. 外观模式
外观模式（Facade Pattern）隐藏系统的复杂性，并向客户端提供了一个客户端可以访问系统的接口。这种模式涉及到一个单一的类，该类提供了客户端请求的简化方法和对现有系统类方法的委托调用。

[facade design pattern](../structural/facade)

## 7. 享元模式
享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。

[flyweight design pattern](../structural/flyweight)