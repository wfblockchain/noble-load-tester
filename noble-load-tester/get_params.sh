export HOME_DIR="~/noblechain/play_sh/noble-1"
export ALICE_ADDR=$(nobled keys show user --home $HOME_DIR --keyring-backend test -a)
export MINTER_ADDR=$(nobled keys show minter --home $HOME_DIR --keyring-backend test -a)
export ALICE_PRIV=$(yes | nobled keys export user --unarmored-hex --unsafe --home $HOME_DIR --keyring-backend test)
export MINTER_PRIV=$(yes | nobled keys export minter --unarmored-hex --unsafe --home $HOME_DIR --keyring-backend test)

export ALICE_ACC_NUM=$((nobled q account $ALICE_ADDR --home $HOME_DIR --output json) | python3 -c "import sys, json; print(json.load(sys.stdin)['account_number'])")
export ALICE_ACC_SEQ=$((nobled q account $ALICE_ADDR --home $HOME_DIR --output json) | python3 -c "import sys, json; print(json.load(sys.stdin)['sequence'])")
export MINTER_ACC_NUM=$((nobled q account $MINTER_ADDR --home $HOME_DIR --output json) | python3 -c "import sys, json; print(json.load(sys.stdin)['account_number'])")
export MINTER_ACC_SEQ=$((nobled q account $MINTER_ADDR --home $HOME_DIR --output json) | python3 -c "import sys, json; print(json.load(sys.stdin)['sequence'])")


echo "Alice acc num: $ALICE_ACC_NUM"
echo "Alice acc seq: $ALICE_ACC_SEQ"
echo "Minter acc num: $MINTER_ACC_NUM"
echo "Minter acc seq: $MINTER_ACC_SEQ"
