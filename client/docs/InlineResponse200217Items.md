# InlineResponse200217Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | id | 
**NetworkName** | **string** | Name | 
**AlertCount** | **int32** | Total Alerts | 
**SeverityCounts** | [**[]InlineResponse200217SeverityCounts**](InlineResponse200217SeverityCounts.md) | Alerts By Severity | 

## Methods

### NewInlineResponse200217Items

`func NewInlineResponse200217Items(networkId string, networkName string, alertCount int32, severityCounts []InlineResponse200217SeverityCounts, ) *InlineResponse200217Items`

NewInlineResponse200217Items instantiates a new InlineResponse200217Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200217ItemsWithDefaults

`func NewInlineResponse200217ItemsWithDefaults() *InlineResponse200217Items`

NewInlineResponse200217ItemsWithDefaults instantiates a new InlineResponse200217Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse200217Items) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200217Items) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200217Items) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse200217Items) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse200217Items) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse200217Items) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetAlertCount

`func (o *InlineResponse200217Items) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *InlineResponse200217Items) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *InlineResponse200217Items) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetSeverityCounts

`func (o *InlineResponse200217Items) GetSeverityCounts() []InlineResponse200217SeverityCounts`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *InlineResponse200217Items) GetSeverityCountsOk() (*[]InlineResponse200217SeverityCounts, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *InlineResponse200217Items) SetSeverityCounts(v []InlineResponse200217SeverityCounts)`

SetSeverityCounts sets SeverityCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


