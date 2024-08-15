# InlineResponse200189

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**Identity** | Pointer to [**InlineResponse200189Identity**](InlineResponse200189Identity.md) |  | [optional] 
**EapolKey** | Pointer to [**InlineResponse200189EapolKey**](InlineResponse200189EapolKey.md) |  | [optional] 

## Methods

### NewInlineResponse200189

`func NewInlineResponse200189() *InlineResponse200189`

NewInlineResponse200189 instantiates a new InlineResponse200189 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200189WithDefaults

`func NewInlineResponse200189WithDefaults() *InlineResponse200189`

NewInlineResponse200189WithDefaults instantiates a new InlineResponse200189 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *InlineResponse200189) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse200189) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse200189) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse200189) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *InlineResponse200189) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *InlineResponse200189) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *InlineResponse200189) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *InlineResponse200189) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetIdentity

`func (o *InlineResponse200189) GetIdentity() InlineResponse200189Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InlineResponse200189) GetIdentityOk() (*InlineResponse200189Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InlineResponse200189) SetIdentity(v InlineResponse200189Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InlineResponse200189) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetEapolKey

`func (o *InlineResponse200189) GetEapolKey() InlineResponse200189EapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *InlineResponse200189) GetEapolKeyOk() (*InlineResponse200189EapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *InlineResponse200189) SetEapolKey(v InlineResponse200189EapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *InlineResponse200189) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


