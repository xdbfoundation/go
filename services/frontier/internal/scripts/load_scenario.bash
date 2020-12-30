#! /usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SCENARIO=$1
CORE_SQL=$DIR/../test/scenarios/$SCENARIO-core.sql
FRONTIER_SQL=$DIR/../test/scenarios/$SCENARIO-frontier.sql

echo "psql $DIGITALBITS_CORE_DATABASE_URL < $CORE_SQL" 
psql $DIGITALBITS_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $FRONTIER_SQL"
psql $DATABASE_URL < $FRONTIER_SQL 
