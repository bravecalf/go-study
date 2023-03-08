# 创建型模式
这些设计模式提供了一种在创建对象的同时隐藏创建逻辑的方式，提供了创建对象的机制， 能够提升已有代码的灵活性和可复用性。
## 1. 工厂方法模式
在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。

[factory design pattern](./creational/factory)

## 2. 抽象工厂模式
在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。

[abstract factory design pattern](./creational/abstract_factory)

## 3. 单例模式
这种模式涉及到一个单一的类，该类负责创建自己的对象，同时确保只有单个对象被创建。这个类提供了一种访问其唯一的对象的方式，可以直接访问，不需要实例化该类的对象。

[singleton design pattern](./creational/singleton)

## 4. 建造者模式
建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

[builder design pattern](./creational/builder)

## 5. 原型模式
原型模式（Prototype Pattern）是用于创建重复的对象，同时又能保证性能。这种模式是实现了一个原型接口，该接口用于创建当前对象的克隆。当直接创建对象的代价比较大时，则采用这种模式。

原型是一种创建型设计模式， 使你能够复制对象， 甚至是复杂对象， 而又无需使代码依赖它们所属的类。

[prototype design pattern](./creational/prototype)

## 6. 对象池模型
对象池设计模式是一种创建型设计模式，其中预先初始化和创建对象池并将其保存在池中。当需要时，客户端可以从池中请求一个对象，使用它，然后将它返回到池中。池中的对象永远不会被销毁。

[object pool design pattern](./creational/prototype)