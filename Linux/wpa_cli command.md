wpa_cli -i wlan0 status (看看目前的狀態)

wpa_cli -i wlan0 help (查看參數與指令)

wpa_cli -i wlan0 terminate (關閉wpa_supplicant)

wpa_cli -i wlan0 interface (查看有那些無線網卡介面)

wpa_cli -i wlan0 list_networks (查看wpa_supplicant.conf檔裡的設定)

wpa_cli -i wlan0 select_network (選擇不同的AP，id為AP代號0,1,2,3...)

wpa_cli -i wlan0 enable_network

wpa_cli -i wlan0 disable_network

wpa_cli -i wlan0 remove_network (刪除AP資訊)

wpa_cli -i wlan0 reconfigure (重新讀取wpa_supplicant.conf設定內容)

wpa_cli -i wlan0 save_config (儲存寫入到wpa_supplicant.conf中，否責變更無效)

wpa_cli -i wlan0 disconnect (斷線)

wpa_cli -i wlan0 reconnect (重新連線)

wpa_cli -i wlan0 scan (掃描附近的AP)

wpa_cli -i wlan0 scan_results (印出附近AP的相關資訊)