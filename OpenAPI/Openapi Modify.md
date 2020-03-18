修改 openapi 文件，以 edge-device-mil 為例：



參考 index.yaml

/openapi/index.yaml :

```
paths:
  /device/network/wan:
    $ref: './module.d/route.yaml#/paths/~1device~1network~1wan'
```

Reference: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#pathsObject



/openapi/module.d/route:

```
paths:
	  /device/network/wan:
    get:
      description: Get device route configuration.
      tags: ["device", "network", "wan"]
      responses:
        200:
          description: Device route configuration.
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/WANInformation'
              example:
                data:
                  name: eth0
                  displayName: LAN1
                  ip: 192.168.3.127
                  netmask: 255.255.255.0
                  gateway: 192.168.3.254
                  dns:
                  - 192.168.3.
```



使用 makefile 的 `make doc` 查看修改結果。

