# Magic Eden Engineering Challenge 
Instructions on how to run and test your code in a local environment through the
command line.
○ A description of how and why you chose the design patterns you did
○ A description of what observability you would add if this was a production system.
What would you monitor for a production rollout?


### Instructions for running locally
After unziping the project run the below commands from inside the project directory
```zsh
make postgres

make createdb

make run
```

To reset and recreate the database run:

```zsh
make freshdb
make run
```

### Design Outline and Motivations

 
