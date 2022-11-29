

### Received Information

Site Name: mysite.com

Port: 8011

Directory: opt/mysite

### Create Directory

```
mkdir -p /opt/mysite
cd /opt/mysite
echo mysite > index.html
mkdir static
cd static
echo static > index.html
```

### Setup NGINX Server

```
server{
	listen      8011;
	server_name mysite.com;
	
	#access_log  /var/log/nginx/host.access.log  main;

  location /{
    root /opt/mysite;
    index index.html index.htm;
  }

  location /static {
    root   /opt/mysite;
    index  index.html index.htm;
  }
}
```

### Start NGINX Server

```
vim /etc/nginx/conf.d/mysite.conf
nginx -t ## to check the correction
systemctl reload nginx
netstat -anlpt | grep 80 ## grep is the keyword to search
```

### Check Port

``` 
vim /etc/nginx/conf.d/mysite.conf
nginx -t (to check got error or not)
systemctl reload nginx
netstat -anlpt | grep 80
```

### Security Setup

``` 
firewall-cmd --add-port=8011/tcp --permanent
systemctl reload firewalld
```

### Domain Setup

local: host

mac: /etc/hosts

[mysite.com](http://mysite.com) 192.168.10.140

(this command - https://www.dalendesign.com/webpress-blog/webmaster-tools/edit-hosts-file-in-mac-terminal/)

(remember add port number) (ex: 8011)

---

### Check port status

```
ps aux | grep nginx
```

```
netstat -anlpt | grep 80 ## (port num)
```

---

### NGINX virtual machine

#### Setup Virtual Server Based On IP

Setup sites with every machine, 

```
server{
	listen      8011;
	server_name mysite.com;
	
	#access_log  /var/log/nginx/mysite.access.log  main;

  location /{
    root /opt/mysite;
    index index.html index.htm;
  }

  location /static {
    root   /opt/mysite;
    index  index.html index.htm;
  }
}
```

#### Setup Virtual Server Based On Port

##### Create Sites

- site1: /opt/site1
- site2: /opt/site2

##### Create Data

```
mkdir /opt/site1
mkdir /opt/site2

echo "my site 1" > /opt/site1/index.html
echo "my site 2" > /opt/site2/index.html
```

##### Configure Site

```
server{
	listen      8012;
	server_name site1;
	
	#access_log  /var/log/nginx/site1.access.log  main;

  location /{
    root /opt/site1;
    index index.html index.htm;
  }
}
```

```
server{
	listen      8013;
	server_name site2;
	
	#access_log  /var/log/nginx/site2.access.log  main;

  location /{
    root /opt/site2;
    index index.html index.htm;
  }
}
```

#### Setup Virtual Server Based On Domain

- mysite1.com /opt/mysite1
- mysite2.com /opt/mysite2

##### Create Data

```
mkdir /opt/mysite1
mkdir /opt/mysite2

echo mysite1.com > /opt/mysite1/index.html
echo mysite2.com > /opt/mysite2/index.html
```

##### Configure Site

```
server{
	listen      8010;
	server_name mysite1.com;
	
	#access_log  /var/log/nginx/site1.access.log  main;

  location /{
    root /opt/mysite1;
    index index.html index.htm;
  }
}
```

```
server{
	listen      8010;
	server_name mysite2.com;
	
	#access_log  /var/log/nginx/site1.access.log  main;

  location /{
    root /opt/mysite2;
    index index.html index.htm;
  }
}
```

#### Change Hosts File

/etc/hosts

```
192.168.10.140 mysite1.com
192.168.10.140 mysite2.com
```

#### Extra

##### Create Data Site

```
mkdir /opt/newsite
mkdir -p /opt/newsite/site1
mkdir -p /opt/newsite/site2
mkdir -p /opt/newsite/site3

echo newsite > /opt/newsite/index.html
echo newsite-site1 > /opt/newsite/site1/index.html
echo newsite-site2 > /opt/newsite/site2/index.html
echo newsite-site3 > /opt/newsite/site3/index.html
```

```
server{
	listen      8014;
	
	#access_log  /var/log/nginx/site1.access.log  main;

  location /{
    root /opt/newsite;
    index index.html index.htm;
  }
}
```

### Conclusion

Based On IP: http://xxx.xxx.xxx.xxx

Based On Port: http://1.1.1.1:8081 http://1.1.1.1:8082

Based On Domain: http://site1.com http://site2.com

Based On Path: http://mysite1.com/site1 http://mysite1.com/site2

---

### systemctl

```
systemctl reload
systemctl stop
systemctl start
```

---

### Setup Proxy Server

One Machine: Process with 1 Proxy site and 2 other sites

- To secure the other port number

#### Create Data

- proxysite1: /opt/proxy1
- proxysite2: /opt/proxy2

```
mkdir /opt/proxy1
mkdir /opt/proxy2

echo "my proxy site 1" > /opt/proxy1/index.html
echo "my proxy site 2" > /opt/proxy2/index.html
```

#### Create Site

proxy1.conf

```
server{
	listen	8015;
	
	location / {
		root /opt/proxy1;
		index index.html index.htm;
	}
}
```

proxy2.conf

```
server{
	listen	8016;
	
	location / {
		root /opt/proxy2;
		index index.html index.htm;
	}
}
```

proxy.conf

```
server{
	listen	8017;
	
	location / {
		proxy_pass http://192.168.10.140:8016;
	}
}
```

#### Open Port

```
firewall-cmd --add-port=8017/tcp --permanent
systemctl reload firewalld
```

---

### Proxy keyword

```
GET POST PUT DELETE
proxy_method GET;
proxy_set_header Host		   $host;
proxu_set_header X-Real-IP $remote_addr;
proxy_buffer_size 4k;
proxy_buffering on;
proxy_buffers
proxy_busy_buffers_size
proxy_cache ## Define share the cache in the same area

```

---

### Static and dynamic separation

jsp, php

nginx tomcat

nginx + php-fpm

```
server{
	listen      8013;
	server_name site2;
	
	#access_log  /var/log/nginx/site2.access.log  main;

  location /{
    root /opt/site2;
    index index.html index.htm;
  }
}
```

#### Ready for php page

```php
<?php
  	echo phpinfo();
?>
```

#### Setup php-fpm

```
/etc/php/7.4/fpm/pool.d

vim ...
listen: 9000
systemctl restart php7.4-fpm
systemctl status php7.4-fpm
systemctl enable php7.4-fpm
```

```
server{
	listen      8013;
	server_name site2;
	
	#access_log  /var/log/nginx/site2.access.log  main;

  location /{
    root /opt/site2;
    index index.html index.htm;
  }
  
  location ~ \.php$ {
  	root /opt/site2;
    index index.php;
    fastcgi_pass  127.0.0.1:9000;
    fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    fastcgi_param QUERY_STRING    $query_string;
    include fastcgi_params;
  }
}
```

---

### Configuration NGINX to achieve load balancing

Setup to /etc/nginx/nginx.conf

```
http {
	upstream webservers {
		server mysite.com;
    server mysite1.com;
    server mysite2.com;
	}
	
	server {
		location / {
			proxy_pass http://backend;
		}
	}
}
```

Setup to /etc/nginx/conf.d/proxy.conf

```
server{
	listen	8017;
	
	location / {
		proxy_pass http://webservers;
	}
}
```

---

### Nginx http configure

```
worker_processes 1;

events {
	worker_connections    1024;
}
http {
	include         mine.types;
	default_type    application/octet-stream;
	
	sendfile        on;
	
	keepalive_timeout   65;
	
	server {
	listen 80;
	server_name localhost;
	default_type text/html;
		location / { ### all path location then trigger this function
			proxy_pass http://backend;
		}
		
		location = /a { ### www.localhost.com/a but www.localhost.com/a/a not trigger ## highest priority
			echo "=/a";
		}
		
		location ^~ /a { ### start with /a so www.localhost.com/a/a/asd also trigger ## higher priority
			echo "^~ /a"
		}
		
		location ~ /\w { ### third priority
			echo "~ /\w";
		}
	}
}
```

---

### Reverse proxy

```
worker_processes 1;

events {
	worker_connections    1024;
}
http {
	include         mine.types;
	default_type    application/octet-stream;
	
	sendfile        on;
	
	keepalive_timeout   65;
	
	server {
	listen 80;
	server_name localhost;
	default_type text/html;
		location / { 
			proxy_pass http://backend;
		}
	}
}
```

