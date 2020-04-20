```
export APPMAN_HOST=10.144.48.73 APPMAN_SSH_USER=moxa APPMAN_SSH_PASS=moxa APPMAN_HTTP_USER=admin APPMAN_HTTP_PASS=admin@123
http --session=/tmp/session-${APPMAN_HOST}.json POST "$APPMAN_HOST/api/v1/auth" name=admin password=admin@123
http --session=/tmp/session1.json POST "${APPMAN_HOST}/api/v1/events" category=network event="uplink change" origin="device" message="aaaaaaabbbbbbbb" severity=warning
http --session=/tmp/session-${APPMAN_HOST}.json POST "${APPMAN_HOST}/api/v1/events" category=network event="uplink change" origin="device" message="aaaaaaabbbbbbbb" severity=warning
```

```
root@Moxa:/home/moxa# appman log show | grep uplink | tail
{"category":"network","createdAt":"2020-03-24T14:27:40.35028164+08:00","event":"uplink change","id":661,"message":"No default route","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T14:27:45.19228796+08:00","event":"uplink change","id":663,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T15:14:33.81771768+08:00","event":"uplink change","id":665,"message":"No default route","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T15:14:43.07229764+08:00","event":"uplink change","id":666,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T16:56:26.9044944+08:00","event":"uplink change","id":677,"message":"No default route","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T16:56:31.58360308+08:00","event":"uplink change","id":678,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T17:50:28.38026564+08:00","event":"uplink change","id":681,"message":"No default route","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T17:50:31.24072432+08:00","event":"uplink change","id":682,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-24T18:59:38.65840424+08:00","event":"uplink change","id":684,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
{"category":"network","createdAt":"2020-03-25T12:21:44.66154168+08:00","event":"uplink change","id":696,"message":"Change to LAN1 (10.144.48.73)","origin":"Device App","severity":"warning","user":""}
```

```
root@Moxa:/home/moxa# appman log profile ls | grep uplink
|  3 | network        | uplink change                   | warning  | false  |
```