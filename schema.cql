CREATE KEYSPACE test13004 WITH replication = {'class': 'NetworkTopologyStrategy', 'datacenter1': '3'}  AND durable_writes = true;

CREATE TYPE test13004.guild_role (
    id bigint,
    name text,
    changed_permissions bigint,
    permissions bigint,
    position int,
    color int,
    hoist boolean,
    tags frozen<set<text>>,
    mentionable boolean
);

CREATE TABLE test13004.guilds (
    id bigint PRIMARY KEY,
    name text,
    roles map<bigint, frozen<guild_role>>,
) WITH bloom_filter_fp_chance = 0.01
    AND caching = {'keys': 'ALL', 'rows_per_partition': 'NONE'}
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy', 'max_threshold': '32', 'min_threshold': '4'}
    AND compression = {'chunk_length_in_kb': '64', 'class': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND crc_check_chance = 1.0
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 0.0
    AND speculative_retry = '99PERCENTILE';
