package services

import (
	"net"
	"strings"
	"time"
)

type UserContact struct {
	AccountID AccountId
	Ip        net.IP
	Port      int32
	Time      time.Time
}

func UpdateUserContact(appCtx AppCtx, accountId AccountId, ip net.IP, port int32) error {
	log := appCtx.Log("UpdateUserContact")
	log.Printf("Received: %v/%v/%v", accountId, ip, port)

	stmt := `insert into user_contacts (id,ip,port,timestamp) VALUES ($1,$2,$3,now()) ON CONFLICT(id) DO UPDATE SET ip=$2, port=$3, timestamp=now() `
	result, err := appCtx.db.Exec(stmt, accountId.String(), ip.String(), port)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("Removed %d rows\n", rows)

	return nil
}

func GetUserContacts(appCtx AppCtx, accountIDs []AccountId) ([]UserContact, error) {
	log := appCtx.Log("GetUserContacts")
	log.Printf("Received: %v", accountIDs)

	stmt := `select id, ip, port, timestamp from "user_contacts" where id in ANY($1::uuid[])`

	params := make([]string, len(accountIDs))
	for _, accountId := range accountIDs {
		params = append(params, accountId.String())
	}
	param := "{" + strings.Join(params, ",") + "}"

	rows, err := appCtx.db.Query(stmt, param)
	if err != nil {
		return []UserContact{}, err
	}

	contacts := make([]UserContact, 0)
	defer rows.Close()
	for rows.Next() {
		var id string
		var ip net.IP
		var port int32
		var timestamp time.Time

		rows.Scan(&id, &ip, &port, &timestamp)
		contacts = append(contacts, UserContact{
			AccountID: NewAccountId(id),
			Ip:        ip,
			Port:      port,
			Time:      timestamp,
		})
	}

	return contacts, nil

}
