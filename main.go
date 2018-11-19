/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 19-11-2018
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

var daysOfMonth = map[string]int{
	"Mehr": 30,
	"Aban": 30,
}

func main() {
	excelFileName := "Aban-1397.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	sheet := xlFile.Sheets[0]
	for i := 6; i < 6+daysOfMonth["Aban"]; i++ {
		row := sheet.Rows[i]

		// the day of the month
		day, err := strconv.Atoi(row.Cells[19].Value)
		if err != nil {
			panic(err)
		}

		// the maximum, minimum and average temperature
		minimumTemperature, err := strconv.ParseFloat(row.Cells[18].Value, 64)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Day: %d, MinTemp: %f\n", day, minimumTemperature)

	}
}
