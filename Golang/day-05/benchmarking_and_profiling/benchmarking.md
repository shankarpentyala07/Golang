1. What is Benchmarking?
A: benchmarking refers to measuring the performance of certain parts of the code, typically functions , to assess their execution time and resource usage

### Writing Benchmarks in Go:
- Benchmarks in Go are written similar to tests , but they use the testing.B type
- Place benchmark functions in _test.go files.
- Benchmark function names should start with Benchmark.

### Running Benchmarks:
- Use the go test command with the -bench flag:
go test -bench=.

- The output will show the number of iterations and the time per iteration.