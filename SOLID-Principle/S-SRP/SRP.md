# SOLID - SRP

> SRP 即：The Single Responsibility Principle （单一职责原则）

单一职责原则：一个类/模块只负责完成一个职责/功能。

### 违反示例：一个类承担多个职责
```C++
// C++
class User{
public:
    void login(const std::string& username, const std::string& password) { /* 登录逻辑 */ }
    void saveToDB(const User& user) { /* 数据库操作 */ }
    void sendEmail(const User& user) { /* 发送邮件 */ }
};
```

```go
// Go
type User struct {}

func (u User) Login(username, password string) { /* 登录逻辑 */ }
func (u User) SaveToDB(user User) { /* 数据库操作 */ }
func (u User) SendEmail(user User) { /* 发送邮件 */ }
```


> 问题：修改登录逻辑、数据库存储或发送邮件中的任意一个部分，都需要修改 User 类。


---
### 单一职责设计

```C++
// C++
class UserAuth { /* 仅处理登录 */ };
class UserRepository { /* 仅处理存储 */ };
class EmailService { /* 仅处理通知 */ };
```

```go
type UserAuth struct{}
type UserRepository struct{}
type EmailService struct{}
```

### 简单示例请看同级目录 Go / Cpp
PS：

Go 语言中含有符合和不符合单一职责设计简单示例

C++ 中直接给出符合单一职责设计简单示例