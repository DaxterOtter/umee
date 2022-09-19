package ibctransfer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"

	"github.com/umee-network/umee/v3/x/ibctransfer/keeper"
)

// IBCModule embeds the ICS-20 transfer IBCModule where we only override specific
// methods.
type IBCModule struct {
	// embed the ICS-20 transfer's AppModule
	ibcporttypes.IBCModule

	keeper keeper.Keeper
}

func NewIBCModule(transferModule ibcporttypes.IBCModule, k keeper.Keeper) IBCModule {
	return IBCModule{
		IBCModule: transferModule,
		keeper:    k,
	}
}

// OnRecvPacket delegates the OnRecvPacket call to the embedded ICS-20 transfer
// IBCModule and updates metadata if successful.
func (am IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	ack := am.IBCModule.OnRecvPacket(ctx, packet, relayer)
	if ack.Success() {
		var data ibctransfertypes.FungibleTokenPacketData
		if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err == nil {
			// track metadata
			am.keeper.PostOnRecvPacket(ctx, packet, data)
		}
	}

	return ack
}
