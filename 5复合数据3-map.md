### 一：map概述
1. map是一种无序的键值对的集合，通过key快速检索数据
2. map是无序的，无法决定它的返回顺序，是因为map使用hash表实现

### 二：定义map
1. 声明变量：var 变量名称 map[key值类型]value值类型
2. make函数：变量名称 := make(map[key值类型]value值类型)
3. 如果不初始化map, 那么会创建一个nil map。
4. nil map不能用来存放键值对
