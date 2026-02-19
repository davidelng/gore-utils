# copy

`copy` is a high-performance byte-shoveler, a fancy title to describe a clone of the `cp` command.

## Considerations

While Go has an inbuilt `io.Copy` and `bytes.Buffer`, we will stick with a simple `[]byte` slice as buffer to perform the copy, we don't need abstractions yet, maybe in the future when the "educational" part is over.

What I want the program to be able to do:

- copy files, obviously
- copy entire directories, and recursive should be the default option, since I find myself using `cp -r` often, but a no-recursive mode should be present
- perform a dry run without executing the copy
- have a verbose mode with percentage of completion, a real progress bar would be neat in the future
- prompt the user for confirmation, basically an interactive mode
- have a user-defined buf size, at least at the beginning, I will probabily remove the feature for reasons
- maybe use concurrency, even if this is I\O
- copy from a url to a local file
- copy from/to a server using ssh (scp like)

Lastly, the program and the package should be tested, both unit tests and blackbox tests should be present.

