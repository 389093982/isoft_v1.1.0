package flyway

import (
	"testing"
)

func Test_Flyway(t *testing.T) {
	flyway := new(FlyWay)
	flyway.Dsn = "root:123456@tcp(106.15.186.139:3306)/isoft_flyway"
	flyway.MigrationLocation = `D:\zhourui\program\go\goland_workspace\src\isoft\isoft\common\flyway\migrations\CREATE_DEMO_201901101845.sql`
	flyway.Migrate()
}
