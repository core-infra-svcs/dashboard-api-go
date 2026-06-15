# InlineResponse200133ProductsApplianceNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**ToVersion** | Pointer to [**InlineResponse200133ProductsWirelessNextUpgradeToVersion**](InlineResponse200133ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200133ProductsApplianceNextUpgrade

`func NewInlineResponse200133ProductsApplianceNextUpgrade() *InlineResponse200133ProductsApplianceNextUpgrade`

NewInlineResponse200133ProductsApplianceNextUpgrade instantiates a new InlineResponse200133ProductsApplianceNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200133ProductsApplianceNextUpgradeWithDefaults

`func NewInlineResponse200133ProductsApplianceNextUpgradeWithDefaults() *InlineResponse200133ProductsApplianceNextUpgrade`

NewInlineResponse200133ProductsApplianceNextUpgradeWithDefaults instantiates a new InlineResponse200133ProductsApplianceNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) GetToVersion() InlineResponse200133ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) GetToVersionOk() (*InlineResponse200133ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) SetToVersion(v InlineResponse200133ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200133ProductsApplianceNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


