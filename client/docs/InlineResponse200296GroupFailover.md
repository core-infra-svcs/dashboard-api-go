# InlineResponse200296GroupFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectToInternet** | Pointer to **bool** | [optional] When both primary and backup tunnels are down, direct traffic to the internet. Traffic will be routed via the WAN | [optional] 

## Methods

### NewInlineResponse200296GroupFailover

`func NewInlineResponse200296GroupFailover() *InlineResponse200296GroupFailover`

NewInlineResponse200296GroupFailover instantiates a new InlineResponse200296GroupFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200296GroupFailoverWithDefaults

`func NewInlineResponse200296GroupFailoverWithDefaults() *InlineResponse200296GroupFailover`

NewInlineResponse200296GroupFailoverWithDefaults instantiates a new InlineResponse200296GroupFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectToInternet

`func (o *InlineResponse200296GroupFailover) GetDirectToInternet() bool`

GetDirectToInternet returns the DirectToInternet field if non-nil, zero value otherwise.

### GetDirectToInternetOk

`func (o *InlineResponse200296GroupFailover) GetDirectToInternetOk() (*bool, bool)`

GetDirectToInternetOk returns a tuple with the DirectToInternet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectToInternet

`func (o *InlineResponse200296GroupFailover) SetDirectToInternet(v bool)`

SetDirectToInternet sets DirectToInternet field to given value.

### HasDirectToInternet

`func (o *InlineResponse200296GroupFailover) HasDirectToInternet() bool`

HasDirectToInternet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


