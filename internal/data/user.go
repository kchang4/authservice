package data

type user struct {
	email        string
	username     string
	passwordhash string
	fullname     string
	createDate   string
	role         int
}

// Sample users, our current db
var userList = []user{
	{
		email:        "abc@gmail.com",
		username:     "abc12",
		passwordhash: "hashedme1",
		fullname:     "abc def",
		createDate:   "1631600786",
		role:         1,
	},
	{
		email:        "chekme@example.com",
		username:     "checkme34",
		passwordhash: "hashedme2",
		fullname:     "check me",
		createDate:   "1631600837",
		role:         0,
	},
}

// Methods

// Validates a user's passwordhash
func (u *user) ValidatePasswordHash(passwordhash string) bool {
	return u.passwordhash == passwordhash
}

// Functions

// Retrieves a user from the database
func GetUser(email string) (user, bool) {
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
	}

	return user{}, false
}

// Adds a user to the database
func AddUser(email string, username string, passwordhash string, fullname string, role int) bool {
	for _, user := range userList {
		if user.email == email || user.username == username {
			return false
		}
	}

	userList = append(userList, user{
		email:        email,
		passwordhash: passwordhash,
		username:     username,
		fullname:     fullname,
		role:         role,
	})

	return true
}
