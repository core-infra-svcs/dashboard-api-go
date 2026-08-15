# InlineResponse200488SsidRadiusPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failover** | Pointer to [**InlineResponse200488SsidRadiusPoliciesFailover**](InlineResponse200488SsidRadiusPoliciesFailover.md) |  | [optional] 
**LoadBalancing** | Pointer to [**InlineResponse200488SsidRadiusPoliciesLoadBalancing**](InlineResponse200488SsidRadiusPoliciesLoadBalancing.md) |  | [optional] 
**Fallback** | Pointer to [**InlineResponse200488SsidRadiusPoliciesFallback**](InlineResponse200488SsidRadiusPoliciesFallback.md) |  | [optional] 
**Attribute** | Pointer to **string** | RADIUS attribute for policies | [optional] 
**Timeout** | Pointer to **int32** | The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds). | [optional] 
**NumAttempts** | Pointer to **int32** | The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5). | [optional] 

## Methods

### NewInlineResponse200488SsidRadiusPolicies

`func NewInlineResponse200488SsidRadiusPolicies() *InlineResponse200488SsidRadiusPolicies`

NewInlineResponse200488SsidRadiusPolicies instantiates a new InlineResponse200488SsidRadiusPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidRadiusPoliciesWithDefaults

`func NewInlineResponse200488SsidRadiusPoliciesWithDefaults() *InlineResponse200488SsidRadiusPolicies`

NewInlineResponse200488SsidRadiusPoliciesWithDefaults instantiates a new InlineResponse200488SsidRadiusPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailover

`func (o *InlineResponse200488SsidRadiusPolicies) GetFailover() InlineResponse200488SsidRadiusPoliciesFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetFailoverOk() (*InlineResponse200488SsidRadiusPoliciesFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200488SsidRadiusPolicies) SetFailover(v InlineResponse200488SsidRadiusPoliciesFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200488SsidRadiusPolicies) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetLoadBalancing

`func (o *InlineResponse200488SsidRadiusPolicies) GetLoadBalancing() InlineResponse200488SsidRadiusPoliciesLoadBalancing`

GetLoadBalancing returns the LoadBalancing field if non-nil, zero value otherwise.

### GetLoadBalancingOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetLoadBalancingOk() (*InlineResponse200488SsidRadiusPoliciesLoadBalancing, bool)`

GetLoadBalancingOk returns a tuple with the LoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancing

`func (o *InlineResponse200488SsidRadiusPolicies) SetLoadBalancing(v InlineResponse200488SsidRadiusPoliciesLoadBalancing)`

SetLoadBalancing sets LoadBalancing field to given value.

### HasLoadBalancing

`func (o *InlineResponse200488SsidRadiusPolicies) HasLoadBalancing() bool`

HasLoadBalancing returns a boolean if a field has been set.

### GetFallback

`func (o *InlineResponse200488SsidRadiusPolicies) GetFallback() InlineResponse200488SsidRadiusPoliciesFallback`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetFallbackOk() (*InlineResponse200488SsidRadiusPoliciesFallback, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *InlineResponse200488SsidRadiusPolicies) SetFallback(v InlineResponse200488SsidRadiusPoliciesFallback)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *InlineResponse200488SsidRadiusPolicies) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetAttribute

`func (o *InlineResponse200488SsidRadiusPolicies) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *InlineResponse200488SsidRadiusPolicies) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *InlineResponse200488SsidRadiusPolicies) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200488SsidRadiusPolicies) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200488SsidRadiusPolicies) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200488SsidRadiusPolicies) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetNumAttempts

`func (o *InlineResponse200488SsidRadiusPolicies) GetNumAttempts() int32`

GetNumAttempts returns the NumAttempts field if non-nil, zero value otherwise.

### GetNumAttemptsOk

`func (o *InlineResponse200488SsidRadiusPolicies) GetNumAttemptsOk() (*int32, bool)`

GetNumAttemptsOk returns a tuple with the NumAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAttempts

`func (o *InlineResponse200488SsidRadiusPolicies) SetNumAttempts(v int32)`

SetNumAttempts sets NumAttempts field to given value.

### HasNumAttempts

`func (o *InlineResponse200488SsidRadiusPolicies) HasNumAttempts() bool`

HasNumAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


