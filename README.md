# Config file structure

Some fields are optional

```
global:
    ssh_key: ~/.ssh/id_rsa # Will use this by default if nothing else is specified
    ssh_key_password: id_rsa_pwd # Selects a credential from credentials-file.
    ssh_password: password # Password can also be specified instead of ssh
    credentials_file: ~/dbman-credentials # encrypted file
    editor: nvim

    # Tunnels can be specified on global level, or on specific connections.
    tunnels:
      - target: 10.0.1.15
        host: 10.0.1.4
        ssh_key: ~/.ssh/id_rsa
        ssh_key_password: id_rsa_pwd # Selects a credential from credentials-file.
        jumphost: 10.0.1.5 
        jumphost_ssh_key: ~/.ssh/id_rsa
        jumphost_ssh_key_password: id_rsa_pwd # Selects a credential from credentials-file.
        ssh_password: password # Can also specify password instead of SSH
      - target: 10.0.1.16
        cmd: "ssh -i ~/.ssh/id_rsa -L 127.0.0.1:3306:10.0.1.16:3306" # Optionally, the SSH command can simply be specified directly.

connections:
    - name: prod-db
      group: prod
      client: mysqlsh
      host: 10.0.1.15
      port: 3306
      cmd_opts: "--sslmode=require"
      credentials: prod-app-user # Selects a credential from credentials-file.
      read_only: false
      tunnel:
          host: 10.0.1.4
          ssh_key: ~/.ssh/id_rsa
          ssh_key_password: id_rsa_pwd # Selects a credential from credentials-file.
          jumphost: 10.0.1.5 
          jumphost_ssh_key: ~/.ssh/id_rsa
          jumphost_ssh_key_password: id_rsa_pwd # Selects a credential from credentials-file.
          ssh_password: password # Can also specify password instead of SSH
          cmd_opts: something # custom cmd opts can also be specified to the SSH command

    # Connections can be "inherited" to select specific DBs or override certain fields
    - name: prod-db-app
      base: prod-db
      db: myapp 

    - name: prod-db-app-select-items
      base: prod-db
      db: myapp 
      exec: "select * from items;" # A connection can also be made to execute a command

    - name: local-test
      cmd: "mysql -h 0.0.0.0 -u test -p" # It's also possible to simply specify a cli command that this connectionn will run when triggered.
      tunnel: 
```

# Credentials file structure:

```
credentials:
    prod_app:
        username: prod_app
        password: abc123
    id_rsa:
        password: abc123
```

# Functionality

- The tool will print the command it will use to connect before running it, hiding any password ofc.
- Should support at least mysql and postgresql
- All configuration files should contain a thorough documentation at the top, with commented examples of how to use the file.
- Credentials file is stored encrypted on disk.
- Credentials and config are mainly edited with a TUI text editor. The preferred editor might be configurable so any editor can be used.

When starting the program with no args, it should let the user select a connection with fzf:

```
> dbc
Select connection:
  db-dev
> db-prod-app
  db-prod
  3/3 
```

Unless the connection is specified:

```
> dbc db-prod
mysql -h 10.0.1.15 -u user -p******
```

New credentials should be added using a "credentials" command, which opens up the credentials file in the users preferred editor.

```
> dbc credentials
Enter credentials file password: ******
1 credentials:
2     prod_app:
3          username: prod_app
4          password: abc123
-- INSERT --
```

New credentials can also be added through the tool:

```
> dbc credentials add
Enter credentials file password: ******
Adding new credential:
Username: dbuser
Password: ********
Name of credential (default: dbuser): db_user
```

Connections configs can be added by configuring the config file directly, or using a command:

```
> dbc config
1 connections:
2    - name: local-test
3      cmd: "mysql -h 0.0.0.0 -u test -p" # It's also possible to simply specify a cli command that this connectionn will run when triggered.
-- INSERT --
```

```
> dbc config add
Adding new config:
Config name: stage_db
Host: 10.0.1.20
User: stage_user
...
```
