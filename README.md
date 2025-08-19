# pprof tutorial

## Profile Descriptions

Available in `/debug/pprof/`.

- profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
- heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
- goroutine: Stack traces of all current goroutines. Use debug=2 as a query parameter to export in the same format as an unrecovered panic.
- allocs: A sampling of all past memory allocations
- block: Stack traces that led to blocking on synchronization primitives
- cmdline: The command line invocation of the current program
- mutex: Stack traces of holders of contended mutexes
- threadcreate: Stack traces that led to the creation of new OS threads
- trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.

## How to enable pprof

- [net/http/pprof](https://pkg.go.dev/net/http/pprof): Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.
- [runtime/pprof](https://pkg.go.dev/runtime/pprof): Package pprof writes runtime profiling data in the format expected by the pprof visualization tool.


## Getting Started

1. **Run the server:**
   ```bash
   go run .
   ```

2. **Access the server:**
   - Server runs on `http://localhost:8080`
   - Usage information: `http://localhost:8080/`
   - Profiling dashboard: `http://localhost:8080/debug/pprof/`

## How to collect profiles

```bash
curl http://localhost:8080/cpu
curl 'http://localhost:8080/heap?size_mb=500' && curl 'http://localhost:8080/debug/pprof/heap' --output heap.out
```

## How to view profiles

```bash
go tool pprof -http 0.0.0.0:8081 ./cpu.out
```

Note: you can directly collect and view in one command: `go tool pprof -http 0.0.0.0:8081 http://localhost:8080/debug/pprof/profile`

## Examples

- Outlier block profiling [Nitro](https://github.com/OffchainLabs/nitro/blob/90570c4bd330bd23321b9e4ca9e41440ab544d2a/execution/gethexec/executionengine.go#L512-L527)
