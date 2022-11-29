```
top ### the command for showing the whole process in cpu
ps -ef ## show non-utilization process
kill (...) ### to remove the process
kill -9 (...) ### remove only the specific process
ps -ef | grep (...) | grep -v 'grep' | awk '{print $2}'
```

