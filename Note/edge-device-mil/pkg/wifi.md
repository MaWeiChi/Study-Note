# wifi

There are two mode of Wi-Fi, the ap mode and the client mode. And the backend application are the `hostapd` and the `wpa_supplicant`.

To achieve that sync the Wi-Fi routers in the disabled ap mode.

The Wi-Fi interface is only using with the `hostapd` in the **enabled ap mode**.
Besides, it is using with the `wpa_supplicant` including **disable ap mode**.

---

## AP

```[go]
const (
	HostapdConfDir      = "/host/etc/hostapd/"
	HostapdOmitPidDir   = "/run/sendsigs.omit.d/"
	HostapdIfUpDownSh   = "/etc/hostapd/ifupdown.sh"
	HostapdConfTemplate = `interface=%s
driver=nl80211
ssid=%s
hw_mode=%s
channel=%d
country_code=%s
ignore_broadcast_ssid=%d
max_num_sta=2
macaddr_acl=0
`
)
```

[hostapd docs](https://w1.fi/cgit/hostap/plain/hostapd/hostapd.conf)



config(20Mhz channel):

```[bash]
interface=wlan0
driver=nl80211
ssid=moxa-20_40_80Mhz
hw_mode=a
channel=40
country_code=US
ignore_broadcast_ssid=0
max_num_sta=2
macaddr_acl=0
ieee80211n=1
ieee80211ac=1
auth_algs=1
wpa=2
wpa_passphrase=admin@123
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP
```

[40-80Mhz](https://www.notion.so/isd/Wi-Fi-AP-hostapd-20-40-80Mhz-channels-72e03c33664547d98ff37c641196e241)

---
## Client

We use the package [`wpac-go`](https://gitlab.com/moxa/ibg/software/platform/thingspro/wpac-go/-/tree/master) to communicate with the `wpa_supplicant` via D-BUS API.