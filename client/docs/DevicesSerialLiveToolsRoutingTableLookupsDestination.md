# DevicesSerialLiveToolsRoutingTableLookupsDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The destination IP address to find a longest prefix match route | [optional] 
**Subnet** | Pointer to **string** | The destination subnet in CIDR format to find matching routes | [optional] 

## Methods

### NewDevicesSerialLiveToolsRoutingTableLookupsDestination

`func NewDevicesSerialLiveToolsRoutingTableLookupsDestination() *DevicesSerialLiveToolsRoutingTableLookupsDestination`

NewDevicesSerialLiveToolsRoutingTableLookupsDestination instantiates a new DevicesSerialLiveToolsRoutingTableLookupsDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsRoutingTableLookupsDestinationWithDefaults

`func NewDevicesSerialLiveToolsRoutingTableLookupsDestinationWithDefaults() *DevicesSerialLiveToolsRoutingTableLookupsDestination`

NewDevicesSerialLiveToolsRoutingTableLookupsDestinationWithDefaults instantiates a new DevicesSerialLiveToolsRoutingTableLookupsDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSubnet

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DevicesSerialLiveToolsRoutingTableLookupsDestination) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


