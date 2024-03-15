# Today Budget App: Backend

A web application that allows users to have a better understanding of their income and expenses. Each user will be able to add expenses and earnings, being able to categorize and personalize them.
[Check source document.](https://docs.google.com/document/d/1m-8xde6kF-XrUtxaU37rSK2FoNM-H_ge0xn6mkHIjxg)

## Requirements
1.  Compatibility: Browsers (must: Chrome);
2.  Account creation;
3.  Create expected earnings;
4.  Create expected expenses;
5.  Add actual expenses;
6.  Add actual earnings;
7.  See net balance for month;
8.  See past months expenses, earnings, and net balance;
9.  See yearly net expenses, earnings, and balance;
10.  See current remaining budget;
11.  See current month net balance;
12.  See expense/income graph by type (pizza graph);

## Stretch goals:
1.  SSO
2.  Connectivity with card activities (automatically add expenses)
3.  Mobile app

## Technology stack
1.  Frontend: React
2.  Backend: Go (Gin, Gorm)
3.  Database: Postgresql

## Expected routes:
### Accounts:

POST/signup - create an account

### Login:

POST/signin - authenticate user using email and password

### Transactions:

POST/transactions - creates a transaction

GET/transactions - gets all user transactions for period

GET/transactions/{transaction_id} - gets selected transaction

DELETE/transactions/{transaction_id} - deletes the selected transaction

PUT or PATCH/transactions/{transaction_id} - updates transaction
