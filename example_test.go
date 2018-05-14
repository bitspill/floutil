package floutil_test

import (
	"fmt"
	"math"

	"github.com/bitspill/floutil"
)

func ExampleAmount() {

	a := floutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = floutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = floutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 FLO
	// 100,000,000 Satoshis: 1 FLO
	// 100,000 Satoshis: 0.001 FLO
}

func ExampleNewAmount() {
	amountOne, err := floutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := floutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := floutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := floutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 FLO
	// 0.01234567 FLO
	// 0 FLO
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := floutil.Amount(44433322211100)

	fmt.Println("Satoshi to kFLO:", amount.Format(floutil.AmountKiloFLO))
	fmt.Println("Satoshi to FLO:", amount)
	fmt.Println("Satoshi to MilliFLO:", amount.Format(floutil.AmountMilliFLO))
	fmt.Println("Satoshi to MicroFLO:", amount.Format(floutil.AmountMicroFLO))
	fmt.Println("Satoshi to Satoshi:", amount.Format(floutil.AmountSatoshi))

	// Output:
	// Satoshi to kFLO: 444.333222111 kFLO
	// Satoshi to FLO: 444333.222111 FLO
	// Satoshi to MilliFLO: 444333222.111 mFLO
	// Satoshi to MicroFLO: 444333222111 μFLO
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
