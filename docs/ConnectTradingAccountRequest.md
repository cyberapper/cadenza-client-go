# ConnectTradingAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Venue** | Pointer to [**Venue**](Venue.md) |  | [optional] 
**CredentialIds** | Pointer to **[]string** | Credential IDs for exchange venues. Not required for Fermata. | [optional] 
**ExternalTradingAccountId** | Pointer to **string** | External trading account ID. Not required for Fermata. | [optional] 
**DealerAccountId** | Pointer to **string** | UUID string | [optional] 
**Nickname** | Pointer to **string** | Nickname of the trading account | [optional] 

## Methods

### NewConnectTradingAccountRequest

`func NewConnectTradingAccountRequest() *ConnectTradingAccountRequest`

NewConnectTradingAccountRequest instantiates a new ConnectTradingAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectTradingAccountRequestWithDefaults

`func NewConnectTradingAccountRequestWithDefaults() *ConnectTradingAccountRequest`

NewConnectTradingAccountRequestWithDefaults instantiates a new ConnectTradingAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVenue

`func (o *ConnectTradingAccountRequest) GetVenue() Venue`

GetVenue returns the Venue field if non-nil, zero value otherwise.

### GetVenueOk

`func (o *ConnectTradingAccountRequest) GetVenueOk() (*Venue, bool)`

GetVenueOk returns a tuple with the Venue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenue

`func (o *ConnectTradingAccountRequest) SetVenue(v Venue)`

SetVenue sets Venue field to given value.

### HasVenue

`func (o *ConnectTradingAccountRequest) HasVenue() bool`

HasVenue returns a boolean if a field has been set.

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

### HasCredentialIds

`func (o *ConnectTradingAccountRequest) HasCredentialIds() bool`

HasCredentialIds returns a boolean if a field has been set.

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

### HasExternalTradingAccountId

`func (o *ConnectTradingAccountRequest) HasExternalTradingAccountId() bool`

HasExternalTradingAccountId returns a boolean if a field has been set.

### GetDealerAccountId

`func (o *ConnectTradingAccountRequest) GetDealerAccountId() string`

GetDealerAccountId returns the DealerAccountId field if non-nil, zero value otherwise.

### GetDealerAccountIdOk

`func (o *ConnectTradingAccountRequest) GetDealerAccountIdOk() (*string, bool)`

GetDealerAccountIdOk returns a tuple with the DealerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealerAccountId

`func (o *ConnectTradingAccountRequest) SetDealerAccountId(v string)`

SetDealerAccountId sets DealerAccountId field to given value.

### HasDealerAccountId

`func (o *ConnectTradingAccountRequest) HasDealerAccountId() bool`

HasDealerAccountId returns a boolean if a field has been set.

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


