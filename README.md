# Get-token assumes an AWS IAM role and provides temporary credentials for accessing services on behalf of that role


```console
get-token
 --client-id=AKIA2TMCDD........
 --client-secret=YI12aYUnOQ0ODcs.......
 --role-arn=arn:aws:iam::7XXXXXXXXXXXX:role/some-access-role 
 --credentials-file-path=text
```
## creates a file called "credentials" based on the following format

```console
[default]
aws_access_key_id = ASIA2TMCDD1234567890
aws_secret_access_key = JSDUpsR4u32NoWwG6vJtNj0L1234567890111213
aws_session_token = IQoJb3JpZ2luX2VjEJj////sdfsadfsadfgy4h4tghtg
```

## example use case
    -- automate the provisioning of greengrass core instances as containers
    -- any edge application that requires an assume role temp credentials
