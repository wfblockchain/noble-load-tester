cd proto
buf generate
cd ..

cp -r github.com/wfblockchain/noble-fiattokenfactory/* ./
rm -rf github.com
