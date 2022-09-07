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

For a full description of how roman numerals work, take a look at the [Wikipedia article](https://en.wikipedia.org/wiki/Roman_numerals).

---

## Running the project

Install Go, clone the repo and run the application.
### Running the program

```bash
# In the repo directory
go run .
```

### Running the tests


```bash
# In the repo directory
go test -v .
```
---
## Functionality supported

Supports roman numbers to X̅ (10000). Runs for bigger numbers also but does not use the large numeral symbols after X̅.

---
## Changelog

- 2022-09-07: Refactor the logic to use a map instead of if-else statements. Add support for numbers up to 10000.
- 2022-09-03: First kata workout done
