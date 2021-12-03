# develop environment

under edge-device-mil folder which cloned from repo:

```[bash]
git clone git@gitlab.com:moxa/ibg/software/platform/thingspro/dlm-agent.git
```

make build environment with debug container:

```[bash]
# armhf
$ ARCH=armhf make debug

# amd64
$ ARCH=amd64 make debug
```

success:

```[bash]
....
docker start dlm-agent-amd64
dlm-agent-amd64
docker attach dlm-agent-amd64
root@d431267a8835:/data# 
```

attach the debug container after building:
```[bash]
# armhf
$ ARCH=armhf make debug

# amd64
$ ARCH=amd64 make debug
```



## feature

### build binary to alter the running device

```[bash]
# armhf
ARCH=armhf make build

# amd64
ARCH=amd64 make build
```

scp the binary to deviced

```[bash]
scp build/armhf/dlmd moxa@192.168.3.127:/tmp
```

stop app

```[bash]
systemctl stop dlm-agent

ssh to your device and change the running binary

```[bash]
cp /tmp/dlmd /usr/sbin/dlmd 
```

start app

```[bash]
systemctl start dlm-agent
```

log

```[bash]
journalctl -u dlm-agent -f -a
```
