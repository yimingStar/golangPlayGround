# Reference
---
Youtube video: https://www.youtube.com/watch?v=YS4e4q9oBaU&t

# Variables
---
Variables have been ask to be nice and clean, we are able to add shadow variables, but not redeclare in the same statement.

- Always have to be use
- Declare variables

  ```
  var i int = 2
  ```
  or to let the compiler decide for us, as an auto type in C++
  
  ```
  i := 2
  ```
  
- Variables in **lower case** will be expose out side outside of the package.

# Primitive
---
Go is very hesitant about implicit data conversion. So, we must be clear of the data type.

## Integer 
- **int8**, **int16**, **int32**, **int64** from different platform which the number represent the bits.  
- **int** is Platform dependent in Golang.  
- **uint** as unsigned integer.  

## Bit operators
- **&, |, ^, &^** bits operator, let's see the example below.

  ```
    a:= 10 // 1010
    b:= 3 // 0011

    // a & b = 0010
    // a | b = 1011
    // a ^ b = 1001
    // a &^ b = oposite of or = 0100
  ```

- **>>, <<** bit shifting

  ```
    a:=8 // 2^3
    // a << 3 = 0100 shift 3 bit left = 2^3 * 2^3 = 2^6
    // a >> 3 = shift 3 bit right = 2^0
  ```

## Floating point numbers   
  Follows IEEE 754 standard, with 32 bit and 64 bit.
  
  ```
    n := 3.14 (var n float32 = 3.14) // Decimal
    n = 13.7e72 // Exponential
    n = 2.1E14 // 13.7e12
  ```
  
## Complext number
  ```
    var n complext64 = 1 + 2i
  ```
  - use real(v) to get the real part (float)
  - use imag(v) to get the imagine part (float)

## string
  - string can be as array just like C++
  - ***immutable***
  - Adding string just use + operator   
  
  Covert the string to byte collection (slice of byte)  
	  ```
	    s := "We are the best"
	    b := []byte(s)
	    fmt.Printf("%v, %T\n", b, b)
	  ```
  
  We get a byte slice with ascii value or utf value, the type result is []unit8  
  ***Many function in Golang is work as slice of byte (string, files)***
  
## rune
  - Golang doesn't have a char data type. It uses byte and rune to represent character values. 
  - Alias for int32
  - Respresent utf32 charater - check this https://pkg.go.dev/strings#Reader.ReadRune
    - if we are working utf32 we can use the above function to get the information we need.   
    <br/>
	
	```  
		r := 'a'
		fmt.Printf("%v, %T\n", r, r)

		// 97, int32
	```


# Constant 
---

# Function
---
- https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11
- Method (Function) in Golang https://go.dev/tour/methods/1

# Array & Slice
---

# Map & Structs
---
## Map
- **Referrence to same underlying data**, not copy.
- Create via literals or make function.
- Check for presence with "value, ok" form of result.


	```
	demoMap := make(map[string]int)
	if _, ok := demoMap["a"]; !ok {
		fmt.Println("unfound key a")
	}
	```
	
## Structs 
- Fields can points to any data type in go.
- Normally created as types, but anonymous structs are allowed.
- Structs are value type, will **copy to new struct**.
- Field can be tagged.
- No inheritance, but use **Composition** to embed.
	- Difference between composition and inheritance in Go, is a struct which inherits from another struct can directly access the methods and fields of the parent struct.
	<br/>
	
	```
	type author struct {
		firstName string
		lastName  string
		bio       string
	}

	func (a author) fullName() string {
		return fmt.Sprintf("%s %s", a.firstName, a.lastName)
	}

	type blogPost struct {
		title   string
		content string
		author
	}

	func (b blogPost) details() {
		fmt.Println("Title: ", b.title)
		fmt.Println("Content: ", b.content)
		fmt.Println("Author: ", b.author.fullName())
		fmt.Println("Bio: ", b.author.bio)
	}

	func main() {
		author1 := author{
			"Naveen",
			"Ramanathan",
			"Golang Enthusiast",
		}

		fmt.Println(author1)
		fmt.Println(author1.fullName())

		blogPost1 := blogPost{
			"This is fantastic",
			"Read through this more than once, promise me.",
			author1,
		}

		fmt.Println(blogPost1)
		fmt.Println(blogPost1.author.fullName())
		blogPost1.details()
	}
	```

# Condition Statement
---
- If Statement

	```
	// num:=9 is an inialization.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	```
- Switch Statement
	Input the tag of switch for checking the condition.
	Tag can be empty input, and also same as If statment with an initializer.
	"break" is already implied.
	
	```
	switch i := 9; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("unmatched")
	}
	```
	
	Use "fallthrough if we want the statement just continue execute the next statment, be aware it is logicless.

# Looping
---
- Also, with **break** and **continue**

	```
	sum := 0
	for i := 1; i < 5; i++ {
	    sum += i
	}
	```
- Another literal format	

	```
	for idx, value := range LIST {
		...
	}
	
	for key, value := range MAP {
		...
	}
	```
	
- With Loop Label using break and continue.

	```
	guestList := []string{"bill", "jill", "joan", "andy", "kelly"}
	arrived := []string{"sally", "jill", "joan", "kelly"}
	CheckList:
		for _, guest := range guestList {
			for _, person := range arrived {
				fmt.Printf("Guest[%s] Person[%s]\n", guest, person)

				if person == guest {
					fmt.Printf("Let %s In\n", person)
					continue CheckList
				}
			}
		}
	```

# Control Flow
---
- Defer 
	- Execute after the function are done
	- Execute last first, first last
	- Common use case, use it at closing a source in the beginning of using the source
	  
		```
		resp, err := ...
		defer resp.body.close()
		\\ continue to use the resource
	  	```
	  
- Panic
	- Use it when the code cannot be recovered. (Webserver port has been occupied)
	- Generate a panic object, which will stop the code when panic pops. (write then abort)
	- Unlike assert in C++, panic does not combine with conditions.
	- The print out, executes after defer 
	
		```
		fmt.Println("start")
		panic("smth bad happens")
		fmt.Println("end")
		```
- Recover
	- Function is built to recover from panic
	
		```
		func main() {
			fmt.Println("start")
			panicked()
			fmt.Println("end")
		}

		func panicked() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("Err:", err)
				}
			}()
			panic("smth bad happens")
		}
		```
# Pointer
---
```
var a int = 32
var b *int = &a
a = 42

c := &a
fmt.Printf("%d, %d, %p, %d, %p", a, *b, b, *c, c)
```

If we want to do pointer arithmetic us the unsafe package
But go seems this as a complex code, so they save it in unsafe package to go through compiler

```
a := [3]int{1, 100, 300}
b := &a[0]
c := &a[1]
fmt.Printf("%v, %p, %p", a, b, c)
```

Also it works with a struct
```
type demoStruct struct {
	test int
} 

var aPtr *demoStruct 
structPtr = &demoStruct{test: 10}
fmt.Println(aPtr)

var bPtr *demoStruct
fmt.Println(bPtr)
bPtr = new(demoStruct)
fmt.Println(bPtr) // initial an empty struct
bPtr.test = 15
fmt.Println(bPtr)
```

Reminder: Map, Slice is using pointer to point to the same address.

# Function
---
- Go technically has only pass-by-value, which is a copy of the input. Typically, we pass a variable or a pointer.
- Wrap input variable into slice

	```
	func main() {
		ans := sum(1, 2, 3)
		fmt.Println(ans)
	}

	func sum(elements ...int) int {
		fmt.Println(elements)
		result := 0
		for _, v := range elements {
			result += v
		}
		return result
	}
	```	

- Return a local variable as a pointer, this will be store in the heap memory (share) in run time for avoiding been clear.<br/>
  (In most of language, function result will be store in stack, and pop as the function returns)
	
	```
	func main() {
		ans := sum(1, 2, 3)
		fmt.Println(*ans)
	}

	func sum(elements ...int) *int {
		fmt.Println(elements)
		result := 0
		for _, v := range elements {
			result += v
		}
		return &result
	}
	```

- type Method Function

	```
	func main() {
		a := price(5)
		a.checkPrice()
		a.double()
		a.checkPrice()
	}

	type price int

	func (p price) checkPrice() {
		fmt.Println(p)
	}

	func (p *price) double() {
		*p = (*p) * 2
	}
	```
	
	Most likely using with Struct
	
	```
	func main() {
	alpha := company{
		name:           "Alphabet",
		headQuarter:    "MTV",
		employeeNumber: 20030,
	}

	alpha.printCompanyName()

	gtech := techCompany{
		parent:        alpha,
		name:          "Google",
		concentration: "Search Engine",
	}

		gtech.printTechCompany()
	}

	type company struct {
		name           string
		headQuarter    string
		employeeNumber int
	}

	type techCompany struct {
		parent        company
		name          string
		concentration string
	}

	func (c company) printCompanyName() {
		fmt.Println(c.name)
	}

	func (c techCompany) printTechCompany() {
		fmt.Println(c)
	}
	```

# Interface
---
- Why using Interface?

	Interface typically can easily build object with Polymorphism charateristic.
	(Referrence: https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang)

	Interfaces are too big of a topic to give an all-depth answer here, but some things to make their use clear.
	Interfaces are a tool. Whether you use them or not is up to you, but they can make code clearer, shorter, more readable, and they can provide a nice API between packages, or clients (users) and servers (providers).

	For example: 

	```
	type Cat struct{}

	func (c Cat) Say() string { return "meow" }

	type Dog struct{}

	func (d Dog) Say() string { return "woof" }

	func main() {
		c := Cat{}
		fmt.Println("Cat says:", c.Say())
		d := Dog{}
		fmt.Println("Dog says:", d.Say())
	}
	```

	In this example, we had two different animal that contains the same action function.
	Assume now we have a input that contains multiple animal object, and our goal is to print let all the animal object say something. 
	What is the possible way?

	```
	// our input must be divided into different type of slice
	c1 := Cat{}
	c2 := Cat{}
	c3 := Cat{}
	var catBox := []Cat{c1, c2, c3}

	for _, a := range catBox {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}


	... Do the same thing to Dogs
	```

	Can we make a upper level type and and wrap the same method into one caller interface.
	(Be aware this is different if we make a new struct and Animal struct and let Cat and Dog use Composition, we need the implementations contain different behaviors)
	Using interface and set both object Cat and Dog into an container

	```
	type Sayer interface {
		Say() string
	}

	animals := []Sayer{c, d}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}
	```

	Typically, single method usually, use the caller function + er. 
	But acually we can also called it Animal and with multiple interface functions.
	This shows the polymorphism in Golang.

- Compose Interface 
	
	We can also create a interface composing difference interface
	
	```
	type FileController interface {
		Writer
		Closer
	}

	type Writer interface {
		Write(byte[]) (int, error)
	}

	type Closer interface {
		Close() error
	}
	```

- Using interface for type switch, a interface can respresent diverse variable type.
	
	```
	var i interface{} = "test"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	default:
		fmt.Printf("i is not an integer, it is a %T", i)
	}
	```

# GoRoutine
---
- Use for concurrent and parallel programs.
- It creates **green thread (virtual thread)**, more lightly than os thread

- How to join and wait for a thread?
	Let's start a simple example. This example can simply execute the print function by go routine.
	However, if we did not put the sleep function the main will exit before the go routine print out.

	```
	func demomain() {
		msg := "Hello!"

		go func() {
			fmt.Println(msg)
		}()

		// race condition occurs, this is bad, use argument
		// this example go routine will go function stack looking for msg
		msg = "Goodbye"

		// this is not a good pratice, do not use sleep call, use waitGroup
		time.Sleep(100 * time.Millisecond)
	}
	```

	Now, according to above exmaple, we should use wait group to wait for the go routine.
	(Kind of thread join in Python)

	```
	var wg = sync.WaitGroup{}

	func main() {
		routineNum := 1
		wg.Add(routineNum)
		msg := "Hello!"
		go shoutOut(msg)

		wg.Wait()
	}

	func shoutOut(msg string) {
		fmt.Println(msg)
		wg.Done()
	}
	```
	
- How to avoid race condition? Using **Mutex**, which are able to lock resources.
	
	Down below, the example will definitely encounters race condition.
	
	```
	var wg = sync.WaitGroup{}
	var counter int = 0

	func main() {
		for i := 0; i < 10; i++ {
			wg.Add(2)
			go printCounter()
			go increaseCounter()
		}
		wg.Wait()
	}

	func printCounter() {
		fmt.Printf("The counter: %d\n", counter)
		wg.Done()
	}

	func increaseCounter() {
		counter++
		wg.Done()
	}
	```
	
	Hence, let's use mutex to lock the variable.
	
	```
	var wg = sync.WaitGroup{}

	// RW mutex allow mutiple read but only one write
	var m = sync.RWMutex{}
	var counter int = 0

	func main() {
		for i := 0; i < 10; i++ {
			wg.Add(2)
			m.RLock()
			go printCounter()
			m.Lock()
			go increaseCounter()
		}
		wg.Wait()
	}

	func printCounter() {
		fmt.Printf("The counter: %d\n", counter)
		m.RUnlock()
		wg.Done()
	}

	func increaseCounter() {
		counter++
		m.Unlock()
		wg.Done()
	}
	```

- Limiting the thread numbers, the default thread numbers will equal to the core numbers.
	This number is a little tricky, you wont want it too high
	```
	fmt.Printf("Thread numbers %v\n", runtime.GOMAXPROCS(-1))
	```

- Important things to remind ourselves
	1. Do not use go routine in library, let the consumer control the concurrency.
	2. Know exactly when the routine is going to end.
	3. Check race condition at compile time - adding race while compiling

		```
		go run -race main.go
		```

# Channels
---
Accordingly, now we know how go routine works, then next coming up question is - how does go routines communicate to each other?
	- Access same variable?
	- Passing messages? (Socket? Queue?)
	
As a result, golang provides a message passing mechanism using **Channel**
- It is a unbuffered FIFO queue, only one message will be send through the Channel at one time
- Mostly, we want go routine dedicate to be a sender or receiver 

	```
	func main() {
		// using make function to create channel
		ch := make(chan int) // create a channel passing message type int
		wg.Add(2)
		// this is a receiver, receive int from ch
		go func(ch <-chan int) {
			i := <-ch
			fmt.Printf("receive %d from channel %v", i, ch)
			wg.Done()
		}(ch)

		// this is a sender, the input limit this routine only can be a sender
		go func(ch chan<- int) {
			ch <- 28
			// this will be invalid in this go routine
			// i := <-ch
			// fmt.Printf("receive %d from channel %v", i, ch)
			wg.Done()
		}(ch)
		wg.Wait()
	}
	```

- Since there is no buffer in the channel we create above if there is no receiver matching the sender things will go wrong, the routine will be locked.

	```
	func main() {
		// using make function to create channel
		ch := make(chan int) // create a channel passing message type int
		wg.Add(2)
		// this is a receiver, receive int from ch
		go func(ch <-chan int) {
			i := <-ch
			fmt.Printf("receive %d from channel %v", i, ch)
			wg.Done()
		}(ch)

		// this is a sender, the input limit this routine only can be a sender
		go func(ch chan<- int) {
			ch <- 28
			ch <- 33
			wg.Done()
		}(ch)
		wg.Wait()
	}
	```
	
	Let's create buffer to deal with the problem.
	
	```
	ch := make(chan int. 50) // create 50 integer buffer
	```
	
	This may solve the data operate in different frequency in different sides.
	But this still not a ideal way to solve the problem. Since it will still be blocked when the buffer is full.
	For example, a application with data peak. It might be running out of spaces.
	
	```
	go func(ch <-chan int) {
		// we need a close() to inform this loop need to close end exit the go routine
		for i := range ch {
			fmt.Printf("receive %d from channel %v\n", i, ch)
		}
		wg.Done()
	}(ch)

	// this is a sender, the input limit this routine only can be a sender
	go func(ch chan<- int) {
		ch <- 28
		ch <- 33
		close(ch)
		wg.Done()
	}(ch)
	```
	
	Another way to write the receiver in this format
	```
	go func(ch <-chan int) {
		for {
			if i, ok <-ch; ok {
				fmt.Printf("receive %d from channel %v\n", i, ch)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	```
- What if no obvious timing to close the go routine? (https://youtu.be/YS4e4q9oBaU?t=23245)
	- Using defer to close the channel
	- Using signal to ping the receiver and close the channel with **select statement** 
		- https://go.dev/tour/concurrency/5
		- https://stackoverflow.com/questions/38821491/what-is-the-difference-between-switch-and-select-in-go
	<br/>
	
	
	```	  
	// @Example 3
	// building a signal channel -> it's use to ping our function
	logCh := make(chan char, 50)
	signalDoneCh := make(chan struct{})

	// how to deal with implicit close timespot using "select statement"
	func main() {
		// using make function to create channel
		ch := make(chan int, 50) // create a channel passing message type int

		logCh <- "send something"
		logCh <- "send something"

		signalDoneCh <- struct{}{}
	}

	func logger() {
		for {
			// blocking select statement
			select {
			case entry := <-logCh:
				fmt.Printf("receive, do something")
			case <- signalDoneCh:
				break;
			}
		}
	}
	```
