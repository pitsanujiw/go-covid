# Simple JSON API to summarize COVID-19 stats using this public API

## How to get dependencies

```bash
 go mod tidy
```

## How to run service

```bash
 make covid-serv
```

## How to run test

```bash
 make test
```

## How to use

After run a local service, you can use a api on browser or postman

```bash
curl http://localhost:5000/covid/summary
```

a result example

```json
{
    "province": {
        "Samut Sakhon": 3613,
        "Bangkok": 2774
    },
    "ageGroup": {
        "0-30": 300,
        "31-60": 150,
        "61+": 250,
        "N/A": 4
    }
}
```

Assignment by LMWN
