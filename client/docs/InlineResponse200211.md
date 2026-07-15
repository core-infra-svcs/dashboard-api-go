# InlineResponse200211

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]InlineResponse200211Nodes**](InlineResponse200211Nodes.md) | List of nodes in the network topology | [optional] 
**Links** | Pointer to [**[]InlineResponse200211Links**](InlineResponse200211Links.md) | List of links between nodes | [optional] 
**Errors** | Pointer to **[]string** | List of errors encountered while building the topology | [optional] 

## Methods

### NewInlineResponse200211

`func NewInlineResponse200211() *InlineResponse200211`

NewInlineResponse200211 instantiates a new InlineResponse200211 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200211WithDefaults

`func NewInlineResponse200211WithDefaults() *InlineResponse200211`

NewInlineResponse200211WithDefaults instantiates a new InlineResponse200211 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *InlineResponse200211) GetNodes() []InlineResponse200211Nodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *InlineResponse200211) GetNodesOk() (*[]InlineResponse200211Nodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *InlineResponse200211) SetNodes(v []InlineResponse200211Nodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *InlineResponse200211) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLinks

`func (o *InlineResponse200211) GetLinks() []InlineResponse200211Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineResponse200211) GetLinksOk() (*[]InlineResponse200211Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineResponse200211) SetLinks(v []InlineResponse200211Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineResponse200211) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200211) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200211) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200211) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200211) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


