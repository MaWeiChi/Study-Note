```
-- Logs begin at Mon 2021-05-17 07:46:58 UTC. --
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] GET    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDashboard-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] PUT    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putDashboard-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 09:48:04 Moxa dlm-agent[28271]: [INFO]May 19 09:48:04.087 - [origin:remoted] - service dlm started
May 19 09:48:04 Moxa dlm-agent[28271]: [INFO]May 19 09:48:04.087 - [origin:remoted] - all started
May 19 09:53:00 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:00 | 200 |      1.1873ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:00 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:00 | 200 |     223.517µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:00 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:00 | 200 |     231.276µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:00 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:00 | 200 |     219.236µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:01 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:01 | 200 |     220.876µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:01 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:01 | 200 |     221.716µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:03 Moxa dlm-agent[28271]: [monitor] run check
May 19 09:53:12 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:12 | 200 |     244.556µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:53:16 Moxa dlm-agent[28271]: [INFO]May 19 09:53:16.688 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:53:17 Moxa dlm-agent[28271]: [INFO]May 19 09:53:17.340 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:53:17 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:17 | 200 |  1.434255594s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"
May 19 09:53:18 Moxa dlm-agent[28271]: [INFO]May 19 09:53:18.676 - [origin:dlm] - skip stop because connection status: false
May 19 09:53:19 Moxa dlm-agent[28271]: [INFO]May 19 09:53:19.488 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:53:20 Moxa dlm-agent[28271]: [INFO]May 19 09:53:20.187 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:53:20 Moxa dlm-agent[28271]: [INFO]May 19 09:53:20.190 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 09:53:20 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:20 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:20 | 200 |   17.076241ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:53:20 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:20 | 200 |   43.276854ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 09:53:20 Moxa dlm-agent[28271]: [INFO]May 19 09:53:20.266 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 09:53:20 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:20 | 200 |     216.317µs |       127.0.0.1 | GET      "/api/v1/_/info"




May 19 09:53:22 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:22.885 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 09:53:22 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:22.888 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 09:53:22 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.929 - [origin:dlm] - latest configuration:  {
May 19 09:53:22 Moxa dlm-agent[28271]:   "connection": {
May 19 09:53:22 Moxa dlm-agent[28271]:     "enable": false,
May 19 09:53:22 Moxa dlm-agent[28271]:     "retryDelaySec": 10,
May 19 09:53:22 Moxa dlm-agent[28271]:     "picTarget": "stage"
May 19 09:53:22 Moxa dlm-agent[28271]:   },
May 19 09:53:22 Moxa dlm-agent[28271]:   "certificates": {
May 19 09:53:22 Moxa dlm-agent[28271]:     "caCertFile": "dev.crt",
May 19 09:53:22 Moxa dlm-agent[28271]:     "caPkFile": "dev.key",
May 19 09:53:22 Moxa dlm-agent[28271]:     "certificateInfo": {
May 19 09:53:22 Moxa dlm-agent[28271]:       "organization": "Moxa Inc.",
May 19 09:53:22 Moxa dlm-agent[28271]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:53:22 Moxa dlm-agent[28271]:       "notBefore": "May 19, 2021 09:53:20",
May 19 09:53:22 Moxa dlm-agent[28271]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:53:22 Moxa dlm-agent[28271]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:53:22 Moxa dlm-agent[28271]:       "macAddress": "0090E882F2FB",
May 19 09:53:22 Moxa dlm-agent[28271]:       "serialNumber": "TAIJB1072828"
May 19 09:53:22 Moxa dlm-agent[28271]:     }
May 19 09:53:22 Moxa dlm-agent[28271]:   },
May 19 09:53:22 Moxa dlm-agent[28271]:   "connectionStatus": {
May 19 09:53:22 Moxa dlm-agent[28271]:     "projectName": "Kevin Test - AGT",
May 19 09:53:22 Moxa dlm-agent[28271]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:53:22 Moxa dlm-agent[28271]:     "status": "",
May 19 09:53:22 Moxa dlm-agent[28271]:     "message": "",
May 19 09:53:22 Moxa dlm-agent[28271]:     "lastConnectedTime": "",
May 19 09:53:22 Moxa dlm-agent[28271]:     "lastConnectTime": "",
May 19 09:53:22 Moxa dlm-agent[28271]:     "lastDisconnectTime": "",
May 19 09:53:22 Moxa dlm-agent[28271]:     "statuslastUpdateTime": ""
May 19 09:53:22 Moxa dlm-agent[28271]:   }
May 19 09:53:22 Moxa dlm-agent[28271]: }
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.943 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.943 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.944 - [origin:dlm] - latest configuration:  {
May 19 09:53:23 Moxa dlm-agent[28271]:   "connection": {
May 19 09:53:23 Moxa dlm-agent[28271]:     "enable": true,
May 19 09:53:23 Moxa dlm-agent[28271]:     "retryDelaySec": 10,
May 19 09:53:23 Moxa dlm-agent[28271]:     "picTarget": "stage"
May 19 09:53:23 Moxa dlm-agent[28271]:   },
May 19 09:53:23 Moxa dlm-agent[28271]:   "certificates": {
May 19 09:53:23 Moxa dlm-agent[28271]:     "caCertFile": "dev.crt",
May 19 09:53:23 Moxa dlm-agent[28271]:     "caPkFile": "dev.key",
May 19 09:53:23 Moxa dlm-agent[28271]:     "certificateInfo": {
May 19 09:53:23 Moxa dlm-agent[28271]:       "organization": "Moxa Inc.",
May 19 09:53:23 Moxa dlm-agent[28271]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:53:23 Moxa dlm-agent[28271]:       "notBefore": "May 19, 2021 09:53:20",
May 19 09:53:23 Moxa dlm-agent[28271]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:53:23 Moxa dlm-agent[28271]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:53:23 Moxa dlm-agent[28271]:       "macAddress": "0090E882F2FB",
May 19 09:53:23 Moxa dlm-agent[28271]:       "serialNumber": "TAIJB1072828"
May 19 09:53:23 Moxa dlm-agent[28271]:     }
May 19 09:53:23 Moxa dlm-agent[28271]:   },
May 19 09:53:23 Moxa dlm-agent[28271]:   "connectionStatus": {
May 19 09:53:23 Moxa dlm-agent[28271]:     "projectName": "Kevin Test - AGT",
May 19 09:53:23 Moxa dlm-agent[28271]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:53:23 Moxa dlm-agent[28271]:     "status": "",
May 19 09:53:23 Moxa dlm-agent[28271]:     "message": "",
May 19 09:53:23 Moxa dlm-agent[28271]:     "lastConnectedTime": "",
May 19 09:53:23 Moxa dlm-agent[28271]:     "lastConnectTime": "",
May 19 09:53:23 Moxa dlm-agent[28271]:     "lastDisconnectTime": "",
May 19 09:53:23 Moxa dlm-agent[28271]:     "statuslastUpdateTime": ""
May 19 09:53:23 Moxa dlm-agent[28271]:   }
May 19 09:53:23 Moxa dlm-agent[28271]: }
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.946 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:22 | 200 |     789.867µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.967 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.967 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.967 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [ERRO]May 19 09:53:22.969 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 09:53:23 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:22 | 404 |       10.16µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 09:53:23 Moxa dlm-agent[28271]: [ERRO]May 19 09:53:22.983 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.983 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.983 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [ERRO]May 19 09:53:22.984 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.984 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.984 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.984 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:22 | 200 |  4.308086577s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"

May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.992 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.992 - [origin:dlm] - register target: stage
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.992 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:22.992 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success
May 19 09:53:23 Moxa dlm-agent[28271]: [INFO]May 19 09:53:23.953 - [origin:dlm] - tasks[3/6] name:check certificate, result:success


May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.962 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.965 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.967 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.974 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 09:53:24 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:24.978 - [origin:dlm-1] - try to connect rule server ...
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.987 - [origin:dlm-1] - connect to rule server success
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.990 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.992 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 09:53:24 Moxa dlm-agent[28271]: [INFO]May 19 09:53:24.992 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.002 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.005 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.013 - [origin:dlm-1] - monitor stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [WARN]May 19 09:53:25.016 - [origin:dlm-1] - waiting for self registration...
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.018 - [origin:dlm-1] - twin stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [WARN]May 19 09:53:25.019 - [origin:dlm-1] - waiting for self registration...
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.019 - [origin:dlm-1] - message stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.128 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.128 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.131 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.132 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.141 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.141 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.142 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.142 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.143 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.147 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.144 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.144 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.152 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.155 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.156 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.156 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.156 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.157 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.165 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.166 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.167 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.167 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.167 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.167 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.168 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.168 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.168 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.169 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.169 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.169 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:53:25 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:25.169 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.169 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.853 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.856 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.856 - [origin:dlm-1] - [connection callback], connection = true
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.866 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 09:53:25 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:25 | 200 |     888.305µs |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 09:53:25 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:25 | 200 |    2.669477ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:53:25 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:25 | 200 |    4.106972ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 09:53:25 Moxa dlm-agent[28271]: [INFO]May 19 09:53:25.903 - [origin:message-group:dlm-1:2] - stopping
May 19 09:53:26 Moxa dlm-agent[28271]: [INFO]May 19 09:53:26.018 - [origin:dlm-1] - monitor stream registration success
May 19 09:53:26 Moxa dlm-agent[28271]: [INFO]May 19 09:53:26.019 - [origin:dlm-1] - twin stream registration success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.143 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.144 - [origin:message-group:dlm-1:2] - message-routine:exit, reason: source canceled
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.144 - [origin:message-group:dlm-1:2] - tasks[1/5] name:stop ipc connection, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.144 - [origin:message-group:dlm-1:2] - outputservice exit
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.145 - [origin:message-group:dlm-1:2] - outputservice state: stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.145 - [origin:message-group:dlm-1:2] - timer stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.145 - [origin:message-group:dlm-1:2] - formater stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.145 - [origin:message-group:dlm-1:2] - tasks[2/5] name:stop output control service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.145 - [origin:message-group:dlm-1:2] - tasks[3/5] name:stop tag service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.153 - [origin:message-group:dlm-1:2] - event origin:OS patch listener exit
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.155 - [origin:message-group:dlm-1:2] - tasks[4/5] name:stop event service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.156 - [origin:message-group:dlm-1:2] - tasks[5/5] name:stop remote monitor service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.156 - [origin:message-group:dlm-1:2] - stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |  2.992983518s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process"
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.155 - [origin:message-group:dlm-1:2] - message:exit, reason: source canceled
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.161 - [origin:message-group:dlm-1:2] - message stream out
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |  2.266454295s |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.169 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.170 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.170 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.322 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.324 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.325 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.325 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.325 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.326 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.327 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.328 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.329 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.329 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.347 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.347 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.348 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.348 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   11.804967ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.432 - [origin:message-group:dlm-1:4] - stopping
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.444 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.454 - [origin:message-group:dlm-1:4] - message-routine:exit, reason: source canceled
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.462 - [origin:message-group:dlm-1:4] - tasks[1/5] name:stop ipc connection, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.463 - [origin:message-group:dlm-1:4] - outputservice exit
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.465 - [origin:message-group:dlm-1:4] - outputservice state: stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.466 - [origin:message-group:dlm-1:4] - timer stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.466 - [origin:message-group:dlm-1:4] - formater stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.467 - [origin:message-group:dlm-1:4] - tasks[2/5] name:stop output control service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.472 - [origin:message-group:dlm-1:4] - tasks[3/5] name:stop tag service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.456 - [origin:message-group:dlm-1:4] - message:exit, reason: source canceled
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.474 - [origin:message-group:dlm-1:4] - message stream out
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |  3.289959912s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.477 - [origin:message-group:dlm-1:4] - event origin:OS patch listener exit
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.478 - [origin:message-group:dlm-1:4] - tasks[4/5] name:stop event service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.482 - [origin:message-group:dlm-1:4] - tasks[5/5] name:stop remote monitor service, result:success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.482 - [origin:message-group:dlm-1:4] - stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   68.330325ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.495 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.497 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.497 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.515 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.515 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.516 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.516 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.516 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.516 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.518 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.518 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.544 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.544 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.573 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.573 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.573 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.573 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:53:28 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   54.893624ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   68.143048ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.614 - [origin:store] - [connectionID:dlm-1] - start
May 19 09:53:28 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   83.891311ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |  103.348753ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.717 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 09:53:28 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |  111.825095ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:53:28 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:28 | 200 |   68.319365ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.788 - [origin:dlm] - last ota result empty
May 19 09:53:28 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:28.791 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.797 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.797 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: waiting
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.797 - [connectionID:dlm-1] - [origin:store] - get connection success, start register
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.801 - [origin:dlm-1] - start store message
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.808 - [origin:store] - [connectionID:dlm-1] - register receiver success
May 19 09:53:28 Moxa dlm-agent[28271]: [INFO]May 19 09:53:28.808 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 09:53:29 Moxa dlm-agent[28271]: [INFO]May 19 09:53:29.363 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 09:53:29 Moxa dlm-agent[28271]: [INFO]May 19 09:53:29.950 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 09:53:29 Moxa dlm-agent[28271]: [INFO]May 19 09:53:29.952 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 09:53:30 Moxa dlm-agent[28271]: [INFO]May 19 09:53:30.115 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 09:53:30 Moxa dlm-agent[28271]: [INFO]May 19 09:53:30.155 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=2cHuJXjLRBQhxMLhEUbsnZ2yDg8%3D&Expires=1621425206"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=LYGWkLKiVcGvm8b%2FmMHwSn4JOus%3D&Expires=1621425206"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=MM9abWqRKG0ipeZeL8%2F7gmbIudA%3D&Expires=1621425206"}}
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.156 - [origin:dlm-1] - [twin callback] exit
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.157 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.158 - [origin:dlm-1] - desired push type:$completeDesired, queue: 0/50
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.158 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=2cHuJXjLRBQhxMLhEUbsnZ2yDg8%3D&Expires=1621425206"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=LYGWkLKiVcGvm8b%2FmMHwSn4JOus%3D&Expires=1621425206"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=MM9abWqRKG0ipeZeL8%2F7gmbIudA%3D&Expires=1621425206"}}
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.160 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.160 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.160 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     348.034µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:53:30 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:30.173 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:53:30 Moxa dlm-agent[28271]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 09:53:30 Moxa dlm-agent[28271]: toolbox.NewTask
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     772.707µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |    5.042838ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |   16.132017ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:53:30 Moxa dlm-agent[28271]: [INFO]May 19 09:53:30.201 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:53:30 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |   35.227585ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |   37.465109ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:53:30 Moxa dlm-agent[28271]: GetDeviceGeneral
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |    16.52725ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |   60.624011ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     649.669µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     386.114µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     921.865µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     273.195µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |      605.35µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |    8.924455ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |     160.038µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 09:53:30 Moxa dlm-agent[28271]: [INFO]May 19 09:53:30.438 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 09:53:30 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:30 | 200 |    28.79953ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:53:31 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:31.166 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"speed":"100 Mbps","subnet":"10.123.12.0","dns":["10.123.200.11","10.123.200.12"],"enable":true,"gateway":"10.123.12.1","id":1,"name":"eth0","netmask":"255.255.254.0","broadcast":"10.123.13.255","displayName":"eth0","ip":"10.123.13.23","mac":"00:90:e8:82:f2:fb","status":"connected"},{"mac":"00:90:e8:82:f2:fc","netmask":"255.255.255.0","speed":"--","status":"disconnected","displayName":"eth1","dns":[],"gateway":"","ip":"192.168.4.127","subnet":"192.168.4.0","broadcast":"192.168.4.255","enable":true,"id":2,"name":"eth1"}]}
May 19 09:53:31 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:31 | 200 |     236.836µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:53:31 Moxa dlm-agent[28271]: [INFO]May 19 09:53:31.202 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 09:53:31 Moxa dlm-agent[28271]: [INFO]May 19 09:53:31.532 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 09:53:31 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:31.536 - [origin:dlm-1] - [command callback] $aws/things/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/jobs/$next/get/accepted {"timestamp":1621418011}
May 19 09:53:31 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:31.849 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:53:31 Moxa dlm-agent[28271]: [GIN] 2021/05/19 - 09:53:31 | 200 |     696.709µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:53:32 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:32.537 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"size":0,"version":"","detail":{"depends":null,"description":"","label":"","origin":""},"name":"","newVersion":""}}
May 19 09:53:32 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:32.538 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 09:53:32 Moxa dlm-agent[28271]: [INFO]May 19 09:53:32.541 - [origin:dlm-1] - [reported], {"osPatchControl":{"status":"ready"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=LYGWkLKiVcGvm8b%2FmMHwSn4JOus%3D\u0026Expires=1621425206"},"edition":"2.0","sshserver":{"enable":true,"port":22},"general":{"hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","modelName":"UC-3111-T-EU-LX","osType":"Linux","serialNumber":"TAIJB1072828","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)","firmwareVersion":"1.4","thingsProProduct":""},"osPatchCronJob":{"updateSchedule":"0 14 * * 0,1,2,3,4,5,6","enable":true},"osPatchDetail":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=MM9abWqRKG0ipeZeL8%2F7gmbIudA%3D\u0026Expires=1621425206","status":"success","version":"","desiredURL":""},"application-dlm-dashboard":{"version":"","desiredURL":"","reportedURL":"","status":""},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/b2217f42-eb3f-4932-b9ff-0354f5ac92ee/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=2cHuJXjLRBQhxMLhEUbsnZ2yDg8%3D\u0026Expires=1621425206"},"applications":{"serialNumber":"TAIJB1072828","cpu":"ARMv7 Processor rev 2 (v7l)","firmwareVersion":"1.4","hostName":"Moxa","osType":"Linux","memorySize":524054528,"modelName":"UC-3111-T-EU-LX","thingsProProduct":"","biosVersion":"1.4.0S02","description":"","disk":[{"name":"System","percent":82.16348920732285,"total":449529856,"used":346475520,"device":"/dev/root","free":75214848,"mount":"/"}],"lastBootTime":"2021-05-17T03:13:07Z"}}
May 19 09:53:32 Moxa dlm-agent[28271]: [INFO]May 19 09:53:32.554 - [origin:dlm-1] - [reported] applications
May 19 09:53:32 Moxa dlm-agent[28271]: [INFO]May 19 09:53:32.720 - [origin:dlm-1] - [reported callback], success
May 19 09:53:32 Moxa dlm-agent[28271]: [INFO]May 19 09:53:32.721 - [origin:dlm-1] - [reported] classic
May 19 09:53:32 Moxa dlm-agent[28271]: [INFO]May 19 09:53:32.890 - [origin:dlm-1] - [reported callback], success
May 19 09:53:58 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:58.519 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:53:58 Moxa dlm-agent[28271]: [DBUG]May 19 09:53:58.522 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
```

