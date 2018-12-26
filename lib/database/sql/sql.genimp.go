package shadow_sql

import "database/sql"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Ctor["ColumnType"] = Shadow_NewStruct_ColumnType
    Ctor["Conn"] = Shadow_NewStruct_Conn
    Ctor["DB"] = Shadow_NewStruct_DB
    Ctor["DBStats"] = Shadow_NewStruct_DBStats
    Pkg["Drivers"] = sql.Drivers
    Pkg["ErrConnDone"] = sql.ErrConnDone
    Pkg["ErrNoRows"] = sql.ErrNoRows
    Pkg["ErrTxDone"] = sql.ErrTxDone
    Pkg["Named"] = sql.Named
    Ctor["NamedArg"] = Shadow_NewStruct_NamedArg
    Ctor["NullBool"] = Shadow_NewStruct_NullBool
    Ctor["NullFloat64"] = Shadow_NewStruct_NullFloat64
    Ctor["NullInt64"] = Shadow_NewStruct_NullInt64
    Ctor["NullString"] = Shadow_NewStruct_NullString
    Pkg["Open"] = sql.Open
    Pkg["OpenDB"] = sql.OpenDB
    Ctor["Out"] = Shadow_NewStruct_Out
    Pkg["Register"] = sql.Register
    Pkg["Result"] = Shadow_InterfaceConvertTo2_Result
    Ctor["Row"] = Shadow_NewStruct_Row
    Ctor["Rows"] = Shadow_NewStruct_Rows
    Pkg["Scanner"] = Shadow_InterfaceConvertTo2_Scanner
    Ctor["Stmt"] = Shadow_NewStruct_Stmt
    Ctor["Tx"] = Shadow_NewStruct_Tx
    Ctor["TxOptions"] = Shadow_NewStruct_TxOptions

}
func Shadow_NewStruct_ColumnType(src *sql.ColumnType) *sql.ColumnType {
    if src == nil {
	   return &sql.ColumnType{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Conn(src *sql.Conn) *sql.Conn {
    if src == nil {
	   return &sql.Conn{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_DB(src *sql.DB) *sql.DB {
    if src == nil {
	   return &sql.DB{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_DBStats(src *sql.DBStats) *sql.DBStats {
    if src == nil {
	   return &sql.DBStats{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NamedArg(src *sql.NamedArg) *sql.NamedArg {
    if src == nil {
	   return &sql.NamedArg{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NullBool(src *sql.NullBool) *sql.NullBool {
    if src == nil {
	   return &sql.NullBool{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NullFloat64(src *sql.NullFloat64) *sql.NullFloat64 {
    if src == nil {
	   return &sql.NullFloat64{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NullInt64(src *sql.NullInt64) *sql.NullInt64 {
    if src == nil {
	   return &sql.NullInt64{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_NullString(src *sql.NullString) *sql.NullString {
    if src == nil {
	   return &sql.NullString{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Out(src *sql.Out) *sql.Out {
    if src == nil {
	   return &sql.Out{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Result(x interface{}) (y sql.Result, b bool) {
	y, b = x.(sql.Result)
	return
}

func Shadow_InterfaceConvertTo1_Result(x interface{}) sql.Result {
	return x.(sql.Result)
}


func Shadow_NewStruct_Row(src *sql.Row) *sql.Row {
    if src == nil {
	   return &sql.Row{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Rows(src *sql.Rows) *sql.Rows {
    if src == nil {
	   return &sql.Rows{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Scanner(x interface{}) (y sql.Scanner, b bool) {
	y, b = x.(sql.Scanner)
	return
}

func Shadow_InterfaceConvertTo1_Scanner(x interface{}) sql.Scanner {
	return x.(sql.Scanner)
}


func Shadow_NewStruct_Stmt(src *sql.Stmt) *sql.Stmt {
    if src == nil {
	   return &sql.Stmt{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Tx(src *sql.Tx) *sql.Tx {
    if src == nil {
	   return &sql.Tx{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_TxOptions(src *sql.TxOptions) *sql.TxOptions {
    if src == nil {
	   return &sql.TxOptions{}
    }
    a := *src
    return &a
}

