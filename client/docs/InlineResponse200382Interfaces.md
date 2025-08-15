# InlineResponse200382Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Changes** | Pointer to [**[]InlineResponse200379Changes**](InlineResponse200379Changes.md) | The statuses of layer 3 interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200382Interfaces

`func NewInlineResponse200382Interfaces() *InlineResponse200382Interfaces`

NewInlineResponse200382Interfaces instantiates a new InlineResponse200382Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200382InterfacesWithDefaults

`func NewInlineResponse200382InterfacesWithDefaults() *InlineResponse200382Interfaces`

NewInlineResponse200382InterfacesWithDefaults instantiates a new InlineResponse200382Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200382Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200382Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200382Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200382Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200382Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200382Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200382Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200382Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200382Interfaces) GetChanges() []InlineResponse200379Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200382Interfaces) GetChangesOk() (*[]InlineResponse200379Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200382Interfaces) SetChanges(v []InlineResponse200379Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200382Interfaces) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


