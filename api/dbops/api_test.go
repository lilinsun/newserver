package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetuser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("lls", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("lls")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("lls", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetuser(t *testing.T) {
	pwd, err := GetUserCredential("lls")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Delete user test failed!")
	}
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}
