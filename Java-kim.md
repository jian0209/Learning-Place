## Java Virtual Machine

### Architecture

寄存器 ｜ PC

#### 需求

程序计数器， 指向下一条指令位置

object 对象放哪里？ 》 堆 heap

---

3个线程

- PC Register
- JVM Stack
- Native Method Stack

Heap

Method Area

- Runtime Constant Pool - 可以共享全局变量

---

### Program Counter Register

- 很小空间（几b罢了）

---

### JVM Stack

- 每个线程私有的
- 会被回收 ｜ 释放
- 默认1mb
- 全局变量不存在 （存在其他基本类型）
- boolean, char, byte, short, int, long, float, double
- Stack 不是先进先出， 是后进先出
- 每个函数有自己的战争
- 操作数 -  1 + n = 4
- - n = 3 （暂存）
  - 操作数就是临时出来的
  - 临时变数

---

### Native Method Stack

- 线程私有的
- 生命周期和线程相同
- 每个线程都有一个

---

### Java Heap

- 真正的内存在heap里面
- heap在开机的时候就会产生大空间
- 公用的
- 用来储存对象object
- 垃圾收集器管理的主要区域， 也被称为“gc堆”
- 垃圾收集器会把堆分成新生代和老年代： eden空间， from survivor空间， to survivor空间
- 大小可调节
- 通过XMS和-XMX控制

---

### Method Area

- 存放每个class的源信息
- A对象存放在里面， method area寻找A对象 加载他的信息
- 共享
- 运行时常量池(Runtime Constant Pool)
- 有效减少空间，垃圾次数

---

### Method Area-Instruction

```java
public class Demo {
	public static void foo(){
		int a = 1;
		int b = 2;
		int c = (a + b) * 5;
	}
}
```

---

### Class Loader Subsystem

- 双亲委派机制
  - 加载一个class的时候， 会默认用自定义的class loader来加载
  - 如果这层没有就往上
  - 如果有了， 把这层东西返回

---

### Excecution Engine

- 所有指令从方法群拿
- 加载到pc指令 然后一条一条执行
- cpu的指令级是二进制（谁来做转换工作）（java在执行指令的时候是Execution Engine转换成cpu可以看得懂的指令）
- 引进即使编译器
  - 会编译成local
  - 不用一行一行指令
  - 释放性能低
  - 这就是为什么java效能高

---

### Garbage Collection

- 如何判断对象是垃圾
  - 一个对象没有被引用
  - 判定（旧版）
    - 引用一次object 就 + 1
    - 取消一次object 就 — 1
    - 到 0 就被回收
    - 造成负向引用
  - 可达性分析（新版）
    - 用的object放在堆里面
    - 一个一个往下走
    - 如果发现已用着
    - 树状结构
    - 梗 - 往下 - 找无法用的 - 被回收
    - 静态变量。。。(Extend)
- 怎么回收
  - 分为 年轻代 ｜ 老年代 ｜ 永久代
  - 年轻代
    - front
    - survivor
    - 触发机制
      - 一直使用
      - 移至老年代
    - GC将不会解读
- permanent存放---

---

java真正的参数有哪些

怎么配置垃圾回收器

怎么优化垃圾回收的时间

怎么导出heap内存的数据

当进程的cpu满了，怎么排查哪个占了很大的应用

---

## 第二集： 数据库

---

## 需自行了解

JVM 的参数

1.zabbix 监控怎么监控系统指标的，alert 怎么做的，docker 容器是怎么做监控的，各类中间件（mysql,redis,rocketmq）怎么做监控的
2.loki 的数据怎么采集到grafana 的，grafana 怎么保存展示日志的
3.java 的代码从gitlab 仓库怎么打包发布到，开发，测试，生产环境上的
4.aws 的elb 接入进来的请求,怎么给nginx,到java 程序的
5.jump 跳板机怎么接入各环境的机器的，怎么保证操作可审计的

---

## 注意事项

```
LambdaUpdateWrapper<Order> updateWrapper = Wrappers.lambdaUpdate();
updateWrapper.eq(Order::getOrderNo, orderVO.getOrderNo());
updateWrapper.set(Order::getErrorStatus, orderVO.getErrorStatus());
```

