# SOLID - ISP

> ISP 即：The Interface Segregation Principle （接口隔离原则原则）

接口隔离原则：
- 客户端不应该依赖它不需要的接口；一个类对另一个类的依赖应该建立在最小的接口上。

问题由来：类 A 通过接口 I 依赖类 B，类 C 通过接口 I 依赖类 D，如果接口 I 对于类 A 和类 B 来说不是最小接口，则类 B 和类 D 必须去实现他们不需要的方法。

解决方案：将臃肿的接口I拆分为独立的几个接口，类A和类C分别与他们需要的接口建立依赖关系。也就是采用接口隔离原则。

---



### 违反示例：臃肿的接口
```C++
// C++
// Animal 接口包含不需要的方法
class Animal {
public:
    virtual void eat() = 0;     // 动物都需要进食
    virtual void swim() = 0;    // 不是所有动物都会游泳，如：鸟
    virtual void fly() = 0;     // 不是所有动物都会飞，如：狗
};

class Dog: public Animal {
    void fly() override {   // 被迫实现无用方法
        /* 抛出异常 或 空实现 */
    }

    // ...
};
```

```go
// Go
type Animal interface {
    Eat()
    Swim()
    Fly()
}

type Dog struct {}
func (d Dog) Fly() {
    // 被迫实现无用方法
    panic("狗不会飞")
}
```


> 问题：客户端被迫依赖不需要的方法。


---
### 冗余接口分离设计

```C++
// C++
class Eater { virtual void eat() = 0; };
class Swimmer { virtual void swim() = 0; };
class Dog: public Enter, public Swimmer {}; // 按需组合
```

```go
// Go
type Eater interface {
    Eat()
}

type Swimmer interface {
    Swim()
}

type Dog struct {}
// 仅实现需要的
func (d Dog) Eat() { /* 实现 */ }
func (d Dog) Swimmer() { /* 实现 */ }
```