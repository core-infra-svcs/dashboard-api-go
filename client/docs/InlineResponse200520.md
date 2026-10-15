# InlineResponse200520

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the SSID profile | [optional] 
**Name** | Pointer to **string** | Name of the SSID profile | [optional] 
**Precedence** | Pointer to [**InlineResponse200517Precedence**](InlineResponse200517Precedence.md) |  | [optional] 
**Ssid** | Pointer to [**InlineResponse200517Ssid**](InlineResponse200517Ssid.md) |  | [optional] 
**IsAttached** | **bool** | True when this profile is attached to the selected SSID (&#x60;networkIds&#x60; + &#x60;ssidNumbers&#x60;). | 

## Methods

### NewInlineResponse200520

`func NewInlineResponse200520(isAttached bool, ) *InlineResponse200520`

NewInlineResponse200520 instantiates a new InlineResponse200520 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200520WithDefaults

`func NewInlineResponse200520WithDefaults() *InlineResponse200520`

NewInlineResponse200520WithDefaults instantiates a new InlineResponse200520 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200520) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200520) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200520) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200520) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200520) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200520) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200520) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200520) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrecedence

`func (o *InlineResponse200520) GetPrecedence() InlineResponse200517Precedence`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *InlineResponse200520) GetPrecedenceOk() (*InlineResponse200517Precedence, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *InlineResponse200520) SetPrecedence(v InlineResponse200517Precedence)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *InlineResponse200520) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse200520) GetSsid() InlineResponse200517Ssid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200520) GetSsidOk() (*InlineResponse200517Ssid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200520) SetSsid(v InlineResponse200517Ssid)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200520) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetIsAttached

`func (o *InlineResponse200520) GetIsAttached() bool`

GetIsAttached returns the IsAttached field if non-nil, zero value otherwise.

### GetIsAttachedOk

`func (o *InlineResponse200520) GetIsAttachedOk() (*bool, bool)`

GetIsAttachedOk returns a tuple with the IsAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttached

`func (o *InlineResponse200520) SetIsAttached(v bool)`

SetIsAttached sets IsAttached field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


