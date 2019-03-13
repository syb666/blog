##  1 列举Http请求中常见的请求方式
##  2 谈谈你对HTTP协议的认识。1.1 长连接
##  3 简述MVC模式和MVT模式
##  4 简述Django请求生命周期
##  5 简述什么是FBV和CBV
##  6 谈一谈你对ORM的理解
##  7 rest_framework 认证组件的流程
##  8 什么是中间件并简述其作用
##  9 django 中间件生命周期
##  10 django中怎么写原生SQL
##  11 如何使用django orm批量创建数据
##  12 命令migrate 和makemigrations的差别
##  14 常用视图响应的方式是什么？
##  15 HTTP响应常见状态码分类
##  16 路由匹配原则是什么？
##  17 缓存系统类型有哪些
##  18 解决跨域的常用方式是什么？
##  19 信号的作用是什么？
##  20 Django的Model的继承有几种形式，分别是什么
##  21 Django中查询queryset时什么情况下用Q
##  22 Django中想验证表单提交是否格式正确需要用到Form中的哪个函数
##  23 orm如何取消级联
##  24 Django中如何读取和保存session，整个session的运行机制是什么
##  25 简述Django对http请求的执行流程
##  25 Django中当用户登录到A服务器进入登陆状态，下次被nginx代理到B服务器会出现什么影响
##  26 跨域请求Django是如何处理的
##  27 查询集的两大特性？什么是惰性执行
##  28 查询集返回的列表过滤器有哪些
##  29 如何获取django urlpatterns里面注册的所有url?
##  30 django路由系统中include是干嘛用的？
##  31 django2.0中的path与django1.xx里面的url有什么区别？
##  32 urlpatterns中的name与namespace有什么作用？你是如何使用的？
##  34 如何给一个字段设置一个主键？
##  35 如何设置一个带有枚举值的字典？
##  36 DateTimeField类型中的auto_now与auto_now_add有什么区别
##  37 values()与values_list()有什么区别？
##  38 selected_related与prefetch_related有什么区别？
##  39 当删除一个外键的时候，如何把与其关联的对应关系删除
##  40 class Meta中的元信息字段有哪些
##  41 多对多关联的表，如何插入数据？如何删除数据？如何更新数据？
##  42 django的M2M关系，如何手动生成第三张表？
##  43 在Django中，服务端给客户端响应信息有几种方式？分别是什么？
##  44 在视图函数中，常用的验证装饰器有哪些？
##  45 如何给一个视图函数加上缓存？
##  46 web框架的本质是什么？
##  47 创建Django工程、Django app、以及运行的命令
##  48 django中csrf的实现机制
##  49 Django App的目录结构
##  50 Django 获取用户前端请求数据的几种方式
##  51 描述下 自定义simple_tag
##  52 什么是Cookie、如何获取、设置Cookie
##  53 什么是session，与cookie的对比、设置、获取、清空session
##  54 什么是CSRF，及防范方式
##  55 get请求和post请求的区别
##  56 图书管理系统的表结构是怎么设计的？
##  57 WSGI / uwsgi/ uWSGI区分
##  59 解释blank和null
##  60 QueryDict和dict区别
##  61 谈谈你对restful规范的认识？
##  62 Django 本身提供了 runserver，为什么不能用来部署？  
##  63 Tornado 的核是什么？ 
##  64  Django重定向你是如何实现的？用的什么状态码？  
##  65 Django中如何加载初始化数据  
##  66 简述Django下的（内建）缓存机制


##  1 列举Http请求中常见的请求方式
HTTP请求的方法：
HTTP/1.1协议中共定义了八种方法（有时也叫“动作”），来表明Request-URL指定的资源不同的操作方式
 

* 注意：
  * 1）方法名称是区分大小写的，当某个请求所针对的资源不支持对应的请求方法的时候，服务器应当返回状态码405（Mothod Not Allowed）；当服务器不认识或者不支持对应的请求方法时，应返回状态码501（Not Implemented）。
  * 2）HTTP服务器至少应该实现GET和HEAD/POST方法，其他方法都是可选的，此外除上述方法，特定的HTTP服务器支持扩展自定义的方法。
                 
* GET
  * 向特定的路径资源发出请求
  * 注意：GET方法不应当被用于产生“副作用”的操作中，例如在WebApplication中，其中一个原因是GET可能会被网络蜘蛛等随意访问。Loadrunner中对应get请求函数：web_link和web_url
      
* POST
  * 向指定路径资源提交数据进行处理请求（一般用于提交表单或者上传文件）
  * 数据被包含在请求体中，POST请求可能会导致新的资源的建立和/或已有资源的修改。Loadrunner中对应POST请求函数：web_submit_data,web_submit_form
                   
* OPTIONS
  * 返回服务器针对特定资源所支持的HTTP请求方法
  * 允许客户端查看服务器的性能，也可以利用向web服务器发送‘*’的请求来测试服务器的功能性
     
* HEAD
  * 向服务器索与GET请求相一致的响应，只不过响应体将不会被返回
  * 这一方法可以再不必传输整个响应内容的情况下，就可以获取包含在响应小消息头中的元信息。
     
* PUT
  * 从客户端向服务器传送的数据取代指定的文档的内容
      
* DELETE
  * 请求服务器删除指定的页面
     
* TRACE
  * 回回显服务器收到的请求，主要用于测试或诊断
    
* CONNECT
  * HTTP/1.1协议中预留给能够将连接改为管道方式的代理服务


## 2 谈谈你对HTTP协议的认识。1.1 长连接
HTTP是一个属于应用层的面向对象的协议</br>
HTTP协议工作于客户端-服务端架构为上。浏览器作为HTTP客户端通过URL向HTTP服务端即WEB服务器发送所有请求。</br>
Web服务器根据接收到的请求后，向客户端发送响应信息。<br>
* 基于TCP/IP
双方建立通信的顺序,以及Web页面显示需要 处理的步骤,等等。像这样把与互联网相关联的协议集合起来总称为TCP/IP。
而http协议是基于TCP/IP协议之上的应用层协议。
* 基于请求－响应模式
HTTP协议规定,请求从客户端发出,最后服务器端响应该请求并 返回
* 无状态保存
HTTP是一种不保存状态,即无状态(stateless)协议。HTTP协议自身不对请求和响应之间的通信状态进行保存。
也就是说在HTTP这个级别,协议对于发送过的请求或响应都不做持久化处理。

 使用HTTP协议,每当有新的请求发送时,就会有对应的新响应产生。协议本身并不保留之前一切的请求或响应报文的信息。
 这是为了更快地处理大量事务,确保协议的可伸缩性,而特意把HTTP协议设计成 如此简单的。可是,随着Web的不断发展,
 因无状态而导致业务处理变得棘手 的情况增多了。比如,用户登录到一家购物网站,即使他跳转到该站的 其他页面后,
 也需要能继续保持登录状态。针对这个实例,网站为了能 够掌握是谁送出的请求,需要保存用户的状态。HTTP/1.1虽然是
 无状态协议,但为了实现期望的保持状态功能, 于是引入了Cookie技术。有了Cookie再用HTTP协议通信,就可以管 理状态了。

##  3 简述MVC模式和MVT模式
  所谓MVC就是把Web应用分为模型(M)，控制器(C)和视图(V)三层,他们之间以一种插件式的、松耦合的方式连接在一起，
模型负责业务对象与数据库的映射(ORM)，视图负责与用户的交互(页面)，控制器接受用户的输入调用模型和视图完成用户的请求

图一

  Django的MTV模式本质上和MVC是一样的，也是为了各组件间保持松耦合关系，只是定义上有些许不同，Django的MTV分别是值：
* M 代表模型（Model）： 负责业务对象和数据库的关系映射(ORM)。
* T 代表模板 (Template)：负责如何把页面展示给用户(html)。
* V 代表视图（View）： 负责业务逻辑，并在适当时候调用Model和Template。</br>
  除了以上三层之外，还需要一个URL分发器，它的作用是将一个个URL的页面请求分发给不同的View处理，
View再调用相应的Model和Template，MTV的响应模式如下所示：
  
图二

一般是用户通过浏览器向我们的服务器发起一个请求(request)，这个请求回去访问视图函数，
（如果不涉及到数据调用，那么这个时候视图函数返回一个模板也就是一个网页给用户），视图函数调用模型，模型去数据库查找数据，
然后逐级返回，视图函数把返回的数据填充到模板中空格中，最后返回网页给用户。


##  4 简述Django请求生命周期
一般是用户通过浏览器向我们的服务器发起一个请求(request)，这个请求回去访问视图函数，（如果不涉及到数据调用，
那么这个时候视图函数返回一个模板也就是一个网页给用户），视图函数调用模型，模型去数据库查找数据，然后逐级返回，
视图函数把返回的数据填充到模板中空格中，最后返回网页给用户。
* 1.wsgi,请求封装后交给web框架 （Flask、Django）     
* 2.中间件，对请求进行校验或在请求对象中添加其他相关数据，例如：csrf、request.session - 
* 3.路由匹配 根据浏览器发送的不同url去匹配不同的视图函数    
* 4.视图函数，在视图函数中进行业务逻辑的处理，可能涉及到：orm、templates => 渲染 - 
* 5.中间件，对响应的数据进行处理。 
* 6.wsgi,将响应的内容发送给浏览器。


图三

## 5 简述什么是FBV和CBV
* FBV（function base views）就是在视图里面使用函数处理请求
* CBV（class base views）就是在视图里面使用类处理请求

## 6 谈一谈你对ORM的理解
ORM是“对象-关系-映射”的简称。
MVC或者MVT框架中包括一个重要的部分，就是ORM，它实现了数据模型与数据库的解耦，即数据模型的设计不需要依赖于特定的数据库，
通过简单的配置就可以轻松更换数据库，这极大的减轻了开发人员的工作量，不需要面对因数据库变更而导致的无效劳动

## 7 rest_framework 认证组件的流程
认证组件 写一个类并注册到认证类(authentication_classes)，在类的的authticate方法中编写认证逻辑代码

## 8 什么是中间件并简述其作用
中间件是一个用来处理Django的请求和响应的框架级别的钩子。它是一个轻量、低级别的插件系统，用于在全局范围内改变Django的输入和输出。
每个中间件组件都负责做一些特定的功能。中间件是介于request与response处理之间的一道处理过程，相对比较轻量级，并且在全局上改变django的输入与输出。


## 9 django 中间件生命周期

图四

请求过来：
中间件：拦截一部分请求；比如验证session, 没有登录的 请求一些页面，跳转至登录页；(图片为中间件的请求过程.)再到 urls ，
分发请求到views 视图 ，通过 CBV(dispatch反射) 和 FBV 的 get 请求 讲 template 页面渲染返回给用户；渲染之前 可以从数据库拿出数据，
放到render 的参数里面传递过去， locals() 表示 把所有参数传递还可以 实例化 其他 form 类，并渲染给前端。

## 10 django中怎么写原生SQL

列举django orm中三种能写sql语句的方法
使用extra：查询人民邮电出版社出版并且价格大于50元的书籍
```python
Book.objects.filter(publisher__name='人民邮电出版社').extra(where=['price>50']) 
#使用raw
books=Book.objects.raw('select * from hello_book')  
for book in books:  
   print book  
#自定义sql
from django.db import connection  
  
cursor = connection.cursor()  
cursor.execute("insert into hello_author(name) VALUES ('郭敬明')")  
cursor.execute("update hello_author set name='韩寒' WHERE name='郭敬明'")  
cursor.execute("delete from hello_author where name='韩寒'")  
cursor.execute("select * from hello_author")  
cursor.fetchone()  
cursor.fetchall() 
```
## 11 如何使用django orm批量创建数据

使用django.db.models.query.QuerySet.bulk_create()批量创建对象，减少SQL查询次数。改进如下：
```python 
querysetlist=[]
for i in resultlist:
    querysetlist.append(Account(name=i))        
Account.objects.bulk_create(querysetlist)
```

## 12 命令migrate 和makemigrations的差别
1 .生成迁移文件
2 .执行迁移

## 13 视图函数中，接收的请求对象常用方法和属性有哪些
* path属性，获取请求页面的全路径，不包括域名
* method属性，获取请求中使用的HTTP方式的字符串表示。全大写表示
* GET属性，获取HTTP GET方式请求传参，的参数（字典类型）


如：http://127.0.0.1:8000/bug/articles/?mch=123 & mim=456
复制代码
```python
from django.shortcuts import render,HttpResponse
def special(request):
    print(request.GET)
    return render(request,'index.html') #向用户显示一个html页面
```
返回：<QueryDict: {' mim': ['456'], 'mch': ['123 ']}>
```python
# POST： 包含所有HTTP POST参数的类字典对象
#
# 服务器收到空的POST请求的情况也是可能发生的，也就是说，表单form通过
# HTTP POST方法提交请求，但是表单中可能没有数据，因此不能使用
# if req.POST来判断是否使用了HTTP POST 方法；应该使用 if req.method=="POST"
#COOKIES: 包含所有cookies的标准Python字典对象；keys和values都是字符串。
#
# FILES： 包含所有上传文件的类字典对象；FILES中的每一个Key都是<input type="file" name="" />标签中name属性的值，FILES中的每一个value同时也是一个标准的python字典对象，包含下面三个Keys：
#
# filename： 上传文件名，用字符串表示
# content_type: 上传文件的Content Type
# content： 上传文件的原始内容
#
#
# user： 是一个django.contrib.auth.models.User对象，代表当前登陆的用户。如果访问用户当前
# 没有登陆，user将被初始化为django.contrib.auth.models.AnonymousUser的实例。你
# 可以通过user的is_authenticated()方法来辨别用户是否登陆：
# if req.user.is_authenticated();只有激活Django中的AuthenticationMiddleware
# 时该属性才可用
#
# session： 唯一可读写的属性，代表当前会话的字典对象；自己有激活Django中的session支持时该属性才可用。
# get_full_path()方法，获取HTTP GET方式请求传参，的URL地址
```
## 14 常用视图响应的方式是什么？

视图的响应返回使用HttpResponse
```python
HttpResponse(content=响应体, content_type=响应体数据类型, status=状态码) 
一般不用这种方式，我习惯使用： 
 response = HttpResponse(‘language python’) 
 response.status_code = 400 
 response[‘languaget’] = ‘Python’ 
 return response
```
返回json数据，可以使用JsonResponse来构造响应对象
 
帮助我们将数据转换为json字符串 
设置响应头Content-Type为 application/json 
例子：
```pytyhon
from django.http import JsonResponse 
           def demo_view(request): 
               return JsonResponse({‘city’: ‘beijing’, ‘subject’: ‘python’})
```
redirect重定向 
我们也可以将返回的结果重定向到另一个页面或接口， 
```python
例子：from django.shortcuts import redirect 
           def demo_view(request): 
               return redirect(‘/index.html’)
```

## 15.HTTP响应常见状态码分类

图五

## 16 路由匹配原则是什么？
1、关于正则匹配优先级
* 在url匹配列表中，如果第一条和第二条同时满足匹配规则，则优先匹配第一条。
* 在url匹配列表中，如果第一条为正则模糊匹配，第二条为精确匹配，则优先匹配第一条

## 17 缓存系统类型有哪些

* 全站缓存
```python
MIDDLEWARE_CLASSES = (
    ‘django.middleware.cache.UpdateCacheMiddleware’, #第一
    'django.middleware.common.CommonMiddleware',
    ‘django.middleware.cache.FetchFromCacheMiddleware’, #最后
)
```
* 视图缓存
```python
from django.views.decorators.cache import cache_page
import time
  
@cache_page(15) #超时时间为15秒
def index(request):
 t=time.time() #获取当前时间
 return render(request,"index.html",locals())
```
* 模板缓存
```python
{% load cache %}
 <h3 style="color: green">不缓存:-----{{ t }}</h3>
  
{% cache 2 'name' %} # 存的key
 <h3>缓存:-----:{{ t }}</h3>
{% endcache %}
```

## 18 解决跨域的常用方式是什么？


跨域是指一个域下的文档或脚本试图去请求另一个域下的资源，这里跨域是广义的。
广义的跨域：
1.) 资源跳转： A链接、重定向、表单提交</br>
2.) 资源嵌入： link script img frame等dom标签，还有样式中background:url()、@font-face()等文件外链</br>
3.) 脚本请求： js发起的ajax请求、dom和js对象的跨域操作等</br>
 
### 1.2那么是什么同源策略呢？

同源策略/SOP（Same origin policy）是一种约定，由Netscape公司1995年引入浏览器，它是浏览器最核心也最基本的安全功能，如果缺少了同源策略，浏览器很容易受到XSS、CSFR等攻击。所谓同源是指"-- 协议+域名+端口 --"三者相同，即便两个不同的域名指向同一个ip地址，也非同源。</br>

### 同源策略限制以下几种行为：
* 1.) Cookie、LocalStorage 和 IndexDB 无法读取 
* 2.) DOM 和 Js对象无法获得
* 3.) AJAX 请求不能发送

图6

### 二、从安全性而言选择那种跨域方式最好，为什么？
一般最安全的是WINDOW.NAME，因为iframe会销毁

### 三、JSONP的缺点
jsonp有个缺陷就是只能get
而且会把请求的内容发送到url中导致安全性极低

跨域：
浏览器从一个域名的网页去请求另一个域名的资源时,浏览器处于安全的考虑,不允许不同源的请求</br>
处理方法：
通过JSONP跨域
1. JSON是一种数据交换格式
2. JSONP是一种非官方的跨域数据交互协议
3. jsonp是包含在函数调用中的json
4. script标签不受同源策略的影响，手动创建一个script标签,传递URL,同时传入一个回调函数的名字
服务器得到名字后,返回数据时会用这个函数名来包裹住数据,客户端获取到数据之后，立即把script标签删掉
cors：跨域资源共享
使用自定义的HTTP头部允许浏览器和服务器相互通信
1.如果是简单请求,直接设置允许访问的域名：
允许你的域名来获取我的数据                         
response['Access-Control-Allow-Origin'] = "*"

## 19 信号的作用是什么？

图7

## 20 Django的Model的继承有几种形式，分别是什么

1.通常，你只是想用父 model 来保存那些你不想在子 model 中重复录入的信息。父类是不使用的也就是不生成单独的数据表，这种情况下使用抽象基类继承 Abstract base classes。

2.如果你想从现有的Model继承并让每个Model都有自己的数据表，那么使用多重表继承Multi-table inheritance。
 
3.最后，如果你只想在 model 中修改 Python-level 级的行为，而不涉及字段改变。 代理 model (Proxy models) 适用于这种场合

## 21 Django中查询queryset时什么情况下用Q

* F:对数据本身的不同字段进行操作 如:比较和更新，对数据进行加减操作
* Q：用于构造复杂的查询条件 如：& |操作

## 22 Django中想验证表单提交是否格式正确需要用到Form中的哪个函数

form.is_valid() :返回布尔值

## 23 orm如何取消级联
Django取消级联删除

图8

并且SET_NULL只有在null为True的时候，才可以使用。

## 24 Django中如何读取和保存session，整个session的运行机制是什么

```python
 更新
     在django—session表中创建一条记录:
       session-key                                     session-data
       ltv8zy1kh5lxj1if1fcs2pqwodumr45t                  更新数据else:
    1 生成随机字符串   ltv8zy1kh5lxj1if1fcs2pqwodumr45t
    2 response.set_cookie("sessionid",ltv8zy1kh5lxj1if1fcs2pqwodumr45t)
    3 在django—session表中创建一条记录:
       session-key                                     session-data
       ltv8zy1kh5lxj1if1fcs2pqwodumr45t       {"is_login":True,"username":"yuan"}

```

## 25 简述Django对http请求的执行流程

  在接受一个Http请求之前的准备,启动一个支持WSGI网关协议的服务器监听端口等待外界的Http请求，比如Django自带的开发者服务器或者uWSGI服务器。
  
  服务器根据WSGI协议指定相应的Handler来处理Http请求，并且初始化该Handler，在Django框架中由框架自身负责实现这一个Handler。 此时服务器已处于监听状态，可以接受外界的Http请求 当一个http请求到达服务器的时候   
  
  服务器根据WSGI协议从Http请求中提取出必要的参数组成一个字典（environ）并传入Handler中进行处理。  
  
在Handler中对已经符合WSGI协议标准规定的http请求进行分析，比如加载Django提供的中间件，路由分配，调用路由匹配的视图等。 返回一个可以被浏览器解析的符合Http协议的HttpResponse。

## 25 Django中当用户登录到A服务器进入登陆状态，下次被nginx代理到B服务器会出现什么影响

如果用户在A应用服务器登陆的session数据没有共享到B应用服务器，那么之前的登录状态就没有了。

## 26 跨域请求Django是如何处理的

* 启用中间件
* post请求
* 验证码
* 表单中添加{%csrf_token%}标签

## 27 查询集的两大特性？什么是惰性执行

* 惰性执行、缓存 。  
* 创建查询集不会访问数据库，直到调用数据时，才会访问数据库，调用数据的情况包括迭代、序列化、与if合用

## 28 查询集返回的列表过滤器有哪些
* all()：返回所有数据  
* filter()：返回满足条件的数据 
* exclude()：返回满足条件之外的数据，相当于sql语句中where部分的not关键字  
* order_by()：排序

## 30 django路由系统中include是干嘛用的？
 
include路由转发
通常，我们会在每个app里，各自创建一个urls.py路由模块，然后从根路由出发，将app所属的url请求，全部转发到相应的urls.py模块中。

## 31 django2.0中的path与django1.xx里面的url有什么区别？

2.0内的path匹配正则时候无效，导入re_path即可匹配正则         url = re_path
url()        是django.urls.re_path()别名       
url(regex, view, kwargs=None, name=None)[source]
This function is an alias to django.urls.re_path(). It's likely to be deprecated in a future release.

## 32 urlpatterns中的name与namespace有什么作用？你是如何使用的？

name:别名，给路由起一个别名
namespace:名称空间，防止多个应用之间的路由重复

## 33 如何根据urlpatterns中的name反向生成url,这样反向生成url的方式有几种？
使用HttpResponseRedirect redirect和reverse 状态码：302,301

## 35 如何设置一个带有枚举值的字典？

图9

## 36 DateTimeField类型中的auto_now与auto_now_add有什么区别

DateTimeField.auto_now
这个参数的默认值为false，设置为true时，能够在保存该字段时，将其值设置为当前时间，并且每次修改model，都会自动更新。因此这个参数在需要存储“最后修改时间”的场景下，十分方便。需要注意的是，设置该参数为true时，并不简单地意味着字段的默认值为当前时间，而是指字段会被“强制”更新到当前时间，你无法程序中手动为字段赋值；如果使用django再带的admin管理器，那么该字段在admin中是只读的。

DateTimeField.auto_now_add
这个参数的默认值也为False，设置为True时，会在model对象第一次被创建时，将字段的值设置为创建时的时间，以后修改对象时，字段的值不会再更新。该属性通常被用在存储“创建时间”的场景下。与auto_now类似，auto_now_add也具有强制性，一旦被设置为True，就无法在程序中手动为字段赋值，在admin中字段也会成为只读的。

## 37 values()与values_list()有什么区别？

* values : 取字典的queryset
* values_list : 取元组的queryset

## 38 selected_related与prefetch_related有什么区别？

  在Django中，所有的Queryset都是惰性的，意思是当创建一个查询集的时候，并没有跟数据库发生任何交互。因此我们可以对查询集进行级联的filter等操作，只有在访问Queryset的内容的时候，Django才会真正进行数据库的访问。而多频率、复杂的数据库查询往往是性能问题最大的根源。
 
  不过我们实际开发中，往往需要访问到外键对象的其他属性。如果按照默认的查询方式去遍历取值，那么会造成多次的数据库查询，效率可想而知。

  在查询对象集合的时候，把指定的外键对象也一并完整查询加载，避免后续的重复查询。
 
1，select_related适用于外键和多对一的关系查询；
2，prefetch_related适用于一对多或者多对多的查询。
## 39 当删除一个外键的时候，如何把与其关联的对应关系删除
1. 删除关联表中的数据时,当前表与其关联的field的操作
2. django2.0之后，表与表之间关联的时候,必须要写on_delete参数,否则会报异常

## 40 class Meta中的元信息字段有哪些

通过一个内嵌类 "class Meta" 给你的 model 定义元数据, 类似下面这样:
```python
class Foo(models.Model): 
    bar = models.CharField(maxlength=30)
 
    class Meta: 
        # ...
```
Model 元数据就是 "不是一个字段的任何数据" -- 比如排序选项, admin 选项等等.</br>
 
下面是所有可能用到的 Meta 选项. 没有一个选项是必需的. 是否添加 class Meta 到你的 model 完全是可选的.</br>
 
app_label</br>
app_label这个选项只在一种情况下使用，就是你的模型类不在默认的应用程序包下的models.py文件中，这时候你需要指定你这个模型类是那个应用程序的。比如你在其他地方写了一个模型类，而这个模型类是属于myapp的，那么你这是需要指定为：</br>
 
app_label='myapp'</br>
db_table</br>
db_table是用于指定自定义数据库表名的。Django有一套默认的按照一定规则生成数据模型对应的数据库表名，如果你想使用自定义的表名，就通过这个属性指定，比如：</br>
table_name='my_owner_table'   </br>
若不提供该参数, Django 会使用 app_label + '_' + module_name 作为表的名字.</br>
若你的表的名字是一个 SQL 保留字, 或包含 Python 变量名不允许的字符--特别是连字符 --没关系. Django 会自动在幕后替你将列名字和表名字用引号引起来.
 
db_tablespace</br>
有些数据库有数据库表空间，比如Oracle。你可以通过db_tablespace来指定这个模型对应的数据库表放在哪个数据库表空间。</br>
 
get_latest_by</br>
由于Django的管理方法中有个lastest()方法，就是得到最近一行记录。如果你的数据模型中有 DateField 或 DateTimeField 类型的字段，你可以通过这个选项来指定lastest()是按照哪个字段进行选取的。</br>
一个 DateField 或 DateTimeField 字段的名字. 若提供该选项, 该模块将拥有一个 get_latest() 函数以得到 "最新的" 对象(依据那个字段):</br>
get_latest_by = "order_date"</br>
 
managed</br>
由于Django会自动根据模型类生成映射的数据库表，如果你不希望Django这么做，可以把managed的值设置为False。</br>
默认值为True,这个选项为True时Django可以对数据库表进行 migrate或migrations、删除等操作。在这个时间Django将管理数据库中表的生命周期</br>
如果为False的时候，不会对数据库表进行创建、删除等操作。可以用于现有表、数据库视图等，其他操作是一样的。</br>
 
order_with_respect_to</br>
这个选项一般用于多对多的关系中，它指向一个关联对象。就是说关联对象找到这个对象后它是经过排序的。指定这个属性后你会得到一个get_XXX_order()和set_XXX_order（）的方法,通过它们你可以设置或者回去排序的对象。</br>
 
举例来说, 如果一个 PizzaToppping 关联到一个 Pizza 对象, 这样做:</br>
 
order_with_respect_to = 'pizza'</br>
...就允许 toppings 依照相关的 pizza 来排序.</br>
 
ordering</br>
这个字段是告诉Django模型对象返回的记录结果集是按照哪个字段排序的。比如下面的代码：</br>
 
ordering=['order_date'] </br>
按订单升序排列</br>
ordering=['-order_date'] </br>
按订单降序排列，-表示降序</br>
ordering=['?order_date'] </br>
随机排序，？表示随机</br>
ordering = ['-pub_date', 'author']</br>
对 pub_date 降序,然后对 author 升序</br>
需要注意的是:不论你使用了多少个字段排序, admin 只使用第一个字段</br>
 
permissions</br>
permissions主要是为了在Django Admin管理模块下使用的，如果你设置了这个属性可以让指定的方法权限描述更清晰可读。</br>
 
要创建一个对象所需要的额外的权限. 如果一个对象有 admin 设置, 则每个对象的添加,删除和改变权限会人(依据该选项)自动创建.下面这个例子指定了一个附加权限</br> 
can_deliver_pizzas:</br>
 
permissions = (("can_deliver_pizzas", "Can deliver pizzas"),)</br>
这是一个2-元素 tuple 的tuple或列表, 其中两2-元素 tuple 的格式为:(permission_code, human_readable_permission_name).</br>
 
unique_together</br>
unique_together这个选项用于：当你需要通过两个字段保持唯一性时使用。这会在 Django admin 层和数据库层同时做出限制(也就是相关的 UNIQUE 语句会被包括在 CREATE TABLE 语句中)。</br>
比如：一个Person的FirstName和LastName两者的组合必须是唯一的，那么需要这样设置：</br>
 
unique_together = (("first_name", "last_name"),)</br>
verbose_name</br>
verbose_name的意思很简单，就是给你的模型类起一个更可读的名字：</br>
 
verbose_name = "pizza"</br>
若未提供该选项, Django 则会用一个类名字的 munged 版本来代替: CamelCase becomes camel case.</br>
 
verbose_name_plural</br>
这个选项是指定，模型的复数形式是什么，比如：</br>
 
verbose_name_plural = "stories"</br>
若未提供该选项, Django 会使用 verbose_name + "s".</br>

## 41 多对多关联的表，如何插入数据？如何删除数据？如何更新数据？


## 42 django的M2M关系，如何手动生成第三张表？
```python
tags = models.ManyToManyField(
        to="Tag",
        through='Article2Tag',
        through_fields=('article', 'tag'),
    )
```

## 43 在Django中，服务端给客户端响应信息有几种方式？分别是什么？
* HTTPresponse，
* jsonresponse,
* redirect

## 44 在视图函数中，常用的验证装饰器有哪些？

## 45 如何给一个视图函数加上缓存？

## 46 web框架的本质是什么？
本质上其实就是一个socket服务端，用户的浏览器其实就是一个socket客户端。

## 47 创建Django工程、Django app、以及运行的命令


## 48 django中csrf的实现机制

* 第一步：django第一次响应来自某个客户端的请求时,后端随机产生一个token值，把这个token保存在SESSION状态中;同时,后端把这个token放到cookie中交给前端页面；
* 第二步：下次前端需要发起请求（比如发帖）的时候把这个token值加入到请求数据或者头信息中,一起传给后端；Cookies:{csrftoken:xxxxx}
* 第三步：后端校验前端请求带过来的token和SESSION里的token是否一致；

## 49 Django App的目录结构

图10

## 50 Django 获取用户前端请求数据的几种方式

@get和@post使用
1：在views模板下编写测试函数(记得在urls.py文件中进行相应配置) 
2：将刚刚封装的函数所在模板引入views.py 
3：使用@get进行拦截
 
@params，response_success，response_failure使用
第一种
```python
@login_required
def simple_view(request):
       return HttpResponse()
``` 
2 通过对基于函数视图或者基于类视图使用一个装饰器实现控制：  
```@login_required(MyView.as_view())```
3 通过覆盖mixin的类视图的dispatch方法实现控制：

## 51 描述下 自定义simple_tag

* 自定义filter：{{ 参数1|filter函数名:参数2 }}
  * 1.可以与if标签来连用
  * 2.自定义时需要写两个形参
* simple_tag:{% simple_tag函数名 参数1 参数2 %}
  * 1.可以传多个参数,没有限制
  * 2.不能与if标签来连用
```python
@register.simple_tag
def multi_tag(x,y):
    return x*y
```

## 52 什么是Cookie、如何获取、设置Cookie

会话跟踪技术，保留用户
Cookie是由服务器创建，然?后通过响应发送给客户端?的一个键值对。
具体一个浏览器针对一个服务器存储的key-value({ })
```python
response.set_cookie("is_login",True) 
request.COOKIES.get("is_login")
```

## 53 什么是session，与cookie的对比、设置、获取、清空session

   Session是服务器端技术，利用这个技术，服务器在运行时可以 为每一个用户的浏览器创建一个其独享的session对象，由于 session为用户浏览器独享，所以用户在访问服务器的web资源时 ，可以把各自的数据放在各自的session中，当用户再去访问该服务器中的其它web资源时，其它web资源再从用户各自的session中 取出数据为用户服务。
   

启用session
编辑settings.py中的一些配置
MIDDLEWARE_CLASSES 确保其中包含以下内容

`'django.contrib.sessions.middleware.SessionMiddleware',`

INSTALLED_APPS 是包含
`'django.contrib.sessions',`

```python
# 创建或修改 session：
request.session[key] = value
# 获取 session：
request.session.get(key,default=None)
# 删除 session
del request.session[key] # 不存在时报错
```

## 54 什么是CSRF，及防范方式

1. 启用中间件
2. post请求
3. 验证码
4. 表单中添加{%csrf_token%}标签

## 55 get请求和post请求的区别

请求方式: get与post请求
* GET提交的数据会放在URL之后，以?分割URL和传输数据，参数之间以&相连，如EditBook?name=test1&id=123456. POST方法是把提交的数据放在HTTP包的Body中.
* GET提交的数据大小有限制（因为浏览器对URL的长度有限制），而POST方法提交的数据没有限制.
* GET与POST请求在服务端获取请求数据方式不同。
* GET方式提交数据，会带来安全问题，比如一个登录页面，通过GET方式提交数据时，用户名和密码将出现在URL上，如果页面可以被缓存或者其他人可以访问这台机器，就可以从历史记录获得该用户的账号和密码.


## 57 WSGI / uwsgi/ uWSGI区分
1. WSGI
WSGI的全称是Web Server Gateway Interface（Web服务器网关接口），它不是服务器、python模块、框架、API或者任何软件，只是一种描述web服务器（如nginx，uWSGI等服务器）如何与web应用程序（如用Django、Flask框架写的程序）通信的规范。
server和application的规范在PEP3333中有具体描述，要实现WSGI协议，必须同时实现web server和web application，当前运行在WSGI协议之上的web框架有Bottle, Flask, Django。
2. uWSGI
uWSGI是一个全功能的HTTP服务器，实现了WSGI协议、uwsgi协议、http协议等。它要做的就是把HTTP协议转化成语言支持的网络协议。比如把HTTP协议转化成WSGI协议，让Python可以直接使用。
3. uwsgi
 与WSGI一样，是uWSGI服务器的独占通信协议，用于定义传输信息的类型(type of information)。每一个uwsgi packet前4byte为传输信息类型的描述，与WSGI协议是两种东西，据说该协议是fcgi【FCGI：fast common gateway interface 快速通用网关接口协议的10倍快。

## 58 如何使用django加密

Django 内置的User类提供了用户密码的存储、验证、修改等功能，默认使用pbkdf2_sha256方式来存储和管理用的密码。
django通过setting.py文件中的PASSWORD_HASHERS来设置选择要使用的算法，列表的第一个元素 (即settings.PASSWORD_HASHERS[0]) 会用于储存密码， 所有其它元素都是用于验证的哈希值，它们可以用于检查现有的密码。意思是如果你打算使用不同的算法，你需要修改PASSWORD_HASHERS，来将你最喜欢的算法在列表中放在首位。
 
一个settings中的Password_hashers看起来是这样的：
 
```python 
PASSWORD_HASHERS = (
 
    'django.contrib.auth.hashers.PBKDF2PasswordHasher',
    'django.contrib.auth.hashers.PBKDF2SHA1PasswordHasher',
    'django.contrib.auth.hashers.BCryptSHA256PasswordHasher',
    'django.contrib.auth.hashers.BCryptPasswordHasher',
    'django.contrib.auth.hashers.SHA1PasswordHasher',
    'django.contrib.auth.hashers.MD5PasswordHasher',
    'django.contrib.auth.hashers.CryptPasswordHasher',
)
```
具体的密码生成以及验证实现
```python
from django.contrib.auth.hashers import make_password,check_password
pwd='4562154'
mpwd=make_password(pwd,None,'pbkdf2_sha256') # 创建django密码，第三个参数为加密算法
pwd_bool=check_password(pwd,mpwd) # 返回的是一个bool类型的值，验证密码正确与否
``` 
 
Django之密码加密
通过django自带的类库，来加密解密很方便，下面来简单介绍下；
 
导入包：
 
`from django.contrib.auth.hashers import make_password, check_password`
从名字就可以看出来他们的作用了。
 
一个是生成密码，一个是核对密码。
 
例如：
 
`make_password("123456")`
得到结果：
 
`u'pbkdf2_sha25615000MAjic3nDGFoi$qbclz+peplspCbRF6uoPZZ42aJIIkMpGt6lQ+Iq8nfQ='`
另外也可以通过参数来生成密码：
 
`>>> make_password("123456", None, 'pbkdf2_sha256')`
校验:
 
校验就是通过check_password(原始值, 生成的密文)来校验密码的。
 
`>>> check_password("123456","pbkdf2_sha25615000MAjic3nDGFoi$qbclz+peplspCbRF6uoPZZ42aJIIkMpGt6lQ+Iq8nfQ=")`
`True`


## 59 解释blank和null

1. blank
设置为True时，字段可以为空。设置为False时，字段是必须填写的。字符型字段CharField和TextField是用空字符串来存储空值的。如果为True，字段允许为空，默认不允许。
 
2. null
设置为True时，django用Null来存储空值。日期型、时间型和数字型字段不接受空字符串。所以设置IntegerField，DateTimeField型字段可以为空时，需要将blank，null均设为True。

如果为True，空值将会被存储为NULL，默认为False。

如果想设置BooleanField为空时可以选用NullBooleanField型字段。

一句话概括
null 是针对数据库而言，如果 null=True, 表示数据库的该字段可以为空。NULL represents non-existent data.
blank 是针对表单的，如果 blank=True，表示你的表单填写该字段的时候可以不填。比如 admin 界面下增加 model 一条记录的时候。直观的看到就是该字段不是粗体

## 60 QueryDict和dict区别

* 在HttpRequest对象中, GET和POST属性是django.http.QueryDict类的实例。
* QueryDict类似字典的自定义类，用来处理单键对应多值的情况。在 HttpRequest 对象中,属性 GET 和 POST 得到的都是 django.http.QueryDict 所创建的实例。这是一个
* django 自定义的类似字典的类，用来处理同一个键带多个值的情况。
* 在 python 原始的字典中，当一个键出现多个值的时候会发生冲突，只保留最后一个值。而在 HTML 表单中，通常会发生一个键有多个值的情况，例如 <selectmultiple> （多选框）就是一个很常见情况。
* request.POST 和request.GET 的QueryDict 在一个正常的请求/响应循环中是不可变的。若要获得可变的版本，需要使用.copy()方法。
 
django QuerySet对象转换成字典对象

```python
manage.py shell
from django.contrib.auth.models import User
from django.forms.models import model_to_dict
u = User.objects.get(id=1)
u_dict = model_to_dict(u)
```
`type(u)</br>`
`<class 'django.contrib.auth.models.User'>`
`type(u_dict)</br>`
`<type 'dict'></br>`

`QueryDict.__init__(query_string=None, mutable=False, encoding=None)`
这是一个构造函数，其中 query_string 需要一个字符串，例如：</br> 
`QueryDict('a=1&a=2&c=3')</br>`
`<QueryDict: {'a': ['1', '2'], 'c': ['3']}></br>`

## 61 谈谈你对restful规范的认识？

首先restful是一种软件架构风格或者说是一种设计风格，并不是标准，它只是提供了一组设计#原则和约束条件，主要用于客户端和服务器交互类的软件。</br>     
就像设计模式一样，并不是一定要遵循这些原则，而是基于这个风格设计的软件可以更简洁，更#有层次，我们可以根据开发的实际情况，做相应的改变。</br>
它里面提到了一些规范，例如：</br>

* 1. restful 提倡面向资源编程,在url接口中尽量要使用名词，不要使用动词</br>   

* 2. 在url接口中推荐使用Https协议，让网络接口更加安全</br>
   `https://www.bootcss.com/v1/mycss？page=3`</br>
    （Https是Http的安全版，即HTTP下加入SSL层，HTTPS的安全基础是SSL，</br>
   因此加密的详细内容就需要SSL（安全套接层协议））</br>    

* 3. 在url中可以体现版本号</br>
   https://v1.bootcss.com/mycss</br>
   不同的版本可以有不同的接口，使其更加简洁，清晰</br>
   
* 4. url中可以体现是否是API接口 </br>
   https://www.bootcss.com/api/mycss  </br>    
   
* 5. url中可以添加条件去筛选匹配</br>
   https://www.bootcss.com/v1/mycss？page=3      </br>      
   
* 6. 、可以根据Http不同的method，进行不同的资源操作</br>
   （5种方法：GET / POST / PUT / DELETE / PATCH）</br>      
   
* 7. 响应式应该设置状态码</br>

* 8. 有返回值，而且格式为统一的json格式 </br>         

* 9. 返回错误信息</br>
   返回值携带错误信息  </br>        
   
* 10. 返回结果中要提供帮助链接，即API最好做到Hypermedia</br>
   如果遇到需要跳转的情况 携带调转接口的URL</br>
   
```python
ret = {
     code: 1000,
     data:{
     id:1,
     name:'小强',
     depart_id:http://www.luffycity.com/api/v1/depart/8/
     }
}
```

## 62 Django 本身提供了 runserver，为什么不能用来部署？

   runserver 方法是调试 Django 时经常用到的运行方式，它使用 Django 自带的  WSGI Server 运行，主要在测试和开发中使用，并且 runserver 开启的方式也是单进程 。  
   uWSGI 是一个 Web 服务器，它实现了 WSGI 协议、uwsgi、http 等协议。注意 uwsgi 是一种通信协议，而 uWSGI 是实现 uwsgi 协议和 WSGI 协议的 Web 服务器。
   uWSGI 具有超快的性能、低内存占用和多 app 管理等优点，并且搭配着 Nginx  就是一个生产环境了，能够将用户访问请求与应用 app 隔离开，实现真正的部署 。相比来讲，支持的并发量更高，方便管理多进程，发挥多核的优势，提升性能。
   
## 63 Tornado 的核是什么？   

Tornado 的核心是 ioloop 和 iostream 这两个模块，前者提供了一个高效的 I/O 事件循环，后者则封装了 一个无阻塞的 socket 。通过向 ioloop 中添加网络 I/O 事件，利用无阻塞的 socket ，再搭配相应的回调 函数，便可达到梦寐以求的高效异步执行。

## 64  Django重定向你是如何实现的？用的什么状态码？
1. 使用HttpResponseRedirect 
2. redirect 和reverse 
3. 状态码：302,301

## 65 Django中如何加载初始化数据
Django在创建对象时在掉用save()方法后，ORM框架会把对象的属性转换为写入到数据库中，实现对数据库的初始化；通过操作对象，查询数据库，将查询集返回给视图函数，通过模板语言展现在前端页面


## 66 简述Django下的（内建）缓存机制  

 Django根据设置的缓存方式，浏览器第一次请求时，cache会缓存单个变量或整个网页等内容到硬盘或者内存中，同时设置response头部，当浏览器再次发起请求时，附带f-Modified-Since请求时间到Django，Django 发现f-Modified-Since会先去参数之后，会与缓存中的过期时间相比较，如果缓存时间比较新，则会重新请求数据，并缓存起来然后返回response给客户端，如果缓存没有过期，则直接从缓存中提取数据，返回给response给客户端。


