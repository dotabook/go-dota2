package dota2

import (
	bgcm "github.com/dotabook/go-dota2/protocol"
)

// LeaveParty attempts to leave the current party.
func (d *Dota2) LeaveParty() {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLeaveParty), &bgcm.CMsgLeaveParty{})
}

func (d *Dota2) InviteToParty(steamid *uint64) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToParty), &bgcm.CMsgInviteToParty{
		SteamId: steamid,
	})
}

// RespondPartyInvite attempts to respond to a party invite.
func (d *Dota2) RespondPartyInvite(partyId uint64, accept bool) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCPartyInviteResponse), &bgcm.CMsgPartyInviteResponse{
		PartyId: &partyId,
		Accept:  &accept,
	})
}
