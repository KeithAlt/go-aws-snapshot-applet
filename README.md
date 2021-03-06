# go-aws-snapshot-applet
A small go app. intended for an AWS EC2 instance. Performs a snapshot of an AWS RDS database with a configured .env
___
<h3>DEPENDENCIES / USE:</h3>
· Requires aws-sdk-go
(https://github.com/aws/aws-sdk-go)

· Requires godotenv
(https://github.com/joho/godotenv)

· Requires active AWS RDS database

You will also need to have an AWS account with appropriate security groups configured for use of the AWS SDK.
Instructions on how to set this up can be found here: https://aws.github.io/aws-sdk-go-v2/docs/getting-started/
___
<h3>SETUP INSTRUCTIONS:</h3>

**1.)** Configure the .env with your AWS RDS database info
```
AWS_DB_ID=<INSERT_HERE>
AWS_CONFIG_REGION=<INSERT_HERE>
```
  
**2.)** cd into _'./go-aws-snapshot-applet'_

**3.)** run in TERMINAL:
```
$ go get
```
___
Upon run/build of the app; will create a new snapshot of the specified database.
Feel free to fork or modify for your own use.
