# 第一个Django项目
## 创建 Django 项目：
用命令行的方式：
1. 创建项目：打开终端，使用命令： django-admin startproject [项目名称] 即可创建。比
如： django-admin startproject first_project 。
2. 创建应用（app）：一个项目类似于是一个架子，但是真正起作用的还是 app 。在终端进入到
项目所在的路径，然后执行 python manage.py startapp [app名称] 创建一个app。
用 pycharm 的方式：
用 pycharm 新建一个 Django 项目，新建项目的截图如下：
使用 pycharm 创建完项目后，还是需要重新进入到命令行单独创建 app 的。
## 运行Django项目：
1. 通过命令行的方式： python manage.py runserver 。这样可以在本地访问你的网站，默认端
口号是 8000 ，这样就可以在浏览器中通过 http://127.0.0.1:8000/ 来访问你的网站啦。如
果想要修改端口号，那么在运行的时候可以指定端口号， python manage.py runserver
9000 这样就可以通过 9000 端口来访问啦。另外，这样运行的项目只能在本机上能访问，如
果想要在其他电脑上也能访问本网站，那么需要指定 ip 地址为 0.0.0.0 。示例为： python
11
manage.py runserver 0.0.0.0:8000 。
2. 通过 pycharm 运行。直接点击右上角的绿色箭头按钮即可运行。
## 项目结构介绍：
1. manage.py ：以后和项目交互基本上都是基于这个文件。一般都是在终端输入 python
manage.py [子命令] 。可以输入 python manage.py help 看下能做什么事情。除非你知道你自
己在做什么，一般情况下不应该编辑这个文件。
2. settings.py ：本项目的设置项，以后所有和项目相关的配置都是放在这个里面。
3. urls.py ：这个文件是用来配置URL路由的。比如访问 http://127.0.0.1/news/ 是访问新闻
列表页，这些东西就需要在这个文件中完成。
4. wsgi.py ：项目与 WSGI 协议兼容的 web 服务器入口，部署的时候需要用到的，一般情况下
也是不需要修改的。
## project和app的关系：
app 是 django 项目的组成部分。一个 app 代表项目中的一个模块，所有 URL 请求的响应都是
由 app 来处理。比如豆瓣，里面有图书，电影，音乐，同城等许许多多的模块，如果站
在 django 的角度来看，图书，电影这些模块就是 app ，图书，电影这些 app 共同组成豆瓣这个
项目。因此这里要有一个概念， django 项目由许多 app 组成，一个 app 可以被用到其他项
目， django 也能拥有不同的 app 。
12
# URL分发器
## 视图：
视图一般都写在 app 的 views.py 中。并且视图的第一个参数永远都是 request （一个
HttpRequest）对象。这个对象存储了请求过来的所有信息，包括携带的参数以及一些头部信息
等。在视图中，一般是完成逻辑相关的操作。比如这个请求是添加一篇博客，那么可以通过
request来接收到这些数据，然后存储到数据库中，最后再把执行的结果返回给浏览器。视图函数
的返回结果必须是 HttpResponseBase 对象或者子类的对象。示例代码如下：
` ``python
from django.http import HttpResponse
` ``
` ``python
def book_list(request):
` ``
` ``python
return HttpResponse("书籍列表！")</br>
` ``
## URL映射：
视图写完后，要与URL进行映射，也即用户在浏览器中输入什么 url 的时候可以请求到这个视图
函数。在用户输入了某个 url ，请求到我们的网站的时候， django 会从项目的 urls.py 文件中
寻找对应的视图。在 urls.py 文件中有一个 urlpatterns 变量，以后 django 就会从这个变量中
读取所有的匹配规则。匹配规则需要使用 django.urls.path 函数进行包裹，这个函数会根据传入
的参数返回 URLPattern 或者是 URLResolver 的对象。示例代码如下：
` ``python
from django.contrib import admin
from django.urls import path
from book import views
urlpatterns = [
path('admin/', admin.site.urls),
path('book/',views.book_list)
]
` ``
## URL中添加参数：
有时候， url 中包含了一些参数需要动态调整。比如简书某篇文章的详情页的url，
是 https://www.jianshu.com/p/a5aab9c4978e 后面的 a5aab9c4978e 就是这篇文章的 id ，那么简
书的文章详情页面的url就可以写成 https://www.jianshu.com/p/<id> ，其中id就是文章的id。那么
如何在 django 中实现这种需求呢。这时候我们可以在 path 函数中，使用尖括号的形式来定义一
个参数。比如我现在想要获取一本书籍的详细信息，那么应该在 url 中指定这个参数。示例代码
如下：
'''from django.contrib import admin
from django.urls import path
from book import views
urlpatterns = [
path('admin/', admin.site.urls),
path('book/',views.book_list),
path('book/<book_id>/',views.book_detail)
]'''
而 views.py 中的代码如下：
def book_detail(request,book_id):
text = "您输入的书籍的id是：%s" % book_id
return HttpResponse(text)
当然，也可以通过查询字符串的方式传递一个参数过去。示例代码如下：
'''urlpatterns = [
path('admin/', admin.site.urls),
path('book/',views.book_list),
path('book/detail/',views.book_detail)
]'''
在 views.py 中的代码如下：
def book_detail(request):
book_id = request.GET.get("id")
text = "您输入的书籍id是：%s" % book_id
return HttpResponse(text)
以后在访问的时候就是通过 /book/detail/?id=1 即可将参数传递过去。
## URL中包含另外一个urls模块：
在我们的项目中，不可能只有一个 app ，如果把所有的 app 的 views 中的视图都放
在 urls.py 中进行映射，肯定会让代码显得非常乱。因此 django 给我们提供了一个方法，可以
在 app 内部包含自己的 url 匹配规则，而在项目的 urls.py 中再统一包含这个 app 的 urls 。
使用这个技术需要借助 include 函数。示例代码如下：
first_project/urls.py文件：
from django.contrib import admin
from django.urls import path,include
14
urlpatterns = [
path('admin/', admin.site.urls),
path('book/',include("book.urls"))
]
在 urls.py 文件中把所有的和 book 这个 app 相关的 url 都移动到 app/urls.py 中了，然后
在 first_project/urls.py 中，通过 include 函数包含 book.urls ，以后在请求 book 相关的url
的时候都需要加一个 book 的前缀。
# book/urls.py文件：
from django.urls import path
from . import views
urlpatterns = [
path('list/',views.book_list),
path('detail/<book_id>/',views.book_detail)
]
以后访问书的列表的 url 的时候，就通过 /book/list/ 来访问，访问书籍详情页面的 url 的时候
就通过 book/detail/<id> 来访问。
## path函数：
path 函数的定义为： path(route,view,name=None,kwargs=None) 。以下对这几个参数进行讲解。
1. route 参数： url 的匹配规则。这个参数中可以指定 url 中需要传递的参数，比如在访问文
章详情页的时候，可以传递一个 id 。传递参数是通过 <> 尖括号来进行指定的。并且在传递
参数的时候，可以指定这个参数的数据类型，比如文章的 id 都是 int 类型，那么可以这样
写 <int:id> ，以后匹配的时候，就只会匹配到 id 为 int 类型的 url ，而不会匹配其他
的 url ，并且在视图函数中获取这个参数的时候，就已经被转换成一个 int 类型了。其中还
有几种常用的类型：
str：非空的字符串类型。默认的转换器。但是不能包含斜杠。
int：匹配任意的零或者正数的整形。到视图函数中就是一个int类型。
slug：由英文中的横杠 - ，或者下划线 _ 连接英文字符或者数字而成的字符串。
uuid：匹配 uuid 字符串。
path：匹配非空的英文字符串，可以包含斜杠。
2. view 参数：可以为一个视图函数或者是 类视图.as_view() 或者是 django.urls.include() 函
数的返回值。
3. name 参数：这个参数是给这个 url 取个名字的，这在项目比较大， url 比较多的时候用处
很大。
15
4. kwargs 参数：有时候想给视图函数传递一些额外的参数，就可以通过 kwargs 参数进行传
递。这个参数接收一个字典。传到视图函数中的时候，会作为一个关键字参数传过去。比如以
下的 url 规则：
from django.urls import path
from . import views
urlpatterns = [
path('blog/<int:year>/', views.year_archive, {'foo': 'bar'}),
]
那么以后在访问 blog/1991/ 这个url的时候，会将 foo=bar 作为关键字参数传
给 year_archive 函数。
## re_path函数：
有时候我们在写url匹配的时候，想要写使用正则表达式来实现一些复杂的需求，那么这时候我们
可以使用 re_path 来实现。 re_path 的参数和 path 参数一模一样，只不过第一个参数也就
是 route 参数可以为一个正则表达式。
一些使用 re_path 的示例代码如下：
from django.urls import path, re_path
from . import views
urlpatterns = [
path('articles/2003/', views.special_case_2003),
re_path(r'articles/(?P<year>[0-9]{4})/', views.year_archive),
re_path(r'articles/(?P<year>[0-9]{4})/(?P<month>[0-9]{2})/', views.month_archiv
e),
re_path(r'articles/(?P<year>[0-9]{4})/(?P<month>[0-9]{2})/(?P<slug>[\w-_]+)/',
views.article_detail),
]
以上例子中我们可以看到，所有的 route 字符串前面都加了一个 r ，表示这个字符串是一个原生
字符串。在写正则表达式中是推荐使用原生字符串的，这样可以避免在 python 这一层面进行转
义。而且，使用正则表达式捕获参数的时候，是用一个圆括号进行包裹，然后这个参数的名字是通
过尖括号 <year> 进行包裹，之后才是写正则表达式的语法。
## include函数：
在项目变大以后，经常不会把所有的 url 匹配规则都放在项目的 urls.py 文件中，而是每
个 app 都有自己的 urls.py 文件，在这个文件中存储的都是当前这个 app 的所有 url 匹配规
则。然后再统一注册到项目的 urls.py 文件中。 include 函数有多种用法，这里讲下两种常用的
16
用法。
1. include(pattern,namespace=None) ：直接把其他 app 的 urls 包含进来。示例代码如下：
from django.contrib import admin
from django.urls import path,include
urlpatterns = [
path('admin/', admin.site.urls),
path('book/',include("book.urls"))
]
当然也可以传递 namespace 参数来指定一个实例命名空间，但是在使用实例命名空间之前，
必须先指定一个应用命名空间。示例代码如下：
主urls.py文件：
from django.urls import path,include
urlpatterns = [
path('movie/',include('movie.urls',namespace='movie'))
]
然后在 movie/urls.py 中指定应用命名空间。实例代码如下：
from django.urls import path
from . import views
应用命名空间
app_name = 'movie'
urlpatterns = [
path('',views.movie,name='index'),
path('list/',views.movie_list,name='list'),
]
2. include(pattern_list) ：可以包含一个列表或者一个元组，这个元组或者列表中又包含的
是 path 或者是 re_path 函数。
3. include((pattern,app_namespace),namespace=None) ：在包含某个 app 的 urls 的时候，可
以指定命名空间，这样做的目的是为了防止不同的 app 下出现相同的 url ，这时候就可以通
过命名空间进行区分。示例代码如下：
from django.contrib import admin
from django.urls import path,include
urlpatterns = [
path('admin/', admin.site.urls),
17
path('book/',include(("book.urls",'book')),namespace='book')
]
但是这样做的前提是已经包含了应用命名空间。即在 myapp.urls.py 中添加一个
和 urlpatterns 同级别的变量 app_name 。
## 指定默认的参数：
使用 path 或者是 re_path 的后，在 route 中都可以包含参数，而有时候想指定默认的参数，这
时候可以通过以下方式来完成。示例代码如下：
from django.urls import path
from . import views
urlpatterns = [
path('blog/', views.page),
path('blog/page<int:num>/', views.page),
]
def page(request, num=1):
...
当在访问 blog/ 的时候，因为没有传递 num 参数，所以会匹配到第一个url，这时候就执
行 view.page 这个视图函数，而在 page 函数中，又有 num=1 这个默认参数。因此这时候就可以
不用传递参数。而如果访问 blog/1 的时候，因为在传递参数的时候传递了 num ，因此会匹配到
第二个 url ，这时候也会执行 views.page ，然后把传递进来的参数传给 page 函数中的 num 。
## url反转：
之前我们都是通过url来访问视图函数。有时候我们知道这个视图函数，但是想反转回他的url。这
时候就可以通过 reverse 来实现。示例代码如下：
reverse("list")
> /book/list/
如果有应用命名空间或者有实例命名空间，那么应该在反转的时候加上命名空间。示例代码如下：
reverse('book:list')
> /book/list/
如果这个url中需要传递参数，那么可以通过 kwargs 来传递参数。示例代码如下：
18
reverse("book:detail",kwargs={"book_id":1})
> /book/detail/1
因为 django 中的 reverse 反转 url 的时候不区分 GET 请求和 POST 请求，因此不能在反转的时
候添加查询字符串的参数。如果想要添加查询字符串的参数，只能手动的添加。示例代码如下：
login_url = reverse('login') + "?next=/"
## 自定义URL转换器：
之前已经学到过一些django内置的 url 转换器，包括有 int 、 uuid 等。有时候这些内置的 url
转换器 并不能满足我们的需求，因此django给我们提供了一个接口可以让我们自己定义自己的url
转换器。
自定义 url 转换器按照以下五个步骤来走就可以了：
1. 定义一个类。
2. 在类中定义一个属性 regex ，这个属性是用来保存 url 转换器规则的正则表达式。
3. 实现 to_python(self,value) 方法，这个方法是将 url 中的值转换一下，然后传给视图函数
的。
4. 实现 to_url(self,value) 方法，这个方法是在做 url 反转的时候，将传进来的参数转换后拼
接成一个正确的url。
5. 将定义好的转换器，注册到django中。
比如写一个匹配四个数字年份的 url 转换器。示例代码如下：
1. 定义一个类
class FourDigitYearConverter:
2. 定义一个正则表达式
regex = '[0-9]{4}'
3. 定义to_python方法
def to_python(self, value):
return int(value)
4. 定义to_url方法
def to_url(self, value):
5. 注册到django中
from django.urls import register_converter
register_converter(converters.FourDigitYearConverter, 'yyyy')
urlpatterns = [
path('articles/2003/', views.special_case_2003),
使用注册的转换器
19
path('articles/<yyyy:year>/', views.year_archive),
...
]
