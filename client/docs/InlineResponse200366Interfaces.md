# InlineResponse200366Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Changes** | Pointer to [**[]InlineResponse200363Changes**](InlineResponse200363Changes.md) | The statuses of layer 3 interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200366Interfaces

`func NewInlineResponse200366Interfaces() *InlineResponse200366Interfaces`

NewInlineResponse200366Interfaces instantiates a new InlineResponse200366Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200366InterfacesWithDefaults

`func NewInlineResponse200366InterfacesWithDefaults() *InlineResponse200366Interfaces`

NewInlineResponse200366InterfacesWithDefaults instantiates a new InlineResponse200366Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200366Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200366Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200366Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200366Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200366Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200366Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200366Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200366Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200366Interfaces) GetChanges() []InlineResponse200363Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200366Interfaces) GetChangesOk() (*[]InlineResponse200363Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200366Interfaces) SetChanges(v []InlineResponse200363Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200366Interfaces) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


