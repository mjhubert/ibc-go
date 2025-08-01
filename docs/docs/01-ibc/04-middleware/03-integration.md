---
title: Integrating IBC middleware into a chain
sidebar_label: Integrating IBC middleware into a chain
sidebar_position: 4
slug: /ibc/middleware/integration
---


# Integrating IBC middleware into a chain

Learn how to integrate IBC middleware(s) with a base application to your chain. The following document only applies for Cosmos SDK chains.

If the middleware is maintaining its own state and/or processing SDK messages, then it should create and register its SDK module with the module manager in `app.go`.

All middleware must be connected to the IBC router and wrap over an underlying base IBC application. An IBC application may be wrapped by many layers of middleware, only the top layer middleware should be hooked to the IBC router, with all underlying middlewares and application getting wrapped by it.

The order of middleware **matters**, function calls from IBC to the application travel from top-level middleware to the bottom middleware and then to the application. Function calls from the application to IBC goes through the bottom middleware in order to the top middleware and then to core IBC handlers. Thus the same set of middleware put in different orders may produce different effects.

## Example integration

```go
// app.go pseudocode



// middleware 1 and middleware 3 are stateful middleware, 
// perhaps implementing separate sdk.Msg and Handlers
// NOTE: NewKeeper returns a pointer so that we can modify
// the keepers later after initialization
// They are all initialized to use the channelKeeper directly at the start
mw1Keeper := mw1.NewKeeper(storeKey1, ..., channelKeeper) // in stack 1 & 3
// middleware 2 is stateless
mw3Keeper1 := mw3.NewKeeper(storeKey3,..., channelKeeper) //  in stack 1
mw3Keeper2 := mw3.NewKeeper(storeKey3,..., channelKeeper) //  in stack 2

// Only create App Module **once** and register in app module
// if the module maintains independent state and/or processes sdk.Msgs
app.moduleManager = module.NewManager(
  ...
  mw1.NewAppModule(mw1Keeper),
  mw3.NewAppModule(mw3Keeper1),
  mw3.NewAppModule(mw3Keeper2),
  transfer.NewAppModule(transferKeeper),
  custom.NewAppModule(customKeeper)
)

// NOTE: IBC Modules may be initialized any number of times provided they use a separate
// Keeper and underlying port.

customKeeper1 := custom.NewKeeper(..., KeeperCustom1, ...)
customKeeper2 := custom.NewKeeper(..., KeeperCustom2, ...)

// initialize base IBC applications
// if you want to create two different stacks with the same base application,
// they must be given different Keepers and assigned different ports.
transferIBCModule := transfer.NewIBCModule(transferKeeper)
customIBCModule1 := custom.NewIBCModule(customKeeper1, "portCustom1")
customIBCModule2 := custom.NewIBCModule(customKeeper2, "portCustom2")

// create IBC stacks by combining middleware with base application
// IBC Stack builders are initialized with the IBC ChannelKeeper which is the top-level ICS4Wrapper
// NOTE: since middleware2 is stateless it does not require a Keeper
// stack 1 contains mw1 -> mw3 -> transfer
stack1 := porttypes.NewStackBuilder(ibcChannelKeeper).
  Base(transferIBCModule).
  Next(mw3).
  Next(mw1).
  Build()
// stack 2 contains mw3 -> mw2 -> custom1
stack2 := porttypes.NewStackBuilder(ibcChannelKeeper).
  Base(customIBCModule1).
  Next(mw2).
  Next(mw3).
  Build()
// stack 3 contains mw2 -> mw1 -> custom2
stack3 := porttypes.NewStackBuilder(ibcChannelKeeper).
  Base(customIBCModule2).
  Next(mw1).
  Next(mw2).
  Build()

// associate each stack with the moduleName provided by the underlying Keeper
ibcRouter := porttypes.NewRouter()
ibcRouter.AddRoute("transfer", stack1)
ibcRouter.AddRoute("custom1", stack2)
ibcRouter.AddRoute("custom2", stack3)
app.IBCKeeper.SetRouter(ibcRouter)
```
