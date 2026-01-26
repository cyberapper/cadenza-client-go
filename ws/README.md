# WebSocket Client (ws)

This package re-exports the [Centrifugo Go SDK](https://github.com/centrifugal/centrifuge-go) for WebSocket connections to the Cadenza API.

## Installation

```go
import "github.com/cyberapper/cadenza-client-go/ws"
```

## Usage

### Connect to WebSocket

```go
client := ws.NewJsonClient("ws://localhost:8001/connection/websocket", ws.Config{
    Token: "your-jwt-token",
})

// Set up event handlers
client.OnConnected(func(e ws.ConnectedEvent) {
    fmt.Println("Connected to server")
})

client.OnDisconnected(func(e ws.DisconnectedEvent) {
    fmt.Printf("Disconnected: %s\n", e.Reason)
})

client.OnError(func(e ws.ErrorEvent) {
    fmt.Printf("Error: %v\n", e.Error)
})

// Connect
err := client.Connect()
if err != nil {
    log.Fatal(err)
}
defer client.Disconnect()
```

### RPC Calls

```go
// Send RPC request
data, _ := json.Marshal(map[string]interface{}{
    "venue":   "binance",
    "symbols": []string{"BTCUSDT"},
})

result, err := client.RPC(ctx, "query.instrument.list", data)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(result.Data))
```

### Subscribe to Channels

```go
// Create subscription
sub, err := client.NewSubscription("market:orderbook:binance:BTCUSDT")
if err != nil {
    log.Fatal(err)
}

// Handle publications
sub.OnPublication(func(e ws.PublicationEvent) {
    fmt.Printf("Received: %s\n", string(e.Data))
})

sub.OnSubscribed(func(e ws.SubscribedEvent) {
    fmt.Println("Subscribed successfully")
})

// Subscribe
err = sub.Subscribe()
if err != nil {
    log.Fatal(err)
}
```

## Exported Types

### Client & Configuration
- `Client` - WebSocket client
- `Config` - Client configuration
- `Subscription` - Channel subscription
- `SubscriptionConfig` - Subscription configuration

### Events
- `ConnectedEvent`, `ConnectingEvent`, `DisconnectedEvent`
- `PublicationEvent`, `JoinEvent`, `LeaveEvent`
- `SubscribedEvent`, `SubscribingEvent`, `UnsubscribedEvent`
- `ErrorEvent`, `MessageEvent`
- `ServerPublicationEvent`, `ServerJoinEvent`, `ServerLeaveEvent`

### Handlers
- `ConnectedHandler`, `ConnectingHandler`, `DisconnectHandler`
- `PublicationHandler`, `JoinHandler`, `LeaveHandler`
- `SubscribedHandler`, `SubscribingHandler`, `UnsubscribedHandler`
- `ErrorHandler`, `MessageHandler`

### Results
- `RPCResult` - RPC call result
- `PublishResult` - Publish result
- `HistoryResult` - Channel history
- `PresenceResult`, `PresenceStatsResult` - Presence info

### Errors
- `Error`, `ConfigurationError`, `ConnectError`
- `RefreshError`, `TransportError`
- `SubscriptionRefreshError`, `SubscriptionSubscribeError`

### Constants
- `StateDisconnected`, `StateConnecting`, `StateConnected`
- `SubStateUnsubscribed`, `SubStateSubscribing`, `SubStateSubscribed`
- `LogLevelNone`, `LogLevelDebug`, `LogLevelTrace`

## References

- [Centrifugo Documentation](https://centrifugal.dev/)
- [centrifuge-go SDK](https://github.com/centrifugal/centrifuge-go)
