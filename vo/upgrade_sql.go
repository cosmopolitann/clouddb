package vo

const (
	DBversion = "db-version"
)

var Version = "3"

var UpgradeSql = map[int][]string{
	1: {
		`INSERT OR REPLACE INTO sys_user(id, peer_id, name, phone, sex, nickname, img) VALUES ('416418922095452160', 'QmdeUHUGuydkN2r7fV1Efq9EKjXdPuss2G1fwabRMcQCjr', '小龙客服', '', 0, '小龙客服', '')`,
	},
	2: {
		`INSERT OR REPLACE INTO sys_user(id, peer_id, name, phone, sex, nickname, img) VALUES ('416418922095452160', 'QmUwp5Qaeb7Chyorzm6jffmSE65mske7DsxhHwvxuyQfSe', '小龙客服', '', 0, '小龙客服', '')`,
	},
}
