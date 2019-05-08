# 宏：
模板中的宏跟python中的函数类似，可以传递参数，但是不能有返回值，可以将一些经常用到的代码片段放到宏中，然后把一些不固定的值抽取出来当成一个变量。
使用宏的时候，参数可以为默认值。相关示例代码如下：
1. 定义宏：
    ```html
    {% macro input(name, value='', type='text') %}
    <input type="{{ type }}" name="{{ name }}" value="{{
    value }}">
    {% endmacro %}
    ```
2. 使用宏：
    ```html
    <p>{{ input('username') }}</p>
    <p>{{ input('password', type='password') }}</p>
    ```

### 导入宏：
1. `import "宏文件的路径" as xxx`。
2. `from '宏文件的路径' import 宏的名字 [as xxx]`。
3. 宏文件路径，不要以相对路径去寻找，都要以`templates`作为绝对路径去找。
4. 如果想要在导入宏的时候，就把当前模版的一些参数传给宏所在的模版，那么就应该在导入的时候使用`with context`。示例：`from 'xxx.html' import input with context`。
# include标签：
1. 这个标签相当于是直接将指定的模版中的代码复制粘贴到当前位置。
2. `include`标签，如果想要使用父模版中的变量，直接用就可以了，不需要使用`with context`。
3. `include`的路径，也是跟`import`一样，直接从`templates`根目录下去找，不要以相对路径去找。
