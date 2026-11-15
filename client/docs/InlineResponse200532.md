# InlineResponse200532

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the SSID profile | [optional] 
**Name** | Pointer to **string** | Name of the SSID profile | [optional] 
**Precedence** | Pointer to [**InlineResponse200529Precedence**](InlineResponse200529Precedence.md) |  | [optional] 
**Ssid** | Pointer to [**InlineResponse200529Ssid**](InlineResponse200529Ssid.md) |  | [optional] 
**IsAttached** | **bool** | True when this profile is attached to the selected SSID (&#x60;networkIds&#x60; + &#x60;ssidNumbers&#x60;). | 

## Methods

### NewInlineResponse200532

`func NewInlineResponse200532(isAttached bool, ) *InlineResponse200532`

NewInlineResponse200532 instantiates a new InlineResponse200532 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200532WithDefaults

`func NewInlineResponse200532WithDefaults() *InlineResponse200532`

NewInlineResponse200532WithDefaults instantiates a new InlineResponse200532 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200532) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200532) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200532) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200532) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200532) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200532) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200532) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200532) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrecedence

`func (o *InlineResponse200532) GetPrecedence() InlineResponse200529Precedence`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *InlineResponse200532) GetPrecedenceOk() (*InlineResponse200529Precedence, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *InlineResponse200532) SetPrecedence(v InlineResponse200529Precedence)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *InlineResponse200532) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse200532) GetSsid() InlineResponse200529Ssid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200532) GetSsidOk() (*InlineResponse200529Ssid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200532) SetSsid(v InlineResponse200529Ssid)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200532) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetIsAttached

`func (o *InlineResponse200532) GetIsAttached() bool`

GetIsAttached returns the IsAttached field if non-nil, zero value otherwise.

### GetIsAttachedOk

`func (o *InlineResponse200532) GetIsAttachedOk() (*bool, bool)`

GetIsAttachedOk returns a tuple with the IsAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttached

`func (o *InlineResponse200532) SetIsAttached(v bool)`

SetIsAttached sets IsAttached field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


