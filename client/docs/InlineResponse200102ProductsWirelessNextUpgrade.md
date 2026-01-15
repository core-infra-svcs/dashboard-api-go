# InlineResponse200102ProductsWirelessNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**ToVersion** | Pointer to [**InlineResponse200102ProductsWirelessNextUpgradeToVersion**](InlineResponse200102ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200102ProductsWirelessNextUpgrade

`func NewInlineResponse200102ProductsWirelessNextUpgrade() *InlineResponse200102ProductsWirelessNextUpgrade`

NewInlineResponse200102ProductsWirelessNextUpgrade instantiates a new InlineResponse200102ProductsWirelessNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102ProductsWirelessNextUpgradeWithDefaults

`func NewInlineResponse200102ProductsWirelessNextUpgradeWithDefaults() *InlineResponse200102ProductsWirelessNextUpgrade`

NewInlineResponse200102ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse200102ProductsWirelessNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) GetToVersion() InlineResponse200102ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse200102ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse200102ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200102ProductsWirelessNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


