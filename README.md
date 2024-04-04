# Kata Presentation

This kata had two phases for me, the research phase and the implementation phase, the `src/list` directory contains the research phase, a lot benchmarks and tests to find out how the allocations behave, and the `pkg` directory, where I implemented my final result.

## Research

I started with a benchmark for the `.Map` functionality, to determine which is the most effective way to create a list and map through it.

The Benchmark will add 1000 entries (an object with one integer value) to a list, .Map to a new list with a function that multiplies the value of the object by 2, and check the value.

![benchmark image](./images/bench_map_0.png)

The untyped list solution in [this file](./src/list/untyped_list.go), has similar performance profiles as the native solution with pointers, however, it still uses a quarter less allocations as direct pointers.

![benchmark comparision for untyped and pointers](./images/bench_map_1.png)

When I use cap or len on the slices, there are around 12 allocations less per slice, because go has to extend the slice on demand when `append` is called.

![benchmark highlights for allocations differences with cap and len](./images/bench_map_2.png)

A last test, was an optimal setup, where the native and the typed list implementation both have the initial len aswell as the len on `.Map` applied.

![benchmark with an optimal setup](./images/map_optimal_bench.png)

These benchmarks were done before any lock mechanism was implemented, this is the result after implementing lock:

![benchmark with an optimal setup after lock implementation](./images/map_optimal_after_lock.png)

Maybe I did something wrong, but I received a huge performance loss (around 18 times worse) after that. Currently using `Rlock` on every read operation like `.Get` and `.Lock` on any `.Set`, `.Map` and `.Add` operation.

I decided for now, to create `Sync` versions of the functions, like `Set` and `SetSync`, where `SetSync` is actually using the lock. The `locking` seems use so much CPU, that I wonder when is it actually better to go for goroutines and parallelism.

I tried to do some changes and use less pointers of my TypedList, and just work with the copy of the struct, but I wasn't able to get much performance improvements there, I was able to eliminate another heap allocation, but that's about it.

![benchmark with copy](./images/bench_map_3.png)

## Implementation

After researching some of the performance, I needed to make compromises between DX and Performance:

Create a new List with LEN vs. CAP vs. Uncapped

> Giving the possibility to initialize a list (or returning a list after map) with len, could lead to out of bounds issues, also, the performance increase of len vs. cap is not very remarkable. I will use cap whenever it is possible (for mapping), and try to avoid len for now.

Map on the struct vs. extra Map function vs. any response

> The `.Map` function on the struct makes sense, naturally, coming from TypeScript, the hell of generic complexity, I feel like I want to have the possibility to put one generic type into the `.Map` function, and get another one out. For example `func (l *list[R]) Map[T](mapper func (obj R) T) *list[T]`, we would transform a list of `R` into a list of `T` here. But it seems like it's not possible to put another generic on a function that is callable on a struct that already has a generic. A solution was that I would take an external `Map` function, that get's the `list` as first parameter, but that would not be `chainable`. Returning `any` as result and casting it afterwards would be another solution. Update: I tried `any`, and it increased the allocations agai but ~800, so this is probably not a valid solution when the focus is on performance.

## Resources

- https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d
- https://medium.com/@ankur_anand/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed
