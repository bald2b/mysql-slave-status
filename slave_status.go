package main

import "database/sql"

type SlaveStatus struct {
	Slave_IO_State                sql.NullString
	Master_Host                   sql.NullString
	Master_User                   sql.NullString
	Master_Port                   sql.NullInt64
	Connect_Retry                 sql.NullInt64
	Master_Log_File               sql.NullString
	Read_Master_Log_Pos           sql.NullInt64
	Relay_Log_File                sql.NullString
	Relay_Log_Pos                 sql.NullInt64
	Relay_Master_Log_File         sql.NullString
	Slave_IO_Running              sql.NullString
	Slave_SQL_Running             sql.NullString
	Replicate_Do_DB               sql.NullString
	Replicate_Ignore_DB           sql.NullString
	Replicate_Do_Table            sql.NullString
	Replicate_Ignore_Table        sql.NullString
	Replicate_Wild_Do_Table       sql.NullString
	Replicate_Wild_Ignore_Table   sql.NullString
	Last_Errno                    sql.NullInt64
	Last_Error                    sql.NullString
	Skip_Counter                  sql.NullInt64
	Exec_Master_Log_Pos           sql.NullInt64
	Relay_Log_Space               sql.NullInt64
	Until_Condition               sql.NullString
	Until_Log_File                sql.NullString
	Until_Log_Pos                 sql.NullInt64
	Master_SSL_Allowed            sql.NullString
	Master_SSL_CA_File            sql.NullString
	Master_SSL_CA_Path            sql.NullString
	Master_SSL_Cert               sql.NullString
	Master_SSL_Cipher             sql.NullString
	Master_SSL_Key                sql.NullString
	Seconds_Behind_Master         sql.NullInt64
	Master_SSL_Verify_Server_Cert sql.NullString
	Last_IO_Errno                 sql.NullInt64
	Last_IO_Error                 sql.NullString
	Last_SQL_Errno                sql.NullInt64
	Last_SQL_Error                sql.NullString
	Replicate_Ignore_Server_Ids   sql.NullString
	Master_Server_Id              sql.NullInt64
	Master_UUID                   sql.NullString
	Master_Info_File              sql.NullString
	SQL_Delay                     sql.NullInt64
	SQL_Remaining_Delay           sql.NullInt64
	Slave_SQL_Running_State       sql.NullString
	Master_Retry_Count            sql.NullInt64
	Master_Bind                   sql.NullString
	Last_IO_Error_Timestamp       sql.NullString
	Last_SQL_Error_Timestamp      sql.NullString
	Master_SSL_Crl                sql.NullString
	Master_SSL_Crlpath            sql.NullString
	Retrieved_Gtid_Set            sql.NullString
	Executed_Gtid_Set             sql.NullString
	Auto_Position                 sql.NullInt64
	Replicate_Rewrite_DB          sql.NullString
	Channel_Name                  sql.NullString
	Master_TLS_Version            sql.NullString
	Master_public_key_path        sql.NullString
	Get_master_public_key         sql.NullInt64
	Network_Namespace             sql.NullString
	Error                         error
}