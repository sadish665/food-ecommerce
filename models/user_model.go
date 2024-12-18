package models

type UserModel struct{
	CommonModel
	FirstName string
	LastName string
	Email string
	Address string
	Password string
	Age int
}

func GetAllUsers()[]UserModel{
		c := UserModel{
			FirstName: "abc",
			LastName: "xyz",
			Email: "abaxyz@gmail.com",
			Address: "BKT",
			Age: 20,
		}
		b := UserModel{
			FirstName: "def",
			LastName: "uvw",
			Email: "defuvw@gmail.com",
			Address: "KTM",
			Age: 21,
		}
		a := UserModel{
			FirstName: "ghi",
			LastName: "rst",
			Email: "ghirst@gmail.com",
			Address: "LTP",
			Age: 22,
		}
	return []UserModel{
		a,
		b,
		c,
	}
}