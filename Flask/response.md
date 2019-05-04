# response笔记：

### 视图函数中可以返回哪些值：
1. 可以返回字符串：返回的字符串其实底层将这个字符串包装成了一个`Response`对象。
2. 可以返回元组：元组的形式是(响应体,状态码,头部信息)，也不一定三个都要写，写两个也是可以的。返回的元组，其实在底层也是包装成了一个`Response`对象。
3. 可以返回`Response`及其子类。


### 实现一个自定义的`Response`对象：
1. 继承自`Response`类。
2. 实现方法`force_type(cls,rv,environ=None)`。
3. 指定`app.response_class`为你自定义的`Response`对象。
4. 如果视图函数返回的数据，不是字符串，也不是元组，也不是Response对象，那么就会将返回值传给`force_type`，然后再将`force_type`的返回值返回给前端。
```python
from flask import Flask,Response,jsonify
from  werkzeug.wrappers import Response
import json

app = Flask(__name__)

#如果将视图函数中返回的字典，转换成json对象，然后返回

class JSONResponse(Response):
    @classmethod
    def force_type(cls, response, environ=None):
        """
        这个方法只有返回非字符串，非元祖，非Response对象才会调用
        response：视图函数的返回值
        :param response:
        :param environ:
        :return:
        """
        if isinstance(response,dict):
            #jsonify除了将字典转换为json，还将该对象包装成了一个response对象
            response = jsonify(response)#jsonify转换为json对象，然后再调用那个对象
        return super(JSONResponse,cls).force_type(response,environ)
        # print(response)
        # print(type(response))
        # return Response("hello")

app.response_class = JSONResponse

@app.route('/')
def hello_world():
    return Response('hello',status=200,mimetype='text/html')
    # return 'Hello World!'

@app.route('/about/')
def about():
    resp = Response(response='about page',status=200,content_type='text/html;charset=utf-8')
    return resp
@app.route('/list/')
def list1():
    return {'username':'sunyabo'}#这里是去执行上面自定义的JSONResponse中的force_type然后返回一个"hello"

@app.route('/list2/')
def list2():
    resp = Response("hello")
    resp.set_cookie('country','china')
    return resp#这里是去执行上面自定义的JSONResponse中的force_type然后返回一个"hello"


if __name__ == '__main__':
    app.run()

```
