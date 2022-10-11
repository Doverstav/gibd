# GIBD (git interactive branch delete)
A small CLI that will allow a user to comfortably delete several git branches.

## Commands
Usecases:
Delete local branches => DONE
Delete local branches with no remote => DONE
Delete local branches with no remote & prune remote beforehand => DONE

There should be a "force" flag to remove branches even if there is a warning (for example, they haven't been merged) => DONE
If the force flag is not set and branch cannot be delete, user should be able to interactively select which branch to force delete => DONE

Do not include main/master => DONE

### Extensions to functionality
Allow user to specify default branch => DONE
Allow user to include main/master => DONE
Allow user to prune before running delete => DONE

### Git commands used
`git for-each-ref --format '%(refname) %(upstream:track)' refs/heads` => Output each branch ref and [gone] if the remote is gone (should language be set to ensure [gone] is shown?)

`git branch -d <branch>` => Delete branch, uses `-D` when doing force delete

`git symbolic-ref refs/remotes/origin/HEAD` => Tries to find default branch so it cna be exluded from list of branches. Does not always work. It proably makes more sense to ask the remote, but that takes more time.

`git prune remote origin` => Used when pruning the remote before running remote-gone

## Improvements
- [X] Pretty error output
- [X] Do not include default branch
    - [X] Allow user to specify default branch & remote
    - [X] Allow user to include main/master
- [X] Allow user to prune before running delete
- [X] Print something if there are no branches to delete
- [ ] Fix docs
- [ ] Clean up/keep things DRY
- [X] Release with Goreleaser?
    - [ ] ~~Or is it enough to just publish the package?~~