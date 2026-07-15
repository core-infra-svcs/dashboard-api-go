# InlineObject218

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**Identity** | Pointer to [**InlineResponse200246Identity**](InlineResponse200246Identity.md) |  | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**EapolKey** | Pointer to [**InlineResponse200246EapolKey**](InlineResponse200246EapolKey.md) |  | [optional] 

## Methods

### NewInlineObject218

`func NewInlineObject218() *InlineObject218`

NewInlineObject218 instantiates a new InlineObject218 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject218WithDefaults

`func NewInlineObject218WithDefaults() *InlineObject218`

NewInlineObject218WithDefaults instantiates a new InlineObject218 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *InlineObject218) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject218) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject218) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject218) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIdentity

`func (o *InlineObject218) GetIdentity() InlineResponse200246Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InlineObject218) GetIdentityOk() (*InlineResponse200246Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InlineObject218) SetIdentity(v InlineResponse200246Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InlineObject218) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaxRetries

`func (o *InlineObject218) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *InlineObject218) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *InlineObject218) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *InlineObject218) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetEapolKey

`func (o *InlineObject218) GetEapolKey() InlineResponse200246EapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *InlineObject218) GetEapolKeyOk() (*InlineResponse200246EapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *InlineObject218) SetEapolKey(v InlineResponse200246EapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *InlineObject218) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


