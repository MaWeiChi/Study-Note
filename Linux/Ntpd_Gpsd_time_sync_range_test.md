```    go
const (
// NTP_GPS_SOURCE_CONFIG_WITH_GATEKEEPER - Set a sync range by time2 (unit:sec, default 11400 sec, 4 hours)
NTP_GPS_SOURCE_CONFIG_WITH_GATEKEEPER = `# GPS Serial data reference 
server 127.127.28.0 minpoll 4 maxpoll 4
fudge 127.127.28.0 time1 0.0 time2 28800.0 refid GPS`
)


```

### First log

```bash
-- Logs begin at Wed 2020-05-20 13:29:55 CST. --
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:32:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602093056,"used":999700480,"usedPercent":15.142861929088964,"inodesTotal":1607808,"inodesUsed":35097,"inodesFree":1572711,"inodesUsedPercent":2.1829098996895153}
May 20 13:33:12 Moxa device_app_1[1243]:
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |      3.9165ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |    4.067375ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:33:12 Moxa device_app_1[1243]: INFO[ip] commander.go:38: exec: [ip a s eth0]
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |     1.37825ms |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |    6.492875ms |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |      579.75µs |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |    1.455625ms |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:33:12 Moxa device_app_1[1243]: INFO[ip] commander.go:38: exec: [ip a s eth1]
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |     303.125µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |     633.625µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:33:12 Moxa device_app_1[1243]: INFO[ip] commander.go:38: exec: [ip a s wlan0]
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |    51.19925ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:33:12 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:12 | 200 |   52.285625ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:33:12 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:33:16 Moxa device_app_1[1243]: INFO[cellular 1] commander.go:140: Non-zero exit code: signal: killed
May 20 13:33:16 Moxa device_app_1[1243]: INFO[cellular 1] commander.go:141: stderr: /usr/sbin/cell_mgmt: line 1505: mx-module-ctl: command not found
May 20 13:33:16 Moxa device_app_1[1243]:
May 20 13:33:19 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"available":true,"status":"nosim","currentProfileId":1,"operatorName":"","name":"usb0","module":"u-blox TOBY-L2 series","imei":"352255062111038","iccid":"","imsi":"","pinRetryRemain":0,"mac":"02:01:05:19:00:0b","rat":"","signal":{"rat":"","csq":0,"rssi":0,"rxqual":0,"rscp":0,"ecio":0,"rsrp":0,"rsrq":0,"indicator":"","level":0},"type":"cellulars","wan":true,"displayName":"Cellular1","capabilities":{"sim":1},"id":1,"profileTimeout":60,"enable":true,"keepalive":{"intervalSec":60,"enable":true,"targetHost":"8.8.8.8"},"autoDetect":false,"profiles":[{"id":1,"name":"SIM1","pinCode":"0000","init":["sim:1"],"pdpContext":{"id":1,"type":"ipv4","static":true,"apn":"internet","auth":{"protocol":"none","username":"","password":""}}}]}]
May 20 13:33:19 Moxa device_app_1[1243]: time="2020-05-20T13:33:19+08:00" level=warning msg="Cellular set PIN code fail. PIN: 0000" category=network event="cellular set PIN code fail" origin="Device App" user=
May 20 13:33:19 Moxa device_app_1[1243]: INFO[route] commander.go:38: exec: [ip route show 0.0.0.0/0]
May 20 13:33:19 Moxa device_app_1[1243]: INFO[route] dns.go:20: Current DNS:
May 20 13:33:19 Moxa device_app_1[1243]: INFO[route] dns.go:21:
May 20 13:33:19 Moxa device_app_1[1243]: INFO[route] commander.go:38: exec: [cp /etc/resolv.conf /host/etc/resolv.conf]
May 20 13:33:19 Moxa device_app_1[1243]: INFO[cellulars] profile.go:61: [1:1] stopped
May 20 13:33:19 Moxa device_app_1[1243]: INFO[route] commander.go:38: exec: [ip route del default]
May 20 13:33:20 Moxa device_app_1[1243]: time="2020-05-20T13:33:20+08:00" level=warning msg="No default route" category=network event="uplink change" origin="Device App" user=
May 20 13:33:20 Moxa device_app_1[1243]: INFO[route] commander.go:49: exec return code : 2,
May 20 13:33:20 Moxa device_app_1[1243]: error : exit status 2
May 20 13:33:20 Moxa device_app_1[1243]: output : RTNETLINK answers: No such process
May 20 13:33:20 Moxa device_app_1[1243]:
May 20 13:33:20 Moxa device_app_1[1243]:
May 20 13:33:20 Moxa device_app_1[1243]: INFO[route] route.go:72: current default route is not found
May 20 13:33:20 Moxa device_app_1[1243]:
May 20 13:33:20 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"type":"wan","name":"","displayName":"","ip":"","netmask":"","gateway":"","dns":[]}]
May 20 13:33:49 Moxa device_app_1[1243]: INFO[time] time.go:79: Gps enable
May 20 13:33:49 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 13:33:49 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 13:33:49 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":300}}]
May 20 13:33:49 Moxa device_app_1[1243]: time="2020-05-20T13:33:49+08:00" level=info msg="Time request on {\"ntp\": {\"enable\": true, \"source\": \"gps\", \"server\": \"pool.ntp.org\", \"interval\": 300}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 13:33:49 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:49 | 200 |   100.96275ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:33:49 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:33:49 | 200 |  101.153625ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:33:50 Moxa device_app_1[1243]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:33:50 Moxa device_app_1[1243]: time="2020-05-20T13:33:50+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:33:50 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:34:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:34:01 | 200 |      599.75µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:34:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:34:01 | 200 |         770µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |     619.625µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |       823.5µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:35:09 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602093056,"used":999700480,"usedPercent":15.142861929088964,"inodesTotal":1607808,"inodesUsed":35097,"inodesFree":1572711,"inodesUsedPercent":2.1829098996895153}
May 20 13:35:09 Moxa device_app_1[1243]:
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |    3.908125ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |    4.057375ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |     364.875µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |      846.25µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |     3.55725ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:35:09 | 200 |    3.750625ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:35:09 Moxa device_app_1[1243]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:35:14 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T01:35:09.000+08:00]
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] time.go:39: Update Time: 2020-05-20T01:35:09.000+08:00
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] time.go:136: Time updated to 2020-05-20T01:35:09.000+08:00
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 01:35:09 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 01:35:09 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":300}}]
May 20 01:35:09 Moxa device_app_1[1243]: time="2020-05-20T01:35:09+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"2020-05-20T13:33:50+08:00\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":300,\"server\":\"\"},\"time\":\"2020-05-20T01:35:09.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 01:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:35:09 | 200 |    956.7785ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 01:35:09 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:35:09 | 200 |  957.156625ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 01:35:51 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 01:35:51 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 01:35:51 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 01:35:51 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 01:35:51 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":300}}]
May 20 01:35:51 Moxa device_app_1[1243]: time="2020-05-20T01:35:51+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":300,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 01:35:51 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:35:51 | 200 |   412.21825ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 01:35:51 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:35:51 | 200 |  412.451625ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 01:36:13 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602088960,"used":999704576,"usedPercent":15.142923972834765,"inodesTotal":1607808,"inodesUsed":35101,"inodesFree":1572707,"inodesUsedPercent":2.1831586856142025}
May 20 01:36:13 Moxa device_app_1[1243]:
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |    4.287125ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |     4.43425ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |      380.75µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |         730µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |      464.25µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |      746.25µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |     6.69575ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 01:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 01:36:13 | 200 |    6.872125ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 01:36:21 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 01:36:21 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 01:35:52 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:53 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:54 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:55 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:56 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:57 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:58 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:35:59 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:00 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:01 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:02 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:03 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:04 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:05 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:06 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:07 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:08 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:09 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:10 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:11 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:12 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:13 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:14 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:15 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:16 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:17 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:18 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:19 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]: 20 May 01:36:20 ntpd[1395]: SHM: difference limit exceeded, delta=43206s
May 20 01:36:21 Moxa device_app_1[1243]:
May 20 01:36:21 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 01:36:21 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 01:36:50 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T07:36:13.000+08:00]
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] time.go:39: Update Time: 2020-05-20T07:36:13.000+08:00
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] time.go:136: Time updated to 2020-05-20T07:36:13.000+08:00
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 07:36:13 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 07:36:13 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":300}}]
May 20 07:36:13 Moxa device_app_1[1243]: time="2020-05-20T07:36:13+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":300,\"server\":\"\"},\"time\":\"2020-05-20T07:36:13.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 07:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 07:36:13 | 200 |  976.581875ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 07:36:13 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 07:36:13 | 200 |     976.766ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 07:36:14 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 07:36:14 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 01:36:22 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:23 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:24 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:25 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:26 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:27 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:28 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:29 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:30 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:31 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:32 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:33 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:34 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:35 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:36 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:37 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:38 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:39 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:40 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:41 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:42 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:43 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:44 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:45 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:46 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:47 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:48 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 01:36:49 ntpd[1398]: SHM: difference limit exceeded, delta=43206s
May 20 07:36:14 Moxa device_app_1[1243]: 20 May 07:36:13 ntpd[1398]: SHM: stale/bad receive time, delay=21564s
May 20 07:36:14 Moxa device_app_1[1243]:
May 20 07:36:14 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 07:36:14 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 13:36:57 Moxa device_app_1[1243]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:36:57 Moxa device_app_1[1243]: time="2020-05-20T13:36:57+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:36:57 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:37:10 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T08:36:13.000+08:00]
May 20 08:36:13 Moxa device_app_1[1243]: INFO[time] time.go:39: Update Time: 2020-05-20T08:36:13.000+08:00
May 20 08:36:13 Moxa device_app_1[1243]: INFO[time] time.go:136: Time updated to 2020-05-20T08:36:13.000+08:00
May 20 08:36:13 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 08:36:14 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 08:36:14 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 08:36:14 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 08:36:14 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":300}}]
May 20 08:36:14 Moxa device_app_1[1243]: time="2020-05-20T08:36:14+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":300,\"server\":\"\"},\"time\":\"2020-05-20T08:36:13.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 08:36:14 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 08:36:14 | 200 |   1.60645375s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 08:36:14 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 08:36:14 | 200 |    1.6067385s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 08:36:20 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 08:36:20 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 08:36:20 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 08:36:20 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 08:36:20 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 08:36:20 Moxa device_app_1[1243]: time="2020-05-20T08:36:20+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 08:36:20 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 08:36:20 | 200 |     160.114ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 08:36:20 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 08:36:20 | 200 |  160.303375ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:37:19 Moxa device_app_1[1243]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:37:19 Moxa device_app_1[1243]: time="2020-05-20T13:37:19+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:37:19 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |      3.7605ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |    3.970625ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |      492.25µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |     747.875µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |         179µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |     370.875µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:37:34 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602067456,"used":999726080,"usedPercent":15.143249702500238,"inodesTotal":1607808,"inodesUsed":35101,"inodesFree":1572707,"inodesUsedPercent":2.1831586856142025}
May 20 13:37:34 Moxa device_app_1[1243]:
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |     4.18275ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:37:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:37:34 | 200 |      4.3225ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:37:52 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T03:01:34.000+08:00]
May 20 03:01:34 Moxa device_app_1[1243]: INFO[time] time.go:39: Update Time: 2020-05-20T03:01:34.000+08:00
May 20 03:01:34 Moxa device_app_1[1243]: INFO[time] time.go:136: Time updated to 2020-05-20T03:01:34.000+08:00
May 20 03:01:34 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 03:01:35 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 03:01:35 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 03:01:35 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 03:01:35 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 03:01:35 Moxa device_app_1[1243]: time="2020-05-20T03:01:35+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"2020-05-20T13:37:19+08:00\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T03:01:34.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 03:01:35 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 03:01:35 | 200 |  1.907757875s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 03:01:35 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 03:01:35 | 200 |  1.908041625s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 03:01:38 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 03:01:38 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 03:01:39 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 03:01:39 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 03:01:39 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 03:01:39 Moxa device_app_1[1243]: time="2020-05-20T03:01:39+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 03:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 03:01:39 | 200 |   1.29606375s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 03:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 03:01:39 | 200 |  1.296352625s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 03:02:09 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 03:02:09 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 03:01:40 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:41 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:42 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:43 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:44 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:45 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:46 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:47 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:48 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:49 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:50 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:51 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:52 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:53 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:54 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:55 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:56 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:57 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:58 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:01:59 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:00 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:01 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:02 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:03 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:04 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:05 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:06 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:07 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]: 20 May 03:02:08 ntpd[1416]: SHM: difference limit exceeded, delta=38179s
May 20 03:02:09 Moxa device_app_1[1243]:
May 20 03:02:09 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 03:02:09 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 03:02:33 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T06:01:34.000+08:00]
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] time.go:39: Update Time: 2020-05-20T06:01:34.000+08:00
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] time.go:136: Time updated to 2020-05-20T06:01:34.000+08:00
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 06:01:34 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 06:01:34 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 06:01:34 Moxa device_app_1[1243]: time="2020-05-20T06:01:34+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T06:01:34.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 06:01:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:34 | 200 |    934.2165ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 06:01:34 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:34 | 200 |    934.3925ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:53: status code: 200
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 06:01:36 ntpd[1424]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 06:01:36 Moxa device_app_1[1243]:
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 06:01:36 ntpd[1426]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 06:01:36 Moxa device_app_1[1243]:
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 06:01:36 Moxa device_app_1[1243]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 06:01:36 ntpd[1429]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 06:01:36 Moxa device_app_1[1243]:
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 06:01:36 Moxa device_app_1[1243]: INFO[time] ntp.go:88: NTP update failed
May 20 06:01:36 Moxa device_app_1[1243]: time="2020-05-20T06:01:36+08:00" level=warning msg="NTP Time Sync Failed" category="device setting" event="time sync failed" origin="Device App" user=
May 20 06:01:37 Moxa device_app_1[1243]: time="2020-05-20T06:01:36+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 06:01:37 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:37 | 200 |  173.928375ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 06:01:37 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:37 | 200 |  174.221875ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |     506.375µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |     691.625µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 06:01:39 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602067456,"used":999726080,"usedPercent":15.143249702500238,"inodesTotal":1607808,"inodesUsed":35101,"inodesFree":1572707,"inodesUsedPercent":2.1831586856142025}
May 20 06:01:39 Moxa device_app_1[1243]:
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |     3.82475ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |     4.14175ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |     135.125µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |      318.25µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |       4.151ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 06:01:39 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 06:01:39 | 200 |    4.549625ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 06:01:40 Moxa device_app_1[1243]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 06:01:40 Moxa device_app_1[1243]: INFO[time] commander.go:141: stderr: 20 May 03:02:10 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:11 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:12 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:13 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:14 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:15 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:16 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:17 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:18 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:19 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:20 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:21 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:22 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:23 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:24 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:25 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:26 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:27 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:28 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:29 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:30 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:31 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 03:02:32 ntpd[1419]: SHM: difference limit exceeded, delta=38179s
May 20 06:01:40 Moxa device_app_1[1243]: 20 May 06:01:34 ntpd[1419]: SHM: stale/bad receive time, delay=10741s
May 20 06:01:40 Moxa device_app_1[1243]:
May 20 06:01:40 Moxa device_app_1[1243]: INFO[time] ntp.go:77: NTP update timeout
May 20 06:01:40 Moxa device_app_1[1243]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 13:38:59 Moxa device_app_1[1243]: time="2020-05-20T13:38:59+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:38:59 Moxa device_app_1[1243]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:38:59 Moxa device_app_1[1243]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |      499.75µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |      678.25µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |      191.25µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |     382.375µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |     3.31375ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |    3.520875ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:39:01 Moxa device_app_1[1243]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602067456,"used":999726080,"usedPercent":15.143249702500238,"inodesTotal":1607808,"inodesUsed":35101,"inodesFree":1572707,"inodesUsedPercent":2.1831586856142025}
May 20 13:39:01 Moxa device_app_1[1243]:
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |     3.57225ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:39:01 Moxa device_app_1[1243]: [GIN] 2020/05/20 - 13:39:01 | 200 |    3.703375ms |   192.168.4.111 | GET      /api/v1/device/general
packet_write_wait: Connection to 192.168.4.127 port 22: Broken pipe
```





### Second log

```bash
-- Logs begin at Wed 2020-05-20 13:43:17 CST. --
May 20 13:45:20 Moxa device_app_1[1242]: INFO[route] commander.go:38: exec: [ip route del default]
May 20 13:45:20 Moxa device_app_1[1242]: INFO[route] commander.go:49: exec return code : 2,
May 20 13:45:20 Moxa device_app_1[1242]: error : exit status 2
May 20 13:45:20 Moxa device_app_1[1242]: output : RTNETLINK answers: No such process
May 20 13:45:20 Moxa device_app_1[1242]:
May 20 13:45:20 Moxa device_app_1[1242]:
May 20 13:45:20 Moxa device_app_1[1242]: INFO[route] route.go:72: current default route is not found
May 20 13:45:20 Moxa device_app_1[1242]:
May 20 13:45:20 Moxa device_app_1[1242]: time="2020-05-20T13:45:20+08:00" level=warning msg="No default route" category=network event="uplink change" origin="Device App" user=
May 20 13:45:20 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"type":"wan","name":"","displayName":"","ip":"","netmask":"","gateway":"","dns":[]}]
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |       460.5µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |      693.25µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:45:51 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s eth0]
May 20 13:45:51 Moxa device_app_1[1242]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602065408,"used":999728128,"usedPercent":15.14328072437314,"inodesTotal":1607808,"inodesUsed":35099,"inodesFree":1572709,"inodesUsedPercent":2.183034292651859}
May 20 13:45:51 Moxa device_app_1[1242]:
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |      24.729ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |   27.239125ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |      4.7905ms |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |     5.15925ms |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:45:51 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s eth1]
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |      406.25µs |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |     609.625µs |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:45:51 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s wlan0]
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |     68.1615ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:45:51 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:51 | 200 |    68.47575ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:51 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |         588µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |         765µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |    4.172875ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |     4.37475ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:45:54 Moxa device_app_1[1242]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602065408,"used":999728128,"usedPercent":15.14328072437314,"inodesTotal":1607808,"inodesUsed":35099,"inodesFree":1572709,"inodesUsedPercent":2.183034292651859}
May 20 13:45:54 Moxa device_app_1[1242]:
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |     4.15325ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |    4.315875ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |     348.125µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:45:54 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:45:54 | 200 |     627.375µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:54 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:45:55 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:46:06 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T23:45:54.000+08:00]
May 20 23:45:54 Moxa device_app_1[1242]: INFO[time] time.go:39: Update Time: 2020-05-20T23:45:54.000+08:00
May 20 23:45:54 Moxa device_app_1[1242]: INFO[time] time.go:136: Time updated to 2020-05-20T23:45:54.000+08:00
May 20 23:45:54 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 23:45:55 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 23:45:55 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 23:45:55 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 23:45:55 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 23:45:55 Moxa device_app_1[1242]: time="2020-05-20T23:45:55+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T23:45:54.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 23:45:55 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:45:55 | 200 |  1.956171125s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:45:55 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:45:55 | 200 |   1.95636075s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:45:59 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 23:45:59 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 23:45:59 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 23:45:59 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:45:59 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 23:45:59 Moxa device_app_1[1242]: time="2020-05-20T23:45:59+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 23:45:59 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:45:59 | 200 |     152.403ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:45:59 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:45:59 | 200 |   152.70675ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:46:24 Moxa device_app_1[1242]: INFO[cellular 1] commander.go:140: Non-zero exit code: signal: killed
May 20 23:46:24 Moxa device_app_1[1242]: INFO[cellular 1] commander.go:141: stderr: /usr/sbin/cell_mgmt: line 1505: mx-module-ctl: command not found
May 20 23:46:24 Moxa device_app_1[1242]:
May 20 23:46:27 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"available":true,"status":"nosim","currentProfileId":1,"operatorName":"","name":"usb0","module":"u-blox TOBY-L2 series","imei":"352255062111038","iccid":"","imsi":"","pinRetryRemain":0,"mac":"02:01:05:19:00:0b","rat":"","signal":{"rat":"","csq":0,"rssi":0,"rxqual":0,"rscp":0,"ecio":0,"rsrp":0,"rsrq":0,"indicator":"","level":0},"type":"cellulars","wan":true,"displayName":"Cellular1","capabilities":{"sim":1},"id":1,"profileTimeout":60,"enable":true,"keepalive":{"intervalSec":60,"enable":true,"targetHost":"8.8.8.8"},"autoDetect":false,"profiles":[{"id":1,"name":"SIM1","pinCode":"0000","init":["sim:1"],"pdpContext":{"id":1,"type":"ipv4","static":true,"apn":"internet","auth":{"protocol":"none","username":"","password":""}}}]}]
May 20 23:46:27 Moxa device_app_1[1242]: time="2020-05-20T23:46:27+08:00" level=warning msg="Cellular set PIN code fail. PIN: 0000" category=network event="cellular set PIN code fail" origin="Device App" user=
May 20 23:46:27 Moxa device_app_1[1242]: INFO[cellulars] profile.go:61: [1:1] stopped
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] commander.go:38: exec: [ip route show 0.0.0.0/0]
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] dns.go:20: Current DNS:
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] dns.go:21:
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] commander.go:38: exec: [cp /etc/resolv.conf /host/etc/resolv.conf]
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] commander.go:38: exec: [ip route del default]
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] commander.go:49: exec return code : 2,
May 20 23:46:27 Moxa device_app_1[1242]: error : exit status 2
May 20 23:46:27 Moxa device_app_1[1242]: output : RTNETLINK answers: No such process
May 20 23:46:27 Moxa device_app_1[1242]:
May 20 23:46:27 Moxa device_app_1[1242]:
May 20 23:46:27 Moxa device_app_1[1242]: INFO[route] route.go:72: current default route is not found
May 20 23:46:27 Moxa device_app_1[1242]:
May 20 23:46:27 Moxa device_app_1[1242]: time="2020-05-20T23:46:27+08:00" level=warning msg="No default route" category=network event="uplink change" origin="Device App" user=
May 20 23:46:27 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"type":"wan","name":"","displayName":"","ip":"","netmask":"","gateway":"","dns":[]}]
May 20 23:46:29 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 23:46:29 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:46:00 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:01 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:02 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:03 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:04 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:05 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:06 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:07 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:08 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:09 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:10 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:11 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:12 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:13 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:14 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:15 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:16 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:17 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:18 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:19 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:20 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:21 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:22 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:23 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:24 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:25 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:26 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:27 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]: 20 May 23:46:28 ntpd[1285]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:29 Moxa device_app_1[1242]:
May 20 23:46:29 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:46:29 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:46:30 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:31 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:32 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:33 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:34 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:35 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:36 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:37 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:38 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:39 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:40 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:41 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:42 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:43 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:44 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:45 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:46 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:47 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:48 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:49 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:50 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:51 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:52 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:53 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:54 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:55 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:56 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:57 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]: 20 May 23:46:58 ntpd[1363]: SHM: difference limit exceeded, delta=35987s
May 20 23:46:59 Moxa device_app_1[1242]:
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:46:59 ntpd[1365]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 23:46:59 Moxa device_app_1[1242]:
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:46:59 ntpd[1369]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 23:46:59 Moxa device_app_1[1242]:
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:46:59 Moxa device_app_1[1242]: time="2020-05-20T23:46:59+08:00" level=warning msg="NTP Time Sync Failed" category="device setting" event="time sync failed" origin="Device App" user=
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:46:59 ntpd[1371]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 23:46:59 Moxa device_app_1[1242]:
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:46:59 Moxa device_app_1[1242]: INFO[time] ntp.go:88: NTP update failed
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |     573.625µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |     753.375µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |    3.846625ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |      4.0715ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 23:47:22 Moxa device_app_1[1242]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602064384,"used":999729152,"usedPercent":15.143296235309592,"inodesTotal":1607808,"inodesUsed":35099,"inodesFree":1572709,"inodesUsedPercent":2.183034292651859}
May 20 23:47:22 Moxa device_app_1[1242]:
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |       5.571ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |      6.0585ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |      293.75µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 23:47:22 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:22 | 200 |       584.5µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 23:47:29 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 23:47:29 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:47:00 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:01 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:02 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:03 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:04 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:05 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:06 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:07 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:08 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:09 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:10 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:11 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:12 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:13 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:14 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:15 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:16 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:17 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:18 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:19 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:20 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:21 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:22 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:23 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:24 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:25 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:26 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:27 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]: 20 May 23:47:28 ntpd[1366]: SHM: difference limit exceeded, delta=35987s
May 20 23:47:29 Moxa device_app_1[1242]:
May 20 23:47:29 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:47:29 Moxa device_app_1[1242]: INFO[time] ntp.go:88: NTP update failed
May 20 23:47:32 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T18:47:22.000+08:00]
May 20 18:47:22 Moxa device_app_1[1242]: INFO[time] time.go:39: Update Time: 2020-05-20T18:47:22.000+08:00
May 20 18:47:22 Moxa device_app_1[1242]: INFO[time] time.go:136: Time updated to 2020-05-20T18:47:22.000+08:00
May 20 18:47:22 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 18:47:23 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 18:47:23 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 18:47:23 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 18:47:23 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 18:47:23 Moxa device_app_1[1242]: time="2020-05-20T18:47:23+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T18:47:22.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 18:47:23 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 18:47:23 | 200 |    1.9494215s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 18:47:23 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 18:47:23 | 200 |  1.949896625s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 18:47:27 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 18:47:27 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 18:47:28 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 18:47:28 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 18:47:28 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 18:47:28 Moxa device_app_1[1242]: time="2020-05-20T18:47:28+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 18:47:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 18:47:28 | 200 |  781.411875ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 18:47:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 18:47:28 | 200 |     781.674ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:47:52 Moxa device_app_1[1242]: time="2020-05-20T13:47:52+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:47:52 Moxa device_app_1[1242]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:47:52 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |     555.875µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |       768.5µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |      192.25µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |     386.375µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:47:53 Moxa device_app_1[1242]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602060288,"used":999733248,"usedPercent":15.143358279055397,"inodesTotal":1607808,"inodesUsed":35103,"inodesFree":1572705,"inodesUsedPercent":2.1832830785765465}
May 20 13:47:53 Moxa device_app_1[1242]:
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |     5.49275ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |     5.70975ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |       3.905ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:47:53 | 200 |     4.11625ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:48:08 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T23:47:53.000+08:00]
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] time.go:39: Update Time: 2020-05-20T23:47:53.000+08:00
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] time.go:136: Time updated to 2020-05-20T23:47:53.000+08:00
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 23:47:53 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 23:47:53 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 23:47:53 Moxa device_app_1[1242]: time="2020-05-20T23:47:53+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"2020-05-20T13:47:52+08:00\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T23:47:53.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 23:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:53 | 200 |  602.423125ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:53 | 200 |  602.623625ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:47:55 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 23:47:55 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 23:47:56 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 23:47:56 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:47:56 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 23:47:56 Moxa device_app_1[1242]: time="2020-05-20T23:47:56+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 23:47:56 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:56 | 200 |  1.494386875s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:47:56 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 23:47:56 | 200 |  1.494677125s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 23:48:26 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 23:48:26 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:47:57 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:47:58 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:47:59 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:00 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:01 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:02 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:03 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:04 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:05 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:06 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:07 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:08 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:09 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:10 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:11 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:12 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:13 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:14 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:15 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:16 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:17 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:18 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:19 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:20 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:21 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:22 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:23 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:24 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]: 20 May 23:48:25 ntpd[1383]: SHM: difference limit exceeded, delta=35984s
May 20 23:48:26 Moxa device_app_1[1242]:
May 20 23:48:26 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 23:48:26 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 23:48:56 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [date --set 2020-05-20T20:47:53.000+08:00]
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] time.go:39: Update Time: 2020-05-20T20:47:53.000+08:00
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] time.go:136: Time updated to 2020-05-20T20:47:53.000+08:00
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: signal: killed
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 23:48:27 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:28 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:29 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:30 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:31 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:32 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:33 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:34 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:35 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:36 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:37 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:38 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:39 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:40 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:41 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:42 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:43 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:44 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:45 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:46 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:47 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:48 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:49 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:50 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:51 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:52 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:53 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:54 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]: 20 May 23:48:55 ntpd[1386]: SHM: difference limit exceeded, delta=35984s
May 20 20:47:53 Moxa device_app_1[1242]:
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 20:47:53 Moxa device_app_1[1242]: time="2020-05-20T20:47:53+08:00" level=warning msg="NTP Time Sync Failed" category="device setting" event="time sync failed" origin="Device App" user=
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:140: Non-zero exit code: exit status 1
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander.go:141: stderr: 20 May 20:47:53 ntpd[1391]: unable to bind to wildcard address :: - another process may be running - EXITING
May 20 20:47:53 Moxa device_app_1[1242]:
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] ntp.go:77: NTP update timeout
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] ntp.go:88: NTP update failed
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 20:47:53 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 20:47:53 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":false,"source":"gps","server":"","interval":60}}]
May 20 20:47:53 Moxa device_app_1[1242]: time="2020-05-20T20:47:53+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":false,\"interval\":60,\"server\":\"\"},\"time\":\"2020-05-20T20:47:53.000+08:00\"}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 20:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 20:47:53 | 200 |  962.731125ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 20:47:53 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 20:47:53 | 200 |   963.04025ms |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:49:12 Moxa device_app_1[1242]: time="2020-05-20T13:49:12+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:49:12 Moxa device_app_1[1242]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:49:12 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:49:13 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:53: status code: 200
May 20 13:49:13 Moxa device_app_1[1242]: INFO[time] commander_bruno.go:54: output: {"message":"","return":0}
May 20 13:49:16 Moxa device_app_1[1242]: INFO[time] time.go:105: Setup NTP (GPS)
May 20 13:49:16 Moxa device_app_1[1242]: INFO[time] commander.go:109: exec: [ntpd -c /host/etc/ntp.conf -g -q]
May 20 13:49:16 Moxa device_app_1[1242]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 20 13:49:16 Moxa device_app_1[1242]: time="2020-05-20T13:49:16+08:00" level=info msg="Time request on {\"lastUpdateTime\":\"-\",\"timezone\":\"Asia/Taipei\",\"ntp\":{\"enable\":true,\"interval\":60,\"server\":\"\"}}" category="device setting" event="configuration update success" origin="Device App" user=
May 20 13:49:16 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:16 | 200 |  3.254957875s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:49:16 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:16 | 200 |   3.25515425s |   192.168.4.111 | PATCH    /api/v1/device/time
May 20 13:49:17 Moxa device_app_1[1242]: time="2020-05-20T13:49:17+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 20 13:49:17 Moxa device_app_1[1242]: INFO[time] ntp.go:82: NTP update successfully
May 20 13:49:17 Moxa device_app_1[1242]: INFO[time] commander.go:38: exec: [hwclock -w]
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     196.375µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     381.875µs |   192.168.4.111 | GET      /api/v1/device/gps
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |      165.25µs |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     353.875µs |   192.168.4.111 | GET      /api/v1/device/network/wan
May 20 13:49:28 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s eth0]
May 20 13:49:28 Moxa device_app_1[1242]: INFO[general] device.go:211: Disk Usage: {"path":"/","fstype":"","total":6890492928,"free":5602056192,"used":999737344,"usedPercent":15.143420322801202,"inodesTotal":1607808,"inodesUsed":35103,"inodesFree":1572705,"inodesUsedPercent":2.1832830785765465}
May 20 13:49:28 Moxa device_app_1[1242]:
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     4.06075ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |    4.211375ms |   192.168.4.111 | GET      /api/v1/device/general
May 20 13:49:28 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s eth1]
May 20 13:49:28 Moxa device_app_1[1242]: INFO[ip] commander.go:38: exec: [ip a s wlan0]
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     694.125µs |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |      921.75µs |   192.168.4.111 | GET      /api/v1/device/cellulars
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |    58.35075ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:49:28 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:28 | 200 |     58.5355ms |   192.168.4.111 | GET      /api/v1/device/ethernets
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:28 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:29 | 200 |       3.611ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:49:29 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:29 | 200 |       3.807ms |   192.168.4.111 | GET      /api/v1/device/zoneinfo
May 20 13:49:29 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:29 | 200 |         515µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:49:29 Moxa device_app_1[1242]: [GIN] 2020/05/20 - 13:49:29 | 200 |       695.5µs |   192.168.4.111 | GET      /api/v1/device/time
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/cpuUsage
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/memoryUsage
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/systemDiskUsed
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkUsage
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkTx
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:375: redis ....  key: system/status/networkRx
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/cellular1Signal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/status/gpsSignal)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/network/networkUsage)system/network/networkTx)system/network/networkRx)
May 20 13:49:29 Moxa device_app_1[1242]: [manager] tag.go:377: hooks: system/storage/systemDiskUsed)system/storage/systemDiskFree)system/storage/systemDiskPercent)
packet_write_wait: Connection to 192.168.4.127 port 22: Broken pipe

```

