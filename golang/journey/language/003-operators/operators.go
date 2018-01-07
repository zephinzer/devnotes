package main

import "fmt"

func main() {
	a, b, c, d := 1, 2, 2, 3
	fmt.Printf("premise: a = %d, b = %d, c = %d, d = %d\n", a, b, c, d)
	fmt.Printf("  a < b           == %t\n", a < b)
	fmt.Printf("  a > b           == %t\n", a > b)
	fmt.Printf("  b <= c          == %t\n", b <= c)
	fmt.Printf("  b >= c          == %t\n", b >= c)
	fmt.Printf("  b != c          == %t\n", b != c)
	fmt.Printf("  a == b          == %t\n", a == b)
	fmt.Printf("  b == c          == %t\n", b == c)
	d++
	fmt.Printf("  d post inc      == %d\n", d)
	d--
	fmt.Printf("  d post dev      == %d\n", d)
	fmt.Printf("  d + 1           == %d\n", d+1)
	d += 1
	fmt.Printf("  d after += 1    == %d\n", d)
	fmt.Printf("  d - 1           == %d\n", d-1)
	d -= 1
	fmt.Printf("  d after -= 1    == %d\n", d)
	d = 4
	fmt.Printf("  d = %d\n", d)
	fmt.Printf("    d      = %4.b\n", d)
	fmt.Printf("    15     = %4.b\n", 15)
	fmt.Printf("    d & 15 = %4.b\n", d&15)
	d &= 15
	fmt.Printf("  d after &= 15   == %d\n", d)
	d = 3
	fmt.Printf("  d = %d\n", d)
	fmt.Printf("    d      = %4.b\n", d)
	fmt.Printf("    13     = %4.b\n", 13)
	fmt.Printf("    d & 13 = %4.b\n", d&13)
	d &= 13
	fmt.Printf("  d after &= 13   == %d\n", d)
	d = 3
	fmt.Printf("  d = %d\n", d)
	fmt.Printf("    d      = %4.b\n", d)
	fmt.Printf("    15     = %4.b\n", 15)
	fmt.Printf("    d & 15 = %4.b\n", d|15)
	d |= 15
	fmt.Printf("  d after |= 15   == %d\n", d)
	d = 3
	fmt.Printf("  d = %d\n", d)
	fmt.Printf("    d      = %4.b\n", d)
	fmt.Printf("    12     = %4.b\n", 12)
	fmt.Printf("    d | 12 = %4.b\n", d|12)
	d |= 12
	fmt.Printf("  d after |= 12   == %d\n", d)

	d = 4
	fmt.Printf("  d = %d = %b\n", d, d)
	fmt.Printf("  d >> 1          == %d == %b\n", d>>1, d>>1)
	fmt.Printf("  d << 1          == %d == %b\n", d<<1, d<<1)
}
