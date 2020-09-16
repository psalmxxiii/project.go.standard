package user

import (
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/pkg/zerolog"
)

func Delete(userUuid string) bool {
	Db := pg.ConnectNew()
	dell := `DELETE FROM public.ad_phone WHERE logi_uuid=$1`
	_, err := Db.Exec(dell, userUuid)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"delete.user.go",
			17,
			"api.crud.user.singleton.com.br",
			"Delete",
			err.Error())
		return false
	}
	return true
}
