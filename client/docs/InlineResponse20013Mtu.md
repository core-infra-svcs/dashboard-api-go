# InlineResponse20013Mtu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **NullableInt32** | The MTU limit for the uplink interface, in bytes, from 1280 through 1500. A null value clears the configured limit. | 

## Methods

### NewInlineResponse20013Mtu

`func NewInlineResponse20013Mtu(limit NullableInt32, ) *InlineResponse20013Mtu`

NewInlineResponse20013Mtu instantiates a new InlineResponse20013Mtu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20013MtuWithDefaults

`func NewInlineResponse20013MtuWithDefaults() *InlineResponse20013Mtu`

NewInlineResponse20013MtuWithDefaults instantiates a new InlineResponse20013Mtu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *InlineResponse20013Mtu) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineResponse20013Mtu) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineResponse20013Mtu) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *InlineResponse20013Mtu) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *InlineResponse20013Mtu) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


