# Wallet
> Generation Wallet Address

> v1.0.0

# BTC

## POST New Segwit

POST /btc/new_segwit

Generate a new BTC segwit address form default mnemonic

> Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» Address|string|true|none||none|
|»» Private|string|true|none||none|
|»» Public|string|true|none||none|
|» trace_id|string|true|none||none|
|» stack|null|true|none||none|

> Response Examples

> OK

```json
{
  "code": 200,
  "status": "success",
  "message": "RequestSuccess",
  "data": {
    "Address": "2N5oc2o58kHsuvoAvMY1bsvFSXQsd6EvMHi",
    "Private": "cTm3VtnuuuVtQm4FdSG2XxEd4ajU41BEFuVrGxxdJXJnsRPbmw3Y",
    "Public": "031b8f7f6e5778b6aa74b429d3bcc50f5e52aa5ba6f7f0319a050104e69332c02b"
  },
  "trace_id": "",
  "stack": null
}
```

## POST SegwitFromMnemonic

POST /btc/segwit_from_mnemonic

Generate BTC segwit address from post mnemonic & password

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» chain_id|body|integer| yes |Defined blockchain 1.Mainnet 2.Testnet 3.Regtest|
|» mnemonic|body|string| yes |Mnemonic wordlist|
|» pass|body|string¦null| no |Password|
|» account|body|integer¦null| no |Account Number|
|» external|body|boolean| yes |If external chain|
|» address|body|integer| no |Address Number|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» Address|string|true|none||none|
|»» Private|string|true|none||none|
|»» Public|string|true|none||none|
|» trace_id|string|true|none||none|
|» stack|null|true|none||none|

> Parameters

```json
{
  "chain_id": 2,
  "mnemonic": "exit fruit duty weekend romance upper human before nuclear rabbit slim frame",
  "external": true,
  "address": 0,
  "account": 0,
  "pass": ""
}
```

> Response Examples

> OK

```json
{
  "code": 200,
  "status": "success",
  "message": "RequestSuccess",
  "data": {
    "Address": "2N34mTbwU6PwyhtGFQy8iML9fq9C3qgCVDE",
    "Private": "cPppReeEVsy9V6TPyXDUjERFMupHqeEcCF1EQJFNrbkjQ62vues9",
    "Public": "039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda"
  },
  "trace_id": "",
  "stack": null
}
```

## POST SegwitFromSeed

POST /btc/segwit_from_seed

Generate BTC segwit address from post seed

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» chain_id|body|integer| yes |Defined blockchain 1.Mainnet 2.Testnet 3.Regtest|
|» seed|body|string| yes |Seed|
|» account|body|integer¦null| no |Account Number|
|» external|body|boolean| yes |If external chain|
|» address|body|integer| no |Address Number|


### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» Address|string|true|none||none|
|»» Private|string|true|none||none|
|»» Public|string|true|none||none|
|» trace_id|string|true|none||none|
|» stack|null|true|none||none|

> Parameters

```json
{
  "chain_id": 0,
  "seed": "string",
  "account": 0,
  "external": true,
  "address": 0
}
```

> Response Examples

> OK

```json
{
  "code": 200,
  "status": "success",
  "message": "RequestSuccess",
  "data": {
    "Address": "2N34mTbwU6PwyhtGFQy8iML9fq9C3qgCVDE",
    "Private": "cPppReeEVsy9V6TPyXDUjERFMupHqeEcCF1EQJFNrbkjQ62vues9",
    "Public": "039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda"
  },
  "trace_id": "",
  "stack": null
}
```

## POST MultiSig

POST /btc/multi_sig

Generate multiple signatures address

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» chain_id|body|integer| yes |none|
|» required|body|integer| yes |none|
|» public_keys|body|[string]| yes |none|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» status|string|true|none||none|
|» message|string|true|none||none|
|» data|object|true|none||none|
|»» Address|string|true|none||none|
|»» Script|string|true|none||none|
|» trace_id|string|true|none||none|
|» stack|null|true|none||none|

> Parameters

```json
{
  "chain_id": 0,
  "required": 0,
  "public_keys": [
    "string"
  ]
}
```

> Response Examples

> OK

```json
{
  "code": 200,
  "status": "success",
  "message": "RequestSuccess",
  "data": {
    "Address": "3EL8R5HYmjaFVDzybAEUJkY3zYHoPDAbpF",
    "Script": "5221039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda2103ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e52ae"
  },
  "trace_id": "",
  "stack": null
}
```