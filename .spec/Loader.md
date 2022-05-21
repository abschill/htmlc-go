# Loader Module


```go
type HTMLCLoader struct {
	htmlcVersion Version `json: "htmlc_version"`
	config Config `json: "config"`
}
```

The `Loader` interface is what is called by the library at runtime, to be shaped using the inline or saved configuration (via [htmlc.json](Config.md)).
