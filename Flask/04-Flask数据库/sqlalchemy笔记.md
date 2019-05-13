# SQLALchemy笔记

### 使用SQLAlchemy去连接数据库：
使用SQLALchemy去连接数据库，需要使用一些配置信息，然后将他们组合成满足条件的字符串：
```python
HOSTNAME = '127.0.0.1'
PORT = '3306'
DATABASE = 'first_sqlalchemy'
USERNAME = 'root'
PASSWORD = 'root'

# dialect+driver://username:password@host:port/database
DB_URI = "mysql+pymysql://{username}:{password}@{host}:{port}/{db}?charset=utf8".format(username=USERNAME,password=PASSWORD,host=HOSTNAME,port=PORT,db=DATABASE)
```
然后使用`create_engine`创建一个引擎`engine`，然后再调用这个引擎的`connect`方法，就可以得到这个对象，然后就可以通过这个对象对数据库进行操作了：
```python
engine = create_engine(DB_URI)

# 判断是否连接成功
conn = engine.connect()
result = conn.execute('select 1')
print(result.fetchone())
```


### ORM介绍：
1. ORM：Object Relationship Mapping
2. 大白话：对象模型与数据库表的映射

### 将ORM模型映射到数据库中：
1. 用`declarative_base`根据`engine`创建一个ORM基类。
    ```python
    from sqlalchemy import create_engine,Column,Integer,String
    from sqlalchemy.ext.declarative import declarative_base
    HOSTNAME = '127.0.0.1'
    PORT = '3306'
    DATABASE = 'first_sqlalchemy'
    USERNAME = 'root'
    PASSWORD = '123456'
    DB_URI = "mysql+pymysql://{username}:{password}@{host}:{port}/{db}?charset=utf8".format(username=USERNAME,password=PASSWORD, host=HOSTNAME, port=PORT,db=DATABASE)
    engine = create_engine(DB_URI)
    Base = declarative_base(engine)
    ```
2. 用这个`Base`类作为基类来写自己的ORM类。要定义`__tablename__`类属性，来指定这个模型映射到数据库中的表名。
    ```python
    class Person(Base):
        __tablename__ = 'person'
    ```
3. 创建属性来映射到表中的字段，所有需要映射到表中的属性都应该为Column类型：
    ```python
    class Person(Base):
        __tablename__ = 'person'
        # 2. 在这个ORM模型中创建一些属性，来跟表中的字段进行一一映射。这些属性必须是sqlalchemy给我们提供好的数据类型。
        id = Column(Integer,primary_key=True,autoincrement=True)
        name = Column(String(50))
        age = Column(Integer)
    ```
4. 使用`Base.metadata.create_all()`来将模型映射到数据库中。
5. 一旦使用`Base.metadata.create_all()`将模型映射到数据库中后，即使改变了模型的字段，也不会重新映射了。

### 用session做数据的增删改查操作：
1. 构建session对象：所有和数据库的ORM操作都必须通过一个叫做`session`的会话对象来实现，通过以下代码来获取会话对象：
    ```python
    from sqlalchemy.orm import sessionmaker

    engine = create_engine(DB_URI)
    session = sessionmaker(engine)()
    ```
2. 添加对象：
    * 创建对象，也即创建一条数据：
        ```python
        p = Person(name='zhiliao',age=18,country='china')
        ```
    * 将这个对象添加到`session`会话对象中：
        ```python
        session.add(p)
        ```
    * 将session中的对象做commit操作（提交）：
        ```python
        session.commit()
        ```
    * 一次性添加多条数据：
        ```python
        p1 = Person(name='zhiliao1',age=19,country='china')
        p2 = Person(name='zhiliao2',age=20,country='china')
        session.add_all([p1,p2])
        session.commit()
        ```
3. 查找对象：
    ```python
    # 查找某个模型对应的那个表中所有的数据：
    all_person = session.query(Person).all()
    # 使用filter_by来做条件查询
    all_person = session.query(Person).filter_by(name='zhiliao').all()
    # 使用filter来做条件查询
    all_person = session.query(Person).filter(Person.name=='zhiliao').all()
    # 使用get方法查找数据，get方法是根据id来查找的，只会返回一条数据或者None
    person = session.query(Person).get(primary_key)
    # 使用first方法获取结果集中的第一条数据
    person = session.query(Person).first()
    ```
4. 修改对象：首先从数据库中查找对象，然后将这条数据修改为你想要的数据，最后做commit操作就可以修改数据了。
    ```python
    person = session.query(Person).first()
    person.name = 'ketang'
    session.commit()
    ```
5. 删除对象：将需要删除的数据从数据库中查找出来，然后使用`session.delete`方法将这条数据从session中删除，最后做commit操作就可以了。
    ```python
    person = session.query(Person).first()
    session.delete(person)
    session.commit()
    ```

### SQLAlchemy常用数据类型：
1. Integer：整形，映射到数据库中是int类型。
2. Float：浮点类型，映射到数据库中是float类型。他占据的32位。
3. Double：双精度浮点类型，映射到数据库中是double类型，占据64位。
4. String：可变字符类型，映射到数据库中是varchar类型.
5. Boolean：布尔类型，映射到数据库中的是tinyint类型。
6. DECIMAL：定点类型。是专门为了解决浮点类型精度丢失的问题的。在存储钱相关的字段的时候建议大家都使用这个数据类型。并且这个类型使用的时候需要传递两个参数，第一个参数是用来标记这个字段总能能存储多少个数字，第二个参数表示小数点后有多少位。
7. Enum：枚举类型。指定某个字段只能是枚举中指定的几个值，不能为其他值。在ORM模型中，使用Enum来作为枚举，示例代码如下：
    ```python
    class Article(Base):
        __tablename__ = 'article'
        id = Column(Integer,primary_key=True,autoincrement=True)
        tag = Column(Enum("python",'flask','django'))
    ```
    在Python3中，已经内置了enum这个枚举的模块，我们也可以使用这个模块去定义相关的字段。示例代码如下：
    ```python
    class TagEnum(enum.Enum):
        python = "python"
        flask = "flask"
        django = "django"

    class Article(Base):
        __tablename__ = 'article'
        id = Column(Integer,primary_key=True,autoincrement=True)
        tag = Column(Enum(TagEnum))

    article = Article(tag=TagEnum.flask)
    ```
8. Date：存储时间，只能存储年月日。映射到数据库中是date类型。在Python代码中，可以使用`datetime.date`来指定。示例代码如下：
    ```python
    class Article(Base):
        __tablename__ = 'article'
        id = Column(Integer,primary_key=True,autoincrement=True)
        create_time = Column(Date)

    article = Article(create_time=date(2017,10,10))
    ```
9. DateTime：存储时间，可以存储年月日时分秒毫秒等。映射到数据库中也是datetime类型。在Python代码中，可以使用`datetime.datetime`来指定。示例代码如下：
    ```python
    class Article(Base):
        __tablename__ = 'article'
        id = Column(Integer,primary_key=True,autoincrement=True)
        create_time = Column(DateTime)

    article = Article(create_time=datetime(2011,11,11,11,11,11))
    ```
10. Time：存储时间，可以存储时分秒。映射到数据库中也是time类型。在Python代码中，可以使用`datetime.time`来至此那个。示例代码如下：
    ```python
    class Article(Base):
        __tablename__ = 'article'
        id = Column(Integer,primary_key=True,autoincrement=True)
        create_time = Column(Time)

    article = Article(create_time=time(hour=11,minute=11,second=11))
    ```
11. Text：存储长字符串。一般可以存储6W多个字符。如果超出了这个范围，可以使用LONGTEXT类型。映射到数据库中就是text类型。
12. LONGTEXT：长文本类型，映射到数据库中是longtext类型。


### Column常用参数：
1. primary_key：设置某个字段为主键。
2. autoincrement：设置这个字段为自动增长的。
3. default：设置某个字段的默认值。在发表时间这些字段上面经常用。
4. nullable：指定某个字段是否为空。默认值是True，就是可以为空。
5. unique：指定某个字段的值是否唯一。默认是False。
6. onupdate：在数据更新的时候会调用这个参数指定的值或者函数。在第一次插入这条数据的时候，不会用onupdate的值，只会使用default的值。常用的就是`update_time`（每次更新数据的时候都要更新的值）。
7. name：指定ORM模型中某个属性映射到表中的字段名。如果不指定，那么会使用这个属性的名字来作为字段名。如果指定了，就会使用指定的这个值作为参数。这个参数也可以当作位置参数，在第1个参数来指定。
    ```python
    title = Column(String(50),name='title',nullable=False)
    title = Column('my_title',String(50),nullable=False)
    ```

### query可用参数：
1. 模型对象。指定查找这个模型中所有的对象。
```
articles = session.query(Article).all()
for article in articles:
    print(article)
```
2. 模型中的属性。可以指定只查找某个模型的其中几个属性。
```
articles = session.query(Article.title,Article.price).all()
for article in articles:
    print(article)
```
3. 聚合函数。
    * func.count：统计行的数量。
    * func.avg：求平均值。
    * func.max：求最大值。
    * func.min：求最小值。
    * func.sum：求和。
    `func`上，其实没有任何聚合函数。但是因为他底层做了一些魔术，只要mysql中有的聚合函数，都可以通过func调用。

```
result = session.query(func.count(Article.id)).first()
print(result)
result = session.query(func.avg(Article.price)).first()
print(result)
result = session.query(func.max(Article.price)).first()
print(result)
result = session.query(func.min(Article.price)).first()
print(result)
result = session.query(func.sum(Article.price)).first()
print(result)
```


### filter过滤条件：
过滤是数据提取的一个很重要的功能，以下对一些常用的过滤条件进行解释，并且这些过滤条件都是只能通过filter方法实现的：
1. equals：
    ```python
    article = session.query(Article).filter(Article.title == "title0").first()
    print(article)
    ```
2. not equals:
    ```python
    query.filter(User.name != 'ed')
    ```
2. like：
    ```python
    query.filter(User.name.like('%ed%'))
    ```

3. in：
    ```python
    query.filter(User.name.in_(['ed','wendy','jack']))
    # 同时，in也可以作用于一个Query
    query.filter(User.name.in_(session.query(User.name).filter(User.name.like('%ed%'))))
    ```

4. not in：
    ```python
    query.filter(~User.name.in_(['ed','wendy','jack']))
    ```
5.  is null：
    ```python
    query.filter(User.name==None)
    # 或者是
    query.filter(User.name.is_(None))
    ```

6. is not null:
    ```python
    query.filter(User.name != None)
    # 或者是
    query.filter(User.name.isnot(None))
    ```

7. and：
    ```python
    from sqlalchemy import and_
    query.filter(and_(User.name=='ed',User.fullname=='Ed Jones'))
    # 或者是传递多个参数
    query.filter(User.name=='ed',User.fullname=='Ed Jones')
    # 或者是通过多次filter操作
    query.filter(User.name=='ed').filter(User.fullname=='Ed Jones')
    ```

8. or：
    ```python
    from sqlalchemy import or_  
    query.filter(or_(User.name=='ed',User.name=='wendy'))
    ```

如果想要查看orm底层转换的sql语句，可以在filter方法后面不要再执行任何方法直接打印就可以看到了。比如：
    ```python
        articles = session.query(Article).filter(or_(Article.title=='abc',Article.content=='abc'))
        print(articles)
    ```

### 外键：
使用SQLAlchemy创建外键非常简单。在从表中增加一个字段，指定这个字段外键的是哪个表的哪个字段就可以了。从表中外键的字段，必须和父表的主键字段类型保持一致。
示例代码如下：
```python
class User(Base):
    __tablename__ = 'user'
    id = Column(Integer,primary_key=True,autoincrement=True)
    username = Column(String(50),nullable=False)

class Article(Base):
    __tablename__ = 'article'
    id = Column(Integer,primary_key=True,autoincrement=True)
    title = Column(String(50),nullable=False)
    content = Column(Text,nullable=False)

    uid = Column(Integer,ForeignKey("user.id"))
```
外键约束有以下几项： 
1. RESTRICT：父表数据被删除，会阻止删除。默认就是这一项。 
2. NO ACTION：在MySQL中，同RESTRICT。 
3. 3. CASCADE：级联删除。 
4. 4. SET NULL：父表数据被删除，子表数据会设置为NULL。

### ORM关系以及一对多：
mysql级别的外键，还不够ORM，必须拿到一个表的外键，然后通过这个外键再去另外一张表中查找，这样太麻烦了。SQLAlchemy提供了一个`relationship`，这个类可以定义属性，以后在访问相关联的表的时候就直接可以通过属性访问的方式就可以访问得到了。示例代码：
```python
class User(Base):
    __tablename__ = 'user'
    id = Column(Integer,primary_key=True,autoincrement=True)
    username = Column(String(50),nullable=False)

    # articles = relationship("Article")

    def __repr__(self):
        return "<User(username:%s)>" % self.username

class Article(Base):
    __tablename__ = 'article'
    id = Column(Integer,primary_key=True,autoincrement=True)
    title = Column(String(50),nullable=False)
    content = Column(Text,nullable=False)
    uid = Column(Integer,ForeignKey("user.id"))

    author = relationship("User",backref="articles")
```
另外，可以通过`backref`来指定反向访问的属性名称。articles是有多个。他们之间的关系是一个一对多的关系。

### 一对一的关系：
在sqlalchemy中，如果想要将两个模型映射成一对一的关系，那么应该在父模型中，指定引用的时候，要传递一个`uselist=False`这个参数进去。就是告诉父模型，以后引用这个从模型的时候，不再是一个列表了，而是一个对象了。示例代码如下：
```python
class User(Base):
    __tablename__ = 'user'
    id = Column(Integer,primary_key=True,autoincrement=True)
    username = Column(String(50),nullable=False)

    extend = relationship("UserExtend",uselist=False)

    def __repr__(self):
        return "<User(username:%s)>" % self.username

class UserExtend(Base):
    __tablename__ = 'user_extend'
    id = Column(Integer, primary_key=True, autoincrement=True)
    school = Column(String(50))
    uid = Column(Integer,ForeignKey("user.id"))

    user = relationship("User",backref="extend")
```
当然，也可以借助`sqlalchemy.orm.backref`来简化代码：
```python
class User(Base):
    __tablename__ = 'user'
    id = Column(Integer,primary_key=True,autoincrement=True)
    username = Column(String(50),nullable=False)

    # extend = relationship("UserExtend",uselist=False)

    def __repr__(self):
        return "<User(username:%s)>" % self.username

class UserExtend(Base):
    __tablename__ = 'user_extend'
    id = Column(Integer, primary_key=True, autoincrement=True)
    school = Column(String(50))
    uid = Column(Integer,ForeignKey("user.id"))

    user = relationship("User",backref=backref("extend",uselist=False))
    
user = User(username='zhiliao')
extend1 = UserExtend(school='zhiliao ketang')
user.extend = extend1

session.add(user)
session.commit()
    
```

### 多对多的关系：
1. 多对多的关系需要通过一张中间表来绑定他们之间的关系。
2. 先把两个需要做多对多的模型定义出来
3. 使用Table定义一个中间表，中间表一般就是包含两个模型的外键字段就可以了，并且让他们两个来作为一个“复合主键”。
4. 在两个需要做多对多的模型中随便选择一个模型，定义一个relationship属性，来绑定三者之间的关系，在使用relationship的时候，需要传入一个secondary=中间表。
```
article_tag = Table(
    'article_tag',
    Base.metadata,
    Column("article_id",Integer,ForeignKey("article.id"),primary_key=True),
    Column("tag_id",Integer,ForeignKey("tag.id"),primary_key=True),

)
class Article(Base):
    __tablename__ = 'article'
    id = Column(Integer,primary_key=True,autoincrement=True)
    title = Column(String(50),nullable=False)
    tags = relationship("Tag",backref="articles",secondary = article_tag)
    def __repr__(self):
        return "<article:%s>" % self.title

class Tag(Base):
    __tablename__ = 'tag'
    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String(50), nullable=False)
    # articles = relationship("Article", backref="tags", secondary=article_tag)
    def __repr__(self):
        return "<tag:%s>" % self.name
Base.metadata.drop_all()
Base.metadata.create_all()
article1 = Article(title='article1')
article2 = Article(title='article2')
tag1 = Tag(name='tag1')
tag2 = Tag(name='tag2')
article1.tags.append(tag1)
article1.tags.append(tag2)
article2.tags.append(tag1)
article2.tags.append(tag2)
session.add(article1)
session.add(article2)
session.commit()
article = session.query(Article).first()
print(article.tags)
```

### ORM层面的删除数据：
ORM层面删除数据，会无视mysql级别的外键约束。直接会将对应的数据删除，然后将从表中的那个外键设置为NULL。如果想要避免这种行为，应该将从表中的外键的`nullable=False`。
在SQLAlchemy，只要将一个数据添加到session中，和他相关联的数据都可以一起存入到数据库中了。这些是怎么设置的呢？其实是通过relationship的时候，有一个关键字参数cascade可以设置这些属性： 
1. save-update：默认选项。在添加一条数据的时候，会把其他和他相关联的数据都添加到数据库中。这种行为就是save-update属性影响的。 
2. delete：表示当删除某一个模型中的数据的时候，是否也删掉使用relationship和他关联的数据。
3. delete-orphan：表示当对一个ORM对象解除了父表中的关联对象的时候，自己便会被删除掉。当然如果父表中的数据被删除，自己也会被删除。这个选项只能用在一对多上，不能用在多对多以及多对一上。并且还需要在子模型中的relationship中，增加一个single_parent=True的参数。 
4. merge：默认选项。当在使用session.merge，合并一个对象的时候，会将使用了relationship相关联的对象也进行merge操作。 
5. expunge：移除操作的时候，会将相关联的对象也进行移除。这个操作只是从session中移除，并不会真正的从数据库中删除。 
6. all：是对save-update, merge, refresh-expire, expunge, delete几种的缩写。


### 排序：
1. order_by：可以指定根据这个表中的某个字段进行排序，如果在前面加了一个-，代表的是降序排序。
2. 在模型定义的时候指定默认排序：有些时候，不想每次在查询的时候都指定排序的方式，可以在定义模型的时候就指定排序的方式。有以下两种方式：
    * relationship的order_by参数：在指定relationship的时候，传递order_by参数来指定排序的字段。
    * 在模型定义中，添加以下代码：

     __mapper_args__ = {
         "order_by": title
       }
    即可让文章使用标题来进行排序。
3. 正序排序与倒序排序：默认是使用正序排序。如果需要使用倒序排序，那么可以使用这个字段的`desc()`方法，或者是在排序的时候使用这个字段的字符串名字，然后在前面加一个负号。

### limit、offset和切片操作：
1. limit：可以限制每次查询的时候只查询几条数据。
2. offset：可以限制查找数据的时候过滤掉前面多少条。
3. 切片：可以对Query对象使用切片操作，来获取想要的数据。可以使用`slice(start,stop)`方法来做切片操作。也可以使用`[start:stop]`的方式来进行切片操作。一般在实际开发中，中括号的形式是用得比较多的。希望大家一定要掌握。示例代码如下：
```python
articles = session.query(Article).order_by(Article.id.desc())[0:10]
```

### 懒加载：
在一对多，或者多对多的时候，如果想要获取多的这一部分的数据的时候，往往能通过一个属性就可以全部获取了。比如有一个作者，想要或者这个作者的所有文章，那么可以通过user.articles就可以获取所有的。但有时候我们不想获取所有的数据，比如只想获取这个作者今天发表的文章，那么这时候我们可以给relationship传递一个lazy='dynamic'，以后通过user.articles获取到的就不是一个列表，而是一个AppenderQuery对象了。这样就可以对这个对象再进行一层过滤和排序等操作。
通过`lazy='dynamic'`，获取出来的多的那一部分的数据，就是一个`AppenderQuery`对象了。这种对象既可以添加新数据，也可以跟`Query`一样，可以再进行一层过滤。
总而言之一句话：如果你在获取数据的时候，想要对多的那一边的数据再进行一层过滤，那么这时候就可以考虑使用`lazy='dynamic'`。
lazy可用的选项：
1. `select`：这个是默认选项。还是拿`user.articles`的例子来讲。如果你没有访问`user.articles`这个属性，那么sqlalchemy就不会从数据库中查找文章。一旦你访问了这个属性，那么sqlalchemy就会立马从数据库中查找所有的文章，并把查找出来的数据组装成一个列表返回。这也是懒加载。
2. `dynamic`：这个就是我们刚刚讲的。就是在访问`user.articles`的时候返回回来的不是一个列表，而是`AppenderQuery`对象。


### group_by：
根据某个字段进行分组。比如想要根据性别进行分组，来统计每个分组分别有多少人，那么可以使用以下代码来完成：
```python
session.query(User.gender,func.count(User.id)).group_by(User.gender).all()
```

### having：
having是对查找结果进一步过滤。比如只想要看未成年人的数量，那么可以首先对年龄进行分组统计人数，然后再对分组进行having过滤。示例代码如下：
```python
result = session.query(User.age,func.count(User.id)).group_by(User.age).having(User.age >= 18).all()
```

### join：
1. join分为left join（左外连接）和right join（右外连接）以及内连接（等值连接）。
2. 参考的网页：http://www.jb51.net/article/15386.htm
3. 在sqlalchemy中，使用join来完成内连接。在写join的时候，如果不写join的条件，那么默认将使用外键来作为条件连接。
4. query查找出来什么值，不会取决于join后面的东西，而是取决于query方法中传了什么参数。就跟原生sql中的select 后面那一个一样。
比如现在要实现一个功能，要查找所有用户，按照发表文章的数量来进行排序。示例代码如下：
```python
result = session.query(User,func.count(Article.id)).join(Article).group_by(User.id).order_by(func.count(Article.id).desc()).all()
```

### subquery：
子查询可以让多个查询变成一个查询，只要查找一次数据库，性能相对来讲更加高效一点。不用写多个sql语句就可以实现一些复杂的查询。那么在sqlalchemy中，要实现一个子查询，应该使用以下几个步骤：
1. 将子查询按照传统的方式写好查询代码，然后在`query`对象后面执行`subquery`方法，将这个查询变成一个子查询。
2. 在子查询中，将以后需要用到的字段通过`label`方法，取个别名。
3. 在父查询中，如果想要使用子查询的字段，那么可以通过子查询的返回值上的`c`属性拿到。
整体的示例代码如下：
```python
stmt = session.query(User.city.label("city"),User.age.label("age")).filter(User.username=='李A').subquery()
result = session.query(User).filter(User.city==stmt.c.city,User.age==stmt.c.age).all()
```
