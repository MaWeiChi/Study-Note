example:
wpa_supplicant -Dnl80211 -iwlan0 -c/etc/wpa_supplicant.conf
wpa_supplicant -B -i wlan0 -c /etc/wpa_supplicant/wpa_supplicant.conf

```
nl80211: Could not set interface 'p2p-dev-wlan0' UP
nl80211: deinit ifname=p2p-dev-wlan0 disabled_11b_rates=0
p2p-dev-wlan0: Failed to initialize driver interface
P2P: Failed to enable P2P Device interface
```



```
killall wpa_supplicant
```

```
apt-get install psmisc(killall)
```



```sh
echo $(iw dev | awk '$1=="Interface"{print $2}')
```

