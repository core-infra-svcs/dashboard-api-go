# NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Type of this performance class. Must be one of: &#39;builtin&#39; or &#39;custom&#39; | [optional] 
**BuiltinPerformanceClassName** | Pointer to **NullableString** | Name of builtin performance class. Must be present when performanceClass type is &#39;builtin&#39; and value must be one of: &#39;VoIP&#39; | [optional] 
**CustomPerformanceClassId** | Pointer to **NullableString** | ID of created custom performance class, must be present when performanceClass type is \&quot;custom\&quot; | [optional] 

## Methods

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass() *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClassWithDefaults

`func NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClassWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass`

NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClassWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetBuiltinPerformanceClassName() string`

GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field if non-nil, zero value otherwise.

### GetBuiltinPerformanceClassNameOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool)`

GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetBuiltinPerformanceClassName(v string)`

SetBuiltinPerformanceClassName sets BuiltinPerformanceClassName field to given value.

### HasBuiltinPerformanceClassName

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasBuiltinPerformanceClassName() bool`

HasBuiltinPerformanceClassName returns a boolean if a field has been set.

### SetBuiltinPerformanceClassNameNil

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetBuiltinPerformanceClassNameNil(b bool)`

 SetBuiltinPerformanceClassNameNil sets the value for BuiltinPerformanceClassName to be an explicit nil

### UnsetBuiltinPerformanceClassName
`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) UnsetBuiltinPerformanceClassName()`

UnsetBuiltinPerformanceClassName ensures that no value is present for BuiltinPerformanceClassName, not even an explicit nil
### GetCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetCustomPerformanceClassId() string`

GetCustomPerformanceClassId returns the CustomPerformanceClassId field if non-nil, zero value otherwise.

### GetCustomPerformanceClassIdOk

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool)`

GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetCustomPerformanceClassId(v string)`

SetCustomPerformanceClassId sets CustomPerformanceClassId field to given value.

### HasCustomPerformanceClassId

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasCustomPerformanceClassId() bool`

HasCustomPerformanceClassId returns a boolean if a field has been set.

### SetCustomPerformanceClassIdNil

`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetCustomPerformanceClassIdNil(b bool)`

 SetCustomPerformanceClassIdNil sets the value for CustomPerformanceClassId to be an explicit nil

### UnsetCustomPerformanceClassId
`func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) UnsetCustomPerformanceClassId()`

UnsetCustomPerformanceClassId ensures that no value is present for CustomPerformanceClassId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


