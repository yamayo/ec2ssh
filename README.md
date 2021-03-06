# ec2ssh
A CLI tool to easily ssh login to EC2 instances selected by [peco](https://github.com/peco/peco).  

## Installation
On macOS, you can use Homebrew:
```
$ brew tap yamayo/ec2ssh
$ brew install ec2ssh
```

On Linux, download [binary](https://github.com/yamayo/ec2ssh/releases).

## Usage
```
$ ec2ssh [option]
```

### Options
#### `--profile, -p`  
Use a specific profile from your credential file. (default `default`)

#### `--region, -r`  
The region to use. Overrides AWS config/env settings.

#### `--user, -u`  
Specifies the user to login to EC2 machine. (default `ec2-user`)

#### `--version`  
Show version.

## Settings and Examples

### AWS Credentials
Before using ec2ssh you need to first give it your AWS credentials.

#### Named Profiles
The following example shows a credentials file with two profiles:  
`~/.aws/credentials`  
```
[default]
aws_access_key_id = xxxxx
aws_secret_access_key = xxxxxxxxxx

[user2]
aws_access_key_id = xxxxx
aws_secret_access_key = xxxxxxxxxx
```

`~/.aws/config`  
```
[default]
region = us-west-2

[profile user2]
region = us-east-1
```

To use a named profile other than `default`, add the `--profile` option to your command.  
```
$ ec2ssh --profile user2
```

Or you can specify it with an environment variable as `AWS_DEFAULT_PROFILE`.  
```
$ export AWS_DEFAULT_PROFILE=user2
```

#### Environment Variables
ec2ssh supports the following environment variables.

- **AWS_ACCESS_KEY_ID** – AWS access key.

- **AWS_SECRET_ACCESS_KEY** – AWS secret key.

- **AWS_SESSION_TOKEN** – Only required if you are using temporary security credentials.

- **AWS_DEFAULT_REGION** – AWS region.

- **AWS_DEFAULT_PROFILE** – name of the CLI profile to use.

- **AWS_CONFIG_FILE** – path to a CLI config file.

See more: http://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-getting-started.html

### Saving your private key
Finally, do not forget save the private key for login to EC2 in `~/.ssh`.
