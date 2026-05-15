# InlineResponse20063Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]InlineResponse20063FiltersConditions**](InlineResponse20063FiltersConditions.md) | Conditions | [optional] 
**FailureType** | Pointer to **string** | Failure Type | [optional] 
**LookbackWindow** | Pointer to **int32** | Loopback Window (in sec) | [optional] 
**MinDuration** | Pointer to **int32** | Min Duration | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Period** | Pointer to **int32** | Period | [optional] 
**Priority** | Pointer to **string** | Priority | [optional] 
**Regex** | Pointer to **string** | Regex | [optional] 
**Selector** | Pointer to **string** | Selector | [optional] 
**Serials** | Pointer to **[]string** | Serials | [optional] 
**SsidNum** | Pointer to **int32** | SSID Number | [optional] 
**Tag** | Pointer to **string** | Tag | [optional] 
**Threshold** | Pointer to **int32** | Threshold | [optional] 
**Timeout** | Pointer to **int32** | Timeout | [optional] 

## Methods

### NewInlineResponse20063Filters

`func NewInlineResponse20063Filters() *InlineResponse20063Filters`

NewInlineResponse20063Filters instantiates a new InlineResponse20063Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20063FiltersWithDefaults

`func NewInlineResponse20063FiltersWithDefaults() *InlineResponse20063Filters`

NewInlineResponse20063FiltersWithDefaults instantiates a new InlineResponse20063Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *InlineResponse20063Filters) GetConditions() []InlineResponse20063FiltersConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InlineResponse20063Filters) GetConditionsOk() (*[]InlineResponse20063FiltersConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InlineResponse20063Filters) SetConditions(v []InlineResponse20063FiltersConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InlineResponse20063Filters) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFailureType

`func (o *InlineResponse20063Filters) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *InlineResponse20063Filters) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *InlineResponse20063Filters) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *InlineResponse20063Filters) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### GetLookbackWindow

`func (o *InlineResponse20063Filters) GetLookbackWindow() int32`

GetLookbackWindow returns the LookbackWindow field if non-nil, zero value otherwise.

### GetLookbackWindowOk

`func (o *InlineResponse20063Filters) GetLookbackWindowOk() (*int32, bool)`

GetLookbackWindowOk returns a tuple with the LookbackWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackWindow

`func (o *InlineResponse20063Filters) SetLookbackWindow(v int32)`

SetLookbackWindow sets LookbackWindow field to given value.

### HasLookbackWindow

`func (o *InlineResponse20063Filters) HasLookbackWindow() bool`

HasLookbackWindow returns a boolean if a field has been set.

### GetMinDuration

`func (o *InlineResponse20063Filters) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *InlineResponse20063Filters) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *InlineResponse20063Filters) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *InlineResponse20063Filters) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20063Filters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20063Filters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20063Filters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20063Filters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeriod

`func (o *InlineResponse20063Filters) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InlineResponse20063Filters) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InlineResponse20063Filters) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *InlineResponse20063Filters) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse20063Filters) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse20063Filters) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse20063Filters) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse20063Filters) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRegex

`func (o *InlineResponse20063Filters) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *InlineResponse20063Filters) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *InlineResponse20063Filters) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *InlineResponse20063Filters) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetSelector

`func (o *InlineResponse20063Filters) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *InlineResponse20063Filters) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *InlineResponse20063Filters) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *InlineResponse20063Filters) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSerials

`func (o *InlineResponse20063Filters) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineResponse20063Filters) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineResponse20063Filters) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineResponse20063Filters) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSsidNum

`func (o *InlineResponse20063Filters) GetSsidNum() int32`

GetSsidNum returns the SsidNum field if non-nil, zero value otherwise.

### GetSsidNumOk

`func (o *InlineResponse20063Filters) GetSsidNumOk() (*int32, bool)`

GetSsidNumOk returns a tuple with the SsidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNum

`func (o *InlineResponse20063Filters) SetSsidNum(v int32)`

SetSsidNum sets SsidNum field to given value.

### HasSsidNum

`func (o *InlineResponse20063Filters) HasSsidNum() bool`

HasSsidNum returns a boolean if a field has been set.

### GetTag

`func (o *InlineResponse20063Filters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *InlineResponse20063Filters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *InlineResponse20063Filters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *InlineResponse20063Filters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThreshold

`func (o *InlineResponse20063Filters) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InlineResponse20063Filters) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InlineResponse20063Filters) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *InlineResponse20063Filters) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse20063Filters) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse20063Filters) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse20063Filters) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse20063Filters) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


