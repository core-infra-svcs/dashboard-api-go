# InlineResponse200251NaiRealms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | The format for the realm (&#39;1&#39; or &#39;0&#39;) | [optional] 
**Name** | Pointer to **string** | The name of the realm | [optional] 
**Methods** | Pointer to [**[]InlineResponse200251Methods**](InlineResponse200251Methods.md) | An array of EAP methods for the realm. | [optional] 

## Methods

### NewInlineResponse200251NaiRealms

`func NewInlineResponse200251NaiRealms() *InlineResponse200251NaiRealms`

NewInlineResponse200251NaiRealms instantiates a new InlineResponse200251NaiRealms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200251NaiRealmsWithDefaults

`func NewInlineResponse200251NaiRealmsWithDefaults() *InlineResponse200251NaiRealms`

NewInlineResponse200251NaiRealmsWithDefaults instantiates a new InlineResponse200251NaiRealms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *InlineResponse200251NaiRealms) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InlineResponse200251NaiRealms) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InlineResponse200251NaiRealms) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *InlineResponse200251NaiRealms) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200251NaiRealms) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200251NaiRealms) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200251NaiRealms) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200251NaiRealms) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMethods

`func (o *InlineResponse200251NaiRealms) GetMethods() []InlineResponse200251Methods`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *InlineResponse200251NaiRealms) GetMethodsOk() (*[]InlineResponse200251Methods, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *InlineResponse200251NaiRealms) SetMethods(v []InlineResponse200251Methods)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *InlineResponse200251NaiRealms) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


