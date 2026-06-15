# InlineResponse200133ProductsWirelessNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**Predownload** | Pointer to [**InlineResponse200133ProductsWirelessNextUpgradePredownload**](InlineResponse200133ProductsWirelessNextUpgradePredownload.md) |  | [optional] 
**ToVersion** | Pointer to [**InlineResponse200133ProductsWirelessNextUpgradeToVersion**](InlineResponse200133ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200133ProductsWirelessNextUpgrade

`func NewInlineResponse200133ProductsWirelessNextUpgrade() *InlineResponse200133ProductsWirelessNextUpgrade`

NewInlineResponse200133ProductsWirelessNextUpgrade instantiates a new InlineResponse200133ProductsWirelessNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200133ProductsWirelessNextUpgradeWithDefaults

`func NewInlineResponse200133ProductsWirelessNextUpgradeWithDefaults() *InlineResponse200133ProductsWirelessNextUpgrade`

NewInlineResponse200133ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse200133ProductsWirelessNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetPredownload() InlineResponse200133ProductsWirelessNextUpgradePredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetPredownloadOk() (*InlineResponse200133ProductsWirelessNextUpgradePredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) SetPredownload(v InlineResponse200133ProductsWirelessNextUpgradePredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetToVersion() InlineResponse200133ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse200133ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse200133ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200133ProductsWirelessNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


