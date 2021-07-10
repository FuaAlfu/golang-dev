---
stack: GO
lang: All
---

## Timeouts
The network is unreliable and the standard library default clients and servers do not set their main timeouts, and all of them interpret the zero value as infinity to boot. Timeouts are subjective to the use case and the Go core team have steered clear of making any sweeping generalisations.