# sendy-data-ingester
Sendy (http://sendy.co) is aimed to handle big email lists but it lacks core functions to import such big lists into its web app. This tool helps to import big lists directly into its mysql database -- it sleeps X seconds every Y lines.

## Usage
```
GO111MODULE=on go build main.go
```

```
./main -config=config.toml -csv=sample.csv -log=s.log -line=500000 -sleep=3 -logtofile
```
