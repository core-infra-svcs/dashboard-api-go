# InlineResponse200102Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The step of the auto locate process when the error occurred. Possible values: &#39;gnss&#39;, &#39;ranging&#39;, &#39;positioning&#39; | [optional] 
**Type** | Pointer to **string** | The type of error that occurred. Possible values: &#39;failure&#39;, &#39;no neighbors&#39;, &#39;missing anchors&#39;, &#39;wrong anchors&#39;, &#39;missing ranging data&#39;, &#39;calculation failure&#39;, &#39;scheduling failure&#39; | [optional] 

## Methods

### NewInlineResponse200102Errors

`func NewInlineResponse200102Errors() *InlineResponse200102Errors`

NewInlineResponse200102Errors instantiates a new InlineResponse200102Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102ErrorsWithDefaults

`func NewInlineResponse200102ErrorsWithDefaults() *InlineResponse200102Errors`

NewInlineResponse200102ErrorsWithDefaults instantiates a new InlineResponse200102Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *InlineResponse200102Errors) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse200102Errors) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse200102Errors) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InlineResponse200102Errors) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200102Errors) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200102Errors) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200102Errors) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200102Errors) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


