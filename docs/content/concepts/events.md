---
title: Core Events
---

# Introduction

This document describes events emitted from core contracts and FVM. 

# Querying Events

To query events you will need to connect to a Flow access node using the SDK for the language that you are using.

## Go SDK

Documentation for how to handle events can be found here:

[https://github.com/onflow/flow-go-sdk#querying-events](https://github.com/onflow/flow-go-sdk#querying-events)

And sample code for doing so is here:

[https://github.com/onflow/flow-go-sdk/tree/master/examples/query_events](https://github.com/onflow/flow-go-sdk/tree/master/examples/query_events)

## JavaScript SDK

You can find the JavaScript SDK for Flow here:

[https://github.com/onflow/flow-js-sdk](https://github.com/onflow/flow-js-sdk)


## Flow CLI

You can use Flow CLI to query events, learn more about it here:

[https://docs.onflow.org/flow-cli/get-events/](https://docs.onflow.org/flow-cli/get-events/)

# Event Descriptions

## Native Events
Native events are events emitted directly from the FVM (Flow virtual machine). The events have the same name on 
all networks and don't follow the standard naming.

### Account Created

Event that is emitted when a new account gets created.

Event name: `flow.AccountCreated`


```cadence
pub event AccountCreated(address: Address)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the newly created account |


### Account Key Added

Event that is emitted when a key gets added to an account.

Event name: `flow.AccountKeyAdded`

```cadence
pub event AccountKeyAdded(address: Address)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the account the key is added to |
| publicKey       | [UInt8] | Public key added to an account |


### Account Key Removed

Event that is emitted when a key gets removed from an account.

Event name: `flow.AccountKeyRemoved`

```cadence
pub event AccountKeyRemoved(address: Address)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the account the key is removed from |
| publicKey       | [UInt8] | Public key removed from an account |


### Account Contract Added

Event that is emitted when a contract gets deployed to an account.

Event name: `flow.AccountContractAdded`

```cadence
pub event AccountContractAdded(address: Address, codeHash: [UInt8], contract: String)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the account the contract gets deployed to |
| codeHash       | [UInt8] | Hash of the contract source code |
| contract       | String | The name of the the contract |

### Account Contract Updated

Event that is emitted when a contract gets updated on an account.

Event name: `flow.AccountContractUpdated`

```cadence
pub event AccountContractUpdated(address: Address, codeHash: [UInt8], contract: String)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the account the contract gets updated on |
| codeHash       | [UInt8] | Hash of the contract source code |
| contract       | String | The name of the the contract |


### Account Contract Removed

Event that is emitted when a contract gets removed from an account.

Event name: `flow.AccountContractRemoved`

```cadence
pub event AccountContractRemoved(address: Address, codeHash: [UInt8], contract: String)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| address       | Address | The address of the account the contract gets removed from |
| codeHash       | [UInt8] | Hash of the contract source code |
| contract       | String | The name of the the contract |



## Core Contract Events

Flow relies on a set of core contracts that define key portions of the Flow protocol. Those contracts are core contracts 
and are made to emit the events documented bellow. You can read about the core [contracts here](https://docs.onflow.org/core-contracts) 
and view their source code with events definitions.

Events emitted from core contracts follow the standard format:
```
A.{contract address}.{contract name}.{event name}
```
The components of the event name:
- `contract address` - the address of the account the contract has been deployed to
- `contract name` - the name of the contract in the source code
- `event name` - the name of the event found in the source code event definition

### Flow Token Contract
Description of events emitted from the [Flow Token contract](https://docs.onflow.org/core-contracts/flow-token/). 
The contract defines the fungible Flow token. Please note that events for the fungible token contracts are the same 
if deployed to a different account but the `contract address` is 
changed to the address of the account the contract has been deployed to.

### Tokens Initialized

Event that is emitted when the contract gets created.

- Event name: `TokensInitialized`
- Mainnet event: `A.1654653399040a61.FlowToken.TokensInitialized`
- Testnet event: `A.7e60df042a9c0868.FlowToken.TokensInitialized`

```cadence
pub event TokensInitialized(initialSupply: UFix64)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| initialSupply       | UFix64 | The initial supply of the tokens |


### Tokens Withdrawn

Event that is emitted when tokens get withdrawn from a Vault.

- Event name: `TokensWithdrawn`
- Mainnet event: `A.1654653399040a61.FlowToken.TokensWithdrawn`
- Testnet event: `A.7e60df042a9c0868.FlowToken.TokensWithdrawn`

```cadence
pub event TokensWithdrawn(amount: UFix64, from: Address?)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| amount       | UFix64 | The amount of tokens withdrawn |
| from       | Address? | Optional address of the account to withdraw from |


### Tokens Deposited

Event that is emitted when tokens get deposited to a Vault.

- Event name: `TokensDeposited`
- Mainnet event: `A.1654653399040a61.FlowToken.TokensDeposited`
- Testnet event: `A.7e60df042a9c0868.FlowToken.TokensDeposited`

```cadence
pub event TokensDeposited(amount: UFix64, to: Address?)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| amount       | UFix64 | The amount of tokens withdrawn |
| to       | Address? | Optional address of the account the amount is deposited to |

### Tokens Minted

Event that is emitted when new tokens gets minted.

- Event name: `TokensMinted`
- Mainnet event: `A.1654653399040a61.FlowToken.TokensMinted`
- Testnet event: `A.7e60df042a9c0868.FlowToken.TokensMinted`

```cadence
pub event TokensMinted(amount: UFix64)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| amount       | UFix64 | The amount of tokens to mint |

### Tokens Burned

Event that is emitted when tokens get destroyed.

- Event name: `TokensBurned`
- Mainnet event: `A.1654653399040a61.FlowToken.TokensBurned`
- Testnet event: `A.7e60df042a9c0868.FlowToken.TokensBurned`

```cadence
pub event TokensBurned(amount: UFix64)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| amount       | UFix64 | The amount of tokens to burn |


### Minter Created

Event that is emitted when a new minter resource gets created.

- Event name: `MinterCreated`
- Mainnet event: `A.1654653399040a61.FlowToken.MinterCreated`
- Testnet event: `A.7e60df042a9c0868.FlowToken.MinterCreated`

```cadence
pub event MinterCreated(allowedAmount: UFix64)
```

| Field             | Type   | Description                                                            |
| ----------------- | ------ | ---------------------------------------------------------------------- |
| allowedAmount       | UFix64 | The amount of tokens that the minter is allowed to mint |

### Burner Created

Event that is emitted when a new burner resource gets created.

- Event name: `BurnerCreated`
- Mainnet event: `A.1654653399040a61.FlowToken.BurnerCreated`
- Testnet event: `A.7e60df042a9c0868.FlowToken.BurnerCreated`

```cadence
pub event BurnerCreated()
```

### Staking Events
Read about the staking events here: [https://docs.onflow.org/staking/events/](https://docs.onflow.org/staking/events/)

