# InlineResponse200125ProductsWirelessNextUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Timestamp of the next scheduled firmware upgrade | [optional] 
**Predownload** | Pointer to [**InlineResponse200125ProductsWirelessNextUpgradePredownload**](InlineResponse200125ProductsWirelessNextUpgradePredownload.md) |  | [optional] 
**ToVersion** | Pointer to [**InlineResponse200125ProductsWirelessNextUpgradeToVersion**](InlineResponse200125ProductsWirelessNextUpgradeToVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200125ProductsWirelessNextUpgrade

`func NewInlineResponse200125ProductsWirelessNextUpgrade() *InlineResponse200125ProductsWirelessNextUpgrade`

NewInlineResponse200125ProductsWirelessNextUpgrade instantiates a new InlineResponse200125ProductsWirelessNextUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200125ProductsWirelessNextUpgradeWithDefaults

`func NewInlineResponse200125ProductsWirelessNextUpgradeWithDefaults() *InlineResponse200125ProductsWirelessNextUpgrade`

NewInlineResponse200125ProductsWirelessNextUpgradeWithDefaults instantiates a new InlineResponse200125ProductsWirelessNextUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetPredownload() InlineResponse200125ProductsWirelessNextUpgradePredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetPredownloadOk() (*InlineResponse200125ProductsWirelessNextUpgradePredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) SetPredownload(v InlineResponse200125ProductsWirelessNextUpgradePredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetToVersion() InlineResponse200125ProductsWirelessNextUpgradeToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) GetToVersionOk() (*InlineResponse200125ProductsWirelessNextUpgradeToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) SetToVersion(v InlineResponse200125ProductsWirelessNextUpgradeToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200125ProductsWirelessNextUpgrade) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


