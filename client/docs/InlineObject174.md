# InlineObject174

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the VoIP provider | 
**Address** | **string** | The IP address (IPv4 only) or hostname of the media server to monitor | 
**BestEffortMonitoringEnabled** | Pointer to **bool** | Indicates that if the media server doesn&#39;t respond to ICMP pings, the nearest hop will be used in its stead. | [optional] 

## Methods

### NewInlineObject174

`func NewInlineObject174(name string, address string, ) *InlineObject174`

NewInlineObject174 instantiates a new InlineObject174 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject174WithDefaults

`func NewInlineObject174WithDefaults() *InlineObject174`

NewInlineObject174WithDefaults instantiates a new InlineObject174 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject174) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject174) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject174) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *InlineObject174) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineObject174) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineObject174) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBestEffortMonitoringEnabled

`func (o *InlineObject174) GetBestEffortMonitoringEnabled() bool`

GetBestEffortMonitoringEnabled returns the BestEffortMonitoringEnabled field if non-nil, zero value otherwise.

### GetBestEffortMonitoringEnabledOk

`func (o *InlineObject174) GetBestEffortMonitoringEnabledOk() (*bool, bool)`

GetBestEffortMonitoringEnabledOk returns a tuple with the BestEffortMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEffortMonitoringEnabled

`func (o *InlineObject174) SetBestEffortMonitoringEnabled(v bool)`

SetBestEffortMonitoringEnabled sets BestEffortMonitoringEnabled field to given value.

### HasBestEffortMonitoringEnabled

`func (o *InlineObject174) HasBestEffortMonitoringEnabled() bool`

HasBestEffortMonitoringEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

