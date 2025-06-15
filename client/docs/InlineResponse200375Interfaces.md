# InlineResponse200375Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Changes** | Pointer to [**[]InlineResponse200372Changes**](InlineResponse200372Changes.md) | The statuses of layer 3 interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200375Interfaces

`func NewInlineResponse200375Interfaces() *InlineResponse200375Interfaces`

NewInlineResponse200375Interfaces instantiates a new InlineResponse200375Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200375InterfacesWithDefaults

`func NewInlineResponse200375InterfacesWithDefaults() *InlineResponse200375Interfaces`

NewInlineResponse200375InterfacesWithDefaults instantiates a new InlineResponse200375Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200375Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200375Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200375Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200375Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200375Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200375Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200375Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200375Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200375Interfaces) GetChanges() []InlineResponse200372Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200375Interfaces) GetChangesOk() (*[]InlineResponse200372Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200375Interfaces) SetChanges(v []InlineResponse200372Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200375Interfaces) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


