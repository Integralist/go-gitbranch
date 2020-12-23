# go-gitbranch

Simple command line abstraction for managing git branches.

- `gitbranch create -name <branch>`: create a branch.
- `gitbranch checkout -name <branch>`: checkout a branch.
- `gitbranch delete -name <branch>`: delete a branch.
- `gitbranch rename -name <branch> -prefix`: rename a branch.

> Note: the 'prefix' used for creating a branch (as well as renaming a branch) can be controlled by setting the environment variable `GITBRANCH_PREFIX`.

This isn't just a pointless veil over the porcelain `git` commands (I'm _very_ comfortable with git). Each command will output a filtered and mutated list of branches and will enable a user to select a branch to interact with by specifying a unique identifying branch number, which for me is quicker than typing out a full branch name (especially as I dynamically generate branch names with both a username prefix and timestamp combination).

I had this code already written in my `.bashrc` but realized that updating my esoteric bash scripting logic wasn't going to be maintainable in the long term. Although the golang equivalent is more code it's ultimately easier for me to work with.

I will typically create alias' in `.bashrc` to this binary and the various subcommands exposed.
