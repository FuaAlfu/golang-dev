## DEFER
is used to ensure that a function call is performed later in program's 
excution, usually for purposess of clean up.
---
	• A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
---