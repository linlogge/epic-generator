package app

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	populationSize         = 9
	turnamentSelectionSize = 3
	numbOfEliteSchedules   = 1
	mutationRate           = 0.1
)

// RunAlgorithm runs the algorithm
func RunAlgorithm(courses []*Course, students []*Student) *Schedule {
	Courses = courses
	Students = students

	// Set the seeder for our random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	genNumber := 0

	population := newPopulation(populationSize)

	for population.Schedules[0].Fitness <= 0.5 {
		genNumber++
		sort.Sort(sort.Reverse(ByFitness(population.Schedules)))
		population = evolve(population)
		fmt.Println("Fitness:", population.Schedules[0].Fitness)
	}

	sort.Sort(ByName(population.Schedules[0].Classes))

	for _, class := range population.Schedules[0].Classes {
		fmt.Printf("Week: %v, Students: %v\n", class.Week, len(class.Students))
		for _, student := range class.Students {
			fmt.Println("ID:", student.ID, "Kurse:", len(student.Courses))
		}
		fmt.Println()
	}
	return population.Schedules[0]
}

// Evolve evloves the population to a next gen
func evolve(pop *Population) *Population {
	return mutatePopulation(crossoverPopulation(pop))
}

func crossoverPopulation(pop *Population) *Population {
	crossoverPop := newPopulation(0)

	for i := 0; i < numbOfEliteSchedules; i++ {
		crossoverPop.Schedules = append(crossoverPop.Schedules, pop.Schedules[i])
	}

	for i := numbOfEliteSchedules; i < populationSize; i++ {
		scheduleOne := selectTurnamentPopulation(pop).Schedules[0]
		scheduleTwo := selectTurnamentPopulation(pop).Schedules[0]
		crossoverSchedule := crossoverSchedule(scheduleOne, scheduleTwo)
		crossoverPop.Schedules = append(crossoverPop.Schedules, crossoverSchedule)
	}
	return crossoverPop
}

func mutatePopulation(pop *Population) *Population {
	for i := numbOfEliteSchedules; i < populationSize; i++ {
		mutateSchedule(pop.Schedules[i])
	}
	return pop
}

func mutateSchedule(mutSchedule *Schedule) *Schedule {
	schedule := newSchedule()

	for i := 0; i < len(mutSchedule.Classes); i++ {
		if mutationRate > rand.Float32() {
			mutSchedule.Classes[i] = schedule.Classes[i]
		}
	}

	return mutSchedule
}

func crossoverSchedule(first *Schedule, second *Schedule) *Schedule {
	crossoverSchedule := newSchedule()

	for i := 0; i < len(crossoverSchedule.Classes); i++ {
		if rand.Intn(2) == 1 {
			crossoverSchedule.Classes[i] = first.Classes[i]
		} else {
			crossoverSchedule.Classes[i] = second.Classes[i]
		}
	}
	return crossoverSchedule
}

func selectTurnamentPopulation(pop *Population) *Population {
	turnamentPop := newPopulation(0)

	for i := 0; i < turnamentSelectionSize; i++ {
		schedule := pop.Schedules[rand.Intn(populationSize)]
		turnamentPop.Schedules = append(turnamentPop.Schedules, schedule)
	}

	sort.Sort(sort.Reverse(ByFitness(turnamentPop.Schedules)))

	return turnamentPop
}
