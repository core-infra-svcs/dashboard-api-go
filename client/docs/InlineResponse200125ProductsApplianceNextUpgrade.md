# InlineResponse200125ProductsApplianceNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**ToVersion** | Pointer to [**InlineResponse200125ProductsWirelessNextUpgradeToVersion**](InlineResponse200125ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200125ProductsApplianceNextUpgrade

`func NewInlineResponse200125ProductsApplianceNextUpgrade() *InlineResponse200125ProductsApplianceNextUpgrade`

NewInlineResponse200125ProductsApplianceNextUpgrade instantiates a new InlineResponse200125ProductsApplianceNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125ProductsApplianceNextUpgradeWithDefaults

`func NewInlineResponse200125ProductsApplianceNextUpgradeWithDefaults() *InlineResponse200125ProductsApplianceNextUpgrade`

NewInlineResponse200125ProductsApplianceNextUpgradeWithDefaults instantiates a new InlineResponse200125ProductsApplianceNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) GetToVersion() InlineResponse200125ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) GetToVersionOk() (*InlineResponse200125ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) SetToVersion(v InlineResponse200125ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200125ProductsApplianceNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


