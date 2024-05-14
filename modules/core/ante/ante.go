package ante

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/cosmos/ibc-go/v8/modules/core/keeper"
)

type RedundantRelayDecorator struct {
	k *keeper.Keeper
}

func NewRedundantRelayDecorator(k *keeper.Keeper) RedundantRelayDecorator {
	return RedundantRelayDecorator{k: k}
}

// AnteHandle returns an error if a multiMsg tx only contains packet messages (Recv, Ack, Timeout) and additional update messages
// and all packet messages are redundant. If the transaction is just a single UpdateClient message, or the multimsg transaction
// contains some other message type, then the antedecorator returns no error and continues processing to ensure these transactions
// are included. This will ensure that relayers do not waste fees on multiMsg transactions when another relayer has already submitted
// all packets, by rejecting the tx at the mempool layer.
func (rrd RedundantRelayDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// do not run redundancy check on DeliverTx or simulate
	if (ctx.IsCheckTx() || ctx.IsReCheckTx()) && !simulate {
		// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`
		redundancies := 0
		packetMsgs := 0
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *channeltypes.MsgRecvPacket:
				response, err := rrd.k.RecvPacket(ctx, msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies++
				}
				packetMsgs++

			case *channeltypes.MsgAcknowledgement:
				response, err := rrd.k.Acknowledgement(ctx, msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies++
				}
				packetMsgs++

			case *channeltypes.MsgTimeout:
				response, err := rrd.k.Timeout(ctx, msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies++
				}
				packetMsgs++

			case *channeltypes.MsgTimeoutOnClose:
				response, err := rrd.k.TimeoutOnClose(ctx, msg)
				if err != nil {
					return ctx, err
				}
				if response.Result == channeltypes.NOOP {
					redundancies++
				}
				packetMsgs++

			case *clienttypes.MsgUpdateClient:
				if err := rrd.updateClientCheckTx(ctx, msg); err != nil {
					return ctx, err
				}

			default:
				// if the multiMsg tx has a msg that is not a packet msg or update msg, then we will not return error
				// regardless of if all packet messages are redundant. This ensures that non-packet messages get processed
				// even if they get batched with redundant packet messages.
				return next(ctx, tx, simulate)
			}
		}

		// only return error if all packet messages are redundant
		if redundancies == packetMsgs && packetMsgs > 0 {
			return ctx, channeltypes.ErrRedundantTx
		}
	}
	return next(ctx, tx, simulate)
}

// updateClientCheckTx runs a subset of ibc client update logic to be used specifically within the RedundantRelayDecorator AnteHandler.
// The following function performs ibc client message verification for CheckTx only and state updates in both CheckTx and ReCheckTx.
// Note that misbehaviour checks are omitted.
func (rrd RedundantRelayDecorator) updateClientCheckTx(ctx sdk.Context, msg *clienttypes.MsgUpdateClient) error {
	clientMsg, err := clienttypes.UnpackClientMessage(msg.ClientMessage)
	if err != nil {
		return err
	}

	if status := rrd.k.ClientKeeper.GetClientStatus(ctx, msg.ClientId); status != exported.Active {
		return errorsmod.Wrapf(clienttypes.ErrClientNotActive, "cannot update client (%s) with status %s", msg.ClientId, status)
	}

	clientModule, found := rrd.k.ClientKeeper.Route(msg.ClientId)
	if !found {
		return errorsmod.Wrap(clienttypes.ErrRouteNotFound, msg.ClientId)
	}

	if !ctx.IsReCheckTx() {
		if err := clientModule.VerifyClientMessage(ctx, msg.ClientId, clientMsg); err != nil {
			return err
		}
	}

	heights := clientModule.UpdateState(ctx, msg.ClientId, clientMsg)

	ctx.Logger().With("module", "x/"+exported.ModuleName).Debug("ante ibc client update", "consensusHeights", heights)

	return nil
}
