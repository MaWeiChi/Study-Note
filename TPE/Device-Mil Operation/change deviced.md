##### method 1

```bash
appman app stop device
docker cp deviced device_app_1:/var/tpdevice/
appman app start device
journalctl APPNAME=device -f
```

##### method 2

```bash
docker cp device_app_1:/usr/sbin/moxa.sh . 
vim moxa.sh  --> 把這個app entry改掉  不要去執行deviecd 改成while true do sleep(1) done
然後把這份moxa.sh docker cp回去container 
appman app restart device 
docker exec -it device_app_1 bash  
```

```bash
#!/bin/sh
#moxa.sh
# set factory configs if first run.
BOOTSTRAP_INIT="/var/tpdevice/module.d/.initialized"
if [ ! -e "$BOOTSTRAP_INIT" ]; then
        echo "$BOOTSTRAP_INIT is Empty, init start ... "
        cp -r /tmp/tpdevice/product.d/* /var/tpdevice/product.d/
        cp -r /tmp/tpdevice/module.d/* /var/tpdevice/module.d/
        touch "$BOOTSTRAP_INIT"
fi

exec tini /var/tpdevice/deviced daemon
```

