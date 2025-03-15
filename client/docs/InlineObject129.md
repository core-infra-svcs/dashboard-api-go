# InlineObject129

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | **[]string** | ids of applications to be installed | 
**Force** | Pointer to **bool** | By default, installation of an app which is believed to already be present on the device will be skipped. If you&#39;d like to force the installation of the app, set this parameter to true. | [optional] 

## Methods

### NewInlineObject129

`func NewInlineObject129(appIds []string, ) *InlineObject129`

NewInlineObject129 instantiates a new InlineObject129 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject129WithDefaults

`func NewInlineObject129WithDefaults() *InlineObject129`

NewInlineObject129WithDefaults instantiates a new InlineObject129 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *InlineObject129) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *InlineObject129) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *InlineObject129) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.


### GetForce

`func (o *InlineObject129) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *InlineObject129) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *InlineObject129) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *InlineObject129) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


