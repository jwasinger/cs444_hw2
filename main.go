package main

import (
  "fmt"
  "time"
  "sync"
  "runtime"
  "math/rand"
)

func Min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func Max(x, y int) int {
  if x > y {
      return x
  }
  return y
}

//sleep from 1-20 seconds
func think(position int, names [5]string) {
  //sleep_time := rand.Intn(19)+1
  sleep_time := rand.Intn(1)+1

  time.Sleep(time.Second*time.Duration(sleep_time))
}

func get_forks(forks *[5]sync.Mutex, position, minFork int, maxFork int, names [5]string) {
  fmt.Println(names[position], " ", minFork, " ", maxFork)
  forks[minFork].Lock()
  forks[maxFork].Lock()
}

func put_forks(forks *[5]sync.Mutex, minFork int, maxFork int, names [5]string) {
  forks[minFork].Unlock()
  forks[maxFork].Unlock()
}

func eat(position int, names [5]string) {
  fmt.Println(names[position], " is eating")
  time.Sleep(time.Duration(5)*time.Second)
}

func philosopher(forks *[5]sync.Mutex, position int, names [5]string) {
  var right int = 0
  var left int = position

  if position+1 == 5 {
    right = 0
  } else {
    right = position+1
  }


  max := Max(left, right)
  min := Min(left, right)

  for ;; {
    think(position, names)
    get_forks(forks, position, min, max, names)
    eat(position, names)
    put_forks(forks, min, max, names)
    fmt.Println(names[position], " is done")
  }
}

func main() {
  runtime.GOMAXPROCS(2) // max number of CPUs used

  names := [5]string{"1","2","3","4","5"}
  forks := [5]sync.Mutex{}

  wg := sync.WaitGroup{}

  wg.Add(1)

  for i := 0; i < 5; i++ {
    go philosopher(&forks, i, names) // create a new goroutine (thread)
  }

  wg.Wait()
}
