Descrition User Document
========================


| Property | Type | Description | Notes 
| --- | --- | --- | --- |
| name | string | Name of user | required
| userName | string | username to login | required
| email | string | using email to verify account | required
| phone | string | phone number | optional
| password | string | password for authentication | required, encoding base 64
| social | map[string]interface{} | information of social as facebook, google, twitter, linkedIn | optional
