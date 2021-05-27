package main

import (
	"github.com/pkg/profile"
	"hash"
	"hash/fnv"
	_ "runtime/pprof"
)

//var b byte
//
//func httpGet() {
//	resp, err := http.Get("http://www.qq.com/?a=" + string(b))
//	if err != nil {
//		return
//	}
//
//	defer resp.Body.Close()
//	if true {
//		return
//	}
//}

var in []byte
var h hash.Hash64

func hashInt() uint64 {
	//h := fnv.New64a()
	h.Write(in)
	out := h.Sum64()
	return out
}

func mainfunc() {
	in = []byte("aaaaaaaaaa")
	h = fnv.New64a()
	for i := 0; i < 2000; i++ {
		in = append(in, 'a')
		hashInt()
		h.Reset()
	}
}

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	mainfunc()
}
