# Problem Statement


# Djikstra's Solution

1) Number each fork 0-5
2) Each philosopher does the following:
```
        get_lower_fork()
        get_higher_fork()
        eat()
        put_down_forks()
        think()
```
* retrieving a fork blocks until that fork is available
* order in which forks are returned does not matter
* eat() blocks for a fixed amount of time
* think() sleeps for a random amount of time

```
        philosophers[5] X F => mutexes[5]
        
        F(i):
          philosophers[i].lower = min(i-1, i)
          philosophers[i].higher = mutexes[i]
          # assume that i has a range of [0-5] inclusive, 0-1 = 5 (negative wraps around)
```
