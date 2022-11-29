```
docker run -d -e env=prod -p 88:80 --name web -h web --restart=always nginx
```

-d = process in back

-p = publish (provide port number)

-e = setup environment (dev | prod)

-name = setup name for the container

-h = hostname

-restart = when open docker then open the container directly

```
docker run -m="500m" --cpus="1" -d nginx
```

-m = memory (m = mb)

-cpus = the number of cpu core

```
docker run -d --name web -p 88:80 -v /opt/wwwroot/:/usr/share/nginx/html nginx
```

above code mean save the code in local, so when next time reopen the container then not be gone. and continue the previous code.