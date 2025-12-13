package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}

func NewCommitOrRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		_ = tx.Rollback()
		return
	}
	_ = tx.Commit()
}
