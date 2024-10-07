package domain

type Account struct {
    ID        string `json:"_id,omitempty"`
    Rev       string `json:"_rev,omitempty"`
    UserID    string `json:"user_id"`
    Balance   float64 `json:"balance"`
}
