### To connect aws linux

```
cd Downloads
ssh -i "jian0209.pem" ubuntu@ec2-54-67-108-230.us-west-1.compute.amazonaws.com
ssh -i "/Users/jian0209/Downloads/jian0209.pem" ubuntu@ec2-54-67-108-230.us-west-1.compute.amazonaws.com
```

### Zabbix IP Address

``` json
[
    {
        "Id": "62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc",
        "Created": "2022-04-07T03:03:32.471732832Z",
        "Path": "docker-entrypoint.sh",
        "Args": [
            "mysqld"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 3258,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2022-04-07T03:03:33.718240562Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:667ee8fb158e365450fc3f09712208fe44e9f1364a9b130fed95f3f4862f8a63",
        "ResolvConfPath": "/var/lib/docker/containers/62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc/hostname",
        "HostsPath": "/var/lib/docker/containers/62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc/hosts",
        "LogPath": "/var/lib/docker/containers/62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc/62e4aa45a989a989fd9f05f5eee3b6c34d75994f794616c3b49f3b66016418dc-json.log",
        "Name": "/zabbix-mysql",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/5dc1cfe36ee4613ea38cb0dba704f712f249e1e42bc08a778cd906cfe689ac13-init/diff:/var/lib/docker/overlay2/d1b534d9afe650f2b9888408d55e50420014537fc6d5c5f720e989a4deffe3f7/diff:/var/lib/docker/overlay2/7500a3baafd3d1e7b74e70e66fe8c2586aaf0934083f389f9275ded9c791ce4f/diff:/var/lib/docker/overlay2/85f70f984291c3d7eab4a83847686992dba673513ff4569446c67f378f99adbe/diff:/var/lib/docker/overlay2/cbdb8a50ffedfd364a01e681ea216db029c1bea3cd5b43a7c55077d55a8bc34d/diff:/var/lib/docker/overlay2/0d01ed574af36981ccc16c6a7c4199c06ff319ec1e03feace8ad229f658b5133/diff:/var/lib/docker/overlay2/5b8bbb365959f62ae6ad1eb61c5b4dc715a32a18872d46a4b257446a118532c4/diff:/var/lib/docker/overlay2/d42f9b65102bf3dddd25d89808b8bbc410e824ac30801552ccf4cf55a7c4b757/diff:/var/lib/docker/overlay2/cff7ede561de9353bcd1bb1039106673b32e692b69494b28e48522b1f8e22fd8/diff:/var/lib/docker/overlay2/31e616979fcc4c0c29a105faf1efc2dc629c9b1508a94ebe96c174699e706ccd/diff:/var/lib/docker/overlay2/8bffd18cd03eeb20b7b423c2d53b99b0debbf1bda2ce7a30ab3b61f2b7cad42e/diff:/var/lib/docker/overlay2/fc0a1554573c63d7a9b6f87edc3d28ff2e71f7f27d5f620a812f3b827735f777/diff:/var/lib/docker/overlay2/823e831f445c4d8fbc6b15bffe0975434eec08528f8c83c382b365acbce020de/diff",
                "MergedDir": "/var/lib/docker/overlay2/5dc1cfe36ee4613ea38cb0dba704f712f249e1e42bc08a778cd906cfe689ac13/merged",
                "UpperDir": "/var/lib/docker/overlay2/5dc1cfe36ee4613ea38cb0dba704f712f249e1e42bc08a778cd906cfe689ac13/diff",
                "WorkDir": "/var/lib/docker/overlay2/5dc1cfe36ee4613ea38cb0dba704f712f249e1e42bc08a778cd906cfe689ac13/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [
            {
                "Type": "volume",
                "Name": "27fb4d36f080e6eb9f7f0196e7f3e29741e0add1386ec1f155ecf94078d88f31",
                "Source": "/var/lib/docker/volumes/27fb4d36f080e6eb9f7f0196e7f3e29741e0add1386ec1f155ecf94078d88f31/_data",
                "Destination": "/var/lib/mysql",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
        "Config": {
            "Hostname": "62e4aa45a989",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "3306/tcp": {},
                "33060/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "MYSQL_ROOT_PASSWORD=password",
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "GOSU_VERSION=1.14",
                "MYSQL_MAJOR=8.0",
                "MYSQL_VERSION=8.0.28-1debian10"
            ],
            "Cmd": [
                "mysqld"
            ],
            "Image": "mysql",
            "Volumes": {
                "/var/lib/mysql": {}
            },
            "WorkingDir": "",
            "Entrypoint": [
                "docker-entrypoint.sh"
            ],
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "df97082ca123c3811e0364269480aae3e622e528811be7bac706faa84e281853",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "3306/tcp": null,
                "33060/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/df97082ca123",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "36ab248bd8a4b1811b892ec2416f0d472117236b6d7ee9dcfe227de3081324aa",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "40ed2d259d6d9da8ae84abf7338d11ff9df78d85e1d8e285f1a887df19a0a414",
                    "EndpointID": "36ab248bd8a4b1811b892ec2416f0d472117236b6d7ee9dcfe227de3081324aa",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]

```



