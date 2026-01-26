# RpcTradingAccountOperationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountHistoryId** | Pointer to **string** |  | [optional] 
**TradingAccountId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**OperationType** | Pointer to [**TradingAccountOperationType**](TradingAccountOperationType.md) |  | [optional] 
**OperateBy** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**PreviousValues** | Pointer to **map[string]interface{}** |  | [optional] 
**NewValues** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRpcTradingAccountOperationHistory

`func NewRpcTradingAccountOperationHistory() *RpcTradingAccountOperationHistory`

NewRpcTradingAccountOperationHistory instantiates a new RpcTradingAccountOperationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcTradingAccountOperationHistoryWithDefaults

`func NewRpcTradingAccountOperationHistoryWithDefaults() *RpcTradingAccountOperationHistory`

NewRpcTradingAccountOperationHistoryWithDefaults instantiates a new RpcTradingAccountOperationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountHistoryId

`func (o *RpcTradingAccountOperationHistory) GetTradingAccountHistoryId() string`

GetTradingAccountHistoryId returns the TradingAccountHistoryId field if non-nil, zero value otherwise.

### GetTradingAccountHistoryIdOk

`func (o *RpcTradingAccountOperationHistory) GetTradingAccountHistoryIdOk() (*string, bool)`

GetTradingAccountHistoryIdOk returns a tuple with the TradingAccountHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountHistoryId

`func (o *RpcTradingAccountOperationHistory) SetTradingAccountHistoryId(v string)`

SetTradingAccountHistoryId sets TradingAccountHistoryId field to given value.

### HasTradingAccountHistoryId

`func (o *RpcTradingAccountOperationHistory) HasTradingAccountHistoryId() bool`

HasTradingAccountHistoryId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcTradingAccountOperationHistory) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcTradingAccountOperationHistory) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcTradingAccountOperationHistory) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcTradingAccountOperationHistory) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetTimestamp

`func (o *RpcTradingAccountOperationHistory) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RpcTradingAccountOperationHistory) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RpcTradingAccountOperationHistory) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RpcTradingAccountOperationHistory) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOperationType

`func (o *RpcTradingAccountOperationHistory) GetOperationType() TradingAccountOperationType`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RpcTradingAccountOperationHistory) GetOperationTypeOk() (*TradingAccountOperationType, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RpcTradingAccountOperationHistory) SetOperationType(v TradingAccountOperationType)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *RpcTradingAccountOperationHistory) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetOperateBy

`func (o *RpcTradingAccountOperationHistory) GetOperateBy() string`

GetOperateBy returns the OperateBy field if non-nil, zero value otherwise.

### GetOperateByOk

`func (o *RpcTradingAccountOperationHistory) GetOperateByOk() (*string, bool)`

GetOperateByOk returns a tuple with the OperateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperateBy

`func (o *RpcTradingAccountOperationHistory) SetOperateBy(v string)`

SetOperateBy sets OperateBy field to given value.

### HasOperateBy

`func (o *RpcTradingAccountOperationHistory) HasOperateBy() bool`

HasOperateBy returns a boolean if a field has been set.

### GetTags

`func (o *RpcTradingAccountOperationHistory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RpcTradingAccountOperationHistory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RpcTradingAccountOperationHistory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RpcTradingAccountOperationHistory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPreviousValues

`func (o *RpcTradingAccountOperationHistory) GetPreviousValues() map[string]interface{}`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *RpcTradingAccountOperationHistory) GetPreviousValuesOk() (*map[string]interface{}, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *RpcTradingAccountOperationHistory) SetPreviousValues(v map[string]interface{})`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *RpcTradingAccountOperationHistory) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.

### GetNewValues

`func (o *RpcTradingAccountOperationHistory) GetNewValues() map[string]interface{}`

GetNewValues returns the NewValues field if non-nil, zero value otherwise.

### GetNewValuesOk

`func (o *RpcTradingAccountOperationHistory) GetNewValuesOk() (*map[string]interface{}, bool)`

GetNewValuesOk returns a tuple with the NewValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValues

`func (o *RpcTradingAccountOperationHistory) SetNewValues(v map[string]interface{})`

SetNewValues sets NewValues field to given value.

### HasNewValues

`func (o *RpcTradingAccountOperationHistory) HasNewValues() bool`

HasNewValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


