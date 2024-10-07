package domain

type User struct {
    ID    string `json:"_id,omitempty"`
    Rev   string `json:"_rev,omitempty"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
