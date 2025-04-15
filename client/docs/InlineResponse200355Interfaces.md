# InlineResponse200355Interfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the wireless LAN controller interface | [optional] 
**Mac** | Pointer to **string** | The MAC address of the wireless LAN controller interface | [optional] 
**Changes** | Pointer to [**[]InlineResponse200355Changes**](InlineResponse200355Changes.md) | The statuses of layer 2 interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200355Interfaces

`func NewInlineResponse200355Interfaces() *InlineResponse200355Interfaces`

NewInlineResponse200355Interfaces instantiates a new InlineResponse200355Interfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200355InterfacesWithDefaults

`func NewInlineResponse200355InterfacesWithDefaults() *InlineResponse200355Interfaces`

NewInlineResponse200355InterfacesWithDefaults instantiates a new InlineResponse200355Interfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineResponse200355Interfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200355Interfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200355Interfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200355Interfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200355Interfaces) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200355Interfaces) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200355Interfaces) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200355Interfaces) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetChanges

`func (o *InlineResponse200355Interfaces) GetChanges() []InlineResponse200355Changes`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *InlineResponse200355Interfaces) GetChangesOk() (*[]InlineResponse200355Changes, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *InlineResponse200355Interfaces) SetChanges(v []InlineResponse200355Changes)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *InlineResponse200355Interfaces) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


