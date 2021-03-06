package model

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"strconv"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	tMessagesFieldNames          = builderx.RawFieldNames(&TMessages{})
	tMessagesRows                = strings.Join(tMessagesFieldNames, ",")
	tMessagesRowsExpectAutoSet   = strings.Join(stringx.Remove(tMessagesFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	tMessagesRowsWithPlaceHolder = strings.Join(stringx.Remove(tMessagesFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	TMessagesModel interface {
		Insert(data TMessages) (sql.Result, error)
		FindOne(id int64) (*TMessages, error)
		FindManyByDst(device string, deviceId int64, limit ...int) ([]TMessages, error)
		Update(data TMessages) error
		Delete(id int64) error
		DeleteManyByGuid(guids []string) error
		DeleteManyByDestination(destination string) error
		DeleteManyByDestinationDeviceId(destination string, deviceId int64) error
		Remove(destination string,id int64) error
		RemoveByGUID(destination, guid string) (*TMessages, error)
		RemoveBySender(destination string, destinationDevice int64, source string, timestamp int64) (*TMessages, error)

	}

	defaultTMessagesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TMessages struct {
		Id                int64     `db:"id"`
		Type              int64     `db:"type"`
		Relay             string    `db:"relay"`
		Timestamp         int64     `db:"timestamp"`
		Source            string    `db:"source"`
		SourceDevice      int64     `db:"source_device"`
		Destination       string    `db:"destination"`
		DestinationDevice int64     `db:"destination_device"`
		Message           string    `db:"message"`
		Content           string    `db:"content"`
		Guid              string    `db:"guid"`
		ServerTimestamp   int64     `db:"server_timestamp"`
		SourceUuid        string    `db:"source_uuid"`
		Ctime             time.Time `db:"ctime"`
	}
)

func NewTMessagesModel(conn sqlx.SqlConn) TMessagesModel {
	return &defaultTMessagesModel{
		conn:  conn,
		table: "`t_messages`",
	}
}

func (m *defaultTMessagesModel) Insert(data TMessages) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tMessagesRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Type, data.Relay, data.Timestamp, data.Source, data.SourceDevice, data.Destination, data.DestinationDevice, data.Message, data.Content, data.Guid, data.ServerTimestamp, data.SourceUuid, data.Ctime)
	return ret, err
}

func (m *defaultTMessagesModel) FindOne(id int64) (*TMessages, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tMessagesRows, m.table)
	var resp TMessages
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTMessagesModel) RemoveByGUID(destination, guid string) (*TMessages, error)  {
	query := fmt.Sprintf("select %s from %s where `guid` = ? AND `destination` = ? ORDER BY id LIMIT 1", tMessagesRows, m.table)
	var resp TMessages
	err := m.conn.QueryRow(&resp, query, guid,destination)
	if err!=nil{
		return nil, err
	}
	err = m.Delete(resp.Id)
	if err!=nil{
		return nil, err
	}

	return &resp,nil

}

func (m *defaultTMessagesModel) RemoveBySender(
	destination string, destinationDevice int64,
	source string, timestamp int64) (*TMessages, error){

	query := fmt.Sprintf("select %s from %s where `destination` = ? AND `destination_device` = ? AND `source` = ? AND `timestamp` = ? ORDER BY id LIMIT 1", tMessagesRows, m.table)
	var resp TMessages
	err := m.conn.QueryRow(&resp, query,destination,destinationDevice,source,timestamp )
	if err!=nil{
		return nil, err
	}
	err = m.Delete(resp.Id)
	if err!=nil{
		return nil, err
	}

	return &resp,nil

}

func (m *defaultTMessagesModel) FindManyByDst(device string, deviceId int64, limit ...int) ([]TMessages, error) {
	//query := fmt.Sprintf("select %s from %s where destination='%s' and destination_device=%d ", tMessagesRows, m.table,device,deviceId)
	query := fmt.Sprintf("select %s from %s where destination='%s' and destination_device=%d order by id ASC  ", tMessagesRows, m.table, device, deviceId)

	if len(limit) > 0 {
		query += "  limit  "
		query += strconv.Itoa(limit[0])
	}

	var resp []TMessages
	logx.Info(query)
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTMessagesModel) Update(data TMessages) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tMessagesRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Type, data.Relay, data.Timestamp, data.Source, data.SourceDevice, data.Destination, data.DestinationDevice, data.Message, data.Content, data.Guid, data.ServerTimestamp, data.SourceUuid, data.Ctime, data.Id)
	return err
}

func (m *defaultTMessagesModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultTMessagesModel) DeleteManyByGuid(guids []string) error {
	guidList := strings.Join(guids, ",")
	guidList = "(" + guidList + ")"
	query := fmt.Sprintf("delete from %s where `id` in  ?", m.table)
	_, err := m.conn.Exec(query, guidList)
	return err
}

func (m *defaultTMessagesModel)  DeleteManyByDestination(destination string) error{
	query := fmt.Sprintf("delete from %s where `destination`= ?", m.table)
	_, err := m.conn.Exec(query, destination)
	return err
}


func (m *defaultTMessagesModel) DeleteManyByDestinationDeviceId(destination string, deviceId int64) error{
	query := fmt.Sprintf("delete from %s where `destination`= ? and `destination_device`=?", m.table)
	_, err := m.conn.Exec(query, destination,deviceId)
	return err
}


func (m *defaultTMessagesModel) Remove(destination string,id int64) error {
	query := fmt.Sprintf("delete from %s where  `id`=?  and `destination`= ? ", m.table)
	_, err := m.conn.Exec(query, id,destination)
	return err
}
