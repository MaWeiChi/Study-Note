```bash
[manager] manager.go:281: [Update Device DB] [{"type":"gps","mode":"auto","interface":"/dev/ttyACM0","location":{"lat":20.086791,"lng":120.434348167},"capabilities":{"interface":["/dev/ttyACM0"]}}]
```



```bash
redis-c
bash: redis-c: command not found
root@Moxa:/home/moxa# redis-cli
127.0.0.1:6379> KEYS *
 1) "taghub:monitor:system/status/systemDiskFree"
 2) "sysman:hardware:time:1"
 3) "sysman:app:metadata:tagservice"
 4) "taghub:monitor:system/status/systemDiskPercent"
 5) "sysman:api:system:sshserver"
 6) "sysman:app:metadata:edge-web"
 7) "sysman:api:system:httpserver"
 8) "sysman:api:system:log"
 9) "sysman:hardware:dhcpservers:3"
10) "taghub:monitor:system/status/networkTx"
11) "sysman:hardware:dhcpservers:2"
12) "taghub:monitor:system/network/networkUsage)system/network/networkTx)system/network/networkRx)"
13) "sysman:hardware:iptables:1"
14) "taghub:monitor:system/status/cpuUsage"
15) "sysman:hardware:ethernets:1"
16) "sysman:hardware:ethernets:2"
17) "taghub:monitor:system/status/sd1-1DiskFree"
18) "sysman:hardware:serials:1"
19) "sysman:app:metadata:device"
20) "taghub:monitor:system/status/networkUsage"
21) "taghub:monitor:system/status/gpsSignal"
22) "sysman:hardware:general:1"
23) "taghub:monitor:system/status/cellular1Rssi"
24) "sysman:app:metadata:modbusmaster"
25) "sysman:hardware:gps:1"
26) "taghub:monitor:system/status/gps"
27) "sysman:hardware:wifi:1"
28) "sysman:resman:usage:system:disk"
29) "sysman:app:runtime:edge-web"
30) "sysman:hardware:minicards:1"
31) "sysman:hardware:wan:1"
32) "sysman:hardware:serials:2"
33) "sysman:resman:resource:resource"
34) "sysman:app:runtime:tagservice"
35) "sysman:api:system:discovery"
36) "sysman:hardware:firewall:1"
37) "sysman:app:metadata:cloud"
38) "taghub:monitor:system/status/sd1-1DiskPercent"
39) "sysman:app:runtime:cloud"
40) "sysman:app:runtime:device"
41) "taghub:monitor:system/status/memoryUsage"
42) "sysman:hardware:dhcpservers:1"
43) "sysman:hardware:cellulars:1"
44) "taghub:monitor:system/status/sd1-1DiskUsed"
45) "sysman:app:runtime:modbusmaster"
46) "sysman:app:runtime:modbusslave"
47) "sysman:api:system:serialconsole"
48) "taghub:monitor:system/status/networkRx"
49) "sysman:hardware:ethernets:3"
50) "taghub:monitor:system/status/systemDiskUsed"
51) "taghub:monitor:system/status/cellular1Signal"
52) "taghub:monitor:system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)"
53) "sysman:app:metadata:modbusslave"
54) "sysman:hardware:route:1"
127.0.0.1:6379> GET system:hadware:gps:1
(nil)
127.0.0.1:6379> GET system:hardware:gps:1
(nil)
127.0.0.1:6379> GET sysman:hardware:gps:1
"{\"capabilities\":{\"interface\":[\"/dev/ttyACM0\"]},\"id\":1,\"interface\":\"/dev/ttyACM0\",\"location\":{\"lat\":20.086791,\"lng\":120.434348167},\"mode\":\"auto\",\"type\":\"gps\"}"
127.0.0.1:6379> sysman:hardware:*
root@Moxa:/home/moxa# ^C
root@Moxa:/home/moxa# redis-c^C
root@Moxa:/home/moxa# sysman:hardware:*

```

