# GIBD (git interactive branch delete)
A small CLI that will allow a user to comfortably delete several git branches.

## Commands
Usecases:
Delete local branches => DONE
Delete local branches with no remote => DONE
Delete local branches with no remote & prune remote beforehand

There should be a "force" flag to remove branches even if there is a warning (for example, they haven't been merged)
If the force flag is not set and branch cannot be delete, user should be able to interactively select which branch to force delete

Do not include main/master

### Extensions to functionality
Allow user to specify default branch
Allow user to include main/master
Allow user to prune before running delete

### Git commands used
`git for-each-ref --format '%(refname) %(upstream:track)' refs/heads` => Output each branch ref and [gone] if the remote is gone (should language be set to ensure [gone] is shown?)

`got branch -d <branch>` => Delete branch, uses `-D` when doing force delete

## Improvements
- [X] Pretty error output
- [X] Do not include default branch
    - [ ] Allow user to specify default branch & remote
    - [ ] Allow user to include main/master
- [ ] Allow user to prune before running delete