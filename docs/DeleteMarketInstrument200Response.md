# DeleteMarketInstrument200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteMarketInstrument200Response

`func NewDeleteMarketInstrument200Response(success bool, errno int32, error_ NullableString, ) *DeleteMarketInstrument200Response`

NewDeleteMarketInstrument200Response instantiates a new DeleteMarketInstrument200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMarketInstrument200ResponseWithDefaults

`func NewDeleteMarketInstrument200ResponseWithDefaults() *DeleteMarketInstrument200Response`

NewDeleteMarketInstrument200ResponseWithDefaults instantiates a new DeleteMarketInstrument200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteMarketInstrument200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteMarketInstrument200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteMarketInstrument200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *DeleteMarketInstrument200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *DeleteMarketInstrument200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *DeleteMarketInstrument200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *DeleteMarketInstrument200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteMarketInstrument200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteMarketInstrument200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *DeleteMarketInstrument200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DeleteMarketInstrument200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetData

`func (o *DeleteMarketInstrument200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteMarketInstrument200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteMarketInstrument200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *DeleteMarketInstrument200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


