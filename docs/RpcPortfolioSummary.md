# RpcPortfolioSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortfolioSummaryId** | Pointer to **string** |  | [optional] 
**TradingAccountId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | Summary currency | [optional] 
**Leverage** | Pointer to **int32** |  | [optional] 
**Equity** | Pointer to **string** | Total equity | [optional] 
**Margin** | Pointer to **string** | Margin collateral | [optional] 
**MarginLoan** | Pointer to **string** | Borrowed margin | [optional] 
**MarginUsage** | Pointer to **string** | Margin usage rate | [optional] 
**MarginRequirement** | Pointer to **string** | Required margin | [optional] 
**MarginLevel** | Pointer to **string** | Margin level | [optional] 
**Credit** | Pointer to **string** | Available credit | [optional] 
**RiskExposure** | Pointer to **string** | Total risk exposure | [optional] 
**MaxRiskExposure** | Pointer to **string** | Maximum risk exposure | [optional] 
**RiskExposureRate** | Pointer to **string** | Risk exposure rate | [optional] 

## Methods

### NewRpcPortfolioSummary

`func NewRpcPortfolioSummary() *RpcPortfolioSummary`

NewRpcPortfolioSummary instantiates a new RpcPortfolioSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpcPortfolioSummaryWithDefaults

`func NewRpcPortfolioSummaryWithDefaults() *RpcPortfolioSummary`

NewRpcPortfolioSummaryWithDefaults instantiates a new RpcPortfolioSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortfolioSummaryId

`func (o *RpcPortfolioSummary) GetPortfolioSummaryId() string`

GetPortfolioSummaryId returns the PortfolioSummaryId field if non-nil, zero value otherwise.

### GetPortfolioSummaryIdOk

`func (o *RpcPortfolioSummary) GetPortfolioSummaryIdOk() (*string, bool)`

GetPortfolioSummaryIdOk returns a tuple with the PortfolioSummaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolioSummaryId

`func (o *RpcPortfolioSummary) SetPortfolioSummaryId(v string)`

SetPortfolioSummaryId sets PortfolioSummaryId field to given value.

### HasPortfolioSummaryId

`func (o *RpcPortfolioSummary) HasPortfolioSummaryId() bool`

HasPortfolioSummaryId returns a boolean if a field has been set.

### GetTradingAccountId

`func (o *RpcPortfolioSummary) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *RpcPortfolioSummary) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *RpcPortfolioSummary) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.

### HasTradingAccountId

`func (o *RpcPortfolioSummary) HasTradingAccountId() bool`

HasTradingAccountId returns a boolean if a field has been set.

### GetCurrency

`func (o *RpcPortfolioSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RpcPortfolioSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RpcPortfolioSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RpcPortfolioSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLeverage

`func (o *RpcPortfolioSummary) GetLeverage() int32`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *RpcPortfolioSummary) GetLeverageOk() (*int32, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *RpcPortfolioSummary) SetLeverage(v int32)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *RpcPortfolioSummary) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetEquity

`func (o *RpcPortfolioSummary) GetEquity() string`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *RpcPortfolioSummary) GetEquityOk() (*string, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *RpcPortfolioSummary) SetEquity(v string)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *RpcPortfolioSummary) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetMargin

`func (o *RpcPortfolioSummary) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *RpcPortfolioSummary) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *RpcPortfolioSummary) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *RpcPortfolioSummary) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarginLoan

`func (o *RpcPortfolioSummary) GetMarginLoan() string`

GetMarginLoan returns the MarginLoan field if non-nil, zero value otherwise.

### GetMarginLoanOk

`func (o *RpcPortfolioSummary) GetMarginLoanOk() (*string, bool)`

GetMarginLoanOk returns a tuple with the MarginLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLoan

`func (o *RpcPortfolioSummary) SetMarginLoan(v string)`

SetMarginLoan sets MarginLoan field to given value.

### HasMarginLoan

`func (o *RpcPortfolioSummary) HasMarginLoan() bool`

HasMarginLoan returns a boolean if a field has been set.

### GetMarginUsage

`func (o *RpcPortfolioSummary) GetMarginUsage() string`

GetMarginUsage returns the MarginUsage field if non-nil, zero value otherwise.

### GetMarginUsageOk

`func (o *RpcPortfolioSummary) GetMarginUsageOk() (*string, bool)`

GetMarginUsageOk returns a tuple with the MarginUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginUsage

`func (o *RpcPortfolioSummary) SetMarginUsage(v string)`

SetMarginUsage sets MarginUsage field to given value.

### HasMarginUsage

`func (o *RpcPortfolioSummary) HasMarginUsage() bool`

HasMarginUsage returns a boolean if a field has been set.

### GetMarginRequirement

`func (o *RpcPortfolioSummary) GetMarginRequirement() string`

GetMarginRequirement returns the MarginRequirement field if non-nil, zero value otherwise.

### GetMarginRequirementOk

`func (o *RpcPortfolioSummary) GetMarginRequirementOk() (*string, bool)`

GetMarginRequirementOk returns a tuple with the MarginRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRequirement

`func (o *RpcPortfolioSummary) SetMarginRequirement(v string)`

SetMarginRequirement sets MarginRequirement field to given value.

### HasMarginRequirement

`func (o *RpcPortfolioSummary) HasMarginRequirement() bool`

HasMarginRequirement returns a boolean if a field has been set.

### GetMarginLevel

`func (o *RpcPortfolioSummary) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *RpcPortfolioSummary) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *RpcPortfolioSummary) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *RpcPortfolioSummary) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetCredit

`func (o *RpcPortfolioSummary) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *RpcPortfolioSummary) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *RpcPortfolioSummary) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *RpcPortfolioSummary) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetRiskExposure

`func (o *RpcPortfolioSummary) GetRiskExposure() string`

GetRiskExposure returns the RiskExposure field if non-nil, zero value otherwise.

### GetRiskExposureOk

`func (o *RpcPortfolioSummary) GetRiskExposureOk() (*string, bool)`

GetRiskExposureOk returns a tuple with the RiskExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposure

`func (o *RpcPortfolioSummary) SetRiskExposure(v string)`

SetRiskExposure sets RiskExposure field to given value.

### HasRiskExposure

`func (o *RpcPortfolioSummary) HasRiskExposure() bool`

HasRiskExposure returns a boolean if a field has been set.

### GetMaxRiskExposure

`func (o *RpcPortfolioSummary) GetMaxRiskExposure() string`

GetMaxRiskExposure returns the MaxRiskExposure field if non-nil, zero value otherwise.

### GetMaxRiskExposureOk

`func (o *RpcPortfolioSummary) GetMaxRiskExposureOk() (*string, bool)`

GetMaxRiskExposureOk returns a tuple with the MaxRiskExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskExposure

`func (o *RpcPortfolioSummary) SetMaxRiskExposure(v string)`

SetMaxRiskExposure sets MaxRiskExposure field to given value.

### HasMaxRiskExposure

`func (o *RpcPortfolioSummary) HasMaxRiskExposure() bool`

HasMaxRiskExposure returns a boolean if a field has been set.

### GetRiskExposureRate

`func (o *RpcPortfolioSummary) GetRiskExposureRate() string`

GetRiskExposureRate returns the RiskExposureRate field if non-nil, zero value otherwise.

### GetRiskExposureRateOk

`func (o *RpcPortfolioSummary) GetRiskExposureRateOk() (*string, bool)`

GetRiskExposureRateOk returns a tuple with the RiskExposureRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposureRate

`func (o *RpcPortfolioSummary) SetRiskExposureRate(v string)`

SetRiskExposureRate sets RiskExposureRate field to given value.

### HasRiskExposureRate

`func (o *RpcPortfolioSummary) HasRiskExposureRate() bool`

HasRiskExposureRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


