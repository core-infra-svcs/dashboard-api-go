# InlineResponse200327BySignalType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Cellular connection type | [optional] 
**Masked** | Pointer to **[]string** | List of masked band identifiers. For LTE use bands identifiers like &#39;30&#39; and for 5G use band identifiers like &#39;n30&#39;. Maximum 256 bands per signal type. | [optional] 
**Enabled** | Pointer to **[]string** | List of enabled band identifiers. For LTE use bands identifiers like &#39;30&#39; and for 5G use band identifiers like &#39;n30&#39;. Maximum 256 bands per signal type. | [optional] 
**Supported** | Pointer to **[]string** | List of supported band identifiers. For LTE use bands identifiers like &#39;30&#39; and for 5G use band identifiers like &#39;n30&#39;. Maximum 256 bands per signal type. | [optional] 

## Methods

### NewInlineResponse200327BySignalType

`func NewInlineResponse200327BySignalType() *InlineResponse200327BySignalType`

NewInlineResponse200327BySignalType instantiates a new InlineResponse200327BySignalType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200327BySignalTypeWithDefaults

`func NewInlineResponse200327BySignalTypeWithDefaults() *InlineResponse200327BySignalType`

NewInlineResponse200327BySignalTypeWithDefaults instantiates a new InlineResponse200327BySignalType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200327BySignalType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200327BySignalType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200327BySignalType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200327BySignalType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMasked

`func (o *InlineResponse200327BySignalType) GetMasked() []string`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *InlineResponse200327BySignalType) GetMaskedOk() (*[]string, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *InlineResponse200327BySignalType) SetMasked(v []string)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *InlineResponse200327BySignalType) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200327BySignalType) GetEnabled() []string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200327BySignalType) GetEnabledOk() (*[]string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200327BySignalType) SetEnabled(v []string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200327BySignalType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSupported

`func (o *InlineResponse200327BySignalType) GetSupported() []string`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *InlineResponse200327BySignalType) GetSupportedOk() (*[]string, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *InlineResponse200327BySignalType) SetSupported(v []string)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *InlineResponse200327BySignalType) HasSupported() bool`

HasSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


