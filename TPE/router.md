```bash
: [GIN-debug] GET    /api/v1/device/wifi       --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifis-fm (5 handlers)
Jul 27 13:55:37 Moxa device_app_1[1242]: [GIN-debug] PATCH  /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).PatchWifiById-fm (5 handlers)
Jul 27 13:55:37 Moxa device_app_1[1242]: [GIN-debug] GET    /api/v1/device/wifi/:id   --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiById-fm (5 handlers)
Jul 27 13:55:37 Moxa device_app_1[1242]: [GIN-debug] GET    /api/v1/device/wifi/:id/client/scan --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiClientScanById-fm (5 handlers)
Jul 27 13:55:37 Moxa device_app_1[1242]: [GIN-debug] GET    /api/v1/device/wifi/in    --> github.com/MOXA-ISD/edge-device-mil/pkg/wifi.(*Service).GetWifiByInternal-fm (5 handlers)
Jul 27 13:55:37 Moxa device_app_1[1242]: [manager] manager.go:261: 'in' in new path '/api/v1/device/wifi/in' conflicts with existing wildcard ':id' in existing prefix '/api/v1/device/wifi/:id'
Jul 27 13:55:37 Moxa device_app_1[1242]: [manager] manager.go:176: wifi panic, count:2, schedule to restart service after 10 seconds
Jul 27 13:55:44 Moxa device_app_1[1242]: [manager] service.go:38: wifi Stopping
Jul 27 13:55:44 Moxa device_app_1[1242]: [manager] service.go:40: wifi Stopped

```

```

```

