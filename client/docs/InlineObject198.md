# InlineObject198

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**Identity** | Pointer to [**InlineResponse200210Identity**](InlineResponse200210Identity.md) |  | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**EapolKey** | Pointer to [**InlineResponse200210EapolKey**](InlineResponse200210EapolKey.md) |  | [optional] 

## Methods

### NewInlineObject198

`func NewInlineObject198() *InlineObject198`

NewInlineObject198 instantiates a new InlineObject198 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject198WithDefaults

`func NewInlineObject198WithDefaults() *InlineObject198`

NewInlineObject198WithDefaults instantiates a new InlineObject198 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *InlineObject198) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject198) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject198) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject198) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIdentity

`func (o *InlineObject198) GetIdentity() InlineResponse200210Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InlineObject198) GetIdentityOk() (*InlineResponse200210Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InlineObject198) SetIdentity(v InlineResponse200210Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InlineObject198) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaxRetries

`func (o *InlineObject198) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *InlineObject198) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *InlineObject198) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *InlineObject198) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetEapolKey

`func (o *InlineObject198) GetEapolKey() InlineResponse200210EapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *InlineObject198) GetEapolKeyOk() (*InlineResponse200210EapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *InlineObject198) SetEapolKey(v InlineResponse200210EapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *InlineObject198) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


