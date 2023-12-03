# go-quiz
Basic timed quiz app where questions and answers are parsed from a CSV.


## Usage
**Default (30 second timer, problems.csv)**
```
go build . && ./go-quiz
```

**Custom Timer**
```
go build . && ./go-quiz -t <seconds>
```

**Custom CSV**
```
go build . && ./go-quiz -csv <file>
```