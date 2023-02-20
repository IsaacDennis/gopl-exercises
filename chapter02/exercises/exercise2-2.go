package main

import (
	"fmt"
	"os"
	"strconv"
)

func CToF(c float64) float64 { return (c*9/5 + 32) }
func FToC(f float64) float64 { return ((f - 32) * 5 / 9) }

func KgToLbs(kg float64) float64  { return kg * 2.2046 }
func LbsToKg(lbs float64) float64 { return lbs / 2.2046 }

func MToFt(m float64) float64  { return m * 3.28 }
func FtToM(ft float64) float64 { return ft / 3.28 }
func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("%s is not a valid number.\n", arg)
			os.Exit(1)
		}
		fmt.Println("Temperature units:")
		fmt.Printf("%.2f°C = %.2f°F, %.2f°F = %.2f°C\n", num, CToF(num), num, FToC(num))

		// Mass and distance units can only be positive
		if num > 0 {
			fmt.Println("-----------------")
			fmt.Println("Mass units:")
			fmt.Printf("%.2f kg = %.2f lbs, %.2f lbs = %.2f kg\n", num, KgToLbs(num), num, LbsToKg(num))
			fmt.Println("-----------------")
			fmt.Println("Distance units:")
			fmt.Printf("%.2f m = %.2f ft, %.2f ft = %.2f m\n", num, MToFt(num), num, FtToM(num))

		}
		fmt.Println("=================")
	}
}
