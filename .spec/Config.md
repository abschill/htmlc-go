# Config Module

```go
type HTMLCConfig struct {
	Root   string `json: "root"`   // root path that your htmlc files will be in (relative to config)
	Chunks string `json: "chunks"` // path to read in files for runtime / compiler
}

type HTMLCConfigFile = map[string]Config
```
