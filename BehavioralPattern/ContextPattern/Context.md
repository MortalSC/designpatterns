## 上下文模式 —— Context

### 基本认识

上下文模式（Context Pattern）：用于在应用程序中传递和管理请求间的各种数据和信息。通过上下文对象，可以在不同组件之间共享数据、状态和配置信息，以实现解耦和提供灵活性。
> 上下文模式（Context Pattern）是一种 ​行为型设计模式，用于在系统组件间 ​传递和管理请求的上下文信息​（如请求ID、超时、取消信号、元数据等）。该模式通过将上下文对象作为参数显式传递，确保跨层级、跨组件的 ​一致性和可控性，尤其在并发和分布式系统中至关重要。

基本原理：
- **数据共享**：上下文模式用于在应用程序中传递和共享数据，便于不同组件之间访问和处理。
- **解耦组件**：通过上下文对象，不同组件之间可以松耦合地进行通信和数据传递。
- **请求范围**：上下文对象通常在请求的生命周期内存在，用于传递请求参数、认证信息等。

上下文模式是一种灵活且常用的设计模式，有助于组织和管理应用程序中的数据流和状态信息。

### 作用/用途
- **​传递请求级数据**：如用户身份、跟踪ID、语言偏好等。
- ​**控制请求生命周期**：通过超时（Timeout）或取消（Cancellation）信号终止下游操作。
- ​**解耦组件依赖**：避免在函数参数中传递大量不相干参数。
- ​**支持并发安全**：在多线程/协程环境下安全共享上下文状态。

### 特点
- **​不可变性**：上下文对象一旦创建不可修改，新配置需通过派生（Derive）新上下文实现。
- ​**链式传递**：通过父上下文派生子上下文，形成层级关系（如 context.WithTimeout()）。
- ​**统一接口**：提供标准方法（如 Done()、Deadline()）处理超时和取消。

### 解决的问题
- **参数透传污染**：避免在函数调用链中层层传递非核心参数（如请求ID）。
- ​**资源泄漏**：通过取消信号确保后台任务及时终止（如取消数据库查询）。
- **缺乏统一控制**：统一管理跨服务的超时、重试策略。

### 关键点
- 派生上下文：通过 WithCancel、WithTimeout、WithValue 等方法从父上下文派生子上下文。
- ​信号传播：父上下文的取消或超时自动触发所有子上下文的终止。
- ​线程安全：context.Context 本身不可变，协程安全。

​最佳实践：
将 context 作为函数的第一个参数（如 func Do(ctx context.Context, ...)）。避免存储 context 在结构体中，应显式传递。仅使用 WithValue 传递请求级数据，而非函数参数。

### Code
> 见同目录代码

### 进阶用法

#### 元数据传递​（使用 WithValue）：
```go
type key string
const requestIDKey key = "request_id"

// 存储值
ctx := context.WithValue(parentCtx, requestIDKey, "12345")

// 读取值
if id, ok := ctx.Value(requestIDKey).(string); ok {
    fmt.Println("Request ID:", id)
}
```


#### 链路追踪：
```go
func TrackRequest(ctx context.Context) context.Context {
    id := generateID()
    return context.WithValue(ctx, requestIDKey, id)
}
```

#### 组合控制：
```go
// 同时设置超时和手动取消
timeoutCtx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
```