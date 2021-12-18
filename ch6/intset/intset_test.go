// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	x.Remove(9)
	fmt.Println(x.Has(9)) // "false"
	fmt.Println(x.Len())  // 3

	x.Clear()
	fmt.Println(x.Len()) // 0

	var z *IntSet
	z = y.Copy()
	fmt.Println(z.String()) // "{9 42}"

	z.AddAll(1, 2, 3)
	fmt.Println(z.String()) // {"1 2 3 9 42"}

	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
	// false
	// 3
	// 0
	// {9 42}
	// {1 2 3 9 42}
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Example_three() {
	var x, y IntSet
	x.AddAll(1, 2, 3)
	y.AddAll(2, 5, 6)
	var z *IntSet
	z = x.IntersectWith(&y)
	fmt.Println(z.String()) // {"2"}

	z = x.DifferenceWith(&y)
	fmt.Println(z.String()) // {"1 3"}

	z = x.SymmetricDifference(&y)
	fmt.Println(z.String()) // {"1 3 5 6"}

	fmt.Println(z.Elems())
	// Output:
	// {2}
	// {1 3}
	// {1 3 5 6}
	// [1 3 5 6]

}
