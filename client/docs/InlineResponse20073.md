# InlineResponse20073

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Dashboard network ID | [optional] 
**Name** | Pointer to **string** | Dashboard network name | [optional] 
**Enabled** | Pointer to **bool** | Is warm spare enabled | [optional] 
**Mode** | Pointer to **string** | HA mode (disabled|active-passive|active-active) | [optional] 
**Designations** | Pointer to [**[]InlineResponse20073Designations**](InlineResponse20073Designations.md) | Ordered HA roles (1&#x3D;primary, 2&#x3D;spare) | [optional] 
**Uplink** | Pointer to [**InlineResponse20073Uplink**](InlineResponse20073Uplink.md) |  | [optional] 

## Methods

### NewInlineResponse20073

`func NewInlineResponse20073() *InlineResponse20073`

NewInlineResponse20073 instantiates a new InlineResponse20073 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073WithDefaults

`func NewInlineResponse20073WithDefaults() *InlineResponse20073`

NewInlineResponse20073WithDefaults instantiates a new InlineResponse20073 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse20073) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20073) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20073) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20073) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20073) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20073) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20073) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20073) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20073) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20073) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20073) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20073) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse20073) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20073) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20073) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20073) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDesignations

`func (o *InlineResponse20073) GetDesignations() []InlineResponse20073Designations`

GetDesignations returns the Designations field if non-nil, zero value otherwise.

### GetDesignationsOk

`func (o *InlineResponse20073) GetDesignationsOk() (*[]InlineResponse20073Designations, bool)`

GetDesignationsOk returns a tuple with the Designations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignations

`func (o *InlineResponse20073) SetDesignations(v []InlineResponse20073Designations)`

SetDesignations sets Designations field to given value.

### HasDesignations

`func (o *InlineResponse20073) HasDesignations() bool`

HasDesignations returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20073) GetUplink() InlineResponse20073Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20073) GetUplinkOk() (*InlineResponse20073Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20073) SetUplink(v InlineResponse20073Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20073) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


