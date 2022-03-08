# Go SDK for CWB (台灣中央氣象局開放資料平臺)

## Usage

### Install

```
go get github.com/fatsoap/cwb-sdk-go
```

```go
import cwb "github.com/fatsoap/cwb-sdk-go"
client, _ := cwb.New(KEY)
w, err := client.GeneralWeatherForecast36()
if err != nil {
    log.Fatal("Some Thing Went Wrong")
}
log.Print(w)
```

## Develope

### .env

```
API_KEY=YOUR_CWB_API_KEY
```
