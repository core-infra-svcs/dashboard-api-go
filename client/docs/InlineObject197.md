# InlineObject197

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | General EAP timeout in seconds. | [optional] 
**Identity** | Pointer to [**InlineResponse200201Identity**](InlineResponse200201Identity.md) |  | [optional] 
**MaxRetries** | Pointer to **int32** | Maximum number of general EAP retries. | [optional] 
**EapolKey** | Pointer to [**InlineResponse200201EapolKey**](InlineResponse200201EapolKey.md) |  | [optional] 

## Methods

### NewInlineObject197

`func NewInlineObject197() *InlineObject197`

NewInlineObject197 instantiates a new InlineObject197 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject197WithDefaults

`func NewInlineObject197WithDefaults() *InlineObject197`

NewInlineObject197WithDefaults instantiates a new InlineObject197 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *InlineObject197) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *InlineObject197) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *InlineObject197) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *InlineObject197) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetIdentity

`func (o *InlineObject197) GetIdentity() InlineResponse200201Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *InlineObject197) GetIdentityOk() (*InlineResponse200201Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *InlineObject197) SetIdentity(v InlineResponse200201Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *InlineObject197) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMaxRetries

`func (o *InlineObject197) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *InlineObject197) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *InlineObject197) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *InlineObject197) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetEapolKey

`func (o *InlineObject197) GetEapolKey() InlineResponse200201EapolKey`

GetEapolKey returns the EapolKey field if non-nil, zero value otherwise.

### GetEapolKeyOk

`func (o *InlineObject197) GetEapolKeyOk() (*InlineResponse200201EapolKey, bool)`

GetEapolKeyOk returns a tuple with the EapolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapolKey

`func (o *InlineObject197) SetEapolKey(v InlineResponse200201EapolKey)`

SetEapolKey sets EapolKey field to given value.

### HasEapolKey

`func (o *InlineObject197) HasEapolKey() bool`

HasEapolKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


