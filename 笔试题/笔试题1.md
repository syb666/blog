### 1.什么是GIL
Python代码的执行由Python 虚拟机(也叫解释器主循环，CPython版本)来控制，Python 在设计之初就考虑到要在解释器的主循环中，同时只有一个线程在执行，即在任意时刻，只有一个线程在解释器中运行。对Python 虚拟机的访问由全局解释器锁（GIL）来控制，正是这个锁能保证同一时刻只有一个线程在运行。
在多线程环境中，Python 虚拟机按以下方式执行：
1. 设置GIL
2. 切换到一个线程去运行
3. 运行：
    * 指定数量的字节码指令，或者
    * 线程主动让出控制（可以调用time.sleep(0)）
4. 把线程设置为睡眠状态
5. 解锁GIL
6. 再次重复以上所有步骤
 
在调用外部代码（如C/C++扩展函数）的时候，GIL 将会被锁定，直到这个函数结束为止（由于在这期间没有Python 的字节码被运行，所以不会做线程切换）。
### 2.Python中的self、@staticmethod、@classmethod的区别
* 类成员方法必须经过实例化对象后才能调用
* 类成员方法的调用：
类方法：类名.方法名()
实例方法：类名().方法名()
* （1）self：表示是当前类对象本身的实例，写在函数定义的第一个参数位置。
  * 1.在当前函数内要访问当前类的属性和函数，需要通过关键字self.属性或self.方法()来调用。
  * 2.如果是其他类或包需要调用当前函数，可以通过实例方法来调用。
* （2）@staticmethod：装饰器，写在函数定义的前面。

  * 1.在当前函数内要访问当前类的属性和函数，需要通过类名.属性或类名.方法()来调用。
  * 2.如果是其他类或包需要调用当前函数，可以通过类方法和实例方法来调用。
* （3）@classmethod：装饰器，写在函数定义的前面。

  * 1.在当前函数内要访问当前类的属性或函数，需要通过在函数参数中定义一个参数来传递当前类对象，参数名自定，如参数.属性或参数.方法()。
  * 2.如果是其他类或包需要调用当前函数，可以通过类方法和实例方法来调用。
实例
```python
class C:
	name = 'I am 18.'
	def func1(self):
		print(self.name)  # 调用当前类的属性，通过self.属性调用
        #静态方法装饰器
	@staticmethod
	def func2():
		print(C.name)  # 调用当前类的属性，通过类名.属性调用
        #类方法装饰器
	@classmethod
	def func3(cls):
		print(cls.name)  # 调用当前类的属性，通过当前类参数.属性调用
```
```python
>>> C().func1()
I am 18.
>>> C.func2()
I am 18.
>>> C().func2()
I am 18.
>>> C.func3()
I am 18.
>>> C().func3()
I am 18.
```
### 2.Python里面如何拷贝一个对象，并解析深浅拷贝
* 直接赋值：其实就是对象的引用（别名）。
* 浅拷贝(copy)：拷贝父对象，不会拷贝对象的内部的子对象。
* 深拷贝(deepcopy)： copy 模块的 deepcopy 方法，完全拷贝了父对象及其子对象。
字典浅拷贝实例
```python
>>>a = {1: [1,2,3]}
>>> b = a.copy()
>>> a, b
({1: [1, 2, 3]}, {1: [1, 2, 3]})
>>> a[1].append(4)
>>> a, b
({1: [1, 2, 3, 4]}, {1: [1, 2, 3, 4]})
```
#
```python
>>>import copy
>>> c = copy.deepcopy(a)
>>> a, c
({1: [1, 2, 3, 4]}, {1: [1, 2, 3, 4]})
>>> a[1].append(5)
>>> a, c
({1: [1, 2, 3, 4, 5]}, {1: [1, 2, 3, 4]})
```
1. b = a: 赋值引用，a 和 b 都指向同一个对象。<br/>
![](https://github.com/syb666/blog/blob/master/%E7%AC%94%E8%AF%95%E9%A2%98/images/1123031-20170323105332346-341380067.png)
2. b = a.copy(): 浅拷贝, a 和 b 是一个独立的对象，但他们的子对象还是指向统一对象（是引用）<br/>
![](https://github.com/syb666/blog/blob/master/%E7%AC%94%E8%AF%95%E9%A2%98/images/1123031-20170323105350955-24694016.png)
3. b = copy.deepcopy(a): 深度拷贝, a 和 b 完全拷贝了父对象及其子对象，两者是完全独立的。<br/>
![](https://github.com/syb666/blog/blob/master/%E7%AC%94%E8%AF%95%E9%A2%98/images/1123031-20170323105418361-983703074.png)

以下实例是使用 copy 模块的 copy.copy（ 浅拷贝 ）和（copy.deepcopy ）:
```python
#!/usr/bin/python
# -*-coding:utf-8 -*-
import copy
a = [1, 2, 3, 4, ['a', 'b']] #原始对象 
b = a                       #赋值，传对象的引用
c = copy.copy(a)            #对象拷贝，浅拷贝
d = copy.deepcopy(a)        #对象拷贝，深拷贝
a.append(5)                 #修改对象a
a[4].append('c')            #修改对象a中的['a', 'b']数组对象
print( 'a = ', a )
print( 'b = ', b )
print( 'c = ', c )
print( 'd = ', d )
```
结果
```python
a =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
b =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
c =  [1, 2, 3, 4, ['a', 'b', 'c']]
d =  [1, 2, 3, 4, ['a', 'b']]
```
### 4.Python里面的search()和match()的区别
1. match()函数只检测字符串开头位置是否匹配，匹配成功才会返回结果，否则返回None
```python
import re
print(re.match("func", "function"))
# 打印结果 <_sre.SRE_Match object; span=(0, 4), match='func'>

print(re.match("func", "function").span())
# 打印结果  (0, 4)

print(re.match("func1", "function"))
# 打印结果 None

注意：print(re.match("func1", "function").span())会报错，因为取不到span
```
2. search()函数会在整个字符串内查找模式匹配,只到找到第一个匹配然后返回一个包含匹配信息的对象,该对象可以通过调用group()方法得到匹配的字符串,如果字符串没有匹配，则返回None。
```python
import re
print(re.search("tion", "function"))
# 打印结果 <_sre.SRE_Match object; span=(4, 8), match='tion'>

print(re.search("tion", "function").span())
# 打印结果  (4, 8)

print(re.search("tion1", "function"))
# 打印结果 None

注意：print(re.search("tion1", "function").span())会报错，因为取不到tion1
```
### 5.简述迭代器生成器以及他们之间的区别
迭代器是一个更抽象的概念，任何对象，如果它的类有 next 方法和 iter 方法返回自己本身，对于 string、list、dict、tuple 等这类容器对象，使用 for 循环遍历是很方便的。在后台 for 语句对容器对象调用 iter()函数，iter()是 python 的内置函数。iter()会返回一个定义了 next()方法的迭代器对象，它在容器中逐个访问容器内元素，next()也是 python 的内置函数。在没有后续元素时，next()会抛出一个 StopIteration 异常。

生成器（Generator）是创建迭代器的简单而强大的工具。它们写起来就像是正规的函数，只是在需要返回数据的时候使用 yield 语句。每次 next()被调用时，生成器会返回它脱离的位置（它记忆语句最后一次执行的位置和所有的数据值）

区别：生成器能做到迭代器能做的所有事,而且因为自动创建了 iter()和 next()方法,生成器显得特别简洁,而且生成器也是高效的，使用生成器表达式取代列表解析可以同时节省内存。除了创建和保存程序状态的自动方法,当发生器终结时,还会自动抛出 StopIteration 异常。
总结：
#####	生成器具有以下几个特点：
* 因为自动创建了 iter()和 next()方法,生成器显得特别简洁。
* 生成器是高效的。
* 使用生成器表达式取代列表解析可以同时节省内存。
* 会自动抛出 StopIteration 异常。

### 6.什么是协程，Python中的协程是如何实现的
协程不是进程或线程，其执行过程更类似于子例程，或者说不带返回值的函数调用。 一个程序可以包含多个协程，
可以对比与一个进程包含多个线程，因而下面我们来比较协程和线程。我们知道多个线程相对独立，有自己的上下文，
切换受系统控制；而协程也相对独立，有自己的上下文，但是其切换由自己控制，由当前协程切换到其他协程由当前协程来控制。 

简单的来讲就是在一个线程中的某个函数，可以在任何地方保存当前函数的一些临时变量等信息，然后切换到另外一个函数中执行，
注意不是通过调用函数的方式做到的，并且切换的次数以及什么时候再换到原来的函数都由开发者自己确定
协程的实现方式有很多,这里我们来列举三种基本方式. 
第一种 利用 yield 来实现协程:
```python
import time
def work1():
    # 循环打印数字1 
    while True:

        print("-----1-----")
        # yield可以暂时挂起该函数,跳转到调用该函数的下方
        yield
        # 延迟一秒以便观察
        time.sleep(1)
def work2():

    while True:

        print("-----2-----")
        yield
        time.sleep(1)
th1 = work1()
th2 = work2()

while True:
    # 唤醒被挂起的函数
    next(th1)
    next(th2)
```
第二种 利用 greenlet 模块实现:
```python
import time
import greenlet
def work1():
    # 循环打印字符串
    while True:

        print("----1----")
        # 启动th2
        th2.switch()
        time.sleep(1)
def work2():
    # 循环打印字符串
    while True:

        print("----2----")
        # 启动th1
        th1.switch()
        time.sleep(1)
# 创建携程
th1 = greenlet.greenlet(work1)
th2 = greenlet.greenlet(work2)
# 启动携程
th1.switch()
```
greenlet已经实现了协程，但是这个还得人工切换，是不是觉得很麻烦,python还有一个比greenlet更强大的并且能够自动切换任务的模块 gevent 
第三种:
```python
import gevent
import time
def work1():
    # 循环打印
    while True:

        print("----1----")
        # 破解sleep 使sleep不再阻塞
        gevent.sleep(1)
def work2():
    while True:
        print("----2----")
        gevent.sleep(1)
# 创建并开始执行携程
th1 = gevent.spawn(work1)
th2 = gevent.spawn(work2)
# 阻塞等待携程结束
gevent.joinall([th1,th2])
--------------------- 
```
### 7.什么是装饰器，请使用装饰器实现singletion。
装饰器可以抽离出大量函数中与函数功能本身无关的雷同代码并继续重用，为已经存在的对象添加额外的功能
在不改变函数代码、功能结果的前提下，为函数增添某些功能。通过python的闭包及一切皆对象，函数也为对象的特性。将函数对象传入wrapper函数，返回一个装饰后的函数。可以极大简化代码，不必重复写代码。




### 8.请使用Python实现快速排序
```python
#快排的主函数，传入参数为一个列表，左右两端的下标
def QuickSort(list,low,high):
    if high > low:
        #传入参数，通过Partitions函数，获取k下标值
        k = Partitions(list,low,high)
        #递归排序列表k下标左侧的列表
        QuickSort(list,low,k-1)
        # 递归排序列表k下标右侧的列表
        QuickSort(list,k+1,high)
def Partitions(list,low,high):
    left = low
    right = high
    #将最左侧的值赋值给参考值k
    k = list[low]
    #当left下标，小于right下标的情况下，此时判断二者移动是否相交，若未相交，则一直循环
    while left < right :
        #当left对应的值小于k参考值，就一直向右移动
        while list[left] <= k:
            left += 1
        # 当right对应的值大于k参考值，就一直向左移动
        while list[right] > k:
            right = right - 1
        #若移动完，二者仍未相遇则交换下标对应的值
        if left < right:
            list[left],list[right] = list[right],list[left]
    #若移动完，已经相遇，则交换right对应的值和参考值
    list[low] = list[right]
    list[right] = k
    #返回k值
    return right
list_demo = [6,1,2,7,9,3,4,5,10,8]
print(list_demo)
QuickSort(list_demo,0,9)
print(list_demo)
```

### 9.简述select和epoll的原理和区别
当用户进程发起一个read请求时，实际是调用系统接口，让内核去读取数据到内核空间，再copy到用户空间，系统给进程return 一个OK,然后用户进程去取，所以在用户发起请求后socket是阻塞状态，而select和epoll这两个函数，的功能就是，将socket设置为不阻塞后，让操作系统帮忙监控他们所维护的所有socket，只有socket准备就绪就返回，这两函数通过不断的轮询自己所维护的socket，将准备好的socket返回用户进程。select与epoll的区别就是，select有最大维护链接限制，Linux下一班为1024，可以通过修改宏定义改变，但是太大了依然会影响速度，而epoll没有限制

### 10.简述Python的垃圾回收机制
python采用引用计数的方式。来回收不被使用的对象。每隔一段时间检查一次，当对象的引用计数为0是就执行del 析构函数，来释放该内存，到预先申请的内存空间中去，防止内存碎片的产生。

分代回收：存活时间越久的对象，越不可能在后面的程序中变成垃圾。出于信任和效率，对于这样一些“长寿”对象，我们相信它们的用处，所以减少在垃圾回收中扫描它们的频率。Python将所有的对象分为0，1，2三代。所有的新建对象都是0代对象。当某一代对象经历过垃圾回收，依然存活，那么它就被归入下一代对象。垃圾回收启动时，一定会扫描所有的0代对象。如果0代经过一定次数垃圾回收，那么就启动对0代和1代的扫描清理。当1代也经历了一定次数的垃圾回收后，那么会启动对0，1，2，即对所有对象进行扫描。

这两个次数即上面get_threshold()返回的(700, 10, 10)返回的两个10。也就是说，每10次0代垃圾回收，会配合1次1代的垃圾回收；而每10次1代的垃圾回收，才会有1次的2代垃圾回收。

### 11.写一个简单的python socket编程

### 12.简述Python上下文管理器原理，并用上下文管理器实现将“hello world”写入文件的功能。

### 13.简述MyISAM和InnoDB的特点

### 14.简述一致性哈希原理和它要解决的问题

### 15.用python将'123456'反转成'654321'。
```python
old = [x for x in '123456']
new_str = ''.join(old[::-1]) #--->‘654321’
```

### 16.利用python执行shell命令并取得返回结果。
```python
import os
return os.popen(argv[1]) #---->返回结果的
# os.system( command ) 不返回结果的
```
### 17.用python继承process，写一个并行执行任务的类，并写出使用过程。

### 18.请列出你使用过的HA方案。

### 19.请列出你了解的Web服务器负载架构。

### 20.要求：列出一个班出平均分数超过60分的男生女生各占多少位。

表名 Student

字段名 name  sex score

### 21.如何判断一个邮箱是否合法。
```python
pattern = '^[A-Za-z_]\w+@\w+.com$' #标准为以下划线或字母开头
pattern_1 = '^(\w|)\w+@\w+.com$' #数字字母下划线开头都能匹配，如果需要限制长度 '^(\w|)\w+{9,14}@\w+.com$' 长度为9到14
mail = 'lina1994@outlook.com'
if re.match(pattern_1, mail):
print('该邮箱合法')
```

### 22.请实现一个装饰器，限制该函数被调用的频率，如10秒一次。

### 23.请描述一下，tuple,list,dict,set 的特点。
一句话来概括四种数据类型的区别是：
tuple是一个不可改变的list，
set是一个没有Value的dict，list和set的数据是可变的，
tuple和dict的数据是不可变的！

### 24.请说一下对迭代器与生成器的理解。
按照某种算法推算出列表元素，可以一直一边循环一边计算出列表元素的机制，称之为生成器：generator。一个简单生成器构造：
列表是L = [a*2 for a in range(10)]  生成器把中括号 [] 换成小括号()就好了，G = (a*2 for a in range(10)) 这样就得到了一个生成器,[.....]生成列表，(.....)生成生成器

而生成器不但可以作用于for循环，还能被next()函数不断调用返回下一个值，直到抛出StopIteration错误无法继续
可以被next()函数调用并不断返回下一个值的对象就称为迭代器Iterator
list、idct、str虽然可以用for循环，可迭代，但是没有nex()方法就不是Iterator迭代器

```python
>>> g = (a*a for a in range(10))
>>> g
<generator object <genexpr> at 0x0000001D925F3308>     #1
>>> for v in g:         #2
...     print(v)
 1
4
9
16
25
36
49
64
81
```
当然我们还可以用yield来进行改装，把它变成一个生成器对象，yield: 在函数执行时会给函数返回generator对象，可以通过该对象的obj.__next__()方法或next(obj)来启动函数
```python
def fib(n):
    i, a, b = 0, 0, 1
    while i < n:
        yield b         #在这儿用yield就好了
        a, b = b, a+b
        i = i + 1
    return '--done--'    
f = fib(6)
while True:
    try:
        res = next(f)
        print(res)
    except    StopIteration as e:
        print(e.value)
        break
#运行结果为
1
1
2
3
5
--done--
```
### 25.请用python实现单例模式，至少两种方式。


### 26.就你熟悉的Web框架，讲一讲如何维持登录状态的。


### 27.请说一说lambda函数的作用，请使用lambda和reduce实现1到100的累加。（**）

Python支持一种有趣的语法，它允许你快速定义单行的最小函数。这些叫做lambda的函数是从Lisp中借用来的，可以被用在任何需要函数的地方。
lambda 函数是一个可以接收任意多个参数（包括可选参数）并且返回单个表达式值的匿名函数。 （注意：lambda 函数不能包含命令，它们所包含的表达式也不能超过一个）
优点:
1、lambda函数比较轻便，即用即扔，很适合需要完成某一项简单功能，但是这个简单的功能只在此一处使用，连名字都很随意的情况下；
2、lambda是匿名函数，一般用来给filter，map，reduce这样的函数式编程服务；
3、作为回调函数，可以传递给某些应用，比如消息处理等。

```python
from functools import reduce
print(reduce(lambda x,y:x+y,range(1,101)))
```



### 28.用正则实现匹配手机号（包含手机号码前带86和+86的情况）。

### 29.

1
2
3
4
5
6
7
8
```python
import copy
a = [1,2,3,[4,5],6]
b=a
c=copy.copy(a)
d=copy.deepcopy(a)
b.append(10)
c[3].append(11)
d[3].append(12)
#请问a,b,c,d的值为？
```

### 30.现有字典d={'a':26,'g':20,'e':22,'c':24,'d':23,'f':21,'b':25}请按照字段中的value进行排序。
sort()与sorted()的不同在于，sort是在原位重新排列列表，而sorted()是产生一个新的列表。
sorted(d.items(),key = lambda x:x[1])
```python
>>> sorted(d.items(),key = lambda x:x[1])
[('g', 20), ('f', 21), ('e', 22), ('d', 23), ('c', 24), ('b', 25), ('a', 26)]
>>> d={'a':26,'g':20,'e':22,'c':24,'d':23,'f':21,'b':25}
>>> sorted(d.items(),key = lambda x:x[1])
[('g', 20), ('f', 21), ('e', 22), ('d', 23), ('c', 24), ('b', 25), ('a', 26)]
```
### 31.解释top命令和vmstat命令。

### 32.mysql高可用方案有哪些，备份方案有哪些，有什么优缺点？

### 33.linux基础问题：

　　- 怎么查看用户登录日志。

　　- linux中的utmp,wtmp,lastlog,message各文件的作用。

　　- 列举你属性的服务器性能查看命令。

　　- linux服务器间怎么实现无密码登录，列举操作步骤

### 34.画出TCP三次握手，四次挥手断开示意图。

### 35.叙述mysql半同步复制原理。

### 36.有这样一个文本文件，它的路径是baseDir,它的名字是test.txt，要求应with方式进行打开，并打印每一行文本，并要求文件路径考虑跨平台问题。

### 37.Python是如何进行类型转换的。

### 38.请写出一段python代码实现删除一个list里面的重复元素。

### 39.python中类方法，类实例方法，静态方法有何区别？

### 40.python中pass语句作用是什么？

### 41.介绍一下python中range()和xrange()函数的用法。

### 42.用python匹配 HTML Tag 的时候，<.*>和<.*?>有什么区别？

### 43.python中 如何拷贝一个对象？

### 44.如何用python查询和替换一个文本字符串？

### 45.Django里QuerySet的get和filter方法的区别？

### 46.简述Django对HTTP请求的执行流程。

### 47.简述Django下的(内建的)缓存机制。

### 48.Django中Model的slugFied类型字段有什么用途？

### 49.Django中如何加载初始数据？

### 50.python函数中经常有*args和**kwargs这两个参数，它们是什么意思，为什么使用它们？

### 51.python中变量的作用域，变量的查找顺序。

### 52.python中如何动态获取和设置对象的属性？

### 53.描述python中GIL的概念，以及它对python多线程的影响，编写一个多线程抓取网页的程序，并阐述  多线程抓取程序是否比单线程性能有提升，并解释原因。

### 54.mysql有哪些存储引擎，优化mysql数据库的方法有哪些。

### 55.Web开发中，session和cookie的作用与区别。

### 56.Web开发中有哪些技术手段防止SQL注入？

### 57.编写快速排序或者冒泡排序。

### 58.解释下HTTP常见的响应状态码。

### 59.Python是 如何进行内存管理的？

### 60.介绍一下python的异常处理机制和自己开发过程中的体会。

### 61.python中怎么有效读取一个20G大小的文件。

### 62.如何查看占用8080端口的是什么进程？

### 63.DNS解析过程是怎样的？有几种解析方式？各自的区别是什么？

### 64.TCP建立连接三次握手，断开连接四次挥手的过程是怎样的？

### 64.谈谈Django中的中间件。

### 65.谈谈CSRF原理

### 66.谈谈RESTful规范

### 67.谈谈Python中的面向对象

### 68.谈谈Django中CBV原理

### 68.谈谈Django REST freamwork

### 69.python基础数据类型
### 70.ambda表达式
### 71.map,filter,reduce是什么
### 72.写一个排序
### 73.贪婪匹配和非贪婪匹配
### 74.常用的编辑器以及快捷键
