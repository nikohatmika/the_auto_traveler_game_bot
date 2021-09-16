package response

import "auto_traveler/bussiness/player_auth"

type playerAuthResponse struct {
	Token string `json:"token"`
}

func FromDomain(domain *player_auth.Domain) (res *playerAuthResponse) {
	if domain != nil {
		res = &playerAuthResponse{
			Token: domain.Token,
		}
	}

	return res
}
