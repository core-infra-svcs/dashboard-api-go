# InlineResponse200379Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Changes** | Pointer to [**[]InlineResponse200379Changes**](InlineResponse200379Changes.md) | The statuses of layer 2 interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200379Interfaces

`func NewInlineResponse200379Interfaces() *InlineResponse200379Interfaces`

NewInlineResponse200379Interfaces instantiates a new InlineResponse200379Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200379InterfacesWithDefaults

`func NewInlineResponse200379InterfacesWithDefaults() *InlineResponse200379Interfaces`

NewInlineResponse200379InterfacesWithDefaults instantiates a new InlineResponse200379Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200379Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200379Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200379Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200379Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200379Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200379Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200379Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200379Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200379Interfaces) GetChanges() []InlineResponse200379Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200379Interfaces) GetChangesOk() (*[]InlineResponse200379Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200379Interfaces) SetChanges(v []InlineResponse200379Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200379Interfaces) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


