# DEBUG模式
### 为什么需要开启DEBUG模式
1. 如果开启DEBUG模式，那么在代码中如果抛出了异常，在浏览器的页面中可以看到具体的错误信息，以及具体的错误代码位置，方便开发者调试。
2. 如果开DEBUG模式，那么在以后的`Python`代码中修改了任何代码,只要按`ctrl+s`，'flask'就会自动的重新记载整个网站。不需要手动点击重新运行。

### 配置DEBUG模式的四种方式
1. 在`app.run()`中传递一个参数`debug=True`就可以开启`DEBUG`模式。
2. 在`app.debug=True`也可以开启`debug`模式。
3. 通过配置参数的形式设置DEBUG模式: `app.config.update(DEBUG=True)`。
4. 通过配置文件的形式设置DEBUG模式，`app.config.from_object(config)`。

### PIN码
如果想要在网页上调试代码，那么应该输入`PIN`码。
