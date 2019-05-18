package main

var (
	secKillConf = &SecSkillConf{}
)

type SecSkillConf struct {
	redisAddr string
	etcdAddr  string
}
