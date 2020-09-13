package app

import (
	"fmt"

	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ErrWrongNumberOfColumns throws if a table
// has a wrong amount of columns
type ErrWrongNumberOfColumns struct {
	table    string
	colCount int
}

func (e ErrWrongNumberOfColumns) Error() string {
	return fmt.Sprintf("The number of columns in table %v is %v but must be 2", e.table, e.colCount)
}

// Deserialize deserializes an excel file
// containing tables of courses into models
func Deserialize(workbook *excel.File) ([]*Course, []*Student, error) {
	var courses []*Course
	var students []*Student

	// Extract all courses from tables
	for i, table := range workbook.GetSheetList() {
		rows, err := workbook.GetRows(table)
		if err != nil {
			return nil, nil, err
		}

		var students []*Student

		if len(rows) <= 1 {
			continue
		}

		for j, row := range rows {

			if j == 0 {
				continue
			}

			colCount := len(row)
			if colCount != 2 {
				return nil, nil, ErrWrongNumberOfColumns{table, colCount}
			}
			students = append(students, &Student{ID: row[1]})
		}

		course := &Course{Name: table, Students: students, ID: i}
		courses = append(courses, course)
	}

	students = getAllStudents(courses)

	// Append courses on students
	for _, student := range students {
		for _, course := range courses {
			for _, courseStudent := range course.Students {
				if courseStudent.ID == student.ID {
					student.Courses = append(student.Courses, course)
				}
			}
		}
	}

	return courses, students, nil
}
