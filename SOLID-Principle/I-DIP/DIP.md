# SOLID - DIP

> DIP 即：The Dependency Inversion Principle （依赖倒置原则）

依赖倒置原则：
- 高层模块不应该依赖于低层模块，两者都应该依赖于抽象
- 抽象不应该依赖于细节；细节应该依赖于抽象。

依赖倒置原则是一个应用广泛的原则，不仅在面向对象程序框架设计中（是核心原则），还是在架构系统中， 甚至是在社会活动构建组织等方面，都发挥重要作用。

高层、低层、抽象层可以表示为多种意义：

     高层：客户端、服务消费者、方法调用者……（可能变动、依赖性高）

     低层：服务端、服务提供者、方法实现者，细节实现者……（易变动、通常需要扩展）

     抽象：交互行为、策略、契约、流程、业务模型……（稳定、可重用）

---



### 违法示例：高层模块依赖底层模块
```C++
// C++
class LightBulb {
public:
    void turnOn() { /* 开灯 */ }
};

class Switch {
    LightBulb bulb; // 紧耦合
public:
    void press() { bulb.turnOn(); }
};
```

```go
// Go
type LightBulb struct{}
func (lb LightBulb) TurnOn() { /* 开灯 */ }

type Switch struct {
    bulb LightBulb // 直接依赖具体类型
}
func (s Switch) Press() {
    s.bulb.TurnOn()
}
```


> 问题：以电器为例，多数家电，比如点风扇也有开关，以上代码与底层模块紧密耦合，导致无法扩展。


---
### 破解依赖导致设计
> 简而言之，抽象不应该依赖于细节；细节应该依赖于抽象。

```C++
// C++
class Switchable {
    virtual void activate() = 0;
};

class LightBulb : public Switchable {
    void activate() { /* 实现逻辑 */ }
};

class Switch {
    Switchable& device; // 依赖于抽象，被抽象的事物一般具备一类事物的共有特征行为属性（比如：开关）
};
```

```go
// Go
type Switchable interface {
    Activate()
}

type LightBulb type {}
func (l LightBulb) Activate() {
    /* 实现逻辑 */
}

type Switch struct { device Switchable } // 依赖接口
```

### 案例参考：FastGO中的三层架构间的关系
> [FastGO](https://github.com/MortalSC/FastGO/tree/master/internal/apiserver)