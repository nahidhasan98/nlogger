# A Custom Logger

**`nlogger`** is a custom logger package to make a log. The logs will be written inside the `log` directory. Here is a sample log file text.

```json
2021.08.13 20:24:08 [ Error ]		strconv.Atoi: parsing "ab": invalid syntax - 2021.08.13 20:24:08
2021.08.13 20:24:08 [ Warn ]		This is a sample warning message! - 2021.08.13 20:24:08
2021.08.13 20:24:08 [ Update ]		Product - pr1031: Product updated. - 2021.08.13 20:24:08
2021.08.13 20:24:08 [ Delete ]		Order - or4116: Order deleted. - 2021.08.13 20:24:08
2021.08.13 20:24:08 [ Endpoint ]	/api/order/delete/3 - 127.0.0.1 - 2021.08.13 20:24:08

```


## Install

```bash
go get github.com/nahidhasan98/nlogger
```


## Examples

Let's start by instantiating a new logger service.

```go
import "github.com/nahidhasan98/nlogger"

func main() {
    //to write log in a default log file
    cLog := nlogger.NewLogger()

    //to write log in a provided file named file
    cLog := nlogger.NewLoggerName("mylog.log")
}
```

#### 1. Error Log
`Error()` creates an error log describing the provided error and time.


```go
func main() {
    //writing an error log
    cLog.Error(err, time.Now())
}
```

#### 2. Warn Log
`Warn()` creates an warning log describing the provided warning message and time.


```go
func main() {
    //writing an warning log
    cLog.Warn("This is a sample warning message!", time.Now())
}
```

#### 3. Update Log
`Update()` creates a log for an updated item. Here `entity` indicates what type of item is updated such as: user, product, cart etc. And `id` indicates the id of updated item and `msg` for providing a custom message.


```go
func main() {
    //writing an update log
    cLog.Update("Product", "pr1031", "Product updated.", time.Now())
}
```

#### 4. Delete Log
`Delete()` creates a log for a deleted item. Here `entity` indicates what type of item is deleted such as: user, product, cart etc. And `id` indicates the id of deleted item and `msg` for providing a custom message.


```go
func main() {
    //writing a delete log
    cLog.Delete("Order", "or4116", "Order deleted.", time.Now())
}
```

#### 5. Endpoint Log
`Endpoint()` creates a log describing the provided `endpoint` that is visited by a client. It also mention the client's `ip` address and the visiting time.


```go
func main() {
    //writing an endpoint log
    cLog.Endpoint("/api/order/delete/3", "127.0.0.1", time.Now())
}
```