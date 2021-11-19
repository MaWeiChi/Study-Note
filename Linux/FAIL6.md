```
May 19 10:01:01 Moxa systemd[1]: Started dlm-agent service.
May 19 10:01:02 Moxa dlm-agent[2272]: *******************************************************************************
May 19 10:01:02 Moxa dlm-agent[2272]:   _______ _     _                 _____             _____  _                 _
May 19 10:01:02 Moxa dlm-agent[2272]:  |__   __| |   (_)               |  __ \           |  ___|| |               | |
May 19 10:01:02 Moxa dlm-agent[2272]:     | |  | |__  _ _ __   __ _ ___| |__) | __ ___   | |    | | ___  _   _  __| |
May 19 10:01:02 Moxa dlm-agent[2272]:     | |  | '_ \| | '_ \ / _  / __|  ___/ '__/ _ \  | |    | |/ _ \| | | |/ _  |
May 19 10:01:02 Moxa dlm-agent[2272]:     | |  | | | | | | | | (_| \__ \ |   | | | (_) | | |___ | | (_) | |_| | (_| |
May 19 10:01:02 Moxa dlm-agent[2272]:     |_|  |_| |_|_|_| |_|\__, |___/_|   |_|  \___/  |_____||_|\___/\_____/\__,_|
May 19 10:01:02 Moxa dlm-agent[2272]:                          __/ |
May 19 10:01:02 Moxa dlm-agent[2272]:                         |___/
May 19 10:01:02 Moxa dlm-agent[2272]:   Application : ThingsPro Cloud service
May 19 10:01:02 Moxa dlm-agent[2272]:   Version     : 0.1.2-armhf, Time : 2021-05-19T10:01:02Z
May 19 10:01:02 Moxa dlm-agent[2272]: *******************************************************************************
May 19 10:01:02 Moxa dlm-agent[2272]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 10:01:02 Moxa dlm-agent[2272]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 10:01:02 Moxa dlm-agent[2272]:  - using env:        export GIN_MODE=release
May 19 10:01:02 Moxa dlm-agent[2272]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 10:01:02 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/_/info            --> main.getInfo (4 handlers)
May 19 10:01:02 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/devicedb/onchange --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/agent.TransOnChange (4 handlers)
May 19 10:01:02 Moxa dlm-agent[2272]: [openvpn] 2021/05/19 10:01:02 config.go:52: main config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json: no such file or directory,<nil>, try load backup
May 19 10:01:02 Moxa dlm-agent[2272]: [openvpn] 2021/05/19 10:01:02 config.go:62: backup config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json.backup: no such file or directory,<nil>, try load factory
May 19 10:01:02 Moxa dlm-agent[2272]: MOXA api token open failed, open /run/mx-api-token: no such file or directory
May 19 10:01:02 Moxa dlm-agent[2272]: config load err:open /etc/dlm-agent/.data/device/openvpn/openvpn.json: no such file or directory
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/device/general    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/device/applications --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/device/ethernets  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernets-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/device/ethernets/:id --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernetById-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/system/reboot     --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PutSystemReboot-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/openvpn           --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetOpenVpn-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PATCH  /api/v1/openvpn/dlm       --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PatchOpenVpnDlm-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetSSHSettings-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).SetSSHSettings-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: config load err:open /etc/dlm-agent/.data/device/software/todo_UpgradeList.json: no such file or directory
May 19 10:01:03 Moxa dlm-agent[2272]: get todo upgrade list error: unexpected end of JSON input
May 19 10:01:03 Moxa dlm-agent[2272]: config load err:open /etc/dlm-agent/.data/device/software/config.json: no such file or directory
May 19 10:01:03 Moxa dlm-agent[2272]: 0 0 * * * *
May 19 10:01:03 Moxa dlm-agent[2272]: toolbox.NewTask
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftware-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftware-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/software/ospatchlist --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetPatchList-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/software/detail   --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareDetail-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareSchedule-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftwareSchedule-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetLogs (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PostLog (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetProfile (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PatchProfile (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/tags/monitor      --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetMonitor-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/tags/monitor/:args --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).PostMonitor-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/tags/monitor/system/system --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetTags-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 10:01:03 Moxa dlm-agent[2272]:  - using env:        export GIN_MODE=release
May 19 10:01:03 Moxa dlm-agent[2272]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.084 - [origin:remoted] - plugin message uploader created
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.084 - [origin:remoted] - plugin devicedb created
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.084 - [origin:remoted] - plugin appman created
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.086 - [origin:remoted] - plugin apiV1 created
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.086 - [origin:remoted] - plugin store created
May 19 10:01:03 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:03.088 - [origin:dlm] - cert path:
May 19 10:01:03 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:03.088 - [origin:dlm] - key path:
May 19 10:01:03 Moxa dlm-agent[2272]: [ERRO]May 19 10:01:03.169 - [origin:dlm] - tasks[1/2] name:[armhf] check cert files, result:failed, error:exit status 1
May 19 10:01:03 Moxa dlm-agent[2272]: [WARN]May 19 10:01:03.169 - [origin:dlm] - load factory certificates failed, error: exit status 1
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.170 - [origin:remoted] - service dlm created
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).getMessagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).putMessagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).postMessagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] DELETE /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.175 - [origin:remoted] - plugin message uploader started
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.176 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.177 - [origin:remoted] - worker sock binding
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.177 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 10:01:03 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:03.190 - [origin:devicedb:dlm-1] - database subscriber start
May 19 10:01:03 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:03.190 - [origin:devicedb:dlm-1] - database subscriber exit, channel empty
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.190 - [origin:devicedb] - connectionID:dlm-1 db client started
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSource-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSource-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSelections-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSelections-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/onchange --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesOnchange-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/test --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).testPropertiesSelections-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).reported-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/custom/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).userReported-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/clear --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).clear-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.191 - [origin:remoted] - plugin devicedb started
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.191 - [origin:remoted] - plugin appman started
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.191 - [origin:remoted] - plugin apiV1 started
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).getConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).putConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores/restart --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).reload-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.197 - [origin:remoted] - plugin store started
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] HEAD   /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] HEAD   /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] HEAD   /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] HEAD   /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/info    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).info-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/tags    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).tagMonitor-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/remotectl/services --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).servicesInfo-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/message --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).sendMessage-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/state --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).updateState-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.198 - [origin:dlm] - skip start, false
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).postConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PATCH  /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).patchConfiguration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/_/device      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDeviceID-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/monitor       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getMonitor-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/registration  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putRegistration-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/enrollment    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putEnrollment-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/edition       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getEdition-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/control       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).controlThing-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/login         --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).login-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] POST   /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] DELETE /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDashboard-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putDashboard-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.200 - [origin:remoted] - service dlm started
May 19 10:01:03 Moxa dlm-agent[2272]: [INFO]May 19 10:01:03.200 - [origin:remoted] - all started
May 19 10:01:25 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:25 | 200 |    1.614814ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:25 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:25 | 200 |     223.636µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:25 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:25 | 200 |     224.196µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:26 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:26 | 200 |     312.235µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:26 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:26 | 200 |     223.716µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:26 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:26 | 200 |     218.356µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:27 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:27 | 200 |     315.035µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:27 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:27 | 200 |     247.996µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:36 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:36 | 200 |     218.837µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 10:01:40 Moxa dlm-agent[2272]: [INFO]May 19 10:01:40.684 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 10:01:41 Moxa dlm-agent[2272]: [INFO]May 19 10:01:41.347 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 10:01:41 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:41 | 200 |  1.439180338s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"
May 19 10:01:47 Moxa dlm-agent[2272]: [INFO]May 19 10:01:47.661 - [origin:dlm] - skip stop because connection status: false
May 19 10:01:48 Moxa dlm-agent[2272]: [INFO]May 19 10:01:48.557 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 10:01:49 Moxa dlm-agent[2272]: [INFO]May 19 10:01:49.179 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 10:01:49 Moxa dlm-agent[2272]: [INFO]May 19 10:01:49.181 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 10:01:49 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:49 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:49 | 200 |   27.630189ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 10:01:49 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:49 | 200 |  113.082436ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 10:01:49 Moxa dlm-agent[2272]: [INFO]May 19 10:01:49.346 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 10:01:49 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:49 | 200 |     247.036µs |       127.0.0.1 | GET      "/api/v1/_/info"
May 19 10:01:51 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:51.933 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 10:01:51 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:51.936 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 10:01:51 Moxa dlm-agent[2272]: [INFO]May 19 10:01:51.976 - [origin:dlm] - latest configuration:  {
May 19 10:01:51 Moxa dlm-agent[2272]:   "connection": {
May 19 10:01:51 Moxa dlm-agent[2272]:     "enable": false,
May 19 10:01:51 Moxa dlm-agent[2272]:     "retryDelaySec": 10,
May 19 10:01:51 Moxa dlm-agent[2272]:     "picTarget": "stage"
May 19 10:01:51 Moxa dlm-agent[2272]:   },
May 19 10:01:51 Moxa dlm-agent[2272]:   "certificates": {
May 19 10:01:51 Moxa dlm-agent[2272]:     "caCertFile": "dev.crt",
May 19 10:01:51 Moxa dlm-agent[2272]:     "caPkFile": "dev.key",
May 19 10:01:51 Moxa dlm-agent[2272]:     "certificateInfo": {
May 19 10:01:51 Moxa dlm-agent[2272]:       "organization": "Moxa Inc.",
May 19 10:01:51 Moxa dlm-agent[2272]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 10:01:51 Moxa dlm-agent[2272]:       "notBefore": "May 19, 2021 10:01:50",
May 19 10:01:51 Moxa dlm-agent[2272]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 10:01:51 Moxa dlm-agent[2272]:       "modelName": "UC-3111-T-EU-LX",
May 19 10:01:51 Moxa dlm-agent[2272]:       "macAddress": "0090E882F2FB",
May 19 10:01:51 Moxa dlm-agent[2272]:       "serialNumber": "TAIJB1072828"
May 19 10:01:51 Moxa dlm-agent[2272]:     }
May 19 10:01:51 Moxa dlm-agent[2272]:   },
May 19 10:01:51 Moxa dlm-agent[2272]:   "connectionStatus": {
May 19 10:01:51 Moxa dlm-agent[2272]:     "projectName": "Kevin Test - AGT",
May 19 10:01:51 Moxa dlm-agent[2272]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 10:01:51 Moxa dlm-agent[2272]:     "status": "",
May 19 10:01:51 Moxa dlm-agent[2272]:     "message": "",
May 19 10:01:51 Moxa dlm-agent[2272]:     "lastConnectedTime": "",
May 19 10:01:51 Moxa dlm-agent[2272]:     "lastConnectTime": "",
May 19 10:01:51 Moxa dlm-agent[2272]:     "lastDisconnectTime": "",
May 19 10:01:51 Moxa dlm-agent[2272]:     "statuslastUpdateTime": ""
May 19 10:01:51 Moxa dlm-agent[2272]:   }
May 19 10:01:51 Moxa dlm-agent[2272]: }
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:51.983 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:51.983 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:51.984 - [origin:dlm] - latest configuration:  {
May 19 10:01:52 Moxa dlm-agent[2272]:   "connection": {
May 19 10:01:52 Moxa dlm-agent[2272]:     "enable": true,
May 19 10:01:52 Moxa dlm-agent[2272]:     "retryDelaySec": 10,
May 19 10:01:52 Moxa dlm-agent[2272]:     "picTarget": "stage"
May 19 10:01:52 Moxa dlm-agent[2272]:   },
May 19 10:01:52 Moxa dlm-agent[2272]:   "certificates": {
May 19 10:01:52 Moxa dlm-agent[2272]:     "caCertFile": "dev.crt",
May 19 10:01:52 Moxa dlm-agent[2272]:     "caPkFile": "dev.key",
May 19 10:01:52 Moxa dlm-agent[2272]:     "certificateInfo": {
May 19 10:01:52 Moxa dlm-agent[2272]:       "organization": "Moxa Inc.",
May 19 10:01:52 Moxa dlm-agent[2272]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 10:01:52 Moxa dlm-agent[2272]:       "notBefore": "May 19, 2021 10:01:50",
May 19 10:01:52 Moxa dlm-agent[2272]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 10:01:52 Moxa dlm-agent[2272]:       "modelName": "UC-3111-T-EU-LX",
May 19 10:01:52 Moxa dlm-agent[2272]:       "macAddress": "0090E882F2FB",
May 19 10:01:52 Moxa dlm-agent[2272]:       "serialNumber": "TAIJB1072828"
May 19 10:01:52 Moxa dlm-agent[2272]:     }
May 19 10:01:52 Moxa dlm-agent[2272]:   },
May 19 10:01:52 Moxa dlm-agent[2272]:   "connectionStatus": {
May 19 10:01:52 Moxa dlm-agent[2272]:     "projectName": "Kevin Test - AGT",
May 19 10:01:52 Moxa dlm-agent[2272]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 10:01:52 Moxa dlm-agent[2272]:     "status": "",
May 19 10:01:52 Moxa dlm-agent[2272]:     "message": "",
May 19 10:01:52 Moxa dlm-agent[2272]:     "lastConnectedTime": "",
May 19 10:01:52 Moxa dlm-agent[2272]:     "lastConnectTime": "",
May 19 10:01:52 Moxa dlm-agent[2272]:     "lastDisconnectTime": "",
May 19 10:01:52 Moxa dlm-agent[2272]:     "statuslastUpdateTime": ""
May 19 10:01:52 Moxa dlm-agent[2272]:   }
May 19 10:01:52 Moxa dlm-agent[2272]: }
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:51.986 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:52 | 200 |    1.154421ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.012 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.012 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.012 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [ERRO]May 19 10:01:52.014 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 10:01:52 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:52 | 404 |        9.32µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 10:01:52 Moxa dlm-agent[2272]: [ERRO]May 19 10:01:52.016 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.016 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.017 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [ERRO]May 19 10:01:52.017 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.017 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.017 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.017 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:52 | 200 |  4.356336783s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.028 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.028 - [origin:dlm] - register target: stage
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.028 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.029 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success
May 19 10:01:52 Moxa dlm-agent[2272]: [INFO]May 19 10:01:52.980 - [origin:dlm] - tasks[3/6] name:check certificate, result:success
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.890 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.893 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.895 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.904 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 10:01:53 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:53.906 - [origin:dlm-1] - try to connect rule server ...
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.914 - [origin:dlm-1] - connect to rule server success
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.922 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.922 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.922 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.923 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 10:01:53 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:53.919 - [origin:dlm-1] - monitor stream in
May 19 10:01:53 Moxa dlm-agent[2272]: [WARN]May 19 10:01:53.924 - [origin:dlm-1] - waiting for self registration...
May 19 10:01:53 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:53.919 - [origin:dlm-1] - twin stream in
May 19 10:01:53 Moxa dlm-agent[2272]: [WARN]May 19 10:01:53.924 - [origin:dlm-1] - waiting for self registration...
May 19 10:01:53 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:53.919 - [origin:dlm-1] - message stream in
May 19 10:01:53 Moxa dlm-agent[2272]: [INFO]May 19 10:01:53.940 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.182 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.182 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.193 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.194 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.196 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.207 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.209 - [origin:message-group:dlm-1:2] - events state: running
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.209 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.210 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.221 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.211 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.211 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.213 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.232 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.232 - [origin:message-group:dlm-1:4] - events state: running
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.233 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.233 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.233 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.198 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.238 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.239 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.239 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.198 - [origin:message-group:dlm-1:2] - twin stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.239 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.198 - [origin:message-group:dlm-1:2] - message stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.239 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.217 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.240 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.219 - [origin:message-group:dlm-1:4] - twin stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.240 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 10:01:54 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:54.220 - [origin:message-group:dlm-1:4] - message stream in
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.240 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.924 - [origin:dlm-1] - twin stream registration success
May 19 10:01:54 Moxa dlm-agent[2272]: [INFO]May 19 10:01:54.924 - [origin:dlm-1] - monitor stream registration success
May 19 10:01:55 Moxa dlm-agent[2272]: [INFO]May 19 10:01:55.107 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 10:01:55 Moxa dlm-agent[2272]: [INFO]May 19 10:01:55.107 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 10:01:55 Moxa dlm-agent[2272]: [INFO]May 19 10:01:55.108 - [origin:dlm-1] - [connection callback], connection = true
May 19 10:01:55 Moxa dlm-agent[2272]: [INFO]May 19 10:01:55.111 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 10:01:55 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:55 | 200 |     222.117µs |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 10:01:55 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:55 | 200 |     745.108µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 10:01:55 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:55 | 200 |    6.701291ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 10:01:55 Moxa dlm-agent[2272]: [INFO]May 19 10:01:55.150 - [origin:message-group:dlm-1:2] - stopping
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.209 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.210 - [origin:message-group:dlm-1:2] - message-routine:exit, reason: source canceled
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.210 - [origin:message-group:dlm-1:2] - tasks[1/5] name:stop ipc connection, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - outputservice exit
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - outputservice state: stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - timer stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - formater stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - tasks[2/5] name:stop output control service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.211 - [origin:message-group:dlm-1:2] - tasks[3/5] name:stop tag service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.222 - [origin:message-group:dlm-1:2] - event origin:OS patch listener exit
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.223 - [origin:message-group:dlm-1:2] - tasks[4/5] name:stop event service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.223 - [origin:message-group:dlm-1:2] - tasks[5/5] name:stop remote monitor service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.223 - [origin:message-group:dlm-1:2] - stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |  2.984112073s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process"
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.227 - [origin:message-group:dlm-1:2] - message:exit, reason: source canceled
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.227 - [origin:message-group:dlm-1:2] - message stream out
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |  2.080353366s |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.230 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.230 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.231 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   10.933262ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.258 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.258 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.258 - [origin:message-group:dlm-1:2] - events state: running
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.258 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.258 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.259 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.260 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.260 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.275 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.276 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.278 - [origin:message-group:dlm-1:4] - stopping
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.279 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.279 - [origin:message-group:dlm-1:4] - message-routine:exit, reason: source canceled
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.279 - [origin:message-group:dlm-1:4] - tasks[1/5] name:stop ipc connection, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.279 - [origin:message-group:dlm-1:4] - outputservice exit
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.280 - [origin:message-group:dlm-1:4] - outputservice state: stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.280 - [origin:message-group:dlm-1:2] - twin stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.280 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.280 - [origin:message-group:dlm-1:2] - message stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.281 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.281 - [origin:message-group:dlm-1:4] - timer stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.285 - [origin:message-group:dlm-1:4] - formater stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.285 - [origin:message-group:dlm-1:4] - tasks[2/5] name:stop output control service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.285 - [origin:message-group:dlm-1:4] - tasks[3/5] name:stop tag service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.286 - [origin:message-group:dlm-1:4] - event origin:OS patch listener exit
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.286 - [origin:message-group:dlm-1:4] - tasks[4/5] name:stop event service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.286 - [origin:message-group:dlm-1:4] - tasks[5/5] name:stop remote monitor service, result:success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.286 - [origin:message-group:dlm-1:4] - stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   19.154288ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.297 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.297 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.297 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |  3.029896567s |       127.0.0.1 | GET      "/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available"
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.306 - [origin:message-group:dlm-1:4] - message:exit, reason: source canceled
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.306 - [origin:message-group:dlm-1:4] - message stream out
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.319 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.319 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.320 - [origin:message-group:dlm-1:4] - events state: running
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.320 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.320 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.320 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.332 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.332 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.354 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.354 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.359 - [origin:message-group:dlm-1:4] - twin stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.359 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.359 - [origin:message-group:dlm-1:4] - message stream in
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.360 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 10:01:57 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   24.600999ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   52.842459ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.389 - [origin:store] - [connectionID:dlm-1] - start
May 19 10:01:57 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   37.806064ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 10:01:57 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   44.409476ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   78.617358ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.461 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 10:01:57 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:57.480 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.481 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.482 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: waiting
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.482 - [origin:store] - [connectionID:dlm-1] - get connection success, start register
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.487 - [origin:dlm-1] - start store message
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.489 - [origin:store] - [connectionID:dlm-1] - register receiver success
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.489 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 10:01:57 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:57 | 200 |   35.926894ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 10:01:57 Moxa dlm-agent[2272]: [INFO]May 19 10:01:57.513 - [origin:dlm] - last ota result empty
May 19 10:01:58 Moxa dlm-agent[2272]: [INFO]May 19 10:01:58.107 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 10:01:58 Moxa dlm-agent[2272]: [INFO]May 19 10:01:58.690 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 10:01:58 Moxa dlm-agent[2272]: [INFO]May 19 10:01:58.690 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 10:01:58 Moxa dlm-agent[2272]: [INFO]May 19 10:01:58.849 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 10:01:58 Moxa dlm-agent[2272]: [INFO]May 19 10:01:58.950 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=kXzuCaEUbv2FCA9I2aCnZz9PIVQ%3D&Expires=1621425715"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=2gDDdjtE1pFAu%2BOP9GjYPmi7ZQo%3D&Expires=1621425715"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=aJMcufCmD%2BoEoqlO4wo%2FekGQgNU%3D&Expires=1621425715"}}
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.951 - [origin:dlm-1] - [twin callback] exit
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.955 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.957 - [origin:dlm-1] - desired push type:$completeDesired, queue: 0/50
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.957 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=kXzuCaEUbv2FCA9I2aCnZz9PIVQ%3D&Expires=1621425715"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=2gDDdjtE1pFAu%2BOP9GjYPmi7ZQo%3D&Expires=1621425715"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=aJMcufCmD%2BoEoqlO4wo%2FekGQgNU%3D&Expires=1621425715"}}
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.965 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.967 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.967 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 10:01:58 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:58 | 200 |     352.354µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 10:01:58 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:58.982 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 10:01:58 Moxa dlm-agent[2272]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 10:01:58 Moxa dlm-agent[2272]: toolbox.NewTask
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:58 | 200 |     379.674µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |    5.045718ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   15.414109ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 10:01:59 Moxa dlm-agent[2272]: [INFO]May 19 10:01:59.003 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 10:01:59 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   26.686405ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |    29.41884ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 10:01:59 Moxa dlm-agent[2272]: GetDeviceGeneral
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   16.045138ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   59.757986ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |     527.472µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |      603.51µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |    1.101862ms |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |     212.077µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |     586.711µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 10:01:59 Moxa dlm-agent[2272]: [INFO]May 19 10:01:59.188 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   12.787392ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |     176.597µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |     179.437µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 10:01:59 Moxa dlm-agent[2272]: [DBUG]May 19 10:01:59.901 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"detail":{"origin":"","depends":null,"description":"","label":""},"name":"","newVersion":"","size":0,"version":""}}
May 19 10:01:59 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:01:59 | 200 |   26.614006ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 10:01:59 Moxa dlm-agent[2272]: [INFO]May 19 10:01:59.955 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 10:02:00 Moxa dlm-agent[2272]: [DBUG]May 19 10:02:00.288 - [origin:dlm-1] - [command callback] $aws/things/ce6272b2-14fb-462b-8f32-50532663b682/jobs/$next/get/accepted {"timestamp":1621418520}
May 19 10:02:00 Moxa dlm-agent[2272]: [INFO]May 19 10:02:00.288 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 10:02:00 Moxa dlm-agent[2272]: [DBUG]May 19 10:02:00.657 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"dns":["10.123.200.11","10.123.200.12"],"gateway":"10.123.12.1","id":1,"mac":"00:90:e8:82:f2:fb","status":"connected","subnet":"10.123.12.0","broadcast":"10.123.13.255","displayName":"eth0","enable":true,"ip":"10.123.13.23","name":"eth0","netmask":"255.255.254.0","speed":"100 Mbps"},{"status":"disconnected","subnet":"192.168.4.0","displayName":"eth1","dns":[],"gateway":"","ip":"192.168.4.127","name":"eth1","speed":"--","broadcast":"192.168.4.255","enable":true,"id":2,"mac":"00:90:e8:82:f2:fc","netmask":"255.255.255.0"}]}
May 19 10:02:00 Moxa dlm-agent[2272]: [GIN] 2021/05/19 - 10:02:00 | 200 |     288.675µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 10:02:01 Moxa dlm-agent[2272]: [DBUG]May 19 10:02:01.405 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 10:02:01 Moxa dlm-agent[2272]: [DBUG]May 19 10:02:01.410 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 10:02:01 Moxa dlm-agent[2272]: [INFO]May 19 10:02:01.414 - [origin:dlm-1] - [reported], {"edition":"2.0","application-dlm-dashboard":{"status":"","version":"","desiredURL":"","reportedURL":""},"sshserver":{"enable":true,"port":22},"osPatchControl":{"status":"ready"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=2gDDdjtE1pFAu%2BOP9GjYPmi7ZQo%3D\u0026Expires=1621425715"},"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=aJMcufCmD%2BoEoqlO4wo%2FekGQgNU%3D\u0026Expires=1621425715"},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/ce6272b2-14fb-462b-8f32-50532663b682/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=kXzuCaEUbv2FCA9I2aCnZz9PIVQ%3D\u0026Expires=1621425715"},"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"general":{"firmwareVersion":"1.4","hostName":"Moxa","modelName":"UC-3111-T-EU-LX","serialNumber":"TAIJB1072828","thingsProProduct":"","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)","lastBootTime":"2021-05-17T03:13:07Z","osType":"Linux"},"applications":{"cpu":"ARMv7 Processor rev 2 (v7l)","disk":[{"device":"/dev/root","free":75214848,"mount":"/","name":"System","percent":82.16348920732285,"total":449529856,"used":346475520}],"firmwareVersion":"1.4","modelName":"UC-3111-T-EU-LX","thingsProProduct":"","osType":"Linux","serialNumber":"TAIJB1072828","biosVersion":"1.4.0S02","description":"","hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","memorySize":524054528}}
May 19 10:02:01 Moxa dlm-agent[2272]: [INFO]May 19 10:02:01.421 - [origin:dlm-1] - [reported] classic
May 19 10:02:01 Moxa dlm-agent[2272]: [INFO]May 19 10:02:01.587 - [origin:dlm-1] - [reported callback], success
May 19 10:02:01 Moxa dlm-agent[2272]: [INFO]May 19 10:02:01.588 - [origin:dlm-1] - [reported] applications
May 19 10:02:01 Moxa dlm-agent[2272]: [INFO]May 19 10:02:01.748 - [origin:dlm-1] - [reported callback], success
```

