package main

type Philo struct {
	id                  int
	mealCount           int
	rightFork, leftFork Fork
}

type Fork struct {
	fork chan int
}
