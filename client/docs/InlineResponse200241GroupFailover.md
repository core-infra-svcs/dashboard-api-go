# InlineResponse200241GroupFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectToInternet** | Pointer to **bool** | [optional] When both primary and backup tunnels are down, direct traffic to the internet. Traffic will be routed via the WAN | [optional] 

## Methods

### NewInlineResponse200241GroupFailover

`func NewInlineResponse200241GroupFailover() *InlineResponse200241GroupFailover`

NewInlineResponse200241GroupFailover instantiates a new InlineResponse200241GroupFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200241GroupFailoverWithDefaults

`func NewInlineResponse200241GroupFailoverWithDefaults() *InlineResponse200241GroupFailover`

NewInlineResponse200241GroupFailoverWithDefaults instantiates a new InlineResponse200241GroupFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectToInternet

`func (o *InlineResponse200241GroupFailover) GetDirectToInternet() bool`

GetDirectToInternet returns the DirectToInternet field if non-nil, zero value otherwise.

### GetDirectToInternetOk

`func (o *InlineResponse200241GroupFailover) GetDirectToInternetOk() (*bool, bool)`

GetDirectToInternetOk returns a tuple with the DirectToInternet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectToInternet

`func (o *InlineResponse200241GroupFailover) SetDirectToInternet(v bool)`

SetDirectToInternet sets DirectToInternet field to given value.

### HasDirectToInternet

`func (o *InlineResponse200241GroupFailover) HasDirectToInternet() bool`

HasDirectToInternet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


