# InlineResponse200136ProductsApplianceNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**ToVersion** | Pointer to [**InlineResponse200136ProductsWirelessNextUpgradeToVersion**](InlineResponse200136ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200136ProductsApplianceNextUpgrade

`func NewInlineResponse200136ProductsApplianceNextUpgrade() *InlineResponse200136ProductsApplianceNextUpgrade`

NewInlineResponse200136ProductsApplianceNextUpgrade instantiates a new InlineResponse200136ProductsApplianceNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200136ProductsApplianceNextUpgradeWithDefaults

`func NewInlineResponse200136ProductsApplianceNextUpgradeWithDefaults() *InlineResponse200136ProductsApplianceNextUpgrade`

NewInlineResponse200136ProductsApplianceNextUpgradeWithDefaults instantiates a new InlineResponse200136ProductsApplianceNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) GetToVersion() InlineResponse200136ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) GetToVersionOk() (*InlineResponse200136ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) SetToVersion(v InlineResponse200136ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200136ProductsApplianceNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


