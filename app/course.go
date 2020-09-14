package app

// Course represents a course
type Course struct {
	ID       int
	Name     string
	Students []*Student
}

// CoursesByName is a helper type
// to sort courses by name
type CoursesByName []*Course

func (b CoursesByName) Len() int {
	return len(b)
}

func (b CoursesByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b CoursesByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func combineCourses(c [][]*Course) []*Course {
	var courses []*Course

	for _, courseList := range c {
		for _, course := range courseList {
			courses = append(courses, course)
		}
	}

	return courses
}

// CountMembersBy takes a list of students and returns the number
// of students that are a member of the course
func (c *Course) CountMembersBy(students []*Student) (counter int) {

	for _, courseStudent := range c.Students {
		for _, student := range students {
			if courseStudent.ID == student.ID {
				counter++
			}
		}
	}

	return counter
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
