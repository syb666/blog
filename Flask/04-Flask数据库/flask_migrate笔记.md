# flask_migrate笔记：
在实际的开发环境中，经常会发生数据库修改的行为。一般我们修改数据库不会直接手动的去修改，而是去修改ORM对应的模型，然后再把模型映射到数据库中。这时候如果有一个工具能专门做这种事情，就显得非常有用了，而flask-migrate就是做这个事情的。flask-migrate是基于Alembic进行的一个封装，并集成到Flask中，而所有的迁移操作其实都是Alembic做的，他能跟踪模型的变化，并将变化映射到数据库中。

### 安装：
pip install flask-migrate

### 在manage.py中的代码：
```python
from flask_script import Manager
from zhiliao import app
from exts import db
from flask_migrate import Migrate,MigrateCommand

manager = Manager(app)

# 用来绑定app和db到flask_migrate的
Migrate(app,db)
# 添加Migrate的所有子命令到db下
manager.add_command("db",MigrateCommand)


if __name__ == '__main__':
    manager.run()
```

### flask_migrate常用命令：
1. 初始化一个环境：python manage.py db init
2. 自动检测模型，生成迁移脚本：python manage.py db migrate
3. 将迁移脚本映射到数据库中：python manage.py db upgrade
4. 更多命令：python manage.py db --help
