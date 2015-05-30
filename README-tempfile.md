<center>
# tempfile
</center>

## What Is It?

_tempfile_ is a go package that creates files in a controlled way that
allows for easy cleanup (if desired).  

It is forked from the go team's pprof
tempfile package but has some enhancements to make it more efficient when there
are large numbers of files to be created.  The original package makes a bad choice
of initial filename to create and as a result each additional file created takes
longer than the last.  No problem for a dozen files, but if thousands then it's
not so good.  The fix is trivial, at the cost of one additional state variable.

I also added a test program to validate the package works as intended.

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/tempfile```

Will copy github.com/hotei/program to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/tempfile.git```

### Features

* Works with same API as pprof version, just faster

### Limitations

* <font color="red">Limit of 10,000 files in one directory</font> But
thats a constant value that can be changed with a recompile.
 
### Usage

Typical usage is demonstrated in the test program repeated here:

```go
func Test_001(t *testing.T) {
	defer Cleanup()
	for i := 0; i < 5; i++ {
		f, err := New("/home/mdr/tmp", "prefix", "suffix")
		if err != nil {
			fmt.Printf("tempfile create failed\n")
			os.Exit(-1)
		}
		fmt.Printf("Created %s\n", f.Name())
		DeferDelete(f.Name())
	}
}

results in the following output:
```

	Created /home/mdr/tmp/prefix000suffix
	Created /home/mdr/tmp/prefix001suffix
	Created /home/mdr/tmp/prefix002suffix
	Created /home/mdr/tmp/prefix003suffix
	Created /home/mdr/tmp/prefix004suffix
	PASS
	ok  	github.com/hotei/tempfile	0.002s


Calling Cleanup() to remove the files is optional, in which case DeferDelete()
can also be omitted.

### To-Do

* Essential:
 * TBD
* Nice:
 * TBD
* Nice but no immediate need:
 * TBD

### Change Log

* 2015-05-25 Started with go 1.4.2
 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for tempfile] [3]
* [Source for pprof's tempfile] [4]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/tempfile "github.com/hotei/tempfile"
[4]: http://golang.org/pkg/cmd/pprof/internal/tempfile/ "pprof tempfile"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.
License is BSD-two-clause, in file "LICENSE"

License
-------
The 'tempfile' go package/program is distributed under the Simplified BSD License:

> Copyright (c) 2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2015 David Rook 

// EOF README.md  (this is a markdown document and tested OK with blackfriday)
