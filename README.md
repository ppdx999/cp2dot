# cp2dot
Generate dot file from competitive programming input

# usage

```terminal
$ cat input.txt
1 2
2 3
2 3
3 1
5 4
5 5

$ cp2dot < input.txt
digraph G {
    1 -> 2;
    2 -> 3;
    2 -> 3;
    3 -> 1;
    5 -> 4;
    5 -> 5;
}
```

To cooperate with graphviz

```
$ cat input.txt | cp2dot | dot -Tsvg -o output.png
```