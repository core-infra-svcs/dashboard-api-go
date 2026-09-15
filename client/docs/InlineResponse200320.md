# InlineResponse200320

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ID of the primary cluster | [optional] 
**Network** | Pointer to [**InlineResponse200319Network**](InlineResponse200319Network.md) |  | [optional] 
**Failover** | Pointer to [**OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover**](OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover.md) |  | [optional] 

## Methods

### NewInlineResponse200320

`func NewInlineResponse200320() *InlineResponse200320`

NewInlineResponse200320 instantiates a new InlineResponse200320 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200320WithDefaults

`func NewInlineResponse200320WithDefaults() *InlineResponse200320`

NewInlineResponse200320WithDefaults instantiates a new InlineResponse200320 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *InlineResponse200320) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *InlineResponse200320) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *InlineResponse200320) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *InlineResponse200320) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200320) GetNetwork() InlineResponse200319Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200320) GetNetworkOk() (*InlineResponse200319Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200320) SetNetwork(v InlineResponse200319Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200320) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFailover

`func (o *InlineResponse200320) GetFailover() OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *InlineResponse200320) GetFailoverOk() (*OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *InlineResponse200320) SetFailover(v OrganizationsOrganizationIdCampusGatewayClustersFailoverTargetsFailover)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *InlineResponse200320) HasFailover() bool`

HasFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


