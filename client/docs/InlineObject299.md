# InlineObject299

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**OrganizationsOrganizationIdNetworksMovesNetwork**](OrganizationsOrganizationIdNetworksMovesNetwork.md) |  | 
**Organizations** | [**OrganizationsOrganizationIdNetworksMovesOrganizations**](OrganizationsOrganizationIdNetworksMovesOrganizations.md) |  | 
**Simulate** | Pointer to **bool** | If true, simulates the network move and validates the operation without committing changes. The network will remain in the source organization. | [optional] 

## Methods

### NewInlineObject299

`func NewInlineObject299(network OrganizationsOrganizationIdNetworksMovesNetwork, organizations OrganizationsOrganizationIdNetworksMovesOrganizations, ) *InlineObject299`

NewInlineObject299 instantiates a new InlineObject299 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject299WithDefaults

`func NewInlineObject299WithDefaults() *InlineObject299`

NewInlineObject299WithDefaults instantiates a new InlineObject299 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineObject299) GetNetwork() OrganizationsOrganizationIdNetworksMovesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject299) GetNetworkOk() (*OrganizationsOrganizationIdNetworksMovesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject299) SetNetwork(v OrganizationsOrganizationIdNetworksMovesNetwork)`

SetNetwork sets Network field to given value.


### GetOrganizations

`func (o *InlineObject299) GetOrganizations() OrganizationsOrganizationIdNetworksMovesOrganizations`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *InlineObject299) GetOrganizationsOk() (*OrganizationsOrganizationIdNetworksMovesOrganizations, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *InlineObject299) SetOrganizations(v OrganizationsOrganizationIdNetworksMovesOrganizations)`

SetOrganizations sets Organizations field to given value.


### GetSimulate

`func (o *InlineObject299) GetSimulate() bool`

GetSimulate returns the Simulate field if non-nil, zero value otherwise.

### GetSimulateOk

`func (o *InlineObject299) GetSimulateOk() (*bool, bool)`

GetSimulateOk returns a tuple with the Simulate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulate

`func (o *InlineObject299) SetSimulate(v bool)`

SetSimulate sets Simulate field to given value.

### HasSimulate

`func (o *InlineObject299) HasSimulate() bool`

HasSimulate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


