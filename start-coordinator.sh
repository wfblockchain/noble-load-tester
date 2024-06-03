#!/bin/sh

/workdir/noble-load-tester/build/noble-load-tester coordinator --expect-workers ${NUM_OF_WORKERS} --bind 0.0.0.0:26670 -c 1 -T ${TIME} -r ${RATE} -s 250 --broadcast-tx-method async \
    --endpoints ${CHAIN_ENDPOINT} --client-factory noble