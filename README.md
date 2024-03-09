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

> TODO...

## Considerations and Assumptions

* Order of intervals passed can be unordered, but interval definition itself is ordered (`[ <start> , <end> ]`)
* Interval start and end must be whole numbers (`ℤ`)
