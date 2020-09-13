package app

import (
	"fmt"
	"sort"
)

// Class represents a class
type Class struct {
	Week     int
	Students []*Student
}

// CountStudents counts the students in a class
func (c *Class) CountStudents() int {
	return len(c.Students)
}

// ClassesToCourses converts classes into courses
func ClassesToCourses(classes []*Class) []*Course {
	var coursesList [][]*Course
	for _, class := range classes {
		coursesList = append(coursesList, class.ToCourses())
	}

	courses := combineCourses(coursesList)
	sort.Sort(CoursesByName(courses))
	return courses
}

// ToCourses converts a class into its initial course
func (c *Class) ToCourses() []*Course {
	var courses []*Course

	for i, course := range Courses {

		var students []*Student

		for _, student := range c.Students {
			for _, courseStudent := range course.Students {
				if student.ID == courseStudent.ID {
					students = append(students, student)
				}
			}
		}

		if len(students) > 0 {
			courses = append(courses, &Course{
				ID:       (c.Week * len(Courses)) + i,
				Students: students,
				Name:     fmt.Sprintf("%v, Woche %v", course.Name, c.Week+1),
			})
		}
	}

	return courses
}
