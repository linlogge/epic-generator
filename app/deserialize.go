package app

import (
	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
)

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

			if len(row) == 2 {
				students = append(students, &Student{ID: row[1]})
			}
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
