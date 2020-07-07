# zendesk-perma-delete
Permanently delete users in Zendesk in the Deleted Users list

## Instructions

Clone this repository and run go build. In your command line, run:

```
./zendesk-perma-delete -url="<ZENDESK-SUBDOMAIN>" -user="<ZENDESK-API-USER>" -key="<ZENDESK-API-KEY>"
```

You should then see the IDs of each deleted user logged in your terminal.

```
EXAMPLE:

User ID 396768650754 Deleted Successfully
User ID 397173667254 Deleted Successfully
User ID 397201585553 Deleted Successfully
User ID 397780969254 Deleted Successfully
User ID 398619041694 Deleted Successfully
User ID 399366560954 Deleted Successfully
User ID 399382214274 Deleted Successfully
User ID 401130574013 Deleted Successfully
User ID 401366335014 Deleted Successfully
User ID 402228095194 Deleted Successfully
User ID 402542539013 Deleted Successfully
User ID 402542548693 Deleted Successfully
User ID 402735366174 Deleted Successfully
User ID 406053155094 Deleted Successfully
```
