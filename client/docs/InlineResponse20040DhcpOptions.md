# InlineResponse20040DhcpOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The code for DHCP option which should be from 2 to 254 | [optional] 
**Type** | Pointer to **string** | The type of the DHCP option which should be one of (&#39;text&#39;, &#39;ip&#39;, &#39;integer&#39; or &#39;hex&#39;) | [optional] 
**Value** | Pointer to **string** | The value of the DHCP option | [optional] 

## Methods

### NewInlineResponse20040DhcpOptions

`func NewInlineResponse20040DhcpOptions() *InlineResponse20040DhcpOptions`

NewInlineResponse20040DhcpOptions instantiates a new InlineResponse20040DhcpOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20040DhcpOptionsWithDefaults

`func NewInlineResponse20040DhcpOptionsWithDefaults() *InlineResponse20040DhcpOptions`

NewInlineResponse20040DhcpOptionsWithDefaults instantiates a new InlineResponse20040DhcpOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineResponse20040DhcpOptions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20040DhcpOptions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20040DhcpOptions) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20040DhcpOptions) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20040DhcpOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20040DhcpOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20040DhcpOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20040DhcpOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InlineResponse20040DhcpOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse20040DhcpOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse20040DhcpOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineResponse20040DhcpOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


