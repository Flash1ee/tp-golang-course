package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

var ExecutePipeline = func(jobs ...job) {
	wg := &sync.WaitGroup{}

	in := make(chan interface{})

	for _, job := range jobs {
		wg.Add(1)

		out := make(chan interface{})
		go BeforeJob(job, in, out, wg)

		in = out
	}
	wg.Wait()
}

var BeforeJob = func(job job, in, out chan interface{}, wg *sync.WaitGroup) {
	job(in, out)
	wg.Done()
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for el := range in {
		wg.Add(1)
		ProcessSingleHash(el, out, wg)
	}
}

var ProcessSingleHash = func(el interface{}, out chan interface{}, wg *sync.WaitGroup) {
	data := strconv.Itoa((el).(int))

	crc32 := DataSignerCrc32(data)
	crc32Md5 := DataSignerCrc32(DataSignerMd5(data))

	out <- crc32 + "~" + crc32Md5

	wg.Done()
}
var MultiHash = func(in, out chan interface{}) {}

var CombineResults = func(in, out chan interface{}) {
	var data []string
	for i := range in {
		data = append(data, i.(string))
	}
	sort.Strings(data)

	out <- strings.Join(data, "_")
}

func main() {}
