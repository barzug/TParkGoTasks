package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func ExecutePipeline(inJobs ...job) {
	in := make(chan interface{}, MaxInputDataLen)
	defer close(in)

	for _, job := range inJobs {
		out := make(chan interface{}, MaxInputDataLen)
		go job(in, out)
		close(out)
		in = out
	}
	return
}

func SingleHash(in, out chan interface{}) {
	muteForMD5 := &sync.Mutex{}
	waitElems := &sync.WaitGroup{}

	for item := range in {
		waitElems.Add(1)

		intItem, ok := item.(int)
		if !ok {
			panic("cant convert data to string")
		}

		itemString := fmt.Sprint(intItem)

		go singleHashForOneElem(out, waitElems, muteForMD5, itemString)
	}
	waitElems.Wait()
}

func MultiHash(in, out chan interface{}) {
	waitElems := &sync.WaitGroup{}

	for item := range in {
		waitElems.Add(1)

		stringItem, ok := item.(string)
		if !ok {
			panic("cant convert data to string")
		}
		go multiHashForOneElem(out, waitElems, stringItem)
	}

	waitElems.Wait()
}

func CombineResults(in, out chan interface{}) {
	results := make([]string, 0)

	for item := range in {

		stringItem, ok := item.(string)
		if !ok {
			panic("cant convert data to string")
		}

		results = append(results, stringItem)
	}
	sort.Strings(results)
	out <- strings.Join(results, "_")
}

func multiHashForOneElem(out chan interface{}, waitElems *sync.WaitGroup, data string) {
	defer waitElems.Done()

	waitMultiHash := &sync.WaitGroup{}
	result := make([]string, 6, 6)

	for th := 0; th <= 5; th++ {
		waitMultiHash.Add(1)

		go func(th int) {
			defer waitMultiHash.Done()
			thString := fmt.Sprint(th)
			crc32Result := DataSignerCrc32(thString + data)
			result[th] = crc32Result
		}(th)
	}
	waitMultiHash.Wait()
	out <- strings.Join(result, "")
}

func singleHashForOneElem(out chan interface{}, waitElems *sync.WaitGroup, muteForMD5 *sync.Mutex, data string) {
	defer waitElems.Done()

	waitTwoCrcResult := &sync.WaitGroup{}
	waitTwoCrcResult.Add(1)
	crc32Data := ""
	crc32md5 := ""

	go func() {
		defer waitTwoCrcResult.Done()
		muteForMD5.Lock()
		md5Data := DataSignerMd5(data)
		muteForMD5.Unlock()
		crc32md5 = DataSignerCrc32(md5Data)
	}()
	crc32Data = DataSignerCrc32(data)

	waitTwoCrcResult.Wait()
	result := crc32Data + "~" + crc32md5
	out <- result
}
