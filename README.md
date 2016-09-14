# Thrombosis of the Grainmill

Thrombomis of the Grainmill is a little command line tool that outputs a string based on a user-defined template. The template uses Go's text/template package, and currently allows iteration over passed arguments. In so doing, it makes me feel much happier than regular old `alias` ever could, hence the name.

## How to Use It

In your home directory, make a folder called `.thrombosis-of-the-grainmill`. It in, put files that contain your alias templates, like so:

```
$ cat ~/.thrombosis-of-the-grainmill/git-ancestor
COMMON=`comun{{range .}} <(git rev-list {{.}}){{end}}`
if [ $? -eq 0 ]; then
       	git show-branch {{range .}}{{.}} {{end}} $COMMON
else
       	echo "No common ancestor found"
fi
```

The `range .` is iterating over the arguments, which is the only variable passed into the template.

## Note

To actually execute what you output from this, you might want to pipe it through bash:

`thrombosis-of-the-grainmill git-ancestor | /bin/bash`

You could even simplify this, like I have with a lil wrapper bash script:

```
$ cat ~/bin/git-anc
#!/bin/bash
thrombosis-of-the-grainmill git-ancestor $@ | /bin/bash;
```

## So what?

No what.