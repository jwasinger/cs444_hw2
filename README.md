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
          philosophers[i].lower = mutexes[i-1]
          philosophers[i].higher = mutexes[i]
          # this assumes that list indexing operates in the same way as Python
```
