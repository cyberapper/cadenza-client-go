# RpcSaveTickersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tickers** | [**[]Ticker**](Ticker.md) |  | 

## Methods

### NewRpcSaveTickersParams

`func NewRpcSaveTickersParams(tickers []Ticker, ) *RpcSaveTickersParams`

NewRpcSaveTickersParams instantiates a new RpcSaveTickersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveTickersParamsWithDefaults

`func NewRpcSaveTickersParamsWithDefaults() *RpcSaveTickersParams`

NewRpcSaveTickersParamsWithDefaults instantiates a new RpcSaveTickersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTickers

`func (o *RpcSaveTickersParams) GetTickers() []Ticker`

GetTickers returns the Tickers field if non-nil, zero value otherwise.

### GetTickersOk

`func (o *RpcSaveTickersParams) GetTickersOk() (*[]Ticker, bool)`

GetTickersOk returns a tuple with the Tickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickers

`func (o *RpcSaveTickersParams) SetTickers(v []Ticker)`

SetTickers sets Tickers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


