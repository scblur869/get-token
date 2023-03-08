# Get-token provides temporary credentials for accessing services based on an assumed AWS IAM role
## it requires a valid aws_client_id and aws_client_secret from an IAM User and a role to assume


## build
```console
go build -o get-token -ldflags "-s -w"
```
## run
```console
get-token
 --client-id=AKIA2TMCDD........
 --client-secret=YI12aYUnOQ0ODcs.......
 --role-arn=arn:aws:iam::7XXXXXXXXXXXX:role/some-access-role
 --region=us-east-1  (default) 
 --credentials-file-path=text  (default /)
```
## output
```console
AWS Assume Role Credentials
Expires : 2023-03-08 23:57:48 +0000 UTC

[default]
aws_access_key_id = ASIA2TMCD12312342524
aws_secret_access_key = CiJDSlEwe18PqJB8pDzjsb1234567784VAVguCp2
aws_session_token = IQoJb3JpZ2luX2VjEJj4523434564wEaCXVzLWVhc3Qt....
```

## creates a file by default called "./credentials" based on the following format

```console
[default]
aws_access_key_id = ASIA2TMCDD1234567890
aws_secret_access_key = JSDUpsR4u32NoWwG6vJtNj0L1234567890111213
aws_session_token = IQoJb3JpZ2luX2VjEJj////sdfsadfsadfgy4h4tghtg
```

## example use case
    -- automate the provisioning of greengrass core instances as containers
    -- any edge application that requires an assume role temp credential to access cloud resources


