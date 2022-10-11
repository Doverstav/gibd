# GIBD (git interactive branch delete)
A small CLI that will allow a user to comfortably delete several git branches.

Inspired by the problem explained in this [stackoverflow question](https://stackoverflow.com/questions/7726949/remove-tracking-branches-no-longer-on-remote), where I found the answers lacking. While they would no doubt work (I have re-used the commands from them in this small CLI), I felt that thay lacked some interactivity. And I also wanted to play around some with Go!

Although this is not a unique idea, a quick search finds that at least [one project](https://github.com/stefanwille/git-branch-delete) has already solved the same issue in a similar way. But this project wasn't about being unique, but about trying some new things (for me) and hopefully learn something!

## Commands
Executables can be found under [releases](https://github.com/Doverstav/gibd/releases).

Run `gibd --help` to get a nice little help message, which will basically repeat what's written here. 

There are two commands:
- `gibd`: List all branches expect the default branch
- `gibd remote-gone`: List all branches with upstream status [gone]
    - By setting flag `-p`/`--prune` the remote will be pruned first, which should ensure upstream status is set to [gone] on all branches where it should be

Both commands also accept some global flags:
- `-d`: Set the name of the default branch, defalt value master (e.g. `gibd -d main`)
- `-r`: Set the name of the remote, default value origin (e.g. `gibd -r origin`)
- `-i`: Set this flag to include the default branch in the list of branches to delete
- `-f`: Set this flag to force delete branches by default 

### Git commands used
This project uses Go to run some Git commands on your local machine, so in the interest of transparency I will detail what commands are used (and what they are used for) in this section.

`git for-each-ref --format '%(refname) %(upstream:track)' refs/heads`  
Outputs each branch ref and the upsteam track information (e.g [gone])

`git branch -d <branch>`  
Delete branch, uses `-D` when doing force delete

`git symbolic-ref refs/remotes/origin/HEAD`  
Tries to find default branch so it can be exluded from list of branches. Does not always work. It probably makes more sense to ask the remote, but that takes more time.

`git prune remote origin`  
Used when pruning the remote before running remote-gone
