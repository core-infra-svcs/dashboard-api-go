# DevicesSerialSwitchRoutingStaticRoutesVrf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VRF this static route belongs to | [optional] 
**LeakRouteToDefaultVrf** | Pointer to **bool** | Whether or not next-hop IP is reachable via default VRF | [optional] 

## Methods

### NewDevicesSerialSwitchRoutingStaticRoutesVrf1

`func NewDevicesSerialSwitchRoutingStaticRoutesVrf1() *DevicesSerialSwitchRoutingStaticRoutesVrf1`

NewDevicesSerialSwitchRoutingStaticRoutesVrf1 instantiates a new DevicesSerialSwitchRoutingStaticRoutesVrf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchRoutingStaticRoutesVrf1WithDefaults

`func NewDevicesSerialSwitchRoutingStaticRoutesVrf1WithDefaults() *DevicesSerialSwitchRoutingStaticRoutesVrf1`

NewDevicesSerialSwitchRoutingStaticRoutesVrf1WithDefaults instantiates a new DevicesSerialSwitchRoutingStaticRoutesVrf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) GetLeakRouteToDefaultVrf() bool`

GetLeakRouteToDefaultVrf returns the LeakRouteToDefaultVrf field if non-nil, zero value otherwise.

### GetLeakRouteToDefaultVrfOk

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) GetLeakRouteToDefaultVrfOk() (*bool, bool)`

GetLeakRouteToDefaultVrfOk returns a tuple with the LeakRouteToDefaultVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) SetLeakRouteToDefaultVrf(v bool)`

SetLeakRouteToDefaultVrf sets LeakRouteToDefaultVrf field to given value.

### HasLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf1) HasLeakRouteToDefaultVrf() bool`

HasLeakRouteToDefaultVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


