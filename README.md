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


## Document in MongoDB
	{	
		"name": "Peter Parker",
		"userName": "Peter"
		"email": "peter@gmail.com"
		"phone": "0123567876",
		"password": "964409EF3D55E7CCD626DA4E0443AA7341FE51F51BBA4B81A92499E1C9D5DB7E [lkQJ7z1V58zWJtpOBEOqc0H+UfUbukuBqSSZ4cnV234=]"
		"salt":"12aAF14yOqdj"
		"social": {"facebook": `FacebookInfo`, "twitter": `TwitterInfo`, ...}
		"extra": {
			"address": "LosA",
			"age": 20
		}
	}
