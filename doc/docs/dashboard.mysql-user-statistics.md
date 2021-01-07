# *MySQL User Statistics*

This dashboard presents various data related to MySQL users.

**NOTE**: This dashboard requires Percona Server for MySQL 5.1+ or MariaDB
10.1/10.2 with XtraDB. Also `userstat` should be enabled, for example with
the `SET GLOBAL userstat=1` statement. See [Configuring MySQL for Best Results](https://www.percona.com/doc/percona-monitoring-and-management/conf-mysql.html)
for further instructions.

Data is displayed for the 5 top users.

Top Users by Connections Created

    The number of times user’s connections connected using SSL to the server.

Top Users by Traffic

    The number of bytes sent to the user’s connections.

Top Users by Rows Fetched/Read

    The number of rows fetched by the user’s connections.

Top Users by Rows Updated

    The number of rows updated by the user’s connections.

Top Users by Busy Time

    The cumulative number of seconds there was activity on connections from the
    user.

Top Users by CPU Time

    The cumulative CPU time elapsed, in seconds, while servicing connections of
    the user.

<!-- -*- mode: rst -*- -->
<!-- Tips (tip) -->
<!-- Abbreviations (abbr) -->
<!-- Docker commands (docker) -->
<!-- Graphical interface elements (gui) -->
<!-- Options and parameters (opt) -->
<!-- pmm-admin commands (pmm-admin) -->
<!-- SQL commands (sql) -->
<!-- PMM Dashboards (dbd) -->
<!-- * Text labels -->
<!-- Special headings (h) -->
<!-- Status labels (status) -->
