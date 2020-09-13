package app

import (
	"os"
	"strconv"

	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
	tablewriter "github.com/olekukonko/tablewriter"
)

// WriteCoursesAsTable prints out the final result
// as a readable table
func WriteCoursesAsTable(courses []*Course) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Name des Kurses", "Anzahl der Schüler"})

	for _, course := range courses {
		studentCount := strconv.Itoa(course.CountStudents())
		row := []string{course.Name, studentCount}
		table.Append(row)
	}
	table.Render()
}

// WriteCoursesAsFile writes the result into an excel file
func WriteCoursesAsFile(courses []*Course, outputPath string) error {
	f := excel.NewFile()
	f.SetCellValue("Sheet1", "A1", "Name des Kurses")
	f.SetCellValue("Sheet1", "B1", "Anzahl der Schüler")

	for index, course := range courses {
		indexString := strconv.Itoa(index + 1)
		f.SetCellValue("Sheet1", "A"+indexString, course.Name)
	}

	for index, course := range courses {
		indexString := strconv.Itoa(index + 1)
		f.SetCellValue("Sheet1", "B"+indexString, len(course.Students))
	}

	return f.SaveAs(outputPath)
}

// WriteScheduleAsTable writes the result as table to stdout
func WriteScheduleAsTable(schedule *Schedule) {

}
