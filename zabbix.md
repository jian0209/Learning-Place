## Running server

```
### If installed as package
shell> service zabbix-server start

### This will work on most of GNU/Linux systems. On other systems you may need to run:
shell> /etc/init.d/zabbix-server start

shell> service zabbix-server stop
shell> service zabbix-server restart
shell> service zabbix-server status

### Start up manually, find the path for zabbix server
shell> zabbix_server
### can use following command
-c --config <file>              path to the configuration file (default is /usr/local/etc/zabbix_server.conf)
-f --foreground                 run Zabbix server in foreground
-R --runtime-control <option>   perform administrative functions
-h --help                       give this help
-V --version                    display version number
shell> zabbix_server -c /usr/local/etc/zabbix_server.conf
shell> zabbix_server --help
shell> zabbix_server -V

### Runtime option
## Example of using runtime control to reload the server configuration cache:
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R config_cache_reload

## Examples of using runtime control to gather diagnostic information:
shell > zabbix_server -R diaginfo
# gather history cache
shell > zabbix_server -R diaginfo=historycache

## Example of using runtime control to reload the SNMP cache:
shell > zabbix_server -R snmp_cache_reload

## Example of using runtime control to trigger execution of housekeeper:
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R housekeeper_execute

## Examples of using runtime control to change log level:
# all process
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R log_level_increase
# second poller process
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R log_level_increase=poller,2
# PID 1234 log level process
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R log_level_increase=1234
# decrease all http poller processes
shell > zabbix_server -c /usr/local/etc/zabbix_server.conf -R log_level_decrease='http poller'

```

​     Server ---> Database ---> Web front end

(all configuration information store in database)

Ex: When create one item in web front end or api, items will be store in database

### Start-up scripts

The scripts are used to automatically start/stop Zabbix processes during system's start-up/shutdown. The scripts are located under directory misc/init.d.

### Server process types

- `alert manager` - alert queue manager
- `alert syncer` - alert DB writer
- `alerter` - process for sending notifications
- `availability manager` - process for host availability updates
- `configuration syncer` - process for managing in-memory cache of configuration data
- `discoverer` - process for discovery of devices
- `escalator` - process for escalation of actions
- `history poller` - process for handling calculated and internal checks requiring a database connection
- `history syncer` - history DB writer
- `housekeeper` - process for removal of old historical data
- `http poller` - web monitoring poller
- `icmp pinger` - poller for icmpping checks
- `ipmi manager` - IPMI poller manager
- `ipmi poller` - poller for IPMI checks
- `java poller` - poller for Java checks
- `lld manager` - manager process of low-level discovery tasks
- `lld worker` - worker process of low-level discovery tasks
- `poller` - normal poller for passive checks
- `preprocessing manager` - manager of preprocessing tasks
- `preprocessing worker` - process for data preprocessing
- `proxy poller` - poller for passive proxies
- `report manager`- manager of scheduled report generation tasks
- `report writer` - process for generating scheduled reports
- `self-monitoring` - process for collecting internal server statistics
- `snmp trapper` - trapper for SNMP traps
- `task manager` - process for remote execution of tasks requested by other components (e.g. close problem, acknowledge problem, check item value now, remote command functionality)
- `timer` - timer for processing maintenances
- `trapper` - trapper for active checks, traps, proxy communication
- `unreachable poller` - poller for unreachable devices
- `vmware collector` - VMware data collector responsible for data gathering from VMware services

---

## Agent

### Passive and active check

Passive Check - agent responds to a data request. Zabbix server asks for data (cpu load), then agent send back result to server.

Active Check - more complex processing. Retrieve a list of data from zabbix server for independent processing. Then periodically send new value to server.

``` 
### If installed as a package
shell > service zabbi-agent start

shell > service zabbix-agent stop
shell > service zabbix-agent restart
shell > service zabbix-agent status

### startup manually
## find the path of zabbix_agentd binary
shell > zabbix_agentd

### agent options

```

### Agent process types

- `active checks` - process for performing active checks
- `collector` - process for data collection
- `listener` - process for listening to passive checks

---

## Proxy

```
### if installed as package
shell > service zabbix-proxy start

shell > service zabbix-proxy stop
shell > service zabbix-proxy restart
shell > service zabbix-proxy status

### start up manually
shell > zabbix_proxy -c /usr/local/etc/zabbix_proxy.conf
shell > zabbix_proxy --help
shell > zabbix_proxy -V

### Runtime Control
## Reload proxy configuration cache
shell > zabbix_proxy -c /usr/local/etc/zabbix_proxy.conf -R config_cache_reload

## gather diagnostic information
# all available diagnostic information
shell > zabbix_proxy -R diaginfo

# history cache statistics
shell > zabbix_proxy -R diaginfo=historycache

## reload SNMP cache
shell > zabbix_proxy -R snmp_cache_reload

## triger execution of housekeeper
shell > zabbix_proxy -c /usr/local/etc/zabbix_proxy.conf -R housekeeper_execute
```

### Proxy process types

- `availability manager` - process for host availability updates
- `configuration syncer` - process for managing in-memory cache of configuration data
- `data sender` - proxy data sender
- `discoverer` - process for discovery of devices
- `heartbeat sender` - proxy heartbeat sender
- `history poller` - process for handling calculated, aggregated and internal checks requiring a database connection
- `history syncer` - history DB writer
- `housekeeper` - process for removal of old historical data
- `http poller` - web monitoring poller
- `icmp pinger` - poller for icmpping checks
- `ipmi manager` - IPMI poller manager
- `ipmi poller` - poller for IPMI checks
- `java poller` - poller for Java checks
- `poller` - normal poller for passive checks
- `preprocessing manager` - manager of preprocessing tasks
- `preprocessing worker` - process for data preprocessing
- `self-monitoring` - process for collecting internal server statistics
- `snmp trapper` - trapper for SNMP traps
- `task manager` - process for remote execution of tasks requested by other components (e.g. close problem, acknowledge problem, check item value now, remote command functionality)
- `trapper` - trapper for active checks, traps, proxy communication
- `unreachable poller` - poller for unreachable devices
- `vmware collector` - VMware data collector responsible for data gathering from VMware services

---

## Host

Zabbix host is the one i want to monitoring machine(switch, vm, servers...)

Creating host is one of the first monitoring task in Zabbix.

---

## Extra

last(/test-172.31.9.150/net.tcp.listen[3306])<>1

last - key

test-172.31.9.150 - hostname

net.tcp.listen[3306] - return 0 or 1

<>1 - not equal to 1