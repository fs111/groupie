# groupie simple group by for the command line

groupie does one simple thing. Read lines from stdin, counts and groups them. When stdin is closed it prints the groups
and their counts found in the stream. This is very useful for a quick frequency analyis like so:

```
$ some-command |  groupie
1 foo
4 bar
7 quux
```

groupie does only one thing well. If you need sorted output pipe its output to `sort -n` or `sort -nr`.

That's it. Yes, one can do `sort | uniq -c | sort -n`, but groupie is faster since you need one sort less.


## Building

Any recent version of go will work. Simply run:

```
$ go build
```


## License - MIT

Copyright 2022 André Kelpe

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
