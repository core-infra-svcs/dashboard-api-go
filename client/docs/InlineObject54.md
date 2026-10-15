# InlineObject54

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enable warm spare | 
**Mode** | Pointer to **string** | HA mode (disabled|active-passive|active-active) | [optional] 
**Designations** | Pointer to [**[]NetworksNetworkIdApplianceDevicesRedundancyDesignations**](NetworksNetworkIdApplianceDevicesRedundancyDesignations.md) | Ordered warm spare roles | [optional] 
**Uplink** | Pointer to [**NetworksNetworkIdApplianceDevicesRedundancyUplink**](NetworksNetworkIdApplianceDevicesRedundancyUplink.md) |  | [optional] 

## Methods

### NewInlineObject54

`func NewInlineObject54(enabled bool, ) *InlineObject54`

NewInlineObject54 instantiates a new InlineObject54 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject54WithDefaults

`func NewInlineObject54WithDefaults() *InlineObject54`

NewInlineObject54WithDefaults instantiates a new InlineObject54 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject54) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject54) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject54) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMode

`func (o *InlineObject54) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject54) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject54) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject54) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDesignations

`func (o *InlineObject54) GetDesignations() []NetworksNetworkIdApplianceDevicesRedundancyDesignations`

GetDesignations returns the Designations field if non-nil, zero value otherwise.

### GetDesignationsOk

`func (o *InlineObject54) GetDesignationsOk() (*[]NetworksNetworkIdApplianceDevicesRedundancyDesignations, bool)`

GetDesignationsOk returns a tuple with the Designations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignations

`func (o *InlineObject54) SetDesignations(v []NetworksNetworkIdApplianceDevicesRedundancyDesignations)`

SetDesignations sets Designations field to given value.

### HasDesignations

`func (o *InlineObject54) HasDesignations() bool`

HasDesignations returns a boolean if a field has been set.

### GetUplink

`func (o *InlineObject54) GetUplink() NetworksNetworkIdApplianceDevicesRedundancyUplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineObject54) GetUplinkOk() (*NetworksNetworkIdApplianceDevicesRedundancyUplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineObject54) SetUplink(v NetworksNetworkIdApplianceDevicesRedundancyUplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineObject54) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


