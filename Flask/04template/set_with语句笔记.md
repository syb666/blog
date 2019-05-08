# set、with语句笔记：

### set语句：
在模版中，可以使用`set`语句来定义变量。示例如下：
```html
{% set username='知了课堂' %}
<p>用户名：{{ username }}</p>
```
一旦定义了这个变量，那么在后面的代码中，都可以使用这个变量，就类似于Python的变量定义是一样的。

### `with`语句：
`with`语句定义的变量，只能在`with`语句块中使用，超过了这个代码块，就不能再使用了。示例代码如下：
```html
{% with classroom = 'zhiliao1班' %}
<p>班级：{{ classroom }}</p>
{% endwith %}
```
`with`语句也不一定要跟一个变量，可以定义一个空的`with`语句，以后在`with`块中通过`set`定义的变量，就只能在这个`with`块中使用了：
```html
{% with %}
    {% set classroom = 'zhiliao1班' %}
    <p>班级：{{ classroom }}</p>
{% endwith %}
```
