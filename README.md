Adam Cameron Code Challenge 2014-11-07
======================================

On 2014-11-07 Adam Cameron issued a code challenge on his blog found at http://blog.adamcameron.me/2014/11/something-for-weekend-wee-code-quiz-in.html. The description of the challenge is to create code to take an input array of integers and a threshold value, and from that return the longest consecutive sequence of numbers which, when totalled, do not exceed the threshold value. In his example the following array of numbers:

```javascript
[100, 300, 100, 50, 50, 50, 50, 50, 500, 200, 100]
```

Given a threshold of **500** the following would be the accurate answer.

```javascript
[100, 50, 50, 50, 50, 50]
```

This code is a version in Go that is a command-line application that takes two input parameters: **input** and **threshold**. To run Adam's specific example (in Windows) execute the following in PowerShell.

```powershell
./adamCameronCodeChallenge201411.exe -input="100,300,100,50,50,50,50,50,500,200,100" -threshold=500
```

### Running Unit Tests
If you have Go installed and configured execute the following from PowerShell in the *adamCameronCodeChallenge201411* directory.

```powershell
go test ./... -v
```

### Additional Nerdy Information
* This solution runs the summing of array slices in parallel using goroutines
* This means you *could possibly* change the code to use multiple processors, and pass it an obscene number of integers to process, and it might be pretty fast
* I'm using BDD-style tests
* I used GoConvey, a great BDD library for writing tests in Go

### License
The MIT License (MIT)

Copyright (c) 2014 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

