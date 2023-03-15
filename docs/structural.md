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