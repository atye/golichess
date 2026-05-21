# UserActivityPostsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopicUrl** | **string** |  | 
**TopicName** | **string** |  | 
**Posts** | [**[]UserActivityPostsInnerPostsInner**](UserActivityPostsInnerPostsInner.md) |  | 

## Methods

### NewUserActivityPostsInner

`func NewUserActivityPostsInner(topicUrl string, topicName string, posts []UserActivityPostsInnerPostsInner, ) *UserActivityPostsInner`

NewUserActivityPostsInner instantiates a new UserActivityPostsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityPostsInnerWithDefaults

`func NewUserActivityPostsInnerWithDefaults() *UserActivityPostsInner`

NewUserActivityPostsInnerWithDefaults instantiates a new UserActivityPostsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopicUrl

`func (o *UserActivityPostsInner) GetTopicUrl() string`

GetTopicUrl returns the TopicUrl field if non-nil, zero value otherwise.

### GetTopicUrlOk

`func (o *UserActivityPostsInner) GetTopicUrlOk() (*string, bool)`

GetTopicUrlOk returns a tuple with the TopicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicUrl

`func (o *UserActivityPostsInner) SetTopicUrl(v string)`

SetTopicUrl sets TopicUrl field to given value.


### GetTopicName

`func (o *UserActivityPostsInner) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *UserActivityPostsInner) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *UserActivityPostsInner) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.


### GetPosts

`func (o *UserActivityPostsInner) GetPosts() []UserActivityPostsInnerPostsInner`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *UserActivityPostsInner) GetPostsOk() (*[]UserActivityPostsInnerPostsInner, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *UserActivityPostsInner) SetPosts(v []UserActivityPostsInnerPostsInner)`

SetPosts sets Posts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


