# go-aws-snapshot-applet
A small go app. intended for an AWS EC2 instance. Performs a snapshot of an AWS RDS database with a configured .env

INSTRUCTIONS:
- Configure the .env with your AWS RDS database info
- cd into './go-aws-snapshot-applet'
- run: $go get

Upon run of the app; will create a new snapshot of the specified database.
Feel free to fork or modify for your own use.
