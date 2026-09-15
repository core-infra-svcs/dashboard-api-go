# InlineResponse200235

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Desired ESL hostname of the network | [optional] 
**Enabled** | Pointer to **bool** | Turn ESL features on and off for this network | [optional] 
**Mode** | Pointer to **string** | Electronic shelf label mode of the network. Valid options are &#39;Bluetooth&#39;, &#39;high frequency&#39; | [optional] 
**Sepioo** | Pointer to [**InlineResponse200235Sepioo**](InlineResponse200235Sepioo.md) |  | [optional] 

## Methods

### NewInlineResponse200235

`func NewInlineResponse200235() *InlineResponse200235`

NewInlineResponse200235 instantiates a new InlineResponse200235 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200235WithDefaults

`func NewInlineResponse200235WithDefaults() *InlineResponse200235`

NewInlineResponse200235WithDefaults instantiates a new InlineResponse200235 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *InlineResponse200235) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InlineResponse200235) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InlineResponse200235) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InlineResponse200235) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse200235) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200235) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200235) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200235) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200235) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200235) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200235) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200235) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSepioo

`func (o *InlineResponse200235) GetSepioo() InlineResponse200235Sepioo`

GetSepioo returns the Sepioo field if non-nil, zero value otherwise.

### GetSepiooOk

`func (o *InlineResponse200235) GetSepiooOk() (*InlineResponse200235Sepioo, bool)`

GetSepiooOk returns a tuple with the Sepioo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepioo

`func (o *InlineResponse200235) SetSepioo(v InlineResponse200235Sepioo)`

SetSepioo sets Sepioo field to given value.

### HasSepioo

`func (o *InlineResponse200235) HasSepioo() bool`

HasSepioo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


