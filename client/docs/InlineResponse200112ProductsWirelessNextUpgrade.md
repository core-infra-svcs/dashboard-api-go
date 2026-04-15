# InlineResponse200112ProductsWirelessNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**Predownload** | Pointer to [**InlineResponse200112ProductsWirelessNextUpgradePredownload**](InlineResponse200112ProductsWirelessNextUpgradePredownload.md) |  | [optional] 
**ToVersion** | Pointer to [**InlineResponse200112ProductsWirelessNextUpgradeToVersion**](InlineResponse200112ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200112ProductsWirelessNextUpgrade

`func NewInlineResponse200112ProductsWirelessNextUpgrade() *InlineResponse200112ProductsWirelessNextUpgrade`

NewInlineResponse200112ProductsWirelessNextUpgrade instantiates a new InlineResponse200112ProductsWirelessNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200112ProductsWirelessNextUpgradeWithDefaults

`func NewInlineResponse200112ProductsWirelessNextUpgradeWithDefaults() *InlineResponse200112ProductsWirelessNextUpgrade`

NewInlineResponse200112ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse200112ProductsWirelessNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetPredownload() InlineResponse200112ProductsWirelessNextUpgradePredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetPredownloadOk() (*InlineResponse200112ProductsWirelessNextUpgradePredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) SetPredownload(v InlineResponse200112ProductsWirelessNextUpgradePredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetToVersion() InlineResponse200112ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse200112ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse200112ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200112ProductsWirelessNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


