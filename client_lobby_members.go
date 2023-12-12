package dota2

import (
	bgcm "github.com/dotabook/go-dota2/protocol"
	"github.com/dotabook/go-steam/steamid"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Dota2) InviteLobbyMember(playerID steamid.SteamId) {
	steamID := playerID.ToUint64()
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToLobby), &bgcm.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}
