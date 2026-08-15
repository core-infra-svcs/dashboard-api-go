# InlineResponse200299Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | id | 
**NetworkName** | **string** | Name | 
**AlertCount** | **int32** | Total Alerts | 
**LastAlertedAt** | **time.Time** | Last time an alert was seen for this network | 
**SeverityCounts** | [**[]InlineResponse200299SeverityCounts**](InlineResponse200299SeverityCounts.md) | Alerts By Severity | 

## Methods

### NewInlineResponse200299Items

`func NewInlineResponse200299Items(networkId string, networkName string, alertCount int32, lastAlertedAt time.Time, severityCounts []InlineResponse200299SeverityCounts, ) *InlineResponse200299Items`

NewInlineResponse200299Items instantiates a new InlineResponse200299Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200299ItemsWithDefaults

`func NewInlineResponse200299ItemsWithDefaults() *InlineResponse200299Items`

NewInlineResponse200299ItemsWithDefaults instantiates a new InlineResponse200299Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200299Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200299Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200299Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse200299Items) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200299Items) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200299Items) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetAlertCount

`func (o *InlineResponse200299Items) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *InlineResponse200299Items) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *InlineResponse200299Items) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetLastAlertedAt

`func (o *InlineResponse200299Items) GetLastAlertedAt() time.Time`

GetLastAlertedAt returns the LastAlertedAt field if non-nil, zero value otherwise.

### GetLastAlertedAtOk

`func (o *InlineResponse200299Items) GetLastAlertedAtOk() (*time.Time, bool)`

GetLastAlertedAtOk returns a tuple with the LastAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAlertedAt

`func (o *InlineResponse200299Items) SetLastAlertedAt(v time.Time)`

SetLastAlertedAt sets LastAlertedAt field to given value.


### GetSeverityCounts

`func (o *InlineResponse200299Items) GetSeverityCounts() []InlineResponse200299SeverityCounts`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *InlineResponse200299Items) GetSeverityCountsOk() (*[]InlineResponse200299SeverityCounts, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *InlineResponse200299Items) SetSeverityCounts(v []InlineResponse200299SeverityCounts)`

SetSeverityCounts sets SeverityCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


