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
	f.SetCellValue("Kursübersicht", "B1", "#SuS")

	f.SetCellValue("Wochenübersicht", "A1", "Woche 1")
	f.SetCellValue("Wochenübersicht", "B1", "Woche 2")

	var o = schedule.Classes[0].Students
	var t = schedule.Classes[1].Students

	for i := 0; i < biggest(len(o), len(t)); i++ {
		var one string
		var two string

		if len(o) > i {
			one = o[i].ID
		}

		if len(t) > i {
			two = t[i].ID
		}

		f.SetCellValue("Wochenübersicht", "A"+fmt.Sprint(i+2), one)
		f.SetCellValue("Wochenübersicht", "B"+fmt.Sprint(i+2), two)
	}

	courses := ClassesToCourses(schedule.Classes)

	for i, course := range courses {
		f.SetCellValue("Kursübersicht", "A"+fmt.Sprint(i+2), course.Name)
		f.SetCellValue("Kursübersicht", "B"+fmt.Sprint(i+2), course.CountStudents())

		for j, student := range course.Students {
			f.SetCellValue("Kursübersicht", toChar(j+3)+fmt.Sprint(i+2), student.ID)
		}
	}

	return f.SaveAs(outputPath)
}

// WriteScheduleToStdOut writes the result as table to stdout
func WriteScheduleToStdOut(schedule *Schedule) {

	classTable := tablewriter.NewWriter(os.Stdout)
	classTable.SetHeader([]string{"Woche 1", "Woche 2"})

	var o = schedule.Classes[0].Students
	var t = schedule.Classes[1].Students

	for i := 0; i < biggest(len(o), len(t)); i++ {
		var one string
		var two string

		if len(o) > i {
			one = o[i].ID
		}

		if len(t) > i {
			two = t[i].ID
		}

		classTable.Append([]string{one, two})
	}

	courseTable := tablewriter.NewWriter(os.Stdout)
	courseTable.SetHeader([]string{"Kursname", "#SuS"})

	courses := ClassesToCourses(schedule.Classes)

	for _, course := range courses {
		var studentIDs []string
		for _, student := range course.Students {
			studentIDs = append(studentIDs, student.ID)
		}
		row := []string{course.Name, strconv.Itoa(course.CountStudents())}
		for _, id := range studentIDs {
			row = append(row, id)
		}
		courseTable.Append(row)
	}

	classTable.Render()
	courseTable.Render()
}
