### 一：go语言select概述
1. select是go中的一个控制结构，类似于用通信的switch语句
2. 每个case必须是一个通信操作，要么是发送，要么是接收
3. select随机执行一个可运行的case，如果没有case可运行，它将阻塞，直到case可以运行
4. 一个默认的子句应该总是可运行的

### 二：select语句语法
```
select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```
1. 每个case都必须是一个通信
2. 所有channel表达式都会被求值
3. 所有被发送的表达式都会被求值
4. 如果任意某个通信可以进行，它就执行，其他被忽略
5. 如果有多个case都可以运行，select会随机公平的选出一个执行，其他不会执行
6. 否则，如果有default子句，则执行该语句
7. 否则，如果没有default子句，select将阻塞，直到某个通信可以运行; go不会重新对channel或值进行求值