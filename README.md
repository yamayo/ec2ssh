# ec2ssh
A tool to easily login to EC2 instances selected by [peco](https://github.com/peco/peco).  

## Demo


## Installation
On macOS, you can use Homebrew:
```
$ brew tap yamayo/ec2ssh
$ brew install ec2ssh
```

On Linux, Unix download [binary]().


## Usage
```
$ ec2ssh [option]
```

### Options
#### `--profile`  
Use a specific profile from your credential file. (default `default`)

#### `--region`  
The region to use. Overrides AWS config/env settings.

#### `--user`  
Specifies the user to login to EC2 machine. (default `ec2-user`)

#### `--version`  
Show version.

## Usage Example
```
$ ec2ssh --profile myprofile --region ap-northeast-1
or
$ ec2ssh --user centos
```

And `Environment Variables`, `Named Profiles` support.  
```
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=XXX AWS_REGION=us-east-1 ec2ssh
```
See: http://docs.aws.amazon.com/ja_jp/cli/latest/userguide/cli-chap-getting-started.html#cli-multiple-profiles
