## IP地址

ip 网络简单的概念你们有空要去看看，192.168.1.0/24 10.0.0.0/8  172.16.0.0/16 0.0.0.0/0 127.0.0.1/32 这些常用的网络段你们要了解

---

公网ip类似你家的大门门牌号，私有ip类似你家里房间的门。 公网ip是给别人去你家访问的的地址。私有ip是你去另外一个房间的地址

---

## 业务模型

![](/Users/jian0209/Downloads/image_2022_04_08T04_10_00_738Z.png)

---

Telnet [ip address] - check the connection with ip address

---

sh -v -i "jian0209.pem" ubuntu@ec2-54-153-102-178.us-west-1.compute.amazonaws.com ###### -v一定要在前面

---

```
server {
  listen 30333;
  server_name _;
  location / {
    resolver 127.0.0.53 ipv6=off;
    proxy_set_header Host rpc.polkadot.io;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_ssl_server_name on;
    proxy_pass wss://rpc.polkadot.io;
  }
}
```

---

## 流程

业务访问nginx服务器，nginx》转发到test-wallet-comm02 服务器上钱包服务

一样的操作。业务统一访问nginx服务器，nginx转发到对应的钱包服务上


不同钱包服务启动方式和构建方式、安装方式都有点不同，但是其实都是触类旁通的

理解的业务模型就好理解了

---

## 查文件大小 \ 清理磁盘

```
du -sh
du -h -d1
df (disk free)
df -h (for human seeable)
```

du -sh /data 查看/data 文件大小

/lib/docker/image/overlay2/layerdb/ 这个目录是docker 文件挂载目录

还有一个目录是/var/lib/docker/containers 是docker 运行日志产生的日志文件，这个要可以删除下面的.log文件内容。

还有一个是java 程序产生的日志/home/ubuntu/java/logs  这个目录

找出日志， 磁盘占满多数为日志所致

```
docker  images|grep 203402419275|awk '{print $3}'|xargs docker rmi
### 删除旧的镜像

docker images  是显示当前的镜像，然后grep 过滤2034这个关键字， awk 处理获取docker images id 然后再调用docker rmi 镜像删除
```

---

## 钱包的启动

钱包的启动，一般我安装好了我是在相关目录下会写一个run.sh 的命令， 你们改了以后，要重启或者其他操作，要先kill 原来的程序 然后在执行bash  run.sh 就可以了

---

## 自动化

机器部署

应用程序的部署（发布）

---

## Zabbix 报错

- 无法连接sql
  - 密码错误导致
- 查看user

```
etc/nginx/nginx.conf
### user -> www-data
```

- 查看zabbix-server日志

```
var/log/zabbix
```

- http#######不是https
- 内网不行就用localhost

```
mysql > select * from mysql.user
```

Configuration file "conf/zabbix.conf.php" created.

username: Admin

password: zabbix

---

## 安全性

运维的机器组不能用在业务和钱包上。只能放运维工具，而且安全性要保证，不能随便开放端口对外。

---

开发环境是vultr 开头的那个

测试环境是test-front

---

## Configure for openresty

```
set nu (in vim show number line)
### step of configure openresty
Configure in .conf folder
openresty -t (to check)
openresty -s reload (restart openresty)
```

---

## 日志流程

loki是数据库存储服务

日志流程是nginx 产生日志。promtail 收集日志，loki存储 ，grafana 展示

---

## 显示监听

```
netstat -tupln
```

---

