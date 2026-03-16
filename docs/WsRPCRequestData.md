# WsRPCRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeOrder** | [**RpcTradeOrder**](RpcTradeOrder.md) |  | 
**TradingAccountId** | **string** | Filter by trading account ID | 
**IdempotencyKey** | Pointer to **string** | Idempotency key to prevent duplicate orders | [optional] 
**AwaitClosed** | Pointer to **bool** | Wait for order to reach terminal state before responding | [optional] [default to false]
**TradeOrderId** | **string** | Filter by specific trade order ID | 
**ClientOrderId** | Pointer to **string** | Client order ID | [optional] 
**ExternalOrderId** | Pointer to **string** | External order ID | [optional] 
**InstrumentId** | **[]string** |  | 
**Side** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Status** | Pointer to [**CredentialStatus**](CredentialStatus.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | Filter orders created after this time | [optional] 
**EndTime** | Pointer to **time.Time** | Filter orders created before this time | [optional] 
**Pagination** | Pointer to [**RpcPagination**](RpcPagination.md) |  | [optional] 
**TradingAccount** | [**RpcTradingAccount**](RpcTradingAccount.md) |  | 
**TradingAccounts** | [**[]RpcTradingAccount**](RpcTradingAccount.md) |  | 
**Venue** | [**Venue**](Venue.md) |  | 
**SubscriptionType** | [**SubscriptionType**](SubscriptionType.md) |  | 
**SubscriptionId** | Pointer to **string** | Subscription ID to cancel | [optional] 
**CredentialType** | [**CredentialType**](CredentialType.md) |  | 
**ApiKey** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SecretPassphrase** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Credential** | [**RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | 
**CredentialIds** | Pointer to **[]string** |  | [optional] 
**Credentials** | Pointer to [**[]RpcTradingAccountCredential**](RpcTradingAccountCredential.md) |  | [optional] 
**CredentialId** | **string** |  | 
**Currency** | Pointer to **string** | Filter by currency | [optional] 
**Portfolio** | [**RpcPortfolio**](RpcPortfolio.md) |  | 
**Symbols** | Pointer to **[]string** |  | [optional] 
**Instruments** | **[]string** |  | 
**SecurityType** | Pointer to **string** |  | [optional] 
**InstrumentStatus** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Securities** | **[]string** |  | 
**SecurityIds** | **[]string** |  | 
**ExternalSymbols** | Pointer to **[]string** |  | [optional] 
**Symbol** | Pointer to **string** | Symbol (alternative to instrumentId) | [optional] 
**Depth** | Pointer to **int32** | Order book depth | [optional] [default to 10]
**InstrumentIds** | Pointer to **[]string** |  | [optional] 
**OrderBooks** | [**[]RpcOrderBook**](RpcOrderBook.md) |  | 
**Klines** | [**[]RpcKline**](RpcKline.md) |  | 
**Interval** | Pointer to **string** |  | [optional] 
**Tickers** | [**[]RpcTicker**](RpcTicker.md) |  | 

## Methods

### NewWsRPCRequestData

`func NewWsRPCRequestData(tradeOrder RpcTradeOrder, tradingAccountId string, tradeOrderId string, instrumentId []string, tradingAccount RpcTradingAccount, tradingAccounts []RpcTradingAccount, venue Venue, subscriptionType SubscriptionType, credentialType CredentialType, credential RpcTradingAccountCredential, credentialId string, portfolio RpcPortfolio, instruments []string, securities []string, securityIds []string, orderBooks []RpcOrderBook, klines []RpcKline, tickers []RpcTicker, ) *WsRPCRequestData`

NewWsRPCRequestData instantiates a new WsRPCRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsRPCRequestDataWithDefaults

`func NewWsRPCRequestDataWithDefaults() *WsRPCRequestData`

NewWsRPCRequestDataWithDefaults instantiates a new WsRPCRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeOrder

`func (o *WsRPCRequestData) GetTradeOrder() RpcTradeOrder`

GetTradeOrder returns the TradeOrder field if non-nil, zero value otherwise.

### GetTradeOrderOk

`func (o *WsRPCRequestData) GetTradeOrderOk() (*RpcTradeOrder, bool)`

GetTradeOrderOk returns a tuple with the TradeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrder

`func (o *WsRPCRequestData) SetTradeOrder(v RpcTradeOrder)`

SetTradeOrder sets TradeOrder field to given value.


### GetTradingAccountId

`func (o *WsRPCRequestData) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *WsRPCRequestData) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *WsRPCRequestData) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetIdempotencyKey

`func (o *WsRPCRequestData) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WsRPCRequestData) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WsRPCRequestData) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *WsRPCRequestData) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetAwaitClosed

`func (o *WsRPCRequestData) GetAwaitClosed() bool`

GetAwaitClosed returns the AwaitClosed field if non-nil, zero value otherwise.

### GetAwaitClosedOk

`func (o *WsRPCRequestData) GetAwaitClosedOk() (*bool, bool)`

GetAwaitClosedOk returns a tuple with the AwaitClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitClosed

`func (o *WsRPCRequestData) SetAwaitClosed(v bool)`

SetAwaitClosed sets AwaitClosed field to given value.

### HasAwaitClosed

`func (o *WsRPCRequestData) HasAwaitClosed() bool`

HasAwaitClosed returns a boolean if a field has been set.

### GetTradeOrderId

`func (o *WsRPCRequestData) GetTradeOrderId() string`

GetTradeOrderId returns the TradeOrderId field if non-nil, zero value otherwise.

### GetTradeOrderIdOk

`func (o *WsRPCRequestData) GetTradeOrderIdOk() (*string, bool)`

GetTradeOrderIdOk returns a tuple with the TradeOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeOrderId

`func (o *WsRPCRequestData) SetTradeOrderId(v string)`

SetTradeOrderId sets TradeOrderId field to given value.


### GetClientOrderId

`func (o *WsRPCRequestData) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *WsRPCRequestData) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *WsRPCRequestData) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *WsRPCRequestData) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *WsRPCRequestData) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *WsRPCRequestData) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *WsRPCRequestData) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *WsRPCRequestData) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.

### GetInstrumentId

`func (o *WsRPCRequestData) GetInstrumentId() []string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *WsRPCRequestData) GetInstrumentIdOk() (*[]string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *WsRPCRequestData) SetInstrumentId(v []string)`

SetInstrumentId sets InstrumentId field to given value.


### GetSide

`func (o *WsRPCRequestData) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *WsRPCRequestData) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *WsRPCRequestData) SetSide(v OrderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *WsRPCRequestData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *WsRPCRequestData) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *WsRPCRequestData) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *WsRPCRequestData) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *WsRPCRequestData) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetStatus

`func (o *WsRPCRequestData) GetStatus() CredentialStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WsRPCRequestData) GetStatusOk() (*CredentialStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WsRPCRequestData) SetStatus(v CredentialStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WsRPCRequestData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *WsRPCRequestData) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WsRPCRequestData) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WsRPCRequestData) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WsRPCRequestData) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WsRPCRequestData) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WsRPCRequestData) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WsRPCRequestData) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WsRPCRequestData) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPagination

`func (o *WsRPCRequestData) GetPagination() RpcPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *WsRPCRequestData) GetPaginationOk() (*RpcPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *WsRPCRequestData) SetPagination(v RpcPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *WsRPCRequestData) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTradingAccount

`func (o *WsRPCRequestData) GetTradingAccount() RpcTradingAccount`

GetTradingAccount returns the TradingAccount field if non-nil, zero value otherwise.

### GetTradingAccountOk

`func (o *WsRPCRequestData) GetTradingAccountOk() (*RpcTradingAccount, bool)`

GetTradingAccountOk returns a tuple with the TradingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccount

`func (o *WsRPCRequestData) SetTradingAccount(v RpcTradingAccount)`

SetTradingAccount sets TradingAccount field to given value.


### GetTradingAccounts

`func (o *WsRPCRequestData) GetTradingAccounts() []RpcTradingAccount`

GetTradingAccounts returns the TradingAccounts field if non-nil, zero value otherwise.

### GetTradingAccountsOk

`func (o *WsRPCRequestData) GetTradingAccountsOk() (*[]RpcTradingAccount, bool)`

GetTradingAccountsOk returns a tuple with the TradingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccounts

`func (o *WsRPCRequestData) SetTradingAccounts(v []RpcTradingAccount)`

SetTradingAccounts sets TradingAccounts field to given value.


### GetVenue

`func (o *WsRPCRequestData) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *WsRPCRequestData) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *WsRPCRequestData) SetVenue(v Venue)`

SetVenue sets Venue field to given value.


### GetSubscriptionType

`func (o *WsRPCRequestData) GetSubscriptionType() SubscriptionType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *WsRPCRequestData) GetSubscriptionTypeOk() (*SubscriptionType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *WsRPCRequestData) SetSubscriptionType(v SubscriptionType)`

SetSubscriptionType sets SubscriptionType field to given value.


### GetSubscriptionId

`func (o *WsRPCRequestData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *WsRPCRequestData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *WsRPCRequestData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *WsRPCRequestData) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetCredentialType

`func (o *WsRPCRequestData) GetCredentialType() CredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *WsRPCRequestData) GetCredentialTypeOk() (*CredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *WsRPCRequestData) SetCredentialType(v CredentialType)`

SetCredentialType sets CredentialType field to given value.


### GetApiKey

`func (o *WsRPCRequestData) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *WsRPCRequestData) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *WsRPCRequestData) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *WsRPCRequestData) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *WsRPCRequestData) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *WsRPCRequestData) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *WsRPCRequestData) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *WsRPCRequestData) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSecretPassphrase

`func (o *WsRPCRequestData) GetSecretPassphrase() string`

GetSecretPassphrase returns the SecretPassphrase field if non-nil, zero value otherwise.

### GetSecretPassphraseOk

`func (o *WsRPCRequestData) GetSecretPassphraseOk() (*string, bool)`

GetSecretPassphraseOk returns a tuple with the SecretPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPassphrase

`func (o *WsRPCRequestData) SetSecretPassphrase(v string)`

SetSecretPassphrase sets SecretPassphrase field to given value.

### HasSecretPassphrase

`func (o *WsRPCRequestData) HasSecretPassphrase() bool`

HasSecretPassphrase returns a boolean if a field has been set.

### GetNickname

`func (o *WsRPCRequestData) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *WsRPCRequestData) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *WsRPCRequestData) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *WsRPCRequestData) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetCredential

`func (o *WsRPCRequestData) GetCredential() RpcTradingAccountCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *WsRPCRequestData) GetCredentialOk() (*RpcTradingAccountCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *WsRPCRequestData) SetCredential(v RpcTradingAccountCredential)`

SetCredential sets Credential field to given value.


### GetCredentialIds

`func (o *WsRPCRequestData) GetCredentialIds() []string`

GetCredentialIds returns the CredentialIds field if non-nil, zero value otherwise.

### GetCredentialIdsOk

`func (o *WsRPCRequestData) GetCredentialIdsOk() (*[]string, bool)`

GetCredentialIdsOk returns a tuple with the CredentialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIds

`func (o *WsRPCRequestData) SetCredentialIds(v []string)`

SetCredentialIds sets CredentialIds field to given value.

### HasCredentialIds

`func (o *WsRPCRequestData) HasCredentialIds() bool`

HasCredentialIds returns a boolean if a field has been set.

### GetCredentials

`func (o *WsRPCRequestData) GetCredentials() []RpcTradingAccountCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *WsRPCRequestData) GetCredentialsOk() (*[]RpcTradingAccountCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *WsRPCRequestData) SetCredentials(v []RpcTradingAccountCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *WsRPCRequestData) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCredentialId

`func (o *WsRPCRequestData) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *WsRPCRequestData) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *WsRPCRequestData) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.


### GetCurrency

`func (o *WsRPCRequestData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WsRPCRequestData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WsRPCRequestData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WsRPCRequestData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPortfolio

`func (o *WsRPCRequestData) GetPortfolio() RpcPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *WsRPCRequestData) GetPortfolioOk() (*RpcPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *WsRPCRequestData) SetPortfolio(v RpcPortfolio)`

SetPortfolio sets Portfolio field to given value.


### GetSymbols

`func (o *WsRPCRequestData) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *WsRPCRequestData) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *WsRPCRequestData) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *WsRPCRequestData) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetInstruments

`func (o *WsRPCRequestData) GetInstruments() []string`

GetInstruments returns the Instruments field if non-nil, zero value otherwise.

### GetInstrumentsOk

`func (o *WsRPCRequestData) GetInstrumentsOk() (*[]string, bool)`

GetInstrumentsOk returns a tuple with the Instruments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruments

`func (o *WsRPCRequestData) SetInstruments(v []string)`

SetInstruments sets Instruments field to given value.


### GetSecurityType

`func (o *WsRPCRequestData) GetSecurityType() string`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *WsRPCRequestData) GetSecurityTypeOk() (*string, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *WsRPCRequestData) SetSecurityType(v string)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *WsRPCRequestData) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetInstrumentStatus

`func (o *WsRPCRequestData) GetInstrumentStatus() string`

GetInstrumentStatus returns the InstrumentStatus field if non-nil, zero value otherwise.

### GetInstrumentStatusOk

`func (o *WsRPCRequestData) GetInstrumentStatusOk() (*string, bool)`

GetInstrumentStatusOk returns a tuple with the InstrumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentStatus

`func (o *WsRPCRequestData) SetInstrumentStatus(v string)`

SetInstrumentStatus sets InstrumentStatus field to given value.

### HasInstrumentStatus

`func (o *WsRPCRequestData) HasInstrumentStatus() bool`

HasInstrumentStatus returns a boolean if a field has been set.

### GetLimit

`func (o *WsRPCRequestData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WsRPCRequestData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WsRPCRequestData) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WsRPCRequestData) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *WsRPCRequestData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WsRPCRequestData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WsRPCRequestData) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WsRPCRequestData) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSecurities

`func (o *WsRPCRequestData) GetSecurities() []string`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *WsRPCRequestData) GetSecuritiesOk() (*[]string, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *WsRPCRequestData) SetSecurities(v []string)`

SetSecurities sets Securities field to given value.


### GetSecurityIds

`func (o *WsRPCRequestData) GetSecurityIds() []string`

GetSecurityIds returns the SecurityIds field if non-nil, zero value otherwise.

### GetSecurityIdsOk

`func (o *WsRPCRequestData) GetSecurityIdsOk() (*[]string, bool)`

GetSecurityIdsOk returns a tuple with the SecurityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIds

`func (o *WsRPCRequestData) SetSecurityIds(v []string)`

SetSecurityIds sets SecurityIds field to given value.


### GetExternalSymbols

`func (o *WsRPCRequestData) GetExternalSymbols() []string`

GetExternalSymbols returns the ExternalSymbols field if non-nil, zero value otherwise.

### GetExternalSymbolsOk

`func (o *WsRPCRequestData) GetExternalSymbolsOk() (*[]string, bool)`

GetExternalSymbolsOk returns a tuple with the ExternalSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSymbols

`func (o *WsRPCRequestData) SetExternalSymbols(v []string)`

SetExternalSymbols sets ExternalSymbols field to given value.

### HasExternalSymbols

`func (o *WsRPCRequestData) HasExternalSymbols() bool`

HasExternalSymbols returns a boolean if a field has been set.

### GetSymbol

`func (o *WsRPCRequestData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *WsRPCRequestData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *WsRPCRequestData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *WsRPCRequestData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDepth

`func (o *WsRPCRequestData) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WsRPCRequestData) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WsRPCRequestData) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *WsRPCRequestData) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetInstrumentIds

`func (o *WsRPCRequestData) GetInstrumentIds() []string`

GetInstrumentIds returns the InstrumentIds field if non-nil, zero value otherwise.

### GetInstrumentIdsOk

`func (o *WsRPCRequestData) GetInstrumentIdsOk() (*[]string, bool)`

GetInstrumentIdsOk returns a tuple with the InstrumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentIds

`func (o *WsRPCRequestData) SetInstrumentIds(v []string)`

SetInstrumentIds sets InstrumentIds field to given value.

### HasInstrumentIds

`func (o *WsRPCRequestData) HasInstrumentIds() bool`

HasInstrumentIds returns a boolean if a field has been set.

### GetOrderBooks

`func (o *WsRPCRequestData) GetOrderBooks() []RpcOrderBook`

GetOrderBooks returns the OrderBooks field if non-nil, zero value otherwise.

### GetOrderBooksOk

`func (o *WsRPCRequestData) GetOrderBooksOk() (*[]RpcOrderBook, bool)`

GetOrderBooksOk returns a tuple with the OrderBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBooks

`func (o *WsRPCRequestData) SetOrderBooks(v []RpcOrderBook)`

SetOrderBooks sets OrderBooks field to given value.


### GetKlines

`func (o *WsRPCRequestData) GetKlines() []RpcKline`

GetKlines returns the Klines field if non-nil, zero value otherwise.

### GetKlinesOk

`func (o *WsRPCRequestData) GetKlinesOk() (*[]RpcKline, bool)`

GetKlinesOk returns a tuple with the Klines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKlines

`func (o *WsRPCRequestData) SetKlines(v []RpcKline)`

SetKlines sets Klines field to given value.


### GetInterval

`func (o *WsRPCRequestData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *WsRPCRequestData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *WsRPCRequestData) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *WsRPCRequestData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTickers

`func (o *WsRPCRequestData) GetTickers() []RpcTicker`

GetTickers returns the Tickers field if non-nil, zero value otherwise.

### GetTickersOk

`func (o *WsRPCRequestData) GetTickersOk() (*[]RpcTicker, bool)`

GetTickersOk returns a tuple with the Tickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickers

`func (o *WsRPCRequestData) SetTickers(v []RpcTicker)`

SetTickers sets Tickers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


