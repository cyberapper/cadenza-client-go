# ConnectTradingAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialIds** | **[]string** | A list of credential IDs to be used to connect the trading account | 
**ExternalTradingAccountId** | **string** | External trading account ID | 
**Nickname** | Pointer to **string** | Nickname of the trading account | [optional] 

## Methods

### NewConnectTradingAccountRequest

`func NewConnectTradingAccountRequest(credentialIds []string, externalTradingAccountId string, ) *ConnectTradingAccountRequest`

NewConnectTradingAccountRequest instantiates a new ConnectTradingAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectTradingAccountRequestWithDefaults

`func NewConnectTradingAccountRequestWithDefaults() *ConnectTradingAccountRequest`

NewConnectTradingAccountRequestWithDefaults instantiates a new ConnectTradingAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialIds

`func (o *ConnectTradingAccountRequest) GetCredentialIds() []string`

GetCredentialIds returns the CredentialIds field if non-nil, zero value otherwise.

### GetCredentialIdsOk

`func (o *ConnectTradingAccountRequest) GetCredentialIdsOk() (*[]string, bool)`

GetCredentialIdsOk returns a tuple with the CredentialIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIds

`func (o *ConnectTradingAccountRequest) SetCredentialIds(v []string)`

SetCredentialIds sets CredentialIds field to given value.


### GetExternalTradingAccountId

`func (o *ConnectTradingAccountRequest) GetExternalTradingAccountId() string`

GetExternalTradingAccountId returns the ExternalTradingAccountId field if non-nil, zero value otherwise.

### GetExternalTradingAccountIdOk

`func (o *ConnectTradingAccountRequest) GetExternalTradingAccountIdOk() (*string, bool)`

GetExternalTradingAccountIdOk returns a tuple with the ExternalTradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTradingAccountId

`func (o *ConnectTradingAccountRequest) SetExternalTradingAccountId(v string)`

SetExternalTradingAccountId sets ExternalTradingAccountId field to given value.


### GetNickname

`func (o *ConnectTradingAccountRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ConnectTradingAccountRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ConnectTradingAccountRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ConnectTradingAccountRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


