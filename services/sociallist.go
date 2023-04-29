package services

func GetSocialList(appCtx AppCtx, owner AccountId, listType string) ([]AccountId, error) {
	log := appCtx.Log("GetSocialList")
	log.Printf("Received: %v", owner)

	stmt := `select id, entity_id from "lists" where "owner_id" = $1 and "list_type" = $2`
	rows, err := appCtx.db.Query(stmt, owner.String(), listType)
	if err != nil {
		return []AccountId{}, err
	}

	entities := make([]AccountId, 0)
	defer rows.Close()
	for rows.Next() {
		var id, entity_id string
		rows.Scan(&id, &entity_id)
		entities = append(entities, NewAccountId(entity_id))
	}

	return entities, nil
}

func AddToSocialList(appCtx AppCtx, owner AccountId, listType string, idToAdd AccountId) error {
	log := appCtx.Log("AddToSocialList")
	log.Printf("Received: %v/%v/%v", owner, listType, idToAdd)

	stmt := `insert into "lists"("id", "list_type","owner_id","entity_id") values(gen_random_uuid(),$1, $2,$3)`
	row := appCtx.db.QueryRow(stmt, listType, owner.String(), idToAdd.String())
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func RemoveFromSocialList(appCtx AppCtx, owner AccountId, listType string, idToRemove AccountId) error {
	log := appCtx.Log("RemoveFromSocialList")
	log.Printf("Received: %v/%v/%v", owner, listType, idToRemove)

	stmt := `delete from "lists" where "owner_id" = ? and "list_type" = ? and "entity_id" = ?)`
	result, err := appCtx.db.Exec(stmt, owner.String(), listType, idToRemove.String())
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
