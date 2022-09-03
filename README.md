# Roman numerals kata with Golang in TDD
> Learning and practicing TDD and Go

---
## Problem Description
Write a function to convert from normal numbers to Roman Numerals: e.g.

``` 
1 --> I
10 --> X
7 --> VII
```

For a full description of how it works, take a look at the [Wikipedia article](https://en.wikipedia.org/wiki/Roman_numerals).

---

## Running the project

Install Go, clone the repo and run the application.
### Running the program

```bash
# In the repo directory
go run .
```

### Running the tests

Coverage 65,4% as of 2022-09-03

```bash
# In the repo directory
go test .
```
---
## Functionality supported

Supports roman numbers to V̅ (5000).

Uses for loop to check how big the number is and constructs the result by adding the corresponding roman number to the result.

---
## Changelog

- 2022-09-03: First kata workout done
