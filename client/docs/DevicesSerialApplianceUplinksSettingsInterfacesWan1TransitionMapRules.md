# DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **NullableString** | How MAP rules are obtained. HTTPS requires an ISP supported for the transition mode. | [optional] 
**Entries** | Pointer to [**[]DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries**](DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries.md) | Manual MAP rules. At most one entry is supported. | [optional] 

## Methods

### NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules

`func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules() *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules`

NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesWithDefaults

`func NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesWithDefaults() *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules`

NewDevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesWithDefaults instantiates a new DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetEntries

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) GetEntries() []DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) GetEntriesOk() (*[]DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) SetEntries(v []DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *DevicesSerialApplianceUplinksSettingsInterfacesWan1TransitionMapRules) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


