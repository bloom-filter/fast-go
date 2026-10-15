# Fast Go

Example code for the course "Fast Go: Measure, Profile and Optimize Go Programs".

Every lesson follows the same loop: measure, find the problem, change the code,
measure again, and compare.

## Requirements

- Go 1.24 or later (the benchmarks use `b.Loop()`)
- benchstat:

  ```bash
  go install golang.org/x/perf/cmd/benchstat@latest
  ```

## Lessons

| Folder | Lesson |
|---|---|
| [ep01](ep01) | Benchmark: measure before you optimize |

## About the numbers

Benchmark results depend on your machine. The numbers in the lessons were measured
on an AMD Ryzen 9 9950X (linux/amd64, Go 1.27.1). Compare before and after
on the same machine, with the same input, and run each benchmark several times.

## License

MIT
