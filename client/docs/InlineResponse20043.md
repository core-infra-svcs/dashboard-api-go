# InlineResponse20043

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApEslId** | Pointer to **int32** | An identifier for the device used by the ESL system | [optional] 
**Serial** | Pointer to **string** | The serial number of the device | [optional] 
**Channel** | Pointer to **string** | Desired ESL channel for the device, or &#39;Auto&#39; (case insensitive) to use the recommended channel | [optional] 
**Enabled** | Pointer to **bool** | Turn ESL features on and off for this device | [optional] 
**NetworkId** | Pointer to **string** | The identifier for the device&#39;s network | [optional] 
**Hostname** | Pointer to **string** | Hostname of the ESL management service | [optional] 
**Provider** | Pointer to **string** | The service providing ESL functionality | [optional] 

## Methods

### NewInlineResponse20043

`func NewInlineResponse20043() *InlineResponse20043`

NewInlineResponse20043 instantiates a new InlineResponse20043 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20043WithDefaults

`func NewInlineResponse20043WithDefaults() *InlineResponse20043`

NewInlineResponse20043WithDefaults instantiates a new InlineResponse20043 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApEslId

`func (o *InlineResponse20043) GetApEslId() int32`

GetApEslId returns the ApEslId field if non-nil, zero value otherwise.

### GetApEslIdOk

`func (o *InlineResponse20043) GetApEslIdOk() (*int32, bool)`

GetApEslIdOk returns a tuple with the ApEslId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApEslId

`func (o *InlineResponse20043) SetApEslId(v int32)`

SetApEslId sets ApEslId field to given value.

### HasApEslId

`func (o *InlineResponse20043) HasApEslId() bool`

HasApEslId returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse20043) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse20043) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse20043) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse20043) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse20043) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse20043) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse20043) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse20043) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20043) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20043) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20043) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20043) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse20043) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20043) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20043) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20043) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetHostname

`func (o *InlineResponse20043) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InlineResponse20043) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InlineResponse20043) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InlineResponse20043) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProvider

`func (o *InlineResponse20043) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InlineResponse20043) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InlineResponse20043) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *InlineResponse20043) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


