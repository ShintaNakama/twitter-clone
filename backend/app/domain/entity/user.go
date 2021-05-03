package entity

type User struct {
	id    string
	name  string
	email string
	image string
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetImage() string {
	return u.image
}

type UserArgs struct {
	ID    string
	Name  string
	Email string
	Image string
}

func NewUser(args *UserArgs) *User {
	return &User{
		id:    args.ID,
		name:  args.Name,
		email: args.Email,
		image: args.Image,
	}
}
