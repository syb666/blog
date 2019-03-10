# **Chapter02**

## **first_project**

\# 第一个项目笔记：

 

\## 创建项目：

\1. 通过命令行的方式：首先要进入到安装了django的虚拟环境中。然后执行命令：

​    \```

​    django-admin startproject [项目的名称]

​    \```

​    这样就可以在当前目录下创建一个项目了。

\2. 通过pycharm的方式：文件->新建项目->选择django。然后指定项目所在的路径，以及Python解释器，再点击Create就可以创建项目了。

 

\## 运行项目：

\1. 终端：进入到项目文件夹中，然后执行以下命令即可运行：

​    \```

​    python manage.py runserver	

​    \```

\2. pycharm：直接点击右上角的绿色三角箭头按钮就可以了。**注意：用pycharm运行项目，要避免一个项目运行多次。**。在项目配置中，把“只用单一实例”那个选项勾选上，避免以上的问题。

 

\## 改变端口号：

\1. 在终端：运行的时候加上一个端口号就可以了。命令为：`python manage.py runserver 9000`。

\2. 在pycharm中：右上角->项目配置->port。改成你想要的端口号，重新运行。

 

\## 让同局域网中的其他电脑访问本机的项目：

\1. 让项目运行到额时候，host为0.0.0.0。

​    \* 在终端，使用命令：`python manage.py runserver 0.0.0.0:8000`。

​    \* 在pycharm，右上角->项目配置->host。改成`0.0.0.0`。

\2. 在`settings.py`文件中，配置`ALLOWED_HOSTS`，将本机的ip地址添加进去。示例代码如下：

​    \```python

​    ALLOWED_HOSTS = ['192.168.0.103']

​    \```

​    注意：要关闭自己电脑的防火墙才行。

 

\## 项目结构分析：

\1. `manange.py`：以后和项目交互基本上都是基于这个文件。一般都是在终端输入python manage.py [子命令]。可以输入python manage.py help看下能做什么事情。除非你知道你自己在做什么，一般情况下不应该编辑这个文件。

\2. `settings.py`：保存项目所有的配置信息。

\3. `urls.py`：用来做url与视图函数映射的。以后来了一个请求，就会从这个文件中找到匹配的视图函数。

\4. `wsig.py`：专门用来做部署的。不需要修改。

 

 

\## django推荐的项目规范：

按照功能或者模块进行分层，分成一个个app。所有和某个模块相关的视图都写在对应的app的views.py中，并且模型和其他的也是类似。然后django已经提供了一个比较方便创建app的命令叫做`python manage.py startapp [app的名称]`。把所有的代码写在各自的app中。

 

 

\## DEBUG模式：

\1. 如果开启了DEBUG模式，那么以后我们修改了Django项目的代码，然后按下ctrl+s，那么Django就会自动的给我们重启项目，不需要手动重启。

\2. 如果开启了DEBUG模式，那么以后Django项目中的代码出现bug了，那么在浏览器中和控制台会打印出错信息。

\3. 在生产环境中，禁止开启DEBUG模式，不然有很大的安全隐患。

\4. 如果将DEBUG设置为False，那么必须要设置ALLOWED_HOSTS.

 

\## ALLOWED_HOSTS：

这个变量是用来设置以后别人只能通过这个变量中的ip地址或者域名来进行访问。

 

## **view_func_demo**

\# 视图函数：

\1. 视图函数的第一个参数必须是request。这个参数绝对不能少。

\2. 视图函数的返回值必须是`django.http.response.HttpResponseBase`的子类的对象。

 

## **url_params_demo**

\# url传递参数：

 

\## url映射：

\1. 为什么会去urls.py文件中寻找映射呢？

是因为在`settings.py`文件中配置了`ROOT_URLCONF`为`urls.py`。所有django会去`urls.py`中寻找。

\2. 在`urls.py`中我们所有的映射，都应该放在`urlpatterns`这个变量中。

\3. 所有的映射不是随便写的，而是使用`path`函数或者是`re_path`函数进行包装的。

 

\## url传参数：

\1. 采用在url中使用变量的方式：在path的第一个参数中，使用`<参数名>`的方式可以传递参数。然后在视图函数中也要写一个参数，视图函数中的参数必须和url中的参数名称保持一致，不然就找不到这个参数。另外，url中可以传递多个参数。

\2. 采用查询字符串的方式：在url中，不需要单独的匹配查询字符串的部分。只需要在视图函数中使用`request.GET.get('参数名称')`的方式来获取。示例代码如下：

​    \```python

​    def author_detail(request):

​        author_id = request.GET['id']

​        text = '作者的id是：%s' % author_id

​        return HttpResponse(text)

​    \```

​    因为查询字符串使用的是`GET`请求，所以我们通过`request.GET`来获取参数。并且因为`GET`是一个类似于字典的数据类型，所有获取值跟字典的方式都是一样的。

 

\## url参数的转换器：

\1. str：除了斜杠`/`以外所有的字符都是可以的。

\2. int：只有是一个或者多个的阿拉伯数字。

\3. path：所有的字符都是满足的。

\4. uuid：只有满足`uuid.uuid4()`这个函数返回的字符串的格式。

\5. slug：英文中的横杆或者英文字符或者阿拉伯数字或者下划线才满足。

## **url_name_demo**

\# url命名：

 

\## 为什么需要url命名？

因为url是经常变化的。如果在代码中写死可能会经常改代码。给url取个名字，以后使用url的时候就使用他的名字进行反转就可以了，就不需要写死url了。

 

\## 如何给一个url指定名称？

在`path`函数中，传递一个`name`参数就可以指定。示例代码如下：

\```python

urlpatterns = [

​    path('',views.index,name='index'),

​    path('login/',views.login,name='login')

]

\```

 

\## 应用命名空间：

在多个app之间，有可能产生同名的url。这时候为了避免反转url的时候产生混淆，可以使用应用命名空间，来做区分。定义应用命名空间非常简单，只要在`app`的`urls.py`中定义一个叫做`app_name`的变量，来指定这个应用的命名空间即可。示例代码如下：

\```python

\# 应用命名空间

app_name = 'front'

 

urlpatterns = [

​    path('',views.index,name='index'),

​    path('login/',views.login,name='login')

]

\```

以后在做反转的时候就可以使用`应用命名空间:url名称`的方式进行反转。示例代码如下：

\```python

login_url = reverse('front:login')

\```

 

\## 应用(app)命名空间和实例命名空间：

一个app，可以创建多个实例。可以使用多个url映射同一个app。所以这就会产生一个问题。以后在做反转的时候，如果使用应用命名空间，那么就会发生混淆。为了避免这个问题。我们可以使用实例命名空间。实例命名空间也是非常简单，只要在`include`函数中传递一个`namespace`变量即可。示例代码如下：

\```python

urlpatterns = [

​    path('',include('front.urls')),

​    \# 同一个app下有两个实例

​    path('cms1/',include('cms.urls',namespace='cms1')),

​    path('cms2/',include('cms.urls',namespace='cms2')),

]

\```

以后在做反转的时候，就可以根据实例命名空间来指定具体的url。示例代码如下：

\```python

def index(request):

​    username = request.GET.get("username")

​    if username:

​        return HttpResponse('CMS首页')

​    else:

​        \# 获取当前的命名空间

​        current_namespace = request.resolver_match.namespace

​        return redirect(reverse("%s:login"%current_namespace))

\```

## **u****rl_include_demo**

\# urls模块化：

如果项目变得越来越大。那么url会变得越来越多。如果都放在主`urls.py`文件中，那么将不太好管理。因此我们可以将每个app自己的urls放到自己的app中进行管理。一般我们会在app中新建一个urls.py文件用来存储所有和这个app相关的子url。

需要注意的地方：

\1. 应该使用`include`函数包含子`urls.py`，并且这个`urls.py`的路径是相对于项目的路径。示例代码如下：

​    \```python

​    urlpatterns = [

​        path('admin/', admin.site.urls),

​        path('book',include('book.urls'))

​    ]

​    \```

\2. 在`app`的`urls.py`中，所有的url匹配也要放在一个叫做`urlpatterns`的变量中，否则找不到。

\3. `url`是会根据主`urls.py`和app中的`urls.py`进行拼接的，因此注意不要多加斜杠。

 

\# include函数的用法：

\1. include(module,namespace=None)：

​    \* module：子url的模块字符串。

​    \* namespace：实例命名空间。这个地方需要注意一点。如果指定实例命名空间，那么前提必须要先指定应用命名空间。也就是在子`urls.py`中添加`app_name`变量。

\2. include((pattern_list, app_namespace), namespace=None)：`include`函数的第一个参数既可以为一个字符串，也可以为一个元组，如果是元组，那么元组的第一个参数是子`urls.py`模块的字符串，元组的第二个参数是应用命名空间。也就是说，应用命名空间既可以在子`urls.py`中通过`app_name`指定，也可以在`include`函数中指定。

\3. include(pattern_list)：`pattern_list`是一个列表。这个列表中装的是`path`或者`re_path`函数。实例代码如下：

​    \```pyhon

​    path('movie/',include([

​        path('',views.movie),

​        path('list/',views.movie_list),

​    ]))

​    \```

## **re_** **path** **_demo**

\# re_path笔记：

\1. re_path和path的作用都是一样的。只不过`re_path`是在写url的时候可以用正则表达式，功能更加强大。

\2. 写正则表达式都推荐使用原生字符串。也就是以`r`开头的字符串。

\3. 在正则表达式中定义变量，需要使用圆括号括起来。这个参数是有名字的，那么需要使用`?P<参数的名字>`。然后在后面添加正则表达式的规则。示例代码如下：

​    \```python

​    from django.urls import re_path

​    from . import views

 

​    urlpatterns = [

​        \# r""：代表的是原生字符串（raw）

​        re_path(r'^$',views.article),

​        \# /article/list/<year>/

​        re_path(r"^list/(?P<year>\d{4})/$",views.article_list),

​        re_path(r"^list/(?P<month>\d{2})/$",views.article_list_month)

​    ]

​    \```

\4. 如果不是特别要求。直接使用`path`就够了，省的把代码搞的很麻烦（因为正则表达式其实是非常晦涩的，特别是一些比较复杂的正则表达式，今天写的明天可能就不记得了）。除非是url中确实是需要使用正则表达式来解决才使用`re_path`。

 

## **reverse_demo**

\# reverse笔记：

\1. 如果在反转url的时候，需要添加参数，那么可以传递`kwargs`参数到`revers`函数中。示例代码如下：

​    \```python

​    detail_url = reverse('detail',kwargs={"article_id":1,'page':2})

​    \```

\2. 如果想要添加查询字符串的参数，则必须手动的进行拼接。示例代码如下：

​    \```python

​    login_url = reverse('login') + "?next=/"

​    \```

## **p****ath****_converter** **_demo**

\# 自定义URL（PATH）转换器笔记：

 

\## 需求：

实现一个获取文章列表的demo，用户可以根据`/articles/文章分类/`的方式来获取文章。其中文章分类采用的是`分类1+分类2+分类3...`的方式拼接的，并且如果只有一个分类，那就不需要加号。示例如下：

\```

\# 1. 第一种：获取python分类下的文章

/articles/python/

\# 2. 第二种：获取python和django分类下的文章

/articles/python+django/

\# 3. 第三种：获取python和django和flask分类下的文章

/articles/python+django+flask/

以此类推...

\```

 

在“文章分类”参数传到视图函数之前要把这些分类分开来存储到列表中。

比如参数是`python+django`，那么传到视图函数的时候就要变成`['python','django']`。

 

以后在使用reverse反转的时候，限制传递“文章分类”的参数应该是一个列表，并且要将这个列表变成`python+django`的形式。

 

 

\## 自定义URL转换器：

之前已经学到过一些django内置的url转换器，包括有int、uuid等。有时候这些内置的url转换器并不能满足我们的需求，因此django给我们提供了一个接口可以让我们自己定义自己的url转换器。

 

自定义url转换器按照以下五个步骤来走就可以了： 

\1. 定义一个类，直接继承自object就可以了。 

\2. 在类中定义一个属性regex，这个属性是用来限制url转换器规则的正则表达式。 

\3. 实现to_python(self,value)方法，这个方法是将url中的值转换一下，然后传给视图函数的。 

\4. 实现to_url(self,value)方法，这个方法是在做url反转的时候，将传进来的参数转换后拼接成一个正确的url。 5. 将定义好的转换器，使用`django.urls.converters.register_converter`方法注册到django中。

 

示例代码如下：

\```python

from django.urls import register_converter

 

class CategoryConverter(object):

​    regex = r'\w+|(\w+\+\w+)+'

 

​    def to_python(self,value):

​        \# python+django+flask

​        \# ['python','django','flask']

​        result = value.split("+")

​        return result

 

​    def to_url(self,value):

​        \# value：['python','django','flask']

​        \# python+django+flask

​        if isinstance(value,list):

​            result = "+".join(value)

​            return result

​        else:

​            raise RuntimeError("转换url的时候，分类参数必须为列表！")

 

register_converter(CategoryConverter,'cate')

\```

## **default_params_demo**

\# URL映射的时候指定默认参数：

使用path或者是re_path的后，在route中都可以包含参数，而有时候想指定默认的参数，这时候可以通过以下方式来完成。示例代码如下：

\```

from django.urls import path

 

from . import views

 

urlpatterns = [

​    path('blog/', views.page),

​    path('blog/page<int:num>/', views.page),

]

 

\# View (in blog/views.py)

def page(request, num=1):

​    \# Output the appropriate page of blog entries, according to num.

​    ...

\```

当在访问blog/的时候，因为没有传递num参数，所以会匹配到第一个url，这时候就执行view.page这个视图函数，而在page函数中，又有num=1这个默认参数。因此这时候就可以不用传递参数。而如果访问blog/1的时候，因为在传递参数的时候传递了num，因此会匹配到第二个url，这时候也会执行views.page，然后把传递进来的参数传给page函数中的num。

# **Chapter03**

## **template_intro_demo**

\# 模版介绍笔记：

 

在之前的章节中，视图函数只是直接返回文本，而在实际生产环境中其实很少这样用，因为实际的页面大多是带有样式的HTML代码，这可以让浏览器渲染出非常漂亮的页面。目前市面上有非常多的模板系统，其中最知名最好用的就是DTL和Jinja2。DTL是Django Template Language三个单词的缩写，也就是Django自带的模板语言。当然也可以配置Django支持Jinja2等其他模板引擎，但是作为Django内置的模板语言，和Django可以达到无缝衔接而不会产生一些不兼容的情况。因此建议大家学习好DTL。

 

\## DTL与普通的HTML文件的区别：

DTL模板是一种带有特殊语法的HTML文件，这个HTML文件可以被Django编译，可以传递参数进去，实现数据动态化。在编译完成后，生成一个普通的HTML文件，然后发送给客户端。

 

\## 渲染模板：

渲染模板有多种方式。这里讲下两种常用的方式。

 

\1. `render_to_string`：找到模板，然后将模板编译后渲染成Python的字符串格式。最后再通过HttpResponse类包装成一个HttpResponse对象返回回去。示例代码如下：

\```python

from django.template.loader import render_to_string

 from django.http import HttpResponse

 def book_detail(request,book_id):

​     html = render_to_string("detail.html")

​     return HttpResponse(html)

\```

\2. 以上方式虽然已经很方便了。但是django还提供了一个更加简便的方式，直接将模板渲染成字符串和包装成HttpResponse对象一步到位完成。示例代码如下：

\```python

 from django.shortcuts import render

 def book_list(request):

​     return render(request,'list.html')

\```

## **template_path_demo**

\# 模版查找路径：

 

在项目的settings.py文件中。有一个TEMPLATES配置，这个配置包含了模板引擎的配置，模板查找路径的配置，模板上下文的配置等。模板路径可以在两个地方配置。

 

\1. `DIRS`：这是一个列表，在这个列表中可以存放所有的模板路径，以后在视图中使用render或者render_to_string渲染模板的时候，会在这个列表的路径中查找模板。

\2. `APP_DIRS`：默认为True，这个设置为True后，会在INSTALLED_APPS的安装了的APP下的templates文件加中查找模板（是按app先加入的顺序对app中的templates进行查找））。

\3. 查找顺序：比如代码render('list.html')。先会在DIRS这个列表中依次查找路径下有没有这个模板，如果有，就返回。如果DIRS列表中所有的路径都没有找到，那么会先检查当前这个视图所处的app是否已经安装，如果已经安装了，那么就先在当前这个app下的templates文件夹中查找模板，如果没有找到，那么会在其他已经安装了的app中查找。如果所有路径下都没有找到，那么会抛出一个TemplateDoesNotExist的异常。

## **template_variable_demo**

\# 模版变量笔记：

\1. 在模版中使用变量，需要将变量放到`{{ 变量 }}`中。

\2. 如果想要访问对象的属性，那么可以通过`对象.属性名`来进行访问。

​    \```python

​    class Person(object):

​        def __init__(self,username):

​            self.username = username

 

​    context = {

​        'person': p

​    }

​    \```

​    以后想要访问`person`的`username`，那么就是通过`person.username`来访问。

\3. 如果想要访问一个字典的key对应的value，那么只能通过`字典.key`的方式进行访问，不能通过`中括号[]`的形式进行访问。

​    \```python

​    context = {

​        'person': {

​            'username':'zhiliao'

​        }

​    }

​    \```

​    那么以后在模版中访问`username`。就是以下代码`person.username`

\4. 因为在访问字典的`key`时候也是使用`点.`来访问，因此不能在字典中定义字典本身就有的属性名当作`key`，否则字典的那个属性将编程字典中的key了。

​    \```python

​    context = {

​        'person': {

​            'username':'zhiliao',

​            'keys':'abc'

​        }

​    }

​    \```

​    以上因为将`keys`作为`person`这个字典的`key`了。因此以后在模版中访问`person.keys`的时候，返回的不是这个字典的所有key，而是对应的值。

\5. 如果想要访问列表或者元组，那么也是通过`点.`的方式进行访问，不能通过`中括号[]`的形式进行访问。这一点和python中是不一样的。示例代码如下：

​    \```python

​    {{ persons.1 }}

\```

## **template_if_demo**

\# if语句笔记：

\1. 所有的标签都是在`{%%}`之间。

\2. if标签有闭合标签。就是`{% endif %}`。

\3. if标签的判断运算符，就跟python中的判断运算符是一样的。`==、!=、<、<=、>、>=、in、not in、is、is not`这些都可以使用。

\4. 还可以使用`elif`以及`else`等标签。

## **template_for_demo**

\# for...in...笔记：

 

\## `for...in...`标签：

`for...in...`类似于`Python`中的`for...in...`。可以遍历列表、元组、字符串、字典等一切可以遍历的对象。示例代码如下：

 

\```python

{% for person in persons %}

<p>{{ person.name }}</p>

{% endfor %}

\```

 

如果想要反向遍历，那么在遍历的时候就加上一个`reversed`。示例代码如下：

 

\```python

{% for person in persons reversed %}

<p>{{ person.name }}</p>

{% endfor %}

\```

 

遍历字典的时候，需要使用`items`、`keys`和`values`等方法。在`DTL`中，执行一个方法不能使用圆括号的形式。遍历字典示例代码如下：

 

\```python

{% for key,value in person.items %}

<p>key：{{ key }}</p>

<p>value：{{ value }}</p>

{% endfor %}

\```

 

在`for`循环中，`DTL`提供了一些变量可供使用。这些变量如下：

 

\* `forloop.counter`：当前循环的下标。以1作为起始值。

\* `forloop.counter0`：当前循环的下标。以0作为起始值。

\* `forloop.revcounter`：当前循环的反向下标值。比如列表有5个元素，那么第一次遍历这个属性是等于5，第二次是4，以此类推。并且是以1作为最后一个元素的下标。

\* `forloop.revcounter0`：类似于forloop.revcounter。不同的是最后一个元素的下标是从0开始。

\* `forloop.first`：是否是第一次遍历。

\* `forloop.last`：是否是最后一次遍历。

\* `forloop.parentloop`：如果有多个循环嵌套，那么这个属性代表的是上一级的for循环。

 

** 模板中的for...in...没有continue和break语句，这一点和Python中有很大的不同，一定要记清楚！ **

\## `for...in...empty`标签：

这个标签使用跟`for...in...`是一样的，只不过是在遍历的对象如果没有元素的情况下，会执行`empty`中的内容。示例代码如下：

 

\```python

{% for person in persons %}

<li>{{ person }}</li>

{% empty %}

暂时还没有任何人

{% endfor %}

\```

## **template_with_deo**

\# with标签笔记：

\1. 在模板中，想要定义变量，可以通过`with`语句来实现。

\2. `with`语句有两种使用方式，第一种是`with xx=xxx`的形式，第二种是`with xxx as xxx`的形式。

\3. 定义的变量只能在with语句块中使用，在with语句块外面使用取不到这个变量。

示例代码如下：

\```python

​    {% with zs=persons.0%}

        <p>{{ zs }}</p>

        <p>{{ zs }}</p>

​    {% endwith %}

​    下面这个因为超过了with语句块，因此不能使用

    <p>{{ zs }}</p>

 

​    {% with persons.0 as zs %}

        <p>{{ zs }}</p>

​    {% endwith %}

\```

## **template_url_demo**

\# url标签笔记：

 

`url`标签：在模版中，我们经常要写一些`url`，比如某个`a`标签中需要定义`href`属性。当然如果通过硬编码的方式直接将这个`url`写死在里面也是可以的。但是这样对于以后项目维护可能不是一件好事。因此建议使用这种反转的方式来实现，类似于`django`中的`reverse`一样。示例代码如下：

 

\```python

<a href="{% url 'book:list' %}">图书列表页面</a>

\```

 

如果`url`反转的时候需要传递参数，那么可以在后面传递。但是参数分位置参数和关键字参数。位置参数和关键字参数不能同时使用。示例代码如下：

 

\```python

\# path部分

path('detail/<book_id>/',views.book_detail,name='detail')

 

\# url反转，使用位置参数

<a href="{% url 'book:detail' 1 %}">图书详情页面</a>

 

\# url反转，使用关键字参数

<a href="{% url 'book:detail' book_id=1 %}">图书详情页面</a>

\```

 

如果想要在使用`url`标签反转的时候要传递查询字符串的参数，那么必须要手动在在后面添加。示例代码如下：

 

\```python

<a href="{% url 'book:detail' book_id=1 %}?page=1">图书详情页面</a>

\```

 

如果需要传递多个参数，那么通过空格的方式进行分隔。示例代码如下：

 

\```python

<a href="{% url 'book:detail' book_id=1 page=2 %}">图书详情页面</a>

\```

## **templae_autoescape_demo**

\# autoescape自动转义笔记：

 

\1. DTL中默认已经开启了自动转义。会将那些特殊字符进行转义。比如会将`<`转义成`<`等。

\2. 如果你不知道自己在干什么，那么最好是使用DTL的自动转义。这样网站才不容易出现XSS漏洞。

\3. 如果变量确实是可信任的。那么可以使用`autoescape`标签来关掉自动转义。示例代码如下：

​    \```python

​    {% autoescape off %}

​        {{ info }}

​    {% endautoescape %}

\```

## **template_verbatim_demo**

\# verbatim标签笔记：

 

`verbatim`标签：默认在`DTL`模板中是会去解析那些特殊字符的。比如`{%`和`%}`以及`{{`等。如果你在某个代码片段中不想使用`DTL`的解析引擎。那么你可以把这个代码片段放在`verbatim`标签中。示例代码下：

\```python

{% verbatim %}

{{if dying}}Still alive.{{/if}}

{% endverbatim %}

\```

## **template_filter_demo**

\# Django模板过滤器笔记：

 

\## 为什么需要过滤器？

因为在DTL中，不支持函数的调用形式`()`，因此不能给函数传递参数，这将有很大的局限性。而过滤器其实就是一个函数，可以对需要处理的参数进行处理，并且还可以额外接收一个参数（也就是说，最多只能有2个参数）。

 

\## add过滤器：

将传进来的参数添加到原来的值上面。这个过滤器会尝试将`值`和`参数`转换成整形然后进行相加。如果转换成整形过程中失败了，那么会将`值`和`参数`进行拼接。如果是字符串，那么会拼接成字符串，如果是列表，那么会拼接成一个列表。示例代码如下：

 

\```python

{{ value|add:"2" }}

\```

 

如果`value`是等于4，那么结果将是6。如果`value`是等于一个普通的字符串，比如`abc`，那么结果将是`abc2`。`add`过滤器的源代码如下：

 

\```python

def add(value, arg):

"""Add the arg to the value."""

try:

return int(value) + int(arg)

except (ValueError, TypeError):

try:

return value + arg

except Exception:

return ''

\```

 

\## cut过滤器：

移除值中所有指定的字符串。类似于`python`中的`replace(args,"")`。示例代码如下：

 

\```python

{{ value|cut:" " }}

\```

 

以上示例将会移除`value`中所有的空格字符。`cut`过滤器的源代码如下：

 

\```python

def cut(value, arg):

"""Remove all values of arg from the given string."""

safe = isinstance(value, SafeData)

value = value.replace(arg, '')

if safe and arg != ';':

return mark_safe(value)

return value

\```

 

\## `date`过滤器：

将

一个日期按照指定的格式，格式化成字符串。示例代码如下：

 

\```python

\# 数据

context = {

"birthday": datetime.now()

}

 

\# 模版

{{ birthday|date:"Y/m/d" }}

\```

 

那么将会输出`2018/02/01`。其中`Y`代表的是四位数字的年份，`m`代表的是两位数字的月份，`d`代表的是两位数字的日。  

还有更多时间格式化的方式。见下表。

 

| 格式字符 | 描述 | 示例 |

| --- | --- | --- |

| Y | 四位数字的年份 | 2018 |

| m | 两位数字的月份 | 01-12 |

| n | 月份，1-9前面没有0前缀 | 1-12 |

| d | 两位数字的天 | 01-31 |

| j | 天，但是1-9前面没有0前缀 | 1-31 |

| g | 小时，12小时格式的，1-9前面没有0前缀 | 1-12 |

| h | 小时，12小时格式的，1-9前面有0前缀 | 01-12 |

| G | 小时，24小时格式的，1-9前面没有0前缀 | 1-23 |

| H | 小时，24小时格式的，1-9前面有0前缀 | 01-23 |

| i | 分钟，1-9前面有0前缀 | 00-59 |

| s | 秒，1-9前面有0前缀 | 00-59 |

 

 

\#

\## default

 

如果值被评估为`False`。比如`[]`，`""`，`None`，`{}`等这些在`if`判断中为`False`的值，都会使用`default`过滤器提供的默认值。示例代码如下：

 

\```python

{{ value|default:"nothing" }}

\```

 

如果`value`是等于一个空的字符串。比如`""`，那么以上代码将会输出`nothing`。

 

\### default\_if\_none

 

如果值是`None`，那么将会使用`default_if_none`提供的默认值。这个和`default`有区别，`default`是所有被评估为`False`的都会使用默认值。而`default_if_none`则只有这个值是等于`None`的时候才会使用默认值。示例代码如下：

 

\```python

{{ value|default_if_none:"nothing" }}

\```

 

如果`value`是等于`""`也即空字符串，那么以上会输出空字符串。如果`value`是一个`None`值，以上代码才会输出`nothing`。

 

\### first

 

返回列表/元组/字符串中的第一个元素。示例代码如下：

 

\```python

{{ value|first }}

\```

 

如果`value`是等于`['a','b','c']`，那么输出将会是`a`。

 

\### last

 

返回列表/元组/字符串中的最后一个元素。示例代码如下：

 

\```python

{{ value|last }}

\```

 

如果`value`是等于`['a','b','c']`，那么输出将会是`c`。

 

\### floatformat

 

使用四舍五入的方式格式化一个浮点类型。如果这个过滤器没有传递任何参数。那么只会在小数点后保留一个小数，如果小数后面全是0，那么只会保留整数。当然也可以传递一个参数，标识具体要保留几个小数。

 

\1. 如果没有传递参数：

 

| value | 模版代码 | 输出 |

| --- | --- | --- |

| 34.23234 | `{{ value\|floatformat }}` | 34.2 |

| 34.000 | `{{ value\|floatformat }}` | 34 |

| 34.260 | `{{ value\|floatformat }}` | 34.3 |

 

\2. 如果传递参数：

 

| value | 模版代码 | 输出 |

| --- | --- | --- |

| 34.23234 | `{{value\|floatformat:3}}` | 34.232 |

| 34.0000 | `{{value\|floatformat:3}}` | 34.000 |

| 34.26000 | `{{value\|floatformat:3}}` | 34.260 |

 

\### join

 

类似与`Python`中的`join`，将列表/元组/字符串用指定的字符进行拼接。示例代码如下：

 

\```python

{{ value|join:"/" }}

\```

 

如果`value`是等于`['a','b','c']`，那么以上代码将输出`a/b/c`。

 

\### length

 

获取一个列表/元组/字符串/字典的长度。示例代码如下：

 

\```python

{{ value|length }}

\```

 

如果`value`是等于`['a','b','c']`，那么以上代码将输出`3`。如果`value`为`None`，那么以上将返回`0`。

 

\### lower

 

将值中所有的字符全部转换成小写。示例代码如下：

 

\```python

{{ value|lower }}

\```

如果`value`是等于`Hello World`。那么以上代码将输出`hello world`。

\### upper

类似于`lower`，只不过是将指定的字符串全部转换成大写。

\### random

在被给的列表/字符串/元组中随机的选择一个值。示例代码如下：

\```python

{{ value|random }}

\```

如果`value`是等于`['a','b','c']`，那么以上代码会在列表中随机选择一个。

\### safe

标记一个字符串是安全的。也即会关掉这个字符串的自动转义。示例代码如下：

\```python

{{value|safe}}

\```

如果`value`是一个不包含任何特殊字符的字符串，比如`<a>`这种，那么以上代码就会把字符串正常的输入。如果`value`是一串`html`代码，那么以上代码将会把这个`html`代码渲染到浏览器中。

\### slice

类似于`Python`中的切片操作。示例代码如下：

\```python

{{ some_list|slice:"2:" }}

\```

以上代码将会给`some_list`从`2`开始做切片操作。

\### stringtags

删除字符串中所有的`html`标签。示例代码如下：

\```python

{{ value|striptags }}

\```

如果`value`是`<strong>hello world</strong>`，那么以上代码将会输出`hello world`。

\### truncatechars

如果给定的字符串长度超过了过滤器指定的长度。那么就会进行切割，并且会拼接三个点来作为省略号。示例代码如下：

\```python

{{ value|truncatechars:5 }}

\```

如果`value`是等于`北京欢迎您~`，那么输出的结果是`北京...`。可能你会想，为什么不会`北京欢迎您...`呢。因为三个点也占了三个字符，所以`北京`+三个点的字符长度就是5。

\### truncatechars\_html

类似于`truncatechars`，只不过是不会切割`html`标签。示例代码如下：

\```python

{{ value|truncatechars:5 }}

\```

如果`value`是等于`<p>北京欢迎您~</p>`，那么输出将是`<p>北京...</p>`。

## **custom_filter_demo**

\# 自定义过滤器笔记：

\1. 首先在某个app中，创建一个python包，叫做`templatetags`，注意，这个包的名字必须为`templatetags`，不然就找不到。

\2. 在这个`templatetags`包下面，创建一个python文件用来存储过滤器。

\3. 在新建的python文件中，定义过滤器（也就是函数），这个函数的第一个参数永远是被过滤的那个值，并且如果在使用过滤器的时候传递参数，那么还可以定义另外一个参数。但是过滤器最多只能有2个参数。

\4. 在写完过滤器（函数）后，要使用`django.template.Library.filter`进行注册。

\5. 还要把这个过滤器所在的这个app添加到`settings.INSTALLED_APS`中，不然Django也找不到这个过滤器。

\6. 在模板中使用`load`标签加载过滤器所在的python包。

 

\7. 可以使用过滤器了。

\8. `django.template.Library.filter`还可以当作装饰器来使用。如果`filter`函数没有传递任何参数，那么将会使用这个函数的名字来作为过滤器的名字。当然如果你不想使用函数的名字来作为过滤器的名字，也可以传递一个`name`参数。示例代码如下：

​    \```python

​    @register.filter('my_greet')

​    def greet(value,word):

​        return value + word

\```

## **template_include_demo**

\# include笔记：

 

\1. 有些模版代码是重复的。因此可以单独抽取出来，以后哪里需要用到，就直接使用`include`进来就可以了。

\2. 如果想要在`include`子模版的时候，传递一些参数，那么可以使用`with xxx=xxx`的形式。示例代码如下：

​    \```python

​    {% include 'header.html' with username='zhiliao' %}

\```

## **template_inherit_demo**

\# 模版继承笔记：

 

在前端页面开发中。有些代码是需要重复使用的。这种情况可以使用`include`标签来实现。也可以使用另外一个比较强大的方式来实现，那就是模版继承。模版继承类似于`Python`中的类，在父类中可以先定义好一些变量和方法，然后在子类中实现。模版继承也可以在父模版中先定义好一些子模版需要用到的代码，然后子模版直接继承就可以了。并且因为子模版肯定有自己的不同代码，因此可以在父模版中定义一个block接口，然后子模版再去实现。以下是父模版的代码：

 

\```html

{% load static %}

<!DOCTYPE html>

<html lang="en">

<head>

<link rel="stylesheet" href="{% static 'style.css' %}" />

<title>{% block title %}我的站点{% endblock %}</title>

</head>

 

<body>

<div id="sidebar">

{% block sidebar %}

<ul>

<li><a href="/">首页</a></li>

<li><a href="/blog/">博客</a></li>

</ul>

{% endblock %}

</div>

<div id="content">

{% block content %}{% endblock %}

</div>

</body>

</html>

\```

 

这个模版，我们取名叫做`base.html`，定义好一个简单的`html`骨架，然后定义好两个`block`接口，让子模版来根据具体需求来实现。子模板然后通过`extends`标签来实现，示例代码如下：

 

\```html

{% extends "base.html" %}

 

{% block title %}博客列表{% endblock %}

 

{% block content %}

{% for entry in blog_entries %}

<h2>{{ entry.title }}</h2>

<p>{{ entry.body }}</p>

{% endfor %}

{% endblock %}

\```

**需要注意的是：extends标签必须放在模版的第开始的位置**

**子模板中的代码必须放在block中，否则将不会被渲染。**

如果在某个`block`中需要使用父模版的内容，那么可以使用`{{block.super}}`来继承。比如上例，`{%block title%}`，如果想要使用父模版的`title`，那么可以在子模版的`title block`中使用`{{ block.super }}`来实现。

 

在定义`block`的时候，除了在`block`开始的地方定义这个`block`的名字，还可以在`block`结束的时候定义名字。比如`{% block title %}{% endblock title %}`。这在大型模版中显得尤其有用，能让你快速的看到`block`包含在哪里。

## **template_static_demo**

\# 加载静态文件笔记：

 

 

在一个网页中，不仅仅只有一个`html`骨架，还需要`css`样式文件，`js`执行文件以及一些图片等。因此在`DTL`中加载静态文件是一个必须要解决的问题。在`DTL`中，使用`static`标签来加载静态文件。要使用`static`标签，首先需要`{% load static %}`。加载静态文件的步骤如下：

 

\1. 首先确保`django.contrib.staticfiles`已经添加到`settings.INSTALLED_APPS`中。

 

\2. 确保在`settings.py`中设置了`STATIC_URL`。

 

\3. 在已经安装了的`app`下创建一个文件夹叫做`static`，然后再在这个`static`文件夹下创建一个当前`app`的名字的文件夹，再把静态文件放到这个文件夹下。例如你的`app`叫做`book`，有一个静态文件叫做`zhiliao.jpg`，那么路径为`book/static/book/zhiliao.jpg`。（为什么在`app`下创建一个`static`文件夹，还需要在这个`static`下创建一个同`app`名字的文件夹呢？原因是如果直接把静态文件放在`static`文件夹下，那么在模版加载静态文件的时候就是使用`zhiliao.jpg`，如果在多个`app`之间有同名的静态文件，这时候可能就会产生混淆。而在`static`文件夹下加了一个同名`app`文件夹，在模版中加载的时候就是使用`app/zhiliao.jpg`，这样就可以避免产生混淆。）

 

\4. 如果有一些静态文件是不和任何`app`挂钩的。那么可以在`settings.py`中添加`STATICFILES_DIRS`，以后`DTL`就会在这个列表的路径中查找静态文件。比如可以设置为:

 

\```python

STATICFILES_DIRS = [

os.path.join(BASE_DIR,"static")

]

\```

 

\5. 在模版中使用`load`标签加载`static`标签。比如要加载在项目的`static`文件夹下的`style.css`的文件。那么示例代码如下：

 

\```html

{% load static %}

<link rel="stylesheet" href="{% static 'style.css' %}">

\```

 

\6. 如果不想每次在模版中加载静态文件都使用`load`加载`static`标签，那么可以在`settings.py`中的`TEMPLATES/OPTIONS`添加`'builtins':['django.templatetags.static']`，这样以后在模版中就可以直接使用`static`标签，而不用手动的`load`了。

 

\7. 如果没有在`settings.INSTALLED_APPS`中添加`django.contrib.staticfiles`。那么我们就需要手动的将请求静态文件的`url`与静态文件的路径进行映射了。示例代码如下：

 

\```python

from django.conf import settings

from django.conf.urls.static import static

 

urlpatterns = [

\# 其他的url映射

] + static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)

\```

 

# **Chapter04**

## **base_orm_operate**

\# ORM对数据库的基本操作：

\## 添加数据：

只要使用ORM模型创建一个对象。然后再调用这个ORM模型的`save`方法就可以保存了。

示例代码如下：

\```python

book = Book(name='西游记',author='吴承恩',price=100)

book.save()

\```

 

\## 查找数据：

所有的查找工作都是使用模型上的`objects`属性来完成的。当然也可以自定义查询对象。这部分功能会在后面讲到。

\1. 根据主键进行查找：使用主键进行查找。可以使用`objects.get`方法。然后传递`pk=xx`的方式进行查找。示例代码如下：

​    \```python

​    book = Book.objects.get(pk=2)

​    \```

\2. 根据其他字段进行查找：可以使用`objects.filter`方法进行查找。示例代码如下：

​    \```python

​    books = Book.objects.filter(name='三国演义')

​    \```

​    使用`filter`方法返回来的是一个`QuerySet`对象。这个对象类似于列表。我们可以使用这个对象的`first`方法来获取第一个值。

 

 

\## 删除数据：

首先查找到对应的数据模型。然后再执行这个模型的`delete`方法即可删除。示例代码如下：

\```python

book = Book.objects.get(pk=1)

book.delete()

\```

 

 

\## 修改数据：

首先查找到对应的数据模型。然后修改这个模型上的属性的值。再执行`save`方法即可修改完成。示例代码如下：

\```python

​    book = Book.objects.get(pk=2)

​    book.price = 200

​    book.save()

\```

## **orm_field_demo**

\# 常用Field笔记：

 

\## navie时间和aware时间：

\### 什么是navie时间？什么是aware时间？

\1. navie时间：不知道自己的时间表示的是哪个时区的。也就是不知道自己几斤几两。比较幼稚。

\2. aware时间：知道自己的时间表示的是哪个时区的。也就是比较清醒。

 

\### pytz库：

专门用来处理时区的库。这个库会经常更新一些时区的数据，不需要我们担心。并且这个库在安装Django的时候会默认的安装。如果没有安装，那么可以通过`pip install pytz`的方式进行安装。

 

\### astimezone方法：

将一个时区的时间转换为另外一个时区的时间。这个方法只能被`aware`类型的时间调用。不能被`navie`类型的时间调用。

示例代码

\```python

import pytz如下：

from datetime import datetime

now = datetime.now() # 这是一个navie类型的时间

utc_timezone = pytz.timezone("UTC") # 定义UTC的时区对象

utc_now = now.astimezone(utc_timezone) # 将当前的时间转换为UTC时区的时间

\>> ValueError: astimezone() cannot be applied to a naive datetime # 会抛出一个异常，原因就是因为navie类型的时间不能调用astimezone方法

 

 

now = now.replace(tzinfo=pytz.timezone('Asia/Shanghai'))

utc_now = now.astimezone(utc_timezone)

\# 这时候就可以正确的转换。

\```

 

\### replace方法：

可以将一个时间的某些属性进行更改。

 

\### django.utils.timezone.now方法：

会根据`settings.py`中是否设置了`USE_TZ=True`获取当前的时间。如果设置了，那么就获取一个`aware`类型的`UTC`时间。如果没有设置，那么就会获取一个`navie`类型的时间。

 

\### django.utils.timezone.localtime方法：

会根据`setting.py`中的`TIME_ZONE`来将一个`aware`类型的时间转换为`TIME_ZONE`指定时区的时间。

 

\## DateField：

 

日期类型。在`Python`中是`datetime.date`类型，可以记录年月日。在映射到数据库中也是`date`类型。使用这个`Field`可以传递以下几个参数：

\1. `auto_now`：在每次这个数据保存的时候，都使用当前的时间。比如作为一个记录修改日期的字段，可以将这个属性设置为`True`。

\2. `auto_now_add`：在每次数据第一次被添加进去的时候，都使用当前的时间。比如作为一个记录第一次入库的字段，可以将这个属性设置为`True`。

 

\## DateTimeField：

 

日期时间类型，类似于`DateField`。不仅仅可以存储日期，还可以存储时间。映射到数据库中是`datetime`类型。这个`Field`也可以使用`auto_now`和`auto_now_add`两个属性。

 

\## TimeField：

 

时间类型。在数据库中是`time`类型。在`Python`中是`datetime.time`类型。

 

\### navie和aware介绍以及在django中的用法：

https://docs.djangoproject.com/en/2.0/topics/i18n/timezones/

 

 

\## EmailField：

类似于`CharField`。在数据库底层也是一个`varchar`类型。最大长度是254个字符。

 

\## FileField：

用来存储文件的。这个请参考后面的文件上传章节部分。

 

\### ImageField：

用来存储图片文件的。这个请参考后面的图片上传章节部分。

 

\### FloatField：

浮点类型。映射到数据库中是`float`类型。

 

\### IntegerField：

整形。值的区间是`-2147483648——2147483647`。

 

\### BigIntegerField：

大整形。值的区间是`-9223372036854775808——9223372036854775807`。

 

\### PositiveIntegerField：

正整形。值的区间是`0——2147483647`。

 

\### SmallIntegerField：

小整形。值的区间是`-32768——32767`。

 

\### PositiveSmallIntegerField：

正小整形。值的区间是`0——32767`。

 

\### TextField：

大量的文本类型。映射到数据库中是longtext类型。

 

\### UUIDField：

只能存储`uuid`格式的字符串。`uuid`是一个32位的全球唯一的字符串，一般用来作为主键。

 

\### URLField：

类似于`CharField`，只不过只能用来存储`url`格式的字符串。并且默认的`max_length`是200。

 

 

\## Field常用的参数

 

\### null：

 

如果设置为`True`，`Django`将会在映射表的时候指定是否为空。默认是为`False`。在使用字符串相关的`Field`（CharField/TextField）的时候，官方推荐尽量不要使用这个参数，也就是保持默认值`False`。因为`Django`在处理字符串相关的`Field`的时候，即使这个`Field`的`null=False`，如果你没有给这个`Field`传递任何值，那么`Django`也会使用一个空的字符串`""`来作为默认值存储进去。因此如果再使用`null=True`，`Django`会产生两种空值的情形（NULL或者空字符串）。如果想要在表单验证的时候允许这个字符串为空，那么建议使用`blank=True`。如果你的`Field`是`BooleanField`，那么对应的可空的字段则为`NullBooleanField`。

 

\### blank：

 

标识这个字段在表单验证的时候是否可以为空。默认是`False`。

这个和`null`是有区别的，`null`是一个纯数据库级别的。而`blank`是表单验证级别的。

 

\### db\_column：

 

这个字段在数据库中的名字。如果没有设置这个参数，那么将会使用模型中属性的名字。

 

\### default：

 

默认值。可以为一个值，或者是一个函数，但是不支持`lambda`表达式。并且不支持列表/字典/集合等可变的数据结构。

 

\### primary\_key：

 

是否为主键。默认是`False`。

 

\### unique：

 

在表中这个字段的值是否唯一。一般是设置手机号码/邮箱等。

更多`Field`参数请参考官方文档：[https://docs.djangoproject.com/zh-hans/2.0/ref/models/fields/](https://docs.djangoproject.com/zh-hans/2.0/ref/models/fields/)

 

 

 

\## 模型中`Meta`配置：

 

对于一些模型级别的配置。我们可以在模型中定义一个类，叫做`Meta`。然后在这个类中添加一些类属性来控制模型的作用。比如我们想要在数据库映射的时候使用自己指定的表名，而不是使用模型的名称。那么我们可以在`Meta`类中添加一个`db_table`的属性。示例代码如下：

\```python

class Book(models.Model):

​    name = models.CharField(max_length=20,null=False)

​    desc = models.CharField(max_length=100,name='description',db_column="description1")

 

class Meta:

​    db_table = 'book_model'

\```

 

以下将对`Meta`类中的一些常用配置进行解释。

 

\### db_table：

这个模型映射到数据库中的表名。如果没有指定这个参数，那么在映射的时候将会使用模型名来作为默认的表名。

 

\### ordering：

设置在提取数据的排序方式。后面章节会讲到如何查找数据。比如我想在查找数据的时候根据添加的时间排序，那么示例代码如下：

\```python

class Book(models.Model):

name = models.CharField(max_length=20,null=False)

desc = models.CharField(max_length=100,name='description',db_column="description1")

pub_date = models.DateTimeField(auto_now_add=True)

 

class Meta:

db_table = 'book_model'

ordering = ['pub_date']

\```

 

更多的配置后面会慢慢介绍到。

官方文档：<https://docs.djangoproject.com/en/2.0/ref/models/options/>

 

## **orm_relationship_demo**

\# 表关系笔记：

 

\## 一对多：

\1. 应用场景：比如文章和作者之间的关系。一个文章只能由一个作者编写，但是一个作者可以写多篇文章。文章和作者之间的关系就是典型的多对一的关系。

\2. 实现方式：一对多或者多对一，都是通过`ForeignKey`来实现的。还是以文章和作者的案例进行讲解。

 

\```python

class User(models.Model):

​    username = models.CharField(max_length=20)

​    password = models.CharField(max_length=100)

 

class Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

​    author = models.ForeignKey("User",on_delete=models.CASCADE)

\```

 

那么以后在给`Article`对象指定`author`，就可以使用以下代码来完成：

 

\```python

article = Article(title=',abc'content='123')

author = User(username='zhiliao',password='111111')

\# 要先保存到数据库中

author.save()

article.author = author

article.save()

\```

 

并且以后如果想要获取某个用户下所有的文章，可以通过`article_set`来实现。示例代码如下：

 

\```python

user = User.objects.first()

\# 获取第一个用户写的所有文章

articles = user.article_set.all()

for article in articles:

​    print(article)

\```

 

并且如果想要将文章添加到某个分类中。可以使用一下的方式：

\```python

category = Category.objects.first()

 

article = Article(title='bbb',content='vvv')

article.author = FrontUser.objects.first()

 

category.article_set.add(article,bulk=False)

\```

\* 使用`bulk=False`，那么Django会自动的保存article，而不需要在添加到category之前先保存article。

\* 或者是另外一种解决方式是，在添加到`category.article_set`中之前，先将`article`保存到数据库中。但是如果`article.category`不能为空，那么就产生一种死循环了，article没有`category`不能保存，而将article添加到`cateogry.artile_set`中，又需要article之前是已经存储到数据库中的。

\* 如果是上面的那种需求，建议使用`bulk=False`的解决方案。

 

\## 一对一：

\1. 在Django中一对一是通过`models.OnetToOneField`来实现的。这个`OneToOneField`其实本质上就是一个外键，只不过这个外键有一个`唯一约束（unique key）`，来实现一对一。

\2. 以后如果想要反向引用，那么是通过引用的模型的名字转换为小写的形式进行访问。比如以下模型：

​    \```python

​    class FrontUser(models.Model):

​        username = models.CharField(max_length=200)

 

​    class UserExtension(models.Model):

​        school = models.CharField(max_length=100)

​        user = models.OneToOneField("FrontUser",on_delete=models.CASCADE)

 

​    \# 通过userextension来访问UserExtension对象

​    user = FrontUser.objects.first()

​    print(user.userextension)

​    \```

​    `UserExtension`的对象，可以通过`user`来访问到对应的user对象。并且`FrontUser`对象可以使用`userextension`来访问对应的`UserExtension`对象。

​    如果不想使用Django默认的引用属性名字。那么可以在`OneToOneField`中添加一个`related_name`参数。示例代码如下：

​    \```python

​    class FrontUser(models.Model):

​        username = models.CharField(max_length=200)

 

​    class UserExtension(models.Model):

​        school = models.CharField(max_length=100)

​        user = models.OneToOneField("FrontUser",on_delete=models.CASCADE,related_name='extension')

 

​    \# 通过extension来访问到UserExtension对象

​    user = FrontUser.objects.first()

​    print(user.extension)

​    \```

​    那么以后就`FrontUser`的对象就可以通过`extension`属性来访问到对应的`UserExtension`对象。

 

\## 多对多：

\1. 应用场景：比如文章和标签的关系。一篇文章可以有多个标签，一个标签可以被多个文章所引用。因此标签和文章的关系是典型的多对多的关系。

 

\2. 实现方式：`Django`为这种多对多的实现提供了专门的`Field`。叫做`ManyToManyField`。还是拿文章和标签为例进行讲解。示例代码如下：

 

\```python

class Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

​    tags = models.ManyToManyField("Tag",related_name="articles")

 

class Tag(models.Model):

​    name = models.CharField(max_length=50)

\```

 

在数据库层面，实际上`Django`是为这种多对多的关系建立了一个中间表。这个中间表分别定义了两个外键，引用到`article`和`tag`两张表的主键。

 

 

## **orm_lookup_demo**

 

\# 查询条件笔记：

 

\1. exact：在底层会被翻译成`=`。

 

\2. iexact：在底层会被翻译成`LIKE`。

​    \* LIKE和=：大部分情况下都是等价的，只有少数情况下是不等价的。

​    \* exict和iexact：他们的区别其实就是LIKE和=的区别，因为exact会被翻译成=，而iexact会被翻译成LIKE。

​    \* 因为`field__exact=xxx`其实等价于`filed=xxx`，因此我们直接使用`filed=xxx`就可以了，并且因为大部分情况`exact`和`iexact`又是等价的，因此我们以后直接使用`field=xxx`就可以了。

 

\3. QuerySet.query：`query`可以用来查看这个`ORM`查询语句最终被翻译成的`SQL`语句。但是`query`只能被用在`QuerySet`对象上，不能用在普通的`ORM模型`上。因此如果你的查询语句是通过`get`来获取数据的，那么就不能使用`query`，因为`get`返回的是满足条件的`ORM`模型，而不是`QuerySet`。如果你是通过`filter`等其他返回`QuerySet`的方法查询的，那么就可以使用`query`。

 

\4. contains：使用大小写敏感的判断，某个字符串是否在指定的字段中。这个判断条件会使用大小敏感，因此在被翻译成`SQL`语句的时候，会使用`like binary`，而`like binary`就是使用大小写敏感的。

 

\5. icontains：使用大小写不敏感的判断，某个字符串是否被包含在指定的字段中。这个查询语句在被翻译成`SQL`的时候，使用的是`like`，而`like`在`MySQL`层面就是不区分大小写的。

\6. contains和icontains：在被翻译成`SQL`的时候使用的是`%hello%`，就是只要整个字符串中出现了`hello`都能过够被找到，而`iexact`没有百分号，那么意味着只有完全相等的时候才会被匹配到。

 

\7. in：可以直接指定某个字段的是否在某个集合中。示例代码如下：

​    \```python

​    articles = Article.objects.filter(id__in=[1,2,3])

​    \```

​    也可以通过其他的表的字段来判断是否在某个集合中。示例代码如下：

​    \```python

​    categories = Category.objects.filter(article__id__in=[1,2,3])

​    \```

​    如果要判断相关联的表的字段，那么也是通过`__`来连接。并且在做关联查询的时候，不需要写`models_set`，直接使用`模型的名字的小写化`就可以了。比如通过分类去查找相应的文章，那么通过`article__id__in`就可以了，而不是写成`article_set__id__in`的形式。当然如果你不想使用默认的形式，可以在外键定义的时候传递`related_query_name`来指定反向查询的名字。示例代码如下：

​    \```python

​    class Category(models.Model):

​        name = models.CharField(max_length=100)

 

​        class Meta:

​            db_table = 'category'

 

 

​    class Article(models.Model):

​        title = models.CharField(max_length=200)

​        content = models.TextField()

​        cateogry = models.ForeignKey("Category",on_delete=models.CASCADE,null=True,related_query_name='articles')

 

​        class Meta:

​            db_table = 'article'

​    \```

​    因为在`cateogry`的`ForeignKey`中指定了`related_query_name`为`articles`，因此你不能再使用`article`来进行反向查询了。这时候就需要通过`articles__id__in`来进行反向查询。

 

​    反向查询是将模型名字小写化。比如`article__in`。可以通过`related_query_name`来指定自己的方式，而不使用默认的方式。

​    反向引用是将模型名字小写化，然后再加上`_set`，比如`article_set`，可以通过`related_name`来指定自己的方式，而不是用默认的方式。

 

​    并且，如果在做反向查询的时候，如果查询的字段就是模型的主键，那么可以省略掉这个字段，直接写成`article__in`就可以了，不需要这个`id`了。

 

​    `in`不仅仅可以指定列表/元组，还可以为`QuerySet`。比如要查询“文章标题中包含有hello的所有分类”，那么可以通过以下代码来实现：

​    \```python

​    articles = Article.objects.filter(title__icontains='hello')

​    categories = Category.objects.filter(articles__in=articles)

​    for cateogry in categories:

​        print(cateogry)

​    \```

 

\8. gt、gte、lt、lte：代表的是大于、大于等于、小于、小于等于的条件。示例代码如下：

​    \```python

​    articles = Article.objects.filter(id__lte=3)

​    \```

 

\9. startswith、istartswith、endswith、iendswith：表示以某个值开始，不区分大小写的以某个值开始、以某个值结束、不区分大小写的以某个值结束。示例代码如下：

​    \```python

​    articles = Article.objects.filter(title__endswith="hello")

​    \```

 

\10. 关于时间的查询条件：

​    \* range：可以指定一个时间段。并且时间应该标记为`aware`时间，不然django会报警告。示例代码如下：

​        \```python

​        start_time = make_aware(datetime(year=2018,month=4,day=4,hour=17,minute=0,second=0))

​        end_time = make_aware(datetime(year=2018,month=4,day=4,hour=18,minute=0,second=0))

​        articles = Article.objects.filter(create_time__range=(start_time,end_time))

​        print(articles.query)

​        print(articles)

​        \```

 

​    \* date：用年月日来进行过滤。如果想要使用这个过滤条件，那么前提必须要在`MySQL`中添加好那些时区文件。如何添加呢？参考教案。示例代码如下：

​        \```python

​        articles = Article.objects.filter(create_time__date=datetime(year=2018,month=4,day=4))

​        \```

​    \* year/month/day：表示根据年/月/日进行查找。示例代码如下：

​        \```python

​        articles = Article.objects.filter(create_time__year__gte=2018)

​        \```

​    \* week_day：根据星期来进行查找。1表示星期天，7表示星期六，2-6代表的是星期一到星期五。比如要查找星期三的所有文章，那么可以通过以下代码来实现：

​        \```python

​        articles = Article.objects.filter(create_time__week_day=4)

​        \```

​    \* time：根据分时秒来进行查找。如果要具体到秒，一般比较难匹配到，可以使用区间的方式来进行查找。区间使用`range`条件。比如想要获取17时/10分/27-28秒之间的文章，那么可以通过以下代码来实现：

​        \```python

​        start_time = time(hour=17,minute=10,second=27)

​        end_time = time(hour=17,minute=10,second=28)

​        articles = Article.objects.filter(create_time__time__range=(start_time,end_time))

​        \```

## **orm_aggregate_demo**

\# 聚合函数笔记：

 

\1. 所有的聚合函数都是放在`django.db.models`下面。

\2. 聚合函数不能够单独的执行，需要放在一些可以执行聚合函数的方法下面中去执行。比如`aggregate`。示例代码如下：

​    \```python

​    result = Book.objects.aggregate(Avg("price"))

​    \```

\3. 聚合函数执行完成后，给这个聚合函数的值取个名字。取名字的规则，默认是`filed+__+聚合函数名字`形成的。比如以上代码形成的名字叫做`price__avg`。如果不想使用默认的名字，那么可以在使用聚合函数的时候传递关键字参数进去，参数的名字就是聚合函数执行完成的名字。实示例代码如下：

​    \```python

​    result = Book.objects.aggregate(avg=Avg("price"))

​    \```

​    以上传递了关键字参数`avg=Avg("price")`，那么以后`Avg`聚合函数执行完成的名字就叫做`avg`。

\4. `aggregate`：这个方法不会返回一个`QuerySet`对象，而是返回一个字典。这个字典中的key就是聚合函数的名字，值就是聚合函数执行后的结果。

\5. `aggregate`和`annotate`的相同和不同：

​    \* 相同：这两个方法都可以执行聚合函数。

​    \* 不同：

​        \- `aggregate`返回的是一个字典，在这个字典中存储的是这个聚合函数执行的结果。而`annotate`返回的是一个`QuerySet`对象，并且会在查找的模型上添加一个聚合函数的属性。

​        \- `aggregate`不会做分组，而`annotate`会使用`group by`子句进行分组，只有调用了`group by`子句，才能对每一条数据求聚合函数的值。

 

\6. `Count`：用来求某个数据的个数。比如要求所有图书的数量，那么可以使用以下代码：

​    \```python

​    result = Book.objects.aggregate(book_nums=Count("id"))

​    \```

​    并且`Count`可以传递`distinct=True`参数，用来剔除那些重复的值，只保留一个。比如要获取作者表中，不同邮箱的个数，那么这时候可以使用`distinct=True`。示例代码如下：

​    \```python

​    result = Author.objects.aggregate(email_nums=Count('email',distinct=True))

​    \```

 

\7. `Max`和`Min`：求指定字段的最大值和最小值。示例代码如下：

​    \```python

​    result = Author.objects.aggregate(max=Max("age"),min=Min("age"))

​    \```

 

\8. `Sum`：求某个字段值的总和。示例代码如下：

​    \```python

​    result = BookOrder.objects.aggregate(total=Sum('price'))

​    \```

​    `aggregate`和`annotate`方法可以在任何的`QuerySet`对象上调用。因此只要是返回了`QuerySet`对象，那么就可以进行链式调用。比如要获取2018年度的销售总额，那么可以先过滤年份，再求聚合函数。示例代码如下：

​    \```python

​    BookOrder.objects.filter(create_time__year=2018).aggregate(total=Sum('price'))

​    \```

 

\7. `F表达式`： 动态的获取某个字段上的值。并且这个F表达式，不会真正的去数据库中查询数据，他相当于只是起一个标识的作用。比如想要将原来每本图书的价格都在原来的基础之上增加10元，那么可以使用以下代码来实现：

​    \```python

​    from django.db.models import F

​    Book.objects.update(price=F("price")+10)

​    \```

 

\8. `Q表达式`：使用`Q`表达式包裹查询条件，可以在条件之间进行多种操作。与/或非等，从而实现一些复杂的查询操作。例子如下：

​    \* 查找价格大于100，并且评分达到4.85以上的图书：

​        \```python

​        \# 不使用Q表达式的

​        books = Book.objects.filter(price__gte=100,rating__gte=4.85)

​        \# 使用Q表达式的

​        books = Book.objects.filter(Q(price__gte=100)&Q(rating__gte=4.85))

​        \```

​    \* 查找价格低于100元，或者评分低于4分的图书：

​        \```python

​        books = Book.objects.filter(Q(price__gte=100)&Q(rating__gte=4.85))

​        \```

​    \* 获取价格大于100，并且图书名字中不包含”传“字的图书：

​        \```python

​        books = Book.objects.filter(Q(price__gte=100)&~Q(name__icontains='传'))

​        \```

 

 

## **orm_queryset_demo**

\# QuerySet API：

 

\## 模型.objects：

这个对象是`django.db.models.manager.Manager`的对象，这个类是一个空壳类，他上面的所有方法都是从`QuerySet`这个类上面拷贝过来的。因此我们只要学会了`QuerySet`，这个`objects`也就知道该如何使用了。

`Manager`源码解析：

\```python

class_name = "BaseManagerFromQuerySet"

 

class_dict = {

​    '_queryset_class': QuerySet

}

class_dict.update(cls._get_queryset_methods(QuerySet))

\# type动态的时候创建类

\# 第一个参数是用来指定创建的类的名字。创建的类名是：BaseManagerFromQuerySet

\# 第二个参数是用来指定这个类的父类。

\# 第三个参数是用来指定这个类的一些属性和方法

return type(class_name,(cls,),class_dict)

 

_get_queryset_methods：这个方法就是将QuerySet中的一些方法拷贝出来

\```

\## filter/exclude/annotate：过滤/排除满足条件的/给模型添加新的字段。

 

\## order_by：

\```python

\# 根据创建的时间正序排序

articles = Article.objects.order_by("create_time")

\# 根据创建的时间倒序排序

articles = Article.objects.order_by("-create_time")

\# 根据作者的名字进行排序

articles = Article.objects.order_by("author__name")

\# 首先根据创建的时间进行排序，如果时间相同，则根据作者的名字进行排序

articles = Article.objects.order_by("create_time",'author__name')

\```

 

一定要注意的一点是，多个`order_by`，会把前面排序的规则给打乱，而使用后面的排序方式。比如以下代码：

 

\```python

articles = Article.objects.order_by("create_time").order_by("author__name")

\```

 

他会根据作者的名字进行排序，而不是使用文章的创建时间。

当然，也可以在模型定义的在`Meta`类中定义`ordering`来指定默认的排序方式。示例代码如下：

\```python

​    class Meta:

​        db_table = 'book_order'

​        ordering = ['create_time','-price']

\```

 

还可以根据`annotate`定义的字段进行排序。比如要实现图书的销量进行排序，那么示例代码如下：

\```python

books = Book.objects.annotate(order_nums=Count("bookorder")).order_by("-order_nums")

​    for book in books:

​        print('%s/%s'%(book.name,book.order_nums))

\```

 

QuerySet API：

我们通常做查询操作的时候，都是通过模型名字.objects的方式进行操作。其实模型名字.objects是一个django.db.models.manager.Manager对象，而Manager这个类是一个“空壳”的类，他本身是没有任何的属性和方法的。他的方法全部都是通过Python动态添加的方式，从QuerySet类中拷贝过来的。示例图如下：
![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps1.png)
所以我们如果想要学习ORM模型的查找操作，必须首先要学会QuerySet上的一些API的使用。

返回新的QuerySet的方法：

在使用QuerySet进行查找操作的时候，可以提供多种操作。比如过滤完后还要根据某个字段进行排序，那么这一系列的操作我们可以通过一个非常流畅的链式调用的方式进行。比如要从文章表中获取标题为123，并且提取后要将结果根据发布的时间进行排序，那么可以使用以下方式来完成：

articles = Article.objects.filter(title='123').order_by('create_time')

可以看到order_by方法是直接在filter执行后调用的。这说明filter返回的对象是一个拥有order_by方法的对象。而这个对象正是一个新的QuerySet对象。因此可以使用order_by方法。

那么以下将介绍在那些会返回新的QuerySet对象的方法。

 

filter：将满足条件的数据提取出来，返回一个新的QuerySet。具体的filter可以提供什么条件查询。请见查询操作章节。

 

exclude：排除满足条件的数据，返回一个新的QuerySet。示例代码如下：

 

 Article.objects.exclude(title__contains='hello')

 

以上代码的意思是提取那些标题不包含hello的图书。

 

annotate：给QuerySet中的每个对象都添加一个使用查询表达式（聚合函数、F表达式、Q表达式、Func表达式等）的新字段。示例代码如下：

 

 articles = Article.objects.annotate(author_name=F("author__name"))

 

以上代码将在每个对象中都添加一个author__name的字段，用来显示这个文章的作者的年龄。

 

order_by：指定将查询的结果根据某个字段进行排序。如果要倒叙排序，那么可以在这个字段的前面加一个负号。示例代码如下：

 

 \# 根据创建的时间正序排序

 articles = Article.objects.order_by("create_time")

 \# 根据创建的时间倒序排序

 articles = Article.objects.order_by("-create_time")

 \# 根据作者的名字进行排序

 articles = Article.objects.order_by("author__name")

 \# 首先根据创建的时间进行排序，如果时间相同，则根据作者的名字进行排序

 articles = Article.objects.order_by("create_time",'author__name')

 

一定要注意的一点是，多个order_by，会把前面排序的规则给打乱，而使用后面的排序方式。比如以下代码：

 

 articles = Article.objects.order_by("create_time").order_by("author__name")

 

他会根据作者的名字进行排序，而不是使用文章的创建时间。

 

 

values：用来指定在提取数据出来，需要提取哪些字段。默认情况下会把表中所有的字段全部都提取出来，可以使用values来进行指定，并且使用了values方法后，提取出的QuerySet中的数据类型不是模型，而是在values方法中指定的字段和值形成的字典：

 

 articles = Article.objects.values("title",'content')

 for article in articles:

​     print(article)

 

以上打印出来的article是类似于{"title":"abc","content":"xxx"}的形式。
如果在values中没有传递任何参数，那么将会返回这个恶模型中所有的属性。

values_list：类似于values。只不过返回的QuerySet中，存储的不是字典，而是元组。示例代码如下：

 articles = Article.objects.values_list("id","title")

 print(articles)

那么在打印articles后，结果为<ySQueret [(1,'abc'),(2,'xxx'),...]>等。
如果在values_list中只有一个字段。那么你可以传递flat=True来将结果扁平化。示例代码如下：

 articles1 = Article.objects.values_list("title")

 \>> <QuerySet [("abc",),("xxx",),...]>

 articles2 = Article.objects.values_list("title",flat=True)

 \>> <QuerySet ["abc",'xxx',...]>

all：获取这个ORM模型的QuerySet对象。

 

select_related：在提取某个模型的数据的同时，也提前将相关联的数据提取出来。比如提取文章数据，可以使用select_related将author信息提取出来，以后再次使用article.author的时候就不需要再次去访问数据库了。可以减少数据库查询的次数。示例代码如下：

 

 article = Article.objects.get(pk=1)

 \>> article.author # 重新执行一次查询语句

 article = Article.objects.select_related("author").get(pk=2)

 \>> article.author # 不需要重新执行查询语句了

 

select_related只能用在一对多或者一对一中，不能用在多对多或者多对一中。比如可以提前获取文章的作者，但是不能通过作者获取这个作者的文章，或者是通过某篇文章获取这个文章所有的标签。

 

prefetch_related：这个方法和select_related非常的类似，就是在访问多个表中的数据的时候，减少查询的次数。这个方法是为了解决多对一和多对多的关系的查询问题。比如要获取标题中带有hello字符串的文章以及他的所有标签，示例代码如下：

 from django.db import connection

 articles = Article.objects.prefetch_related("tag_set").filter(title__contains='hello')

 print(articles.query) # 通过这条命令查看在底层的SQL语句

 for article in articles:

​     print("title:",article.title)

​     print(article.tag_set.all())

 

 \# 通过以下代码可以看出以上代码执行的sql语句

 for sql in connection.queries:

​     print(sql)

 

但是如果在使用article.tag_set的时候，如果又创建了一个新的QuerySet那么会把之前的SQL优化给破坏掉。比如以下代码：

 tags = Tag.obejcts.prefetch_related("articles")

 for tag in tags:

​     articles = tag.articles.filter(title__contains='hello') #因为filter方法会重新生成一个QuerySet，因此会破坏掉之前的sql优化

 

 \# 通过以下代码，我们可以看到在使用了filter的，他的sql查询会更多，而没有使用filter的，只有两次sql查询

 for sql in connection.queries:

​     print(sql)

 

那如果确实是想要在查询的时候指定过滤条件该如何做呢，这时候我们可以使用django.db.models.Prefetch来实现，Prefetch这个可以提前定义好queryset。示例代码如下：

 

 tags = Tag.objects.prefetch_related(Prefetch("articles",queryset=Article.objects.filter(title__contains='hello'))).all()

 for tag in tags:

​     articles = tag.articles.all()

​     for article in articles:

​         print(article)

 

 for sql in connection.queries:

​     print('='*30)

​     print(sql)

 

因为使用了Prefetch，即使在查询文章的时候使用了filter，也只会发生两次查询操作。

 

defer：在一些表中，可能存在很多的字段，但是一些字段的数据量可能是比较庞大的，而此时你又不需要，比如我们在获取文章列表的时候，文章的内容我们是不需要的，因此这时候我们就可以使用defer来过滤掉一些字段。这个字段跟values有点类似，只不过defer返回的不是字典，而是模型。示例代码如下：

 

articles = list(Article.objects.defer("title"))for sql in connection.queries:

​    print('='*30)

​    print(sql)

 

在看以上代码的sql语句，你就可以看到，查找文章的字段，除了title，其他字段都查找出来了。当然，你也可以使用article.title来获取这个文章的标题，但是会重新执行一个查询的语句。示例代码如下：

 

articles = list(Article.objects.defer("title"))for article in articles:

​    \# 因为在上面提取的时候过滤了title

​    \# 这个地方重新获取title，将重新向数据库中进行一次查找操作

​    print(article.title)for sql in connection.queries:

​    print('='*30)

​    print(sql)

 

defer虽然能过滤字段，但是有些字段是不能过滤的，比如id，即使你过滤了，也会提取出来。

 

only：跟defer类似，只不过defer是过滤掉指定的字段，而only是只提取指定的字段。

 

get：获取满足条件的数据。这个函数只能返回一条数据，并且如果给的条件有多条数据，那么这个方法会抛出MultipleObjectsReturned错误，如果给的条件没有任何数据，那么就会抛出DoesNotExit错误。所以这个方法在获取数据的只能，只能有且只有一条。

 

create：创建一条数据，并且保存到数据库中。这个方法相当于先用指定的模型创建一个对象，然后再调用这个对象的save方法。示例代码如下：

 

article = Article(title='abc')

article.save()

\# 下面这行代码相当于以上两行代码

article = Article.objects.create(title='abc')

 

get_or_create：根据某个条件进行查找，如果找到了那么就返回这条数据，如果没有查找到，那么就创建一个。示例代码如下：

obj,created= Category.objects.get_or_create(title='默认分类')

 

如果有标题等于默认分类的分类，那么就会查找出来，如果没有，则会创建并且存储到数据库中。
这个方法的返回值是一个元组，元组的第一个参数obj是这个对象，第二个参数created代表是否创建的。

 

bulk_create：一次性创建多个数据。示例代码如下：

 

Tag.objects.bulk_create([

​    Tag(name='111'),

​    Tag(name='222'),

])

 

count：获取提取的数据的个数。如果想要知道总共有多少条数据，那么建议使用count，而不是使用len(articles)这种。因为count在底层是使用select count(*)来实现的，这种方式比使用len函数更加的高效。

first和last：返回QuerySet中的第一条和最后一条数据。

aggregate：使用聚合函数。

exists：判断某个条件的数据是否存在。如果要判断某个条件的元素是否存在，那么建议使用exists，这比使用count或者直接判断QuerySet更有效得多。示例代码如下：

if Article.objects.filter(title__contains='hello').exists():

​    print(True)

比使用count更高效：if Article.objects.filter(title__contains='hello').count() > 0:

​    print(True)

也比直接判断QuerySet更高效：if Article.objects.filter(title__contains='hello'):

​    print(True)

distinct：去除掉那些重复的数据。这个方法如果底层数据库用的是MySQL，那么不能传递任何的参数。比如想要提取所有销售的价格超过80元的图书，并且删掉那些重复的，那么可以使用distinct来帮我们实现，示例代码如下：

books = Book.objects.filter(bookorder__price__gte=80).distinct()

需要注意的是，如果在distinct之前使用了order_by，那么因为order_by会提取order_by中指定的字段，因此再使用distinct就会根据多个字段来进行唯一化，所以就不会把那些重复的数据删掉。示例代码如下：

orders = BookOrder.objects.order_by("create_time").values("book_id").distinct()

那么以上代码因为使用了order_by，即使使用了distinct，也会把重复的book_id提取出来。

update：执行更新操作，在SQL底层走的也是update命令。比如要将所有category为空的article的article字段都更新为默认的分类。示例代码如下：

Article.objects.filter(category__isnull=True).update(category_id=3)

 

注意这个方法走的是更新的逻辑。所以更新完成后保存到数据库中不会执行save方法，因此不会更新auto_now设置的字段。

 

delete：删除所有满足条件的数据。删除数据的时候，要注意on_delete指定的处理方式。

 

切片操作：有时候我们查找数据，有可能只需要其中的一部分。那么这时候可以使用切片操作来帮我们完成。QuerySet使用切片操作就跟列表使用切片操作是一样的。示例代码如下：

books = Book.objects.all()[1:3]

for book in books:

​    print(book)

 

切片操作并不是把所有数据从数据库中提取出来再做切片操作。而是在数据库层面使用LIMIE和OFFSET来帮我们完成。所以如果只需要取其中一部分的数据的时候，建议大家使用切片操作。

 

什么时候Django会将QuerySet转换为SQL去执行：

生成一个QuerySet对象并不会马上转换为SQL语句去执行。
比如我们获取Book表下所有的图书：

books = Book.objects.all()

print(connection.queries)

我们可以看到在打印connection.quries的时候打印的是一个空的列表。说明上面的QuerySet并没有真正的执行。
在以下情况下QuerySet会被转换为SQL语句执行：

 

迭代：在遍历QuerySet对象的时候，会首先先执行这个SQL语句，然后再把这个结果返回进行迭代。比如以下代码就会转换为SQL语句：

 

 for book in Book.objects.all():

​     print(book)

使用步长做切片操作：QuerySet可以类似于列表一样做切片操作。做切片操作本身不会执行SQL语句，但是如果如果在做切片操作的时候提供了步长，那么就会立马执行SQL语句。需要注意的是，做切片后不能再执行filter方法，否则会报错。

调用len函数：调用len函数用来获取QuerySet中总共有多少条数据也会执行SQL语句

调用list函数：调用list函数用来将一个QuerySet对象转换为list对象也会立马执行SQL语句。

判断：如果对某个QuerySet进行判断，也会立马执行SQL语句。

## **orm_migrations_demo**

\# migrate怎么判断哪些迁移脚本需要执行：

他会将代码中的迁移脚本和数据库中`django_migrations`中的迁移脚本进行对比，如果发现数据库中，没有这个迁移脚本，那么就会执行这个迁移脚本。

 

\# migrate做了什么事情：

\1. 将相关的迁移脚本翻译成SQL语句，在数据库中执行这个SQL语句。

\2. 如果这个SQL语句执行没有问题，那么就会将这个迁移脚本的名字记录到`django_migrations`中。

 

 

\# 执行migrate命令的时候报错的解决办法：

 

\## 原因：

执行migrate命令会报错的原因是。数据库的`django_migrations`表中的迁移版本记录和代码中的迁移脚本不一致导致的。

 

\## 解决办法：

 

\### 使用--fake参数：

首先对比数据库中的迁移脚本和代码中的迁移脚本。然后找到哪个不同，之后再使用`--fake`，将代码中的迁移脚本添加到`django_migrations`中，但是并不会执行sql语句。这样就可以避免每次执行`migrate`的时候，都执行一些重复的迁移脚本。

 

\### 终极解决方案：

如果代码中的迁移脚本和数据库中的迁移脚本实在太多，就是搞不清了。那么这时候就可以使用以下终极解决方案：

\1. 终极解决方案原理：就是将之前的那些迁移脚本都不用了。重新来过。要将出问题的app下的所有模型和数据库中表保持一致，重新映射。

\2. 将出问题的app下的所有模型，都和数据库中的表保持一致。

\3. 将出问题的app下的所有迁移脚本文件都删掉。再在`django_migrations`表中将出问题的app相关的迁移记录都删掉。

\4. 使用`makemigrations`，重新将模型生成一个迁移脚本。

\5. 使用`migrate --fake-initial`参数，将刚刚生成的迁移脚本，标记为已经完成（因为这些模型相对应的表，其实都已经在数据库中存在了，不需要重复执行了。）

\6. 可以做其他的映射了。

 

根据已有的表自动生成模型：

在实际开发中，有些时候可能数据库已经存在了。如果我们用Django来开发一个网站，读取的是之前已经存在的数据库中的数据。那么该如何将模型与数据库中的表映射呢？根据旧的数据库生成对应的ORM模型，需要以下几个步骤：

 

Django给我们提供了一个inspectdb的命令，可以非常方便的将已经存在的表，自动的生成模型。想要使用inspectdb自动将表生成模型。首先需要在settings.py中配置好数据库相关信息。不然就找不到数据库。示例代码如下：

 

 DATABASES = {

​     'default': {

​         'ENGINE': 'django.db.backends.mysql',

​         'NAME': "migrations_demo",

​         'HOST': '127.0.0.1',

​         'PORT': '3306',

​         'USER': 'root',

​         'PASSWORD': 'root'

​     }

 }

 

比如有以下表：

article表：

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps2.jpg) 

 

tag表：

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps3.jpg) 

article_tag表：

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps4.jpg) 

front_user表：

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps5.jpg) 

那么通过python manage.py inspectdb，就会将表转换为模型后的代码，显示在终端：

 

from django.db import models

class ArticleArticle(models.Model):

  title = models.CharField(max_length=100)

  content = models.TextField(blank=True, null=True)

  create_time = models.DateTimeField(blank=True, null=True)

  author = models.ForeignKey('FrontUserFrontuser', models.DO_NOTHING, blank=True, null=True)

 

  class Meta:

​      managed = False

​      db_table = 'article_article'

class ArticleArticleTags(models.Model):

  article = models.ForeignKey(ArticleArticle, models.DO_NOTHING)

  tag = models.ForeignKey('ArticleTag', models.DO_NOTHING)

 

  class Meta:

​      managed = False

​      db_table = 'article_article_tags'

​      unique_together = (('article', 'tag'),)

class ArticleTag(models.Model):

  name = models.CharField(max_length=100)

 

  class Meta:

​      managed = False

​      db_table = 'article_tag'

class FrontUserFrontuser(models.Model):

  username = models.CharField(max_length=100)

  telephone = models.CharField(max_length=11)

 

  class Meta:

​      managed = False

​      db_table = 'front_user_frontuser'

 

以上代码只是显示在终端。如果想要保存到文件中。那么可以使用>重定向输出到指定的文件。比如让他输出到models.py文件中。示例命令如下：

 

python manage.py inspectdb > models.py

 

以上的命令，只能在终端执行，不能在pycharm->Tools->Run manage.py Task...中使用。

 

如果只是想要转换一个表为模型。那么可以指定表的名字。示例命令如下：

 

python manage.py inspectdb article_article > models.py

 

 

修正模型：新生成的ORM模型有些地方可能不太适合使用。比如模型的名字，表之间的关系等等。那么以下选项还需要重新配置一下：

 

模型名：自动生成的模型，是根据表的名字生成的，可能不是你想要的。这时候模型的名字你可以改成任何你想要的。

模型所属app：根据自己的需要，将相应的模型放在对应的app中。放在同一个app中也是没有任何问题的。只是不方便管理。

模型外键引用：将所有使用ForeignKey的地方，模型引用都改成字符串。这样不会产生模型顺序的问题。另外，如果引用的模型已经移动到其他的app中了，那么还要加上这个app的前缀。

让Django管理模型：将Meta下的managed=False删掉，如果保留这个，那么以后这个模型有任何的修改，使用migrate都不会映射到数据库中。

 

当有多对多的时候，应该也要修正模型。将中间表注视了，然后使用ManyToManyField来实现多对多。并且，使用ManyToManyField生成的中间表的名字可能和数据库中那个中间表的名字不一致，这时候肯定就不能正常连接了。那么可以通过db_table来指定中间表的名字。示例代码如下：

 

class Article(models.Model):

 title = models.CharField(max_length=100, blank=True, null=True)

 content = models.TextField(blank=True, null=True)

 author = models.ForeignKey('front.User', models.SET_NULL, blank=True, null=True)

 \# 使用ManyToManyField模型到表，生成的中间表的规则是：article_tags

 \# 但现在已经存在的表的名字叫做：article_tag

 \# 可以使用db_table，指定中间表的名字

 tags = models.ManyToManyField("Tag",db_table='article_tag')

 

 class Meta:

​     db_table = 'article'

 

表名：切记不要修改表的名字。不然映射到数据库中，会发生找不到对应表的错误。

 

执行命令python manage.py makemigrations生成初始化的迁移脚本。方便后面通过ORM来管理表。这时候还需要执行命令python manage.py migrate --fake-initial，因为如果不使用--fake-initial，那么会将迁移脚本会映射到数据库中。这时候迁移脚本会新创建表，而这个表之前是已经存在了的，所以肯定会报错。此时我们只要将这个0001-initial的状态修改为已经映射，而不真正执行映射，下次再migrate的时候，就会忽略他。

 

 

将Django的核心表映射到数据库中：Django中还有一些核心的表也是需要创建的。不然有些功能是用不了的。比如auth相关表。如果这个数据库之前就是使用Django开发的，那么这些表就已经存在了。可以不用管了。如果之前这个数据库不是使用Django开发的，那么应该使用migrate命令将Django中的核心模型映射到数据库中。

# **Chapter05**

## **method_decorator_demo装饰器**

Django限制请求method

常用的请求method：

GET请求：GET请求一般用来向服务器索取数据，但不会向服务器提交数据，不会对服务器的状态进行更改。比如向服务器获取某篇文章的详情。

POST请求：POST请求一般是用来向服务器提交数据，会对服务器的状态进行更改。比如提交一篇文章给服务器。

限制请求装饰器：

Django内置的视图装饰器可以给视图提供一些限制。比如这个视图只能通过GET的method访问等。以下将介绍一些常用的内置视图装饰器。

 

django.http.decorators.http.require_http_methods：这个装饰器需要传递一个允许访问的方法的列表。比如只能通过GET的方式访问。那么示例代码如下：

 

 from django.views.decorators.http import require_http_methods

 @require_http_methods(['POST','GET'])
*def* index2(*request*):
    *if* *request*.method=='GET':
        *return* render(*request*,'add_article.html')
    *else*:
        title = *request*.POST.get('title')
        content = *request*.POST.get('content')
        Article.objects.create(title=title,content=content)
        *return* HttpResponse("index1")

 

 

django.views.decorators.http.require_GET：这个装饰器相当于是require_http_methods(['GET'])的简写形式，只允许使用GET的method来访问视图。示例代码如下：

 

 from django.views.decorators.http import require_GET

 @require_GET

  *def* index1(*request*):
    articles = Article.objects.all()
    *return* render(*request*,'index.html',context={"articles":articles})

 

django.views.decorators.http.require_POST：这个装饰器相当于是require_http_methods(['POST'])的简写形式，只允许使用POST的method来访问视图。示例代码如下：

 

 from django.views.decorators.http import require_POST

 @require_POST

 *def* index2(*request*):
    title = *request*.POST.get('title')
    content = *request*.POST.get('content')
    Article.objects.create(title=title,content=content)
    *return* HttpResponse("index1")

 

 

 

django.views.decorators.http.require_safe：这个装饰器相当于是require_http_methods(['GET','HEAD'])的简写形式，只允许使用相对安全的方式来访问视图。因为GET和HEAD不会对服务器产生增删改的行为。因此是一种相对安全的请求方式。示例代码如下：

 

 from django.views.decorators.http import require_safe

 @require_safe

 def my_view(request):

​     pass

 

## **redirect_demo重定向**

重定向分为永久性重定向和暂时性重定向，在页面上体现的操作就是浏览器会从一个页面自动跳转到另外一个页面。比如用户访问了一个需要权限的页面，但是该用户当前并没有登录，因此我们应该给他重定向到登录页面。

 

永久性重定向：http的状态码是301，多用于旧网址被废弃了要转到一个新的网址确保用户的访问，最经典的就是京东网站，你输入www.jingdong.com的时候，会被重定向到www.jd.com，因为jingdong.com这个网址已经被废弃了，被改成jd.com，所以这种情况下应该用永久重定向。

 

 

暂时性重定向：http的状态码是302，表示页面的暂时性跳转。比如访问一个需要权限的网址，如果当前用户没有登录，应该重定向到登录页面，这种情况下，应该用暂时性重定向。

 

在Django中，重定向是使用redirect(to, *args, permanent=False, **kwargs)来实现的。to是一个url，permanent代表的是这个重定向是否是一个永久的重定向，默认是False。关于重定向的使用。请看以下例子：

*from* django.urls *import* path
*from* front *import* views
urlpatterns = [
    path('', views.index,name='index'),
    path('signup/', views.signup,name='signup'),
]

 

*def* index(*request*):
    username = *request*.GET.get("username")
    *if* username:
        *return* HttpResponse("首页")
    *else*:
        *return* redirect(reverse("signup"))

*def* signup(*request*):
    *return* HttpResponse("注册页面")

 

 

 

## **WSGIRequest对象**

Django在接收到http请求之后，会根据http请求携带的参数以及报文信息创建一个WSGIRequest对象，并且作为视图函数第一个参数传给视图函数。也就是我们经常看到的request参数。在这个对象上我们可以找到客户端上传上来的所有信息。这个对象的完整路径是django.core.handlers.wsgi.WSGIRequest。

WSGIRequest对象常用属性和方法：

WSGIRequest对象常用属性：

WSGIRequest对象上大部分的属性都是只读的。因为这些属性是从客户端上传上来的，没必要做任何的修改。以下将对一些常用的属性进行讲解：

path：请求服务器的完整“路径”，但不包含域名和参数。比如http://www.baidu.com/xxx/yyy/，那么path就是/xxx/yyy/。

method：代表当前请求的http方法。比如是GET还是POST。

GET：一个django.http.request.QueryDict对象。操作起来类似于字典。这个属性中包含了所有以?xxx=xxx的方式上传上来的参数。

POST：也是一个django.http.request.QueryDict对象。这个属性中包含了所有以POST方式上传上来的参数。

FILES：也是一个django.http.request.QueryDict对象。这个属性中包含了所有上传的文件。

COOKIES：一个标准的Python字典，包含所有的cookie，键值对都是字符串类型。

session：一个类似于字典的对象。用来操作服务器的session。

 

META：存储的客户端发送上来的所有header信息。

 

 

CONTENT_LENGTH：请求的正文的长度（是一个字符串）。

 

CONTENT_TYPE：请求的正文的MIME类型。

HTTP_ACCEPT：响应可接收的Content-Type。

HTTP_ACCEPT_ENCODING：响应可接收的编码。

HTTP_ACCEPT_LANGUAGE： 响应可接收的语言。

HTTP_HOST：客户端发送的HOST值。

HTTP_REFERER：在访问这个页面上一个页面的url。

QUERY_STRING：单个字符串形式的查询字符串（未解析过的形式）。

REMOTE_ADDR：客户端的IP地址。如果服务器使用了nginx做反向代理或者负载均衡，那么这个值返回的是127.0.0.1，这时候可以使用HTTP_X_FORWARDED_FOR来获取，所以获取ip地址的代码片段如下：

  if request.META.has_key('HTTP_X_FORWARDED_FOR'):  

​      ip =  request.META['HTTP_X_FORWARDED_FOR']  

  else:  

​      ip = request.META['REMOTE_ADDR']

 

REMOTE_HOST：客户端的主机名。

REQUEST_METHOD：请求方法。一个字符串类似于GET或者POST。

SERVER_NAME：服务器域名。

SERVER_PORT：服务器端口号，是一个字符串类型。

WSGIRequest对象常用方法：

is_secure()：是否是采用https协议。

is_ajax()：是否采用ajax发送的请求。原理就是判断请求头中是否存在X-Requested-With:XMLHttpRequest。

get_host()：服务器的域名。如果在访问的时候还有端口号，那么会加上端口号。比如www.baidu.com:9000。

get_full_path()：返回完整的path。如果有查询字符串，还会加上查询字符串。比如/music/bands/?print=True。

get_raw_uri()：获取请求的完整url。

## **QueryDict对象：**

我们平时用的request.GET和request.POST都是QueryDict对象，这个对象继承自dict，因此用法跟dict相差无几。其中用得比较多的是get方法和getlist方法。

get方法：用来获取指定key的值，如果没有这个key，那么会返回None,如果要是有默认值就显示默认值。

@require_http_methods(['GET'])
*def* index1(*request*):
    username = *request*.GET.get('username',default=1)
    print(username)
    *return* HttpResponse("index1")

 

getlist方法：如果浏览器上传上来的key对应的值有多个，那么就需要通过这个方法获取。

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps6.png)![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps7.png)<**label**>
    Python<**input** type=**"checkbox"** name=**"tags"** value=**"Python"**/>
</**label**>
<**label**>
    Django<**input** type=**"checkbox"** name=**"tags"** value=**"Django"**/>
</**label**>

 

@require_http_methods(['GET','POST'])
*def* index2(*request*):
    *if* *request*.method=='GET':
        *return* render(*request*, 'index2.html')
    *else*:
        title = *request*.POST.get('title')
        content = *request*.POST.get('cgetlist(ontent')
        tags = *request*.POST.'tags')
        print("title:",title,"content:",content,"tags:",tags)
    *return* HttpResponse("index2")

 

 

 

 

## **HttpResponse对象**

Django服务器接收到客户端发送过来的请求后，会将提交上来的这些数据封装成一个HttpRequest对象传给视图函数。那么视图函数在处理完相关的逻辑后，也需要返回一个响应给浏览器。而这个响应，我们必须返回HttpResponseBase或者他的子类的对象。而HttpResponse则是HttpResponseBase用得最多的子类。那么接下来就来介绍一下HttpResponse及其子类。

常用属性：

content：返回的内容。

status_code：返回的HTTP响应状态码。

content_type：返回的数据的MIME类型，默认为text/html。浏览器会根据这个属性，来显示数据。如果是text/html，那么就会解析这个字符串，如果text/plain，那么就会显示一个纯文本。常用的Content-Type如下：

text/html（默认的，html文件）

text/plain（纯文本）

text/css（css文件）

text/javascript（js文件）

multipart/form-data（文件提交）

application/json（json传输）

application/xml（xml文件）

设置请求头：response['X-Access-Token'] = 'xxxx'。

常用方法：

set_cookie：用来设置cookie信息。后面讲到授权的时候会着重讲到。

delete_cookie：用来删除cookie信息。

write：HttpResponse是一个类似于文件的对象，可以用来写入数据到数据体（content）中。

*def* index3(*request*):
    response = HttpResponse("<h1>孙亚博</h1>",content_type='text/plain;charset=utf-8')
    response['X-XXX-SUN']='YABO'#设置请求头
    response.write("<h5>你是谁</h5>")
    print(response.status_code)
    *return* response

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps8.jpg) 

## **JsonResponse类：**

用来对象dump成json字符串，然后返回将json字符串封装成Response对象返回给浏览器。并且他的Content-Type是application/json。示例代码如下：

from django.http import JsonResponsedef index(request):

​    return JsonResponse({"username":"zhiliao","age":18})

默认情况下JsonResponse只能对字典进行dump，如果想要对非字典的数据进行dump，那么需要给JsonResponse传递一个safe=False参数。示例代码如下：

from django.http import JsonResponsedef index(request):

​    persons = ['张三','李四','王五']

​    return HttpResponse(persons)

以上代码会报错，应该在使用HttpResponse的时候，传入一个safe=False参数，示例代码如下：

return HttpResponse(persons,safe=False)

 

 

## **生成CSV文件：**

有时候我们做的网站，需要将一些数据，生成有一个CSV文件给浏览器，并且是作为附件的形式下载下来。以下将讲解如何生成CSV文件。

生成小的CSV文件：

这里将用一个生成小的CSV文件为例，来把生成CSV文件的技术要点讲到位。我们用Python内置的csv模块来处理csv文件，并且使用HttpResponse来将csv文件返回回去。示例代码如下：

*from* django.http *import* HttpResponse,JsonResponse,StreamingHttpResponse

*def* csv_view(*request*):
    response = HttpResponse(content_type='text/csv')
    response['Content-Disposition'] = 'attachment;filename="abc.csv"'
    writer = csv.writer(response)
    writer.writerow(['username','age'])
    writer.writerow(['zhiliao',18])
    *return* response

这里再来对每个部分的代码进行解释：

我们在初始化HttpResponse的时候，指定了Content-Type为text/csv，这将告诉浏览器，这是一个csv格式的文件而不是一个HTML格式的文件，如果用默认值，默认值就是html，那么浏览器将把csv格式的文件按照html格式输出，这肯定不是我们想要的。

第二个我们还在response中添加一个Content-Disposition头，这个东西是用来告诉浏览器该如何处理这个文件，我们给这个头的值设置为attachment;，那么浏览器将不会对这个文件进行显示，而是作为附件的形式下载，第二个filename="somefilename.csv"是用来指定这个csv文件的名字。

我们使用csv模块的writer方法，将相应的数据写入到response中。

将csv文件定义成模板：

我们还可以将csv格式的文件定义成模板，然后使用Django内置的模板系统，并给这个模板传入一个Context对象，这样模板系统就会根据传入的Context对象，生成具体的csv文件。示例代码如下：

模板文件：

{% for row in rows %}{{ row.0 }},{{ row.1 }}
{% endfor %}

视图函数：

from django.http import HttpResponsefrom django.template import loader, Context

*def* template_csv_view(*request*):
    response = HttpResponse(content_type='text/csv')
    response['Content-Disposition'] = 'attachment;filename="abc.csv"'
    content = {
        'rows': [
            ['username','age'],
            ['孙亚博',18],
        ]
    }
    template = loader.get_template('a.txt')
    csv_template = template.render(content)
    response.content = csv_template
    *return* response

 

生成大的CSV文件：

这是将整体的数据返回回来

*def* large_csv_view1(*request*):
    response = HttpResponse(content_type='text/csv')
    response['Content-Disposition'] = 'attachment;filename="larges.csv"'
    rows = ("row {},{}\n".format(row,row) *for* row *in* range(0,10000000))
    response.streaming_content = rows
    *return* response

 

以上的例子是生成的一个小的csv文件，如果想要生成大型的csv文件，那么以上方式将有可能会发生超时的情况（服务器要生成一个大型csv文件，需要的时间可能会超过浏览器默认的超时时间）。这时候我们可以借助另外一个类，叫做StreamingHttpResponse对象，这个对象是将响应的数据作为一个流返回给客户端，而不是作为一个整体返回。示例代码如下：

这里我们构建了一个非常大的数据集rows，并且将其变成一个迭代器。然后因为StreamingHttpResponse的第一个参数只能是一个生成器，因此我们使用圆括号(writer.writerow(row) for row in rows)，并且因为我们要写的文件是csv格式的文件，因此需要调用writer.writerow将row变成一个csv格式的字符串。而调用writer.writerow又需要一个中间的容器，因此这里我们定义了一个非常简单的类Echo，这个类只实现一个write方法，以后在执行csv.writer(pseudo_buffer)的时候，就会调用Echo.writer方法。
注意：StreamingHttpResponse会启动一个进程来和客户端保持长连接，所以会很消耗资源。所以如果不是特殊要求，尽量少用这种方法。

关于StreamingHttpResponse：

这个类是专门用来处理流数据的。使得在处理一些大型文件的时候，不会因为服务器处理时间过长而到时连接超时。这个类不是继承自HttpResponse，并且跟HttpResponse对比有以下几点区别：

这个类没有属性content，相反是streaming_content。

这个类的streaming_content必须是一个可以迭代的对象。

这个类没有write方法，如果给这个类的对象写入数据将会报错。

注意：StreamingHttpResponse会启动一个进程来和客户端保持长连接，所以会很消耗资源。所以如果不是特殊要求，尽量少用这种方法。

 

*def* large_csv_view(*request*):
    response = HttpResponse(content_type='text/csv')
    response['Content-Disposition'] = 'attachment;filename="large.csv"'
    write = csv.writer(response)
    *for* row *in* range(10000000):
        write.writerow(["row{}".format(row),"row{}".format(row)])
    *return* response

 

 

## **类视图**

在写视图的时候，Django除了使用函数作为视图，也可以使用类作为视图。使用类视图可以使用类的一些特性，比如继承等。

View：

django.views.generic.base.View是主要的类视图，所有的类视图都是继承自他。如果我们写自己的类视图，也可以继承自他。然后再根据当前请求的method，来实现不同的方法。比如这个视图只能使用get的方式来请求，那么就可以在这个类中定义get(self,request,*args,**kwargs)方法。以此类推，如果只需要实现post方法，那么就只需要在类中实现post(self,request,*args,**kwargs)。示例代码如下：

from django.views import Viewclass BookDetailView(View):

​    def get(self,request,*args,**kwargs):

​        return render(request,'detail.html')

类视图写完后，还应该在urls.py中进行映射，映射的时候就需要调用View的类方法as_view()来进行转换。示例代码如下：

urlpatterns = [        

​    path("detail/<book_id>/",views.BookList.as_view(),name='detail')

]

*class* BookList(View):
    *def* get(self,request,*book_id*):
        print("我的图书id：%s" % *book_id*)
        *return* HttpResponse("success")

 

除了get方法，View还支持以下方法['get','post','put','patch','delete','head','options','trace']。

如果用户访问了View中没有定义的方法。比如你的类视图只支持get方法，而出现了post方法，那么就会把这个请求转发给http_method_not_allowed(request,*args,**kwargs)。示例代码如下：

*class* BookList1(View):
    *def* get(self,request,*args,**kwargs):
        *return* HttpResponse("success")
    *def* http_method_not_allowed(self, *request*, **args*, ***kwargs*):
        *return* HttpResponse("您当前采用的method是：%s，本视图只支持"
                            "使用post请求！" % *request*.method)

urls.py中的映射如下：

path("addbook/",views.AddBookView.as_view(),name='add_book')

如果你在浏览器中访问addbook/，因为浏览器访问采用的是get方法，而addbook只支持post方法，因此以上视图会返回您当前采用的method是：GET，本视图只支持使用post请求！。

其实不管是get请求还是post请求，都会走dispatch(request,*args,**kwargs)方法，并且dispatch先执行，所以如果实现这个方法，将能够对所有请求都处理到。

*class* AddBook(View):
    *def* get(self,*request*,*args,**kwargs):
        *return* render(*request*,'index.html')
    *def* dispatch(self, *request*, **args*, ***kwargs*):
        print("我是最牛逼的")
        *return* super(AddBook,self).dispatch(*request*,**args*,***kwargs*)
    *def* post(self,*request*,*args,**kwargs):
        xuehao = *request*.POST.get('xuehao')
        name = *request*.POST.get('name')
        print("学号:{}   姓名:{}".format(xuehao,name))
        *return* HttpResponse("success")

我是最牛逼的

学号:111   姓名:111

 

## **TemplateView：**

django.views.generic.base.TemplateView，这个类视图是专门用来返回模版的。在这个类中，有两个属性是经常需要用到的，一个是template_name，这个属性是用来存储模版的路径，TemplateView会自动的渲染这个变量指向的模版。另外一个是get_context_data，这个方法是用来返回上下文数据的，也就是在给模版传的参数的。示例代码如下：

from django.views.generic.base import TemplateView

*class* AboutView(TemplateView):
    template_name = 'about.html'
    *def* get_context_data(self, ***kwargs*):
        content = {
            'name':'牛逼'
        }
        *return* content  #这里是将content返回到about.html中

在urls.py中的映射代码如下：

from django.urls import path

from myapp.views import HomePageView

urlpatterns = [

​    path('', HomePageView.as_view(), name='home'),

]

如果在模版中不需要传递任何参数，那么可以直接只在urls.py中使用TemplateView来渲染模版。示例代码如下：

from django.urls import pathfrom django.views.generic import TemplateView

urlpatterns = [

​    path('about/', TemplateView.as_view(template_name="about.html")),

]

 

## **ListView：**

在网站开发中，经常会出现需要列出某个表中的一些数据作为列表展示出来。比如文章列表，图书列表等等。在Django中可以使用ListView来帮我们快速实现这种需求。示例代码如下：

*class* AddArticle(ListView):
    model = Article
    template_name = 'add.html'
    paginate_by = 8
    context_object_name = 'articles'
    ordering = 'create_time'
    page_kwarg = 'sun'
    *def* get_context_data(self, ***kwargs*):
        context = super(AddArticle,self).get_context_data(***kwargs*)
        print(context)
        print(context)
        *return* context
    *def* get_queryset(self):
        *return* Article.objects.filter(id__gte=80)

对以上代码进行解释：

首先AddArticle是继承自ListView。

如果不写get_queryset默认把所有的文章都提取出来

model：重写model类属性，指定这个列表是给哪个模型的。

template_name：指定这个列表的模板。

paginate_by：指定这个列表一页中展示多少条数据。

context_object_name：指定这个列表模型在模板中的参数名称。

ordering：指定这个列表的排序方式。

page_kwarg：获取第几页的数据的参数名称。默认是page。

get_context_data：获取上下文的数据。

get_queryset：如果你提取数据的时候，并不是要把所有数据都返回，那么你可以重写这个方法。将一些不需要展示的数据给过滤掉。

## **Paginator和Page类：**

Paginator和Page类都是用来做分页的。他们在Django中的路径为django.core.paginator.Paginator和django.core.paginator.Page。以下对这两个类的常用属性和方法做解释：

## **Paginator常用属性和方法：**

\1. count：总共有多少条数据。

\2. num_pages：总共有多少页。

\3. page_range：页面的区间。比如有三页，那么就range(1,4)。

## **Page常用属性和方法：**

\1. has_next：是否还有下一页。

\2. has_previous：是否还有上一页。

\3. next_page_number：下一页的页码。

\4. previous_page_number：上一页的页码。

\5. number：当前页。

\6. start_index：当前这一页的第一条数据的索引值。

\7. end_index：当前这一页的最后一条数据的索引值。

*def* get_context_data(self, ***kwargs*):
    context = super(AddArticle,self).get_context_data(***kwargs*)
    paginator = context.get('paginator')
    page_obj = context.get('page_obj')
    print(page_obj.has_next())
    print(paginator.count)
    print(paginator.num_pages)
    print(page_obj.page_range)
    print(page_obj.has_previous())
    print(page_obj.next_page_number())
    print(page_obj.previous_page_number())
    print(page_obj.number)
    print(page_obj.start_index())
    print(page_obj.end_index())
    print(context)
    *return* context

 

 

 

 

 

## **手动实现分页**

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps9.jpg) 

首先是‘上一页’的处理

{% **if** page_obj.has_previous %}
    <**li**><**a** href=**"**{% **url** 'front:add1' %}**?sun=**{{ page_obj.previous_page_number }}**"**>**上一页**</**a**></**li**>
{% **else** %}
    <**li** class=**"disabled"**><**a** href=**"javascript:void(0);"**>**上一页**</**a**></**li**>
{% **endif** %}

首先我们通过page_obj.has_previous 我们通过这个来判断是否还有上一页，如果有上一页的话，通过page_obj.previous_page_number 通过它来跳转到上一页，如果没有上一页的话，就让‘上一页’这个字变为disabled。

其次是中间页

{# 中间页 #}
{% **for** page in paginator.page_range %}
    {% **if** page_obj.number == page %}
        <**li** class=**"active"**><**a** href=**"javascript:void(0);"**>{{ page }}</**a**></**li**>
    {% **else** %}
        <**li**><**a** href=**"**{% **url** 'front:add1' %}**?sun=**{{ page }}**"**>{{ page }}</**a**></**li**>
    {% **endif** %}
{% **endfor** %}

通过paginator.page_range 来确定总共有多少页然后通过判断 page_obj.number == page  是否相等，如果相等就不执行什么操作，如果不相等的话，就会跳转到page页

最后是下一页

{# 下一页 #}
{% **if** page_obj.has_next %}
    <**li**><**a** href=**"**{% **url** 'front:add1' %}**?sun=**{{ page_obj.next_page_number }}**"**>**下一页**</**a**></**li**>
{% **else** %}
    <**li** class=**"disabled"**><**a** href=**"javascript:void(0);"**>**下一页**</**a**></**li**>
{% **endif** %}

 

通过page_obj.has_next 判断是否有下一页，如果有的话跳转到下一页，没有的话就不可点击

 

## **手动实现分页算法**

1.首先我们先定一个方法，并且再让get_context_data进行调用，定义的这个函数名为 

get_paginate_data这里我们需要用到paginator,page_obj,还有around_content是两边显示的个数，通过上一节学习的paginator,page_obj的用法，我们先找到当前页两边页的取值范围

left_pages和right_pages这俩取值范围，然后再通过判断当前页跟around的关系来从新设定left_pages和right_pages这俩取值范围，然后我们从新定义俩变量来确定当前页前面是否

需要加上...，如果需要的话我们再在模板中进行渲染，第一页跟最后一页可以直接写死。

{% **if** left_have_more %}
    <**li**><**a** href=**"**{% **url** 'front:add1' %}**?sun=1"**>**1**</**a**></**li**>
    <**li**><**a** href=**"javascript:void(0);"**>**...**</**a**></**li**>
{% **endif** %}

这里我们所到的东西，必须通过get_paginate_data返回给get_context_data的context然后再给模板

模板页如下

<ul class="pagination">
    {# 上一页 #}
    {% if page_obj.has_previous %}
        <li><a href="{% url 'front:add1' %}?sun={{ page_obj.previous_page_number }}">上一页</a></li>
    {% else %}
        <li class="disabled"><a href="javascript:void(0);">上一页</a></li>
    {% endif %}
    {# 左边页 #}
    {% if left_have_more %}
        <li><a href="{% url 'front:add1' %}?sun=1">1</a></li>
        <li><a href="javascript:void(0);">...</a></li>
    {% endif %}
    {% for page in left_pages %}
        <li><a href="{% url 'front:add1' %}?sun={{ page }}">{{ page }}</a></li>
    {% endfor %}

    {# 当前页 #}
    <li class="active"><a href="{% url 'front:add1' %}?sun={{ current_page }}">{{ current_page }}</a></li>

    {# 右边页 #}
    {% for page in right_pages %}
        <li><a href="{% url 'front:add1' %}?sun={{ page }}">{{ page }}</a></li>
    {% endfor %}

    {% if right_have_more %}
        <li><a href="javascript:void(0);">...</a></li>
        <li><a href="{% url 'front:add1' %}?sun={{ totle_pages }}">{{ totle_pages }}</a></li>
    {% endif %}

    {# 下一页 #}
    {% if page_obj.has_next %}
        <li><a href="{% url 'front:add1' %}?sun={{ page_obj.next_page_number }}">下一页</a></li>
    {% else %}
        <li class="disabled"><a href="javascript:void(0);">下一页</a></li>
    {% endif %}
</ul>

试图页如下

*def* get_context_data(self, ***kwargs*):
    context = super(AddArticle,self).get_context_data(***kwargs*)
    paginator = context.get('paginator')
    page_obj = context.get('page_obj')

​    sun = self.get_paginate_data(paginator,page_obj)
​    context.update(sun)#加入到context
​    *return* context
*def* get_paginate_data(self,*paginator*,*page_obj*,*around_content*=2):
​    current_page = *page_obj*.number
​    totle_pages = *paginator*.num_pages
​    left_have_more = *False*    right_have_more = *False*    left_pages = range(current_page-*around_content*,current_page)
​    right_pages = range(current_page+1,current_page+*around_content*+1)

​    *if* current_page<=*around_content*+2:#为了防止出现0，-1页
​        left_pages = range(1, current_page)
​    *else*:
​        left_have_more = *True*        left_pages = range(current_page - *around_content*, current_page)

​    *if* current_page>=totle_pages-*around_content*-1:#为了防止超过最大页数的数字产生
​        right_pages = range(current_page+1, totle_pages + 1)
​    *else*:
​        right_have_more = *True*        right_pages = range(current_page+1, current_page+*around_content*+1)

​    *return* {
​        'left_pages':left_pages,
​        'right_pages':right_pages,
​        'current_page':current_page,
​        'totle_pages':totle_pages,
​        'left_have_more':left_have_more,
​        'right_have_more':right_have_more,
​    }

 

## **给类视图添加装饰器：**

在开发中，有时候需要给一些视图添加装饰器。如果用函数视图那么非常简单，只要在函数的上面写上装饰器就可以了。但是如果想要给类添加装饰器，那么可以通过以下两种方式来实现：

## **装饰dispatch方法：**

from django.utils.decorators import method_decorator

def login_required(func):

​    def wrapper(request,*args,**kwargs):

​        if request.GET.get("username"):

​            return func(request,*args,**kwargs)

​        else:

​            return redirect(reverse('index'))

​    return wrapper

 

class IndexView(View):

​    def get(self,request,*args,**kwargs):

​        return HttpResponse("index")

​    @method_decorator(login_required)

​    def dispatch(self, request, *args, **kwargs):

​        super(IndexView, self).dispatch(request,*args,**kwargs)

## **错误处理**

在一些网站开发中。经常会需要捕获一些错误，然后将这些错误返回比较优美的界面，或者是将这个错误的请求做一些日志保存。那么我们本节就来讲讲如何实现。

常用的错误码：

404：服务器没有指定的url。

403：没有权限访问相关的数据。

405：请求的method错误。

400：bad request，请求的参数错误。

500：服务器内部错误，一般是代码出bug了。

502：一般部署的时候见得比较多，一般是nginx启动了，然后uwsgi有问题。

自定义错误模板：

在碰到比如404，500错误的时候，想要返回自己定义的模板。那么可以直接在templates文件夹下创建相应错误代码的html模板文件。那么以后在发生相应错误后，会将指定的模板返回回去。 ![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps10.png)

错误处理的解决方案：

对于404和500这种自动抛出的错误。我们可以直接在templates文件夹下新建相应错误代码的模板文件。而对于其他的错误，我们可以专门定义一个app，用来处理这些错误。 ![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps11.png)

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps12.jpg) 

*from* django.urls *import* path
*from* . *import* views
app_name='errors'
urlpatterns = [
    path('405.html',views.view_405,name='405'),
    path('403.html',views.view_403,name='403'),
]

*def* view_405(*request*):
    *return* render(*request*,'errors/405.html',status=405)
*def* view_403(*request*):
    *return* render(*request*,'errors/403.html',status=403)

 

 

 

 

 

 

# **Chapter06**

## **表单：**

HTML中的表单：

单纯从前端的html来说，表单是用来提交数据给服务器的,不管后台的服务器用的是Django还是PHP语言还是其他语言。只要把input标签放在form标签中，然后再添加一个提交按钮，那么以后点击提交按钮，就可以将input标签中对应的值提交给服务器了。

Django中的表单：

Django中的表单丰富了传统的HTML语言中的表单。在Django中的表单，主要做以下两件事：

渲染表单模板。

表单验证数据是否合法。

Django中表单使用流程：

在讲解Django表单的具体每部分的细节之前。我们首先先来看下整体的使用流程。这里以一个做一个留言板为例。首先我们在后台服务器定义一个表单类required=*False* 默认为true是必须传，继承自django.forms.Form。示例代码如下：

*from* django *import* forms
*class* MessageBoardForm(forms.Form):
    title = forms.CharField(max_length=100,min_length=2,label='标题',error_messages={"min_length":'最少不能少于2字符'})
    content = forms.CharField(widget=forms.Textarea,label='内容',error_messages={"required":'必须传递content字符'})
    email = forms.EmailField(label='邮箱',error_messages={"required":'必须传递email字符'})
    reply = forms.BooleanField(required=*False*,label='是否回复')

 

然后在视图中，根据是GET还是POST请求来做相应的操作。如果是GET请求，那么返回一个空的表单，如果是POST请求，那么将提交上来的数据进行校验，因为前面设置好MessageBoardForm的格式，这里会进行自动的检测。示例代码如下：

*from* django.shortcuts *import* render
*from* django.views.generic *import* View
*from* .forms *import* MessageBoardForm
*from* django.http *import* HttpResponse
*class* Index_View(View):
    *def* get(self,*request*):
        form = MessageBoardForm()
        *return* render(*request*,'index.html',context={"form":form})
    *def* post(self,*request*):
        form = MessageBoardForm(*request*.POST)#这里把前台传递数据给form
        *if* form.is_valid():
            title = form.cleaned_data.get('title')
            content = form.cleaned_data.get('content')
            email = form.cleaned_data.get('email')
            reply = form.cleaned_data.get('reply')
            *return* HttpResponse('success')
        *else*:
            print(form.errors.get_json_data())#返回所有的错误信息
            *return* HttpResponse("fail")

 

在使用GET请求的时候，我们传了一个form给模板，那么以后模板就可以使用form来生成一个表单的html代码。在使用POST请求的时候，我们根据前端上传上来的数据，构建一个新的表单，这个表单是用来验证数据是否合法的，如果数据都验证通过了，那么我们可以通过cleaned_data来获取相应的数据。在模板中渲染表单的HTML代码如下：

<form action="" method="post">

    <table>

 

​        <tr>

​            <td></td>

​            <td><input type="submit" value="提交"></td>

​        </tr>

​    </table></form>

我们在最外面给了一个form标签，然后在里面使用了table标签来进行美化，在使用form对象渲染的时候，使用的是table的方式，当然还可以使用ul的方式（as_ul），也可以使用p标签的方式（as_p），并且在后面我们还加上了一个提交按钮。这样就可以生成一个表单

 

## **自定义验证：**

有时候对一个字段验证，不是一个长度，一个正则表达式能够写清楚的，还需要一些其他复杂的逻辑，那么我们可以对某个字段，进行自定义的验证。比如在注册的表单验证中，我们想要验证手机号码是否已经被注册过了，那么这时候就需要在数据库中进行判断才知道。对某个字段进行自定义的验证方式是，定义一个方法，这个方法的名字定义规则是：clean_fieldname。如果验证失败，那么就抛出一个验证错误。比如要验证用户表中手机号码之前是否在数据库中存在，那么可以通过以下代码实现：

*from* django *import* forms
*from* django.core *import* validators
*from* .models *import* User

*class* Register(forms.Form):
    username = forms.CharField(max_length=100,min_length=2,label='用户名',error_messages={'min_length':'用户名最少2字符'})
    telephone = forms.CharField(label='电话', validators=[validators.RegexValidator(r'1[345678]\d{9}', message="请输入一个正确为手机号")])
    pwd = forms.CharField(max_length=100)
    pwd1 = forms.CharField(max_length=100)
    *def* clean_telephone(self):
        telphone = self.cleaned_data.get('telephone')
        exits = User.objects.filter(telphone=telphone).exists()
        *if* exits:
            \#return HttpResponse("该手机号已经被注册")
            *raise* forms.ValidationError("该手机号已经被注册")
        *else*: 	
            *return* telphone

以上是对某个字段进行验证，如果验证数据的时候，需要针对多个字段进行验证，那么可以重写clean方法。比如要在注册的时候，要判断提交的两个密码是否相等。那么可以使用以下代码来完成：

*class* Register(forms.Form):
    username = forms.CharField(max_length=100,min_length=2,label='用户名',error_messages={'min_length':'用户名最少2字符'})
    telephone = forms.CharField(label='电话', validators=[validators.RegexValidator(r'1[345678]\d{9}', message="请输入一个正确为手机号")])
    pwd = forms.CharField(max_length=100)
    pwd1 = forms.CharField(max_length=100)
    *def* clean(self):
        pwd = self.cleaned_data.get('pwd')
        pwd1 = self.cleaned_data.get('pwd1')
        *if* pwd!=pwd1:
            *raise* forms.ValidationError("俩次输入密码不一致")

提取错误信息：

如果验证失败了，那么有一些错误信息是我们需要传给前端的。这时候我们可以通过以下属性来获取：

form.errors：这个属性获取的错误信息是一个包含了html标签的错误信息。

form.errors.get_json_data()：这个方法获取到的是一个字典类型的错误信息。将某个字段的名字作为key，错误信息作为值的一个字典。

form.as_json()：这个方法是将form.get_json_data()返回的字典dump成json格式的字符串，方便进行传输。

上述方法获取的字段的错误值，都是一个比较复杂的数据。比如以下：

{'username': [{'message': 'Enter a valid URL.', 'code': 'invalid'}, {'message': 'Ensure this value has at most 4 characters (it has 22).', 'code': 'max_length'}]}

那么如果我只想把错误信息放在一个列表中，而不要再放在一个字典中。这时候我们可以定义一个方法，把这个数据重新整理一份。实例代码如下：

class MyForm(forms.Form):

​    username = forms.URLField(max_length=4)

​    def get_errors(self):

​        errors = self.errors.get_json_data()

​        new_errors = {}

​        for key,message_dicts in errors.items():

​            messages = []

​            for message in message_dicts:

​                messages.append(message['message'])

​            new_errors[key] = messages

​        return new_errors

这样就可以把某个字段所有的错误信息直接放在这个列表中。

## **提取错误信息：**

如果验证失败了，那么有一些错误信息是我们需要传给前端的。这时候我们可以通过以下属性来获取：

form.errors：这个属性获取的错误信息是一个包含了html标签的错误信息。

form.errors.get_json_data()：这个方法获取到的是一个字典类型的错误信息。将某个字段的名字作为key，错误信息作为值的一个字典。

form.as_json()：这个方法是将form.get_json_data()返回的字典dump成json格式的字符串，方便进行传输。

上述方法获取的字段的错误值，都是一个比较复杂的数据。比如以下：

{'username': [{'message': 'Enter a valid URL.', 'code': 'invalid'}, {'message': 'Ensure this value has at most 4 characters (it has 22).', 'code': 'max_length'}]}

那么如果我只想把错误信息放在一个列表中，而不要再放在一个字典中。这时候我们可以定义一个方法，把这个数据重新整理一份。实例代码如下：

class MyForm(forms.Form):

​    username = forms.URLField(max_length=4)

​    def get_errors(self):

​        errors = self.errors.get_json_data()

​        new_errors = {}

​        for key,message_dicts in errors.items():

​            messages = []

​            for message in message_dicts:

​                messages.append(message['message'])

​            new_errors[key] = messages

​        return new_errors

这样就可以把某个字段所有的错误信息直接放在这个列表中。为了能够让更多的类方法调用get_errors我们可以定义一个类，让其他使用get_errors方法的类去继承有get_errors的类

## **ModelForm：**

大家在写表单的时候，会发现表单中的Field和模型中的Field基本上是一模一样的，而且表单中需要验证的数据，也就是我们模型中需要保存的。那么这时候我们就可以将模型中的字段和表单中的字段进行绑定。
比如现在有个Article的模型。示例代码如下：

from django.db import models

rom django.core import validatorsclass 

Article(models.Model):

​    title = models.CharField(max_length=10,validators=[validators.MinLengthValidator(limit_value=3)])

​    content = models.TextField()

​    author = models.CharField(max_length=100)

​    category = models.CharField(max_length=100)

​    create_time = models.DateTimeField(auto_now_add=True)

那么在写表单的时候，就不需要把Article模型中所有的字段都一个个重复写一遍了。示例代码如下：

from django import formsclass MyForm(forms.ModelForm):

​    class Meta:

​        model = Article

​        fields = "__all__"

MyForm是继承自forms.ModelForm，然后在表单中定义了一个Meta类，在Meta类中指定了model=Article，以及fields="__all__"，这样就可以将Article模型中所有的字段都复制过来，进行验证。如果只想针对其中几个字段进行验证，那么可以给fields指定一个列表，将需要的字段写进去。比如只想验证title和content，那么可以使用以下代码实现：

from django import formsclass MyForm(forms.ModelForm):

​    class Meta:

​        model = Article

​        fields = ['title','content']

如果要验证的字段比较多，只是除了少数几个字段不需要验证，那么可以使用exclude来代替fields。比如我不想验证category，那么示例代码如下：

class MyForm(forms.ModelForm):

​    class Meta:

​        model = Article

​        exclude = ['category']

## **save方法：**

ModelForm还有save方法，可以在验证完成后直接调用save方法，就可以将这个数据保存到数据库中了，这是必须保证fields=’__all__’ 否者的话会报错 。示例代码如下：

form = MyForm(request.POST)

if form.is_valid():

​    	form.save()

return HttpResponse('succes')

else:

​    	print(form.get_errors())

​    	return HttpResponse('fail')

这个方法必须要在clean没有问题后才能使用，如果在clean之前使用，会抛出异常。另外，我们在调用save方法的时候，如果传入一个commit=False，那么只会生成这个模型的对象，而不会把这个对象真正的插入到数据库中。比如表单上验证的字段没有包含模型中所有的字段，这时候就可以先创建对象，再根据填充其他字段，把所有字段的值都补充完成后，再保存到数据库中。示例代码如下：

Models中

*class* User(models.Model):
    username = models.CharField(max_length=100)
    password = models.CharField(max_length=16)
    telephone = models.CharField(max_length=12,validators=[validators.RegexValidator(r'1[345678]\d{9}')])

Forms中

*class* Register(forms.ModelForm):
    pwd1 = forms.CharField(max_length=16,min_length=6)
    pwd2 = forms.CharField(max_length=16,min_length=6)
    *class* Meta:
        model = User
        fields = ['username','telephone']
    *def* get_error(self):
        form = self.errors.get_json_data()
        new_error = {}
        *for* key,messages_dict *in* form.items():
            messages = []
            *for* message *in* messages_dict:
                messages.append(message['message'])
            new_error[key] = messages
        *return* new_error

​    *def* clean(self):
​        clean_data = super().clean()
​        pwd1 = clean_data.get('pwd1')
​        pwd2 = clean_data.get('pwd2')
​        *if* pwd1!=pwd2:
​            *raise* forms.ValidationError("俩次密码输入不一样")
​        *else*:
​            *return* clean_data

Views中

*from* django.views.decorators.http *import* require_POST

@require_POST
*def* register(*request*):
    form = Register(*request*.POST)
    *if* form.is_valid():
        user = form.save(commit=*False*)
        user.password = form.cleaned_data.get('pwd1')
        user.save()
        *return* HttpResponse("success")
    *else*:
        print(form.get_error())
        *return* HttpResponse("false")

## **文件上传：**

文件上传是网站开发中非常常见的功能。这里详细讲述如何在Django中实现文件的上传功能。

前端HTML代码实现：

在前端中，我们需要填入一个form标签，然后在这个form标签中指定enctype="multipart/form-data"，不然就不能上传文件。

在form标签中添加一个input标签，然后指定input标签的name，以及type="file"。

以上两步的示例代码如下：

<form action="" method="post" enctype="multipart/form-data">

<input type="file" name="myfile">

</form>

后端的代码实现：

后端的主要工作是接收文件。然后存储文件。接收文件的方式跟接收POST的方式是一样的，只不过是通过FILES来实现。示例代码如下：

def save_file(file):

​    with open('somefile.txt','wb') as fp:

​        for chunk in file.chunks():

​            fp.write(chunk)

def index(request):

​    if request.method == 'GET':

​        form = MyForm()

​        return render(request,'index.html',{'form':form})

​    else:

​        myfile = request.FILES.get('myfile')

​        save_file(myfile)

​        return HttpResponse('success')

以上代码通过request.FILES接收到文件后，再写入到指定的地方。这样就可以完成一个文件的上传功能了。

 

## **使用模型来处理上传的文件：**

在定义模型的时候，我们可以给存储文件的字段指定为FileField，这个Field可以传递一个upload_to参数，用来指定上传上来的文件保存到哪里。比如我们让他保存到项目的files文件夹下，那么示例代码如下：

\# models.pyclass Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

​    thumbnail = models.FileField(upload_to="files")

 

\# views.py

def index(request):

​    if request.method == 'GET':

​        return render(request,'index.html')

​    else:

​        title = request.POST.get('title')

​        content = request.POST.get('content')

​        thumbnail = request.FILES.get('thumbnail') #别忘了是FILES

​        article = Article(title=title, content=content, thumbnail=thumbnail)

​        article.save()

​        return HttpResponse('success')

调用完article.save()方法，就会把文件保存到files下面，并且会将这个文件的路径存储到数据库中。

## **指定MEDIA_ROOT和MEDIA_URL：**

以上我们是使用了upload_to来指定上传的文件的目录。我们也可以指定MEDIA_ROOT，就不需要在FielField中指定upload_to，他会自动的将文件上传到MEDIA_ROOT的目录下。

MEDIA_ROOT = os.path.join(BASE_DIR,'media')

MEDIA_URL = '/media/'

然后我们可以在urls.py中添加MEDIA_ROOT目录下的访问路径。示例代码如下：

from django.urls import path

from front import views

from django.conf.urls.static import static

from django.conf import settings

urlpatterns = [

​    path('', views.index),

] + static(settings.MEDIA_URL,document_root=settings.MEDIA_ROOT)

如果我们同时指定MEDIA_ROOT和upload_to，那么会将文件上传到MEDIA_ROOT下的upload_to文件夹中。示例代码如下：

class Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

thumbnail = models.FileField(upload_to="%Y/%m/%d/")#生成如下图的格式

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps13.jpg) 

*from* django.urls *import* path
*from* front *import* views
*from* django.conf.urls.static *import* static
*from* django.conf *import* settings
urlpatterns = [
    path('',views.Filess)
]+ static(settings.MEDIA_URL,document_root=settings.MEDIA_ROOT)

这样写可以访问上传的文件

## **限制上传的文件拓展名：**

如果想要限制上传的文件的拓展名，那么我们就需要用到表单来进行限制。我们可以使用普通的Form表单，也可以使用ModelForm，直接从模型中读取字段。示例代码如下：

\# models.pyclass Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

​    thumbnial = models.FileField(upload_to='%Y/%m/%d/',validators=[validators.FileExtensionValidator(['txt','pdf'])])

\# forms.pyclass ArticleForm(forms.ModelForm):

​    class Meta:

​        model = Article

​        fields = "__all__"

上传图片：

上传图片跟上传普通文件是一样的。只不过是上传图片的时候Django会判断上传的文件是否是图片的格式（除了判断后缀名，还会判断是否是可用的图片）。如果不是，那么就会验证失败。我们首先先来定义一个包含ImageField的模型。示例代码如下：

class Article(models.Model):

​    title = models.CharField(max_length=100)

​    content = models.TextField()

​    thumbnail = models.ImageField(upload_to="%Y/%m/%d/")

因为要验证是否是合格的图片，因此我们还需要用一个表单来进行验证。表单我们直接就使用ModelForm就可以了。示例代码如下：

class MyForm(forms.ModelForm):

​    class Meta:

​        model = Article

​        fields = "__all__"

注意：使用ImageField，必须要先安装Pillow库：pip install pillow

## **memcached**

什么是memcached：

memcached之前是danga的一个项目，最早是为LiveJournal服务的，当初设计师为了加速LiveJournal访问速度而开发的，后来被很多大型项目采用。官网是www.danga.com或者是memcached.org。

Memcached是一个高性能的分布式的内存对象缓存系统，全世界有不少公司采用这个缓存项目来构建大负载的网站，来分担数据库的压力。Memcached是通过在内存里维护一个统一的巨大的hash表，memcached能存储各种各样的数据，包括图像、视频、文件、以及数据库检索的结果等。简单的说就是将数据调用到内存中，然后从内存中读取，从而大大提高读取速度。

哪些情况下适合使用Memcached：存储验证码（图形验证码、短信验证码）、登录session等所有不是至关重要的数据。

安装和启动memcached：

windows：

安装：memcached.exe -d install。

启动：memcached.exe -d start。

linux（ubuntu）：

安装：sudo apt install memcached

启动：

cd /usr/local/memcached/bin

./memcached -d start

 

可能出现的问题：

提示你没有权限：在打开cmd的时候，右键使用管理员身份运行。

提示缺少pthreadGC2.dll文件：将pthreadGC2.dll文件拷贝到windows/System32.

不要放在含有中文的路径下面。

 

启动memcached：

 

-d：这个参数是让memcached在后台运行。

-m：指定占用多少内存。以M为单位，默认为64M。

-p：指定占用的端口。默认端口是11211。

-l：别的机器可以通过哪个ip地址连接到我这台服务器。如果是通过service memcached start的方式，那么只能通过本机连接。如果想要让别的机器连接，就必须设置-l 0.0.0.0。

如果想要使用以上参数来指定一些配置信息，那么不能使用service memcached start，而应该使用/usr/bin/memcached的方式来运行。比如/usr/bin/memcached -u memcache -m 1024 -p 11222 start。

telnet操作memcached：

telnet ip地址 [11211]

添加数据：

set：

语法：

  set key flas(是否压缩) timeout value_length

  value

示例：

  set username 0 60 7

  zhiliao

add：

语法：

  add key flas(0) timeout value_length

  value

示例：

  add username 0 60 7

  xiaotuo

 

set和add的区别：add是只负责添加数据，不会去修改数据。如果添加的数据的key已经存在了，则添加失败，如果添加的key不存在，则添加成功。而set不同，如果memcached中不存在相同的key，则进行添加，如果存在，则替换。

 

获取数据：

语法：

  get key

示例：

  get username

删除数据：

语法：

  delete key

示例：

  delete username

flush_all：删除memcached中的所有数据。

 

查看memcached的当前状态：

 

语法：stats。

通过python操作memcached：

安装：python-memcached：pip install python-memcached。

建立连接：

 import memcache

 mc = memcache.Client(['127.0.0.1:11211','192.168.174.130:11211'],debug=True)#可以用多ip

设置数据：

 mc.set('username','hello world',time=60*5)

 mc.set_multi({'email':'xxx@qq.com','telphone':'111111'},time=60*5)

获取数据：

 mc.get('telphone')

删除数据：

 mc.delete('email')

自增长：

 mc.incr('read_count')

自减少：

 mc.decr('read_count')

*import* memcache
mc = memcache.Client(['127.0.0.1:11211'],debug=*True*)
sun = mc.get('username')
mc.set_multi({'author':'sunyabo','price':123,'age':18},time=1000)
print(sun)                     #None
print(mc.get('price'))	 #123	
print(mc.get('author'))	 #sunyabo
print(mc.get('age'))		 #18
mc.delete('username')	 
sun = mc.get('username')
mc.incr('age')
print(mc.get('age'))		#19
print(sun)			#None
print(mc.decr('price'))	#122

 

memcached的安全性：

memcached的操作不需要任何用户名和密码，只需要知道memcached服务器的ip地址和端口号即可。因此memcached使用的时候尤其要注意他的安全性。这里提供两种安全的解决方案。分别来进行讲解：

使用-l参数设置为只有本地可以连接：这种方式，就只能通过本机才能连接，别的机器都不能访问，可以达到最好的安全性。

使用防火墙，关闭11211端口，外面也不能访问。

  ufw enable # 开启防火墙

  ufw disable # 关闭防火墙

  ufw default deny # 防火墙以禁止的方式打开，默认是关闭那些没有开启的端口

  ufw deny 端口号 # 关闭某个端口

  ufw allow 端口号 # 开启某个端口

在Django中使用memcached：

首先需要在settings.py中配置好缓存：

CACHES = {

​    'default': {

​        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',

​        'LOCATION': '127.0.0.1:11211',

​    }

}

如果想要使用多台机器，那么可以在LOCATION指定多个连接，示例代码如下：

CACHES = {

​    'default': {

​        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',

​        'LOCATION': [

​            '172.19.26.240:11211',

​            '172.19.26.242:11211',

​        ]

​    }

}

配置好memcached的缓存后，以后在代码中就可以使用以下代码来操作memcached了：

如何查看所有的memcached中的所有的key

\1. stats items

\2. Stats  cachedump  [items_id]  0[代表查看所有的key]

*from* django.http *import* HttpResponse
*from* django.core.cache *import* cache
*from* django.core.cache.backends.memcached *import* MemcachedCache
*def* index(*request*):
    cache.set('passWord','123',1000)
    an = cache.get('abc')
    print(an)
    *return* HttpResponse("success")

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps14.jpg) 

 

需要注意的是，django在存储数据到memcached中的时候，不会将指定的key存储进去，而是会对key进行一些处理。比如会加一个前缀，会加一个版本号。如果想要自己加前缀，那么可以在settings.CACHES中添加KEY_FUNCTION参数：

CACHES = {

​    'default': {

​        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',

​        'LOCATION': '127.0.0.1:11211',

​        'KEY_FUNCTION': lambda key,prefix_key,version:"django:%s"%key 设置指定的key

​    }

}

 

# **Chapter07**

## **cookie和session**

cookie：在网站中，http请求是无状态的。也就是说即使第一次和服务器连接后并且登录成功后，第二次请求服务器依然不能知道当前请求是哪个用户。cookie的出现就是为了解决这个问题，第一次登录后服务器返回一些数据（cookie）给浏览器，然后浏览器保存在本地，当该用户发送第二次请求的时候，就会自动的把上次请求存储的cookie数据自动的携带给服务器，服务器通过浏览器携带的数据就能判断当前用户是哪个了。cookie存储的数据量有限，不同的浏览器有不同的存储大小，但一般不超过4KB。因此使用cookie只能存储一些小量的数据。

 

session: session和cookie的作用有点类似，都是为了存储用户相关的信息。不同的是，cookie是存储在本地浏览器，session是一个思路、一个概念、一个服务器存储授权信息的解决方案，不同的服务器，不同的框架，不同的语言有不同的实现。虽然实现不一样，但是他们的目的都是服务器为了方便存储数据的。session的出现，是为了解决cookie存储数据不安全的问题的。

 

 

cookie和session使用：web开发发展至今，cookie和session的使用已经出现了一些非常成熟的方案。在如今的市场或者企业里，一般有两种存储方式：

 

存储在服务端：通过cookie存储一个sessionid，然后具体的数据则是保存在session中。如果用户已经登录，则服务器会在cookie中保存一个sessionid，下次再次请求的时候，会把该sessionid携带上来，服务器根据sessionid在session库中获取用户的session数据。就能知道该用户到底是谁，以及之前保存的一些状态信息。这种专业术语叫做server side session。Django把session信息默认存储到数据库中，当然也可以存储到其他地方，比如缓存中，文件系统中等。存储在服务器的数据会更加的安全，不容易被窃取。但存储在服务器也有一定的弊端，就是会占用服务器的资源，但现在服务器已经发展至今，一些session信息还是绰绰有余的。

将session数据加密，然后存储在cookie中。这种专业术语叫做client side session。flask框架默认采用的就是这种方式，但是也可以替换成其他形式。

在django中操作cookie和session：

操作cookie：

设置cookie：

设置cookie是设置值给浏览器的。因此我们需要通过response的对象来设置，设置cookie可以通过response.set_cookie来设置，这个方法的相关参数如下：

key：这个cookie的key。

value：这个cookie的value。

max_age：最长的生命周期。单位是秒。

expires：过期时间。跟max_age是类似的，只不过这个参数需要传递一个具体的日期，比如datetime或者是符合日期格式的字符串。如果同时设置了expires和max_age，那么将会使用expires的值作为过期时间。

path：对域名下哪个路径有效。默认是对域名下所有路径都有效。

domain：针对哪个域名有效。默认是针对主域名下都有效，如果只要针对某个子域名才有效，那么可以设置这个属性.

secure：是否是安全的，如果设置为True，那么只能在https协议下才可用。

httponly：默认是False。如果为True，那么在客户端不能通过JavaScript进行操作。

删除cookie：

通过delete_cookie即可删除cookie。实际上删除cookie就是将指定的cookie的值设置为空的字符串，然后使用将他的过期时间设置为0，也就是浏览器关闭后就过期。

 

如果path=/cms/时，我们访问时前面必须有cms，后面可以随便的添加

获取cookie：

获取浏览器发送过来的cookie信息。可以通过request.COOKIES来或者。这个对象是一个字典类型。比如获取所有的cookie，那么示例代码如下：

cookies = request.COOKIES

for cookie_key,cookie_value in cookies.items():

   print(cookie_key,cookie_value)

 

操作session：

django中的session默认情况下是存储在服务器的数据库中的，在表中会根据sessionid来提取指定的session数据，然后再把这个sessionid放到cookie中发送给浏览器存储，浏览器下次在向服务器发送请求的时候会自动的把所有cookie信息都发送给服务器，服务器再从cookie中获取sessionid，然后再从数据库中获取session数据。但是我们在操作session的时候，这些细节压根就不用管。我们只需要通过request.session即可操作。示例代码如下：

def index(request):

   request.session.get('username')

   return HttpResponse('index')

session常用的方法如下：

get：用来从session中获取指定值。

pop：从session中删除一个值。

keys：从session中获取所有的键。

items：从session中获取所有的值。

clear：清除当前这个用户的session数据。

flush：删除session并且删除在浏览器中存储的session_id，一般在注销的时候用得比较多。

set_expiry(value)：设置过期时间。

*def* session_view(*request*):
    *request*.session['username']='sunyabo'
    *request*.session['sex'] = 'nan'
    sex = *request*.session.get('username')
    print(sex)
    name = *request*.session.pop('sex')
    print(name)
    name = *request*.session.get('sex')
    print(name)
    *request*.session.flush()
    sessions = *request*.session.items()
    *for* key,value *in* sessions:
        print(key,value)
    \# for key in request.session.keys():
    \#     print(key)
    *request*.session.clear_expired()
    *return* HttpResponse("success")

整形：代表秒数，表示多少秒后过期。

0：代表只要浏览器关闭，session就会过期。

None：会使用全局的session配置。在settings.py中可以设置SESSION_COOKIE_AGE来配置全局的过期时间。默认是1209600秒，也就是2周的时间。

 

clear_expired：清除过期的session。Django并不会清除过期的session，需要定期手动的清理，或者是在终端，使用命令行python manage.py clearsessions来清除过期的session。

 

修改session的存储机制：

默认情况下，session数据是存储到数据库中的。当然也可以将session数据存储到其他地方。可以通过设置SESSION_ENGINE来更改session的存储位置，这个可以配置为以下几种方案：

django.contrib.sessions.backends.db：使用数据库。默认就是这种方案。

django.contrib.sessions.backends.file：使用文件来存储session。

django.contrib.sessions.backends.cache：使用缓存来存储session。想要将数据存储到缓存中，前提是你必须要在settings.py中配置好CACHES，并且是需要使用Memcached，而不能使用纯内存作为缓存。

django.contrib.sessions.backends.cached_db：在存储数据的时候，会将数据先存到缓存中，再存到数据库中。这样就可以保证万一缓存系统出现问题，session数据也不会丢失。在获取数据的时候，会先从缓存中获取，如果缓存中没有，那么就会从数据库中获取。

首先在setting中配置

CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': '127.0.0.1:11211',
    }
}
SESSION_ENGINE = 'django.contrib.sessions.backends.cached_db'

 

*def* session1(*request*):
    *request*.session["user"] = '孙亚博'
    user = *request*.session.get('user')#从缓存中获取
    print(user)
    *return* HttpResponse("success")

 

django.contrib.sessions.backends.signed_cookies：将session信息加密后存储到浏览器的cookie中。这种方式要注意安全，建议设置SESSION_COOKIE_HTTPONLY=True，那么在浏览器中不能通过js来操作session数据，并且还需要对settings.py中的SECRET_KEY进行保密，因为一旦别人知道这个SECRET_KEY，那么就可以进行解密。另外还有就是在cookie中，存储的数据不能超过4k。

## **上下文处理器**

上下文处理器是可以返回一些数据，在全局模板中都可以使用。比如登录后的用户信息，在很多页面中都需要使用，那么我们可以放在上下文处理器中，就没有必要在每个视图函数中都返回这个对象。

在settings.TEMPLATES.OPTIONS.context_processors中，有许多内置的上下文处理器。这些上下文处理器的作用如下：

django.template.context_processors.debug：增加一个debug和sql_queries变量。在模板中可以通过他来查看到一些数据库查询。在setting中设置如下

INTERNAL_IPS = ['127.0.0.1']

首页{{ request.path }},{{ debug }},{{ sql_queries }}#设置完这样才能显示出来

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps15.jpg) 

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps16.jpg) 

django.template.context_processors.request：增加一个request变量。这个request变量也就是在视图函数的第一个参数。

django.contrib.auth.context_processors.auth：Django有内置的用户系统，这个上下文处理器会增加一个user对象。

django.contrib.messages.context_processors.messages：增加一个messages变量。可以显示出一些错误的信息如下图

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps17.jpg) 

解决方案就是通过上下文处理器来进行处理，首先是在TEMPLATES中的OPTIONS中添加

'django.contrib.messages.context_processors.messages',

这样我们就可以使用messages了，首先是将错误的信息传入到messages中

errors = form.get_errors()
\#print(errors)
*for* error *in* errors:
    messages.info(*request*,error)

​    \#messages.add_message(request,messages.INFO,error)#第二个参数是的级别

将错误信息传进去就可以在模板中进行使用了，这里名字写messages

{% **for** foo in messages %}
    <**li**>{{ foo }}</**li**>
{% **endfor** %}

django.template.context_processors.media：在模板中可以读取MEDIA_URL。比如想要在模板中使用上传的文件，那么这时候就需要使用settings.py中设置的MEDIA_URL来拼接url，并且不要忘了映射到url中。示例代码如下：

MEDIA_ROOT = os.path.join(BASE_DIR,'media')
MEDIA_URL = '/media/'

在这里我们创建一个media文件夹，通过上面的设置我们上传的文件默认传入到media文件中如果想现实图片我们可以通过以下方法

<img src=" {{ MEDIA_URL }}11.png" alt="">这样是为了防止写死方便以后维护

*from* django.conf.urls.static *import* static
*from* django.conf *import* settings

urlpatterns = [
    path('', views.index,name='index'),
    path('signin/', views.SigninView.as_view(),name='signin'),
    path('signup/', views.SignupView.as_view(),name='signup'),
    path('blog/', views.blog,name='blog'),
    path('video/', views.video,name='video'),
] + static(settings.MEDIA_URL,document_root=settings.MEDIA_ROOT)

django.template.context_processors.static：在模板中可以使用STATIC_URL。

django.template.context_processors.csrf：在模板中可以使用csrf_token变量来生成一个csrf token。我们有俩种方式  

方法一:

通过将上下文处理csrf加入到context_processors，然后按照如下写一个input,然后再使用csrf_token,这样就可以提交csrf

<**input** type=**"hidden"** name=**"csrfmiddlewaretoken"** value=**"**{{ csrf_token }}**"**>

'django.template.context_processors.csrf',

方法二:

什么都不添加直接使用{% csrf_token %}这个标签,当我们查看页面源代码时，我们就会看到{% csrf_token %}自动生成一个input(跟上面的input一样)

但是我们有时候想要在head中产生一个csrf，所以我们不希望生成一个input，所以我们不能用{% csrf_token %}这个标签，应该用{{ csrf_token }}这个变量

{% **block** head %}
    <**meta** name=**"csrf-token"** content=**"**{{ csrf_token }}**"**>
{% **endblock** %}

 

自定义上下文处理器：

有时候我们想要返回自己的数据。那么这时候我们可以自定义上下文处理器。自定义上下文处理器的步骤如下：

你可以根据这个上下文处理器是属于哪个app，然后在这个app中创建一个文件专门用来存储上下文处理器。比如context_processors.py。或者是你也可以专门创建一个Python包，用来存储所有的上下文处理器。

在setting中配置:'front.context_processors.front_user',

在你定义的上下文处理器文件中，定义一个函数，这个函数只有一个request参数。这个函数中处理完自己的逻辑后，把需要返回给模板的数据，通过字典的形式返回。如果不需要返回任何数据，那么也必须返回一个空的字典。示例代码如下：

 *from* .models *import* User
*def* front_user(*request*):
    sun = *request*.session.get('user_id')
    content = {}
    *if* sun:
        *try*:
            user = User.objects.get(pk=sun)
            content['front'] = user#这里的front用在模板中
        *except*:
            *pass*    *return* content

模板中的映射

{% if front %}
    <li><a href="#">{{ front.username }}</a></li>
{% else %}
    <li><a href="{% url 'signup' %}">登录</a></li>
    <li><a href="{% url 'signin' %}">注册</a></li>
{% endif %}

 

## **中间件**

中间件是在request和response处理过程中的一个插件。比如在request到达视图函数之前，我们可以使用中间件来做一些相关的事情，比如可以判断当前这个用户有没有登录，如果登录了，就绑定一个user对象到request上。也可以在response到达浏览器之前，做一些相关的处理，比如想要统一在response上设置一些cookie信息等。

自定义中间件：

中间件所处的位置没有规定。只要是放到项目当中即可。一般分为两种情况，如果中间件是属于某个app的，那么可以在这个app下面创建一个python文件用来存放这个中间件，也可以专门创建一个Python包，用来存放本项目的所有中间件。创建中间件有两种方式，一种是使用函数，一种是使用类，接下来对这两种方式做个介绍：

使用函数的中间件：

*def* front_user_middleware(*get_response*):
    print("这是中间件初始化代码")
    *def* middleware(*request*):
        print("request到达view的执行代码")
        user_id = *request*.session.get('user_id')
        *if* user_id:
            *try*:
                user = User.objects.get(pk=user_id)
                *request*.front_user = user
            *except*:
                *request*.front_user = *None*        response = get_response(*request*)
        print("response到达浏览器的代码")
        *return* response
    *return* middleware

然后还需要在setting中的MIDDLEWARE中进行配置

'front.middlewares.front_user_middleware',

当我们把front_user绑定之后，下次我们使用的时候就可以直接用了

如果要查找登陆的用户名直接用resquest.front_user.username

 

使用类的中间件：

*class* FrontUser(object):
    *def* __init__(self,*get_response*):
        print("这是中间件初始化代码")
        self.get_response = *get_response*    *def* __call__(self, *request*):
        print("request到达view的执行代码")
        user_id = *request*.session.get('user_id')
        *if* user_id:
            *try*:
                user = User.objects.get(pk=user_id)
                *request*.front_user = user
            *except*:
                *request*.front_user = *None*        response = self.get_response(*request*)
        print("response到达浏览器的代码")
        *return* response

在写完中间件后，还需要在settings.MIDDLEWARES中配置写好的中间件才可以使用。比如我们写了一个在request到达视图函数之前，判断这个用户是否登录，如果已经登录就绑定一个对象到request上的中间件，那么就可以在settings.MIDDLEWARES下做以下配置：

MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
    'front.middlewares.front_user_middleware',
    'front.middlewares.FrontUser',
]

中间件的执行是有顺序的，他会根据在MIDDLEWARE中存放的顺序来执行。因此如果有些中间件是需要基于其他中间件的，那么就需要放在其他中间件的后面来执行。

## **Django内置的中间件：**

django.middleware.common.CommonMiddleware：通用中间件。他的作用如下：

限制settings.DISALLOWED_USER_AGENTS中指定的请求头来访问本网站。DISALLOWED_USER_AGENT是一个正则表达式的列表。示例代码如下：

*import* re
DISALLOWED_USER_AGENTS = [
    re.compile(r'^\s$|^$'),#\s代表所有的空白字符
    re.compile(r'.*PhantomJS.*')
]

我们可以对这个做一个测试

*import* requests
headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36'
    \# 'User-Agent': ''不能爬到
}
response = requests.get('http://127.0.0.1:8000/',headers=headers)
print(response.text)

 

1.如果开发者在定义url的时候，最后有一个斜杠。但是用户在访问url的时候没有提交这个斜杠，那么CommonMiddleware会自动的重定向到加了斜杠的url上去。

 

2.django.middleware.gzip.GZipMiddleware：将响应数据进行压缩。如果内容长度少于200个长度，那么就不会压缩。注释与不注释上面的代码的效果

![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps18.jpg)![img](file:///C:\Users\ASUS\AppData\Local\Temp\ksohtml6780\wps19.jpg) 

3 .django.contrib.messages.middleware.MessageMiddleware：消息处理相关的中间件。

\4. django.middleware.security.SecurityMiddleware：做了一些安全处理的中间件。比如设置XSS防御的请求头，比如做了http协议转https协议的工作等。

\5. django.contrib.sessions.middleware.SessionMiddleware：session中间件。会给request添加一个处理好的session对象。

\6. django.contrib.auth.middleware.AuthenticationMiddleware：会给request添加一个user对象的中间件。

\7. django.middleware.csrf.CsrfViewMiddleware：CSRF保护的中间件。

\8. django.middleware.clickjacking.XFrameOptionsMiddleware：做了clickjacking攻击的保护。clickjacking保护是攻击者在自己的病毒网站上，写一个诱惑用户点击的按钮，然后使用iframe的方式将受攻击的网站（比如银行网站）加载到自己的网站上去，并将其设置为透明的，用户就看不到，然后再把受攻击的网站（比如银行网站）的转账按钮定位到病毒网站的按钮上，这样用户在点击病毒网站上按钮的时候，实际上点击的是受攻击的网站（比如银行网站）上的按钮，从而实现了在不知不觉中给攻击者转账的功能。

\9. 缓存中间件：用来缓存一些页面的。

django.middleware.cache.UpdateCacheMiddleware。

django.middleware.cache.FetchFromCacheMiddleware。

## **内置中间件放置的顺序：**

\1. SecurityMiddleware：应该放到最前面。因为这个中间件并不需要依赖任何其他的中间件。如果你的网站同时支持http协议和https协议，并且你想让用户在使用http协议的时候重定向到https协议，那么就没有必要让他执行下面一大串中间件再重定向，这样效率更高。

\2. UpdateCacheMiddleware：应该在SessionMiddleware, GZipMiddleware, LocaleMiddleware之前。

\3. GZipMiddleware。

\4. ConditionalGetMiddleware。

\5. SessionMiddleware。

\6. LocaleMiddleware。

\7. CommonMiddleware。

\8. CsrfViewMiddleware。

\9. AuthenticationMiddleware。

10.MessageMiddleware。

11.FetchFromCacheMiddleware。

12.FlatpageFallbackMiddleware。

13.RedirectFallbackMiddleware。

 

 

 

# **Django-project**

## **前端开发环境配置：**

### **nvm安装：**

nvm（Node Version Manager）是一个用来管理node版本的工具。我们之所以需要使用node，是因为我们需要使用node中的npm(Node Package Manager)，使用npm的目的是为了能够方便的管理一些前端开发的包！nvm的安装非常简单，步骤如下：

到这个链接下载nvm的安装包：https://github.com/coreybutler/nvm-windows/releases。

然后点击一顿下一步，安装即可！

安装完成后，还需要配置环境变量。在我的电脑->属性->高级系统设置->环境变量->系统环境变量->Path下新建一个，把nvm所处的路径填入进去即可！

打开cmd，然后输入nvm，如果没有提示没有找不到这个命令。说明已经安装成功！

Mac或者Linux安装nvm请看这里：https://github.com/creationix/nvm。也要记得配置环境变量。

nvm常用命令：

nvm install node：安装最新版的node.js。nvm i == nvm install。

nvm install [version]：安装指定版本的node.js 。

nvm use [version]：使用某个版本的node。

nvm list：列出当前安装了哪些版本的node。

nvm uninstall [version]：卸载指定版本的node。

nvm node_mirror [url]：设置nvm的镜像。

nvm npm_mirror [url]：设置npm的镜像。

### **node安装：**

安装完nvm后，我们就可以通过nvm来安装node了。这里我们安装6.4.0版本的的node.js就可以。因为最新版的node.js的npm是5.0的，上面还有很多坑。安装命令如下：

nvm install 6.4.0

如果你的网络够快，那以上命令在稍等片刻之后会安装成功。如果你的网速很慢，那以上命令可能会发生超时。因为node的服务器地址是https://nodejs.org/dist/，这个域名的服务器是在国外。因此会比较慢。因此我们可以设置一下nvm的源。

nvm node_mirror https://npm.taobao.org/mirrors/node/

nvm npm_mirror https://npm.taobao.org/mirrors/npm/

npm：

npm(Node Package Manager)在安装node的时候就会自动的安装了。当时前提条件是你需要设置当前的node的版本：nvm use 6.4.0。然后就可以使用npm了.
关于npm常用命令以及用法，请看下文。

### **安装包：**

安装包分为全局安装和本地安装。全局安装是安装在当前node环境中，在可以在cmd中当作命令使用。而本地安装是安装在当前项目中，只有当前这个项目能使用，并且可以通过require引用。安装的方式只有-g参数的区别：

npm install express          # 本地安装

npm install express -g   # 全局安装

### **本地安装**

将安装包放在./node_modules下（运行 npm 命令时所在的目录），如果没有node_modules目录，会在当前执行npm命令的目录下生成node_modules目录。

可以通过require()来引入本地安装的包。

全局安装

将安装包放在/usr/local下或者你node的安装目录。

可以直接在命令行里使用。

卸载包：

npm uninstall [package]

更新包：

npm update [package]

搜索包：

npm search [package]

使用淘宝镜像：

npm install -g cnpm --registry=<https://registry.npm.taobao.org> 那么以后就可以使用cnpm来安装包了！

 

 

## **前端项目搭建**

前端我们使用gulp来自动化开发流程。配置好gulp后，可以自动给我们处理好一些工作。比如写完css后，要压缩成.min.css，写完js后，要做混淆和压缩，图片压缩等。这些工作都可以让gulp帮我们完成。

### **安装gulp：**

### **1. 创建本地包管理环境：**

使用npm init命令在本地生成一个package.json文件，package.json是用来记录你当前这个项目依赖了哪些包，以后别人拿到你这个项目后，不需要你的node_modules文件夹（因为node_moduels中的包实在太庞大了）。只需要执行npm install命令，即会自动安装package.json下devDependencies中指定的依赖包。

### **2. 安装gulp：**

gulp的安装非常简单，只要使用npm命令安装即可。但是因为gulp需要作为命令行的方式运行，因此需要在安装在系统级别的目录中。

npm install gulp -g

因为在本地需要使用require的方式gulp。因此也需要在本地安装一份：

npm install gulp --save-dev

以上的--save-dev是将安装的包的添加到package.json下的devDependencies依赖中。以后通过npm install即可自动安装。devDependencies这个是用来记录开发环境下使用的包，如果想要记录生产环境下使用的包，那么在安装包的时候使用npm install xx --save就会记录到package.json下的dependencies中，dependencies是专门用来记录生产环境下的依赖包的！

### **3. 创建gulp任务：**

要使用gulp来流程化我们的开发工作。首先需要在项目的根目录下创建一个gulpfile.js文件。然后在gulpfile.js中填入以下代码：

var gulp = require("gulp")

 

gulp.task("greet",function () {

​    console.log('hello world');

});

这里对代码进行一一解释：

 

通过require语句引用已经安装的第三方依赖包。这个require只能是引用当前项目的，不能引用全局下的。require语法是node.js独有的，只能在node.js环境下使用。

 

 

gulp.task是用来创建一个任务。gulp.task的第一个参数是命令的名字，第二个参数是一个函数，就是执行这个命令的时候会做什么事情，都是写在这个里面的。

 

写完以上代码后，以后如果想要执行greet命令，那么只需要进入到项目所在的路径，然后终端使用gulp greet即可执行。

### **4. 创建处理css文件的任务：**

gulp只是提供一个框架给我们。如果我们想要实现一些更加复杂的功能，比如css压缩，那么我们还需要安装一下gulp-cssnano插件。gulp相关的插件安装也是通过npm命令安装，安装方式跟其他包是一模一样的（gulp插件本身就是一个普通的包）。
对css文件的处理，需要做的事情就是压缩，然后再将压缩后的文件放到指定目录下（不要和原来css文件重合了）！这里我们使用gulp-cssnano来处理这个工作：

npm install gulp-cssnano --save-dev

然后在gulpfile.js中写入以下代码：

var gulp = require("gulp")var cssnano = require("gulp-cssnano")

// 定义一个处理css文件改动的任务

gulp.task("css",function () {

​    gulp.src("./css/*.css")

​    .pipe(cssnano())

​    .pipe(gulp.dest("./css/dist/"))

});

以上对代码进行详细解释：

gulp.task：创建一个css处理的任务。

gulp.src：找到当前css目录下所有以.css结尾的css文件。

pipe：管道方法。将上一个方法的返回结果传给另外一个处理器。比如以上的cssnano。

gulp.dest：将处理完后的文件，放到指定的目录下。不要放在和原文件相同的目录，以免产生冲突，也不方便管理。

### **5. 修改文件名：**

像以上任务，压缩完css文件后，最好是给他添加一个.min.css的后缀，这样一眼就能知道这个是经过压缩后的文件。这时候我们就需要使用gulp-rename来修改了。当然首先也需要安装npm install gulp-rename --save-dev。示例代码如下：

var gulp = require("gulp")var cssnano = require("gulp-cssnano")var rename = require("gulp-rename")

gulp.task("css",function () {

​    gulp.src("./css/*.css")

​    .pipe(cssnano())

​    .pipe(rename({"suffix":".min"}))

​    .pipe(gulp.dest("./css/dist/"))

});

在上述代码中，我们增加了一行.pipe(rename({"suffix":".min"}))，这个我们就是使用rename方法，并且传递一个对象参数，指定修改名字的规则为添加一个.min后缀名。这个gulp-rename还有其他的指定文件名的方式，比如可以在文件名前加个前缀等。更多的教程可以看这个：https://www.npmjs.com/package/gulp-rename。

### **6. 创建处理js文件的任务：**

处理js文件，我们需要使用到gulp-uglify插件。安装命令如下：

npm install gulp-uglify --save-dev

安装完后，我们就可以对js文件进行处理了。示例代码如下：

var gulp = require("gulp")var rename = require("gulp-rename")var uglify = require('gulp-uglify');

gulp.task('script',function(){

​    gulp.src(path.js + '*.js')

​    .pipe(uglify())

​    .pipe(rename({suffix:'.min'}))

​    .pipe(gulp.dest('js/'));

});

这里就是增加了一个.pipe(uglify())的处理，对js文件进行压缩和丑化（修改变量名）等处理。更多关于gulp-uglify的教程。请看：https://github.com/mishoo/UglifyJS2#minify-options。

### **7. 合并多个文件：**

在网页开发中，为了加快网页的渲染速度，有时候我们会将多个文件压缩成一个文件，从而减少请求的次数。要拼接文件，我们需要用到gulp-concat插件。安装命令如下：

npm install gulp-concat --save-dev

比如我们现在有一个nav.js文件用来控制导航条的。有一个index.js文件用来控制首页整体内容的。那么我们可以使用以下代码将这两个文件合并成一个文件：

var gulp = require('gulp');var concat = require('gulp-concat');var uglify = require('gulp-uglify');

gulp.task('vendorjs',function(){

​    gulp.src([

​        './js/nav.js',

​        './js/index.js'

​    ])

​    .pipe(concat('index.min.js'))

​    .pipe(uglify())

​    .pipe(gulp.dest('dist/js/'));

});

### **8. 压缩图片：**

图片是限制网站加载速度的一个主要原因。图片越大，从网站上下载所花费的时间越长。因此对于一些图片，我们可以采取无损压缩，即在不改变图片质量的基础之上进行压缩。在gulp中我们可以通过gulp-imagemin来帮我们实现。安装命令如下：

npm install gulp-imagemin --save-dev

压缩图片也是一个比较大的工作量，对于一些已经压缩过的图片，我们就没必要再重复压缩了。这时候我们可以使用gulp-cache来缓存那些压缩过的图片。安装命令如下：

npm install gulp-cache --save-dev

两个插件结合使用的代码如下：

var imagemin = require('gulp-imagemin');var cache = require('gulp-cache');

gulp.task('image',function(){

​    gulp.src("./images/*.*")

​    .pipe(cache(imagemin()))

​    .pipe(gulp.dest('dist/images/'));

});

### **9. 检测代码修改，自动刷新浏览器：**

以上所有的任务，我们都是需要手动的在终端去执行。这样很不方便我们开发。最好的方式就是我修改了代码后，gulp会自动的执行相应的任务。这个工作我们可以使用gulp内置的watch方法帮我们完成：

var gulp = require("gulp")var cssnano = require("gulp-cssnano")var rename = require("gulp-rename")

// 定义一个处理css文件改动的任务

gulp.task("css",function () {

​    gulp.src("./css/*.css")

​    .pipe(cssnano())

​    .pipe(rename({"suffix":".min"}))

​    .pipe(gulp.dest("./css/dist/"))

​    .pipe(connect.reload())

});

// 定义一个监听的任务

gulp.task("watch",function () {

​    // 监听所有的css文件，然后执行css这个任务

​    gulp.watch("./css/*.css",['css'])

});

以后只要在终端执行gulp watch命令即可自动监听所有的css文件，然后自动执行css的任务，完成相应的工作。

### **10. 更改文件后，自动刷新浏览器：**

以上我们实现了更改一些css文件后，可以自动执行处理css的任务。但是我们还是需要手动的去刷新浏览器，才能看到修改后的效果。有什么办法能在修改完代码后，自动的刷新浏览器呢。答案是使用browser-sync。browser-sync安装的命令如下：

npm install browser-sync --save-dev

browser-sync使用的示例代码如下：

var gulp = require("gulp")var cssnano = require("gulp-cssnano")var rename = require("gulp-rename")var bs = require("browser-sync").create()

 

gulp.task("bs",function () {

​    bs.init({

​        'server': {

​            'baseDir': './'

​        }

​    });

});

// 定义一个处理css文件改动的任务

gulp.task("css",function () {

​    gulp.src("./css/*.css")

​    .pipe(cssnano())

​    .pipe(rename({"suffix":".min"}))

​    .pipe(gulp.dest("./css/dist/"))

​    .pipe(bs.stream())

});

// 定义一个监听的任务

gulp.task("watch",function () {

​    gulp.watch("./css/*.css",['css'])

});

// 执行gulp server开启服务器

gulp.task("server",['bs','watch'])

以上我们创建了一个bs的任务，这个任务会开启一个3000端口，以后我们在访问html页面的时候，就需要通过http://127.0.0.1:3000的方式来访问了。然后接下来我们还定义了一个server任务。这个任务会去执行bs和watch任务，只要修改了css文件，那么就会执行css的任务，然后就会自动刷新浏览器。
browser-sync更多的教程请参考：http://www.browsersync.cn/docs/gulp/。

 
