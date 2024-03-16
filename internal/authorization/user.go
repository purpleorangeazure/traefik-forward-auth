package authorization

// User represents an autorized user
type User struct {
	Name   string
	Groups []string
	Roles []string
}

// GetName returns the user name
func (k *User) GetName() string {
	return k.Name
}

// GetGroups return list of groups the user belongs to
func (k *User) GetGroups() []string {
	return k.Groups
}

func (k *User) GetRoles() []string {
	return k.Roles
}