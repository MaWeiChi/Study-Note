```bash
528  nginx -T -c /var/thingspro/data/nginx/nginx.conf 

529  vi /var/thingspro/data/nginx/locations/50-wificlient.conf 

530  nginx -T -c /var/thingspro/data/nginx/nginx.conf 

531  fg

532  vi /var/thingspro/data/nginx/locations/50-wificlient.conf 

533  nginx -s reload -c /var/thingspro/data/nginx/nginx.conf 

534  nginx -T -c /var/thingspro/data/nginx/nginx.conf 

535  curl -H "mx-api-token:`cat /var/thingspro/data/mx-api-token`" http://127.0.0.1/api/v1/wificlient/status

536  docker ps -a

537  journalctl APPNAME=wificlient -f

538  curl -H "mx-api-token:`cat /var/thingspro/data/mx-api-token`" http://127.0.0.1/api/v1/wificlient/status
539  docker ps -a

540  docker inspect 8d

541  curl -H "mx-api-token:`cat /var/thingspro/data/mx-api-token`" http://172.31.9.18:8818/api/v1/wificlient/status

542  curl -H "mx-api-token:`cat /var/thingspro/data/mx-api-token`" http://172.31.9.18:8818/api/v1/wificlient/reconnect

543  bg

544  ps aux| grep wpa_cli

545  fg

546  fg

547  history
```


