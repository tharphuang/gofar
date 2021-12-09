FROM mysql:5.7.32
EXPOSE 3306
LABEL version="0.1" description="Mysql服务器" by="don"

#设置免密登录
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

#将所需文件放到容器中
#拷贝安装脚本
COPY setup.sh /mysql/setup.sh 

#设置容器启动时执行的命令
CMD ["/mysql/setup.sh"]
ENTRYPOINT ["/bin/bash"]