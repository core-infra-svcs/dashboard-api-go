# InlineResponse200238Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | id | 
**NetworkName** | **string** | Name | 
**AlertCount** | **int32** | Total Alerts | 
**SeverityCounts** | [**[]InlineResponse200238SeverityCounts**](InlineResponse200238SeverityCounts.md) | Alerts By Severity | 

## Methods

### NewInlineResponse200238Items

`func NewInlineResponse200238Items(networkId string, networkName string, alertCount int32, severityCounts []InlineResponse200238SeverityCounts, ) *InlineResponse200238Items`

NewInlineResponse200238Items instantiates a new InlineResponse200238Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200238ItemsWithDefaults

`func NewInlineResponse200238ItemsWithDefaults() *InlineResponse200238Items`

NewInlineResponse200238ItemsWithDefaults instantiates a new InlineResponse200238Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200238Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200238Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200238Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse200238Items) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200238Items) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200238Items) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetAlertCount

`func (o *InlineResponse200238Items) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *InlineResponse200238Items) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *InlineResponse200238Items) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetSeverityCounts

`func (o *InlineResponse200238Items) GetSeverityCounts() []InlineResponse200238SeverityCounts`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *InlineResponse200238Items) GetSeverityCountsOk() (*[]InlineResponse200238SeverityCounts, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *InlineResponse200238Items) SetSeverityCounts(v []InlineResponse200238SeverityCounts)`

SetSeverityCounts sets SeverityCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


