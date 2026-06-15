# InlineResponse2008

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | A port&#39;s identifier as labeled on the device | [optional] 
**Interface** | Pointer to [**InlineResponse2008Interface**](InlineResponse2008Interface.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the port is enabled | [optional] 
**Name** | Pointer to **string** | The name of the interface the port is mapped to | [optional] 
**Personality** | Pointer to [**InlineResponse2008Personality**](InlineResponse2008Personality.md) |  | [optional] 
**Uplink** | Pointer to [**InlineResponse2008Uplink**](InlineResponse2008Uplink.md) |  | [optional] 
**Downlink** | Pointer to [**InlineResponse2008Downlink**](InlineResponse2008Downlink.md) |  | [optional] 

## Methods

### NewInlineResponse2008

`func NewInlineResponse2008() *InlineResponse2008`

NewInlineResponse2008 instantiates a new InlineResponse2008 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008WithDefaults

`func NewInlineResponse2008WithDefaults() *InlineResponse2008`

NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *InlineResponse2008) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InlineResponse2008) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InlineResponse2008) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InlineResponse2008) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetInterface

`func (o *InlineResponse2008) GetInterface() InlineResponse2008Interface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InlineResponse2008) GetInterfaceOk() (*InlineResponse2008Interface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InlineResponse2008) SetInterface(v InlineResponse2008Interface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *InlineResponse2008) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse2008) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse2008) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse2008) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse2008) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2008) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2008) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2008) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2008) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonality

`func (o *InlineResponse2008) GetPersonality() InlineResponse2008Personality`

GetPersonality returns the Personality field if non-nil, zero value otherwise.

### GetPersonalityOk

`func (o *InlineResponse2008) GetPersonalityOk() (*InlineResponse2008Personality, bool)`

GetPersonalityOk returns a tuple with the Personality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonality

`func (o *InlineResponse2008) SetPersonality(v InlineResponse2008Personality)`

SetPersonality sets Personality field to given value.

### HasPersonality

`func (o *InlineResponse2008) HasPersonality() bool`

HasPersonality returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse2008) GetUplink() InlineResponse2008Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse2008) GetUplinkOk() (*InlineResponse2008Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse2008) SetUplink(v InlineResponse2008Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse2008) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetDownlink

`func (o *InlineResponse2008) GetDownlink() InlineResponse2008Downlink`

GetDownlink returns the Downlink field if non-nil, zero value otherwise.

### GetDownlinkOk

`func (o *InlineResponse2008) GetDownlinkOk() (*InlineResponse2008Downlink, bool)`

GetDownlinkOk returns a tuple with the Downlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlink

`func (o *InlineResponse2008) SetDownlink(v InlineResponse2008Downlink)`

SetDownlink sets Downlink field to given value.

### HasDownlink

`func (o *InlineResponse2008) HasDownlink() bool`

HasDownlink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


