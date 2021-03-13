## Cookies - writing and reading
---
## About Cookies
	A cookie is a file that is stored on the client's machine.
	Cookies are written by domains to store information on the client's machine.
	
	Limits
	We can store whatever information we would like in a cookie up to a certain size limit.
	The size limit of a cookie is dependent upon the browser but is usually around 4096 characters.
	There is also a limit to the number of cookies one domain can write.
	This limit is also browser specific.
	See this resource for more information.
	
	Expiring a cookie
	If the Expires or MaxAge field isn't set, then the cookie is deleted when the browser is closed.
	This is colloquially known as a "session cookie."
	You can expire a cookie by setting one of these two fields: Expires or MaxAge
	Expires sets the exact time when the cookie expires. Expires is Deprecated.
---
---
MaxAge sets how long the cookie should live (in seconds).