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
