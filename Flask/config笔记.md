# config笔记
### 使用`app.config.from_object`的方式加载配置文件
1. 导入`import config`。
2. 导入`app.config.from_object(config)`。
### 使用`app.config.from_pyfile`的方式加载配置文件
这种方式不需要`import`,直接使用`app.config.from_object('config.py')`就可以了。
注意这个地方，必须要写文件的全名，后缀名不能少。
1. 这种方式，加载配置文件，不局限于只能使用`py`文件，普通的`txt`文件也可以。
2. 这种方式，可以传递`silent=True`，那么这个静态文件没有找到的时候，不会抛出异常。
