# InlineResponse200216

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]InlineResponse200216Nodes**](InlineResponse200216Nodes.md) | List of nodes in the network topology | [optional] 
**Links** | Pointer to [**[]InlineResponse200216Links**](InlineResponse200216Links.md) | List of links between nodes | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered while building the topology | [optional] 

## Methods

### NewInlineResponse200216

`func NewInlineResponse200216() *InlineResponse200216`

NewInlineResponse200216 instantiates a new InlineResponse200216 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200216WithDefaults

`func NewInlineResponse200216WithDefaults() *InlineResponse200216`

NewInlineResponse200216WithDefaults instantiates a new InlineResponse200216 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *InlineResponse200216) GetNodes() []InlineResponse200216Nodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InlineResponse200216) GetNodesOk() (*[]InlineResponse200216Nodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InlineResponse200216) SetNodes(v []InlineResponse200216Nodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InlineResponse200216) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200216) GetLinks() []InlineResponse200216Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200216) GetLinksOk() (*[]InlineResponse200216Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200216) SetLinks(v []InlineResponse200216Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200216) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200216) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200216) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200216) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200216) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


