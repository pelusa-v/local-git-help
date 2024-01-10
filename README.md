# Local Git help
Tool to change local git account easily
## Configure local environment:
- Create a configuration file in "~" path with your git accounts:

```json
[
    {
        "user_email": "email1@sample1.com",
        "user_name": "sample-user-1",
        "tag": "Github Personal"
    },
    {
        "user_email": "email2@sample2.com",
        "user_name": "sample-user-2",
        "tag": "Github Work"
    },
    {
        "user_email": "email2@sample2.com",
        "user_name": "sample-user-3",
        "tag": "Bitucket Work"
    }
]
```

## Using:
- Clone repository
- Inside root folder of repository, run the following command: ```go run cmd/main.go```
- You look a menu and select the local Git account to use:
```
Select local github user to use (↑ or ↓), 'q' or 'Q' for exit
1) email1@sample1.com / sample-user-1 (Github Personal)
2) email2@sample2.com / sample-user-2 (Github Work)   <--------
3) email2@sample2.com / sample-user-3 (Bitucket Work)
```