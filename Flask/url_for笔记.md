# url_for笔记：

### `url_for`的基本使用：
`url_for`第一个参数，应该是视图函数的名字的字符串。后面的参数就是传递给`url`。
如果传递的参数之前在`url`中已经定义了，那么这个参数就会被当成`path`的形式给
`url`。如果这个参数之前没有在`url`中定义，那么将变成查询字符串的形式放到`url`中。
```python
@app.route('/post/list/<page>/')
def my_list(page):
    return 'my list'

print(url_for('my_list',page=1,count=2))
# 构建出来的url：/my_list/1/?count=2
```

### 为什么需要`url_for`：
1. 将来如果修改了`URL`，但没有修改该URL对应的函数名，就不用到处去替换URL了。
2. `url_for`会自动的处理那些特殊的字符，不需要手动去处理。
    ```python
    url = url_for('login',next='/')
    # 会自动的将/编码，不需要手动去处理。
    # url=/login/?next=%2F
    ```

### 强烈建议以后在使用url的时候，使用`url_for`来反转url。
