# Flask-SQLAlchemy笔记：

### 安装：
```shell
pip install flask-sqlalchemy
```

### 数据库连接：
1. 跟sqlalchemy一样，定义好数据库连接字符串DB_URI。
2. 将这个定义好的数据库连接字符串DB_URI，通过`SQLALCHEMY_DATABASE_URI`这个键放到`app.config`中。示例代码：`app.config["SQLALCHEMY_DATABASE_URI"] = DB_URI`.
3. 使用`flask_sqlalchemy.SQLAlchemy`这个类定义一个对象，并将`app`传入进去。示例代码：`db = SQLAlchemy(app)`。

### 创建ORM模型：
1. 还是跟使用sqlalchemy一样，定义模型。现在不再是需要使用`delarative_base`来创建一个基类。而是使用`db.Model`来作为基类。
2. 在模型类中，`Column`、`String`、`Integer`以及`relationship`等，都不需要导入了，直接使用`db`下面相应的属性名就可以了。
3. 在定义模型的时候，可以不写`__tablename__`，那么`flask_sqlalchemy`会默认使用当前的模型的名字转换成小写来作为表的名字，并且如果这个模型的名字使用了多个单词并且使用了驼峰命名法，那么会在多个单词之间使用下划线来进行连接。**虽然flask_sqlalchemy给我们提供了这个特性，但是不推荐使用。因为明言胜于暗喻**

### 将ORM模型映射到数据库：
1. db.drop_all()
2. db.create_all()

### 使用session：
以后session也不需要使用`sessionmaker`来创建了。直接使用`db.session`就可以了。操作这个session的时候就跟之前的`sqlalchemy`的`session`是iyimoyiyang的。

### 查询数据：
如果查找数据只是查找一个模型上的数据，那么可以通过`模型.query`的方式进行查找。`query`就跟之前的sqlalchemy中的query方法是一样用的。示例代码如下：
    ```python
    users = User.query.order_by(User.id.desc()).all()
    print(users)
    ```
    
    ```
    app.config['SQLALCHEMY_DATABASE_URI'] = DB_URI
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    db = SQLAlchemy(app)
    # 明言胜于暗喻
    class UserModel(db.Model):
        __tablename__ = 'user_model'
        id = db.Column(db.Integer,primary_key=True,autoincrement=True)
        username = db.Column(db.String(50),nullable=False)

        def __repr__(self):
            return "<User(username: %s)>" % self.username

    class Article(db.Model):
        __tablename__ = 'article'
        id = db.Column(db.Integer,primary_key=True,autoincrement=True)
        title = db.Column(db.String(50),nullable=False)
        uid = db.Column(db.Integer,db.ForeignKey("user_model.id"))

        author = db.relationship("User",backref="artiles")

    db.drop_all()
    db.create_all()

    # user = User(username='zhiliao')
    # article = Article(title='title one')
    # # cascade save-update
    # article.author = user
    #
    # db.session.add(article)
    # db.session.commit()
    # order_by
    # filter
    # filter_by
    # group_by
    # join
    # users = User.query.order_by(User.id.desc()).all()
    # print(users)

    # user = User.query.filter(User.username=='zhiliao').first()
    # user.username = 'zhiliao1'
    # db.session.commit()

    # user = User.query.filter(User.username=='zhiliao1').first()
    # db.session.delete(user)
    # db.session.commit

    # db.drop_all()
    # db.create_all()

    # user = User(username='小明')
    # article = Article(title='1111')
    # article.author = user
    # db.session.add(article)
    # db.session.commit()
    ```
