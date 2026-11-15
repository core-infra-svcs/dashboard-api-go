# InlineResponse200499ClientsImpactedByStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | Pointer to **int32** | Number of unique client failures during the association step of connection. | [optional] 
**Authentication** | Pointer to **int32** | Number of unique client failures during the authentication step of connection. | [optional] 
**IpAssignment** | Pointer to **int32** | Number of unique client failures during the IP assignment step of connection. | [optional] 

## Methods

### NewInlineResponse200499ClientsImpactedByStep

`func NewInlineResponse200499ClientsImpactedByStep() *InlineResponse200499ClientsImpactedByStep`

NewInlineResponse200499ClientsImpactedByStep instantiates a new InlineResponse200499ClientsImpactedByStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200499ClientsImpactedByStepWithDefaults

`func NewInlineResponse200499ClientsImpactedByStepWithDefaults() *InlineResponse200499ClientsImpactedByStep`

NewInlineResponse200499ClientsImpactedByStepWithDefaults instantiates a new InlineResponse200499ClientsImpactedByStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociation

`func (o *InlineResponse200499ClientsImpactedByStep) GetAssociation() int32`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *InlineResponse200499ClientsImpactedByStep) GetAssociationOk() (*int32, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *InlineResponse200499ClientsImpactedByStep) SetAssociation(v int32)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *InlineResponse200499ClientsImpactedByStep) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetAuthentication

`func (o *InlineResponse200499ClientsImpactedByStep) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *InlineResponse200499ClientsImpactedByStep) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *InlineResponse200499ClientsImpactedByStep) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *InlineResponse200499ClientsImpactedByStep) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetIpAssignment

`func (o *InlineResponse200499ClientsImpactedByStep) GetIpAssignment() int32`

GetIpAssignment returns the IpAssignment field if non-nil, zero value otherwise.

### GetIpAssignmentOk

`func (o *InlineResponse200499ClientsImpactedByStep) GetIpAssignmentOk() (*int32, bool)`

GetIpAssignmentOk returns a tuple with the IpAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignment

`func (o *InlineResponse200499ClientsImpactedByStep) SetIpAssignment(v int32)`

SetIpAssignment sets IpAssignment field to given value.

### HasIpAssignment

`func (o *InlineResponse200499ClientsImpactedByStep) HasIpAssignment() bool`

HasIpAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


