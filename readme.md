# A Online Judge Based On LXD

## Intrduction

This is a judge system based on LXD, which means all program submitted will run in a Linux Container for safety.
It contains 2 parts, Judge for manage Daemons on LXDs, and serve for Application Server, and Daemon to judge specified submissions.

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

# Daemon For LXD

## Introduction

Daemon is a judgement tool running in LXD, who get a http request and judge, then response.
For Example, I post a command to judge a C Program, I just tell it the compiler is "gcc", where it can get problemInformation, problemId, and a submission.
Then it will tell me whether the source was accepted or compiler error, (even time limit error), and the time memory usage.

## Usage

You must use HTTP POST request like this for request body:
```
{
    "problemId":2,
    "api":"http://127.0.0.1:1234/problem/",
    "compiler":"gcc",
    "submission":"#include <stdio.h>\nint main()\n{\nprintf("Hello world!");\nreturn 0;\n}"
}
```
The api address is that will be call by Judge with adding problemId to its end to get Standard Input And Answer of this problem.
For above, the judge will call `http://127.0.0.1:1234/problem/2` to find informations.So the api address must response somting like below:
```
{
    "code":0,
    "msg":"ok",
    "body":{
        "stdin":"1",
        "stdout":"1\n1"
    }
}
```
The stdout is the standard answer for specified stdin, to judge whether a source is ok.

## Taste

Now, the function is not full and stable, but you can try it by run `check/server/test.go` to run a api server,
then run `main.go` and run `client.go` to try it, also you can edit something to run by your thougths.

## License

All rights reserved By InsZVA