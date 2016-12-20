package main

import (
	"fmt"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var dmp = diffmatchpatch.New()

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func createPatch(lastDoc, currDoc string) (_result string) {
	t0 := time.Now()

	diffs := dmp.DiffMain(lastDoc, currDoc, false)
	diffs = dmp.DiffCleanupSemantic(diffs)

	patch := dmp.PatchMake(diffs)

	result := dmp.PatchToText(patch)

	t1 := time.Now()
	fmt.Print(t1.Sub(t0).Nanoseconds())
	fmt.Println("ns")

	return result
}

func main() {
	fmt.Println(createPatch(text1, text2))
}
