# A path to OOP with Go - Workshop

## Schedule

- 09:00-09:10: Introduction to OOP
- 09:10-09:30: Warm up coding exercise
- 09:30-09:50: Types and Interfaces (types, aliases, interfaces, embedding, composition, type assertions) 
- 09:50-10:00: Break
- 10:00-10:30: Exercise 2 (interfaces)
- 10:30-10:50: Organizing your packages 
  - Who provides interfaces and emerging patterns
  - Inner package, when?
  - Package `internal`
- 10:50-11:00: Break
- 11:00-11:20: More theory: 
  - Short-lived objects, contexts
  - Default variables
  - The options, builder pattern
  - Code generation, why? When?
  - The `[]T{}` to `interface{}...` conversion problem
- 11:30-11:50: Generics
  - 
  - Constraints
    - any 
    - Comparables
    - Indexables
- 11:50-12:00: Break
- 12:00-12:45: Exercise 3 (generics)
- 12:45-13:00: Conclusion

## Introduction to OOP

### Is Go an Object-Oriented language? 
This question is so loaded that all I had to do was advertise this workshop to receive loads of feedback on it from total strangers.
Go doesn't offer classes, which means there are no constructors (or destructors) and no inheritance, etc. These are technical concepts that have become synonymous with OOP. 
However, as we will see, Go has a variety of very strong features for Object Oriented Programming that enable Gophers to express their code in a manner that follows the OO principals. 
I called this workshop a path to OOD with Go and not OOP with Go, because it doesn't follow the same design principals - in particular it should impact how we arrange our packages and when done well it reduces the amount of actual coding. 

### But first, what is OOP?

What we can all agree on: The central idea behind Object Oriented is to divide software into "things" or "objects" or "instances" that communicate via "messages" or "methods" or "member functions".
This core idea has not changed in the 4-5+ decades since it was conceptualized.
It is meant to allow the developer to build code and assign responsibilities just like in the real world, which is what we are familiar with, and how we generally think.

### Do you need OOP?

Just like in the real world, wherever there are things, there can be a mess. *__That's why Marie Kondo.__*
Just as you can write sane procedural code, you can write sane OO code. You and your team should define design best practices that match your needs. 

### Exercise 1:
Where we will understand some OO basics using an example of a gopher and a maze.

*This exercise is heavily inspired by the Intro to CS first home assignment that Prof. Jeff Rosendchein gave my CS class in 2003.

To get a sense of what strong OOP can do for us, solve a maze given a Gopher that can perform 4 actions:
```go
// Gopher is an interface to an object that can move around a maze
type Gopher interface {
	Finished() bool // Has the Gopher reached the target cell?
	Move() error    // The Gopher moves one step in its current direction
	TurnLeft()      // The Gopher will turn left
	TurnRight()     // The Gopher will turn right
}
```

Find the function `SolveMaze(g Gopher)` in cmd/maze/maze.go and implement it, then run:
```bash
make build
make test-maze
````

When your test passes (it only tests very basic navigation), run:
```bash
make run-maze > tmp/maze.html 
```
Open the maze.html file in the browser to see the result.
You can run the app multiple times to see your gopher running through different mazes.

#### Wasn't that simple?

OOP is all about ease and simplicity in being able to express our ideas and translate them into code as simply as possible while delegating responsibilities to things.

#### Let's review the code that made this possible.
The repo started with one package in the pkg directory called maze which offers a basic maze generator and nothing else. 






### Now that we understand how powerful OOP can be (and that Go does offer OOP), let's understand what features Go provides for OOP and how they correlate with those of other languages.





In Java:
To express that type A is B you can use inheritance exactly once, and unlimited number of interfaces
A class has to be aware of the interfaces it implements. 
To use type A as an interface it doesn't explicitly implement you must provide some code (either another class or generic behavior)
You can only inherit once (unlike in C++) so inheritance is very limited.
You have abstract classes

In C++:
You can inherit as many classes as you like
There is 
You have aliases

In Go: any type (except for primitives) can have methods.
You can make anything into a type. 
To express the concept that type A is B use interfaces
To express that type A is made of B - use composition (embedding)
An interface can also be a composition of interfaces
You can force a type to implement an interface if you need to.

In Ruby:
You can inherit exactly once
Ruby is moving more and more from inheritance to modules to express composition

Because interfaces are implicit:
1. Packages never provide the intervaces 
>_The problem with object-oriented languages is they've got all this implicit environment that they carry around with them. You wanted a banana but what you got was a gorilla holding the banana and the entire jungle._
(Joe Armstrong)

What did he mean by that?
The common OOP languages approach is that class A must inherit from class B or implement interface I in order to be used as an instance of B or I.
For instance, the class Banana will have to extend or inherit from Fruit (or a similar Object class) implement a Holdable interface just in case we ever want it to be held, implement a GrowsOnTree just in case we need to figure out where it came from. etc.
What happens if the Banana we imported doesn't implement an interface that we need it to? We create a MyBanana and inherit Banana just to be able to work with the original class but in the context that our application provides for it.
What happens if when the CTORs that were written for the Banana class don't match our context of creation (we cannot provide all fields and we just need a subset of the class)? - The class might be useless.

Inheritance:
To  

In an ideal world an application and all of its components are limited to absolutely necessary types.

###How does Go deal with that?
- Anything can be boxed into a type - primitives, structs, functions, etc.
- All types can have methods.
- Composition for structs and interfaces
- No CTORs: 
  - No more CTORs that don't reflect the reality of the construction of the object.     
  - Objects should be naturally instantiated with values according to the context that they were created for, requiring us to think in actions that lead to the creation of objects where those are appropriate (os.Open will return a File instead of using a File constructors)
  Interfaces are implicit - if it matches its requirements it implements the interface, a class doesn't have to provide all the interfaces that it will implement.







No classes - only structs
Any type can have methods
Mutating objects requires working with pointers


A package never provides an interface to its objects - that practice comes from languages where interfaces are explicite
Interfaces determine what your object is
Embeds determine what your object is made of (composed of)

```go
type UserModel struct{
        Model
        Name string
} 

type Model struct{
        ID int ``
		
}
```




###Why?
It was a long time concensus that "real gophers" don't need generics so much so that around the time the generics draft of 2020 was released, many gophers expressed that they will likely never use this feature.

But before we let any good feature go to waste, this is a good time to examine the Go features related to types in general, OOD in particular and how generics play into them.

###What is ODD?
Extracting functionality into things so that we can perform actions on them.
An object is a set of fields aggregated together and a set of functions (called methods) related to the logical entity that it represents
For instance if we described a person as an object, it will have certain attributes or field 

###What is abstraction?
Decoupling behavior from implementation 

###What is generalization?
It is very similar to abstraction

###What is encapsulation?
Hiding implementation including costs

###Creating objects in different languages
##The struct
The struct datatype comes from Algol, and it is commonly a set of fields aggregated in memory together and therefore instantiated together
In c++ it can also have methods and a constructor
In go it also has methods

##The class
The class is similar to the struct but will have a constructor and methods (there are other differences in c++ between a class and a struct, but they are not relevant to this comparison)
Typically has a constructor, or many constructors that allow defining initial values for its members (a comprehensive state)
Destructors only make sense when the user can predict when they will execute, this paradigm doesn't exist in the GC languages world, but it is very much alive in C++ still when the user manages the heap and therefore can determine when destructors are invoked.
The evolution of moving towards having no destructors and therefore having to maintain explicit control over operations like close might have had something to do with Go's choice not to have constructors. Another is perhaps that the creation of an object is in its own an action, not a reflexive concept, instead a file object is logically the result of opening a file      
have destructors, but since 
Inheritance (in C++ multiple inheritance is possible making it possible to create composites)

##The interface
In C++:
No interface - you inherit 

In C++
In many 
In C++ we have structs



###What is polymorphism?



#packages
Packages are not often looked at when considering OOD because it's assumed that they will naturally fall into place as you work but that's often not the case.
Let's consider this structrue of a package:

- pkg
|--animal
|--|-- 

This is not a powerful package.