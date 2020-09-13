package main

import (
	"errors"
	"fmt"
	"os"

	app "gibhub.com/davincigs/courses/app"
	excel "github.com/360EntSecGroup-Skylar/excelize/v2"
	cli "github.com/urfave/cli/v2"
)

var (
	// ErrInputPathEmpty throws if input file path is not specified
	ErrInputPathEmpty = errors.New("Input file path cannot be empty")
	// ErrInvalidFileFormat throws if the input file could not be read by excelize
	ErrInvalidFileFormat = errors.New("The provided file is not a valid Excel file")
	// ErrMaxStudentsTooSmall throws if the max amount of students is too small to compute
	ErrMaxStudentsTooSmall = errors.New("The provided maximum number of students must be at least 5")
)

const (
	// DefaultFitness is the default fitness the
	// algorithm uses to determine when to stop
	DefaultFitness = 0.2
	// DefaultMaxStudents is the default number of
	// students that can be in a class simultaneously
	DefaultMaxStudents = 15
)

func main() {
	app := &cli.App{
		Name:   "covidc",
		Usage:  "an algorithm to schedule courses",
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
				Name:  "max",
				Usage: "the max number of students in a class",
				Value: DefaultMaxStudents,
			},
			&cli.Float64Flag{
				Name:  "fitness",
				Usage: "the fitness the algorithm should stop",
				Value: DefaultFitness,
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

	// Check if user provided an input file path
	inputPath := cmd.String("input")
	if inputPath == "" {
		return ErrInputPathEmpty
	}

	// Check if user provided a max fitness
	fitness := cmd.Float64("fitness")

	// Check if user provided a max amount of students
	maxStudents := cmd.Int("max")
	if maxStudents < 5 {
		return ErrMaxStudentsTooSmall
	}

	// Open file with provided input file path
	workbook, err := excel.OpenFile(inputPath)
	if err != nil {
		return ErrInvalidFileFormat
	}

	courses, students, err := app.Deserialize(workbook)
	if err != nil {
		return err
	}

	schedule := app.RunAlgorithm(courses, students, maxStudents, float32(fitness))

	outputPath := cmd.String("output")
	if outputPath == "" {
		app.WriteScheduleToStdOut(schedule)
		return nil
	}

	app.WriteScheduleAsFile(schedule, outputPath)

	return nil
}
