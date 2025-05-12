# SOLID - OCP

> OCP 即：The Open Closed Principle （开闭原则）

开闭原则：对扩展开放，对修改关闭（新增功能时不修改已有代码）。

说人话：希望是通过在已有代码基础上扩展代码，而非修改代码的方式完成新功能的添加。

开闭原则，并不是说完全杜绝修改，而是尽可能不修改或者以最小的代码修改代价来完成新功能的添加。

`PS：最容易违反该规则的代码习惯就是大量使用 if else 条件语句，仅通过条件判断，多数情况下的内容新增都是入侵式修改。`

### 违法示例：一个类承担多个职责
```C++
// C++
class AreaCalculate {
public:
    double calculateArea(cosnt std::string& shapeType) {
        if (shapeType == "circle") { /* 计算圆形 */ }
        else if (shapeType == "square") { /* 是计算正方形 */ }

        // TODO：新增其他图形计算
    }
};
```

```go
// Go
func calculateArea(shapeType string) float64 {
    if (shapeType == "circle") { /* 计算圆形 */ }
    else if (shapeType == "square") { /* 是计算正方形 */ }

    // TODO：新增其他图形计算
}
```


> 问题：每次新增图形类型都要修改核心逻辑，容易引入错误。


---
### 开闭设计

```C++
// C++
class Shape {
public: 
    virtual double area() = 0;  // 要求子类必须实现，通过继承来完成高可拓展性
};

class Circle: public Shape {
public:
    double area()  { /* 计算圆形 */ }
};

class Square: public Shape {
public:
    double area()  { /* 是计算正方形 */ }
};

class 拓展图形: public Shape {
public:
    double area() { /* 对应的面积算法 */ }
};
```

```go
type Shape interface {  // 图形接口，只要实现接口下的方法，即被视为实现了对应的接口
    Area() float64
}

type Circle struct {}
func (c Circle) Area() float64 { /* 计算圆形 */ }

type 拓展图形 struct {}
func (x 拓展图形) Area() float64 { /* 对应的面积算法 */ }
```