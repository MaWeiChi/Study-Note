# wpa_supplicant

Using the wpa_supplicat and the wap_cli to control Wi-Fi in our devices:

get interface name:

```c
$ iw dev
phy#0
        Unnamed/non-netdev interface
                wdev 0x2
                addr 00:15:61:20:9d:4b
                type P2P-device
                txpower 0.00 dBm
        Interface wlan0
                ifindex 5
                wdev 0x1
                addr 00:15:61:20:9d:4b
                type managed
                txpower 0.00 dBm
```

## config setting

append the two lines in the ```/etc/network/interfaces``` ([reference](https://manpages.debian.org/stretch/ifupdown2/interfaces.5.en.html)) :

```bash
auto wlan0
iface wlan0 inet dhcp
```

create config `/etc/wpa_supplicant/wpa_supplicant.conf` ([reference](https://w1.fi/cgit/hostap/plain/wpa_supplicant/wpa_supplicant.conf)) :

```bash
ctrl_interface=/var/run/wpa_supplicant
update_config=1
```

luanch

```bash
wpa_supplicant -D nl80211 -i wlan0 -c /etc/wpa_supplicant.conf -B
```
