package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func ConvertBaseNum(input string, srcBase int, destBase int) (string, error) {
	decimalValue, err := strconv.ParseInt(input, srcBase, 64)
	if err != nil {
		return "", err
	}

	converted := ""
	for decimalValue > 0 {
		digit := decimalValue % int64(destBase)
		decimalValue /= int64(destBase)
		converted = strconv.FormatInt(digit, destBase) + converted
	}

	return strings.ToUpper(converted), nil
}

func SolveQuadraticEquation(coeffA, coeffB, coeffC float64) (complex128, complex128) {
	discriminant := coeffB*coeffB - 4*coeffA*coeffC

	if discriminant >= 0 {
		root1 := (-coeffB + math.Sqrt(discriminant)) / (2 * coeffA)
		root2 := (-coeffB - math.Sqrt(discriminant)) / (2 * coeffA)
		return complex(root1, 0), complex(root2, 0)
	}

	realPart := -coeffB / (2 * coeffA)
	imagPart := math.Sqrt(-discriminant) / (2 * coeffA)
	return complex(realPart, imagPart), complex(realPart, -imagPart)
}

func SortAbsValues(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})
}

func CombineArrays(array1, array2 []int) []int {
	result := append(array1, array2...)
	return result
}

func FindSubstringIndex(fullString, subString string) int {
	if len(subString) == 0 {
		return 0
	}
	if len(subString) > len(fullString) {
		return -1
	}
	for i := 0; i <= len(fullString)-len(subString); i++ {
		if strings.HasPrefix(fullString[i:], subString) {
			return i
		}
	}
	return -1
}

func PerformCalculation(val1, val2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return val1 + val2, nil
	case "-":
		return val1 - val2, nil
	case "*":
		return val1 * val2, nil
	case "/":
		if val2 == 0 {
			return 0, errors.New("division by zero")
		}
		return val1 / val2, nil
	case "^":
		return math.Pow(val1, val2), nil
	case "%":
		if val2 == 0 {
			return 0, errors.New("modulo by zero")
		}
		return float64(int(val1) % int(val2)), nil
	default:
		return 0, errors.New("unsupported operator")
	}
}

func CheckPalindrome(text string) bool {
	normalized := ""
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			normalized += strings.ToLower(string(char))
		}
	}
	for i, j := 0, len(normalized)-1; i < j; i, j = i+1, j-1 {
		if normalized[i] != normalized[j] {
			return false
		}
	}
	return true
}

func CheckSegmentOverlap(seg1, seg2, seg3 [2]int) bool {
	leftBound := maxInt(seg1[0], seg2[0], seg3[0])
	rightBound := minInt(seg1[1], seg2[1], seg3[1])
	return leftBound <= rightBound
}

func maxInt(values ...int) int {
	maxVal := values[0]
	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func minInt(values ...int) int {
	minVal := values[0]
	for _, val := range values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func NormalizeAndSplit(sentence string) []string {
	var cleaned strings.Builder
	for _, ch := range sentence {
		if unicode.IsLetter(ch) || unicode.IsSpace(ch) {
			cleaned.WriteRune(ch)
		}
	}
	return strings.Fields(cleaned.String())
}

func FindLargestWord(sentence string) string {
	words := NormalizeAndSplit(sentence)
	largest := ""
	for _, word := range words {
		if len(word) > len(largest) {
			largest = word
		}
	}
	return largest
}

func IsArmstrong(num int) bool {
	original := num
	sum := 0
	numDigits := int(math.Log10(float64(num))) + 1
	for num > 0 {
		digit := num % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		num /= 10
	}
	return sum == original
}

func ArmstrongNumbersInRange(start, end int) []int {
	results := []int{}
	for i := start; i <= end; i++ {
		if IsArmstrong(i) {
			results = append(results, i)
		}
	}
	return results
}

func ReverseText(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ComputeGCD(num1, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
func main() {
	// Пример использования ConvertBaseNum
	result, err := ConvertBaseNum("1010", 2, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("ConvertBaseNum result:", result)
	}

	// Пример решения квадратного уравнения
	root1, root2 := SolveQuadraticEquation(1, -3, 2)
	fmt.Printf("SolveQuadraticEquation roots: %v, %v\n", root1, root2)

	// Пример проверки палиндрома
	isPalindrome := CheckPalindrome("A man, a plan, a canal: Panama")
	fmt.Println("CheckPalindrome result:", isPalindrome)

	// Пример нахождения чисел Армстронга в диапазоне
	armstrongNums := ArmstrongNumbersInRange(100, 500)
	fmt.Println("ArmstrongNumbersInRange:", armstrongNums)

	// Пример обратного текста
	reversed := ReverseText("hello")
	fmt.Println("ReverseText result:", reversed)
}
