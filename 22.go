package main

import (
	"fmt"
	"log"
	"math/big"
)

// struct of object
type Big struct {
	x      string
	y      string
	result string
}

// create BigInt from string
func (selfbig *Big) NewBig() (vals []*big.Int) {
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)
	a.SetString(selfbig.x, 10)
	b.SetString(selfbig.y, 10)
	res.SetString(selfbig.result, 10)
	vals = append(vals, a, b, res)
	return
}

// do all actions
func Acts(x, y string) (vls []string) {
	// get object
	vals := Big{
		x:      x,
		y:      y,
		result: "",
	}
	// get big values from object
	v := vals.NewBig()
	vls = append(vls,
		"\n(+)", v[2].Add(v[0], v[1]).String(),
		"\n(*)", v[2].Mul(v[0], v[1]).String(),
		"\n(-)", v[2].Sub(v[0], v[1]).String(),
	)
	if v[1].Int64() != 0 {
		vls = append(vls, "\n(/)", v[2].Div(v[0], v[1]).String())
		return
	}
	return
}

func getFromUser() {
	var x, y string
	for {
		fmt.Scanln(&x, &y)
		log.Println(Acts(x, y))
	}
}
