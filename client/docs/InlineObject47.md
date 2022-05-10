# InlineObject47

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The VLAN ID of the new VLAN (must be between 1 and 4094) | 
**Name** | **string** | The name of the new VLAN | 
**Subnet** | Pointer to **string** | The subnet of the VLAN | [optional] 
**ApplianceIp** | Pointer to **string** | The local IP of the appliance on the VLAN | [optional] 
**GroupPolicyId** | Pointer to **string** | The id of the desired group policy to apply to the VLAN | [optional] 

## Methods

### NewInlineObject47

`func NewInlineObject47(id string, name string, ) *InlineObject47`

NewInlineObject47 instantiates a new InlineObject47 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject47WithDefaults

`func NewInlineObject47WithDefaults() *InlineObject47`

NewInlineObject47WithDefaults instantiates a new InlineObject47 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject47) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject47) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject47) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InlineObject47) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject47) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject47) SetName(v string)`

SetName sets Name field to given value.


### GetSubnet

`func (o *InlineObject47) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineObject47) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineObject47) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineObject47) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetApplianceIp

`func (o *InlineObject47) GetApplianceIp() string`

GetApplianceIp returns the ApplianceIp field if non-nil, zero value otherwise.

### GetApplianceIpOk

`func (o *InlineObject47) GetApplianceIpOk() (*string, bool)`

GetApplianceIpOk returns a tuple with the ApplianceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceIp

`func (o *InlineObject47) SetApplianceIp(v string)`

SetApplianceIp sets ApplianceIp field to given value.

### HasApplianceIp

`func (o *InlineObject47) HasApplianceIp() bool`

HasApplianceIp returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineObject47) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineObject47) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineObject47) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineObject47) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


