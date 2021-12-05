# internal

## linux

---
## service
Support http/https, openvpn, rds, ssh service interface.

| port    | HTTP | HTTPS | SSH | RDS  |
| ------- | ---- | ----- | --- | ---- |
| Linux   | x    | x     | 22  | x    |
| TPG     | 80   | 443   | 22  | x    |
| TPE     | 80   | 8443  | 22  | x    |
| Windows | x    | x     | 22  | 3389 |

### http/https
HTTP and HTTPS are supported on the TPG and the TPE.




---

## vendor
There are some modified some vendor code, it wil be alter with the `build` command in the Makefile.

