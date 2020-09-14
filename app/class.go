package app

import (
	"fmt"
	"sort"
)

// Week represents a week
type Week struct {
	Week     int
	Students []*Student
}

// CountStudents counts the students of a week
func (w *Week) CountStudents() int {
	return len(w.Students)
}

// WeeksToCourses converts weeks into their initial courses
func WeeksToCourses(weeks []*Week) []*Course {
	var coursesList [][]*Course
	for _, week := range weeks {
		coursesList = append(coursesList, week.ToCourses())
	}

	courses := combineCourses(coursesList)
	sort.Sort(CoursesByName(courses))
	return courses
}

// ToCourses converts a week into its initial courses
func (w *Week) ToCourses() []*Course {
	var courses []*Course

	for i, course := range Courses {

		var students []*Student

		for _, student := range w.Students {
			for _, courseStudent := range course.Students {
				if student.ID == courseStudent.ID {
					students = append(students, student)
				}
			}
		}

		if len(students) > 0 {
			courses = append(courses, &Course{
				ID:       (w.Week * len(Courses)) + i,
				Students: students,
				Name:     fmt.Sprintf("%v, Woche %v", course.Name, w.Week+1),
			})
		}
	}

	return courses
}
