<h1 style="text-align: center;">
    B"H
</h1>



<h2 style="text-align: center;">
Concurrency in Go
</h2>



---

### URL's
- [O'Reilly Learning Book](https://learning.oreilly.com/library/view/concurrency-in-go/9781491941294/ch01.html)
- [Errata](https://www.oreilly.com/catalog/errata.csp?isbn=0636920046189)








---
<h2 style="text-align: center;">
Chapter 1. An Introduction to Concurrency
</h2>


#### Amdahl’s law
- Amdahl’s law describes a way in which to model the potential performance gains from implementing the solution to a problem in a parallel manner. 
- Simply put, it states that the gains are bounded by how much of the program must be written in a sequential manner.


---

#### Why Is Concurrency Hard?

1. **Race Conditions**

    - A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed.

    - Most of the time, data races are introduced because the developers are thinking about the problem sequentially. 

    - They assume that because a line of code falls before another that it will run first.

    - I sometimes find it helpful to imagine a large period of time passing between operations. Imagine an hour passes between the time when the goroutine is invoked, and when it is run. How would the rest of the program behave? 

    - Race conditions are one of the most insidious types of concurrency bugs because they may not show up until years after the code has been placed into production.

2. **Atomicity**
    - When something is considered atomic, this means that within the context that it is operating, it is indivisible, or uninterruptible.

    - The first thing that’s very important is the word “context.” Something may be atomic in one context, but not another.

    - Something that is atomic will happen in its entirety without anything happening in that context simultaneously.

    - Atomicity is important because if something is atomic, implicitly it is safe within concurrent contexts.

    - Most statements are **not** atomic, let alone functions, methods, and programs. 

    - We’ll go into more depth later, but in short we can force atomicity by employing various techniques.

3. **Memory Access Synchronization**

    - There’s a name for a section of your program that needs exclusive access to a shared resource. This is called a **critical section**.

4. **Deadlocks**    
    
    - A deadlocked program is one in which all concurrent processes are waiting on one another. 
    
    - In this state, the program will never recover without outside intervention.

    - The Go runtime attempts to do its part and will detect some deadlocks (all goroutines must be blocked, or “asleep”), but this doesn’t do much to help you prevent deadlocks.


    **Coffman Conditions**
    - There are a few conditions that must be present for deadlocks to arise. 

    - If we ensure that **at least one** of these conditions is not true, we can prevent deadlocks from occurring.
    
        **Mutual Exclusion**
        - A concurrent process holds exclusive rights to a resource at any one time.

        **Wait For Condition**
        - A concurrent process must simultaneously hold a resource and be waiting for an additional resource.

        **No Preemption**
        - A resource held by a concurrent process can only be released by that process, so it fulfills this condition.

        **Circular Wait**
        - A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.

        ![](img/d-lock.png)

5. **Livelocks**

    - Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

    - Have you ever been in a hallway walking toward another person? She moves to one side to let you pass, but you’ve just done the same. So you move to the other side, but she’s also done the same. Imagine this going on forever, and you understand livelocks.        

    - In my opinion, livelocks are more difficult to spot than deadlocks simply because it can appear as if the program is doing work.

    - Livelocks are a subset of a larger set of problems called starvation. We’ll look at that next.

6. Starvation
    - Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.

    -  Livelocks warrant discussion separate from starvation because in a livelock, all the concurrent processes are starved equally, and no work is accomplished. More broadly, starvation usually implies that there are one or more greedy concurrent process that are unfairly preventing one or more concurrent processes from accomplishing work as efficiently as possible, or maybe at all.

    - One of the ways you can detect and solve starvation is by logging when work is accomplished, and then determining if your rate of work is as high as you expect it.

    - We should also consider the case where the starvation is coming from outside the Go process. Keep in mind that starvation can also apply to CPU, memory, file handles, database connections: any resource that must be shared is a candidate for starvation.

---

### Some Advice

![](img/advice-1.png)

---

![](img/advice-2.png)

---

![](img/advice-3.png)













---
<h2 style="text-align: center;">
Chapter 2. Modeling Your Code: Communicating Sequential Processes
</h2>


### The Difference Between Concurrency and Parallelism

Concurrency is a property of the code; parallelism is a property of the running program.

- We do not write parallel code, only concurrent code that we hope will be run in parallel. Parallelism is a property of the runtime of our program, not the code.

---
---

![](img/abstraction-1.png)

![](img/abstraction-2.png)

---

### CSP - Communicating Sequential Processes

In this paper, Hoare suggests that **input** and **output** are two overlooked primitives of programming—particularly in concurrent code.

![](img/csp.png)

---

Goroutines free us from having to think about our problem space in terms of parallelism and instead allow us to model problems closer to their natural level of concurrency.

![](img/go-cool-1.png)

- Go’s runtime multiplexes goroutines onto OS threads automatically and manages their scheduling for us. 

---

![](img/go-cool-2.png)

---

![](img/go-cool-3.png)

---

![](img/when.png)

![](img/when-2.png)


**Go’s philosophy on concurrency can be summed up like this: aim for simplicity, use channels when possible, and treat goroutines like a free resource.**








---
<h2 style="text-align: center;">
Chapter 3. Go’s Concurrency Building Blocks
</h2>

