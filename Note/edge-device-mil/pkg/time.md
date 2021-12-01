# time

```[go]
type NtpEntry struct {
	Enable           bool   `json:"enable"`
	Source           string `json:"source" validate:"omitempty,oneof=gps timeserver"`
	Server           string `json:"server" validate:"omitempty,min=2,max=253,ipv4|fqdn"`
	GPSLongJump      bool   `json:"gpsLongJump"`
	Interval         int    `json:"interval" validate:"omitempty,min=60,max=2592000"`
	updateStatus     int
	gpsServiceStatus bool
	wg               *sync.WaitGroup
	ctx              context.Context
	cancel           context.CancelFunc
	t                *TimeEntry
}

const (
	NTP_CONFIG_PATH       = "/host/etc/ntp.conf"
	NTP_GPS_SOURCE_CONFIG = `# GPS Serial data reference
server 127.127.28.0 minpoll 4 maxpoll 4
fudge 127.127.28.0 time1 0.0 flag1 1 refid GPS`
	NTP_GPS_SOURCE_CONFIG_WITH_LIMIT = `# GPS Serial data reference
server 127.127.28.0 minpoll 4 maxpoll 4
fudge 127.127.28.0 time1 0.0 time2 14400.0 refid GPS`
	NTP_TIMESERVER_SOURCE_CONFIG = `# Time Server
server %s minpoll 4 maxpoll 4
fudge %s stratum 10`
)
```

In the auto mode of time sync, there are two options, one is to use timeservers with the Internet, another is to use the GPS service with the device module.

We use the following command to sync the time with 15 seconds timeout in golang.

```[go]
cmd.RunWithCancel(ctx, 15, true, "ntpd", "-c", NTP_CONFIG_PATH, "-g", "-q")
```

We modify the `/host/etc/ntp.conf` with NTP_GPS_SOURCE_CONFIG, NTP_GPS_SOURCE_CONFIG_WITH_LIMIT, and NTP_TIMESERVER_SOURCE_CONFIG. 

If the value of `GPSLongJump` is false, we use the NTP_GPS_SOURCE_CONFIG_WITH_LIMIT, otherwise we use the NTP_GPS_SOURCE_CONFIG.

The difference between the two options is that GPS sync time with a default value of 14400 sec (4hrs) when the `GPSLongJump` is false. If the time difference is greater than 14400 sec, whether the time of device is in the future or the past, the time sync will fail.

## [Reference](https://www.eecis.udel.edu/~mills/ntp/html/drivers/driver28.html)

**time2**

Maximum allowed difference between remote and local clock, in seconds. Values <1.0 or >86400.0 are ignored, and the default value of 4hrs (14400s) is used instead. See also flag 1.

**flag1**

Skip the difference limit check if set. Useful for systems where the RTC backup cannot keep the time over long periods without power and the SHM clock must be able to force long-distance initial jumps. Check the difference limit if cleared (default).