package app

// Schedule represents the finished
// distributable timetable
type Schedule struct {
	Weeks          []*Week
	Conflicts      int
	Fitness        float32
	FitnessChanged bool
}

// ScheduleByFitness is a helper type
// to sort schedules by fitness
type ScheduleByFitness []*Schedule

func (b ScheduleByFitness) Len() int {
	return len(b)
}

func (b ScheduleByFitness) Less(i, j int) bool {
	return b[i].GetFitness() < b[j].GetFitness()
}

func (b ScheduleByFitness) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// CalculateFitness is the main function to find
// the best combination of students and classes
func (s *Schedule) CalculateFitness() float32 {
	s.Conflicts = 0

	for _, studentInWeekOne := range s.Weeks[0].Students {
		for _, studentInWeekTwo := range s.Weeks[1].Students {
			if studentInWeekOne.ID == studentInWeekTwo.ID {
				s.Conflicts += len(Courses)
			}
		}
	}

	for _, course := range Courses {
		weekOneStudents := course.CountMembersBy(s.Weeks[0].Students)
		weekTwoStudents := course.CountMembersBy(s.Weeks[1].Students)

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

	var weeks []*Week

	studentsInWeekOne, studentsInWeekTwo := splitStudents(Students)

	weekOne := &Week{
		Week:     0,
		Students: studentsInWeekOne,
	}
	weekTwo := &Week{
		Week:     1,
		Students: studentsInWeekTwo,
	}

	weeks = append(weeks, weekOne, weekTwo)

	return &Schedule{
		Weeks:          weeks,
		Fitness:        -1,
		FitnessChanged: true,
	}
}
