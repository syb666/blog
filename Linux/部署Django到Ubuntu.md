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
    
