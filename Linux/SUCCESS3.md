```
May 19 09:37:51 Moxa systemd[1]: Started dlm-agent service.
May 19 09:37:52 Moxa dlm-agent[23477]: *******************************************************************************
May 19 09:37:52 Moxa dlm-agent[23477]:   _______ _     _                 _____             _____  _                 _
May 19 09:37:52 Moxa dlm-agent[23477]:  |__   __| |   (_)               |  __ \           |  ___|| |               | |
May 19 09:37:52 Moxa dlm-agent[23477]:     | |  | |__  _ _ __   __ _ ___| |__) | __ ___   | |    | | ___  _   _  __| |
May 19 09:37:52 Moxa dlm-agent[23477]:     | |  | '_ \| | '_ \ / _  / __|  ___/ '__/ _ \  | |    | |/ _ \| | | |/ _  |
May 19 09:37:52 Moxa dlm-agent[23477]:     | |  | | | | | | | | (_| \__ \ |   | | | (_) | | |___ | | (_) | |_| | (_| |
May 19 09:37:52 Moxa dlm-agent[23477]:     |_|  |_| |_|_|_| |_|\__, |___/_|   |_|  \___/  |_____||_|\___/\_____/\__,_|
May 19 09:37:52 Moxa dlm-agent[23477]:                          __/ |
May 19 09:37:52 Moxa dlm-agent[23477]:                         |___/
May 19 09:37:52 Moxa dlm-agent[23477]:   Application : ThingsPro Cloud service
May 19 09:37:52 Moxa dlm-agent[23477]:   Version     : 0.1.2-armhf, Time : 2021-05-19T09:37:52Z
May 19 09:37:52 Moxa dlm-agent[23477]: *******************************************************************************
May 19 09:37:52 Moxa dlm-agent[23477]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:37:52 Moxa dlm-agent[23477]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:37:52 Moxa dlm-agent[23477]:  - using env:        export GIN_MODE=release
May 19 09:37:52 Moxa dlm-agent[23477]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:37:52 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/_/info            --> main.getInfo (4 handlers)
May 19 09:37:52 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/devicedb/onchange --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/agent.TransOnChange (4 handlers)
May 19 09:37:52 Moxa dlm-agent[23477]: [openvpn] 2021/05/19 09:37:52 config.go:52: main config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json: no such file or directory,<nil>, try load backup
May 19 09:37:52 Moxa dlm-agent[23477]: [openvpn] 2021/05/19 09:37:52 config.go:62: backup config load error: open /etc/dlm-agent/.data/device/openvpn/configuration.json.backup: no such file or directory,<nil>, try load factory
May 19 09:37:52 Moxa dlm-agent[23477]: MOXA api token open failed, open /run/mx-api-token: no such file or directory
May 19 09:37:53 Moxa dlm-agent[23477]: config load err:open /etc/dlm-agent/.data/device/openvpn/openvpn.json: no such file or directory
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/device/general    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/device/applications --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceGeneral-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/device/ethernets  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernets-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/device/ethernets/:id --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetDeviceEthernetById-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/system/reboot     --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PutSystemReboot-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/openvpn           --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetOpenVpn-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PATCH  /api/v1/openvpn/dlm       --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).PatchOpenVpnDlm-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).GetSSHSettings-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/system/sshserver  --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/device.(*Device).SetSSHSettings-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: config load err:open /etc/dlm-agent/.data/device/software/todo_UpgradeList.json: no such file or directory
May 19 09:37:53 Moxa dlm-agent[23477]: get todo upgrade list error: unexpected end of JSON input
May 19 09:37:53 Moxa dlm-agent[23477]: config load err:open /etc/dlm-agent/.data/device/software/config.json: no such file or directory
May 19 09:37:53 Moxa dlm-agent[23477]: 0 0 * * * *
May 19 09:37:53 Moxa dlm-agent[23477]: toolbox.NewTask
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftware-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/software          --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftware-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/software/ospatchlist --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetPatchList-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/software/detail   --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareDetail-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).GetSoftwareSchedule-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/software/schedule --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/software.(*Software).PutSoftwareSchedule-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetLogs (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/events            --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PostLog (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.GetProfile (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/events/profile    --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/internal/event/api.PatchProfile (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/tags/monitor      --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetMonitor-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/tags/monitor/:args --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).PostMonitor-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/tags/monitor/system/system --> gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/pkg/monitor.(*Monitor).GetTags-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
May 19 09:37:53 Moxa dlm-agent[23477]:  - using env:        export GIN_MODE=release
May 19 09:37:53 Moxa dlm-agent[23477]:  - using code:        gin.SetMode(gin.ReleaseMode)
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.786 - [origin:remoted] - plugin message uploader created
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.786 - [origin:remoted] - plugin devicedb created
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.786 - [origin:remoted] - plugin appman created
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.788 - [origin:remoted] - plugin apiV1 created
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.789 - [origin:remoted] - plugin store created
May 19 09:37:53 Moxa dlm-agent[23477]: [DBUG]May 19 09:37:53.790 - [origin:dlm] - cert path:
May 19 09:37:53 Moxa dlm-agent[23477]: [DBUG]May 19 09:37:53.790 - [origin:dlm] - key path:
May 19 09:37:53 Moxa dlm-agent[23477]: [ERRO]May 19 09:37:53.927 - [origin:dlm] - tasks[1/2] name:[armhf] check cert files, result:failed, error:exit status 1
May 19 09:37:53 Moxa dlm-agent[23477]: [WARN]May 19 09:37:53.927 - [origin:dlm] - load factory certificates failed, error: exit status 1
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.927 - [origin:remoted] - service dlm created
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).getMessagesPolicy-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).putMessagesPolicy-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/remotectl/plugins/messages --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).postMessagesPolicy-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] DELETE /api/v1/remotectl/plugins/messages/:id --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/message.(*Plugin).messageEntry-fm (4 handlers)
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.935 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.934 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.935 - [origin:remoted] - worker sock binding
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.942 - [origin:remoted] - plugin message uploader started
May 19 09:37:53 Moxa dlm-agent[23477]: [DBUG]May 19 09:37:53.953 - [origin:devicedb:dlm-1] - database subscriber start
May 19 09:37:53 Moxa dlm-agent[23477]: [DBUG]May 19 09:37:53.954 - [origin:devicedb:dlm-1] - database subscriber exit, channel empty
May 19 09:37:53 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.954 - [origin:devicedb] - connectionID:dlm-1 db client started
May 19 09:37:53 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSource-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/source --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSource-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).getPropertiesSelections-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/properties --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesSelections-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/onchange --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).putPropertiesOnchange-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/test --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).testPropertiesSelections-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).reported-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/custom/reported --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).userReported-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/devicedb/clear --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/devicedb.(*Plugin).clear-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.955 - [origin:remoted] - plugin devicedb started
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.955 - [origin:remoted] - plugin appman started
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.955 - [origin:remoted] - plugin apiV1 started
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).getConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).putConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/plugins/stores/restart --> github.com/MOXA-ISD/edge-thingspro-agent/plugin/store.(*Plugin).reload-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.955 - [origin:remoted] - plugin store started
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] HEAD   /api/v1/remotectl/index/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] HEAD   /api/v1/remotectl/css/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] HEAD   /api/v1/remotectl/js/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] HEAD   /api/v1/remotectl/img/*filepath --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/info    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).info-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/tags    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).tagMonitor-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/remotectl/services --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).servicesInfo-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/message --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).sendMessage-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/connections/:id/state --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/remote.(*Worker).updateState-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.956 - [origin:dlm] - skip start, false
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).postConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PATCH  /api/v1/dlm               --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).patchConfiguration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/_/device      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDeviceID-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/monitor       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getMonitor-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/registration  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putRegistration-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/enrollment    --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putEnrollment-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/edition       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getEdition-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/control       --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).controlThing-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/login         --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).login-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] POST   /api/v1/dlm/messages      --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] DELETE /api/v1/dlm/messages/:id  --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).messagesPolicy-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).getDashboard-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/dashboard     --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).putDashboard-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).appsInfo-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PATCH  /api/v1/dlm/apps          --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).updateAppsInfo-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] GET    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/dlm/store-and-forward --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).storeHandler-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-export --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).exportHandler-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [GIN-debug] PUT    /api/v1/remotectl/commands/dlm-import --> github.com/MOXA-ISD/edge-thingspro-agent/pkg/dlm.(*Dlm).importHandler-fm (4 handlers)
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.972 - [origin:remoted] - service dlm started
May 19 09:37:54 Moxa dlm-agent[23477]: [INFO]May 19 09:37:53.972 - [origin:remoted] - all started
May 19 09:38:15 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:15 | 200 |    1.673173ms |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:15 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:15 | 200 |     226.637µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:15 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:15 | 200 |     228.636µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:17 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:17 | 200 |     326.594µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:17 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:17 | 200 |     218.157µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:17 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:17 | 200 |     220.796µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:24 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:24 | 200 |     227.876µs |       127.0.0.1 | GET      "/api/v1/dlm"
May 19 09:38:29 Moxa dlm-agent[23477]: [INFO]May 19 09:38:29.586 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success
May 19 09:38:30 Moxa dlm-agent[23477]: [INFO]May 19 09:38:30.219 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:38:30 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:30 | 200 |  1.451325744s |       127.0.0.1 | PUT      "/api/v1/dlm/registration?type=AGT"


May 19 09:38:33 Moxa dlm-agent[23477]: [INFO]May 19 09:38:33.312 - [origin:dlm] - skip stop because connection status: false
May 19 09:38:34 Moxa dlm-agent[23477]: [INFO]May 19 09:38:34.086 - [origin:dlm] - tasks[1/2] name:login to dlm, result:success



May 19 09:38:34 Moxa dlm-agent[23477]: [INFO]May 19 09:38:34.709 - [origin:dlm] - tasks[2/2] name:get project list, result:success
May 19 09:38:34 Moxa dlm-agent[23477]: [INFO]May 19 09:38:34.712 - [origin:dlm] - tasks[1/3] name:registration, result:success
May 19 09:38:34 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:34 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:34 | 200 |   15.098633ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:38:34 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:34 | 200 |   66.871906ms |       127.0.0.1 | GET      "/api/v1/device/ethernets/1"
May 19 09:38:34 Moxa dlm-agent[23477]: [INFO]May 19 09:38:34.812 - [origin:dlm] - tasks[1/2] name:prepare enroll parameters, result:success
May 19 09:38:34 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:34 | 200 |     303.836µs |       127.0.0.1 | GET      "/api/v1/_/info"




May 19 09:38:37 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:37.186 - [origin:dlm] - cert path: /etc/dlm-agent/.data/enroll/dev.crt
May 19 09:38:37 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:37.188 - [origin:dlm] - key path: /etc/dlm-agent/.data/enroll/dev.key
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.229 - [origin:dlm] - latest configuration:  {
May 19 09:38:37 Moxa dlm-agent[23477]:   "connection": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "enable": false,
May 19 09:38:37 Moxa dlm-agent[23477]:     "retryDelaySec": 10,
May 19 09:38:37 Moxa dlm-agent[23477]:     "picTarget": "stage"
May 19 09:38:37 Moxa dlm-agent[23477]:   },
May 19 09:38:37 Moxa dlm-agent[23477]:   "certificates": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "caCertFile": "dev.crt",
May 19 09:38:37 Moxa dlm-agent[23477]:     "caPkFile": "dev.key",
May 19 09:38:37 Moxa dlm-agent[23477]:     "certificateInfo": {
May 19 09:38:37 Moxa dlm-agent[23477]:       "organization": "Moxa Inc.",
May 19 09:38:37 Moxa dlm-agent[23477]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:38:37 Moxa dlm-agent[23477]:       "notBefore": "May 19, 2021 09:38:35",
May 19 09:38:37 Moxa dlm-agent[23477]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:38:37 Moxa dlm-agent[23477]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:38:37 Moxa dlm-agent[23477]:       "macAddress": "0090E882F2FB",
May 19 09:38:37 Moxa dlm-agent[23477]:       "serialNumber": "TAIJB1072828"
May 19 09:38:37 Moxa dlm-agent[23477]:     }
May 19 09:38:37 Moxa dlm-agent[23477]:   },
May 19 09:38:37 Moxa dlm-agent[23477]:   "connectionStatus": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "projectName": "Kevin Test - AGT",
May 19 09:38:37 Moxa dlm-agent[23477]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:38:37 Moxa dlm-agent[23477]:     "status": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "message": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastConnectedTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastConnectTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastDisconnectTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "statuslastUpdateTime": ""
May 19 09:38:37 Moxa dlm-agent[23477]:   }
May 19 09:38:37 Moxa dlm-agent[23477]: }
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.243 - [origin:dlm] - tasks[2/2] name:enroll, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.243 - [origin:dlm] - tasks[2/3] name:enrollment, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.244 - [origin:dlm] - latest configuration:  {
May 19 09:38:37 Moxa dlm-agent[23477]:   "connection": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "enable": true,
May 19 09:38:37 Moxa dlm-agent[23477]:     "retryDelaySec": 10,
May 19 09:38:37 Moxa dlm-agent[23477]:     "picTarget": "stage"
May 19 09:38:37 Moxa dlm-agent[23477]:   },
May 19 09:38:37 Moxa dlm-agent[23477]:   "certificates": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "caCertFile": "dev.crt",
May 19 09:38:37 Moxa dlm-agent[23477]:     "caPkFile": "dev.key",
May 19 09:38:37 Moxa dlm-agent[23477]:     "certificateInfo": {
May 19 09:38:37 Moxa dlm-agent[23477]:       "organization": "Moxa Inc.",
May 19 09:38:37 Moxa dlm-agent[23477]:       "issuerCN": "moxa-thingspro-device-intermediate",
May 19 09:38:37 Moxa dlm-agent[23477]:       "notBefore": "May 19, 2021 09:38:35",
May 19 09:38:37 Moxa dlm-agent[23477]:       "notAfter": "Dec 31, 2024 23:59:59",
May 19 09:38:37 Moxa dlm-agent[23477]:       "modelName": "UC-3111-T-EU-LX",
May 19 09:38:37 Moxa dlm-agent[23477]:       "macAddress": "0090E882F2FB",
May 19 09:38:37 Moxa dlm-agent[23477]:       "serialNumber": "TAIJB1072828"
May 19 09:38:37 Moxa dlm-agent[23477]:     }
May 19 09:38:37 Moxa dlm-agent[23477]:   },
May 19 09:38:37 Moxa dlm-agent[23477]:   "connectionStatus": {
May 19 09:38:37 Moxa dlm-agent[23477]:     "projectName": "Kevin Test - AGT",
May 19 09:38:37 Moxa dlm-agent[23477]:     "projectID": "d1515d8a-fba3-494f-be25-53f541379b50",
May 19 09:38:37 Moxa dlm-agent[23477]:     "status": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "message": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastConnectedTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastConnectTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "lastDisconnectTime": "",
May 19 09:38:37 Moxa dlm-agent[23477]:     "statuslastUpdateTime": ""
May 19 09:38:37 Moxa dlm-agent[23477]:   }
May 19 09:38:37 Moxa dlm-agent[23477]: }
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.246 - [origin:dlm] - tasks[3/3] name:setup, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:37 | 200 |    3.521263ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/devicedb"
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.278 - [origin:dlm] - tasks[1/3] name:enable shadow env, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.278 - [origin:dlm] - Job:start thingspro cloud apps monitor doing
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.278 - [origin:dlm] - tasks[1/5] name:init self, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [ERRO]May 19 09:38:37.280 - [origin:dlm] - tasks[2/5] name:connect to appman runtime event channel, result:failed, error:Create app runtime pubsub receive failed(dial tcp 127.0.0.1:6379: connect: connection refused)
May 19 09:38:37 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:37 | 404 |         8.6µs |       127.0.0.1 | GET      "/api/v1/apps"
May 19 09:38:37 Moxa dlm-agent[23477]: [ERRO]May 19 09:38:37.283 - [origin:dlm] - get apps info failed(code:404, msg:404 page not found)
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.283 - [origin:dlm] - tasks[3/5] name:init apps info, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.283 - [origin:dlm] - tasks[4/5] name:init cloud services info, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [ERRO]May 19 09:38:37.283 - [origin:dlm] - tasks[5/5] name:start appman runtime event channel listener, result:failed, error:skip monitor because event channel is nil
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.283 - [origin:dlm] - Job:start thingspro cloud apps monitor done
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.283 - [origin:dlm] - tasks[2/3] name:enable app monitor, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.283 - [origin:dlm] - tasks[3/3] name:start connection, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:37 | 200 |  3.971836325s |       127.0.0.1 | PUT      "/api/v1/dlm/login?project=Kevin%20Test%20-%20AGT&type=AGT"
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.288 - [origin:dlm] - tasks[1/6] name:reset connection state, result:success
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.290 - [origin:dlm] - register target: stage
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.291 - [origin:dlm] - register device: {"mac":"0090E882F2FB","serialNumber":"TAIJB1072828","modelName":"UC-3111-T-EU-LX"}
May 19 09:38:37 Moxa dlm-agent[23477]: [INFO]May 19 09:38:37.292 - [origin:dlm] - tasks[2/6] name:setup register parameters, result:success


May 19 09:38:38 Moxa dlm-agent[23477]: [INFO]May 19 09:38:38.217 - [origin:dlm] - tasks[3/6] name:check certificate, result:success


May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.040 - [origin:dlm] - tasks[4/6] name:register, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.043 - [origin:dlm] - projectName: Kevin Test - AGT
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.045 - [origin:dlm] - projectId: d1515d8a-fba3-494f-be25-53f541379b50
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.073 - [origin:dlm] - tasks[5/6] name:setup target connection parameters, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.076 - [origin:dlm-1] - try to connect rule server ...
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.089 - [origin:dlm-1] - connect to rule server success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.092 - [origin:dlm-1] - tasks[1/5] name:init, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.094 - [origin:dlm-1] - tasks[2/5] name:start tunnel mgmt skip
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.094 - [origin:dlm-1] - tasks[3/5] name:disconnect last session skip
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.106 - [origin:dlm-1] - tasks[4/5] name:init mqtt client, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.114 - [origin:dlm-1] - dlm-1 callbacks service started
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.130 - [origin:dlm-1] - monitor stream in
May 19 09:38:39 Moxa dlm-agent[23477]: [WARN]May 19 09:38:39.133 - [origin:dlm-1] - waiting for self registration...
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.135 - [origin:dlm-1] - twin stream in
May 19 09:38:39 Moxa dlm-agent[23477]: [WARN]May 19 09:38:39.136 - [origin:dlm-1] - waiting for self registration...
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.136 - [origin:dlm-1] - message stream in

May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.899 - [origin:dlm-1] - tasks[5/5] name:connect, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.902 - [origin:dlm] - tasks[6/6] name:connect to target, result:success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.902 - [origin:dlm-1] - [connection callback], connection = true
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.912 - [origin:remoted] - connection[dlm-1] connection status: true
May 19 09:38:39 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:39 | 200 |     919.265µs |       127.0.0.1 | PUT      "/api/v1/remotectl/connections/dlm-1/state?state=0"
May 19 09:38:39 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:39 | 200 |     734.268µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:38:39 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:39 | 200 |     4.91032ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/1"
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.944 - [origin:message-group:dlm-1:2] - stopping
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.947 - [origin:message-group:dlm-1:2] - main function stopped, start shutdown process
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.949 - [origin:message-group:dlm-1:2] - stopped
May 19 09:38:39 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:39 | 200 |   10.772464ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/2"
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.954 - [origin:message-group:dlm-1:2] - waiting target connection ready
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.954 - [origin:message-group:dlm-1:2] - next state connection ready
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.954 - [origin:message-group:dlm-1:2] - try to connect rule server ...
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.966 - [origin:message-group:dlm-1:2] - connect to rule server success
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.966 - [origin:message-group:dlm-1:2] - outputservice state: running
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.972 - [origin:message-group:dlm-1:2] - events state: running
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.972 - [origin:message-group:dlm-1:2] - events state: stopped
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.972 - [origin:message-group:dlm-1:2] - event origin:OS patch listener start
May 19 09:38:39 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.972 - [origin:message-group:dlm-1:2] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20upgrading%20process
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.973 - [origin:message-group:dlm-1:2] - report goroutine exit, plugin is not handle twin stream
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.973 - [origin:message-group:dlm-1:2] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.976 - [origin:message-group:dlm-1:2] - monitor stream in
May 19 09:38:39 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.976 - [origin:message-group:dlm-1:2] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.988 - [origin:message-group:dlm-1:2] - twin stream in
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.988 - [origin:message-group:dlm-1:2] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:39.988 - [origin:message-group:dlm-1:2] - message stream in
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:39.989 - [origin:remoted] - [origin:message-group:dlm-1:2] - connected to connectionID[dlm-1]
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:39 | 200 |     4.24393ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/3"
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.013 - [origin:message-group:dlm-1:4] - stopping
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.014 - [origin:message-group:dlm-1:4] - main function stopped, start shutdown process
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.014 - [origin:message-group:dlm-1:4] - stopped
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |    6.000022ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/messages/4"
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.020 - [origin:message-group:dlm-1:4] - waiting target connection ready
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.021 - [origin:message-group:dlm-1:4] - next state connection ready
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.021 - [origin:message-group:dlm-1:4] - try to connect rule server ...
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.034 - [origin:message-group:dlm-1:4] - connect to rule server success
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.034 - [origin:message-group:dlm-1:4] - outputservice state: running
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.035 - [origin:message-group:dlm-1:4] - events state: running
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.035 - [origin:message-group:dlm-1:4] - events state: stopped
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.035 - [origin:message-group:dlm-1:4] - event origin:OS patch listener start
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.041 - [origin:message-group:dlm-1:4] - listening events url: http://127.0.0.1:58000/api/v1/events?event=true&categories=OS%20patch&eventNames=OS%20patch%20schedule%20job%20skip,OS%20patch%20schedule%20job%20fail,OS%20patch%20upgrade%20success,OS%20patch%20upgrade%20fail,New%20OS%20patch%20available
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.043 - [origin:message-group:dlm-1:4] - report goroutine exit, plugin is not handle twin stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.043 - [origin:message-group:dlm-1:4] - monitor goroutine exit, plugin is not handle monitor stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.049 - [origin:message-group:dlm-1:4] - monitor stream in
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.049 - [origin:message-group:dlm-1:4] - monitor stream out, message: plugin client no need to handle monitor stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.050 - [origin:message-group:dlm-1:4] - twin stream in
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.050 - [origin:message-group:dlm-1:4] - twin stream out, message: plugin client no need to handle twin stream
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.050 - [origin:message-group:dlm-1:4] - message stream in
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.050 - [origin:remoted] - [origin:message-group:dlm-1:4] - connected to connectionID[dlm-1]
May 19 09:38:40 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |   25.115709ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |   50.510494ms |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/stores"
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.112 - [origin:store] - [connectionID:dlm-1] - start
May 19 09:38:40 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |   12.919749ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |    28.75777ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/stores"
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.138 - [origin:dlm-1] - twin stream registration success
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.138 - [origin:dlm-1] - monitor stream registration success
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.139 - [origin:dlm] - last ota state empty, open /etc/dlm-agent/.data/dlm/ota_state.json: no such file or directory
May 19 09:38:40 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |   42.629103ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:38:40 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:40 | 200 |   58.807678ms |       127.0.0.1 | GET      "/api/v1/events?categories=system&eventNames=software%20installation%20completed,software%20installation%20failed,upgrade%20roll%20back%20completed&limit=1"
May 19 09:38:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:40.235 - [origin:store] - [connectionID:dlm-1] - origin: DLM App
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.240 - [origin:store] - [connectionID:dlm-1] - store rountine started
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.250 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: waiting
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.250 - [origin:store] - [connectionID:dlm-1] - get connection success, start register
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.244 - [origin:dlm] - last ota result empty
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.268 - [origin:dlm-1] - start store message
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.269 - [origin:store] - [connectionID:dlm-1] - register receiver success
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.269 - [origin:store] - [connectionID:dlm-1] - connection[dlm-1] state: ready
May 19 09:38:40 Moxa dlm-agent[23477]: [INFO]May 19 09:38:40.795 - [origin:dlm-1] - tasks[1/7] name:subscribe classic shadow, result:success
May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.334 - [origin:dlm-1] - tasks[2/7] name:subscribe sub-shadows, result:success
May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.334 - [origin:dlm-1] - tasks[3/7] name:subscribe tunnels notify skip
May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.430 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_PARTIAL, Payload =  {"application-dlm-dashboard":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/application-dlm-dashboard.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Q1Is93HEkEBD0CbDmnSvuLZagHk%3D&Expires=1621424320"}}
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.431 - [origin:dlm-1] - [twin callback] exit
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.432 - [origin:dlm-1] - desired recv type:
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.433 - [origin:dlm-1] - desired push type:, queue: 0/50
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.433 - [origin:devicedb] - connectionID:dlm-1 desired process=>:{"application-dlm-dashboard":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/application-dlm-dashboard.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Q1Is93HEkEBD0CbDmnSvuLZagHk%3D&Expires=1621424320"}}
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.435 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: application-dlm-dashboard
May 19 09:38:41 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:41 | 200 |      594.07µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:38:41 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:41 | 200 |   18.217622ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.482 - [origin:dlm-1] - tasks[4/7] name:get classic shadow, result:success
May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.520 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_COMPLETE, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}}
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.520 - [origin:dlm-1] - [twin callback] exit
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.521 - [origin:dlm-1] - desired recv type:$completeDesired
May 19 09:38:41 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:41.521 - [origin:dlm-1] - desired push type:$completeDesired, queue: 1/50

May 19 09:38:41 Moxa dlm-agent[23477]: [INFO]May 19 09:38:41.794 - [origin:dlm-1] - tasks[5/7] name:get sub-shadows, result:success
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.149 - [origin:devicedb:dlm-1] - put reported success. data: {"application-dlm-dashboard":[{"onChange":false,"outputTopic":"dlm/d4db0bf0-80e0-41c4-a4de-db15ded68014/dashboard","pollingInterval":5,"properties":[{"value":"dashboard","key":"class"}],"sendOutThreshold":{"time":0,"mode":"bySize","size":5120,"sizeIdleTimer":{"enable":true,"time":30}},"description":"","enable":false,"format":"{tags:{messageTimeStamp:now,(.srcName):{(.tagName):(.dataValue|fromjson)}}}","tags":{"system":{"system":["network","storage","cpu","memory"]}},"samplingMode":"allValues","customSamplingRate":true,"id":3,"minPublishInterval":0}]}
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.154 - [origin:remoted] - connectionID:dlm-1, desired type: done, queue: 1/50
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.163 - [origin:devicedb] - connectionID:dlm-1 desired process=>$completeDesired:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}}
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.158 - [origin:dlm-1] - [reported], {"application-dlm-dashboard":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/application-dlm-dashboard.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=Q1Is93HEkEBD0CbDmnSvuLZagHk%3D\u0026Expires=1621424320"}}
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.166 - [origin:dlm-1] - [reported] applications
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     345.755µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.172 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:38:42 Moxa dlm-agent[23477]: 0 0 14 * * 0,1,2,3,4,5,6
May 19 09:38:42 Moxa dlm-agent[23477]: toolbox.NewTask
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     160.878µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |    2.758155ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   26.157412ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.217 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:38:42 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   17.408355ms |       127.0.0.1 | GET      "/api/v1/device/general"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   26.948559ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:38:42 Moxa dlm-agent[23477]: GetDeviceGeneral
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   16.241494ms |       127.0.0.1 | GET      "/api/v1/device/applications"
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.316 - [origin:dlm-1] - [reported callback], success
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   78.817751ms |       127.0.0.1 | GET      "/api/v1/system/sshserver"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     683.109µs |       127.0.0.1 | GET      "/api/v1/software"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     723.028µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |    1.063063ms |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     168.637µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     727.548µs |       127.0.0.1 | GET      "/api/v1/remotectl/plugins/messages"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |    9.995996ms |       127.0.0.1 | GET      "/api/v1/dlm/dashboard"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     294.195µs |       127.0.0.1 | GET      "/api/v1/dlm/edition"
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.464 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_PARTIAL, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D&Expires=1621424320"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D&Expires=1621424320"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D&Expires=1621424320"}}
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.466 - [origin:dlm-1] - [twin callback] exit
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.473 - [origin:remoted] - connectionID:dlm-1, desired type:$completeDesired done, queue: 0/50
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.476 - [origin:dlm-1] - desired recv type:
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.478 - [origin:dlm-1] - desired push type:, queue: 0/50
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.479 - [origin:devicedb] - connectionID:dlm-1 desired process=>:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D&Expires=1621424320"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D&Expires=1621424320"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D&Expires=1621424320"}}
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.479 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.479 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.479 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.492 - [origin:dlm-1] - [reported], {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"applications":{"description":"","hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","memorySize":524054528,"osType":"Linux","serialNumber":"TAIJB1072828","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)","disk":[{"device":"/dev/root","free":75214848,"mount":"/","name":"System","percent":82.16348920732285,"total":449529856,"used":346475520}],"firmwareVersion":"1.4","modelName":"UC-3111-T-EU-LX","thingsProProduct":""},"general":{"hostName":"Moxa","lastBootTime":"2021-05-17T03:13:07Z","modelName":"UC-3111-T-EU-LX","serialNumber":"TAIJB1072828","thingsProProduct":"","biosVersion":"1.4.0S02","cpu":"ARMv7 Processor rev 2 (v7l)","firmwareVersion":"1.4","osType":"Linux"},"osPatchDetail":{"reportedURL":"","status":"","version":"","desiredURL":""},"osPatchControl":{"status":"ready"},"application-dlm-dashboard":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/application-dlm-dashboard.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=Q1Is93HEkEBD0CbDmnSvuLZagHk%3D\u0026Expires=1621424320"},"edition":"2.0","ethernets":{"desiredURL":"","reportedURL":"","status":"","version":""},"sshserver":{"enable":true,"port":22},"osPatchList":{"status":"","version":"","desiredURL":"","reportedURL":""}}
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.494 - [origin:dlm-1] - [reported] classic
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     153.638µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.508 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |     155.398µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |    5.155835ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |    9.993876ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.527 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.531 - [origin:dlm-1] - tasks[6/7] name:subscribe job topics, result:success
May 19 09:38:42 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:42 | 200 |   28.254778ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.608 - [origin:dlm-1] - [twin callback], DEVICE_TWIN_PARTIAL, Payload =  {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D&Expires=1621424320"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D&Expires=1621424320"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D&Expires=1621424320"}}
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.611 - [origin:dlm-1] - [twin callback] exit
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.612 - [origin:dlm-1] - desired recv type:
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.616 - [origin:dlm-1] - desired push type:, queue: 1/50
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.642 - [origin:dlm-1] - [reported callback], success
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.643 - [origin:dlm-1] - [reported] applications
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.792 - [origin:dlm-1] - [reported callback], success
May 19 09:38:42 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:42.830 - [origin:dlm-1] - [command callback] $aws/things/d4db0bf0-80e0-41c4-a4de-db15ded68014/jobs/$next/get/accepted {"timestamp":1621417122}
May 19 09:38:42 Moxa dlm-agent[23477]: [INFO]May 19 09:38:42.830 - [origin:dlm-1] - tasks[7/7] name:get pending jobs, result:success
May 19 09:38:43 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:43.237 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"displayName":"eth0","gateway":"10.123.12.1","id":1,"ip":"10.123.13.23","speed":"100 Mbps","status":"connected","broadcast":"10.123.13.255","dns":["10.123.200.11","10.123.200.12"],"enable":true,"mac":"00:90:e8:82:f2:fb","name":"eth0","netmask":"255.255.254.0","subnet":"10.123.12.0"},{"dns":[],"name":"eth1","speed":"--","subnet":"192.168.4.0","ip":"192.168.4.127","mac":"00:90:e8:82:f2:fc","netmask":"255.255.255.0","broadcast":"192.168.4.255","displayName":"eth1","enable":true,"gateway":"","id":2,"status":"disconnected"}]}
May 19 09:38:43 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:43 | 200 |     201.237µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:38:43 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:43.955 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:38:43 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:43 | 200 |     314.754µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.610 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"detail":{"label":"","origin":"","depends":null,"description":""},"name":"","newVersion":"","size":0,"version":""}}
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.615 - [origin:remoted] - connectionID:dlm-1, desired type: done, queue: 1/50
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.619 - [origin:devicedb] - connectionID:dlm-1 desired process=>:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D&Expires=1621424320"},"osPatchList":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D&Expires=1621424320"},"osPatchDetail":{"status":"created","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON&Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D&Expires=1621424320"}}
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.622 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: ethernets
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.622 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchList
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.622 - [origin:devicedb:dlm-1] - skip download because desired URL empty, path: osPatchDetail
May 19 09:38:44 Moxa dlm-agent[23477]: [INFO]May 19 09:38:44.618 - [origin:dlm-1] - [reported], {"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D\u0026Expires=1621424320"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D\u0026Expires=1621424320"},"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D\u0026Expires=1621424320"}}
May 19 09:38:44 Moxa dlm-agent[23477]: [INFO]May 19 09:38:44.625 - [origin:dlm-1] - [reported] classic
May 19 09:38:44 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:44 | 200 |     138.597µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:44 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:44.631 - [origin:devicedb:dlm-1] - property:osPatchCronJob new desired payload:{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"}
May 19 09:38:44 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:44 | 200 |     156.997µs |       127.0.0.1 | GET      "/api/v1/software/schedule"
May 19 09:38:44 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:44 | 200 |    2.800674ms |       127.0.0.1 | PUT      "/api/v1/remotectl/plugins/devicedb/onchange?name=osPatch&args=cron"
May 19 09:38:44 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:44 | 200 |   11.570971ms |       127.0.0.1 | PUT      "/api/v1/software/schedule"
May 19 09:38:44 Moxa dlm-agent[23477]: [INFO]May 19 09:38:44.653 - [origin:devicedb:dlm-1] - set desired property:/software/schedule, success
May 19 09:38:44 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:44 | 200 |   34.012164ms |       127.0.0.1 | GET      "/api/v1/device/ethernets"
May 19 09:38:44 Moxa dlm-agent[23477]: [INFO]May 19 09:38:44.774 - [origin:dlm-1] - [reported callback], success
May 19 09:38:45 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:45.390 - [origin:devicedb:dlm-1] - put reported success. data: {"ethernets":[{"id":1,"name":"eth0","subnet":"10.123.12.0","broadcast":"10.123.13.255","displayName":"eth0","dns":["10.123.200.11","10.123.200.12"],"mac":"00:90:e8:82:f2:fb","netmask":"255.255.254.0","speed":"100 Mbps","status":"connected","enable":true,"gateway":"10.123.12.1","ip":"10.123.13.23"},{"dns":[],"id":2,"mac":"00:90:e8:82:f2:fc","netmask":"255.255.255.0","subnet":"192.168.4.0","status":"disconnected","broadcast":"192.168.4.255","displayName":"eth1","enable":true,"gateway":"","ip":"192.168.4.127","name":"eth1","speed":"--"}]}
May 19 09:38:45 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:45 | 200 |     231.796µs |       127.0.0.1 | GET      "/api/v1/software/ospatchlist"
May 19 09:38:46 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:46.054 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchList":[]}
May 19 09:38:46 Moxa dlm-agent[23477]: [GIN] 2021/05/19 - 09:38:46 | 200 |     174.557µs |       127.0.0.1 | GET      "/api/v1/software/detail"
May 19 09:38:46 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:46.768 - [origin:devicedb:dlm-1] - put reported success. data: {"osPatchDetail":{"detail":{"depends":null,"description":"","label":"","origin":""},"name":"","newVersion":"","size":0,"version":""}}
May 19 09:38:46 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:46.772 - [origin:devicedb:dlm-1] - skip reported because payload same as last time, last:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D\u0026Expires=1621424320"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D\u0026Expires=1621424320"},"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDetail.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D\u0026Expires=1621424320"}}, now:{"osPatchCronJob":{"enable":true,"updateSchedule":"0 14 * * 0,1,2,3,4,5,6"},"ethernets":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/ethernets.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=Yfn0ofDnTNmGp6HoITgIVs3iQas%3D\u0026Expires=1621424320"},"osPatchList":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchList.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=DRXMyAMHrjJa75hPBgomWV39oxE%3D\u0026Expires=1621424320"},"osPatchDetail":{"status":"success","version":"","desiredURL":"","reportedURL":"https://dlm-infra-dlmstorage-vporf2en8gcj.s3.amazonaws.com/devices/d4db0bf0-80e0-41c4-a4de-db15ded68014/reported/osPatchDeta
May 19 09:38:46 Moxa dlm-agent[23477]: il.json?AWSAccessKeyId=AKIATM2PBA27IDVPLFON\u0026Signature=3Iepi4rpcNaTXCSx4gL5gN2%2Bcm4%3D\u0026Expires=1621424320"}}
May 19 09:38:46 Moxa dlm-agent[23477]: [DBUG]May 19 09:38:46.777 - [origin:remoted] - connectionID:dlm-1, desired type: done, queue: 0/50
May 19 09:39:10 Moxa dlm-agent[23477]: [DBUG]May 19 09:39:10.044 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:39:10 Moxa dlm-agent[23477]: [DBUG]May 19 09:39:10.047 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
May 19 09:39:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:39:40.044 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:39:40 Moxa dlm-agent[23477]: [DBUG]May 19 09:39:40.044 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
May 19 09:40:10 Moxa dlm-agent[23477]: [DBUG]May 19 09:40:10.044 - [origin:message-group:dlm-1:4] - tag count: 0/0, buffer size: 2/5120
May 19 09:40:10 Moxa dlm-agent[23477]: [DBUG]May 19 09:40:10.044 - [origin:message-group:dlm-1:4] - buffer empty, skip flush
```

```
pytest -m "not wificlient" --tavern-beta-new-traceback --html=all.html
```

```
DLM Cloud connection status         : connectFailure
DLM Cloud connection fail reason    : perform register request failed, err:Post https://pic-dev.thingspro.dev/api/v1/register: dial tcp: lookup pic-dev.thingspro.dev on 192.168.50.1:53: read udp 192.168.50.91:38683->192.168.50.1:53: i/o timeout
```



```
dlm-agent v 1.0.0
linux:
    deb:
        armhf: http://ibg-nas.moxa.com:5000/sharing/yIrv6VbvS
        amd64: http://ibg-nas.moxa.com:5000/sharing/fMITdj0kS
    tar: ( https://gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/-/blob/develop/docs/linux/README.md#install-dlm-agenttar)
        armhf: http://ibg-nas.moxa.com:5000/sharing/LNrAtVsuE
        amd64: http://ibg-nas.moxa.com:5000/sharing/rE6ID6sfq
windows:
    msi: http://ibg-nas.moxa.com:5000/sharing/LCek6QJ3q
changelog:
https://gitlab.com/moxa/ibg/software/platform/thingspro/dlm-agent/-/tags/dlm-agent-v1.0.0
prod: http://ibg-edge-s3.moxa.com/v3/edge/builds/dlm-agent/release/prod/v1.0.0/501/
```

