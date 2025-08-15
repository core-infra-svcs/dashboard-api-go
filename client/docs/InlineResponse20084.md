# InlineResponse20084

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentMode** | Pointer to **string** | Deployment mode for the cellular gateways in the network. (Passthrough/Routed) | [optional] 
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all MGs in  this network. | [optional] 
**Subnets** | Pointer to [**[]InlineResponse20084Subnets**](InlineResponse20084Subnets.md) | List of subnets of all MGs in this network. | [optional] 

## Methods

### NewInlineResponse20084

`func NewInlineResponse20084() *InlineResponse20084`

NewInlineResponse20084 instantiates a new InlineResponse20084 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20084WithDefaults

`func NewInlineResponse20084WithDefaults() *InlineResponse20084`

NewInlineResponse20084WithDefaults instantiates a new InlineResponse20084 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentMode

`func (o *InlineResponse20084) GetDeploymentMode() string`

GetDeploymentMode returns the DeploymentMode field if non-nil, zero value otherwise.

### GetDeploymentModeOk

`func (o *InlineResponse20084) GetDeploymentModeOk() (*string, bool)`

GetDeploymentModeOk returns a tuple with the DeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMode

`func (o *InlineResponse20084) SetDeploymentMode(v string)`

SetDeploymentMode sets DeploymentMode field to given value.

### HasDeploymentMode

`func (o *InlineResponse20084) HasDeploymentMode() bool`

HasDeploymentMode returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20084) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20084) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20084) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20084) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *InlineResponse20084) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineResponse20084) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineResponse20084) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineResponse20084) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20084) GetSubnets() []InlineResponse20084Subnets`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20084) GetSubnetsOk() (*[]InlineResponse20084Subnets, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20084) SetSubnets(v []InlineResponse20084Subnets)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20084) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


