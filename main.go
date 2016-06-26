package main

import (
	"agg"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	a := &agg.Agg{Out: os.Stdout}

	err := agg.ParseFlag(a)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		ts := fields[0]

		if a.Type == agg.Count {
			for _, v := range fields[1:] {
				a.PushStr(ts, v)
			}
		} else {
			for _, v := range fields[1:] {
				f, err := strconv.ParseFloat(v, 64)

				if err != nil {
					log.Fatal(err)
				}

				a.PushNum(ts, f)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
