# CLITASKTRACKER

Task Tracker
Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](roadmap.sh).

How to run
Clone the repository and run the following command:

git clone https://github.com/AnosVoldigoad999/CLITASKTRACKER.git
cd CLITASKTRACKER
Run the following command to build, install and run the project:
#
NOTE: YOU NEED TO HAVE GO INSTALLED, IF YOU DON'T HAVE IT INSTALLED GO HERE: https://go.dev/doc/install
#
go build
#
go install
#
ctt --help # To see the list of available commands

# To add a task
ctt add "Buy groceries"

# To update a task
ctt update 1 "Buy groceries and cook dinner"

# To delete a task
ctt delete 1

# To mark a task as in progress/done/todo
ctt mark-in-progress 1
#
ctt mark-done 1
#
ctt mark-todo 1
#
# To list all tasks
ctt list 
#
ctt list done
#
ctt list todo
#
ctt list in-progress