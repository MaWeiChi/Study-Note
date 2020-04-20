### Under GO/Src :

Settings.json:

```json
    "go.formatTool": "gofmt",
    "go.inferGopath":true,
    "go.useLanguageServer": true,
    "go.toolsEnvVars": {"GO111MODULE": "off",},
```

```go
Import "C"  //could not import C (no package fot import C)
```

```go
listenForInterrupt(chCancel) // undeclared name: listenForInterrupt			
```

Settings.json:

```json
 // "go.toolsEnvVars": {"GO111MODULE": "off",},
		"go.toolsEnvVars": {"GO111MODULE": "on",},  
```

```
Your workspace is misconfigured: go [-e -json -compiled=true -test=true -export=false -deps=true -find=false -- ./]: exit status 1: go: cannot find main module; see 'go help modules' . Please see https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md for more information or file an issue (https://github.com/golang/go/issues/new) if you believe this is a mistake.
```

```vscode
The "inferGopath" setting is disabled for this workspace because Go modules are being used.
```





### Out GO/:

```go
import (
	"github.com/gin-gonic/gin" // could not import github.com/gin-gonic/gin(no package...)
)
```

Any action:

```vscode
You are neither in a module nor in your GOPATH. Please see https://github.com/golang/go/wiki/Modules for information on how to set up your GO project.
```

Settings.json:

```json
 "go.toolsEnvVars": {"GO111MODULE": "on",},
```

```
Your workspace is misconfigured: go [-e -json -compiled=true -test=true -export=false -deps=true -find=false -- ./]: exit status 1: go: cannot find main module; see 'go help modules' . Please see https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md for more information or file an issue (https://github.com/golang/go/issues/new) if you believe this is a mistake.
```

```vscode
The "inferGopath" setting is disabled for this workspace because Go modules are being used.
```



