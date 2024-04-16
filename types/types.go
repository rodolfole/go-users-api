package types

type UserStore interface {
	CheckIfUserExist(correo string, telefono string) (*User, error)
	GetUserByCorreo(correo string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID         int    `json:"id"`
	Usuario    string `json:"usuario"`
	Correo     string `json:"correo"`
	Telefono   string `json:"telefono"`
	Contrasena string `json:"-"`
}

type LoginUserPayload struct {
	Correo     string `json:"correo" validate:"required,email"`
	Contrasena string `json:"contrasena" validate:"required"`
}

type RegisterUserPayload struct {
	Usuario    string `json:"usuario" validate:"required"`
	Correo     string `json:"correo" validate:"required,email"`
	Telefono   string `json:"telefono" validate:"required,min=10,max=10"`
	Contrasena string `json:"contrasena" validate:"required,min=6,max=12"`
}
