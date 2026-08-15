# InlineResponse200488SsidRadiusServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of the RADIUS server | [optional] 
**Address** | Pointer to **string** | The IP address or FQDN of the RADIUS server | [optional] 
**Name** | Pointer to **string** | Name of the RADIUS server | [optional] 
**Modes** | Pointer to [**[]InlineResponse200488SsidRadiusModes**](InlineResponse200488SsidRadiusModes.md) | Available server modes | [optional] 

## Methods

### NewInlineResponse200488SsidRadiusServers

`func NewInlineResponse200488SsidRadiusServers() *InlineResponse200488SsidRadiusServers`

NewInlineResponse200488SsidRadiusServers instantiates a new InlineResponse200488SsidRadiusServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidRadiusServersWithDefaults

`func NewInlineResponse200488SsidRadiusServersWithDefaults() *InlineResponse200488SsidRadiusServers`

NewInlineResponse200488SsidRadiusServersWithDefaults instantiates a new InlineResponse200488SsidRadiusServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200488SsidRadiusServers) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200488SsidRadiusServers) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200488SsidRadiusServers) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200488SsidRadiusServers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200488SsidRadiusServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200488SsidRadiusServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200488SsidRadiusServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200488SsidRadiusServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200488SsidRadiusServers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200488SsidRadiusServers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200488SsidRadiusServers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200488SsidRadiusServers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModes

`func (o *InlineResponse200488SsidRadiusServers) GetModes() []InlineResponse200488SsidRadiusModes`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *InlineResponse200488SsidRadiusServers) GetModesOk() (*[]InlineResponse200488SsidRadiusModes, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *InlineResponse200488SsidRadiusServers) SetModes(v []InlineResponse200488SsidRadiusModes)`

SetModes sets Modes field to given value.

### HasModes

`func (o *InlineResponse200488SsidRadiusServers) HasModes() bool`

HasModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


