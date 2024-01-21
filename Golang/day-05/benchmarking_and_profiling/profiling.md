What is Profiling?
Profiling is used to analyze your program's runtime behavior and identify areas that need optimization.

#### Profiling in Go:
- Go's testing package allows you to generate CPU and memory profiles
- You can enable profiling with flags when running benchmarks or tests.

### CPU profiling:
- Run the benchmarks or tests with the -cpuprofile flag to generate a CPU profile:

` go test -bench=. -cpuprofile=cpu.prof`

### Memory profiling:
- Similarly, for memory profiling, use the -memprofile flag:

`go test -bench=. -memprofile=mem.prof`

### Analyzing the profiles:
- Use the go tool pprof command to analyze the profile. For example:

`go tool pprof cpu.prof`

- For better UI Visualisation try: https://www.speedscope.app/ or below commad: `go tool pprof -http=:8080 cpu.prof`

(Install graphviz for this : https://graphviz.org/download/)