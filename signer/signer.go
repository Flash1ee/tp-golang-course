package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	singleHashSeparator    = "~"
	multiHashSeparator     = ""
	combineResultSeparator = "_"
	th                     = 6
)

var ExecutePipeline = func(jobs ...job) {
	wg := &sync.WaitGroup{}

	in := make(chan interface{})

	for _, job := range jobs {
		wg.Add(1)

		out := make(chan interface{})
		go BeforeJob(wg, job, in, out)

		in = out
	}
	wg.Wait()
}

var BeforeJob = func(wg *sync.WaitGroup, job job, in, out chan interface{}) {
	job(in, out)
	close(out)
	wg.Done()
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}
	for el := range in {
		wg.Add(1)
		go ProcessSingleHash(wg, mu, el, out)
	}
	wg.Wait()
}

var ProcessSingleHash = func(wg *sync.WaitGroup, mu *sync.Mutex, el interface{}, out chan interface{}) {
	data := strconv.Itoa((el).(int))

	dataChan := make(chan string)
	go Crc32Parallel(data, dataChan)

	mu.Lock()
	md5 := DataSignerMd5(data)
	mu.Unlock()

	crc32Md5 := DataSignerCrc32(md5)

	crc32 := <-dataChan
	out <- crc32 + singleHashSeparator + crc32Md5

	wg.Done()
}
var Crc32Parallel = func(data string, out chan string) {
	out <- DataSignerCrc32(data)
}

var MultiHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for el := range in {
		wg.Add(1)
		go ProcessMultiHash(wg, el, out)

	}
	wg.Wait()
}

var ProcessMultiHash = func(wg *sync.WaitGroup, el interface{}, out chan interface{}) {
	InternalWg := &sync.WaitGroup{}
	buf := make([]string, th)

	for i := 0; i < th; i++ {
		InternalWg.Add(1)
		data := strconv.Itoa(i) + el.(string)

		go func(pos int) {
			buf[pos] = DataSignerCrc32(data)
			InternalWg.Done()
		}(i)
	}
	InternalWg.Wait()
	out <- strings.Join(buf, multiHashSeparator)
	wg.Done()
}
var CombineResults = func(in, out chan interface{}) {
	var data []string
	for i := range in {
		data = append(data, i.(string))
	}
	sort.Strings(data)

	out <- strings.Join(data, combineResultSeparator)
}
