# InlineResponse200493SsidRadiusPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failover** | Pointer to [**InlineResponse200493SsidRadiusPoliciesFailover**](InlineResponse200493SsidRadiusPoliciesFailover.md) |  | [optional] 
**LoadBalancing** | Pointer to [**InlineResponse200493SsidRadiusPoliciesLoadBalancing**](InlineResponse200493SsidRadiusPoliciesLoadBalancing.md) |  | [optional] 
**Fallback** | Pointer to [**InlineResponse200493SsidRadiusPoliciesFallback**](InlineResponse200493SsidRadiusPoliciesFallback.md) |  | [optional] 
**Attribute** | Pointer to **string** | RADIUS attribute for policies | [optional] 
**Timeout** | Pointer to **int32** | The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds). | [optional] 
**NumAttempts** | Pointer to **int32** | The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5). | [optional] 

## Methods

### NewInlineResponse200493SsidRadiusPolicies

`func NewInlineResponse200493SsidRadiusPolicies() *InlineResponse200493SsidRadiusPolicies`

NewInlineResponse200493SsidRadiusPolicies instantiates a new InlineResponse200493SsidRadiusPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200493SsidRadiusPoliciesWithDefaults

`func NewInlineResponse200493SsidRadiusPoliciesWithDefaults() *InlineResponse200493SsidRadiusPolicies`

NewInlineResponse200493SsidRadiusPoliciesWithDefaults instantiates a new InlineResponse200493SsidRadiusPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailover

`func (o *InlineResponse200493SsidRadiusPolicies) GetFailover() InlineResponse200493SsidRadiusPoliciesFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetFailoverOk() (*InlineResponse200493SsidRadiusPoliciesFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200493SsidRadiusPolicies) SetFailover(v InlineResponse200493SsidRadiusPoliciesFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200493SsidRadiusPolicies) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetLoadBalancing

`func (o *InlineResponse200493SsidRadiusPolicies) GetLoadBalancing() InlineResponse200493SsidRadiusPoliciesLoadBalancing`

GetLoadBalancing returns the LoadBalancing field if non-nil, zero value otherwise.

### GetLoadBalancingOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetLoadBalancingOk() (*InlineResponse200493SsidRadiusPoliciesLoadBalancing, bool)`

GetLoadBalancingOk returns a tuple with the LoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancing

`func (o *InlineResponse200493SsidRadiusPolicies) SetLoadBalancing(v InlineResponse200493SsidRadiusPoliciesLoadBalancing)`

SetLoadBalancing sets LoadBalancing field to given value.

### HasLoadBalancing

`func (o *InlineResponse200493SsidRadiusPolicies) HasLoadBalancing() bool`

HasLoadBalancing returns a boolean if a field has been set.

### GetFallback

`func (o *InlineResponse200493SsidRadiusPolicies) GetFallback() InlineResponse200493SsidRadiusPoliciesFallback`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetFallbackOk() (*InlineResponse200493SsidRadiusPoliciesFallback, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *InlineResponse200493SsidRadiusPolicies) SetFallback(v InlineResponse200493SsidRadiusPoliciesFallback)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *InlineResponse200493SsidRadiusPolicies) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetAttribute

`func (o *InlineResponse200493SsidRadiusPolicies) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *InlineResponse200493SsidRadiusPolicies) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *InlineResponse200493SsidRadiusPolicies) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetTimeout

`func (o *InlineResponse200493SsidRadiusPolicies) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200493SsidRadiusPolicies) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200493SsidRadiusPolicies) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetNumAttempts

`func (o *InlineResponse200493SsidRadiusPolicies) GetNumAttempts() int32`

GetNumAttempts returns the NumAttempts field if non-nil, zero value otherwise.

### GetNumAttemptsOk

`func (o *InlineResponse200493SsidRadiusPolicies) GetNumAttemptsOk() (*int32, bool)`

GetNumAttemptsOk returns a tuple with the NumAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAttempts

`func (o *InlineResponse200493SsidRadiusPolicies) SetNumAttempts(v int32)`

SetNumAttempts sets NumAttempts field to given value.

### HasNumAttempts

`func (o *InlineResponse200493SsidRadiusPolicies) HasNumAttempts() bool`

HasNumAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


