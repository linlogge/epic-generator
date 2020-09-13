package app

import "math/rand"

var (
	// Courses are the courses to
	// run the algorithm on
	Courses []*Course
	// Students are all students
	// extracted from courses
	Students []*Student
	// MaxStudents is the amount
	// of students that can be in a
	// class simultaneously
	MaxStudents int
)

// Course represents a course
type Course struct {
	ID       int
	Name     string
	Students []*Student
}

// CountStudents counts the number of
// students in a course
func (c *Course) CountStudents() int {
	return len(c.Students)
}

// getAllStudents gets all students of all courses
func getAllStudents(courses []*Course) []*Student {
	var students []*Student

	for _, course := range courses {
		for _, cStudent := range course.Students {
			if !studentsContainStudent(students, cStudent) {
				students = append(students, cStudent)
			}
		}
	}

	return students
}

func studentsContainStudent(students []*Student, student *Student) bool {
	for _, s := range students {
		if s.ID == student.ID {
			return true
		}
	}
	return false
}

func splitStudents(students []*Student) ([]*Student, []*Student) {
	var studentsOne []*Student
	var studentsTwo []*Student

	for _, student := range Students {
		if rand.Intn(2) == 1 {
			studentsOne = append(studentsOne, student)
		} else {
			studentsTwo = append(studentsTwo, student)
		}
	}

	return studentsOne, studentsTwo
}

// Student represents a student
type Student struct {
	ID      string
	Courses []*Course
}

// Population represents a
// generation of schedules
type Population struct {
	Size      int
	Schedules []*Schedule
}

func newPopulation(size int) *Population {

	var schedules []*Schedule

	for i := 0; i < size; i++ {
		schedules = append(schedules, newSchedule())
	}
	return &Population{Size: size, Schedules: schedules}
}
