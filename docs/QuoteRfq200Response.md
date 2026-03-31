# QuoteRfq200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Indicates if the operation was successful | [optional] 
**Errno** | **int32** | Error code (0 for success, non-zero indicates error). Format: AABBB where AA is the module code and BBB is the error code | 
**Error** | Pointer to **NullableString** | Error message (null for successful operations) | [optional] 
**Details** | Pointer to [**NullableBaseResponseDetails**](BaseResponseDetails.md) |  | [optional] 
**Data** | Pointer to [**Quote**](Quote.md) |  | [optional] 

## Methods

### NewQuoteRfq200Response

`func NewQuoteRfq200Response(errno int32, ) *QuoteRfq200Response`

NewQuoteRfq200Response instantiates a new QuoteRfq200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRfq200ResponseWithDefaults

`func NewQuoteRfq200ResponseWithDefaults() *QuoteRfq200Response`

NewQuoteRfq200ResponseWithDefaults instantiates a new QuoteRfq200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *QuoteRfq200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *QuoteRfq200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *QuoteRfq200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *QuoteRfq200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrno

`func (o *QuoteRfq200Response) GetErrno() int32`

GetErrno returns the Errno field if non-nil, zero value otherwise.

### GetErrnoOk

`func (o *QuoteRfq200Response) GetErrnoOk() (*int32, bool)`

GetErrnoOk returns a tuple with the Errno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrno

`func (o *QuoteRfq200Response) SetErrno(v int32)`

SetErrno sets Errno field to given value.


### GetError

`func (o *QuoteRfq200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QuoteRfq200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QuoteRfq200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QuoteRfq200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QuoteRfq200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QuoteRfq200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetDetails

`func (o *QuoteRfq200Response) GetDetails() BaseResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *QuoteRfq200Response) GetDetailsOk() (*BaseResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *QuoteRfq200Response) SetDetails(v BaseResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *QuoteRfq200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *QuoteRfq200Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *QuoteRfq200Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetData

`func (o *QuoteRfq200Response) GetData() Quote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuoteRfq200Response) GetDataOk() (*Quote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuoteRfq200Response) SetData(v Quote)`

SetData sets Data field to given value.

### HasData

`func (o *QuoteRfq200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


