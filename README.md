# Not golang experts: Gopherstalker

## API Endpoints

### Registrations

#### POST `/registrations`

Request body:

``` json
{
  "user" : {
    "email" : "email@example.com",
    "password" : "supersecret",
    "password_confirmation" : "supersecret"
  }
}
```

##### Response codes

Status `201`

``` json
{
  "token" : "authtoken"
}
```

Status `422`

Error messages:

* Email has already been taken
* Passwords don't match

```json
{
  "error" : "error message"
}
```

### Sessions

#### POST `/sessions`

Request body:

``` json
{
  "user" : {
    "email" : "email@example.com",
    "password" : "supersecret"
  }
}
```

##### Response codes

Status `201`

``` json
{
  "token" : "authtoken"
}
```

Status `422`

Error messages:

* Invalid email or password

```json
{
  "error" : "error message"
}
```

#### DELETE `/sessions?token=yourauthtoken`

##### Response codes

Status `201`

``` json
{
  "token" : "authtoken"
}
```

Status `422`

**This request does not require body**

Error messages:

* Invalid email or password

```json
{
  "error" : "error message"
}
```
