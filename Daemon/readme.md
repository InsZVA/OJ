# Daemon For LXD

## Introduction

Daemon is a judgement tool running in LXD, who get a http request and judge, then response.
For Example, I post a command to judge a C Program, I just tell it the compiler is "gcc", where it can get problemInformation, problemId, and a submission.
Then it will tell me whether the source was accepted or compiler error, (even time limit error), and the time memory usage.

## Usage

You must use HTTP POST request like this for request body:
```
    problemId=2&
    api=http://127.0.0.1:1234/problem/"&
    compiler=gcc&
    submission=#include <stdio.h>%10int main()%10{%10printf("Hello world!");%10return 0;%10}"&
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