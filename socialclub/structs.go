package socialclub

import "time"

type CrewHierarchy struct {
	CrewRanks []struct {
		Name            string `json:"name"`
		RankOrder       int    `json:"rankOrder"`
		RankPermissions struct {
			CanWriteOnWall            bool `json:"canWriteOnWall"`
			CanInvite                 bool `json:"canInvite"`
			CanPublishMultipleEmblems bool `json:"canPublishMultipleEmblems"`
			CanDeleteFromWall         bool `json:"canDeleteFromWall"`
			CanDemote                 bool `json:"canDemote"`
			CanPromote                bool `json:"canPromote"`
			CanKick                   bool `json:"canKick"`
			CanViewSettings           bool `json:"canViewSettings"`
			CanEditSettings           bool `json:"canEditSettings"`
			CanPostMessages           bool `json:"canPostMessages"`
			CanUpdateStatus           bool `json:"canUpdateStatus"`
			CanManageRanks            bool `json:"canManageRanks"`
		} `json:"rankPermissions"`
		MemberCount int `json:"memberCount"`
		RankMembers []struct {
			AvatarURL        string    `json:"avatarUrl"`
			Nickname         string    `json:"nickname"`
			RockstarID       int       `json:"rockstarId"`
			DateJoined       time.Time `json:"dateJoined"`
			RankOrder        int       `json:"rankOrder"`
			OnlineService    string    `json:"onlineService"`
			IsGamertagHidden bool      `json:"isGamertagHidden"`
			Gamertag         string    `json:"gamertag"`
			PrimaryClan      struct {
				ID            int    `json:"id"`
				Name          string `json:"name"`
				Tag           string `json:"tag"`
				IsOpenClan    bool   `json:"isOpenClan"`
				IsSystemClan  bool   `json:"isSystemClan"`
				IsPrivateClan bool   `json:"isPrivateClan"`
				IsFounderClan bool   `json:"isFounderClan"`
				Color         string `json:"color"`
				RankOrder     int    `json:"rankOrder"`
			} `json:"primaryClan"`
		} `json:"rankMembers"`
	} `json:"crewRanks"`
	CrewID int  `json:"crewId"`
	Status bool `json:"status"`
}

type CrewBasicInfo struct {
	CrewID               int       `json:"crewId"`
	CrewName             string    `json:"crewName"`
	CrewTag              string    `json:"crewTag"`
	CrewMotto            string    `json:"crewMotto"`
	MemberCount          int       `json:"memberCount"`
	IsPrimary            bool      `json:"isPrimary"`
	IsPrivate            bool      `json:"isPrivate"`
	RankOrder            int       `json:"rankOrder"`
	IsFounderCrew        bool      `json:"isFounderCrew"`
	IsOpen               bool      `json:"isOpen"`
	IsSystem             bool      `json:"isSystem"`
	IsSystemPrivate      bool      `json:"isSystemPrivate"`
	CrewColour           string    `json:"crewColour"`
	Division             string    `json:"division"`
	CanInvite            bool      `json:"canInvite"`
	CanLeave             bool      `json:"canLeave"`
	CanJoin              bool      `json:"canJoin"`
	CanJoinDisabled      bool      `json:"canJoinDisabled"`
	CanRequestInvite     bool      `json:"canRequestInvite"`
	IsJoinRequestPending bool      `json:"isJoinRequestPending"`
	ShowMaxExceededBtn   bool      `json:"showMaxExceededBtn"`
	IsMember             bool      `json:"isMember"`
	CanDeleteFromWall    bool      `json:"canDeleteFromWall"`
	CreatedAt            time.Time `json:"createdAt"`
	Status               bool      `json:"status"`
}
