package convert

import (
	"math"
	"strconv"
	"strings"
)

//* Convert ft and inches into metric float64
func ConvertFtIntoM(ft string, in string) float64 {
	ftNumber, _ := strconv.ParseFloat(ft, 64)
	inNumber, _ := strconv.ParseFloat(in, 64)
	totalInches := ftNumber*12 + inNumber
	return float64(totalInches) / 39.37
}

//* Convert lbs string into a float in metric
func LbsToKg(lbs string) float64 {
	lbsNum, _ := strconv.ParseFloat(lbs, 64)
	return lbsNum / 2.205
}

//* Get BMI from Weight and Height
func ConvertBmiFromKgAndM(m float64, kg float64) float64 {
	return kg / math.Pow(m, 2)
}

//* Trim the '\n' that is listened for with bufio
func Trim(value string) string {
	return strings.TrimSpace(value)
}
