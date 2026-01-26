# RpcSyncInstrumentsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RpcSyncInstrumentsResultData**](RpcSyncInstrumentsResultData.md) |  | [optional] 
**Error** | Pointer to [**RpcError**](RpcError.md) |  | [optional] 

## Methods

### NewRpcSyncInstrumentsResult

`func NewRpcSyncInstrumentsResult() *RpcSyncInstrumentsResult`

NewRpcSyncInstrumentsResult instantiates a new RpcSyncInstrumentsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcSyncInstrumentsResultWithDefaults

`func NewRpcSyncInstrumentsResultWithDefaults() *RpcSyncInstrumentsResult`

NewRpcSyncInstrumentsResultWithDefaults instantiates a new RpcSyncInstrumentsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RpcSyncInstrumentsResult) GetData() RpcSyncInstrumentsResultData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RpcSyncInstrumentsResult) GetDataOk() (*RpcSyncInstrumentsResultData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RpcSyncInstrumentsResult) SetData(v RpcSyncInstrumentsResultData)`

SetData sets Data field to given value.

### HasData

`func (o *RpcSyncInstrumentsResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *RpcSyncInstrumentsResult) GetError() RpcError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RpcSyncInstrumentsResult) GetErrorOk() (*RpcError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RpcSyncInstrumentsResult) SetError(v RpcError)`

SetError sets Error field to given value.

### HasError

`func (o *RpcSyncInstrumentsResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


