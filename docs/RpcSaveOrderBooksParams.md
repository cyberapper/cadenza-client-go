# RpcSaveOrderBooksParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderBooks** | [**[]RpcOrderBook**](RpcOrderBook.md) |  | 

## Methods

### NewRpcSaveOrderBooksParams

`func NewRpcSaveOrderBooksParams(orderBooks []RpcOrderBook, ) *RpcSaveOrderBooksParams`

NewRpcSaveOrderBooksParams instantiates a new RpcSaveOrderBooksParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSaveOrderBooksParamsWithDefaults

`func NewRpcSaveOrderBooksParamsWithDefaults() *RpcSaveOrderBooksParams`

NewRpcSaveOrderBooksParamsWithDefaults instantiates a new RpcSaveOrderBooksParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderBooks

`func (o *RpcSaveOrderBooksParams) GetOrderBooks() []RpcOrderBook`

GetOrderBooks returns the OrderBooks field if non-nil, zero value otherwise.

### GetOrderBooksOk

`func (o *RpcSaveOrderBooksParams) GetOrderBooksOk() (*[]RpcOrderBook, bool)`

GetOrderBooksOk returns a tuple with the OrderBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBooks

`func (o *RpcSaveOrderBooksParams) SetOrderBooks(v []RpcOrderBook)`

SetOrderBooks sets OrderBooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


