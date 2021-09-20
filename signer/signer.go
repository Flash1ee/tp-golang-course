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

	mu.Lock()
	md5 := DataSignerMd5(data)
	mu.Unlock()

	//dataChanCrc32 := make(chan string)
	//go Crc32Parallel(data, dataChanCrc32)
	crc32Md5 := DataSignerCrc32(md5)
	//crc32 := <- dataChanCrc32
	crc32 := DataSignerCrc32(data)
	out <- crc32 + "~" + crc32Md5

	wg.Done()
}
var Crc32Parallel = func(data string, out chan string) {
	out <- DataSignerCrc32(data)
}

var MultiHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}
	for el := range in {
		wg.Add(1)
		go ProcessMultiHash(el, out, wg, mu)

	}
	wg.Wait()
}

var ProcessMultiHash = func(el interface{}, out chan interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	const th = 6

	buf := make([]string, th)
	dataChan := make(chan string)
	for i := 0; i < th; i++ {
		data := strconv.Itoa(i) + el.(string)

		go Crc32Parallel(data, dataChan)
		//crc32 := DataSignerCrc32(data)
		buf[i] = <-dataChan
	}
	out <- strings.Join(buf, "")
	wg.Done()
}

var CombineResults = func(in, out chan interface{}) {
	var data []string
	for i := range in {
		data = append(data, i.(string))
	}
	sort.Strings(data)

	out <- strings.Join(data, "_")
}

func main() {}
