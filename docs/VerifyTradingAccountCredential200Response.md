# VerifyTradingAccountCredential200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates if the operation was successful | 
**Errno** | **int32** | Error code (0 for success, negative for errors) | 
**Error** | **NullableString** | Error message (null for successful operations) | 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**[]TradingAccount**](TradingAccount.md) |  | [optional] 

## Methods

### NewVerifyTradingAccountCredential200Response

`func NewVerifyTradingAccountCredential200Response(success bool, errno int32, error_ NullableString, ) *VerifyTradingAccountCredential200Response`

NewVerifyTradingAccountCredential200Response instantiates a new VerifyTradingAccountCredential200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyTradingAccountCredential200ResponseWithDefaults

`func NewVerifyTradingAccountCredential200ResponseWithDefaults() *VerifyTradingAccountCredential200Response`

NewVerifyTradingAccountCredential200ResponseWithDefaults instantiates a new VerifyTradingAccountCredential200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *VerifyTradingAccountCredential200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *VerifyTradingAccountCredential200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *VerifyTradingAccountCredential200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrno

`func (o *VerifyTradingAccountCredential200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *VerifyTradingAccountCredential200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *VerifyTradingAccountCredential200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *VerifyTradingAccountCredential200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VerifyTradingAccountCredential200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VerifyTradingAccountCredential200Response) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *VerifyTradingAccountCredential200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *VerifyTradingAccountCredential200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *VerifyTradingAccountCredential200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerifyTradingAccountCredential200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerifyTradingAccountCredential200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerifyTradingAccountCredential200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *VerifyTradingAccountCredential200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *VerifyTradingAccountCredential200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *VerifyTradingAccountCredential200Response) GetData() []TradingAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VerifyTradingAccountCredential200Response) GetDataOk() (*[]TradingAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VerifyTradingAccountCredential200Response) SetData(v []TradingAccount)`

SetData sets Data field to given value.

### HasData

`func (o *VerifyTradingAccountCredential200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


