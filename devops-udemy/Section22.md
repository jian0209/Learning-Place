```
wget -O /etc/yum.repos.d/jenkins.repo \ ### -O = / \ = next line / to specific the folder and install
https://pkg.jenkins.io/redhat-stable/jenkins.repo ### 
curl url -o package name ### -o = output / download folder to the current folder
apt install
```

```
### apt repos files located
/etc/apt/sources.list
/etc/apy/source.list.d
```

```
### for ubuntu
wget url...
dpkg -i tree(...).deb

### if want find package
dpkg -l

### remove dpkg package
dpkg -r (package name)

apt update
apt search (package name)
apt install (package name)
apt upgrade
apt remove (package name)
apt purge (package name) ### fully remove
```

