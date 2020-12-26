# go-gitbranch

Simple command line abstraction for managing git branches.

- `gitbranch create`: create a branch.
- `gitbranch checkout`: checkout a branch.
- `gitbranch delete`: delete a branch.
- `gitbranch rename`: rename a branch.

```
gitbranch create --help

Usage of create:
  -branch string
        branch to create

gitbranch checkout --help

Usage of checkout:
  -branch string
        branch to checkout

gitbranch rename --help

Usage of rename:
  -branch string
        branch to rename
  -name string
        new branch name
  -normalize
        whether to normalize the given branch name
  -prefix
        whether to generate a unique prefix for the branch name

gitbranch delete --help

Usage of delete:
  -branch string
        branch to delete
```

When creating a branch we automatically modify the branch name to have unique prefix: 

```
<username>_<yyyymmdd>_<branch_name_with_hyphens_replaced_by_underscores>
```

This is the same logic used when renaming a branch and passing the `--prefix` and/or `--normalize` flags. 

> Note: the 'username' part can be controlled by setting the environment variable `GITBRANCH_PREFIX`.

## Why?

This isn't just a pointless veil over the porcelain `git` commands (I'm _very_ comfortable with git). 

Firstly, each subcommand (when used without the relevant flags) will display a _filtered_ and _mutated_ list of branches and will enable a user to select a branch to interact with by specifying a unique identifying branch number. This (for me) is quicker than typing out a full branch name (especially as I dynamically generate branch names with both a username prefix and date combination).

You might be wondering why I even need to type out a full branch name manually when there is git autocomplete? Well, nearly every command I use I put inside a bash shell `alias` which means the git autocomplete no longer works! So, as I dynamically generate a rather long git branch, typing that out manually is _tedious_ to say the least.

I had this code already written in my `.bashrc` but realized that updating my esoteric bash scripting logic (the more complicated scripting was done to handle the filtering/numerical prefixing) wasn't going to be maintainable in the long term. It worked, but if I needed to make any changes then it was going to be much harder to remember how it all worked.

On top of that bash scripting doesn't lend itself to code reuse like a more robust 'general programming' language. This means, although the golang equivalent is more code it's ultimately easier to work with than the more _complex_ bash script equivalent.

> Note: I then create alias' in `.bashrc` to this binary and the various subcommands exposed. Refer to my [dotfiles](https://github.com/Integralist/dotfiles) for examples.

### The trouble with git alias

The issue with git related alias' losing autocomplete is possible to workaround, but not without some caveats.

Here's the code necessary to get autocomplete to work:

```bash
alias g="git"
__git_complete g _git
alias gb="git branch"
__git_complete gb _git_branch
alias gc="git checkout"
__git_complete gc _git_checkout
alias gu="git push"
__git_complete gp _git_push
alias gd="git pull"
__git_complete gp _git_pull
```

But the `__git_complete` command is lazy loaded internally by the bash shell, so calling it from within shell configuration files (e.g. `.bash_profile` and `.bashrc`) will cause an error as the command isn't yet loaded. It's not possible to access the function either because it's an _internal_ (i.e. _private_) function as far as the bash shell is concerned.

To workaround _that_ issue, you'll need to grab the code yourself from the bash source code! See: https://github.com/git/git/blob/master/contrib/completion/git-completion.bash which of course is an ugly manual step to have to take.

Now you've _still_ got to remember the various flags because (using the above alias' as an example) imagine you want to delete a branch. You execute `gb -<Tab>` expecting to see a bunch of flags like `-D` for deleting a specified branch but you in fact get nothing!? You have to use the `--` flag equivalent (e.g. `gb --<Tab>` and then you'll see `--delete` as a possible option).

But there's still problems with the fact that only `--` flag equivalents have tab completion (and this is just an issue with git completion in general and not related to alias'). The biggest issue is that there is no long form flag equivalent to `git checkout -b` if you want to create a branch and check it out in one command you can only use `-b` and so you have to remember it as there's no autocomplete for it.

## RELEASES

A release is generated automatically via the [`goreleaser-action`](https://github.com/goreleaser/goreleaser-action) which is triggered whenever a semver git tag is pushed (e.g. `git tag v1.0.0 -m "v1.0.0" && git push origin v1.0.0`).
