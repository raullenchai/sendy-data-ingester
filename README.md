# sendy-data-ingester
Sendy is aimed to handle big email lists but it lacks the core functions to import such big lists into its web app. This tool means to help to import big lists into its mysql database directly.

## Usage
```
GO111MODULE=on go build main.go
```

```
./main -config=config.toml -csv=sample.csv -log=s.log -line=500000 -sleep=3 -logtofile
```
