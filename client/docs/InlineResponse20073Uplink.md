# InlineResponse20073Uplink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Uplink mode (public|virtual) | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse20073UplinkInterfaces**](InlineResponse20073UplinkInterfaces.md) | List of WAN interfaces | [optional] 
**Sharing** | Pointer to [**NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing**](NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing.md) |  | [optional] 

## Methods

### NewInlineResponse20073Uplink

`func NewInlineResponse20073Uplink() *InlineResponse20073Uplink`

NewInlineResponse20073Uplink instantiates a new InlineResponse20073Uplink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073UplinkWithDefaults

`func NewInlineResponse20073UplinkWithDefaults() *InlineResponse20073Uplink`

NewInlineResponse20073UplinkWithDefaults instantiates a new InlineResponse20073Uplink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20073Uplink) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20073Uplink) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20073Uplink) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20073Uplink) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse20073Uplink) GetInterfaces() []InlineResponse20073UplinkInterfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse20073Uplink) GetInterfacesOk() (*[]InlineResponse20073UplinkInterfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse20073Uplink) SetInterfaces(v []InlineResponse20073UplinkInterfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse20073Uplink) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetSharing

`func (o *InlineResponse20073Uplink) GetSharing() NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *InlineResponse20073Uplink) GetSharingOk() (*NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *InlineResponse20073Uplink) SetSharing(v NetworksNetworkIdApplianceDevicesRedundancyUplinkSharing)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *InlineResponse20073Uplink) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


