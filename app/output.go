package app

import (
	"fmt"
	"os"
	"strconv"

	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
	tablewriter "github.com/olekukonko/tablewriter"
)

func toChar(i int) string {
	return string(rune('A' - 1 + i))
}

// biggest returns the largest of two integers
func biggest(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// WriteScheduleAsFile writes the result into an excel file
func WriteScheduleAsFile(schedule *Schedule, outputPath string) error {
	f := excel.NewFile()

	f.NewSheet("Wochenübersicht")
	f.NewSheet("Kursübersicht")
	f.DeleteSheet("Sheet1")

	f.SetCellValue("Kursübersicht", "A1", "Kursname")
	f.SetCellValue("Kursübersicht", "B1", "Kursgröße")

	f.SetCellValue("Wochenübersicht", "A1", "Woche 1")
	f.SetCellValue("Wochenübersicht", "B1", "Woche 2")

	var o = schedule.Weeks[0].Students
	var t = schedule.Weeks[1].Students

	for i := 0; i < biggest(len(o), len(t)); i++ {
		var one string
		var two string

		if len(o) > i {
			one = o[i].ID
		}

		if len(t) > i {
			two = t[i].ID
		}

		f.SetCellValue("Wochenübersicht", fmt.Sprint("A", i+2), one)
		f.SetCellValue("Wochenübersicht", fmt.Sprint("B", i+2), two)
	}

	courses := WeeksToCourses(schedule.Weeks)

	for i, course := range courses {
		f.SetCellValue("Kursübersicht", fmt.Sprint("A", i+2), course.Name)
		f.SetCellValue("Kursübersicht", fmt.Sprint("B", i+2), course.CountStudents())

		for j, student := range course.Students {
			f.SetCellValue("Kursübersicht", fmt.Sprint(toChar(j+3), i+2), student.ID)
		}
	}

	return f.SaveAs(outputPath)
}

// WriteScheduleToStdOut writes the result as table to stdout
func WriteScheduleToStdOut(schedule *Schedule) {

	weeksTable := tablewriter.NewWriter(os.Stdout)
	weeksTable.SetHeader([]string{"Woche 1", "Woche 2"})

	var o = schedule.Weeks[0].Students
	var t = schedule.Weeks[1].Students

	for i := 0; i < biggest(len(o), len(t)); i++ {
		var one string
		var two string

		if len(o) > i {
			one = o[i].ID
		}

		if len(t) > i {
			two = t[i].ID
		}

		weeksTable.Append([]string{one, two})
	}

	coursesTable := tablewriter.NewWriter(os.Stdout)
	coursesTable.SetHeader([]string{"Kursname", "Kursgröße"})

	courses := WeeksToCourses(schedule.Weeks)

	for _, course := range courses {
		var studentIDs []string
		for _, student := range course.Students {
			studentIDs = append(studentIDs, student.ID)
		}

		row := []string{course.Name, strconv.Itoa(course.CountStudents())}
		for _, id := range studentIDs {
			row = append(row, id)
		}

		coursesTable.Append(row)
	}

	weeksTable.Render()
	coursesTable.Render()
}
