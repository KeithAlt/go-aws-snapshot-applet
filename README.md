# go-aws-snapshot-applet
A small go app. intended for an AWS EC2 instance. Performs a snapshot of an AWS RDS database with a configured .env
___
<h3>INSTRUCTIONS:</h3>

**1.)** Configure the .env with your AWS RDS database info
```
AWS_DB_ID=<INSERT REGION>
AWS_CONFIG_REGION=<INSERT REGION>
```
  
**2.)** cd into _'./go-aws-snapshot-applet'_

**3.)** run in TERMINAL:
```
$go get
```
___
Upon run of the app; will create a new snapshot of the specified database.
Feel free to fork or modify for your own use.
