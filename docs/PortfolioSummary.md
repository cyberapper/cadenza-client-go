# PortfolioSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradingAccountId** | **string** | UUID string | 
**Currency** | **string** | Base currency for the portfolio summary | 
**Leverage** | **NullableInt32** | Leverage multiplier | 
**Equity** | **string** | Decimal value as string to preserve precision | 
**Margin** | **string** | Decimal value as string to preserve precision | 
**MarginLoan** | **string** | Decimal value as string to preserve precision | 
**MarginUsage** | **string** | Decimal value as string to preserve precision | 
**MarginRequirement** | **string** | Decimal value as string to preserve precision | 
**MarginLevel** | **string** | Decimal value as string to preserve precision | 
**Credit** | **string** | Decimal value as string to preserve precision | 
**RiskExposure** | **string** | Decimal value as string to preserve precision | 
**MaxRiskExposure** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 
**RiskExposureRate** | Pointer to **string** | Decimal value as string to preserve precision | [optional] 

## Methods

### NewPortfolioSummary

`func NewPortfolioSummary(tradingAccountId string, currency string, leverage NullableInt32, equity string, margin string, marginLoan string, marginUsage string, marginRequirement string, marginLevel string, credit string, riskExposure string, ) *PortfolioSummary`

NewPortfolioSummary instantiates a new PortfolioSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioSummaryWithDefaults

`func NewPortfolioSummaryWithDefaults() *PortfolioSummary`

NewPortfolioSummaryWithDefaults instantiates a new PortfolioSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradingAccountId

`func (o *PortfolioSummary) GetTradingAccountId() string`

GetTradingAccountId returns the TradingAccountId field if non-nil, zero value otherwise.

### GetTradingAccountIdOk

`func (o *PortfolioSummary) GetTradingAccountIdOk() (*string, bool)`

GetTradingAccountIdOk returns a tuple with the TradingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingAccountId

`func (o *PortfolioSummary) SetTradingAccountId(v string)`

SetTradingAccountId sets TradingAccountId field to given value.


### GetCurrency

`func (o *PortfolioSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PortfolioSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PortfolioSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLeverage

`func (o *PortfolioSummary) GetLeverage() int32`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *PortfolioSummary) GetLeverageOk() (*int32, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *PortfolioSummary) SetLeverage(v int32)`

SetLeverage sets Leverage field to given value.


### SetLeverageNil

`func (o *PortfolioSummary) SetLeverageNil(b bool)`

 SetLeverageNil sets the value for Leverage to be an explicit nil

### UnsetLeverage
`func (o *PortfolioSummary) UnsetLeverage()`

UnsetLeverage ensures that no value is present for Leverage, not even an explicit nil
### GetEquity

`func (o *PortfolioSummary) GetEquity() string`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *PortfolioSummary) GetEquityOk() (*string, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *PortfolioSummary) SetEquity(v string)`

SetEquity sets Equity field to given value.


### GetMargin

`func (o *PortfolioSummary) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *PortfolioSummary) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *PortfolioSummary) SetMargin(v string)`

SetMargin sets Margin field to given value.


### GetMarginLoan

`func (o *PortfolioSummary) GetMarginLoan() string`

GetMarginLoan returns the MarginLoan field if non-nil, zero value otherwise.

### GetMarginLoanOk

`func (o *PortfolioSummary) GetMarginLoanOk() (*string, bool)`

GetMarginLoanOk returns a tuple with the MarginLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLoan

`func (o *PortfolioSummary) SetMarginLoan(v string)`

SetMarginLoan sets MarginLoan field to given value.


### GetMarginUsage

`func (o *PortfolioSummary) GetMarginUsage() string`

GetMarginUsage returns the MarginUsage field if non-nil, zero value otherwise.

### GetMarginUsageOk

`func (o *PortfolioSummary) GetMarginUsageOk() (*string, bool)`

GetMarginUsageOk returns a tuple with the MarginUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginUsage

`func (o *PortfolioSummary) SetMarginUsage(v string)`

SetMarginUsage sets MarginUsage field to given value.


### GetMarginRequirement

`func (o *PortfolioSummary) GetMarginRequirement() string`

GetMarginRequirement returns the MarginRequirement field if non-nil, zero value otherwise.

### GetMarginRequirementOk

`func (o *PortfolioSummary) GetMarginRequirementOk() (*string, bool)`

GetMarginRequirementOk returns a tuple with the MarginRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRequirement

`func (o *PortfolioSummary) SetMarginRequirement(v string)`

SetMarginRequirement sets MarginRequirement field to given value.


### GetMarginLevel

`func (o *PortfolioSummary) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *PortfolioSummary) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *PortfolioSummary) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.


### GetCredit

`func (o *PortfolioSummary) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *PortfolioSummary) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *PortfolioSummary) SetCredit(v string)`

SetCredit sets Credit field to given value.


### GetRiskExposure

`func (o *PortfolioSummary) GetRiskExposure() string`

GetRiskExposure returns the RiskExposure field if non-nil, zero value otherwise.

### GetRiskExposureOk

`func (o *PortfolioSummary) GetRiskExposureOk() (*string, bool)`

GetRiskExposureOk returns a tuple with the RiskExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposure

`func (o *PortfolioSummary) SetRiskExposure(v string)`

SetRiskExposure sets RiskExposure field to given value.


### GetMaxRiskExposure

`func (o *PortfolioSummary) GetMaxRiskExposure() string`

GetMaxRiskExposure returns the MaxRiskExposure field if non-nil, zero value otherwise.

### GetMaxRiskExposureOk

`func (o *PortfolioSummary) GetMaxRiskExposureOk() (*string, bool)`

GetMaxRiskExposureOk returns a tuple with the MaxRiskExposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskExposure

`func (o *PortfolioSummary) SetMaxRiskExposure(v string)`

SetMaxRiskExposure sets MaxRiskExposure field to given value.

### HasMaxRiskExposure

`func (o *PortfolioSummary) HasMaxRiskExposure() bool`

HasMaxRiskExposure returns a boolean if a field has been set.

### GetRiskExposureRate

`func (o *PortfolioSummary) GetRiskExposureRate() string`

GetRiskExposureRate returns the RiskExposureRate field if non-nil, zero value otherwise.

### GetRiskExposureRateOk

`func (o *PortfolioSummary) GetRiskExposureRateOk() (*string, bool)`

GetRiskExposureRateOk returns a tuple with the RiskExposureRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposureRate

`func (o *PortfolioSummary) SetRiskExposureRate(v string)`

SetRiskExposureRate sets RiskExposureRate field to given value.

### HasRiskExposureRate

`func (o *PortfolioSummary) HasRiskExposureRate() bool`

HasRiskExposureRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


