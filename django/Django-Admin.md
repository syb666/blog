# 一、创建项目

使用pycharm创建Django项目,使用命令创建应用 python manage.py startapp goods (创建应用goods)

# 二、配置文件

Settings.py文件INSTALLED_APPS中加载应用配置实现注册应用,并配置数据库

```
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'test_1',
        'USER':'root',
        'PASSWORD':'123456',
        'PORT':'3306',
        'HOST':'127.0.0.1'
    }
}
```
使用pymsql驱动,在 项 目 同 名 文 件 夹 下 的__init__.py 文件初始化mysql驱动。

```
import pymysql
pymysql.install_as_MySQLdb()
```
# 三、goods应用定义模型

商店-------------------->商品(一对多)
```
from django.db import models

# Create your models here.
class Shop(models.Model):
    name = models.CharField(max_length=20,verbose_name='店铺名称')
    location = models.CharField(max_length=30,verbose_name='店铺地址')
    create_date = models.DateField(verbose_name='开业时间')
    is_delete = models.BooleanField(default=False,verbose_name='是否删除')
    class Meta:
        db_table = 'shop'
        verbose_name_plural = '商店'
    def __str__(self):
        return self.name

class Goods(models.Model):
    choices_gender = (
        (0,'男'),
        (1,'女')
    )
    name = models.CharField(max_length=20,verbose_name='商品名称')
    brand_name = models.CharField(max_length=20,verbose_name='品牌名称')
    show_date = models.DateField(verbose_name='上市日期')
    gender = models.IntegerField(choices=choices_gender,default=0,verbose_name='性别')
    comment = models.CharField(max_length=200,null=True,blank=True,verbose_name='商品描述')
    is_show = models.BooleanField(default=True,verbose_name='是否展示')
    shop = models.ForeignKey('Shop',on_delete=models.CASCADE,verbose_name='店铺')

    class Meta:
        db_table = 'goods'
        #本来是默认的goods+s，下面这个可以修改名字
        verbose_name_plural = '商品'
    def __str__(self):
        return self.name
```
# 四、数据库迁移,创建对象

数 据 库 迁 移 : python manage.py makemigrations 生成迁移文件pyhton manage.py migrate 同步到数据库进 入 交 互 环 境 , 创 建 对 象 python manage.pyshell (导入模型)启 动 python 两 种 方 法 ,1.pyhton 2.pythonmanage.py shell,都会启动交互解释器,但是第二种启动解释器会告诉Django使用哪个配置文件,Django框架大部分子系统都依赖于配置文件.，然后在交互解释器中创建对象。

# 五、创建超级管理员

```python manage.py createsuperuser。```然后输入一些信息

# 六、界面本地化

因为我们在登录站点时，他默认的是Django自带的信息，都是英文的。这是我们只需要修改settings里的信息，就可以改变。
```
LANGUAGE_CODE = 'zh-hans'
TIME_ZONE = 'Asia/Shanghai'
```
# 七、注册模型

在应用goods下的admin.py文件中注册模型,可以实现对数据进行增删改管理操作。

注册1：使用自带的站点
```
from django.contrib import admin
from .models import Shop,Goods

# 直接使用自带的站点
admin.site.register(Shop)
admin.site.register(Goods)
```
# 八、站点显示备注信息

为模型字段选项添加verbose_name显示字段备注信息,verbose_name_plural显示表名备注信息
![](https://img-blog.csdnimg.cn/20190418153310870.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
![](https://img-blog.csdnimg.cn/20190418153529808.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)


修 改 应 用 名 : 在 应 用 下 的 app.py 文 件 下 添加verbose_name
![](https://img-blog.csdnimg.cn/20190418153933614.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)


# 九、定义与使用admin管理类

ModelAdmin类控制admin站点的展示效果,我们可以自定义管理类,继承ModelAdmin

注册1：
```
# 直接使用自带的站点
admin.site.register(Shop)
admin.site.register(Goods)
```

注册2：使用装饰器注册
```
@admin.register(Shop)
class ShopAdmin(admin.ModelAdmin):
    pass
    
@admin.register(Goods)
class GoodsAdmin(admin.ModelAdmin):
    pass
```
# 十、自定义编辑页面显示

fileds:编辑页面显示可编辑字段,如果嵌套,会在同一行显示多个字段
![](https://img-blog.csdnimg.cn/20190418154403152.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)


exclude:编辑页面不显示可编辑字段
![](https://img-blog.csdnimg.cn/20190418154513551.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)


filedsets:编辑页面对编辑字段进行分组，filedsets不能与fileds同时使用。
![](https://img-blog.csdnimg.cn/20190418155600678.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)


如果字段为日期或时间类型,并选项auto_now或auto_now_add设置为True,那么这个字段是不可编辑的,默认不会显示,也不能手动添加

# 十一、InlineModelAdmin内联(关联对象)

在管理界面能够在与父模型相同的页面上编辑模型,在一对多的关系中,我们可以在一的 一方编辑多的对象两 个 子 类:TabularInline 以 表 格 的 形 式 嵌 入StackenInline 以块状形式嵌入这两种的方法仅仅只是用于渲染他们的模板。

```
class GoodsTabularInline(admin.TabularInline):
    #关联模型
    model = Goods

class GoodsStackedInline(admin.StackedInline):
    # 关联模型
    model = Goods
```
1. 以块状形式显示
```
@admin.register(Shop)
class ShopAdmin(admin.ModelAdmin):
    # pass
    inlines = [
        GoodsStackedInline
    ]
```
![](https://img-blog.csdnimg.cn/20190418161035362.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
2. 以表格形式显示
```
@admin.register(Shop)
class ShopAdmin(admin.ModelAdmin):
    # pass
    inlines = [
        GoodsTabularInline
    ]
```
![](https://img-blog.csdnimg.cn/20190418161229177.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
# 十二、自定义列表页显示

list_display : 定 义 列 表 页 显 示 的 列 , 未 定 义list_display 默认显示的是定义模型的__str__方法返回的对象字符串
![](https://img-blog.csdnimg.cn/20190418161429269.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)

如果字段是 ForeignKey,Django 将显示__str__相关对象(在我们的Goods 模型 中,shop是Goods模型的外键,那么shop字段显示的就是我们Shop模型中 的__str__()对象)如果字段是 BooleanFiled,Django 将显示一个开关，而不是 True 或 False如果字段的值为 None,Django 默认显示 - 破折号, empty_value_display 属性设置值为 None的显示
```
#通过列表显示
list_display = ['name','brand_name','show_date','gender','shop','comment','is_show']```
# 设置字段为NUll的显示
empty_value_display = '未填
'```
![](https://img-blog.csdnimg.cn/20190418161619612.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
```
list_display_links:设置进入编辑页面的链接,默认第一列链接进入到编辑页面,将其设置 为 None 完全没有链接,将其设置为 list_display 要转换为链接的列的列表或元组, list_display_links 必须和list_display 配合使用

#通过列表显示
list_display = ['name','brand_name','show_date','gender','shop','comment','is_show']```
# 设置字段为NUll的显示
empty_value_display = '未填'
#设置进入编辑页的链接，默认第一列，必须和list_display搭配使用
list_display_links = ['name','brand_name','show_date']
```


list_editable:允许字段在列表页进行编辑,必须和list_display 配 合 使 用 , 无 法 编 辑 未 显 示 的 字段 ,list_editable 列 表 字 段 , 不 能 和list_display_links 字段名冲突,不能有进入到编辑页面的链 接.
```
#通过列表显示
list_display = ['name','brand_name','show_date','gender','shop','comment','is_show']
# 设置字段为NUll的显示
empty_value_display = '未填'
# #设置进入编辑页的链接，默认第一列，必须和list_display搭配使用
list_display_links = ['name','brand_name','show_date']
# 允许字段可以再列表页可以编辑,这里的属性不能与上面的字段名有冲突的
list_editable = ['comment','is_show']

list_filter 列表页显示右侧的过滤器, 关联对象过滤：外键名__关联对象字段名
![](https://img-blog.csdnimg.cn/20190418162224443.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)

list_per_page:控制列表页显示的条数,默认100

list_max_show_all:控制’显示全部’的列表页的显示条数,默认 200 条

search_fields : 在 列 表 页 启 用 搜 索 框search_fields 字段名列表,关联对象外键查询 search_fields=[‘外键名__字段名’]
```
# 十三、 AdminSite 站点信息
```
调整 adminsite 基本信息
admin.site.site_title= 站点标题
admin.site.site_header= 站点头
admin.site.index_title= 首页标题
```
![](https://img-blog.csdnimg.cn/20190418162428955.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
![](https://img-blog.csdnimg.cn/20190418162658869.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hlbnVzeWI=,size_16,color_FFFFFF,t_70)
