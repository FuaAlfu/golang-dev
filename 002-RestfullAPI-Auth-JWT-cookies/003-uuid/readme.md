## UUID - Universally unique identifier 
---
---
A universally unique identifier (UUID) is an identifier standard used in software construction. A UUID is simply a 128-bit value. 
The meaning of each bit is defined by any of several variants. 
For human-readable display, many systems use a canonical format using hexadecimal text with inserted hyphen characters.

For example: 123e4567-e89b-12d3-a456-426655440000 The intent of UUIDs is to enable distributed systems to uniquely identify information without significant central coordination. 
In this context the word unique should be taken to mean "practically unique" rather than "guaranteed unique". 
Since the identifiers have a finite size, it is possible for two differing items to share the same identifier. 
This is a form of hash collision. 
The identifier size and generation process need to be selected so as to make this sufficiently improbable in practice. 
Anyone can create a UUID and use it to identify something with reasonable confidence that the same identifier will never be unintentionally created by anyone to identify something else. Information labeled with UUIDs can therefore be later combined into a single database without needing to resolve identifier (ID) conflicts. Adoption of UUIDs is widespread with many computing platforms providing support for generating UUIDs and for parsing/generating their textual representation. 
Only after generating 1 billion UUIDs every second for the next 100 years would the probability of creating just one duplicate would be about 50%.
---