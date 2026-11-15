# InlineResponse200529

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the SSID profile | [optional] 
**Name** | Pointer to **string** | Name of the SSID profile | [optional] 
**Precedence** | Pointer to [**InlineResponse200529Precedence**](InlineResponse200529Precedence.md) |  | [optional] 
**Ssid** | Pointer to [**InlineResponse200529Ssid**](InlineResponse200529Ssid.md) |  | [optional] 

## Methods

### NewInlineResponse200529

`func NewInlineResponse200529() *InlineResponse200529`

NewInlineResponse200529 instantiates a new InlineResponse200529 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529WithDefaults

`func NewInlineResponse200529WithDefaults() *InlineResponse200529`

NewInlineResponse200529WithDefaults instantiates a new InlineResponse200529 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200529) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200529) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200529) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200529) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200529) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200529) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200529) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200529) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrecedence

`func (o *InlineResponse200529) GetPrecedence() InlineResponse200529Precedence`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *InlineResponse200529) GetPrecedenceOk() (*InlineResponse200529Precedence, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *InlineResponse200529) SetPrecedence(v InlineResponse200529Precedence)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *InlineResponse200529) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetSsid

`func (o *InlineResponse200529) GetSsid() InlineResponse200529Ssid`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *InlineResponse200529) GetSsidOk() (*InlineResponse200529Ssid, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *InlineResponse200529) SetSsid(v InlineResponse200529Ssid)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *InlineResponse200529) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


