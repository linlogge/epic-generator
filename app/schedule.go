package app

// Schedule represents the finished
// distributable timetable
type Schedule struct {
	Classes        []*Class
	Conflicts      int
	Fitness        float32
	FitnessChanged bool
}

// ByFitness is a helper type
// to sort schedules by fitness
type ByFitness []*Schedule

func (b ByFitness) Len() int {
	return len(b)
}

func (b ByFitness) Less(i, j int) bool {
	return b[i].GetFitness() < b[j].GetFitness()
}

func (b ByFitness) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// CalculateFitness is the main function to find
// the best combination of students and classes
func (s *Schedule) CalculateFitness() float32 {
	s.Conflicts = 0

	for _, studentOne := range s.Classes[0].Students {
		for _, studentTwo := range s.Classes[1].Students {
			if studentOne.ID == studentTwo.ID {
				s.Conflicts += len(Courses)
			}
		}
	}

	for _, course := range Courses {
		weekOneStudents := course.CountMembersBy(s.Classes[0].Students)
		weekTwoStudents := course.CountMembersBy(s.Classes[1].Students)

		if weekOneStudents > MaxStudents || weekTwoStudents > MaxStudents {
			s.Conflicts += len(Courses)
		}

		if weekOneStudents < weekTwoStudents {
			s.Conflicts += (weekTwoStudents - weekOneStudents)
		} else if weekTwoStudents != weekOneStudents {
			s.Conflicts += (weekOneStudents - weekTwoStudents)
		}
	}

	return 1 / (float32(s.Conflicts) + 1)
}

// GetFitness checks if the fitness has
// changed and calculates it if necessary
func (s *Schedule) GetFitness() float32 {
	if s.FitnessChanged {
		s.Fitness = s.CalculateFitness()
		s.FitnessChanged = false
	}

	return s.Fitness
}

func newSchedule() *Schedule {

	var classes []*Class

	studentsOne, studentsTwo := splitStudents(Students)

	classOne := &Class{
		Week:     0,
		Students: studentsOne,
	}
	classTwo := &Class{
		Week:     1,
		Students: studentsTwo,
	}

	classes = append(classes, classOne, classTwo)

	return &Schedule{
		Classes:        classes,
		Fitness:        -1,
		FitnessChanged: true,
	}
}
