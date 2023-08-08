# Engineering Challenge

## Challenge

Build a production-level real-time system that emulates the indexing of data on the blockchain. In
this system, Solana account updates will stream continuously in real time.

For this challenge, the data will be similar to Solana accounts, but no blockchain knowledge
needed!

## SystemDetails

Create classes as you see fit, each having appropriate encapsulation, attributes, and well defined
interfaces.

In a production system, this would be connected to a blockchain node, and stream data updates
continuously. For this challenge, account updates can be asynchronously read from a json file.

```json
[{
"id": "GzbXUY1JQwRVUf3j3myg2NbDRwD5i4jD4HJpYhVNfiDm",
"accountType": "escrow",
"tokens": 500000 ,
"callbackTimeMs": 400 ,
"data": {
"subtype_field1": true,
"subtype_field2": 999
},
"version": 123
}
]
```

Each account comes into the system at a continuous uniform (random) distribution between 0
and 1000ms.

Each account has the following information:

**ID** - Unique identifier of the account
**AccountType** - Type of the account.
**Data** - Data of the account. All accounts that share the same Account Type have the same data
schema. This is the information in which clients are most interested in. You can assume these
schemas are fixed.
**Tokens** - Amount of tokens in the account.
**Version** - Version of the account on chain. If two updates for the same account come in,the old
version should be erased.
**CallbackTimeMs** - Time at which we’d like to print the contents of the account to console after it’s
been ingested.

Display a short message log message to console when each (accountId+version) tuple has been
indexed.

Display a callback log when an account’s call_back_time_ms has expired. If the same account is
ingested with a newer version number, and the old callback has not fired yet, cancel the older
version’s active callback. If an old version of the same account is ingested, ignore that update.

Display a message when an old callback is canceled in favor of a new one.

Once all events and callbacks have completed, print the highest token-value accounts by
AccountType(taking into account write version), and gracefully shut-down the system.

## ExampleScenarios

These scenarios only cover a single accountID, but demonstrate expected ingestion/callback behaviors:

**Scenario 1 - SingleUpdate**
0ms - simulation starts - ID1 scheduled to be ingested 550ms (0-1000ms random) later
550ms - ID1 v1 is “ingested”, we print it as indexed
950ms - ID1 v1 callbackfires(and we log with version 1)

**Scenario 2 - Updates with Cancellation**
0ms - simulation starts - ID1 scheduled to be ingested 550ms (0-1000ms random) later
550ms - ID1 v1 is “ingested”, we print it asindexed
650ms - ID1 v3 is “ingested”, print ID1 v3 indexed, cancel active ID1 v1 callback
950ms - ID1 callback fires (and we log with version 1)
1050ms - ID1 v3 callback fires


## Deliverables

Please take your time delivering a quality solution that shows your ability.Include:
* A README file that contains:
    * Instructions on how to run and test your code in a local environment through the commandline.
    * A description of how and why you chose the design patterns you did
    * A description of what observability you would add if this was a production system. What would you monitor for a production rollout?

* Production ready code that:
    * Follows community standard syntax and style
    * Has no debug logging,TODOs,or FIXMEs
    * Has test coverage to ensure quality and safety
* All code zipped into a folder named me-challenege-YOUR-NAME.zip


Please make sure that you’ve showed us your best code before mimicking production-like infrastructure.
While docker is allowed, we discourage spending too much time on setting up real databases/other infrastructure.

## Rubric

This challenge isi meant to help us see your best code,and to showcase your judgment.When
we evaluate the challenge, we look at how focused you were on meeting the requirements,at
the simplicity and correctness of your architecture, at your useofappropriatedesignpatterns,
yourchoicesofthreadinganddatastructures,andyouruseofbestpracticesforcode,testing,
anddocumentation.



