# InlineResponse20099

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**Identity** | Pointer to [**InlineResponse20099Identity**](InlineResponse20099Identity.md) |  | [optional] 
**EapolKey** | Pointer to [**InlineResponse20099EapolKey**](InlineResponse20099EapolKey.md) |  | [optional] 

## Methods

### NewInlineResponse20099

`func NewInlineResponse20099() *InlineResponse20099`

NewInlineResponse20099 instantiates a new InlineResponse20099 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20099WithDefaults

`func NewInlineResponse20099WithDefaults() *InlineResponse20099`

NewInlineResponse20099WithDefaults instantiates a new InlineResponse20099 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *InlineResponse20099) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineResponse20099) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineResponse20099) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineResponse20099) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *InlineResponse20099) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *InlineResponse20099) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *InlineResponse20099) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *InlineResponse20099) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetIdentity

`func (o *InlineResponse20099) GetIdentity() InlineResponse20099Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InlineResponse20099) GetIdentityOk() (*InlineResponse20099Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InlineResponse20099) SetIdentity(v InlineResponse20099Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InlineResponse20099) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetEapolKey

`func (o *InlineResponse20099) GetEapolKey() InlineResponse20099EapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *InlineResponse20099) GetEapolKeyOk() (*InlineResponse20099EapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *InlineResponse20099) SetEapolKey(v InlineResponse20099EapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *InlineResponse20099) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


