// Package ws re-exports the Centrifugo client SDK for WebSocket connections.
//
// Usage:
//
//	import "github.com/cyberapper/cadenza-client-go/ws"
//
//	client := ws.NewJsonClient(wsURL, ws.Config{Token: token})
//	client.Connect()
package ws

import "github.com/centrifugal/centrifuge-go"

// Client and configuration types
type (
	Client             = centrifuge.Client
	Config             = centrifuge.Config
	Subscription       = centrifuge.Subscription
	SubscriptionConfig = centrifuge.SubscriptionConfig
	ClientInfo         = centrifuge.ClientInfo
	StreamPosition     = centrifuge.StreamPosition
	Publication        = centrifuge.Publication
)

// Event types
type (
	PublicationEvent        = centrifuge.PublicationEvent
	JoinEvent               = centrifuge.JoinEvent
	LeaveEvent              = centrifuge.LeaveEvent
	SubscribedEvent         = centrifuge.SubscribedEvent
	UnsubscribedEvent       = centrifuge.UnsubscribedEvent
	SubscribingEvent        = centrifuge.SubscribingEvent
	SubscriptionErrorEvent  = centrifuge.SubscriptionErrorEvent
	SubscriptionTokenEvent  = centrifuge.SubscriptionTokenEvent
	ConnectedEvent          = centrifuge.ConnectedEvent
	ConnectingEvent         = centrifuge.ConnectingEvent
	DisconnectedEvent       = centrifuge.DisconnectedEvent
	ErrorEvent              = centrifuge.ErrorEvent
	MessageEvent            = centrifuge.MessageEvent
	ConnectionTokenEvent    = centrifuge.ConnectionTokenEvent
	ServerPublicationEvent  = centrifuge.ServerPublicationEvent
	ServerJoinEvent         = centrifuge.ServerJoinEvent
	ServerLeaveEvent        = centrifuge.ServerLeaveEvent
	ServerSubscribedEvent   = centrifuge.ServerSubscribedEvent
	ServerUnsubscribedEvent = centrifuge.ServerUnsubscribedEvent
	ServerSubscribingEvent  = centrifuge.ServerSubscribingEvent
)

// Handler types
type (
	PublicationHandler          = centrifuge.PublicationHandler
	JoinHandler                 = centrifuge.JoinHandler
	LeaveHandler                = centrifuge.LeaveHandler
	SubscribedHandler           = centrifuge.SubscribedHandler
	UnsubscribedHandler         = centrifuge.UnsubscribedHandler
	SubscribingHandler          = centrifuge.SubscribingHandler
	SubscriptionErrorHandler    = centrifuge.SubscriptionErrorHandler
	ConnectedHandler            = centrifuge.ConnectedHandler
	ConnectingHandler           = centrifuge.ConnectingHandler
	DisconnectHandler           = centrifuge.DisconnectHandler
	ErrorHandler                = centrifuge.ErrorHandler
	MessageHandler              = centrifuge.MessageHandler
	ServerPublicationHandler    = centrifuge.ServerPublicationHandler
	ServerJoinHandler           = centrifuge.ServerJoinHandler
	ServerLeaveHandler          = centrifuge.ServerLeaveHandler
	ServerSubscribedHandler     = centrifuge.ServerSubscribedHandler
	ServerUnsubscribedHandler   = centrifuge.ServerUnsubscribedHandler
	ServerSubscribingHandler    = centrifuge.ServerSubscribingHandler
	LogHandler                  = centrifuge.LogHandler
)

// Result types
type (
	RPCResult           = centrifuge.RPCResult
	PublishResult       = centrifuge.PublishResult
	HistoryResult       = centrifuge.HistoryResult
	PresenceResult      = centrifuge.PresenceResult
	PresenceStatsResult = centrifuge.PresenceStatsResult
	PresenceStats       = centrifuge.PresenceStats
	UnsubscribeResult   = centrifuge.UnsubscribeResult
)

// Error types
type (
	Error                        = centrifuge.Error
	ConfigurationError           = centrifuge.ConfigurationError
	ConnectError                 = centrifuge.ConnectError
	RefreshError                 = centrifuge.RefreshError
	TransportError               = centrifuge.TransportError
	SubscriptionRefreshError     = centrifuge.SubscriptionRefreshError
	SubscriptionSubscribeError   = centrifuge.SubscriptionSubscribeError
)

// Other types
type (
	LogEntry  = centrifuge.LogEntry
	LogLevel  = centrifuge.LogLevel
	State     = centrifuge.State
	SubState  = centrifuge.SubState
	DeltaType = centrifuge.DeltaType
	HistoryOption  = centrifuge.HistoryOption
	HistoryOptions = centrifuge.HistoryOptions
)

// Constructors
var (
	NewJsonClient     = centrifuge.NewJsonClient
	NewProtobufClient = centrifuge.NewProtobufClient
)

// History options
var (
	WithHistoryLimit   = centrifuge.WithHistoryLimit
	WithHistoryReverse = centrifuge.WithHistoryReverse
	WithHistorySince   = centrifuge.WithHistorySince
)

// Constants
const (
	LogLevelNone  = centrifuge.LogLevelNone
	LogLevelDebug = centrifuge.LogLevelDebug
	LogLevelTrace = centrifuge.LogLevelTrace
)

const (
	StateDisconnected = centrifuge.StateDisconnected
	StateConnecting   = centrifuge.StateConnecting
	StateConnected    = centrifuge.StateConnected
	StateClosed       = centrifuge.StateClosed
)

const (
	SubStateUnsubscribed = centrifuge.SubStateUnsubscribed
	SubStateSubscribing  = centrifuge.SubStateSubscribing
	SubStateSubscribed   = centrifuge.SubStateSubscribed
)

const (
	DeltaTypeNone   = centrifuge.DeltaTypeNone
	DeltaTypeFossil = centrifuge.DeltaTypeFossil
)

// Errors
var (
	ErrTimeout = centrifuge.ErrTimeout
)
