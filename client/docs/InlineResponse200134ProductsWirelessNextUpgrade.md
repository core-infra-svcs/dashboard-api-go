# InlineResponse200134ProductsWirelessNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**Strategy** | Pointer to **string** | The strategy that network devices will use to perform the upgrade. If unspecified, the scheduled upgrade will use the default strategy for the network. | [optional] 
**Predownload** | Pointer to [**InlineResponse200134ProductsWirelessNextUpgradePredownload**](InlineResponse200134ProductsWirelessNextUpgradePredownload.md) |  | [optional] 
**ToVersion** | Pointer to [**InlineResponse200134ProductsWirelessNextUpgradeToVersion**](InlineResponse200134ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200134ProductsWirelessNextUpgrade

`func NewInlineResponse200134ProductsWirelessNextUpgrade() *InlineResponse200134ProductsWirelessNextUpgrade`

NewInlineResponse200134ProductsWirelessNextUpgrade instantiates a new InlineResponse200134ProductsWirelessNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200134ProductsWirelessNextUpgradeWithDefaults

`func NewInlineResponse200134ProductsWirelessNextUpgradeWithDefaults() *InlineResponse200134ProductsWirelessNextUpgrade`

NewInlineResponse200134ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse200134ProductsWirelessNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStrategy

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetPredownload() InlineResponse200134ProductsWirelessNextUpgradePredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetPredownloadOk() (*InlineResponse200134ProductsWirelessNextUpgradePredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) SetPredownload(v InlineResponse200134ProductsWirelessNextUpgradePredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetToVersion() InlineResponse200134ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse200134ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse200134ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200134ProductsWirelessNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


