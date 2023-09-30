# notes
notes app

## setting up the database
- Log into mysql as the root user by entering the command `mysql -u root -p` into the terminal.
- Enter your mysql root password.
- Create the notes database in the mysql cli by entering the command `CREATE DATABASE notes;`.
- Grant all privileges on the notes database to your user by entering the command `GRANT ALL PRIVILEGES ON notes.* TO '{user}'@'localhost';`, where `{user}` is replaced by your mysql username.
- Exit mysql by entering `exit;`.

## setting up the environment
- Create a file named ".env" in the root directory of this application.
- Type the following into the .env file where `{user}` is the mysql username with privileges to the notes database and `{pass}` is that mysql user's password:
```
DB_USER = "{user}"
DB_PASS = "{pass}"
```