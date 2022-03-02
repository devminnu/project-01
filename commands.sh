# initialize empty git repo
git init

# set default branch name on every new repo initialization with git init command
 git config --global init.defaultBranch BRANCH_NAME
 git config --global init.defaultBranch main

#  change/rename branch name
git branch -m BRANCH_NAME
git branch -m main

# git status
git status

# to stage files
git add FILE_1 FILE_2

# to unstage files
git rm --cached FILE_1 FILE_2

# git restore - restores changes to a file
git restore FILE_NAME
