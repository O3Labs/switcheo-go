# Switcheo APIs wrapper written in Go

switcheo-go is a go client library for Switcheo APIs

## Installation

```sh
go get github.com/o3labs/switcheo-go
```

For details on all the functionality in this library, see the [GoDoc](https://godoc.org/github.com/O3Labs/switcheo-go/switcheo) documentation.


### Example usage
``` go
api := client.New(switcheo.TestNetAPI, httpClient, switcheo.UserAgent)
c := general.Client{API: *api}
response, err := c.GetTimestamp()
if err != nil {
	fmt.Printf("%v", err)
	return
}
log.Printf("%+v", response)
```

### Test
Each package contains test cases and including sample usage of apis