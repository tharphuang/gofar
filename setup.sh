#!/bin/bash
set -e

echo '1.启动mysql....'
#启动mysql
service mysql start

service mysql status

echo '2.开始导入数据....'
#导入数据
mysql <<EOF
create database if not exists test default character set utf8 collate utf8_general_ci;
use test;
DROP TABLE IF EXISTS user;

CREATE TABLE user (
  id bigint(20) NOT NULL,
  username varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB;

INSERT INTO user (id,username) VALUES (0,'张三');
INSERT INTO user (id,username) VALUES (1,'李四');
EOF

echo '3.导入数据完毕....'

sleep 2
echo service mysql status

#重新设置mysql密码
echo "4.开始修改密码...."
mysql <<EOF
use mysql;
select host, user from user;
create user docker identified by '123456';
grant all on test.* to docker@'%' identified by '123456' with grant option;
flush privileges;
EOF

echo "5.修改密码完毕...."

#sleep 3
echo service mysql status
echo "mysql容器启动完毕,且数据导入成功"

tail -f /dev/null