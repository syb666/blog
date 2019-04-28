## 部署项目时使用的知识点
1. 在刚进系统时需要先更新一下资源，要不然会出现找不到资源。使用sudo apt update进行更新
2. 安装完memcached之后可以使用telnet 127.0.0.1 11211 进行测试。
    ```
    set username 0 60 7 
    zhiliao
    STORED
    get username
    VALUE username 0 7
    zhiliao
    END 
    quitr^H    
    ERROR
    quit
    ```
3. 当我在Xshell中修改完一部分代码时，当我再向码云中拉取代码时，会出现问题。这时我只需要```git add``` .然后再提交一下 ```git commit -m 'nothing'```然后再根据提示
输一下邮箱和用户名。然后再```git pull origin master```
4. 查看运行文件```ps aux|grep uwsgi``` kill-9 端口号
5. 退出虚拟环境deactivate
6. 在使用uwsgi时会出现一些问题，出现过主urls找不到，然后使用python manage.py runserver 0.0.0.0:8000发现端口占用，
然后kill之后再进行uwsgi --ini uwsgi.ini,就可以使用了。
7. 在运行```supervisord -c supervisor.conf```出现下面的错误时,只需要将supervisor.conf中的第一行注释去掉就好了
    ```
    Error: File contains no section headers.
    file: '/etc/supervisor/conf.d/imgweb.conf', line: 1
    '；/etc/supervisor/conf.d/imgweb.conf\n'
    For help, use /home/ejior/.virtualenvs/begin/bin/supervisord -h
    ```

8. 在向码云上传项目和拉取项目时常用的git命令如下
    ```
    初始化git init
    提交到缓存git commit -m '注释'
    提交项目git push origin master
    拉取项目git pull origin master
    ```
9. 当部署完数据库时，要使用本地的Navicat连接服务器的mysql时，需要将mysql里的user为root中的Host 改为%，然后一定一定要记得重启数据库.
   ```
   mysql> USE mysql; -- 切换到 mysql DB
   Database changed
   mysql> SELECT User, Password, Host FROM user; -- 查看现有用户,密码及允许连接的主机
   +------+----------+-----------+
   | User | Password | Host      |
   +------+----------+-----------+
   | root |          | localhost |
   +------+----------+-----------+
   row in set (0.00 sec)
   如果没有host创建host
   如果没有"%"这个host值,就执行下面这两句:
   mysql> update user set host='%' where user='root';
   mysql> flush privileges;
   mysql> SELECT User, Password, Host FROM user; -- 查看现有用户,密码及允许连接的主机
   +------+----------+-----------+
   | User | Password | Host      |
   +------+----------+-----------+
   | root |          |   %       |
   +------+----------+-----------+
   授权用户
   任意主机以用户root和密码mypwd连接到mysql服务器
   mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'mypwd' WITH GRANT OPTION;
   mysql> flush privileges;
   一定要记得重启数据库。`service mysql restart` 
   ```
10. 在本地项目中修改为```DEBUG = False```时导致整个页面的静态文件加载不出来，修改为```DEBUG = True```时静态文件成功加载

11. 在部署到服务器之后，登录账户时出现403错误，然后只需要在注册登录页面加上下面的代码就行了，这里就是设置一下csrf值。
    ```
    <script>
            if(getCookie('csrftoken')){
            }else{
                $.ajax({
                    data:{csrfmiddlewaretoken:'{{ csrf_token }}'},
                })
            }
        </script>
    ```
12. 当时在部署完项目时，可以使用公网ip访问，但是使用域名一直访问不了。老是出错显示让我把域名加入到```ALLOWED_HOSTS```这里面
但是加上完后，重启了nginx还是不行，最后重启supervisor才可以。对于每次修改参数后一定要记得重启supervisor。```supervisord -c supervisor.conf```.
13. 在修改IP访问者的信息进行分页时，修改完信息后。在服务器上不能显示修改的东西，这时重新启动一下uwsgi就好啦。```uwsgi --ini uwsgi.ini```。因为我们可以使用supervisor管理uwsgi，这时我们可以通过supervisor重新启动一下uwsgi。
    * status # 查看状态
    * start program_name #启动程序
    * restart program_name #重新启动程序
    * stop program_name # 关闭程序
    * reload # 重新加载配置文件
    * quit # 退出控制台


