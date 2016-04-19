# Response Available for SignUp

## Success Case
Happens when user provide all data needed to create a user

| Name | Code | Description
| --- | --- | ---
| Created | 201 | User created successfully

## Fail Case
Happens when there is an error happen during signup process

| Name | Code | Description
| --- | --- | ---
| Username Missing | 400 | The username is missing or empty.
| Password Missing | 400 | The password is missing or empty.
| Username Taken | 400 | The username has already been taken.
