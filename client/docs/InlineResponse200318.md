# InlineResponse200318

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the primary cluster | [optional] 
**Network** | Pointer to [**InlineResponse200317Network**](InlineResponse200317Network.md) |  | [optional] 
**Failover** | Pointer to [**OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover**](OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover.md) |  | [optional] 

## Methods

### NewInlineResponse200318

`func NewInlineResponse200318() *InlineResponse200318`

NewInlineResponse200318 instantiates a new InlineResponse200318 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200318WithDefaults

`func NewInlineResponse200318WithDefaults() *InlineResponse200318`

NewInlineResponse200318WithDefaults instantiates a new InlineResponse200318 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200318) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200318) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200318) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200318) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200318) GetNetwork() InlineResponse200317Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200318) GetNetworkOk() (*InlineResponse200317Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200318) SetNetwork(v InlineResponse200317Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200318) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFailover

`func (o *InlineResponse200318) GetFailover() OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200318) GetFailoverOk() (*OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200318) SetFailover(v OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200318) HasFailover() bool`

HasFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


