# InlineResponse200179

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]InlineResponse200179Nodes**](InlineResponse200179Nodes.md) | List of nodes in the network topology | [optional] 
**Links** | Pointer to [**[]InlineResponse200179Links**](InlineResponse200179Links.md) | List of links between nodes | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered while building the topology | [optional] 

## Methods

### NewInlineResponse200179

`func NewInlineResponse200179() *InlineResponse200179`

NewInlineResponse200179 instantiates a new InlineResponse200179 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200179WithDefaults

`func NewInlineResponse200179WithDefaults() *InlineResponse200179`

NewInlineResponse200179WithDefaults instantiates a new InlineResponse200179 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *InlineResponse200179) GetNodes() []InlineResponse200179Nodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InlineResponse200179) GetNodesOk() (*[]InlineResponse200179Nodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InlineResponse200179) SetNodes(v []InlineResponse200179Nodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InlineResponse200179) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200179) GetLinks() []InlineResponse200179Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200179) GetLinksOk() (*[]InlineResponse200179Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200179) SetLinks(v []InlineResponse200179Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200179) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200179) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200179) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200179) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200179) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


