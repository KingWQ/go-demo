### 一：ticker概述
1. ticker是一个定时触发的计时器
2. 它会以一个间隔往channel发送一个事件(当前时间)
3. 而channel的接收者可以以固定的时间间隔从channel中读取事件