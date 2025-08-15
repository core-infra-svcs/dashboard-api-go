# DevicesSerialSwitchRoutingStaticRoutesVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VRF this static route belongs to | [optional] 
**LeakRouteToDefaultVrf** | Pointer to **bool** | Whether or not next-hop IP is reachable via default VRF | [optional] 

## Methods

### NewDevicesSerialSwitchRoutingStaticRoutesVrf

`func NewDevicesSerialSwitchRoutingStaticRoutesVrf() *DevicesSerialSwitchRoutingStaticRoutesVrf`

NewDevicesSerialSwitchRoutingStaticRoutesVrf instantiates a new DevicesSerialSwitchRoutingStaticRoutesVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialSwitchRoutingStaticRoutesVrfWithDefaults

`func NewDevicesSerialSwitchRoutingStaticRoutesVrfWithDefaults() *DevicesSerialSwitchRoutingStaticRoutesVrf`

NewDevicesSerialSwitchRoutingStaticRoutesVrfWithDefaults instantiates a new DevicesSerialSwitchRoutingStaticRoutesVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) GetLeakRouteToDefaultVrf() bool`

GetLeakRouteToDefaultVrf returns the LeakRouteToDefaultVrf field if non-nil, zero value otherwise.

### GetLeakRouteToDefaultVrfOk

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) GetLeakRouteToDefaultVrfOk() (*bool, bool)`

GetLeakRouteToDefaultVrfOk returns a tuple with the LeakRouteToDefaultVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) SetLeakRouteToDefaultVrf(v bool)`

SetLeakRouteToDefaultVrf sets LeakRouteToDefaultVrf field to given value.

### HasLeakRouteToDefaultVrf

`func (o *DevicesSerialSwitchRoutingStaticRoutesVrf) HasLeakRouteToDefaultVrf() bool`

HasLeakRouteToDefaultVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


