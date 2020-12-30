#!/usr/bin/env bash

set -e

source /etc/profile

echo "using config:"
cat digitalbits-core.cfg

# initialize new db
digitalbits-core new-db

if [ "$1" = "standalone" ]; then
  # start a network from scratch
  digitalbits-core force-scp

  # initialze history archive for stand alone network
  digitalbits-core new-hist vs

  # serve history archives to frontier on port 1570
  pushd /history/vs/
  python3 -m http.server 1570 &
  popd
fi

exec /init -- digitalbits-core run