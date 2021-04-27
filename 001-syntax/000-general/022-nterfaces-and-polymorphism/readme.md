---
stack: Go
lag: Interfaces & polymorphism
---

## Interfaces & polymorphism
In Go, values can be of more than one type. An interface allows a value to be of more than one type. We create an interface using this syntax: “keyword identifier type” so for an interface it would be: “type human interface” We then define which method(s) a type must have to implement that interface. If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}), then that TYPE implicitly implements the interface and is also of that interface type. In Go, values can be of more than one 