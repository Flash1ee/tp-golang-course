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
		go BeforeJob(job, in, out, wg)

		in = out
	}
	wg.Wait()
}

var BeforeJob = func(job job, in, out chan interface{}, wg *sync.WaitGroup) {
	job(in, out)
	close(out)
	wg.Done()
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}
	for el := range in {
		wg.Add(1)
		go ProcessSingleHash(el, out, wg, mu)
	}
	wg.Wait()
}

var ProcessSingleHash = func(el interface{}, out chan interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
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
		go ProcessMultiHash(el, out, wg)

	}
	wg.Wait()
}

var ProcessMultiHash = func(el interface{}, out chan interface{}, wg *sync.WaitGroup) {
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

func main() {}
