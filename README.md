# agg

Aggregate log records by time.

[![Build Status](https://travis-ci.org/winebarrel/agg.svg?branch=master)](https://travis-ci.org/winebarrel/agg)

## Usage

```
Usage of agg:
  -t string
      count,sum,avg,max,min (default "count")
```

```
$ cat access.log
12:10  FOO  100
12:10  FOO  200
12:10  BAR  300
12:11  FOO  200
12:11  BAR  300
12:11  BAR  400
12:12  ZOO  500

$ awk '{print $1, $2}' access.log | agg
12:10 BAR:1 FOO:2
12:11 BAR:2 FOO:1
12:12 ZOO:1

$ awk '{sub(/[0-9]$/, "0", $1); print $1, $2}' access.log | agg
12:10 BAR:3 FOO:3 ZOO:1

$ awk '{print $1, $3}' access.log | agg -t sum
12:10 600.000000
12:11 900.000000
12:12 500.000000

$ awk '{print $1, $3}' access.log | agg -t avg
12:10 200.000000
12:11 300.000000
12:12 500.000000

$ awk '{print $1, $3}' access.log | agg -t max
12:10 300.000000
12:11 400.000000
12:12 500.000000

$ awk '{print $1, $3}' access.log | agg -t min
12:10 100.000000
12:11 200.000000
12:12 500.000000
```
