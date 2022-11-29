#!/bin/bash

install_gitlabrunner() {
	apt update -y
	apt install -y docker.io awscli libxml-xpath-perl
	# 安裝gitlab runner
	curl -L --output /usr/local/bin/gitlab-runner https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-linux-amd64
	chmod +x /usr/local/bin/gitlab-runner
	useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash
	gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner
	gitlab-runner start
	gitlab-runner register
	usermod -a -G docker  gitlab-runner
	# 安裝java
	wget  https://otc365-master-codepline.s3-ap-southeast-1.amazonaws.com/java/jdk-8u251-linux-x64.tar.gz
	tar -zxvf jdk-8u251-linux-x64.tar.gz -C /usr/local
	mv  /usr/local/jdk1.8.0_251 /usr/local/jdk1.8.0
	# 安裝mvn
	wget https://ftp.wayne.edu/apache/maven/maven-3/3.8.1/binaries/apache-maven-3.8.1-bin.tar.gz
	tar -zxvf apache-maven-3.8.1-bin.tar.gz -C  /usr/local/
	mv /usr/local/apache-maven-3.8.1 /usr/local/maven-3.8.1
	# 添加 環境變量
	echo 'export JAVA_HOME=/usr/local/jdk1.8.0' >>/etc/profile
	echo 'export PATH=$PATH:/usr/local/jdk1.8.0/bin:/usr/local/maven-3.8.1/bin' >>/etc/profile
	# 安裝 node
	wget https://nodejs.org/dist/v12.18.3/node-v12.18.3-linux-x64.tar.xz
	xz -d node-v12.18.3-linux-x64.tar.xz
	tar -xvf node-v12.18.3-linux-x64.tar -C /usr/local/
	ln -sf  /usr/local/node-v12.18.3-linux-x64/bin/node /usr/bin/
	ln -sf  /usr/local/node-v12.18.3-linux-x64/bin/npm /usr/bin/
	npm install -g npm-cache
	npm install -g @quasar/cli
	ln -sf /usr/local/node-v12.18.3-linux-x64/bin/npm-cache /usr/bin/
	ln -sf /usr/local/node-v12.18.3-linux-x64/bin/quasar /usr/bin/
	ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	# 安裝fluter
	wget https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_2.2.3-stable.tar.xz
	tar xf flutter_linux_2.2.3-stable.tar.xz -C /usr/local
	chown gitlab-runner:gitlab-runner /usr/local/flutter -R
	echo 'export PATH=$PATH:/usr/local/jdk1.8.0/bin:/usr/local/maven-3.8.1/bin' >>/etc/profile
}	



install_openresty(){
	apt update -y
	ln -sf  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	apt-get -y install --no-install-recommends wget gnupg ca-certificates
	wget -O - https://openresty.org/package/pubkey.gpg | sudo apt-key add -
	echo "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/openresty.list
	apt-get update
	apt-get -y install --no-install-recommends openresty
	cp  /usr/bin/openresty /usr/bin/nginx
	cp /etc/init.d/openresty /etc/init.d/nginx
	systemctl enable openresty
}

install_java_test() {
	apt update -y
	ln -sf  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	apt install docker.io  docker-compose awscli gnupg2 pass -y
	usermod -a -G docker  ubuntu
	useradd -m -G docker -s /bin/bash soda
	su - soda 
	mkdir .ssh
	cd .ssh
	echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDMTiKSh2osvy7WDaqJaoCIcCLLisFnSm6F4CKRtv0OVGZiHZarozhYaYgrUIBorhWrINadiKcKqm6uGZQ7CBI9+Yz41E0pjzQWF1KrqdW0oJzdmfcNv+Idz20TEMcbAQplDTgyc2f8h6uzXOv/Ui42aVSvVscgEaXqspm4Io958Hg0WtErUAMRsT7+ymFIQcuU7+2QUah67lU4POcxAURw/VP9CzwxYwNzuS/o7ZQf8icKqU+wFbe1WkGmdTkxGfRrSzSQA+1zZihUe9vtSiiqfnJyuaiV5v94T52F2p6XmlY7a60mmLKgVWjdK10iqWsCbWM7FfIELFnfqdOCZ3xZ soda@ip-192-168-0-18' >> authorized_keys
	chmod 600 authorized_keys
	apt install -y awscli gnupg2 pass
	aws configure
	docker network create --subnet 172.18.0.0/16 mynetwork
}

install_java_prod() {
	apt update -y
	ln -sf  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
	apt install -y docker.io  docker-compose gnupg2 pass awscli
	systemctl enable docker.service
	docker network create --subnet 172.18.0.0/16 mynetwork
}

install zabbix(){
	wget https://repo.zabbix.com/zabbix/5.4/ubuntu/pool/main/z/zabbix-release/zabbix-release_5.4-1+ubuntu18.04_all.deb && dpkg -i zabbix-release_5.4-1+ubuntu18.04_all.deb && apt update -y && apt install zabbix-agent
	usermod -a  -G docker zabbix
	sed -i 's/ServerActive=127.0.0.1/ServerActive=18.141.17.69/g' /etc/zabbix/zabbix_agentd.conf
	ip=$(ifconfig ens5|awk -F 'inet ' '{print $2}'|awk '{print $1}'|xargs)
	sed -i "s/Hostname=Zabbix server/Hostname=ops-$(ifconfig ens5|awk -F 'inet ' '{print $2}'|awk '{print $1}'|xargs)/g" /etc/zabbix/zabbix_agentd.conf
	sudo wget https://raw.githubusercontent.com/xiezhm/ops/master/docker.py -O  /etc/zabbix/docker.py
	sudo chmod +x /etc/zabbix/docker.py
	sudo wget https://raw.githubusercontent.com/xiezhm/ops/master/aws_redis.py -O  /etc/zabbix/aws_redis.py
	sudo chmod +x  /etc/zabbix/aws_redis.py
	sudo wget https://raw.githubusercontent.com/xiezhm/ops/master/linux_zabbix_agent.conf -O  /etc/zabbix/zabbix_agentd.d/linux_zabbix_agent.conf
	sudo systemctl restart zabbix-agent.service
}

sudo systemctl enable docker.service
sudo systemctl enable zabbix-agent.service
sudo systemctl enable openresty
sudo systemctl enable nginx
sudo sed -i 's/1/0/g' /etc/apt/apt.conf.d/10periodic
sudo sed -i 's/1/0/g' /etc/apt/apt.conf.d/20auto-upgrades



install webrtc(){
	~/.acme.sh/acme.sh --issue -d webrtc.Otce.cc --standalone --keylength ec-256 --force --fullchain-file /data/configs/certs/cert.pem --key-file /data/configs/certs/key.pem --register-account -m anna@ynjbzdglhtym.com

}

install docker_templete() {
	sudo useradd -m -G docker -s /usr/sbin/nologin ops && \
	sudo apt update -y && \
	sudo apt install python3-pip -y && \
	sudo pip3 install boto3 && \
	sudo pip3 install --upgrade awscli

}