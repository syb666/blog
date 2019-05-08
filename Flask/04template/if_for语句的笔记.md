# if条件判断语句：
`if`条件判断语句必须放在`{% if statement %}`中间，并且还必须有结束的标签`{% endif %}`。和`python`中的类似，
可以使用`>，<，<=，>=，==，!=`来进行判断，也可以通过`and，or，not，()`来进行逻辑合并操作。
# for循环语句笔记：
在`jinja2`中的`for`循环，跟`python`中的`for`循环基本上是一模一样的。也是`for...in...`的形式。并且也可以遍历所有的序列以及迭代器。
但是唯一不同的是，`jinja2`中的`for`循环没有`break`和`continue`语句。
```python
1. if ：if语句和 python 中的类似，可以使用 >，<，<=，>=，==，!= 来进行判断，也可以通
过 and，or，not，() 来进行逻辑合并操作，以下看例子：
{% if kenny.sick %}
 Kenny is sick.
{% elif kenny.dead %}
 You killed Kenny! You bastard!!!
{% else %}
 Kenny looks okay --- so far
{% endif %}
2. for...in... ： for 循环可以遍历任何一个序列包括列表、字典、元组。并且可以进行反向
遍历，以下将用几个例子进行解释：
普通的遍历：
<ul>
 {% for user in users %}
<li>{{ user.username|e }}</li>
 {% endfor %}
</ul>
遍历字典：
<dl>
 {% for key, value in my_dict.iteritems() %}
<dt>{{ key|e }}</dt>
<dd>{{ value|e }}</dd>
 {% endfor %}
</dl>
如果序列中没有值的时候，进入 else ：
<ul>
 {% for user in users %}
<li>{{ user.username|e }}</li>
 {% else %}
<li><em>no users found</em></li>
 {% endfor %}
</ul>
```
并且 Jinja 中的 for 循环还包含以下变量，可以用来获取当前的遍历状态：
| 变量 | 描述 | | --- | --- | | loop.index | 当前迭代的索引（从1开始） | | loop.index0 | 当前迭代的
索引（从0开始） | | loop.first | 是否是第一次迭代，返回True或False | | loop.last | 是否是最后
一次迭代，返回True或False | | loop.length | 序列的长度 |
另外，不可以使用 continue 和 break 表达式来控制循环的执行。
