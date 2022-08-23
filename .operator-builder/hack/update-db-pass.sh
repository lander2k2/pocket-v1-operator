#!/bin/bash

set -e

VALIDATOR=$1
MANIFEST=/tmp/validator.yaml
PWD=$(kubectl get secret validator.${VALIDATOR}-database.credentials.postgresql.acid.zalan.do -ojson | jq -r '.data.password' | base64 -d)

if [[ "$VALIDATOR" == "v1-validator1" ]]; then
    PRIVATE_KEY="2e00000000000000000000000000000000000000000000000000000000000000264a0707979e0d6691f74b055429b5f318d39c2883bb509310b67424252e9ef2"
elif [[ "$VALIDATOR" == "v1-validator2" ]]; then
    PRIVATE_KEY="2d00000000000000000000000000000000000000000000000000000000000000ee37d8c8e9cf42a34cfa75ff1141e2bc0ff2f37483f064dce47cb4d5e69db1d4"
elif [[ "$VALIDATOR" == "v1-validator3" ]]; then
    PRIVATE_KEY="2b000000000000000000000000000000000000000000000000000000000000001ba66c6751506850ae0787244c69476b6d45fb857a914a5a0445a24253f7b810"
elif [[ "$VALIDATOR" == "v1-validator4" ]]; then
    PRIVATE_KEY="2c00000000000000000000000000000000000000000000000000000000000000f868bcc508133899cc47b612e4f7d9d5dacc90ce1f28214a97b651baa00bf6e4"
fi

LAST_APPLIED=$(kubectl get cm ${VALIDATOR}-config -ojson | jq -r '.metadata.annotations."kubectl.kubernetes.io/last-applied-configuration"')

cat > $MANIFEST <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: $VALIDATOR-config
  namespace: primary
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      $LAST_APPLIED
data:
  config.json: |
    {
      "root_dir": "/data",
      "genesis_source": {
        "file": {
          "path": "/genesis.json"
        }
      },
      "private_key": "$PRIVATE_KEY",
      "enable_telemetry": true,
      "p2p": {
        "consensus_port": 8221,
        "use_raintree": true,
        "connection_type": "tcp",
        "protocol": "tcp",
        "address": "0.0.0.0:8222",
        "peers": [
          "v1-validator1:8222",
          "v1-validator2:8222",
          "v1-validator3:8222",
          "v1-validator4:8222"
        ]
      },
      "consensus": {
        "max_mempool_bytes": 500000000,
        "max_block_bytes": 4000000,
        "pacemaker": {
          "timeout_msec": 5000,
          "manual": true,
          "debug_time_between_steps_msec": 1000
        }
      },
      "pre_persistence": {
        "capacity": 99999,
        "mempool_max_bytes": 99999,
        "mempool_max_txs": 99999
      },
      "persistence": {
        "postgres_url": "postgres://validator:$PWD@$VALIDATOR-database:5432/validatordb",
        "schema": "validator",
        "block_store_path": "/blockstore"
      },
      "utility": {},
      "telemetry": {
        "address": "0.0.0.0:9000",
        "endpoint": "/metrics"
      }
    }
EOF

kubectl apply -f $MANIFEST

exit 0

