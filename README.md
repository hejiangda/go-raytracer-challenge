# The-Ray-Tracer-Challenge

春节期间阅读 **《The Ray Tracer Challenge》** 时用 **go** 语言写的代码，该书在oreilly在线图书馆的地址为：

https://learning.oreilly.com/library/view/the-ray-tracer/9781680506778/

官方地址：

https://pragprog.com/book/jbtracer/the-ray-tracer-challenge

这本书没有给出参考的代码，它完全要求读者自己写出代码完成挑战。它使用cucumber这个BDD（行为驱动开发）工具的Gherkin语言来提供测试案例。

在go语言中可以用godog来解析feature文件中的测试用例，但是由于自己水平有限，没弄懂godog怎么用最方便，就只用了go官方的testing包来写单元测试。

由于每次都需要先根据书中的测试样例先写测试，这些代码都是以TDD（测试驱动开发）的方式编写的。


## 目录结构说明
各章节对应的文件夹：

### 1. Tuples, Points, and Vectors

实现 Tuple 类型，Tuple 是 Point 和 Vector 的底层数据结构。w分量为1的为点，为0的为向量。

#### 对应的代码文件夹
- tuples
- chapter01

### 2. Drawing on a canvas

实现 canvas 并能输出 ppm 文件

该章最后会画出子弹发射后在风和重力作用下的运动轨迹图:

![轨迹](chapter02/chapter02.png)

#### 对应的代码文件夹
- canvas
- chapter02

## 3. Matrices

建立基础矩阵，矩阵的运算操作

#### 对应的代码文件夹
- matrices

## 4. Matrix Transformation

各类矩阵的变换：平移、旋转、放缩等

该章结束大作业是利用各类变换画一个表盘：

![表盘](chapter04/chapter04.png)

#### 对应的代码文件夹
- transformations
- chapter04

## 5. Ray-Sphere Intersections

实现球、光线投射、球的相交检测等

该章结束大作业是实现最基本的球的渲染：

![红球](chapter05/chapter05.png)

#### 对应的代码文件夹
- chapter05
- intersections
- objects
- rays
