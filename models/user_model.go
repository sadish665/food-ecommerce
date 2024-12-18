package models

type UserModel struct{
	CommonModel
	UserDto
}

type UserDto struct{
	UserName string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Address string `json:"address"`
	Password string `json:"password"`
	Age int `json:"age"`
	CommonModelDto
}

func GetAllUsers()[]UserModel{
	c := UserModel{
		UserDto: UserDto{
			FirstName: "abc",
			LastName: "xyz",
			Email: "abaxyz@gmail.com",
			Address: "BKT",
			Age: 20,
		},
	}	
	return []UserModel{
		c,
	}
}