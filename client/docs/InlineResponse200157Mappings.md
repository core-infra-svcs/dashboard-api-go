# InlineResponse200157Mappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dscp** | Pointer to **int32** | The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive. | [optional] 
**Cos** | Pointer to **int32** | The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive. | [optional] 
**Title** | Pointer to **string** | Label for the mapping (optional). | [optional] 

## Methods

### NewInlineResponse200157Mappings

`func NewInlineResponse200157Mappings() *InlineResponse200157Mappings`

NewInlineResponse200157Mappings instantiates a new InlineResponse200157Mappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200157MappingsWithDefaults

`func NewInlineResponse200157MappingsWithDefaults() *InlineResponse200157Mappings`

NewInlineResponse200157MappingsWithDefaults instantiates a new InlineResponse200157Mappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscp

`func (o *InlineResponse200157Mappings) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *InlineResponse200157Mappings) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *InlineResponse200157Mappings) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *InlineResponse200157Mappings) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetCos

`func (o *InlineResponse200157Mappings) GetCos() int32`

GetCos returns the Cos field if non-nil, zero value otherwise.

### GetCosOk

`func (o *InlineResponse200157Mappings) GetCosOk() (*int32, bool)`

GetCosOk returns a tuple with the Cos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCos

`func (o *InlineResponse200157Mappings) SetCos(v int32)`

SetCos sets Cos field to given value.

### HasCos

`func (o *InlineResponse200157Mappings) HasCos() bool`

HasCos returns a boolean if a field has been set.

### GetTitle

`func (o *InlineResponse200157Mappings) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InlineResponse200157Mappings) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InlineResponse200157Mappings) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InlineResponse200157Mappings) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


