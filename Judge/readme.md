# Judge Manager

## Introduction

Judge Manager is a management of judges in LXC/LXD, who accept http request to judge problems and notify them when judge begins and judge results.
Judge Manager can set up many kinds of compilers and many LXC containers for each compilers so that we can judge many submission at the same time.
When a judge request come, Judge Manager will select a free container to judge it, and restore the container to a snapshot when judge end.
Judge Manager will install the replyment packages and softwares automatically, and there will be a high experience for user.

## Usage

Befor Run Judge Manager, you must install lxd first:
```
sudo add-apt-repository ppa:ubuntu-lxc/lxd-git-master && sudo apt-get update
sudo apt-get install lxd
```

And then run Judge Manager is ok.

Judge Manager listen to http://127.0.0.1:2020 and it accept request body like below:
```
    problemId=2&
    api=http://127.0.0.1:1234/problem/"&
    compiler=gcc&
    submission=#include <stdio.h>%10int main()%10{%10printf("Hello world!");%10return 0;%10}"&
    notify=http://127.0.0.1/notify
```
`problemId`, `api`, `compiler`, and `submission` will be pass to container, and `notify` is the address to notify judge start and result.
Judge Manager will get `notify`/problem/`problemId`/start to notify the start, and post a json like below to `notify`/problem/`problemId`/result:
```
{
    "code":   0,
    "msg":    "ok",
    "body":   ResponseBody{
        Accepted:   "false",
        ErrorType:      "Time Limit Error",
        ErrorMsg:       "",
        MemoryUse:      126,
        TimeUse:        999,
    },
}
```

## License

All rights reserved by InsZVA