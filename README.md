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

This isn't just a pointless veil over the porcelain `git` commands (I'm _very_ comfortable with git). When used without the relevant subcommand flags, each subcommand will display a _filtered_ and _mutated_ list of branches and will enable a user to select a branch to interact with by specifying a unique identifying branch number, which for me is quicker than typing out a full branch name (especially as I dynamically generate branch names with both a username prefix and date combination).

I had this code already written in my `.bashrc` but realized that updating my esoteric bash scripting logic wasn't going to be maintainable in the long term. Although the golang equivalent is more code it's ultimately easier to work with than the more _complex_ bash script equivalent.

> Note: I then create alias' in `.bashrc` to this binary and the various subcommands exposed.
