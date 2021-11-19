```
-- Logs begin at Mon 2021-05-17 07:46:58 UTC. --
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 09:32:14 Moxa dlm-agent[20708]: [INFO]May 19 09:32:13.968 - [origin:remoted] - service dlm started
May 19 09:32:14 Moxa dlm-agent[20708]: [INFO]May 19 09:32:13.968 - [origin:remoted] - all started
May 19 09:32:14 Moxa dlm-agent[20708]: [INFO]May 19 09:32:13.956 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:32:14 Moxa dlm-agent[20708]: [INFO]May 19 09:32:13.957 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:33:50 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:50 | 200 |    1.139102ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:33:50 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:50 | 200 |     219.957µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:33:50 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:50 | 200 |     230.236µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:33:51 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:51 | 200 |     313.075µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:33:51 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:51 | 200 |     322.194µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:33:51 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:33:51 | 200 |     220.116µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:34:00 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:00 | 200 |     217.476µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:34:04 Moxa dlm-agent[20708]: [INFO]May 19 09:34:04.189 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:34:04 Moxa dlm-agent[20708]: [INFO]May 19 09:34:04.855 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:34:04 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:04 | 200 |  1.514870386s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"
May 19 09:34:08 Moxa dlm-agent[20708]: [INFO]May 19 09:34:08.051 - [origin:dlm] - skip stop because connection status: false
May 19 09:34:08 Moxa dlm-agent[20708]: [INFO]May 19 09:34:08.792 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:34:09 Moxa dlm-agent[20708]: [INFO]May 19 09:34:09.439 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:34:09 Moxa dlm-agent[20708]: [INFO]May 19 09:34:09.441 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 09:34:09 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:09 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:09 | 200 |   15.183397ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:34:09 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:09 | 200 |    66.88357ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 09:34:09 Moxa dlm-agent[20708]: [INFO]May 19 09:34:09.541 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 09:34:09 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:09 | 200 |     430.513µs |       127.0.0.1 | GET      "/api/v1/_/info"

May 19 09:34:12 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:12.010 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 09:34:12 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:12.013 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.054 - [origin:dlm] - latest configuration:  {
May 19 09:34:12 Moxa dlm-agent[20708]:   "connection": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "enable": false,
May 19 09:34:12 Moxa dlm-agent[20708]:     "retryDelaySec": 10,
May 19 09:34:12 Moxa dlm-agent[20708]:     "picTarget": "stage"
May 19 09:34:12 Moxa dlm-agent[20708]:   },
May 19 09:34:12 Moxa dlm-agent[20708]:   "certificates": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "caCertFile": "dev.crt",
May 19 09:34:12 Moxa dlm-agent[20708]:     "caPkFile": "dev.key",
May 19 09:34:12 Moxa dlm-agent[20708]:     "certificateInfo": {
May 19 09:34:12 Moxa dlm-agent[20708]:       "organization": "Moxa Inc.",
May 19 09:34:12 Moxa dlm-agent[20708]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:34:12 Moxa dlm-agent[20708]:       "notBefore": "May 19, 2021 09:34:10",
May 19 09:34:12 Moxa dlm-agent[20708]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:34:12 Moxa dlm-agent[20708]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:34:12 Moxa dlm-agent[20708]:       "macAddress": "0090E882F2FB",
May 19 09:34:12 Moxa dlm-agent[20708]:       "serialNumber": "TAIJB1072828"
May 19 09:34:12 Moxa dlm-agent[20708]:     }
May 19 09:34:12 Moxa dlm-agent[20708]:   },
May 19 09:34:12 Moxa dlm-agent[20708]:   "connectionStatus": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "projectName": "Kevin Test - AGT",
May 19 09:34:12 Moxa dlm-agent[20708]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:34:12 Moxa dlm-agent[20708]:     "status": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "message": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastConnectedTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastConnectTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastDisconnectTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "statuslastUpdateTime": ""
May 19 09:34:12 Moxa dlm-agent[20708]:   }
May 19 09:34:12 Moxa dlm-agent[20708]: }
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.063 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.063 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.064 - [origin:dlm] - latest configuration:  {
May 19 09:34:12 Moxa dlm-agent[20708]:   "connection": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "enable": true,
May 19 09:34:12 Moxa dlm-agent[20708]:     "retryDelaySec": 10,
May 19 09:34:12 Moxa dlm-agent[20708]:     "picTarget": "stage"
May 19 09:34:12 Moxa dlm-agent[20708]:   },
May 19 09:34:12 Moxa dlm-agent[20708]:   "certificates": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "caCertFile": "dev.crt",
May 19 09:34:12 Moxa dlm-agent[20708]:     "caPkFile": "dev.key",
May 19 09:34:12 Moxa dlm-agent[20708]:     "certificateInfo": {
May 19 09:34:12 Moxa dlm-agent[20708]:       "organization": "Moxa Inc.",
May 19 09:34:12 Moxa dlm-agent[20708]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:34:12 Moxa dlm-agent[20708]:       "notBefore": "May 19, 2021 09:34:10",
May 19 09:34:12 Moxa dlm-agent[20708]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:34:12 Moxa dlm-agent[20708]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:34:12 Moxa dlm-agent[20708]:       "macAddress": "0090E882F2FB",
May 19 09:34:12 Moxa dlm-agent[20708]:       "serialNumber": "TAIJB1072828"
May 19 09:34:12 Moxa dlm-agent[20708]:     }
May 19 09:34:12 Moxa dlm-agent[20708]:   },
May 19 09:34:12 Moxa dlm-agent[20708]:   "connectionStatus": {
May 19 09:34:12 Moxa dlm-agent[20708]:     "projectName": "Kevin Test - AGT",
May 19 09:34:12 Moxa dlm-agent[20708]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:34:12 Moxa dlm-agent[20708]:     "status": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "message": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastConnectedTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastConnectTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "lastDisconnectTime": "",
May 19 09:34:12 Moxa dlm-agent[20708]:     "statuslastUpdateTime": ""
May 19 09:34:12 Moxa dlm-agent[20708]:   }
May 19 09:34:12 Moxa dlm-agent[20708]: }
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.066 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:12 | 200 |    3.083391ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.088 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.088 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.088 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [ERRO]May 19 09:34:12.091 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 09:34:12 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:12 | 404 |         9.6µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 09:34:12 Moxa dlm-agent[20708]: [ERRO]May 19 09:34:12.097 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.097 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.097 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [ERRO]May 19 09:34:12.097 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.097 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.097 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.098 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:12 | 200 |  4.047451038s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.110 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.110 - [origin:dlm] - register target: stage
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.110 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 09:34:12 Moxa dlm-agent[20708]: [INFO]May 19 09:34:12.110 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.074 - [origin:dlm] - tasks[3/6] name:check certificate, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.827 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.830 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.832 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.841 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:13.843 - [origin:dlm-1] - try to connect rule server ...
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.851 - [origin:dlm-1] - connect to rule server success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.859 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.859 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.859 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.860 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 09:34:13 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:13.856 - [origin:dlm-1] - monitor stream in
May 19 09:34:13 Moxa dlm-agent[20708]: [WARN]May 19 09:34:13.860 - [origin:dlm-1] - waiting for self registration...
May 19 09:34:13 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:13.856 - [origin:dlm-1] - twin stream in
May 19 09:34:13 Moxa dlm-agent[20708]: [WARN]May 19 09:34:13.860 - [origin:dlm-1] - waiting for self registration...
May 19 09:34:13 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:13.857 - [origin:dlm-1] - message stream in
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.879 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 09:34:13 Moxa dlm-agent[20708]: [INFO]May 19 09:34:13.994 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:34:13 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:13.994 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.000 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.006 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.003 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.017 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.018 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.018 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.018 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.019 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.005 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.022 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.023 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.023 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.005 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.041 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.005 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.041 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.035 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.044 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.045 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.045 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.045 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.056 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.056 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.058 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.058 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.059 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.060 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.060 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:34:14 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:14.061 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.061 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.861 - [origin:dlm-1] - monitor stream registration success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.863 - [origin:dlm-1] - twin stream registration success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.979 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.979 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 09:34:14 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.979 - [origin:dlm-1] - [connection callback], connection = true
May 19 09:34:15 Moxa dlm-agent[20708]: [INFO]May 19 09:34:14.983 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 09:34:15 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:14 | 200 |     234.717µs |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 09:34:15 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:14 | 200 |    2.790276ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:34:15 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:15 | 200 |    4.947401ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 09:34:15 Moxa dlm-agent[20708]: [INFO]May 19 09:34:15.024 - [origin:message-group:dlm-1:2] - stopping
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.045 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.046 - [origin:message-group:dlm-1:2] - message-routine:exit, reason: source canceled
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - tasks[1/5] name:stop ipc connection, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - outputservice exit
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - outputservice state: stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - timer stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - formater stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - tasks[2/5] name:stop output control service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.047 - [origin:message-group:dlm-1:2] - tasks[3/5] name:stop tag service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.048 - [origin:message-group:dlm-1:2] - event origin:OS patch listener exit
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.048 - [origin:message-group:dlm-1:2] - tasks[4/5] name:stop event service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.048 - [origin:message-group:dlm-1:2] - tasks[5/5] name:stop remote monitor service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.048 - [origin:message-group:dlm-1:2] - stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.062 - [origin:message-group:dlm-1:2] - message:exit, reason: source canceled
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.062 - [origin:message-group:dlm-1:2] - message stream out
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   2.98865611s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process"
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |  2.042131184s |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.066 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.066 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.067 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |     3.75558ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.085 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.086 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.086 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.086 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.086 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.086 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.088 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.088 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.105 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.105 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.106 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.106 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.107 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.107 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.121 - [origin:message-group:dlm-1:4] - stopping
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.121 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - message-routine:exit, reason: source canceled
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - tasks[1/5] name:stop ipc connection, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - outputservice exit
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - outputservice state: stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - timer stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.122 - [origin:message-group:dlm-1:4] - formater stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.123 - [origin:message-group:dlm-1:4] - tasks[2/5] name:stop output control service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.123 - [origin:message-group:dlm-1:4] - tasks[3/5] name:stop tag service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.123 - [origin:message-group:dlm-1:4] - event origin:OS patch listener exit
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.123 - [origin:message-group:dlm-1:4] - tasks[4/5] name:stop event service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.123 - [origin:message-group:dlm-1:4] - tasks[5/5] name:stop remote monitor service, result:success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.124 - [origin:message-group:dlm-1:4] - stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   21.477536ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.130 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.130 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.131 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |  3.091561423s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available"
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.135 - [origin:message-group:dlm-1:4] - message:exit, reason: source canceled
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.135 - [origin:message-group:dlm-1:4] - message stream out
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.151 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.151 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.152 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.152 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.152 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.152 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.153 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.154 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.188 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.188 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.188 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.188 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.189 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.189 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:34:17 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   31.594014ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   48.681181ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.223 - [origin:store] - [connectionID:dlm-1] - start
May 19 09:34:17 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   39.704005ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:34:17 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   24.606126ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |    59.29509ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.288 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 09:34:17 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:17.321 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.323 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.324 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: waiting
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.324 - [origin:store] - [connectionID:dlm-1] - get connection success, start register
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.328 - [origin:dlm-1] - start store message
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.343 - [origin:store] - [connectionID:dlm-1] - register receiver success
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.344 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 09:34:17 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:17 | 200 |   83.007391ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.376 - [origin:dlm] - last ota result empty
May 19 09:34:17 Moxa dlm-agent[20708]: [INFO]May 19 09:34:17.913 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.499 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.502 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.651 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.689 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=RnQpcYqMV5OlcTPrUWAVrLfVNLA%3D&Expires=1621424055"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=GWq8XZFpzycjmMbXiPTsdJxT1vs%3D&Expires=1621424055"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=XFhCKh4xW6Ik6bpsoPJssgUVJ4Y%3D&Expires=1621424055"}}
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.690 - [origin:dlm-1] - [twin callback] exit
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.692 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.693 - [origin:dlm-1] - desired push type:$completeDesired, queue: 0/50
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.694 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=RnQpcYqMV5OlcTPrUWAVrLfVNLA%3D&Expires=1621424055"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=GWq8XZFpzycjmMbXiPTsdJxT1vs%3D&Expires=1621424055"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=XFhCKh4xW6Ik6bpsoPJssgUVJ4Y%3D&Expires=1621424055"}}
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.703 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.703 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.703 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     359.514µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:34:18 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:18.710 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:34:18 Moxa dlm-agent[20708]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 09:34:18 Moxa dlm-agent[20708]: toolbox.NewTask
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     192.036µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |    3.056591ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   14.279811ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.739 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:34:18 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   20.502712ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   27.766155ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:34:18 Moxa dlm-agent[20708]: GetDeviceGeneral
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   15.923945ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   59.653284ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     481.552µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     443.193µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     895.746µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     304.795µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     599.911µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |    9.076775ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |     177.677µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 09:34:18 Moxa dlm-agent[20708]: [INFO]May 19 09:34:18.955 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 09:34:18 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:18 | 200 |   30.796427ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:34:19 Moxa dlm-agent[20708]: [INFO]May 19 09:34:19.676 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 09:34:19 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:19.680 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"ip":"10.123.13.23","netmask":"255.255.254.0","subnet":"10.123.12.0","displayName":"eth0","dns":["10.123.200.11","10.123.200.12"],"enable":true,"gateway":"10.123.12.1","id":1,"mac":"00:90:e8:82:f2:fb","name":"eth0","speed":"100 Mbps","broadcast":"10.123.13.255","status":"connected"},{"broadcast":"192.168.4.255","dns":[],"id":2,"mac":"00:90:e8:82:f2:fc","speed":"--","displayName":"eth1","enable":true,"gateway":"","ip":"192.168.4.127","name":"eth1","netmask":"255.255.255.0","status":"disconnected","subnet":"192.168.4.0"}]}
May 19 09:34:19 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:19 | 200 |     157.878µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:34:19 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:19.992 - [origin:dlm-1] - [command callback] $aws/things/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/jobs/$next/get/accepted {"timestamp":1621416859}
May 19 09:34:20 Moxa dlm-agent[20708]: [INFO]May 19 09:34:20.007 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 09:34:20 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:20.336 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:34:20 Moxa dlm-agent[20708]: [GIN] 2021/05/19 - 09:34:20 | 200 |     272.116µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:34:21 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:21.001 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"name":"","newVersion":"","size":0,"version":"","detail":{"depends":null,"description":"","label":"","origin":""}}}
May 19 09:34:21 Moxa dlm-agent[20708]: [DBUG]May 19 09:34:21.006 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 09:34:21 Moxa dlm-agent[20708]: [INFO]May 19 09:34:21.011 - [origin:dlm-1] - [reported], {"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=RnQpcYqMV5OlcTPrUWAVrLfVNLA%3D\u0026Expires=1621424055"},"applications":{"disk":[{"percent":82.16348920732285,"total":449529856,"used":346475520,"device":"/dev/root","free":75214848,"mount":"/","name":"System"}],"biosVersion":"1.4.0S02","description":"","hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","memorySize":524054528,"modelName":"UC-3111-T-EU-LX","osType":"Linux","serialNumber":"TAIJB1072828","cpu":"ARMv7 Processor rev 2 (v7l)","firmwareVersion":"1.4","thingsProProduct":""},"osPatchControl":{"status":"ready"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=GWq8XZFpzycjmMbXiPTsdJxT1vs%3D\u0026Expires=1621424055"},"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/eabe5b13-e37b-4ca7-b739-ca7eb507eac6/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=XFhCKh4xW6Ik6bpsoPJssgUVJ4Y%3D\u0026Expires=1621424055"},"application-dlm-dashboard":{"version":"","desiredURL":"","reportedURL":"","status":""},"edition":"2.0","general":{"biosVersion":"1.4.0S02","firmwareVersion":"1.4","modelName":"UC-3111-T-EU-LX","osType":"Linux","thingsProProduct":"","cpu":"ARMv7 Processor rev 2 (v7l)","hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","serialNumber":"TAIJB1072828"},"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"sshserver":{"port":22,"enable":true}}
May 19 09:34:21 Moxa dlm-agent[20708]: [INFO]May 19 09:34:21.017 - [origin:dlm-1] - [reported] classic
May 19 09:34:21 Moxa dlm-agent[20708]: [INFO]May 19 09:34:21.171 - [origin:dlm-1] - [reported callback], success
May 19 09:34:21 Moxa dlm-agent[20708]: [INFO]May 19 09:34:21.171 - [origin:dlm-1] - [reported] applications
May 19 09:34:21 Moxa dlm-agent[20708]: [INFO]May 19 09:34:21.321 - [origin:dlm-1] - [reported callback], success
```

