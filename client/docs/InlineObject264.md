# InlineObject264

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | **string** | The private key and certificate used to inspect HTTPS traffic. The certificate must be in .pem format. | 
**Serials** | **[]string** | Serial numbers of the security appliances that will receive the new HTTPS certificate for HTTPS Inspection | 

## Methods

### NewInlineObject264

`func NewInlineObject264(contents string, serials []string, ) *InlineObject264`

NewInlineObject264 instantiates a new InlineObject264 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject264WithDefaults

`func NewInlineObject264WithDefaults() *InlineObject264`

NewInlineObject264WithDefaults instantiates a new InlineObject264 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *InlineObject264) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *InlineObject264) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *InlineObject264) SetContents(v string)`

SetContents sets Contents field to given value.


### GetSerials

`func (o *InlineObject264) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject264) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject264) SetSerials(v []string)`

SetSerials sets Serials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


