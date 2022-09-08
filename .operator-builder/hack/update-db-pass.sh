#!/bin/bash

set -e

VALIDATOR=$1
MANIFEST=/tmp/validator.yaml
PWD=$(kubectl get secret validator.${VALIDATOR}-database.credentials.postgresql.acid.zalan.do -ojson | jq -r '.data.password' | base64 -d)

if [[ "$VALIDATOR" == "v1-validator1" ]]; then
    PRIVATE_KEY="ccec19df8fe866280e41da68d52d0ecdb07b01e85eeef45f400fd3a89b71c26a79254a4bc46bf1182826145b0b01b48bab4240cd30e23ba90e4e5e6b56961c6d"
elif [[ "$VALIDATOR" == "v1-validator2" ]]; then
    PRIVATE_KEY="8a76f99e3bf132f1d61bb4a123f495e00c169ac7c55fa1a2aa1b34196020edb191dd4fd53e8e27020d62796fe68b469fad5fa5a7abc61d3eb2bd98ba16af1e29"
elif [[ "$VALIDATOR" == "v1-validator3" ]]; then
    PRIVATE_KEY="c7bd1bd027e76b31534c3f5226c8e3c3f2a034ba9fa11017b65191f7f9ef0d253e5e4bbed5f98721163bb84445072a9202d213f1e348c5e9e0e2ea83bbb7e3aa"
elif [[ "$VALIDATOR" == "v1-validator4" ]]; then
    PRIVATE_KEY="ddff03df6c525e551c5e9cd0e31ac4ec99dd6aa5d62185ba969bbf2e62db7e2c6c207cea1b1bf45dad8f8973d57291d3da31855254d7f1ed83ec3e06cabfe6b7"
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
      "base": {
        "root_directory": "/go/src/github.com/pocket-network",
        "private_key": "$PRIVATE_KEY"
      },
      "consensus": {
        "max_mempool_bytes": 500000000,
        "pacemaker_config": {
          "timeout_msec": 5000,
          "manual": true,
          "debug_time_between_steps_msec": 1000
        }
      },
      "utility": {},
      "persistence": {
        "postgres_url": "postgres://validator:$PWD@$VALIDATOR-database:5432/validatordb",
        "node_schema": "validator",
        "block_store_path": "/blockstore"
      },
      "p2p": {
        "consensus_port": 8080,
        "use_rain_tree": true,
        "connection_type": 1
      },
      "telemetry": {
        "enabled": true,
        "address": "0.0.0.0:9000",
        "endpoint": "/metrics"
      }
    }
EOF

kubectl apply -f $MANIFEST

exit 0

