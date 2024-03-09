# Merge Intervals

This package merges overlapping number intervals.

## Usage

```bash
go run . <input intervals>
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
BenchmarkParseInputArgs-12    1254063      977.8 ns/op    1136 B/op    21 allocs/op
BenchmarkMerge-12             9526440      120.7 ns/op     312 B/op     6 allocs/op
BenchmarkMerge1000-12         1798074      661.0 ns/op      88 B/op     3 allocs/op
BenchmarkMerge10000-12         192378     6181.0 ns/op      88 B/op     3 allocs/op
BenchmarkMerge100000-12         19362    61873.0 ns/op      88 B/op     3 allocs/op
PASS
ok      github.com/LucaBernstein/merge-intervals-go     8.281s
```

## Considerations and Assumptions

* Order of intervals passed can be unordered, but interval definition itself is ordered (`[ <start> , <end> ]`)
* Interval start and end must be whole numbers (`ℤ`)
