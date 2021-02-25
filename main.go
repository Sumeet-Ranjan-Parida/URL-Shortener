package main

import (
	"fmt"
	"time"

	"github.com/speps/go-hashids"
)

func main() {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)

	now := time.Now()
	e, _ := h.Encode([]int{int(now.Unix())})

	d, _ := h.DecodeWithError(e)

	fmt.Println(e)
	fmt.Println(d)
}
