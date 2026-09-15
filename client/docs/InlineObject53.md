# InlineObject53

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable warm spare | 
**Mode** | Pointer to **string** | HA mode (disabled|active-passive|active-active) | [optional] 
**Designations** | Pointer to [**[]NetworksNetworkIdApplianceDevicesRedundancyDesignations**](NetworksNetworkIdApplianceDevicesRedundancyDesignations.md) | Ordered warm spare roles | [optional] 
**Uplink** | Pointer to [**NetworksNetworkIdApplianceDevicesRedundancyUplink**](NetworksNetworkIdApplianceDevicesRedundancyUplink.md) |  | [optional] 

## Methods

### NewInlineObject53

`func NewInlineObject53(enabled bool, ) *InlineObject53`

NewInlineObject53 instantiates a new InlineObject53 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject53WithDefaults

`func NewInlineObject53WithDefaults() *InlineObject53`

NewInlineObject53WithDefaults instantiates a new InlineObject53 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject53) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject53) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject53) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMode

`func (o *InlineObject53) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject53) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject53) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject53) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDesignations

`func (o *InlineObject53) GetDesignations() []NetworksNetworkIdApplianceDevicesRedundancyDesignations`

GetDesignations returns the Designations field if non-nil, zero value otherwise.

### GetDesignationsOk

`func (o *InlineObject53) GetDesignationsOk() (*[]NetworksNetworkIdApplianceDevicesRedundancyDesignations, bool)`

GetDesignationsOk returns a tuple with the Designations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignations

`func (o *InlineObject53) SetDesignations(v []NetworksNetworkIdApplianceDevicesRedundancyDesignations)`

SetDesignations sets Designations field to given value.

### HasDesignations

`func (o *InlineObject53) HasDesignations() bool`

HasDesignations returns a boolean if a field has been set.

### GetUplink

`func (o *InlineObject53) GetUplink() NetworksNetworkIdApplianceDevicesRedundancyUplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineObject53) GetUplinkOk() (*NetworksNetworkIdApplianceDevicesRedundancyUplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineObject53) SetUplink(v NetworksNetworkIdApplianceDevicesRedundancyUplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineObject53) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


