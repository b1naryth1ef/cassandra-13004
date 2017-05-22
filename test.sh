cqlsh -e 'DROP KEYSPACE test13004'
cqlsh -f schema.cql
cqlsh -e "COPY test13004.guilds FROM 'data.csv' WITH skiprows=0 AND maxrows=1;"
cqlsh -e "COPY test13004.guilds FROM 'data.csv' WITH skiprows=1 AND maxrows=2;"
