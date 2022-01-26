module github.com/rajathagasthya/gomod-retract-experiment

go 1.17

retract (
    [v1.4.0-pre-alpha-1, v1.4.0-pre-alpha-3]
)

require github.com/google/go-cmp v0.5.7

require golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
