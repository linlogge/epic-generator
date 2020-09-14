package main

import (
	"errors"
	"fmt"
	"os"

	app "gibhub.com/davincigs/courses/app"
	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
	cli "github.com/urfave/cli/v2"
)

const (
	// DefaultMaxStudents is the default maximum number of
	// students that can be in a class simultaneously
	DefaultMaxStudents = 15
	// DefaultMaxGenerations is the default maximum number
	// of generations that are the same before the algorithm
	// stops to evolve
	DefaultMaxGenerations = 2500
)

var (
	// ErrInputPathEmpty throws if input file path is not specified
	ErrInputPathEmpty = errors.New("Input file path must be set with --input")
	// ErrMaxStudentsTooSmall throws if the max amount of students is too small to compute
	ErrMaxStudentsTooSmall = errors.New("The provided maximum number of students must be at least 5")
)

// ErrCourseTooLarge throws if the provied max amount of students
// in a class is less than the half of the courses student count
type ErrCourseTooLarge struct {
	Name     string
	Students int
}

func (e ErrCourseTooLarge) Error() string {
	return fmt.Sprintf("The course %v has too many students (%v)", e.Name, e.Students)
}

func main() {
	app := &cli.App{
		Name:   "epic-generator",
		Usage:  "a generator to generate a schedule for seperate learning",
		Action: Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "input",
				Usage: "the input file path",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "the output file path",
			},
			&cli.IntFlag{
				Name:  "students",
				Usage: "the max number of students in a class",
				Value: DefaultMaxStudents,
			},
			&cli.IntFlag{
				Name:  "generations",
				Usage: "the max number generations that are the same",
				Value: DefaultMaxGenerations,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Run runs the application
func Run(cmd *cli.Context) error {

	inputPath := cmd.String("input")
	if inputPath == "" {
		return ErrInputPathEmpty
	}

	maxStudents := cmd.Int("students")
	if maxStudents < 5 {
		return ErrMaxStudentsTooSmall
	}

	workbook, err := excel.OpenFile(inputPath)
	if err != nil {
		return err
	}

	courses, students, err := app.Deserialize(workbook)
	if err != nil {
		return err
	}

	for _, course := range courses {
		if course.CountStudents()/2 > maxStudents {
			overflow := course.CountStudents() - (maxStudents * 2)
			return ErrCourseTooLarge{Name: course.Name, Students: overflow}
		}
	}

	maxGenerations := cmd.Int("generations")

	schedule, generations, duration := app.RunAlgorithm(&app.Algorithm{
		Courses:        courses,
		Students:       students,
		MaxStudents:    maxStudents,
		MaxGenerations: maxGenerations,
	})

	fmt.Println()
	fmt.Println("Fitness:", schedule.Fitness)
	fmt.Println("Conflicts:", schedule.Conflicts)
	fmt.Println("Generations:", generations)
	fmt.Println("Duration:", duration)
	fmt.Println()

	outputPath := cmd.String("output")
	if outputPath == "" {
		app.WriteScheduleToStdOut(schedule)
		return nil
	}

	return app.WriteScheduleAsFile(schedule, outputPath)
}
