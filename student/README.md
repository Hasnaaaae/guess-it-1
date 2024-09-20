# <span style="color:magenta; size : 20px"> Guess-it-1 </span>

### <span style="color:pink"> Objectives
The goal of this project is to create a program that predicts a range for the next number based on a series of given inputs. Based on previous mathematical skills the program should take input numbers and determine a range within which the next number should be.

### <span style="color:pink"> Author 

- Hasnae Lamrani

### <span style="color:pink"> How to run ?

To run your program, use the following command:
```
$ go run . 
```

### <span style="color:pink"> How to test ?

- After downloading the test file, copy the contents of the file into the working directory.

- Install docker by those commands: 
```
$ curl -fsSL https://get.docker.com/rootless | sh 
```
```
$ export PATH=$HOME/bin:$PATH
```
```
$ export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
```

- Run this command to have the dependencies needed and to start the webpage on the port 3000:
```
$ docker-compose up
```

- After opening your browser of preference in the port [3000](http://localhost:3000/)

- Add in the URL a guesser , in other words, the name of one of the files present in the `ai/` folder:
```console
?guesser=<name_of_guesser>
```
For example:

```console
?guesser=big-range
```

- After that, choose which of the random data set to test, click `Quick` to skip the waiting and be presented with the results.

- lear the displays clicking on the `Clean` button after each test.


### <span style="color:pink"> Learning Outcomes

This project will enhance our understanding of:

- Statistical and Probabilities Calculation
- Scripting and automation