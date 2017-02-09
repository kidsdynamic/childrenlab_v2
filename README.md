## Overview API

### User API
* [POST   /v1/user/login](#v1userlogin---post)
* [POST   /v1/user/register      ](#v1userregister---post)
* [GET   /v1/user/isTokenValid  ](#v1useristokenvalid---get)
* [GET   /v1/user/isEmailAvailableToRegister  ](#v1userisemailavailabletoregister---get)
* [PUT    /v1/user/updateProfile ](#v1userupdateprofile---put)
* [GET    /v1/user/retrieveUserProfile](#v1userretrieveuserprofile---get)
* [PUT    /v1/user/updateIOSRegistrationId](#v1userupdateiosregistrationid---put)
* [GET    /v1/user/findByEmail](#v1userfindbyemail---get)


### Kid API
* [POST   /v1/kids/add             ](#v1kidsadd---post)
* [PUT    /v1/kids/update          ](#v1kidsupdate---put)
* [GET    /v1/kids/list          ](#v1kidslist---get)
* [GET    /v1/kids/whoRegisteredMacID          ](#v1kidswhoregisteredmacid---get)
* [DELETE    /v1/kids/delete          ](#v1kidsdelete---delete)

### Avatar
* [POST   /v1/user/avatar/upload   ](#v1useravatarupload---post)
* [POST   /v1/user/avatar/uploadKid](#v1useravataruploadkid---post)

### Activity
* [POST   /v1/activity/uploadRawData](#v1activityuploadrawdata---post)
* [GET   /v1/activity/retrieveData](#v1activityretrievedata---get)
* [GET   /v1/activity/retrieveDataByTime](#v1activityretrievedatabytime---get)

### Event
* [POST   /v1/event/add](#v1eventadd---post)
* [PUT   /v1/event/update](#v1eventupdate---put)
* [DELETE   /v1/event/delete](#v1eventdelete---delete)
* [GET   /v1/event/retrieveEvents](#v1eventretrieveevents---get)
* [GET   /v1/event/retrieveAllEventsWithTodo](#v1eventretrievealleventswithtodo---get)
* [GET   /v1/event/retrieveAllEventsByKid](#v1eventretrieveeventsbykid---get)

#### TODO
* [PUT   /v1/event/todo/done](#v1eventtododone---put)

### Multi-Host API
* [POST   /v1/subHost/add](#v1subhostadd---post)
* [PUT   /v1/subHost/accept](#v1subhostaccept---put)
* [PUT   /v1/subHost/deny](#v1subhostdeny---put)
* [GET   /v1/subHost/list](#v1subhostlist---get)

## /v1/user/login - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| :------------- |:-------------:|:-------------:| :-----|
| email     | Yes | String |   test@kidsdynami.com |
| password     | Yes | String |   aaaaaa |

#### Response Status
| Status Code    | Meaning      |
| :------------- |:-------------|
| 200     | Success Login |
| 400     | Bad Request. Login failed |

* Success response body:
```
{
  "access_token": "2f69e1d858a04b4f35dd477c76921b34",
  "username": "teste"
}
```

## /v1/user/isTokenValid - GET

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |
| token     | Yes | String |   2f69e1d858a04b4f35dd477c76921b34 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | The token is valid |
| 403     | Forbidden. The token is not valid |
* The API doesn't return any JSON

## /v1/user/register - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |
| password     | Yes | String |   aaaaaa |
| firstName     | Yes | String |   Jay |
| lastName     | Yes | String |   Chen |
| phoneNumber     | No | String |   3442314231 |
| zipCode     | No | String |   11101 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Register success |
| 400     | Bad request. Missing some parameters |
| 409     | Conflict. The email is already registered |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it doesn't return JSON
* Fail - response body:
```
{
  "message": "The email is already registered"
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/user/isEmailAvailableToRegister - GET
* The API doesn't return any body

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | The email is able to register |
| 409     | Conflict. The email is already registered |

### curl
```
curl -X GET "http://localhost:8111/v1/user/isEmailAvailableToRegister?email=jack083001@gmail.com"
```



## /v1/user/register - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |
| password     | Yes | String |   aaaaaa |
| firstName     | Yes | String |   Jay |
| lastName     | Yes | String |   Chen |
| phoneNumber     | No | String |   3442314231 |
| zipCode     | No | String |   11101 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Register success |
| 400     | Bad request. Missing some parameters |
| 409     | Conflict. The email is already registered |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it doesn't return JSON
* Fail - response body:
```
{
  "message": "The email is already registered"
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/user/updateProfile - PUT
* Content-Type: application/json
* The API doesn't support email and password update

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| firstName    | No | String |   Jay |
| lastName     | No | String |   Chen |
| phoneNumber  | No | String |   3442314231 |
| zipCode      | No | String |   11101 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |
* Success - 
```
{
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "KIDLLE",
    "lastName": "YES",
    "lastUpdate": "2016-12-18T21:24:57Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "3444943214",
    "profile": ""
  }
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/user/retrieveUserProfile - GET
* Retrieve user data by the header token
* No any parameter required, just need x-auth-token on the header
* It retrieves with kids data that belong to the user

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '' "http://localhost:8111/v1/user/retrieveUserProfile"
```
* Success - with kids data
```
{
  "kids": [
    {
      "id": 18,
      "name": "Jay",
      "dateCreated": "2016-12-11T22:37:15Z",
      "macId": "13031FCFE5E02",
      "profile": ""
    },
    {
      "id": 19,
      "name": "KIDLLE",
      "dateCreated": "2016-12-18T04:17:35Z",
      "macId": "hgweorahgbkljwhnpi",
      "profile": ""
    },
    {
      "id": 20,
      "name": "KIDLLE",
      "dateCreated": "2016-12-18T21:19:54Z",
      "macId": "hgweorahgbkljwhnpi2",
      "profile": ""
    }
  ],
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "name": "KIDLLE",
    "lastUpdate": "2016-12-18T21:24:57Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  }
}
```

## /v1/user/updateIOSRegistrationId - PUT
* It's for IOS registration ID - push notification

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| registrationId     | Yes | String |   TestRegistrationID |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update the registration id successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" -d '{
  "registrationId": "TestRegistrationID"
}' "http://localhost:8111/v1/user/updateIOSRegistrationId"
```
* Success - Return user profile
```
{
  "id": 6,
  "email": "jack08300@gmail.com",
  "firstName": "JJJ",
  "lastName": "TTT",
  "lastUpdate": "2017-01-11T03:17:31Z",
  "dateCreated": "2017-01-11T03:17:31Z",
  "zipCode": "11111",
  "phoneNumber": "",
  "profile": "",
  "registrationId": "TestRegistrationID"
}
```

## /v1/user/findByEmail - GET
* Retrieve user data by the email

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Get user successfully |
| 400     | Bad request. The email parameter is missing |
| 404     | Not found. Can't find user by the email |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 4b0d53ff1953af73aff2fec2478bd1fe" "http://localhost:8111/v1/user/findByEmail?email=jack0830421@gmail.com"
```
* Success -
```
{
  "id": 2,
  "email": "jack08300@gmail.com",
  "firstName": "Jay",
  "lastName": "Chen",
  "lastUpdate": "0001-01-01T00:00:00Z",
  "dateCreated": "0001-01-01T00:00:00Z",
  "zipCode": "",
  "phoneNumber": "",
  "profile": ""
}
```

## /v1/user/avatar/upload - POST
* Upload user profile image
* Content-Type: multipart/form-data
* The file will be on Amazon S3 - ask me the path
* The file name will be /userProfile/avatar_{userID}

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| upload     | Yes | file |    |

### curl Example
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637"  -H "Content-Type: multipart/form-data; -F "upload=@CwEiSDRVMAE4IO4.jpg" "http://localhost:8111/v1/user/avatar/upload"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | File upload successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it updates profile property
```
{
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "KIDLLE",
    "lastName": "YES",
    "lastUpdate": "2016-12-18T04:11:02Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": "avatar_29.jpg"
  }
}
```
* Bad request error - response body:
```
{
  "error": {
    "ErrorString": "request Content-Type isn't multipart/form-data"
  },
  "message": "upload parameter is required"
}
```

## /v1/user/avatar/uploadKid - POST
* Upload kid profile image
* Content-Type: multipart/form-data
* The file will be on Amazon S3 - ask me the path
* The file name will be /userProfile/kid_avatar_{kidID}

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| upload     | Yes | File |    |
| kidId      | Yes | Integer | |

### curl Example
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: multipart/form-data; -F "kidId=19" -F "upload=@CwEiSDRVMAE4IO4.jpg" "http://localhost:8111/v1/user/avatar/uploadKid"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | File upload successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it updates profile property
```
{
  "kid": {
    "id": 13,
    "name": "kid13",
    "dateCreated": "2017-01-29T23:07:38Z",
    "macId": "Mac_ID3",
    "profile": "kid_avatar_13.jpg",
    "parent": {
      "id": 5,
      "email": "jack08301@gmail.com",
      "name": "Jay",
      "lastUpdate": "2017-01-11T04:16:44Z",
      "dateCreated": "2017-01-11T04:16:44Z",
      "zipCode": "11111",
      "phoneNumber": "11111",
      "profile": "avatar_5.jpg",
      "registrationId": "123test"
    }
  }
}
```
* Bad request error - response body:
```
{
  "error": {
    "ErrorString": "request Content-Type isn't multipart/form-data"
  },
  "message": "upload parameter is required"
}
```

## /v1/kids/add - POST
* Add kid
* It only responses added kid
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| name     | Yes | string | Kiddle   |
| macId      | Yes | string | 8D071FCFE5E0 |

### curl Example
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: 8158050a26569ff9c8c5297351c4f56a" -d '{
  "name": "Kids3",
  "macId": "Mac_ID4"
}' "http://localhost:8111/v1/kids/add"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Data added successfully |
| 400     | Bad request. Missing some parameters |
| 409     | Conflict. The mac ID is already registered |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it returns all of kids belong to the user
```
{
  "id": 2,
  "name": "Kids3",
  "dateCreated": "2017-02-02T01:14:16Z",
  "macId": "Mac_ID4",
  "profile": "",
  "parent": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "",
    "phoneNumber": "",
    "profile": ""
  }
}
```
* Conflict error - response body:
```
{
  "message": "The device is already registered"
}
```

## /v1/kids/update - PUT
* Content-Type: application/json
* Update kid info

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| kidId    | Yes | String |   19 |
| name    | No | String |   Jay |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: application/json" -d '{
	"lastName": "yes",
	"kidId": 19
}' "http://localhost:8111/v1/kids/update"
```

* Success - 
```
{
  "kid": {
    "id": 13,
    "name": "kid13",
    "dateCreated": "2017-01-29T23:07:38Z",
    "macId": "Mac_ID3",
    "profile": "",
    "parent": {
      "id": 5,
      "email": "jack08301@gmail.com",
      "firstName": "JJJ",
      "lastName": "TTT",
      "lastUpdate": "2017-01-11T04:16:44Z",
      "dateCreated": "2017-01-11T04:16:44Z",
      "zipCode": "11111",
      "phoneNumber": "412412412412",
      "profile": "avatar_5.jpg"
    }
  }
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/kids/delete - DELETE
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***
* It can only delete logged in user's kid data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| kidId    | Yes | Integer |   3 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X DELETE -H "Content-Type: application/json" -H "x-auth-token: 1c8306065183d7829fab224be6e43257" "http://localhost:8111/v1/kids/delete?kidId=2"
```

* Success - no response body

## /v1/kids/list - GET
* Retrieve kids belong to the signed in user
* Doesn't need any parameter

#### Request Parameters
None


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 403     | User is not signed in |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 28e20ffb974205c0747b5aa35d53e53" "http://localhost:8111/v1/kids/list"
```

* Success - 
```
[
  {
    "id": 3,
    "name": "kid13",
    "dateCreated": "2017-02-02T10:25:20Z",
    "macId": "012345678915",
    "profile": "kid_avatar_3.jpg"
  }
]
```

## /v1/kids/whoRegisteredMacID - GET
* Retrieve kid and user information by MAC ID
* Doesn't need x-auth-token

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| macId    | Yes | String |   hgweorahgbkljwhnpi2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 404     | The user not found which means the MAC ID is not registered by anyone |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/kids/whoRegisteredMacID?macId=hgweorahgbkljwhnpi2"
```

* Success - 
```
{
  "kid": {
    "id": 2,
    "name": "KIDLLE123124",
    "dateCreated": "2017-01-16T22:18:26Z",
    "macId": "hgweorahgbkljwhnpi1",
    "profile": "qwe",
    "parent": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "11111",
      "phoneNumber": "11111",
      "profile": "",
      "registrationId": "123test"
    }
  }
}
```

## /v1/activity/uploadRawData - POST
* Content-Type: application/json
* Activity data
* When the data duplicate which is when server response ***409*** status, you can ***ignore*** it and process next data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| indoorActivity      | Yes | String | 1481299119,0,216,2,3,4 |
| outdoorActivity     | Yes | String | 1481299119,1,0,0,0,0 |
| time                | Yes | Long |   1470885849 |
| macId               | Yes | String |   hgweorahgbkljwhnpi2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 409     | Conflict. The data is already exist |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: application/json" -d '{
	"indoorActivity": "1481299119,0,216,2,3,4",
	"outdoorActivity": "1481299119,1,0,0,0,0",
	"time": 1470885849,
	"macId": "hgweorahgbkljwhnpi2"
}' "http://localhost:8111/v1/activity/uploadRawData"
```

* Success - No any JSON response

* 409 status - duplicate data:
```
{
  "message": "This is a duplicate data"
}
```


## /v1/activity/retrieveData - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId      | Yes | Integer | 20 |
| period     | Yes | String | DAILY, WEEKLY MONTHLY, YEARLY  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/activity/dailyData?kidId=20&period=DAILY"
```

* Success - 
```
{
  "activities": [
    {
      "id": 121,
      "macId": "hgweorahgbkljwhnpi2",
      "kidId": "20",
      "type": "INDOOR",
      "steps": 3298,
      "distance": 0,
      "receivedDate": "2016-12-13T19:55:02Z"
    },
    {
      "id": 122,
      "macId": "hgweorahgbkljwhnpi2",
      "kidId": "20",
      "type": "OUTDOOR",
      "steps": 42,
      "distance": 0,
      "receivedDate": "2016-12-13T19:55:02Z"
    }
  ]
}
```

* 400 Bad Request - 
```
{
  "error": {
    "Func": "ParseInt",
    "Num": "",
    "Err": {}
  },
  "message": "kidId should be int type."
}
```

## /v1/activity/retrieveDataByTime - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| start      | Yes | Long ***Timestamp*** | 1491599032 |
| end        | Yes | Long ***Timestamp*** | 1498089090  |
| kidId      | Yes | Integer | 1  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: 27e217ae4d2907188b8c92cdcf9c85ac" "http://localhost:8111/v1/activity/retrieveDataByTime?start=1491599032&end=1498089090&kidId=1"
```

* Success - 
```
{
  "activities": [
    {
      "id": 5,
      "macId": "hgweorahgbkljwhnpi3",
      "kidId": "1",
      "type": "INDOOR",
      "steps": 10,
      "distance": 0,
      "receivedDate": "2017-05-03T10:58:10Z"
    },
    {
      "id": 6,
      "macId": "hgweorahgbkljwhnpi3",
      "kidId": "1",
      "type": "OUTDOOR",
      "steps": 24,
      "distance": 0,
      "receivedDate": "2017-06-21T23:51:30Z"
    }
  ]
}
```


## /v1/event/add - POST
* Content-Type: application/json
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId             | Yes | Integer | 20 |
| name               | Yes | String | Test event name |
| startDate           | Yes | String |   2015-08-30T08:20:00 |
| endDate             | Yes | String |   2015-08-31T08:20:00 |
| color               | Yes | String |   #F05D25 |
| description         | No | String |   Test description |
| alert               | No | Integer |   49 |
| city               | No | String |   New York |
| state               | No | String |   New York |
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1", "test todo 2"] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Added successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "kidId": 20,
  "Name": "Test event name",
  "startDate": "2015-08-30T08:20:00",
  "endDate": "2015-08-31T08:20:00",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "Hahah",
  "alert": 49,
  "city": "New York",
  "state": "New York",
  "todo": [
  	"test todo 1", "test todo 2"	
  ]
  
}' "http://localhost:8111/v1/event/add"
```

* Success - Returns added event
```
{
  "event": {
    "id": 398,
    "userId": 29,
    "kidId": 20,
    "name": "Test event name",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2016-12-26T06:25:19Z",
    "lastUpdated": "2016-12-26T06:25:19Z",
    "todo": [
      {
        "id": 38,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2016-12-26T06:25:19Z",
        "lastUpdated": "2016-12-26T06:25:19Z"
      },
      {
        "id": 39,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2016-12-26T06:25:19Z",
        "lastUpdated": "2016-12-26T06:25:19Z"
      }
    ]
  }
}
```

## /v1/event/update - PUT
* Content-Type: application/json
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***
* Send the parameter even user does not change it

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId             | Yes | Integer | 414 |
| name               | Yes | String | Test event name |
| startDate           | Yes | String |   2015-08-30T08:20:00 |
| endDate             | Yes | String |   2015-08-31T08:20:00 |
| color               | Yes | String |   #F05D25 |
| description         | No | String |   Test description |
| alert               | No | Integer |   49 |
| city               | No | String |   New York |
| state               | No | String |   New York |
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1"] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Updated successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "eventId": 414,
  "Name": "Test event name2",
  "startDate": "2015-08-30T08:20:00",
  "endDate": "2015-08-31T08:20:00",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "Hahah",
  "alert": 49,
  "city": "New York",
  "state": "New York",
  "todo": [
  	"test todo 1"	
  ]
  
}' "http://localhost:8111/v1/event/update"
```

* Success - Returns updated event
```
{
  "event": {
    "id": 414,
    "userId": 29,
    "kidId": 20,
    "name": "Test event name2",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2016-12-26T21:21:53Z",
    "lastUpdated": "2016-12-27T00:59:28Z",
    "todo": [
      {
        "id": 83,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2016-12-27T00:59:28Z",
        "lastUpdated": "2016-12-27T00:59:28Z"
      }
    ]
  }
}
```

## /v1/event/delete - DELETE
* URL Query

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId             | Yes | Integer | 414 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X DELETE -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" "http://localhost:8111/v1/event/delete?eventId=2"
```

* Success - Returns empty json
```
{}
```

## /v1/event/retrieveEvents - GET
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***
* If trying to retrieve month events, do not use '00' as month value. Example: 2016-12-01T00:00:00 to retrieve 2016-12 events.


#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| period        | Yes | String | MONTH, DAY |
| date          | Yes | String | 2016-12-26T01:00:00 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/event/retrieveEvents?period=MONTH&date=2016-12-01T00:00:00"
```

* Success - Returns updated event
```
{
  "event": [
    {
      "id": 410,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-27T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:30Z",
      "lastUpdated": "2016-12-26T21:17:30Z",
      "todo": [
        {
          "id": 66,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:30Z",
          "lastUpdated": "2016-12-26T21:17:30Z"
        },
        {
          "id": 67,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:30Z",
          "lastUpdated": "2016-12-26T21:17:30Z"
        }
      ]
    },
    {
      "id": 411,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:35Z",
      "lastUpdated": "2016-12-26T21:17:35Z",
      "todo": [
        {
          "id": 68,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:35Z",
          "lastUpdated": "2016-12-26T21:17:35Z"
        },
        {
          "id": 69,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:35Z",
          "lastUpdated": "2016-12-26T21:17:35Z"
        }
      ]
    },
    {
      "id": 412,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "",
      "state": "",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:45Z",
      "lastUpdated": "2016-12-26T21:17:45Z",
      "todo": [
        {
          "id": 70,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:45Z",
          "lastUpdated": "2016-12-26T21:17:45Z"
        },
        {
          "id": 71,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:45Z",
          "lastUpdated": "2016-12-26T21:17:45Z"
        }
      ]
    },
    {
      "id": 413,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:19:16Z",
      "lastUpdated": "2016-12-26T21:19:16Z",
      "todo": [
        {
          "id": 72,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:19:16Z",
          "lastUpdated": "2016-12-26T21:19:16Z"
        },
        {
          "id": 73,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:19:16Z",
          "lastUpdated": "2016-12-26T21:19:16Z"
        }
      ]
    }
  ]
}
```

## /v1/event/retrieveAllEventsWithTodo - GET
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/event/retrieveAllEvents"
```

* Success - Returns events
```
[
  {
    "id": 1,
    "userId": 3,
    "kidId": 3,
    "name": "Test event name",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "DAILY",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-09T12:22:46Z",
    "lastUpdated": "2017-02-09T12:22:46Z",
    "todo": [
      {
        "id": 1,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:46Z",
        "lastUpdated": "2017-02-09T17:22:46Z"
      },
      {
        "id": 2,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:46Z",
        "lastUpdated": "2017-02-09T17:22:46Z"
      }
    ]
  },
  {
    "id": 2,
    "userId": 3,
    "kidId": 3,
    "name": "Test event name",
    "startDate": "2016-08-30T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "DAILY",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-09T12:22:58Z",
    "lastUpdated": "2017-02-09T12:22:58Z",
    "todo": [
      {
        "id": 3,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:58Z",
        "lastUpdated": "2017-02-09T17:22:58Z"
      },
      {
        "id": 4,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:58Z",
        "lastUpdated": "2017-02-09T17:22:58Z"
      }
    ]
  }
]
```

## /v1/event/retrieveAllEventsByKid - GET

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 403     | The user doesn't have access to the kid
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 42142142421421312312312" "http://localhost:8111/v1/event/retrieveEventsByKid?kidId=3"```

* Success - Returns events
```
[
  {
    "id": 1,
    "userId": 3,
    "kidId": 3,
    "name": "Test event name",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "DAILY",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-09T12:22:46Z",
    "lastUpdated": "2017-02-09T12:22:46Z",
    "todo": [
      {
        "id": 1,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:46Z",
        "lastUpdated": "2017-02-09T17:22:46Z"
      },
      {
        "id": 2,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:46Z",
        "lastUpdated": "2017-02-09T17:22:46Z"
      }
    ]
  },
  {
    "id": 2,
    "userId": 3,
    "kidId": 3,
    "name": "Test event name",
    "startDate": "2016-08-30T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "DAILY",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-09T12:22:58Z",
    "lastUpdated": "2017-02-09T12:22:58Z",
    "todo": [
      {
        "id": 3,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:58Z",
        "lastUpdated": "2017-02-09T17:22:58Z"
      },
      {
        "id": 4,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-09T17:22:58Z",
        "lastUpdated": "2017-02-09T17:22:58Z"
      }
    ]
  }
]
```


## /v1/event/todo/done - PUT
* Set status = done to todo

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId        | Yes | Integer | 2 |
| todoId          | Yes | Integer | 2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | updated successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" -H "Content-Type: application/json" -d '{
  "eventId": 2,
  "todoId": 2
}' "http://localhost:8111/v1/event/todo/done"
```

* Success - 
```
{}
```


## /v1/subHost/add - POST
* Send sub host request to the host account
* Content-Type: application/json
* Get ```hostId``` from [/v1/kids/whoRegisteredMacID](#v1kidswhoregisteredmacid---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| hostId            | Yes | Integer | 35 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Added successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 409     | Conflict. The request is already exists |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: 4094345fecef667df936a632bafabbe6" -d '{
  "hostId": 63
}' "http://localhost:8111/v1/subHost/add"
```

* Success - Returns added request
```
{
  "id": 3,
  "requestFromUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "11111",
    "phoneNumber": "11111",
    "profile": "",
    "registrationId": "123test"
  },
  "requestToUser": {
    "id": 63,
    "email": "jack08301@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-01-17T00:56:06Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "PENDING",
  "createdDate": "2017-01-18T04:03:22Z",
  "lastUpdated": "2017-01-18T04:03:22Z"
}
```

* Fail (409) -
```
{
  "message": "The request is already exist"
}
```

## /v1/subHost/accept - PUT
* Content-Type: application/json
* Accept the sub host request by allowing sub host to view the kids
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId             | Yes | Integer | 10 |
| KidId             | Yes | Integer Array | [22, 10] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Accept successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong. Or the logged in user doesn't have permission |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: 6d312cbc54ce05a279dcc554b984061c" -d '{
  "subHostId": 1,
  "kidId": [
  	9, 10
  ]
}' "http://localhost:8111/v1/subHost/accept"
```

* Success - Returns updated request
```
{
  "id": 3,
  "requestFromUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "11111",
    "phoneNumber": "11111",
    "profile": "",
    "registrationId": "123test"
  },
  "requestToUser": {
    "id": 63,
    "email": "jack08301@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-01-17T00:56:06Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "ACCEPTED",
  "createdDate": "2017-01-18T04:03:22Z",
  "lastUpdated": "2017-01-18T04:03:22Z",
  "kids": [
    {
      "id": 9,
      "name": "Another KID",
      "dateCreated": "2017-01-17T00:56:58Z",
      "macId": "Test2",
      "profile": ""
    },
    {
      "id": 10,
      "name": "KIDLLE123124",
      "dateCreated": "2017-01-17T02:39:36Z",
      "macId": "Test3",
      "profile": ""
    }
  ]
}
```

## /v1/subHost/deny - PUT
* Content-Type: application/json
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId             | Yes | Integer | 2 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Accept successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: 58lkp329ejbr4498st59ur2na7e0rmtg" -d '{
  "subHostId": 2
}' "http://localhost:8111/v1/subHost/deny"
```

* Success - Returns updated request
```
{
  "id": 3,
  "requestFromUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "11111",
    "phoneNumber": "11111",
    "profile": "",
    "registrationId": "123test"
  },
  "requestToUser": {
    "id": 63,
    "email": "jack08301@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-01-17T00:56:06Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "DENIED",
  "createdDate": "2017-01-18T04:03:22Z",
  "lastUpdated": "2017-01-17T23:16:14.456787738-05:00"
}
```

## /v1/subHost/list - GET
* Content-Type: application/json
* If no status parameter, the API returns ALL of sub host belong to the user

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| status        | No | String | PENDING, ACCEPTED, DENIED |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receive list successfully |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 58lkp329ejbr4498st59ur2na7e0rmtg" "http://localhost:8111/v1/subHost/list"
```

* Success - Returns updated request
```
[
  {
    "id": 3,
    "requestFromUser": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "11111",
      "phoneNumber": "11111",
      "profile": "",
      "registrationId": "123test"
    },
    "requestToUser": {
      "id": 63,
      "email": "jack08301@gmail.com",
      "firstName": "JJJ",
      "lastName": "TTT",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "2017-01-17T00:56:06Z",
      "zipCode": "11111",
      "phoneNumber": "",
      "profile": ""
    },
    "status": "DENIED",
    "createdDate": "2017-01-18T04:03:22Z",
    "lastUpdated": "2017-01-18T04:16:14Z",
    "kids": [
      {
        "id": 9,
        "name": "Another KID",
        "dateCreated": "2017-01-17T00:56:58Z",
        "macId": "Test2",
        "profile": ""
      },
      {
        "id": 10,
        "name": "KIDLLE123124",
        "dateCreated": "2017-01-17T02:39:36Z",
        "macId": "Test3",
        "profile": ""
      }
    ]
  }
]
```