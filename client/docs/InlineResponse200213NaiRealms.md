# InlineResponse200213NaiRealms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | The format for the realm (&#39;1&#39; or &#39;0&#39;) | [optional] 
**Name** | Pointer to **string** | The name of the realm | [optional] 
**Methods** | Pointer to [**[]InlineResponse200213Methods**](InlineResponse200213Methods.md) | An array of EAP methods for the realm. | [optional] 

## Methods

### NewInlineResponse200213NaiRealms

`func NewInlineResponse200213NaiRealms() *InlineResponse200213NaiRealms`

NewInlineResponse200213NaiRealms instantiates a new InlineResponse200213NaiRealms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200213NaiRealmsWithDefaults

`func NewInlineResponse200213NaiRealmsWithDefaults() *InlineResponse200213NaiRealms`

NewInlineResponse200213NaiRealmsWithDefaults instantiates a new InlineResponse200213NaiRealms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *InlineResponse200213NaiRealms) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InlineResponse200213NaiRealms) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InlineResponse200213NaiRealms) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *InlineResponse200213NaiRealms) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200213NaiRealms) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200213NaiRealms) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200213NaiRealms) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200213NaiRealms) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMethods

`func (o *InlineResponse200213NaiRealms) GetMethods() []InlineResponse200213Methods`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *InlineResponse200213NaiRealms) GetMethodsOk() (*[]InlineResponse200213Methods, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *InlineResponse200213NaiRealms) SetMethods(v []InlineResponse200213Methods)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *InlineResponse200213NaiRealms) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


