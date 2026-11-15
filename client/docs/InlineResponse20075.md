# InlineResponse20075

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Dashboard network ID | [optional] 
**Name** | Pointer to **string** | Dashboard network name | [optional] 
**Enabled** | Pointer to **bool** | Is warm spare enabled | [optional] 
**Mode** | Pointer to **string** | HA mode (disabled|active-passive|active-active) | [optional] 
**Designations** | Pointer to [**[]InlineResponse20075Designations**](InlineResponse20075Designations.md) | Ordered HA roles (1&#x3D;primary, 2&#x3D;spare) | [optional] 
**Uplink** | Pointer to [**InlineResponse20075Uplink**](InlineResponse20075Uplink.md) |  | [optional] 

## Methods

### NewInlineResponse20075

`func NewInlineResponse20075() *InlineResponse20075`

NewInlineResponse20075 instantiates a new InlineResponse20075 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20075WithDefaults

`func NewInlineResponse20075WithDefaults() *InlineResponse20075`

NewInlineResponse20075WithDefaults instantiates a new InlineResponse20075 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse20075) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20075) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20075) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20075) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20075) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20075) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20075) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20075) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20075) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20075) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20075) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20075) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse20075) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20075) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20075) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20075) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDesignations

`func (o *InlineResponse20075) GetDesignations() []InlineResponse20075Designations`

GetDesignations returns the Designations field if non-nil, zero value otherwise.

### GetDesignationsOk

`func (o *InlineResponse20075) GetDesignationsOk() (*[]InlineResponse20075Designations, bool)`

GetDesignationsOk returns a tuple with the Designations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignations

`func (o *InlineResponse20075) SetDesignations(v []InlineResponse20075Designations)`

SetDesignations sets Designations field to given value.

### HasDesignations

`func (o *InlineResponse20075) HasDesignations() bool`

HasDesignations returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20075) GetUplink() InlineResponse20075Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20075) GetUplinkOk() (*InlineResponse20075Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20075) SetUplink(v InlineResponse20075Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20075) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


