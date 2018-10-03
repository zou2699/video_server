package dbops

import (
	"testing"
	"strconv"
	"time"
	"fmt"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

var tempvid string

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// user
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Delete", testDeleteUser)
	t.Run("ReGet", testReGetUserCredential)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("Tsou", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("Tsou")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("Tsou", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testReGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("Tsou")
	if err != nil {
		t.Errorf("Error of ReDeleteUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

// video
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddNewVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddNewVideoInfo:%v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo:%v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo:%v", err)
	}
}

// comment
func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "i love golang"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddNewComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
