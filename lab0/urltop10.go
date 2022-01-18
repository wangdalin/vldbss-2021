package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// K Top K
const K = 10

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs

	// round 1: both map phase and reduce phase do the url count
	args = append(args, RoundArgs{
		MapFunc:    URLTop10CountMap,
		ReduceFunc: URLTop10CountReduce,
		NReduce:    nWorkers,
	})

	// round 2: both map phase and reduce do sort and topK filter
	args = append(args, RoundArgs{
		MapFunc:    URLTop10SortFilterMap,
		ReduceFunc: URLTop10SortFilterReduce,
		NReduce:    1,
	})

	return args
}

// URLTop10CountMap count url and combine same key to reduce IO
func URLTop10CountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	kvMap := make(map[string]int)
	for _, url := range lines {
		url = strings.TrimSpace(url)
		if len(url) == 0 {
			continue
		}
		if _, exist := kvMap[url]; !exist {
			kvMap[url] = 0
		}
		kvMap[url] = kvMap[url] + 1
	}

	kvs := make([]KeyValue, 0)
	for k, v := range kvMap {
		kvs = append(kvs, KeyValue{Key: k, Value: strconv.Itoa(v)})
	}
	return kvs
}

// URLTop10CountReduce calculate key's total count
func URLTop10CountReduce(key string, values []string) string {
	count := 0
	for _, v := range values {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		count += vInt
	}

	return fmt.Sprintf("%s: %d\n", key, count)
}

// URLTop10SortFilterMap do sort and topK filter
func URLTop10SortFilterMap(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	cnts := make(map[string]int)
	for _, v := range lines {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, ": ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts[tmp[0]] = n
	}

	us, cs := TopN(cnts, K)
	kvs := make([]KeyValue, 0, K)
	for i := range us {
		kvs = append(kvs, KeyValue{Key: "", Value: fmt.Sprintf("%s: %d", us[i], cs[i])})
	}

	return kvs
}

// URLTop10SortFilterReduce do sort and topK filter
func URLTop10SortFilterReduce(key string, values []string) string {
	kvMap := make(map[string]int, len(values))
	for _, value := range values {
		tmp := strings.Split(value, ": ")
		url := strings.TrimSpace(tmp[0])
		if len(url) == 0 {
			continue
		}
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		if _, exist := kvMap[url]; !exist {
			kvMap[url] = 0
		}
		kvMap[url] = kvMap[url] + n
	}

	us, cs := TopN(kvMap, K)
	buf := new(bytes.Buffer)
	for i := range us {
		fmt.Fprintf(buf, "%s: %d\n", us[i], cs[i])
	}
	return buf.String()
}
