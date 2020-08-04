package main

import "fmt"

func main() {
	var (
		x int
		y int
		z int
	)
	fmt.Println("Menu Calculator")
	fmt.Println("1. Kali")
	fmt.Println("2. Bagi")
	fmt.Println("3. Tambah")
	fmt.Println("4. Kurang")
	fmt.Println("5. Pangkat")
	fmt.Println("6. Akar")
	fmt.Println("7. Luas Persegi")
	fmt.Println("8. Luas Lingkaran")
	fmt.Println("9. Volume Balok")
	fmt.Println("10.Volume Tabung")
	fmt.Println("11. Volume Prisma")
	fmt.Println("Pilih [1-10]")
	fmt.Scanf("%d", &z)

	switch z {
	case 1:
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		fmt.Println(Tambah(x, y))

	}

}

func Kali(a, b int) int {
	return a * b
}

func Bagi(a, b int) int {
	return a / b
}

func Tambah(a, b int) int {
	return a + b
}

func Kurang(a, b int) int {
	return a - b
}
