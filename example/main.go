package main

import (
	"fmt"
	"github.com/ysugimoto/go-kakasi"
)

func main() {
	ret, err := kakasi.Transform("群馬県の有名なご当地料理", kakasi.WithKanji(kakasi.ASCII))
	fmt.Println(ret, err)
}
