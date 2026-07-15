# InlineResponse200297Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | id | 
**NetworkName** | **string** | Name | 
**AlertCount** | **int32** | Total Alerts | 
**LastAlertedAt** | **time.Time** | Last time an alert was seen for this network | 
**SeverityCounts** | [**[]InlineResponse200297SeverityCounts**](InlineResponse200297SeverityCounts.md) | Alerts By Severity | 

## Methods

### NewInlineResponse200297Items

`func NewInlineResponse200297Items(networkId string, networkName string, alertCount int32, lastAlertedAt time.Time, severityCounts []InlineResponse200297SeverityCounts, ) *InlineResponse200297Items`

NewInlineResponse200297Items instantiates a new InlineResponse200297Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200297ItemsWithDefaults

`func NewInlineResponse200297ItemsWithDefaults() *InlineResponse200297Items`

NewInlineResponse200297ItemsWithDefaults instantiates a new InlineResponse200297Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200297Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200297Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200297Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse200297Items) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200297Items) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200297Items) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetAlertCount

`func (o *InlineResponse200297Items) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *InlineResponse200297Items) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *InlineResponse200297Items) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetLastAlertedAt

`func (o *InlineResponse200297Items) GetLastAlertedAt() time.Time`

GetLastAlertedAt returns the LastAlertedAt field if non-nil, zero value otherwise.

### GetLastAlertedAtOk

`func (o *InlineResponse200297Items) GetLastAlertedAtOk() (*time.Time, bool)`

GetLastAlertedAtOk returns a tuple with the LastAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAlertedAt

`func (o *InlineResponse200297Items) SetLastAlertedAt(v time.Time)`

SetLastAlertedAt sets LastAlertedAt field to given value.


### GetSeverityCounts

`func (o *InlineResponse200297Items) GetSeverityCounts() []InlineResponse200297SeverityCounts`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *InlineResponse200297Items) GetSeverityCountsOk() (*[]InlineResponse200297SeverityCounts, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *InlineResponse200297Items) SetSeverityCounts(v []InlineResponse200297SeverityCounts)`

SetSeverityCounts sets SeverityCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


