// user ->api service -> delete video
// api service -> scheduler-> write video deletion record
// timer
// timer -> runnner -> read video deletion record -> delete video from folder
package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddvideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}
	_,err= stmtIns.Exec(vid)
	if err!=nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
