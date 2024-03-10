# Merge Intervals

This package merges overlapping number intervals.

## Usage

```bash
go run . <input intervals>
```

Larger datasets of intervals to merge can be provided via file and env var:

```bash
INTERVALS_FILE=<filename> go run .
```

## Example

The input `[25,30] [2,19] [14,23] [4,8]` should yield the output `[2,23] [25,30]`.

```bash
# Input
go run . "[25,30] [2,19] [14,23] [4,8]"

# Output
[2,23] [25,30]
```

## Performance and Footprint

This package includes benchmark tests that show how different parts behave with different loads (input parsing, interval merge).

Sample output of `go test -bench=. -run=XXX -benchmem`:

```
pkg: github.com/LucaBernstein/merge-intervals-go
BenchmarkParseInputArgs-12           1230621        965.1 ns/op    1136 B/op    21 allocs/op
BenchmarkMerge-12                    9905893        120.6 ns/op     312 B/op     6 allocs/op
BenchmarkMergeOverlap1000-12         1868359        643.1 ns/op      88 B/op     3 allocs/op
BenchmarkMergeOverlap10000-12         199348       5978.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeOverlap100000-12         20149      60515.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeOverlap1000000-12         1995     613224.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeDistinct1000-12        1847392        643.8 ns/op      88 B/op     3 allocs/op
BenchmarkMergeDistinct10000-12        201364       6022.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeDistinct100000-12        19869      59441.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeDistinct1000000-12        1990     602862.0 ns/op      88 B/op     3 allocs/op
BenchmarkMergeDistinct10000000-12        196    6131859.0 ns/op      88 B/op     3 allocs/op
PASS
ok      github.com/LucaBernstein/merge-intervals-go     17.673s
```

In the current implementation, the calculations are executed in a single thread only.
Hence, for a larger amount of intervals provided, the runtime increases linearly.
The memory footprint also increases linearly for storing the intervals internally.

For example, during the `BenchmarkMergeDistinct1000000` benchmark run, a memory allocation of around `40 MB` during execution has been observed.
Likewise, during the `BenchmarkMergeDistinct10000000` benchmark run, the memory allocation observed has been around `400 MB`.

## Considerations and Assumptions

* Order of intervals passed can be unordered, but interval definition itself is ordered (`[ <start> , <end> ]` with `start <= end`)
* Interval start and end must be whole numbers (`ℤ`)
