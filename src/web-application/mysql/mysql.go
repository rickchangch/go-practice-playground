package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Go 連結資料庫的 driver，類似於 Java JDBC，需實作指定的 interface (database/sql/driver)
*/

// Member ...
type Member struct {
	ID       int
	Name     string
	Info     string
	Birthday time.Time // MySQL min Date: 1000-01-01
	Register time.Time // MySQL min DateTime: 1000-01-01 00:00:00
	Login    time.Time // MySQL min DateTime: 1000-01-01 00:00:00
	VIP      string
	Created  time.Time
	Updated  time.Time
}

func (m *Member) String() string {
	memberBytes, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}

	return string(memberBytes)
}

// Connect ...
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:rootpass@tcp(localhost:3307)/mytest?charset=utf8mb4,utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetMember ...
func GetMember(db *sql.DB, id int64) (*Member, error) {
	mem := &Member{}
	err := db.QueryRow("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?", id).Scan(&mem.ID, &mem.Name, &mem.Info, &mem.Birthday, &mem.Register, &mem.Login, &mem.VIP, &mem.Created, &mem.Updated)

	if err != nil {
		return nil, err
	}
	return mem, nil
}

func Run() {
	fmt.Println("- MySQL Demo")

	db, err := Connect()
	if err != nil {
		log.Fatal("connect:", err)
	}
	defer db.Close()

	birthday := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC) //time.Parse("2006-01-02", "1000-01-01")
	register := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	login := time.Now()

	m1 := &Member{
		Name:     "小明",
		Info:     "小明",
		Birthday: birthday,
		Register: register,
		Login:    login,
		VIP:      "A",
	}
	log.Println("member:", m1)

	ins, err := db.Prepare("insert into member(name, info, birthday, register, login, vip) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("prepare insert:", err)
		return
	}
	defer ins.Close()

	result, err := ins.Exec(m1.Name, m1.Info, m1.Birthday, m1.Register, m1.Login, m1.VIP)
	if err != nil {
		log.Println("insert:", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("last id:", err)
		return
	}

	sel, err := db.Prepare("select id, name, info, birthday, register, login, vip, created, updated from member where id = ?")
	if err != nil {
		log.Println("prepare select:", err)
		return
	}
	defer sel.Close()

	rows, err := sel.Query(id)
	if err != nil {
		log.Println("query:", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		m2 := &Member{}
		err = rows.Scan(&m2.ID, &m2.Name, &m2.Info, &m2.Birthday, &m2.Register, &m2.Login, &m2.VIP, &m2.Created, &m2.Updated)
		if err != nil {
			log.Println("scan:", err)
			return
		}
		fmt.Println("get member:", m2)
	} else {
		log.Printf("data (%d) not found\n", id)
	}

	other, err := GetMember(db, 100)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("id 100 not found")
		} else {
			fmt.Println("other error:", err)
		}
	} else {
		fmt.Println("other member:", other)
	}

	fmt.Println("end")
}
