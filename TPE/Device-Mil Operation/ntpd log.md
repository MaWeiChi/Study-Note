```bash
-- Logs begin at Fri 2020-05-15 12:25:30 CST. --
May 15 12:26:42 Moxa device_app_1[1234]: INFO[route] commander.go:38: exec: [ip route show 0.0.0.0/0]
May 15 12:26:42 Moxa device_app_1[1234]: INFO[route] route.go:58: No need to change
May 15 12:27:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:27:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:27:32 Moxa device_app_1[1234]: INFO[time] ntp.go:92: NTP update successfully
May 15 12:27:32 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [hwclock -w]
May 15 12:28:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:28:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:28:33 Moxa device_app_1[1234]: INFO[time] ntp.go:92: NTP update successfully
May 15 12:28:33 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [hwclock -w]
May 15 12:29:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:29:25 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:29:32 Moxa device_app_1[1234]: INFO[time] ntp.go:92: NTP update successfully
May 15 12:29:32 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [hwclock -w]
May 15 12:29:40 Moxa device_app_1[1234]: INFO[time] time.go:96: Setup NTP (GPS)
May 15 12:29:40 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:29:40 Moxa device_app_1[1234]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"gps","server":"","interval":60}}]
May 15 12:29:40 Moxa device_app_1[1234]: time="2020-05-15T12:29:40+08:00" level=info msg="Time request on {\"ntp\": {\"enable\": true, \"source\": \"gps\", \"server\": \"tock.stdtime.gov.tw\", \"interval\": 60}}" category="device setting" event="configuration update success" origin="Device App" user=
May 15 12:29:40 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:29:40 | 200 |  4.248382045s |  192.168.50.248 | PATCH    /api/v1/device/time
May 15 12:29:40 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:29:40 | 200 |  4.249936794s |  192.168.50.248 | PATCH    /api/v1/device/time
May 15 12:29:40 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:30:40 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:30:40 Moxa device_app_1[1234]: INFO[time] ntp.go:70: another NTP porcess is running. msg:job 628
May 15 12:30:40 Moxa device_app_1[1234]: 
May 15 12:30:40 Moxa device_app_1[1234]: 
May 15 12:30:40 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c kill $(netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}')]
May 15 12:30:40 Moxa device_app_1[1234]: INFO[time] ntp.go:72: another NTP porcess is killed
May 15 12:30:40 Moxa device_app_1[1234]: INFO[time] ntp.go:95: NTP update failed. msg:<nil>
May 15 12:30:40 Moxa device_app_1[1234]: 
May 15 12:30:40 Moxa device_app_1[1234]: time="2020-05-15T12:30:40+08:00" level=warning msg="NTP Time Sync Failed" category="device setting" event="time sync failed" origin="Device App" user=
May 15 12:30:43 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] time.go:98: Setup NTP (Time Server:tock.stdtime.gov.tw)
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:31:30 Moxa device_app_1[1234]: [manager] manager.go:281: [Update Device DB] [{"timezone":"Asia/Taipei","ntp":{"enable":true,"source":"timeserver","server":"tock.stdtime.gov.tw","interval":60}}]
May 15 12:31:30 Moxa device_app_1[1234]: time="2020-05-15T12:31:30+08:00" level=info msg="Time request on {\"ntp\": {\"enable\": true, \"source\": \"timeserver\", \"server\": \"tock.stdtime.gov.tw\", \"interval\": 60}}" category="device setting" event="configuration update success" origin="Device App" user=
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] ntp.go:70: another NTP porcess is running. msg:job 640
May 15 12:31:30 Moxa device_app_1[1234]: 
May 15 12:31:30 Moxa device_app_1[1234]: 
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c kill $(netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}')]
May 15 12:31:30 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:31:30 | 200 |  2.252229375s |  192.168.50.248 | PATCH    /api/v1/device/time
May 15 12:31:30 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:31:30 | 200 |   2.25248375s |  192.168.50.248 | PATCH    /api/v1/device/time
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] ntp.go:72: another NTP porcess is killed
May 15 12:31:30 Moxa device_app_1[1234]: INFO[time] ntp.go:95: NTP update failed. msg:<nil>
May 15 12:31:30 Moxa device_app_1[1234]: 
May 15 12:31:30 Moxa device_app_1[1234]: time="2020-05-15T12:31:30+08:00" level=warning msg="NTP Time Sync Failed" category="device setting" event="time sync failed" origin="Device App" user=
May 15 12:31:33 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:31:41 Moxa device_app_1[1234]: INFO[time] ntp.go:92: NTP update successfully
May 15 12:31:41 Moxa device_app_1[1234]: time="2020-05-15T12:31:41+08:00" level=info msg="NTP Time Sync Successed" category="device setting" event="time sync success" origin="Device App" user=
May 15 12:31:41 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [hwclock -w]
May 15 12:31:57 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:31:57 | 200 |     551.375µs |  192.168.50.248 | GET      /api/v1/device/time
May 15 12:31:57 Moxa device_app_1[1234]: [GIN] 2020/05/15 - 12:31:57 | 200 |      723.75µs |  192.168.50.248 | GET      /api/v1/device/time
May 15 12:32:30 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [sh -c netstat -lanp | grep ':123' | awk 'FNR==1 split($6,a,"/") {print a[1]}']
May 15 12:32:30 Moxa device_app_1[1234]: INFO[time] commander.go:38: exec: [ntpd -c /host/etc/ntp.conf -g -q 30]
May 15 12:32:37 Moxa device_app_1[1234]: INFO[time] ntp.go:92: NTP update successfully

```

