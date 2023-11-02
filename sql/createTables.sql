CREATE TABLE IF NOT EXISTS expenses(
    id INTEGER PRIMARY KEY,
    label TEXT,
    amount REAL,
    frequency TEXT,
    tag TEXT,
    expensedate TEXT,
    submissiondate TEXT,
    userid TEXT
);

CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    username TEXT,
    hashedpassword TEXT,
    email TEXT,
    creationdate TEXT,
    lastlogin TEXT
)