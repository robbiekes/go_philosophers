package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	allPhilos = 5
	think     = time.Millisecond * 200
	mustEat   = 10
)

func closeChannels(forks []Fork) {
	for i := 0; i < allPhilos; i++ {
		close(forks[i].fork)
	}
}

func createForks(forks []Fork) {
	for i := 0; i < allPhilos; i++ {
		forks[i].fork = make(chan int, 1)
	}
}

func createPhilos(philos []*Philo, forks []Fork) {
	wg := sync.WaitGroup{}
	wg.Add(allPhilos)
	for i := 0; i < allPhilos; i++ {
		philos[i] = &Philo{i + 1, 0, forks[i], forks[(i+1)%allPhilos]}
		go dinner(philos[i], &wg)
	}
	wg.Wait()
}

func dinner(philo *Philo, wg *sync.WaitGroup) {
	for philo.mealCount < mustEat {
		fmt.Printf("philosopher %d is thinking\n", philo.id)
		time.Sleep(think)
		philo.leftFork.fork <- 1
		philo.rightFork.fork <- 1
		fmt.Printf("philosopher %d took forks\n", philo.id)
		fmt.Printf("philosopher %d is eating\n", philo.id)
		philo.mealCount++
		<-philo.leftFork.fork
		<-philo.rightFork.fork
	}
	fmt.Println("philosopher", philo.id, "finished. he had dinner", philo.mealCount, "times")
	wg.Done()
}

func DiningPhilosophers() {
	philosophers := make([]*Philo, allPhilos)
	forks := make([]Fork, allPhilos)

	createForks(forks)
	createPhilos(philosophers, forks)
	closeChannels(forks)
}
