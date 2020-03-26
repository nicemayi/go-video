package session

import (
	"go-video/api/dbops"
	"go-video/api/defs"
	"go-video/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map // 线程安全的map golang 1.9以后支持,读性能很好，写还是有问题

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 24*60*60*1000 // Service side session valid time: 24 hours
	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}

	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
