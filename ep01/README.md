# ep01: Benchmark

## Folders

| Folder | Count5xx |
|---|---|
| [before](before) | Splits the whole line with `strings.Split()` |
| [after](after) | Cuts only up to the status code with `strings.Cut()` |
| [results](results) | The results shown in the lesson |

The tests and benchmarks are the same in both folders. Only `Count5xx()` is different.

## Run it

Start in the `ep01` folder.

Check that the code works, run the benchmark, then add memory allocations:

```bash
cd before
go test
go test -bench=.
go test -bench=. -benchmem
cd ..
```

Save 10 runs of each version:

```bash
cd before
go test -bench=BenchmarkCount5xx -benchmem -count=10 > ../before.txt
cd ../after
go test -bench=BenchmarkCount5xx -benchmem -count=10 > ../after.txt
cd ..
```

Compare them:

```bash
benchstat before.txt after.txt
```
