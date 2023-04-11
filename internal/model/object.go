package model

import (
	"WowjoyProject/WADO_URL/global"
)

// 获取影像路径通信
func GetImagePath(objuid string) string {
	global.Logger.Info("***开始获取影像的路径***")
	sql := `select i.file_name,sl.ip,sl.s_virtual_dir
	from instance i
	inner join study_location sl on i.location_code = sl.n_station_code
	where i.sop_instance_uid =?;`
	// global.Logger.Info("***sql:***", sql, objuid)

	row := global.DBEngine.QueryRow(sql, objuid)
	key := FileData{}
	if err := row.Scan(&key.file, &key.ip, &key.virpath); err != nil {
		global.Logger.Error(err)
		return ""
	}
	if key.file.String == "" || key.ip.String == "" || key.virpath.String == "" {
		global.Logger.Error("***影像路径为空***")
		return ""
	}

	return "\\\\" + key.ip.String + "\\" + key.virpath.String + "\\" + key.file.String
}
