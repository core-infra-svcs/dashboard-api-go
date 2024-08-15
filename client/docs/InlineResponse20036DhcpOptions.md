# InlineResponse20036DhcpOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The code for DHCP option which should be from 2 to 254 | [optional] 
**Type** | Pointer to **string** | The type of the DHCP option which should be one of (&#39;text&#39;, &#39;ip&#39;, &#39;integer&#39; or &#39;hex&#39;) | [optional] 
**Value** | Pointer to **string** | The value of the DHCP option | [optional] 

## Methods

### NewInlineResponse20036DhcpOptions

`func NewInlineResponse20036DhcpOptions() *InlineResponse20036DhcpOptions`

NewInlineResponse20036DhcpOptions instantiates a new InlineResponse20036DhcpOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036DhcpOptionsWithDefaults

`func NewInlineResponse20036DhcpOptionsWithDefaults() *InlineResponse20036DhcpOptions`

NewInlineResponse20036DhcpOptionsWithDefaults instantiates a new InlineResponse20036DhcpOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineResponse20036DhcpOptions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20036DhcpOptions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20036DhcpOptions) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20036DhcpOptions) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20036DhcpOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20036DhcpOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20036DhcpOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20036DhcpOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *InlineResponse20036DhcpOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse20036DhcpOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse20036DhcpOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineResponse20036DhcpOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


