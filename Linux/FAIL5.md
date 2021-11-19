```
May 19 09:56:59 Moxa dlm-agent[32634]: *******************************************************************************
May 19 09:56:59 Moxa dlm-agent[32634]:   _______ _     _                 _____             _____  _                 _
May 19 09:56:59 Moxa dlm-agent[32634]:  |__   __| |   (_)               |  __ \           |  ___|| |               | |
May 19 09:56:59 Moxa dlm-agent[32634]:     | |  | |__  _ _ __   __ _ ___| |__) | __ ___   | |    | | ___  _   _  __| |
May 19 09:56:59 Moxa dlm-agent[32634]:     | |  | '_ \| | '_ \ / _  / __|  ___/ '__/ _ \  | |    | |/ _ \| | | |/ _  |
May 19 09:56:59 Moxa dlm-agent[32634]:     | |  | | | | | | | | (_| \__ \ |   | | | (_) | | |___ | | (_) | |_| | (_| |
May 19 09:56:59 Moxa dlm-agent[32634]:     |_|  |_| |_|_|_| |_|\__, |___/_|   |_|  \___/  |_____||_|\___/\_____/\__,_|
May 19 09:56:59 Moxa dlm-agent[32634]:                          __/ |
May 19 09:56:59 Moxa dlm-agent[32634]:                         |___/
May 19 09:56:59 Moxa dlm-agent[32634]:   Application : ThingsPro Cloud service
May 19 09:56:59 Moxa dlm-agent[32634]:   Version     : 0.1.2-armhf, Time : 2021-05-19T09:56:59Z
May 19 09:56:59 Moxa dlm-agent[32634]: *******************************************************************************
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:56:59 Moxa dlm-agent[32634]:  - using env:        export GIN_MODE=release
May 19 09:56:59 Moxa dlm-agent[32634]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/_/info            --> main.getInfo (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/devicedb/onchange --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/agent.TransOnChange (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [openvpn] 2021/05/19 09:56:59 config.go:52: main config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json: no such file or directory,<nil>, try load backup
May 19 09:56:59 Moxa dlm-agent[32634]: [openvpn] 2021/05/19 09:56:59 config.go:62: backup config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json.backup: no such file or directory,<nil>, try load factory
May 19 09:56:59 Moxa dlm-agent[32634]: MOXA api token open failed, open /run/mx-api-token: no such file or directory
May 19 09:56:59 Moxa dlm-agent[32634]: config load err:open /etc/dlm-agent/.data/device/openvpn/openvpn.json: no such file or directory
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/device/general    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/device/applications --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/device/ethernets  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernets-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/device/ethernets/:id --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernetById-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/system/reboot     --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PutSystemReboot-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/openvpn           --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetOpenVpn-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PATCH  /api/v1/openvpn/dlm       --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PatchOpenVpnDlm-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetSSHSettings-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).SetSSHSettings-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: config load err:open /etc/dlm-agent/.data/device/software/todo_UpgradeList.json: no such file or directory
May 19 09:56:59 Moxa dlm-agent[32634]: get todo upgrade list error: unexpected end of JSON input
May 19 09:56:59 Moxa dlm-agent[32634]: config load err:open /etc/dlm-agent/.data/device/software/config.json: no such file or directory
May 19 09:56:59 Moxa dlm-agent[32634]: 0 0 * * * *
May 19 09:56:59 Moxa dlm-agent[32634]: toolbox.NewTask
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftware-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftware-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/software/ospatchlist --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetPatchList-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/software/detail   --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareDetail-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareSchedule-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftwareSchedule-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetLogs (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PostLog (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetProfile (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PatchProfile (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/tags/monitor      --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetMonitor-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/tags/monitor/:args --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).PostMonitor-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/tags/monitor/system/system --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetTags-fm (4 handlers)
May 19 09:56:59 Moxa dlm-agent[32634]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:57:00 Moxa dlm-agent[32634]:  - using env:        export GIN_MODE=release
May 19 09:57:00 Moxa dlm-agent[32634]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:56:59.954 - [origin:remoted] - plugin message uploader created
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:56:59.954 - [origin:remoted] - plugin devicedb created
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:56:59.954 - [origin:remoted] - plugin appman created
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:56:59.956 - [origin:remoted] - plugin apiV1 created
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:56:59.956 - [origin:remoted] - plugin store created
May 19 09:57:00 Moxa dlm-agent[32634]: [DBUG]May 19 09:56:59.958 - [origin:dlm] - cert path:
May 19 09:57:00 Moxa dlm-agent[32634]: [DBUG]May 19 09:56:59.958 - [origin:dlm] - key path:
May 19 09:57:00 Moxa dlm-agent[32634]: [ERRO]May 19 09:57:00.039 - [origin:dlm] - tasks[1/2] name:[armhf] check cert files, result:failed, error:exit status 1
May 19 09:57:00 Moxa dlm-agent[32634]: [WARN]May 19 09:57:00.039 - [origin:dlm] - load factory certificates failed, error: exit status 1
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.040 - [origin:remoted] - service dlm created
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores/restart --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).reload-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.041 - [origin:remoted] - plugin store started
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).getMessagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).putMessagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.047 - [origin:remoted] - worker sock binding
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.047 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.046 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).postMessagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] DELETE /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.064 - [origin:remoted] - plugin message uploader started
May 19 09:57:00 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:00.069 - [origin:devicedb:dlm-1] - database subscriber start
May 19 09:57:00 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:00.069 - [origin:devicedb:dlm-1] - database subscriber exit, channel empty
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.069 - [origin:devicedb] - connectionID:dlm-1 db client started
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSource-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSource-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSelections-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSelections-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/onchange --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesOnchange-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/test --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).testPropertiesSelections-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).reported-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/custom/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).userReported-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/clear --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).clear-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.070 - [origin:remoted] - plugin devicedb started
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.070 - [origin:remoted] - plugin appman started
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.070 - [origin:remoted] - plugin apiV1 started
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] HEAD   /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] HEAD   /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] HEAD   /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] HEAD   /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/info    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).info-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/tags    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).tagMonitor-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/remotectl/services --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).servicesInfo-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/message --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).sendMessage-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/state --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).updateState-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.071 - [origin:dlm] - skip start, false
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).postConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PATCH  /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).patchConfiguration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/_/device      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDeviceID-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/monitor       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getMonitor-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/registration  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putRegistration-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/enrollment    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putEnrollment-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/edition       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getEdition-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/control       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).controlThing-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/login         --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).login-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] POST   /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] DELETE /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDashboard-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putDashboard-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.073 - [origin:remoted] - service dlm started
May 19 09:57:00 Moxa dlm-agent[32634]: [INFO]May 19 09:57:00.073 - [origin:remoted] - all started
May 19 09:57:15 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:15 | 200 |     1.23786ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:15 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:15 | 200 |     220.396µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:15 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:15 | 200 |     238.516µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:15 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:15 | 200 |     320.035µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:15 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:15 | 200 |     220.036µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:16 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:16 | 200 |     220.036µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:22 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:22 | 200 |     240.516µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:57:26 Moxa dlm-agent[32634]: [INFO]May 19 09:57:26.039 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:57:26 Moxa dlm-agent[32634]: [INFO]May 19 09:57:26.665 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:57:26 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:26 | 200 |  1.459015523s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"
May 19 09:57:31 Moxa dlm-agent[32634]: [INFO]May 19 09:57:31.224 - [origin:dlm] - skip stop because connection status: false
May 19 09:57:32 Moxa dlm-agent[32634]: [INFO]May 19 09:57:32.039 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:57:32 Moxa dlm-agent[32634]: [INFO]May 19 09:57:32.696 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:57:32 Moxa dlm-agent[32634]: [INFO]May 19 09:57:32.699 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 09:57:32 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:32 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:32 | 200 |   23.885171ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:57:32 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:32 | 200 |  129.474449ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 09:57:32 Moxa dlm-agent[32634]: [INFO]May 19 09:57:32.919 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 09:57:32 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:32 | 200 |     227.156µs |       127.0.0.1 | GET      "/api/v1/_/info"
May 19 09:57:35 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:35.494 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 09:57:35 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:35.497 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.537 - [origin:dlm] - latest configuration:  {
May 19 09:57:35 Moxa dlm-agent[32634]:   "connection": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "enable": false,
May 19 09:57:35 Moxa dlm-agent[32634]:     "retryDelaySec": 10,
May 19 09:57:35 Moxa dlm-agent[32634]:     "picTarget": "stage"
May 19 09:57:35 Moxa dlm-agent[32634]:   },
May 19 09:57:35 Moxa dlm-agent[32634]:   "certificates": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "caCertFile": "dev.crt",
May 19 09:57:35 Moxa dlm-agent[32634]:     "caPkFile": "dev.key",
May 19 09:57:35 Moxa dlm-agent[32634]:     "certificateInfo": {
May 19 09:57:35 Moxa dlm-agent[32634]:       "organization": "Moxa Inc.",
May 19 09:57:35 Moxa dlm-agent[32634]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:57:35 Moxa dlm-agent[32634]:       "notBefore": "May 19, 2021 09:57:33",
May 19 09:57:35 Moxa dlm-agent[32634]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:57:35 Moxa dlm-agent[32634]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:57:35 Moxa dlm-agent[32634]:       "macAddress": "0090E882F2FB",
May 19 09:57:35 Moxa dlm-agent[32634]:       "serialNumber": "TAIJB1072828"
May 19 09:57:35 Moxa dlm-agent[32634]:     }
May 19 09:57:35 Moxa dlm-agent[32634]:   },
May 19 09:57:35 Moxa dlm-agent[32634]:   "connectionStatus": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "projectName": "Kevin Test - AGT",
May 19 09:57:35 Moxa dlm-agent[32634]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:57:35 Moxa dlm-agent[32634]:     "status": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "message": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastConnectedTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastConnectTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastDisconnectTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "statuslastUpdateTime": ""
May 19 09:57:35 Moxa dlm-agent[32634]:   }
May 19 09:57:35 Moxa dlm-agent[32634]: }
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.545 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.546 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.546 - [origin:dlm] - latest configuration:  {
May 19 09:57:35 Moxa dlm-agent[32634]:   "connection": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "enable": true,
May 19 09:57:35 Moxa dlm-agent[32634]:     "retryDelaySec": 10,
May 19 09:57:35 Moxa dlm-agent[32634]:     "picTarget": "stage"
May 19 09:57:35 Moxa dlm-agent[32634]:   },
May 19 09:57:35 Moxa dlm-agent[32634]:   "certificates": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "caCertFile": "dev.crt",
May 19 09:57:35 Moxa dlm-agent[32634]:     "caPkFile": "dev.key",
May 19 09:57:35 Moxa dlm-agent[32634]:     "certificateInfo": {
May 19 09:57:35 Moxa dlm-agent[32634]:       "organization": "Moxa Inc.",
May 19 09:57:35 Moxa dlm-agent[32634]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:57:35 Moxa dlm-agent[32634]:       "notBefore": "May 19, 2021 09:57:33",
May 19 09:57:35 Moxa dlm-agent[32634]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:57:35 Moxa dlm-agent[32634]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:57:35 Moxa dlm-agent[32634]:       "macAddress": "0090E882F2FB",
May 19 09:57:35 Moxa dlm-agent[32634]:       "serialNumber": "TAIJB1072828"
May 19 09:57:35 Moxa dlm-agent[32634]:     }
May 19 09:57:35 Moxa dlm-agent[32634]:   },
May 19 09:57:35 Moxa dlm-agent[32634]:   "connectionStatus": {
May 19 09:57:35 Moxa dlm-agent[32634]:     "projectName": "Kevin Test - AGT",
May 19 09:57:35 Moxa dlm-agent[32634]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:57:35 Moxa dlm-agent[32634]:     "status": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "message": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastConnectedTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastConnectTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "lastDisconnectTime": "",
May 19 09:57:35 Moxa dlm-agent[32634]:     "statuslastUpdateTime": ""
May 19 09:57:35 Moxa dlm-agent[32634]:   }
May 19 09:57:35 Moxa dlm-agent[32634]: }
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.548 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:35 | 200 |    1.169461ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.571 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.572 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.572 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [ERRO]May 19 09:57:35.574 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 09:57:35 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:35 | 404 |        9.08µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 09:57:35 Moxa dlm-agent[32634]: [ERRO]May 19 09:57:35.576 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.576 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.576 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [ERRO]May 19 09:57:35.576 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.576 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.577 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.577 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:35 | 200 |  4.353335036s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.592 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.592 - [origin:dlm] - register target: stage
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.592 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 09:57:35 Moxa dlm-agent[32634]: [INFO]May 19 09:57:35.592 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success
May 19 09:57:36 Moxa dlm-agent[32634]: [INFO]May 19 09:57:36.642 - [origin:dlm] - tasks[3/6] name:check certificate, result:success
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.718 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.721 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.723 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.738 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 09:57:37 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:37.745 - [origin:dlm-1] - try to connect rule server ...
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.762 - [origin:dlm-1] - connect to rule server success
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.770 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.771 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.778 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.772 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 09:57:37 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:37.769 - [origin:dlm-1] - monitor stream in
May 19 09:57:37 Moxa dlm-agent[32634]: [WARN]May 19 09:57:37.789 - [origin:dlm-1] - waiting for self registration...
May 19 09:57:37 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:37.769 - [origin:dlm-1] - twin stream in
May 19 09:57:37 Moxa dlm-agent[32634]: [WARN]May 19 09:57:37.790 - [origin:dlm-1] - waiting for self registration...
May 19 09:57:37 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:37.770 - [origin:dlm-1] - message stream in
May 19 09:57:37 Moxa dlm-agent[32634]: [INFO]May 19 09:57:37.811 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.785 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.786 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.786 - [origin:dlm-1] - [connection callback], connection = true
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.789 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 09:57:38 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:38 | 200 |     232.876µs |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.803 - [origin:dlm-1] - twin stream registration success
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.803 - [origin:dlm-1] - monitor stream registration success
May 19 09:57:38 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:38 | 200 |     745.508µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:57:38 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:38 | 200 |   53.631966ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.923 - [origin:message-group:dlm-1:2] - stopping
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.931 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.933 - [origin:message-group:dlm-1:2] - stopped
May 19 09:57:38 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:38 | 200 |   18.723935ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.945 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.967 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:57:38 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:38.970 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.989 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.992 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:57:38 Moxa dlm-agent[32634]: [INFO]May 19 09:57:38.999 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.014 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.014 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.015 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.005 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.011 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.031 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.034 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |    8.718818ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.055 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.067 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.055 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.068 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.088 - [origin:message-group:dlm-1:4] - stopping
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.089 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.090 - [origin:message-group:dlm-1:4] - stopped
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |   28.252339ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.111 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.140 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.141 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.148 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.148 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.149 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.150 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.150 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.150 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.191 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.192 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:57:39 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |   76.504433ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.206 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.206 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.206 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.206 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.207 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.207 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |  102.305052ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.253 - [origin:store] - [connectionID:dlm-1] - start
May 19 09:57:39 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |   70.373212ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:57:39 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |   51.695517ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |  102.428729ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.355 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 09:57:39 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:39 | 200 |   46.919635ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 09:57:39 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:39.408 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.413 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.414 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: waiting
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.414 - [connectionID:dlm-1] - [origin:store] - get connection success, start register
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.412 - [origin:dlm] - last ota result empty
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.420 - [origin:dlm-1] - start store message
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.420 - [origin:store] - [connectionID:dlm-1] - register receiver success
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.420 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 09:57:39 Moxa dlm-agent[32634]: [INFO]May 19 09:57:39.986 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.556 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.559 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.718 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.762 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}}
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.763 - [origin:dlm-1] - [twin callback] exit
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.764 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.764 - [origin:dlm-1] - desired push type:$completeDesired, queue: 0/50
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.764 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}}
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     332.395µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.768 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:57:40 Moxa dlm-agent[32634]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 09:57:40 Moxa dlm-agent[32634]: toolbox.NewTask
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     165.997µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |    2.786235ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |   14.038491ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.798 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:57:40 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |   21.190935ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |   27.195317ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:57:40 Moxa dlm-agent[32634]: GetDeviceGeneral
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |   16.165016ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |   59.376152ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     476.592µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     648.189µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     810.027µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     140.478µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |      587.23µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |    8.866335ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:57:40 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:40 | 200 |     196.877µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 09:57:40 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:40.994 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 09:57:40 Moxa dlm-agent[32634]: [INFO]May 19 09:57:40.998 - [origin:dlm-1] - [reported], {"application-dlm-dashboard":{"status":"","version":"","desiredURL":"","reportedURL":""},"ethernets":{"status":"","version":"","desiredURL":"","reportedURL":""},"osPatchDetail":{"status":"","version":"","desiredURL":"","reportedURL":""},"osPatchControl":{"status":"ready"},"osPatchList":{"reportedURL":"","status":"","version":"","desiredURL":""},"edition":"2.0","applications":{"firmwareVersion":"1.4","hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","memorySize":524054528,"serialNumber":"TAIJB1072828","description":"","disk":[{"device":"/dev/root","free":75214848,"mount":"/","name":"System","percent":82.16348920732285,"total":449529856,"used":346475520}],"modelName":"UC-3111-T-EU-LX","osType":"Linux","thingsProProduct":"","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)"},"sshserver":{"enable":true,"port":22},"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"general":{"cpu":"ARMv7 Processor rev 2 (v7l)","firmwareVersion":"1.4","lastBootTime":"2021-05-17T03:13:07Z","modelName":"UC-3111-T-EU-LX","biosVersion":"1.4.0S02","osType":"Linux","serialNumber":"TAIJB1072828","thingsProProduct":"","hostName":"Moxa"}}
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.002 - [origin:dlm-1] - [reported] classic
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.041 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.163 - [origin:dlm-1] - [reported callback], success
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.170 - [origin:dlm-1] - [reported] applications
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.282 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_PARTIAL, Payload =  {"ethernets":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=5s%2FqxF4Bvw5pWXbH25vyWSk%2Bk8A%3D&Expires=1621425459","status":"created"},"osPatchDetail":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=NSiiQFRzfZkK7TqG3tx5zPbReuI%3D&Expires=1621425459","status":"created"},"osPatchList":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=RCjst0Xu0ZBpUFTik4L3t%2FdlKU8%3D&Expires=1621425459","status":"created"}}
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.282 - [origin:dlm-1] - [twin callback] exit
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.283 - [origin:dlm-1] - desired recv type:
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.283 - [origin:dlm-1] - desired push type:, queue: 0/50
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.283 - [origin:devicedb] - connectionID:dlm-1 desired process=>:{"ethernets":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=5s%2FqxF4Bvw5pWXbH25vyWSk%2Bk8A%3D&Expires=1621425459","status":"created"},"osPatchDetail":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=NSiiQFRzfZkK7TqG3tx5zPbReuI%3D&Expires=1621425459","status":"created"},"osPatchList":{"reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=RCjst0Xu0ZBpUFTik4L3t%2FdlKU8%3D&Expires=1621425459","status":"created"}}
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.285 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.285 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:57:41 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:41.285 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:57:41 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:41 | 200 |   39.386158ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.336 - [origin:dlm-1] - [reported callback], success
May 19 09:57:41 Moxa dlm-agent[32634]: [INFO]May 19 09:57:41.799 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 09:57:42 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:42.063 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"status":"connected","broadcast":"10.123.13.255","dns":["10.123.200.11","10.123.200.12"],"gateway":"10.123.12.1","mac":"00:90:e8:82:f2:fb","netmask":"255.255.254.0","speed":"100 Mbps","displayName":"eth0","enable":true,"id":1,"ip":"10.123.13.23","name":"eth0","subnet":"10.123.12.0"},{"speed":"--","dns":[],"enable":true,"gateway":"","id":2,"ip":"192.168.4.127","netmask":"255.255.255.0","broadcast":"192.168.4.255","displayName":"eth1","mac":"00:90:e8:82:f2:fc","name":"eth1","status":"disconnected","subnet":"192.168.4.0"}]}
May 19 09:57:42 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:42 | 200 |     198.916µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:57:42 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:42.122 - [origin:dlm-1] - [command callback] $aws/things/e38751c7-c439-4387-914e-117114e37c31/jobs/$next/get/accepted {"timestamp":1621418262}
May 19 09:57:42 Moxa dlm-agent[32634]: [INFO]May 19 09:57:42.124 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 09:57:42 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:42.754 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:57:42 Moxa dlm-agent[32634]: [GIN] 2021/05/19 - 09:57:42 | 200 |     277.555µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:57:43 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:43.444 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"name":"","newVersion":"","size":0,"version":"","detail":{"origin":"","depends":null,"description":"","label":""}}}
May 19 09:57:43 Moxa dlm-agent[32634]: [DBUG]May 19 09:57:43.446 - [origin:remoted] - connectionID:dlm-1, desired type: done, queue: 0/50
May 19 09:57:43 Moxa dlm-agent[32634]: [INFO]May 19 09:57:43.447 - [origin:dlm-1] - [reported], {"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=NSiiQFRzfZkK7TqG3tx5zPbReuI%3D\u0026Expires=1621425459"},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=5s%2FqxF4Bvw5pWXbH25vyWSk%2Bk8A%3D\u0026Expires=1621425459"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/e38751c7-c439-4387-914e-117114e37c31/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=RCjst0Xu0ZBpUFTik4L3t%2FdlKU8%3D\u0026Expires=1621425459"}}
May 19 09:57:43 Moxa dlm-agent[32634]: [INFO]May 19 09:57:43.447 - [origin:dlm-1] - [reported] classic
May 19 09:57:43 Moxa dlm-agent[32634]: [INFO]May 19 09:57:43.605 - [origin:dlm-1] - [reported callback], success
May 19 09:58:09 Moxa dlm-agent[32634]: [DBUG]May 19 09:58:09.193 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:58:09 Moxa dlm-agent[32634]: [DBUG]May 19 09:58:09.196 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
```