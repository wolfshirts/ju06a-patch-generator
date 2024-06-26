package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"

	dict "github.com/wolfshirts/ju0ga-generator/internal/dictionary"
	"github.com/wolfshirts/ju0ga-generator/internal/ranges"
)

func main() {
	patchCount := flag.Int("patches", 1, "number of patches, 1-128")
	flag.Parse()
	if *patchCount > 128 {
		fmt.Println("Max is 128 patches.")
		return
	}
	dir := strings.TrimRight(dict.GetPatchName(1), " ")
	err := os.Mkdir(dir, 0770)
	if err != nil {
		panic("failed to create storage directory")
	}
	fmt.Println("Created directory " + dir + "...")
	for i := 0; i < *patchCount; i++ {
		p := generatePatch()
		fileName := fmt.Sprintf("JU06A_PATCH%d.PRM", i+1)
		st := convertAndNamePatch(p)
		f, err := os.Create(fmt.Sprintf("%s/%s", dir, fileName))
		if err != nil {
			panic(err)
		}
		f.Write([]byte(st))
		f.Chmod(0777)
		f.Close()
	}
}

func generatePatch() map[string]int {
	patchRanges := ranges.GetRangeMap()
	patch := map[string]int{}
	for k, v := range patchRanges {
		num := rand.Intn(v[1]+1-v[0]) + v[0]
		patch[k] = num
	}
	return patch
}

func convertAndNamePatch(p map[string]int) string {
	patchName := dict.GetPatchName(rand.Intn(2) + 1)
	fmt.Println("Creating... " + patchName)
	serialized := ""
	for _, v := range ranges.GetOrderedParams() {
		pad := 16 - len(v)
		fixed := v
		for i := 0; i < pad; i++ {
			fixed += " "
		}
		line := fmt.Sprintf("%s(%d);\r\n", fixed, p[v])
		serialized += line
	}
	serialized += fmt.Sprintf("PATCH_NAME(%s);", patchName)
	return serialized
}
