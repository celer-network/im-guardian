## Celer Inter-chain Message (IM) App Guardian

[Celer IM](https://im-docs.celer.network/) app guardian is run by the dApp community to ensure their application security.

The app guardian monitors the [message `Executed` events](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/message/messagebus/MessageBusReceiver.sol#L27-L34) emitted from the `MessageBus` contracts on the destination chains, and uses the `srcChainId` and `srcTxHash` fields in the event to look for the matched [`Message` events](https://github.com/celer-network/sgn-v2-contracts/blob/main/contracts/message/messagebus/MessageBusSender.sol#L15) from the `MessageBus` contracts on the source chains. If it fails to find a matched event on the source chain, it will try to [pause the message receiver (dApp) contracts](https://github.com/celer-network/im-guardian/blob/main/guardian/message.go#L35-L39) or execute any dApp-specific logics if added.

To run the guardian, go to `cmd/main` and do `go run main.go start --home <path>`. 

A `config.toml` file needs to be provided in the `<home>/config` folder. See this [config.toml](https://github.com/celer-network/im-guardian/blob/main/test/config/config.toml) for example.

Note that Celer IM is already secured by the Celer State Guardian Network (SGN), which is a proven secure decentralized platform that has processed a [large volume](https://cbridge-analytics.celer.network/) of cross-chain asset transfers and tons of cross-chain messages without any security incident. This app guardian is for dApp communities who do not fully trust Celer SGN and want further safety guarantees even if Celer IM is compromised.