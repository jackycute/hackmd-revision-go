package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func createPatch(lastDoc, currDoc string) (_result string) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	m0 := mem.HeapAlloc
	t0 := time.Now()

	diffs := dmp.DiffMain(lastDoc, currDoc, false)
	diffs = dmp.DiffCleanupSemantic(diffs)

	patch := dmp.PatchMake(diffs)

	result := dmp.PatchToText(patch)

	t1 := time.Now()
	fmt.Print(t1.Sub(t0).Nanoseconds())
	fmt.Println(" ns")
	runtime.ReadMemStats(&mem)
	m1 := mem.HeapAlloc
	fmt.Print((m1 - m0) / 1024)
	fmt.Println(" KB")

	return result
}

func main() {
	createPatch(text1, text2)
	//fmt.Println(createPatch(text1, text2))
	d1, err := ioutil.ReadFile("speedtest1.txt")
	if err != nil {
		panic(err)
	}
	d2, err := ioutil.ReadFile("speedtest2.txt")
	if err != nil {
		panic(err)
	}
	createPatch(string(d1), string(d2))
	//fmt.Println(createPatch(string(d1), string(d2)))
}
