# alembic笔记：

使用alembic的步骤：
1. 定义好自己的模型。
2. 使用alembic创建一个仓库：`alembic init [仓库的名字，推荐使用alembic]`。
3. 修改配置文件：
    * 在`alembic.ini`中，给`sqlalchemy.url`设置数据库的连接方式。这个连接方式跟sqlalchemy的方式一样的。
    * 在`alembic/env.py`中的`target_metadata`设置模型的`Base.metadata`。但是要导入`models`，需要将models所在的路径添加到这个文件中。示例代码如下：
        ```python
        import sys,os
        sys.path.append(os.path.dirname(os.path.dirname(__file__)))
        import moders#导入时一定要将models.py所在的文件夹导入
        ```
4. 将ORM模型生成迁移脚本：`alembic revision --autogenerate -m 'message'`。
5. 将生成的脚本映射到数据库中：`alembic upgrade head`。
6. 以后如果修改了模型，重复4、5步骤。
7. 注意事项：在终端中，如果想要使用alembic，则需要首先进入到安装了alembic的虚拟环境中，不然就找不到这个命令。


### 常用命令：
1. init：创建一个alembic仓库。
2. revision：创建一个新的版本文件。
3. --autogenerate：自动将当前模型的修改，生成迁移脚本。
4. -m：本次迁移做了哪些修改，用户可以指定这个参数，方便回顾。
5. upgrade：将指定版本的迁移文件映射到数据库中，会执行版本文件中的upgrade函数。如果有多个迁移脚本没有被映射到数据库中，那么会执行多个迁移脚本。
6. [head]：代表最新的迁移脚本的版本号。
7. downgrade：会执行指定版本的迁移文件中的downgrade函数。
8. heads：展示head指向的脚本文件版本号。
9. history：列出所有的迁移版本及其信息。
10. current：展示当前数据库中的版本号。

### 经典错误：
1. FAILED: Target database is not up to date.
    * 原因：主要是heads和current不相同。current落后于heads的版本。
    * 解决办法：将current移动到head上。alembic upgrade head
2. FAILED: Can't locate revision identified by '77525ee61b5b'
    * 原因：数据库中存的版本号不在迁移脚本文件中
    * 解决办法：删除数据库的alembic_version表中的数据，重新执行alembic upgrade head
3. 执行`upgrade head`时报某个表已经存在的错误：
    * 原因：执行这个命令的时候，会执行所有的迁移脚本，因为数据库中已经存在了这个表。然后迁移脚本中又包含了创建表的代码。
    * 解决办法：（1）删除versions中所有的迁移文件。（2）修改迁移脚本中创建表的代码。
