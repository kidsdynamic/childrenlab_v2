## /v1/user/login -POST
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

## /v1/user/isTokenValid - POST
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

