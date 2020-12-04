#!/bin/bash
http post https://rpc.mainnet.near.org  jsonrpc=2.0 id=dontcare method=block  params:='{"finality":"final"}'
