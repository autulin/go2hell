# go2hell

## Usage

Implement your any stress job (HTTP, RPC, etc.) from IJob interface
```go
type IJob interface {
	Init()  // do before stress job start
	Exe()   // do stress thing, http request, rpc, etc.
	End()   // do after stress job finished
}
```

then run.