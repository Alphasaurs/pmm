# MySQL MyISAM Aria Metrics Dashboard

The MySQL MyISAM Aria Metrics dashboard describes the specific features
of MariaDB MySQL server: [Aria storage engine](https://mariadb.com/kb/en/the-mariadb-library/aria-storage-engine/), [Online DDL (online alter table)](https://mariadb.com/kb/en/the-mariadb-library/alter-table/),

> and [InnoDB defragmentation patch](https://mariadb.com/kb/en/the-mariadb-library/defragmenting-innodb-tablespaces/). This dashboard contains the following metrics:

## Aria Storage Engine

Aria storage is specific for MariaDB Server. Aria has most of the same
variables that MyISAM has, but with an Aria prefix. If you use Aria
instead of MyISAM, then you should make `key_buffer_size` smaller and
aria-pagecache-buffer-size bigger.

## Aria Pagecache Reads/Writes

This graph is similar to InnoDB buffer pool reads and
writes. `aria-pagecache-buffer-size` is the main cache for aria storage
engine. If you see high reads and writes (physical IO), i.e. reads is close to
read requests or writes are close to write requests you may need to increase the
`aria-pagecache-buffer-size` (you may need to decrease other buffers:
`key_buffer_size`, `innodb_buffer_pool_size` etc)

## Aria Pagecache Blocks

This graphs shows the utilization for the aria pagecache.  This is similar to
InnoDB buffer pool graph. If you see all blocks are used you may consider
increasing `aria-pagecache-buffer-size` (you may need to decrease other
buffers: `key_buffer_size`, `innodb_buffer_pool_size` etc)

## Aria Transactions Log Syncs

This metric is similar to InnoDB log file syncs. If you see lots of log syncs
and want to relax the durability settings you can change (in seconds) from 30
(default) to a higher number. It is good to look at the disk IO dashboard as
well.

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
