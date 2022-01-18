# Lab 0 Map-Reduce Task

中国人民大学 王大林 sxwangdalin@ruc.edu.cn
## 实验结果
### 1. 完成 Map-Reduce 框架

运行结果截图如下

![image-20220109174657696](C:\Users\wdl\AppData\Roaming\Typora\typora-user-images\image-20220109174657696.png)

测试通过！总用时1059.919s

### 2.  基于 Map-Reduce 框架编写 Map-Reduce 函数

运行结果截图如下

![image-20220109180849484](C:/Users/wdl/AppData/Roaming/Typora/typora-user-images/image-20220109180849484.png)

测试通过！总用时139.656s

## 实验总结

### 1. debug总结

- 预留磁盘空间不足。实测本测试生成的tmp文件夹需要47.7GB磁盘空间，开始时因为磁盘空间不足而报错

![image-20220109181214594](C:/Users/wdl/AppData/Roaming/Typora/typora-user-images/image-20220109181214594.png)

- 运行时间超出go test的默认最大等待时间

开始时遇到这样的报错信息

![image-20220109181305768](C:/Users/wdl/AppData/Roaming/Typora/typora-user-images/image-20220109181305768.png)调查发现是因为go test 设定了默认的最大运行时间为10min，而由于机器性能的限制，本次测试的最大运行时间可能会超过10min，因此，需要通过```-testout 30m```来指定时间，防止因为超时而报错。

### 2. Map-Reduce总结

Example版本的方法如下：

- 第一轮
  - map：把<url，“”>这样的组合全部存储下来
  - reduce：把url相同的“”字符的个数记为value
- 第二轮
  - map：<'''', "url count">
  - reduce：构建map，把用空格分隔的url和count写入，随后排序取top10

对于上面的做法可以做如下改进：

- 在第一轮map中就可以开始进行局部count，然后在reduce中把各个局部和累加为最终的count
- 在第二轮map中，只输出局部top10，在reduce时仍就能保证结果不变

