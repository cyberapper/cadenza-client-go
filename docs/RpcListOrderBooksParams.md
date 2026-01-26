# RpcListOrderBooksParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentIds** | Pointer to **[]string** | List of instrument IDs | [optional] 
**Venue** | Pointer to **string** | Filter by venue | [optional] 
**Symbols** | Pointer to **[]string** | List of symbols | [optional] 
**Depth** | Pointer to **int32** | Order book depth | [optional] [default to 10]

## Methods

### NewRpcListOrderBooksParams

`func NewRpcListOrderBooksParams() *RpcListOrderBooksParams`

NewRpcListOrderBooksParams instantiates a new RpcListOrderBooksParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcListOrderBooksParamsWithDefaults

`func NewRpcListOrderBooksParamsWithDefaults() *RpcListOrderBooksParams`

NewRpcListOrderBooksParamsWithDefaults instantiates a new RpcListOrderBooksParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrumentIds

`func (o *RpcListOrderBooksParams) GetInstrumentIds() []string`

GetInstrumentIds returns the InstrumentIds field if non-nil, zero value otherwise.

### GetInstrumentIdsOk

`func (o *RpcListOrderBooksParams) GetInstrumentIdsOk() (*[]string, bool)`

GetInstrumentIdsOk returns a tuple with the InstrumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentIds

`func (o *RpcListOrderBooksParams) SetInstrumentIds(v []string)`

SetInstrumentIds sets InstrumentIds field to given value.

### HasInstrumentIds

`func (o *RpcListOrderBooksParams) HasInstrumentIds() bool`

HasInstrumentIds returns a boolean if a field has been set.

### GetVenue

`func (o *RpcListOrderBooksParams) GetVenue() string`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *RpcListOrderBooksParams) GetVenueOk() (*string, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *RpcListOrderBooksParams) SetVenue(v string)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *RpcListOrderBooksParams) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

### GetSymbols

`func (o *RpcListOrderBooksParams) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *RpcListOrderBooksParams) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *RpcListOrderBooksParams) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *RpcListOrderBooksParams) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetDepth

`func (o *RpcListOrderBooksParams) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *RpcListOrderBooksParams) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *RpcListOrderBooksParams) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *RpcListOrderBooksParams) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


