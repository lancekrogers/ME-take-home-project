# Magic Eden Engineering Challenge 
Instructions on how to run and test your code in a local environment through the
command line.
○ A description of how and why you chose the design patterns you did
○ A description of what observability you would add if this was a production system.
What would you monitor for a production rollout?


## Overview
This repository contains the codebase for the Magic Eden Engineering Challenge. It includes various packages for handling accounts, callbacks, database operations, and ingestion of data.

### Instructions for running locally
After unziping the project run the below commands from inside the project directory:
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

To rebuild the binary, ensure you have installed go then run ```make build```

### Design Overview and Motivations
```zsh
├── cmd
│   ├── config.yaml
│   └── main.go
├── data
│   └── challenge-input.json
├── pkg
    ├── accounts
    ├── callbacks
    ├── db
    ├── ingestion
    ├── setup.go
    └── utils
```
#### Singleton Pattern

I used a singleton pattern to implement the entry point to the service in main.go. This pattern helps ensure that the main application object is instantiated only once, providing a single point of access to the resources.

#### Package Structure

- accounts: Handles account-related operations
    - Accounts are managed in two tables Accounts and AccountUpdates
        - Accounts contains the latest state of the account
        - AccounUpdates contains all new account updates
- callbacks: Manages callback functionality
    - Utilizes a mutex to provide concurrent access to the timer hashmap
- db: Contains database-related files, including migrations and SQL queries.
    - Used sqlc to generate boilerplate database code from SQL and gomock to mock the database in test
    - I used postgres database.  I have a lot of experience with postgres and postgres based databases.  For the scope of this project I originally started with sqlite, but I switched to postgres due to a lack of sqlc support for certain features I planned to use in sqlite.
- ingestion: Responsible for data ingestion.
    - asyncrously processes account updates
    - Tables are updated following the observer patter by upserting to Accounts while inserting into AccountUpdates
- utils: Includes utility functions like logging and configuration loading.
    - The logging utility allows for mocking of logging and using a custom logger so that I could show milliseconds


This modular design ensures that each part of the application is contained within its own package, promoting separation of concerns and ease of testing.
