# SOLID - LSP

> LSP 即：The Liskov Subsitution Principle （里氏替换原则）

里氏替换原则：
- 里氏替换原则：如果 S 是 T 的子类型，则类型 T 的对象可以替换为类型 S 的对象，而不会破坏程序。


简单来说，里氏替换原则要求子类(派生类)能够替换父类(基类)并且不影响程序的行为。也就是说，子类应该继承父类的所有属性和行为，并且可以在不改变程序逻辑的情况下进行扩展。

---



### 违反示例：子类破坏父类行为
```C++
// C++
// Bird 鸟类：特殊点，鸵鸟不会飞
class Bird {
public:
    virtual void fly() { /* 默认飞行 */ }
};

class Ostrich {
public:
    void fly() override {
        throw "鸵鸟不能飞~" // 破坏父类行为
    }
};
```

```go
// Go
type Bird struct{}
func (b Bird) Fly() { /* 默认飞行 */ }

type Ostrich struct{ Bird }
func (o Ostrich) Fly() {
    panic("鸵鸟不能飞！") // 使用时会导致运行时崩溃
}
```


> 问题：子类无法替换父类，导致运行时错误。


---
### 设计

```C++
// C++
class Bird {};
class FlyingBird : public Bird { void fly(); }; // 飞行能力单独分层
class Ostrich : public Bird {}; // 不继承 fly()
```

```go
// Go
type Bird struct{}
type FlyingBird struct{ Bird }
func (fb FlyingBird) Fly() { /* 飞行 */ }

type Ostrich struct{ Bird } // 无 Fly 方法
```

> 具体看.go文件
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

func Write(w Writer, p []byte) (int, error) {
    return w.Write(p)
}

// 里氏替换：
// ReadWriter(子类) 替换 Writer(父类) 
func Write(rw ReadWriter, p []byte) (int, error) {
    return rw.Write(p)
}

```