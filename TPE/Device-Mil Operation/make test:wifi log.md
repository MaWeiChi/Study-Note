```
root@b476d9c49e72:/data# make test/wifi 
DEVICE_DEBUG=1 TP_DEVICE_DEBUG=1 go test -mod=vendor -ldflags "-s -w" -tags "bruno" -v github.com/MOXA-ISD/edge-device-mil/pkg/wifi
=== RUN   TestWifiInitialize
--- PASS: TestWifiInitialize (0.00s)
=== RUN   Test_wifi_s1
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/devicedb          --> github.com/MOXA-ISD/edge-device-mil/internal/test.(*MockDevDB).dbTestFunc-fm (3 handlers)
2020/05/20 08:54:23 listening :59000
2020/05/20 08:54:23 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:23 | 200 |        44.4µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:25 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:25 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[GIN] 2020/05/20 - 08:54:25 | 200 |       223.8µs |       127.0.0.1 | GET      /api/v1/device/wifi
[GIN] 2020/05/20 - 08:54:25 | 200 |       307.2µs |       127.0.0.1 | GET      /api/v1/device/wifi
2020/05/20 08:54:25 [checker] compare start =>
Expected: 200
Received: 200
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:25  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : get wifi info                                                                                                          |
|Result : Success                                                                                                                |
|Elapsed: 2 ms                                                                                                                   |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |2              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s1 (2.04s)
=== RUN   Test_wifi_s2
2020/05/20 08:54:25 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:25 | 200 |        66.5µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:27 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:27 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[GIN] 2020/05/20 - 08:54:27 | 200 |         382µs |       127.0.0.1 | GET      /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:27 | 200 |       448.4µs |       127.0.0.1 | GET      /api/v1/device/wifi/1
2020/05/20 08:54:27 [checker] compare start =>
Expected: 200
Received: 200
2020/05/20 08:54:27 [checker] compare start =>
Key: data.enable
Expected: false
Received: false
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:27  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : get wifi info by id                                                                                                    |
|Result : Success                                                                                                                |
|Elapsed: 3 ms                                                                                                                   |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |3              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|3  |                              |data.enable                   |false                         |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s2 (2.04s)
=== RUN   Test_wifi_s3
2020/05/20 08:54:27 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:27 | 200 |        80.7µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:29 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:29 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":true,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":"@!@!@!@!"},"clients":null}}]
[GIN] 2020/05/20 - 08:54:29 | 200 |        62.8µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN] 2020/05/20 - 08:54:29 | 200 |     19.1661ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:29 | 200 |      19.235ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:29 [checker] compare start =>
Expected: 200
Received: 200
2020/05/20 08:54:29 [checker] compare start =>
Key: data.ap.security.password
Expected: @!@!@!@!
Received: @!@!@!@!
[manager] service.go:67: wifi stoped
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:29  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP password                                                                                               |
|Result : Success                                                                                                                |
|Elapsed: 22 ms                                                                                                                  |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |22             |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|3  |                              |data.ap.security.password     |@!@!@!@!                      |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s3 (2.05s)
=== RUN   Test_wifi_s4
2020/05/20 08:54:29 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:29 | 200 |        43.8µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:31 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:31 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
INFO[wifi] api.go:27: Key: 'WifiEntry.WifiReadWriteEntry.Ap.Security.Password' Error:Field validation for 'Password' failed on the 'min' tag
[GIN] 2020/05/20 - 08:54:31 | 400 |       347.9µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:31 | 400 |       466.7µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:31 [checker] compare start =>
Expected: 400
Received: 400
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:31  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP password by invalid length                                                                             |
|Result : Success                                                                                                                |
|Elapsed: 3 ms                                                                                                                   |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |3              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |400                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s4 (2.03s)
=== RUN   Test_wifi_s5
2020/05/20 08:54:31 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:31 | 200 |        54.9µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:33 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:33 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":true,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-test-ap_1","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:33 | 200 |        53.1µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN] 2020/05/20 - 08:54:33 | 200 |      13.962ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:33 | 200 |     14.1326ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:33 [checker] compare start =>
Expected: 200
Received: 200
2020/05/20 08:54:33 [checker] compare start =>
Key: data.ap.ssid
Expected: moxa-test-ap_1
Received: moxa-test-ap_1
[manager] service.go:38: wifi Stopping
[manager] service.go:67: wifi stoped
[manager] service.go:40: wifi Stopped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:33  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP SSID                                                                                                   |
|Result : Success                                                                                                                |
|Elapsed: 16 ms                                                                                                                  |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |16             |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|3  |                              |data.ap.ssid                  |moxa-test-ap_1                |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s5 (2.04s)
=== RUN   Test_wifi_s6
2020/05/20 08:54:33 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:33 | 200 |        69.1µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:35 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:35 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":true,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":5,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:35 | 200 |        41.6µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN] 2020/05/20 - 08:54:35 | 200 |     15.2116ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:35 | 200 |     15.3221ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:35 [checker] compare start =>
Expected: 200
Received: 200
2020/05/20 08:54:35 [checker] compare start =>
Key: data.ap.channel
Expected: 5
Received: 5
[manager] service.go:38: wifi Stopping
[manager] service.go:67: wifi stoped
[manager] service.go:40: wifi Stopped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:35  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP SSID by invalid schema                                                                                 |
|Result : Success                                                                                                                |
|Elapsed: 17 ms                                                                                                                  |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |17             |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|3  |                              |data.ap.channel               |5                             |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s6 (2.05s)
=== RUN   Test_wifi_s7
2020/05/20 08:54:35 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:35 | 200 |          24µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:37 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:37 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
INFO[wifi] api.go:93: The configured channel is not supported in the region
[GIN] 2020/05/20 - 08:54:37 | 400 |       298.9µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:37 | 400 |       352.8µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:37 [checker] compare start =>
Expected: 400
Received: 400
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:37  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP as invalid channel                                                                                     |
|Result : Success                                                                                                                |
|Elapsed: 2 ms                                                                                                                   |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |2              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |400                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s7 (2.03s)
=== RUN   Test_wifi_s8
2020/05/20 08:54:37 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:37 | 200 |        51.3µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:39 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:39 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":true,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":5,"region":"US","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:39 | 200 |        23.4µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN] 2020/05/20 - 08:54:39 | 200 |     13.0817ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:39 | 200 |     13.1214ms |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:39 [checker] compare start =>
Expected: 200
Received: 200
2020/05/20 08:54:39 [checker] compare start =>
Key: data.ap.region
Expected: US
Received: US
2020/05/20 08:54:39 [checker] compare start =>
Key: data.ap.channel
Expected: 5
Received: 5
2020/05/20 08:54:39 [checker] [{"type":"wifi","name":"wlan0","available":true,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":5,"region":"US","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:39  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP region and channel                                                                                     |
|Result : Success                                                                                                                |
|Elapsed: 15 ms                                                                                                                  |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |15             |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |200                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|3  |                              |data.ap.region                |US                            |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|4  |                              |data.ap.channel               |5                             |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s8 (2.05s)
=== RUN   Test_wifi_s9
2020/05/20 08:54:39 MOXA api token open failed, open /run/mx-api-token: no such file or directory
[manager] tag.go:87: System tag client init failed.
[manager] tag.go:121: Create pubsub receive exit because timeout.
[wifi] config.go:171: main config load error: open /data/data/module.d/wifi/configuration.json: no such file or directory,<nil>, try load backup
[wifi] config.go:181: backup config load error: open /data/data/module.d/wifi/configuration.json.backup: no such file or directory,<nil>, try load factory
INFO[wifi] index.go:100: Load config code: factory config

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[manager] tag.go:360: pubsub channel is null
[manager] service.go:58: wifi main process running
[manager] manager.go:281: [Update Device DB] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
[GIN] 2020/05/20 - 08:54:39 | 200 |        36.2µs |       127.0.0.1 | POST     /api/v1/devicedb
[GIN-debug] PATCH  /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifis-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
[GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
2020/05/20 08:54:41 [checker] [{"type":"wifi","name":"wlan0","available":false,"status":false,"capabilities":{"mode":["ap"],"band":[]},"regions":[{"name":"Australia","countryCode":"AU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Bahrain","countryCode":"BH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Brazil","countryCode":"BR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Switzerland","countryCode":"CH","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"China","countryCode":"CN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Europe","countryCode":"GB","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Israel","countryCode":"IL","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48]},{"name":"India","countryCode":"IN","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Japan","countryCode":"JP","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13,14],"band50":[8,12,16,36,38,40,42,44,46,48]},{"name":"Korea","countryCode":"KR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"New Zealand","countryCode":"NZ","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Russia","countryCode":"RU","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Singapore","countryCode":"SG","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"Turkey","countryCode":"TR","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48]},{"name":"Taiwan","countryCode":"TW","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[34,36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"North America","countryCode":"US","band24":[1,2,3,4,5,6,7,8,9,10,11],"band50":[36,38,40,42,44,46,48,149,153,157,161,165]},{"name":"South Africa","countryCode":"ZA","band24":[1,2,3,4,5,6,7,8,9,10,11,12,13],"band50":[36,38,40,42,44,46,48,100,104,108,112,116,120,124,128,132,136,140]}],"id":1,"enable":false,"mode":"ap","ap":{"band":"band24","broadcastSsid":true,"ssid":"moxa-sample-ap","channel":6,"region":"TW","security":{"mode":"wpa2","encryption":"aes","password":""},"clients":null}}]
2020/05/20 08:54:41 [checker] compare start =>
Key: 0.enable
Expected: false
Received: false
INFO[wifi] api.go:27: unexpected end of JSON input
[GIN] 2020/05/20 - 08:54:41 | 400 |       404.8µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
[GIN] 2020/05/20 - 08:54:41 | 400 |       488.1µs |       127.0.0.1 | PATCH    /api/v1/device/wifi/1
2020/05/20 08:54:41 [checker] compare start =>
Expected: 400
Received: 400
[manager] service.go:38: wifi Stopping
[manager] service.go:40: wifi Stopped
[manager] service.go:67: wifi stoped
[manager] manager.go:101: Server exit success
2020/05/20 08:54:41  --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|Case   : modify Wi-Fi AP by invalid region                                                                                      |
|Result : Success                                                                                                                |
|Elapsed: 3 ms                                                                                                                   |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|   |Description                   |Key                           |Expected                      |Received       |Elapsed        |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|1  |                              |0.enable                      |false                         |ok             |3              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------
|2  |                              |$CODE                         |400                           |ok             |0              |
 --- ------------------------------ ------------------------------ ------------------------------ --------------- ---------------

--- PASS: Test_wifi_s9 (2.03s)
PASS
ok      github.com/MOXA-ISD/edge-device-mil/pkg/wifi    18.381s
```

