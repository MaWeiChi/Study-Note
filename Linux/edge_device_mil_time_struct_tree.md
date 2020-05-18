```go
type Service struct {
	name       string
	logger     *log.Logger
	timeout    int
	manager    types.ManagerCallback
	configPath string
	time       TimeEntry
	zoneInfo   TimeZoneEntry
	validate   *validator.Validate
	db         *platform.DBClient
	locker     *service.Locker
	mxlog      elog.EventLog
}
```

```go
type NtpEntry struct {
	Enable         bool   `json:"enable"`
	GPS            bool   `json:"gps"`
	Server         string `json:"server" validate:"omitempty,min=2,max=253,ipv4|fqdn"`
	Interval       int    `json:"interval" validate:"omitempty,min=60,max=2592000"`
	updateStatus   int
	cancel         bool
	wg             *sync.WaitGroup
	t              *TimeEntry
}
```

```go
type TimeConfigEntry struct {
	Type     string   `json:"type"`
	Timezone string   `json:"timezone" validate:"omitempty,min=1,max=64,alpha|contains=/"`
	Ntp      NtpEntry `json:"ntp"`
}

type TimeReadOnlyEntry struct {
	LastUpdateTime string `json:"lastUpdateTime" validate:"omitempty"`
}

type TimeEntry struct {
	TimeConfigEntry
	TimeReadOnlyEntry
	Time     string `json:"time" validate:"omitempty,min=20,max=35,timerfc3339"`
	s        *Service
}
```

```go
type ZoneEntry struct {
	Cca2   string `json:"cca2"`
	Name   string `json:"name"`
	Offset string `json:"offset"`
}

type Iso3166Entry struct {
	Cca2 string `json:"cca2"`
	Name string `json:"name"`
}

type TimeZoneEntry struct {
	Iso3166 []Iso3166Entry `json:"iso3166"`
	Zone    []ZoneEntry    `json:"zone"`
}
```

