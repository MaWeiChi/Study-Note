```
May 19 09:25:50 Moxa dlm-agent[17597]: *******************************************************************************
May 19 09:25:50 Moxa dlm-agent[17597]:   _______ _     _                 _____             _____  _                 _
May 19 09:25:50 Moxa dlm-agent[17597]:  |__   __| |   (_)               |  __ \           |  ___|| |               | |
May 19 09:25:50 Moxa dlm-agent[17597]:     | |  | |__  _ _ __   __ _ ___| |__) | __ ___   | |    | | ___  _   _  __| |
May 19 09:25:50 Moxa dlm-agent[17597]:     | |  | '_ \| | '_ \ / _  / __|  ___/ '__/ _ \  | |    | |/ _ \| | | |/ _  |
May 19 09:25:50 Moxa dlm-agent[17597]:     | |  | | | | | | | | (_| \__ \ |   | | | (_) | | |___ | | (_) | |_| | (_| |
May 19 09:25:50 Moxa dlm-agent[17597]:     |_|  |_| |_|_|_| |_|\__, |___/_|   |_|  \___/  |_____||_|\___/\_____/\__,_|
May 19 09:25:50 Moxa dlm-agent[17597]:                          __/ |
May 19 09:25:50 Moxa dlm-agent[17597]:                         |___/
May 19 09:25:50 Moxa dlm-agent[17597]:   Application : ThingsPro Cloud service
May 19 09:25:50 Moxa dlm-agent[17597]:   Version     : 0.1.2-armhf, Time : 2021-05-19T09:25:50Z
May 19 09:25:50 Moxa dlm-agent[17597]: *******************************************************************************
May 19 09:25:50 Moxa dlm-agent[17597]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:25:50 Moxa dlm-agent[17597]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:25:50 Moxa dlm-agent[17597]:  - using env:        export GIN_MODE=release
May 19 09:25:50 Moxa dlm-agent[17597]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:25:50 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/_/info            --> main.getInfo (4 handlers)
May 19 09:25:50 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/devicedb/onchange --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/agent.TransOnChange (4 handlers)
May 19 09:25:50 Moxa dlm-agent[17597]: [openvpn] 2021/05/19 09:25:50 config.go:52: main config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json: no such file or directory,<nil>, try load backup
May 19 09:25:50 Moxa dlm-agent[17597]: [openvpn] 2021/05/19 09:25:50 config.go:62: backup config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json.backup: no such file or directory,<nil>, try load factory
May 19 09:25:50 Moxa dlm-agent[17597]: MOXA api token open failed, open /run/mx-api-token: no such file or directory
May 19 09:25:51 Moxa dlm-agent[17597]: config load err:open /etc/dlm-agent/.data/device/openvpn/openvpn.json: no such file or directory
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/device/general    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/device/applications --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/device/ethernets  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernets-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/device/ethernets/:id --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernetById-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/system/reboot     --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PutSystemReboot-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/openvpn           --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetOpenVpn-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] PATCH  /api/v1/openvpn/dlm       --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PatchOpenVpnDlm-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetSSHSettings-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).SetSSHSettings-fm (4 handlers)
May 19 09:25:51 Moxa dlm-agent[17597]: config load err:open /etc/dlm-agent/.data/device/software/todo_UpgradeList.json: no such file or directory
May 19 09:25:51 Moxa dlm-agent[17597]: get todo upgrade list error: unexpected end of JSON input
May 19 09:25:51 Moxa dlm-agent[17597]: config load err:open /etc/dlm-agent/.data/device/software/config.json: no such file or directory
May 19 09:25:51 Moxa dlm-agent[17597]: 0 0 * * * *
May 19 09:25:51 Moxa dlm-agent[17597]: toolbox.NewTask
May 19 09:25:51 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftware-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftware-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/software/ospatchlist --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetPatchList-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/software/detail   --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareDetail-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareSchedule-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftwareSchedule-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetLogs (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PostLog (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetProfile (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PatchProfile (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/tags/monitor      --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetMonitor-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/tags/monitor/:args --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).PostMonitor-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/tags/monitor/system/system --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetTags-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:25:52 Moxa dlm-agent[17597]:  - using env:        export GIN_MODE=release
May 19 09:25:52 Moxa dlm-agent[17597]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.012 - [origin:remoted] - plugin message uploader created
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.012 - [origin:remoted] - plugin devicedb created
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.013 - [origin:remoted] - plugin appman created
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.014 - [origin:remoted] - plugin apiV1 created
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.015 - [origin:remoted] - plugin store created
May 19 09:25:52 Moxa dlm-agent[17597]: [DBUG]May 19 09:25:52.017 - [origin:dlm] - cert path:
May 19 09:25:52 Moxa dlm-agent[17597]: [DBUG]May 19 09:25:52.017 - [origin:dlm] - key path:
May 19 09:25:52 Moxa dlm-agent[17597]: [ERRO]May 19 09:25:52.154 - [origin:dlm] - tasks[1/2] name:[armhf] check cert files, result:failed, error:exit status 1
May 19 09:25:52 Moxa dlm-agent[17597]: [WARN]May 19 09:25:52.154 - [origin:dlm] - load factory certificates failed, error: exit status 1
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.155 - [origin:remoted] - service dlm created
May 19 09:25:52 Moxa dlm-agent[17597]: [DBUG]May 19 09:25:52.159 - [origin:devicedb:dlm-1] - database subscriber start
May 19 09:25:52 Moxa dlm-agent[17597]: [DBUG]May 19 09:25:52.159 - [origin:devicedb:dlm-1] - database subscriber exit, channel empty
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.159 - [origin:devicedb] - connectionID:dlm-1 db client started
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSource-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSource-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSelections-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSelections-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/onchange --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesOnchange-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/test --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).testPropertiesSelections-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).reported-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/custom/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).userReported-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/clear --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).clear-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.160 - [origin:remoted] - plugin devicedb started
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.160 - [origin:remoted] - plugin appman started
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.160 - [origin:remoted] - plugin apiV1 started
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores/restart --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).reload-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.161 - [origin:remoted] - plugin store started
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.166 - [origin:remoted] - worker sock binding
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).getMessagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).putMessagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).postMessagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] DELETE /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.177 - [origin:remoted] - plugin message uploader started
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] HEAD   /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] HEAD   /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] HEAD   /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] HEAD   /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/info    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).info-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/tags    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).tagMonitor-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/remotectl/services --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).servicesInfo-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/message --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).sendMessage-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/state --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).updateState-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.181 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.182 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.178 - [origin:dlm] - skip start, false
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).postConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PATCH  /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).patchConfiguration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/_/device      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDeviceID-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/monitor       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getMonitor-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/registration  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putRegistration-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/enrollment    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putEnrollment-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/edition       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getEdition-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/control       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).controlThing-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/login         --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).login-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] POST   /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] DELETE /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDashboard-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putDashboard-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.192 - [origin:remoted] - service dlm started
May 19 09:25:52 Moxa dlm-agent[17597]: [INFO]May 19 09:25:52.192 - [origin:remoted] - all started
May 19 09:26:00 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:00 | 200 |    1.115582ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:00 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:00 | 200 |     218.917µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:00 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:00 | 200 |     222.316µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:01 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:01 | 200 |     311.075µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:01 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:01 | 200 |     219.077µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:01 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:01 | 200 |     218.397µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:11 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:11 | 200 |     220.837µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:26:15 Moxa dlm-agent[17597]: [INFO]May 19 09:26:15.665 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:26:16 Moxa dlm-agent[17597]: [INFO]May 19 09:26:16.291 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:26:16 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:16 | 200 |  1.507422319s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"
May 19 09:26:17 Moxa dlm-agent[17597]: [INFO]May 19 09:26:17.676 - [origin:dlm] - skip stop because connection status: false
May 19 09:26:18 Moxa dlm-agent[17597]: [INFO]May 19 09:26:18.483 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:26:19 Moxa dlm-agent[17597]: [INFO]May 19 09:26:19.140 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:26:19 Moxa dlm-agent[17597]: [INFO]May 19 09:26:19.143 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 09:26:19 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:19 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:19 | 200 |    15.01696ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:26:19 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:19 | 200 |   66.597094ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 09:26:19 Moxa dlm-agent[17597]: [INFO]May 19 09:26:19.242 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 09:26:19 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:19 | 200 |     300.915µs |       127.0.0.1 | GET      "/api/v1/_/info"
May 19 09:26:21 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:21.694 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 09:26:21 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:21.697 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.774 - [origin:dlm] - latest configuration:  {
May 19 09:26:21 Moxa dlm-agent[17597]:   "connection": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "enable": false,
May 19 09:26:21 Moxa dlm-agent[17597]:     "retryDelaySec": 10,
May 19 09:26:21 Moxa dlm-agent[17597]:     "picTarget": "stage"
May 19 09:26:21 Moxa dlm-agent[17597]:   },
May 19 09:26:21 Moxa dlm-agent[17597]:   "certificates": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "caCertFile": "dev.crt",
May 19 09:26:21 Moxa dlm-agent[17597]:     "caPkFile": "dev.key",
May 19 09:26:21 Moxa dlm-agent[17597]:     "certificateInfo": {
May 19 09:26:21 Moxa dlm-agent[17597]:       "organization": "Moxa Inc.",
May 19 09:26:21 Moxa dlm-agent[17597]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:26:21 Moxa dlm-agent[17597]:       "notBefore": "May 19, 2021 09:26:19",
May 19 09:26:21 Moxa dlm-agent[17597]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:26:21 Moxa dlm-agent[17597]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:26:21 Moxa dlm-agent[17597]:       "macAddress": "0090E882F2FB",
May 19 09:26:21 Moxa dlm-agent[17597]:       "serialNumber": "TAIJB1072828"
May 19 09:26:21 Moxa dlm-agent[17597]:     }
May 19 09:26:21 Moxa dlm-agent[17597]:   },
May 19 09:26:21 Moxa dlm-agent[17597]:   "connectionStatus": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "projectName": "Kevin Test - AGT",
May 19 09:26:21 Moxa dlm-agent[17597]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:26:21 Moxa dlm-agent[17597]:     "status": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "message": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastConnectedTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastConnectTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastDisconnectTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "statuslastUpdateTime": ""
May 19 09:26:21 Moxa dlm-agent[17597]:   }
May 19 09:26:21 Moxa dlm-agent[17597]: }
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.783 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.783 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.784 - [origin:dlm] - latest configuration:  {
May 19 09:26:21 Moxa dlm-agent[17597]:   "connection": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "enable": true,
May 19 09:26:21 Moxa dlm-agent[17597]:     "retryDelaySec": 10,
May 19 09:26:21 Moxa dlm-agent[17597]:     "picTarget": "stage"
May 19 09:26:21 Moxa dlm-agent[17597]:   },
May 19 09:26:21 Moxa dlm-agent[17597]:   "certificates": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "caCertFile": "dev.crt",
May 19 09:26:21 Moxa dlm-agent[17597]:     "caPkFile": "dev.key",
May 19 09:26:21 Moxa dlm-agent[17597]:     "certificateInfo": {
May 19 09:26:21 Moxa dlm-agent[17597]:       "organization": "Moxa Inc.",
May 19 09:26:21 Moxa dlm-agent[17597]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:26:21 Moxa dlm-agent[17597]:       "notBefore": "May 19, 2021 09:26:19",
May 19 09:26:21 Moxa dlm-agent[17597]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:26:21 Moxa dlm-agent[17597]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:26:21 Moxa dlm-agent[17597]:       "macAddress": "0090E882F2FB",
May 19 09:26:21 Moxa dlm-agent[17597]:       "serialNumber": "TAIJB1072828"
May 19 09:26:21 Moxa dlm-agent[17597]:     }
May 19 09:26:21 Moxa dlm-agent[17597]:   },
May 19 09:26:21 Moxa dlm-agent[17597]:   "connectionStatus": {
May 19 09:26:21 Moxa dlm-agent[17597]:     "projectName": "Kevin Test - AGT",
May 19 09:26:21 Moxa dlm-agent[17597]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:26:21 Moxa dlm-agent[17597]:     "status": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "message": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastConnectedTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastConnectTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "lastDisconnectTime": "",
May 19 09:26:21 Moxa dlm-agent[17597]:     "statuslastUpdateTime": ""
May 19 09:26:21 Moxa dlm-agent[17597]:   }
May 19 09:26:21 Moxa dlm-agent[17597]: }
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.787 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 09:26:21 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:21 | 200 |    1.155661ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.822 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.822 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.823 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 09:26:21 Moxa dlm-agent[17597]: [ERRO]May 19 09:26:21.825 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 09:26:21 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:21 | 404 |        9.48µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 09:26:21 Moxa dlm-agent[17597]: [ERRO]May 19 09:26:21.827 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 09:26:21 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.827 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.827 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [ERRO]May 19 09:26:21.827 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.827 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.827 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.827 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.841 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.842 - [origin:dlm] - register target: stage
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.842 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 09:26:22 Moxa dlm-agent[17597]: [INFO]May 19 09:26:21.842 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success
May 19 09:26:22 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:21 | 200 |  4.152368633s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"
May 19 09:26:24 Moxa dlm-agent[17597]: [INFO]May 19 09:26:24.379 - [origin:dlm] - tasks[3/6] name:check certificate, result:success
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.163 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.166 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.168 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.177 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.179 - [origin:dlm-1] - try to connect rule server ...
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.186 - [origin:dlm-1] - connect to rule server success
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.189 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.189 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.189 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.200 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.203 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.214 - [origin:dlm-1] - monitor stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [WARN]May 19 09:26:25.216 - [origin:dlm-1] - waiting for self registration...
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.217 - [origin:dlm-1] - twin stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [WARN]May 19 09:26:25.223 - [origin:dlm-1] - waiting for self registration...
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.218 - [origin:dlm-1] - message stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.228 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.228 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.234 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.234 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.236 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.236 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.237 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.238 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.239 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.240 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.246 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.247 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.247 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.247 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:26:25 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:25.248 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:26:25 Moxa dlm-agent[17597]: [INFO]May 19 09:26:25.248 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.054 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.057 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.057 - [origin:dlm-1] - [connection callback], connection = true
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.067 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 09:26:26 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:26 | 200 |    2.604119ms |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 09:26:26 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:26 | 200 |    1.096582ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:26:26 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:26 | 200 |    4.115894ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.103 - [origin:message-group:dlm-1:2] - stopping
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.106 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.107 - [origin:message-group:dlm-1:2] - stopped
May 19 09:26:26 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:26 | 200 |   10.674109ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.113 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.113 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.113 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.129 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.132 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.134 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.134 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.134 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.135 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.136 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.136 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:26:26 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:26 | 200 |   12.927593ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.152 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.152 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.154 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.154 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:26:26 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:26.154 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.154 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.173 - [origin:message-group:dlm-1:4] - stopping
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.222 - [origin:dlm-1] - monitor stream registration success
May 19 09:26:26 Moxa dlm-agent[17597]: [INFO]May 19 09:26:26.223 - [origin:dlm-1] - twin stream registration success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.236 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.238 - [origin:message-group:dlm-1:4] - message-routine:exit, reason: source canceled
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.238 - [origin:message-group:dlm-1:4] - tasks[1/5] name:stop ipc connection, result:success
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.238 - [origin:message-group:dlm-1:4] - outputservice exit
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.238 - [origin:message-group:dlm-1:4] - outputservice state: stopped
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.238 - [origin:message-group:dlm-1:4] - timer stopped
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.239 - [origin:message-group:dlm-1:4] - formater stopped
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.239 - [origin:message-group:dlm-1:4] - tasks[2/5] name:stop output control service, result:success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.239 - [origin:message-group:dlm-1:4] - tasks[3/5] name:stop tag service, result:success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.240 - [origin:message-group:dlm-1:4] - event origin:OS patch listener exit
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.242 - [origin:message-group:dlm-1:4] - tasks[4/5] name:stop event service, result:success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.242 - [origin:message-group:dlm-1:4] - tasks[5/5] name:stop remote monitor service, result:success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.243 - [origin:message-group:dlm-1:4] - stopped
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |  2.996551809s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available"
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.241 - [origin:message-group:dlm-1:4] - message:exit, reason: source canceled
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.247 - [origin:message-group:dlm-1:4] - message stream out
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |  2.092696792s |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.253 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.255 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.255 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.287 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.287 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.288 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.288 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.288 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.289 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.290 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.291 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.310 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.310 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.315 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.316 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.316 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.316 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:26:28 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |   49.774323ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |   65.980064ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.342 - [origin:store] - [connectionID:dlm-1] - start
May 19 09:26:28 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |    44.97192ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:26:28 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |   23.819539ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |   83.495944ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.424 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 09:26:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:28.446 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.452 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.452 - [connectionID:dlm-1] - [origin:store] - connection[dlm-1] state: waiting
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.452 - [origin:store] - [connectionID:dlm-1] - get connection success, start register
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.463 - [origin:dlm-1] - start store message
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.466 - [connectionID:dlm-1] - [origin:store] - register receiver success
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.467 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 09:26:28 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:28 | 200 |   52.036327ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 09:26:28 Moxa dlm-agent[17597]: [INFO]May 19 09:26:28.480 - [origin:dlm] - last ota result empty
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.035 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.583 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.583 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.735 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.774 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=k8Fel7tJVqHClIWr1osl%2B%2F6skjA%3D&Expires=1621423586"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=MYsFeKVXqpe3%2BK644DbYjTi55gQ%3D&Expires=1621423586"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=0TvujGg0EPIWwkPqCHFsIMksQYg%3D&Expires=1621423586"}}
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.775 - [origin:dlm-1] - [twin callback] exit
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.776 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.776 - [origin:dlm-1] - desired push type:$completeDesired, queue: 0/50
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.776 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=k8Fel7tJVqHClIWr1osl%2B%2F6skjA%3D&Expires=1621423586"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=MYsFeKVXqpe3%2BK644DbYjTi55gQ%3D&Expires=1621423586"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=0TvujGg0EPIWwkPqCHFsIMksQYg%3D&Expires=1621423586"}}
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.778 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.779 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.779 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     321.555µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:26:29 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:29.792 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:26:29 Moxa dlm-agent[17597]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 09:26:29 Moxa dlm-agent[17597]: toolbox.NewTask
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     197.757µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |    2.981352ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |   14.154454ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:26:29 Moxa dlm-agent[17597]: [INFO]May 19 09:26:29.815 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:26:29 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |   24.221012ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |   27.250924ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:26:29 Moxa dlm-agent[17597]: GetDeviceGeneral
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |   16.187981ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |   59.578887ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     487.673µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     577.071µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     677.509µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:26:29 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:29 | 200 |     144.078µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:26:30 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:30 | 200 |     600.671µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:26:30 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:30 | 200 |    8.962377ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:26:30 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:30 | 200 |     161.038µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 09:26:30 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:30 | 200 |   24.314051ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:26:30 Moxa dlm-agent[17597]: [INFO]May 19 09:26:30.058 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 09:26:30 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:30.712 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"gateway":"10.123.12.1","ip":"10.123.13.23","name":"eth0","netmask":"255.255.254.0","displayName":"eth0","dns":["10.123.200.11","10.123.200.12"],"id":1,"mac":"00:90:e8:82:f2:fb","speed":"100 Mbps","status":"connected","subnet":"10.123.12.0","broadcast":"10.123.13.255","enable":true},{"displayName":"eth1","gateway":"","status":"disconnected","mac":"00:90:e8:82:f2:fc","name":"eth1","netmask":"255.255.255.0","broadcast":"192.168.4.255","dns":[],"enable":true,"id":2,"ip":"192.168.4.127","speed":"--","subnet":"192.168.4.0"}]}
May 19 09:26:30 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:30 | 200 |     229.396µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:26:30 Moxa dlm-agent[17597]: [INFO]May 19 09:26:30.795 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 09:26:31 Moxa dlm-agent[17597]: [INFO]May 19 09:26:31.106 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 09:26:31 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:31.109 - [origin:dlm-1] - [command callback] $aws/things/a249a099-554d-4e33-badf-8ad1c8372d04/jobs/$next/get/accepted {"timestamp":1621416391}
May 19 09:26:31 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:31.420 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:26:31 Moxa dlm-agent[17597]: [GIN] 2021/05/19 - 09:26:31 | 200 |     301.596µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:26:32 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:32.089 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"detail":{"description":"","label":"","origin":"","depends":null},"name":"","newVersion":"","size":0,"version":""}}
May 19 09:26:32 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:32.094 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 09:26:32 Moxa dlm-agent[17597]: [INFO]May 19 09:26:32.099 - [origin:dlm-1] - [reported], {"osPatchCronJob":{"updateSchedule":"0 14 * * 0,1,2,3,4,5,6","enable":true},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=MYsFeKVXqpe3%2BK644DbYjTi55gQ%3D\u0026Expires=1621423586"},"applications":{"hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","serialNumber":"TAIJB1072828","description":"","disk":[{"percent":82.16348920732285,"total":449529856,"used":346475520,"device":"/dev/root","free":75214848,"mount":"/","name":"System"}],"firmwareVersion":"1.4","memorySize":524054528,"modelName":"UC-3111-T-EU-LX","osType":"Linux","thingsProProduct":"","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)"},"edition":"2.0","osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=0TvujGg0EPIWwkPqCHFsIMksQYg%3D\u0026Expires=1621423586"},"general":{"firmwareVersion":"1.4","lastBootTime":"2021-05-17T03:13:07Z","osType":"Linux","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)","hostName":"Moxa","modelName":"UC-3111-T-EU-LX","serialNumber":"TAIJB1072828","thingsProProduct":""},"ethernets":{"desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/a249a099-554d-4e33-badf-8ad1c8372d04/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=k8Fel7tJVqHClIWr1osl%2B%2F6skjA%3D\u0026Expires=1621423586","status":"success","version":""},"sshserver":{"enable":true,"port":22},"osPatchControl":{"status":"ready"},"application-dlm-dashboard":{"version":"","desiredURL":"","reportedURL":"","status":""}}
May 19 09:26:32 Moxa dlm-agent[17597]: [INFO]May 19 09:26:32.103 - [origin:dlm-1] - [reported] classic
May 19 09:26:32 Moxa dlm-agent[17597]: [INFO]May 19 09:26:32.261 - [origin:dlm-1] - [reported callback], success
May 19 09:26:32 Moxa dlm-agent[17597]: [INFO]May 19 09:26:32.263 - [origin:dlm-1] - [reported] applications
May 19 09:26:32 Moxa dlm-agent[17597]: [INFO]May 19 09:26:32.420 - [origin:dlm-1] - [reported callback], success
May 19 09:26:58 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:58.292 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:26:58 Moxa dlm-agent[17597]: [DBUG]May 19 09:26:58.296 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
May 19 09:27:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:27:28.292 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:27:28 Moxa dlm-agent[17597]: [DBUG]May 19 09:27:28.292 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
```

