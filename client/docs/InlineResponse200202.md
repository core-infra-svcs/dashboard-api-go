# InlineResponse200202

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]InlineResponse200202Nodes**](InlineResponse200202Nodes.md) | List of nodes in the network topology | [optional] 
**Links** | Pointer to [**[]InlineResponse200202Links**](InlineResponse200202Links.md) | List of links between nodes | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered while building the topology | [optional] 

## Methods

### NewInlineResponse200202

`func NewInlineResponse200202() *InlineResponse200202`

NewInlineResponse200202 instantiates a new InlineResponse200202 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200202WithDefaults

`func NewInlineResponse200202WithDefaults() *InlineResponse200202`

NewInlineResponse200202WithDefaults instantiates a new InlineResponse200202 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *InlineResponse200202) GetNodes() []InlineResponse200202Nodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InlineResponse200202) GetNodesOk() (*[]InlineResponse200202Nodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InlineResponse200202) SetNodes(v []InlineResponse200202Nodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InlineResponse200202) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200202) GetLinks() []InlineResponse200202Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200202) GetLinksOk() (*[]InlineResponse200202Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200202) SetLinks(v []InlineResponse200202Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200202) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200202) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200202) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200202) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200202) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


