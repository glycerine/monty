package shadow_driver

import "database/sql/driver"

var Pkg = make(map[string]interface{})
var Ctor = make(map[string]interface{})

func init() {
    Pkg["Bool"] = driver.Bool
    Pkg["ColumnConverter"] = Shadow_InterfaceConvertTo2_ColumnConverter
    Pkg["Conn"] = Shadow_InterfaceConvertTo2_Conn
    Pkg["ConnBeginTx"] = Shadow_InterfaceConvertTo2_ConnBeginTx
    Pkg["ConnPrepareContext"] = Shadow_InterfaceConvertTo2_ConnPrepareContext
    Pkg["Connector"] = Shadow_InterfaceConvertTo2_Connector
    Pkg["DefaultParameterConverter"] = driver.DefaultParameterConverter
    Pkg["Driver"] = Shadow_InterfaceConvertTo2_Driver
    Pkg["DriverContext"] = Shadow_InterfaceConvertTo2_DriverContext
    Pkg["ErrBadConn"] = driver.ErrBadConn
    Pkg["ErrRemoveArgument"] = driver.ErrRemoveArgument
    Pkg["ErrSkip"] = driver.ErrSkip
    Pkg["Execer"] = Shadow_InterfaceConvertTo2_Execer
    Pkg["ExecerContext"] = Shadow_InterfaceConvertTo2_ExecerContext
    Pkg["Int32"] = driver.Int32
    Pkg["IsScanValue"] = driver.IsScanValue
    Pkg["IsValue"] = driver.IsValue
    Ctor["NamedValue"] = Shadow_NewStruct_NamedValue
    Pkg["NamedValueChecker"] = Shadow_InterfaceConvertTo2_NamedValueChecker
    Ctor["NotNull"] = Shadow_NewStruct_NotNull
    Ctor["Null"] = Shadow_NewStruct_Null
    Pkg["Pinger"] = Shadow_InterfaceConvertTo2_Pinger
    Pkg["Queryer"] = Shadow_InterfaceConvertTo2_Queryer
    Pkg["QueryerContext"] = Shadow_InterfaceConvertTo2_QueryerContext
    Pkg["Result"] = Shadow_InterfaceConvertTo2_Result
    Pkg["ResultNoRows"] = driver.ResultNoRows
    Pkg["Rows"] = Shadow_InterfaceConvertTo2_Rows
    Pkg["RowsColumnTypeDatabaseTypeName"] = Shadow_InterfaceConvertTo2_RowsColumnTypeDatabaseTypeName
    Pkg["RowsColumnTypeLength"] = Shadow_InterfaceConvertTo2_RowsColumnTypeLength
    Pkg["RowsColumnTypeNullable"] = Shadow_InterfaceConvertTo2_RowsColumnTypeNullable
    Pkg["RowsColumnTypePrecisionScale"] = Shadow_InterfaceConvertTo2_RowsColumnTypePrecisionScale
    Pkg["RowsColumnTypeScanType"] = Shadow_InterfaceConvertTo2_RowsColumnTypeScanType
    Pkg["RowsNextResultSet"] = Shadow_InterfaceConvertTo2_RowsNextResultSet
    Pkg["SessionResetter"] = Shadow_InterfaceConvertTo2_SessionResetter
    Pkg["Stmt"] = Shadow_InterfaceConvertTo2_Stmt
    Pkg["StmtExecContext"] = Shadow_InterfaceConvertTo2_StmtExecContext
    Pkg["StmtQueryContext"] = Shadow_InterfaceConvertTo2_StmtQueryContext
    Pkg["String"] = driver.String
    Pkg["Tx"] = Shadow_InterfaceConvertTo2_Tx
    Ctor["TxOptions"] = Shadow_NewStruct_TxOptions
    Pkg["Value"] = Shadow_InterfaceConvertTo2_Value
    Pkg["ValueConverter"] = Shadow_InterfaceConvertTo2_ValueConverter
    Pkg["Valuer"] = Shadow_InterfaceConvertTo2_Valuer

}
func Shadow_InterfaceConvertTo2_ColumnConverter(x interface{}) (y driver.ColumnConverter, b bool) {
	y, b = x.(driver.ColumnConverter)
	return
}

func Shadow_InterfaceConvertTo1_ColumnConverter(x interface{}) driver.ColumnConverter {
	return x.(driver.ColumnConverter)
}


func Shadow_InterfaceConvertTo2_Conn(x interface{}) (y driver.Conn, b bool) {
	y, b = x.(driver.Conn)
	return
}

func Shadow_InterfaceConvertTo1_Conn(x interface{}) driver.Conn {
	return x.(driver.Conn)
}


func Shadow_InterfaceConvertTo2_ConnBeginTx(x interface{}) (y driver.ConnBeginTx, b bool) {
	y, b = x.(driver.ConnBeginTx)
	return
}

func Shadow_InterfaceConvertTo1_ConnBeginTx(x interface{}) driver.ConnBeginTx {
	return x.(driver.ConnBeginTx)
}


func Shadow_InterfaceConvertTo2_ConnPrepareContext(x interface{}) (y driver.ConnPrepareContext, b bool) {
	y, b = x.(driver.ConnPrepareContext)
	return
}

func Shadow_InterfaceConvertTo1_ConnPrepareContext(x interface{}) driver.ConnPrepareContext {
	return x.(driver.ConnPrepareContext)
}


func Shadow_InterfaceConvertTo2_Connector(x interface{}) (y driver.Connector, b bool) {
	y, b = x.(driver.Connector)
	return
}

func Shadow_InterfaceConvertTo1_Connector(x interface{}) driver.Connector {
	return x.(driver.Connector)
}


func Shadow_InterfaceConvertTo2_Driver(x interface{}) (y driver.Driver, b bool) {
	y, b = x.(driver.Driver)
	return
}

func Shadow_InterfaceConvertTo1_Driver(x interface{}) driver.Driver {
	return x.(driver.Driver)
}


func Shadow_InterfaceConvertTo2_DriverContext(x interface{}) (y driver.DriverContext, b bool) {
	y, b = x.(driver.DriverContext)
	return
}

func Shadow_InterfaceConvertTo1_DriverContext(x interface{}) driver.DriverContext {
	return x.(driver.DriverContext)
}


func Shadow_InterfaceConvertTo2_Execer(x interface{}) (y driver.Execer, b bool) {
	y, b = x.(driver.Execer)
	return
}

func Shadow_InterfaceConvertTo1_Execer(x interface{}) driver.Execer {
	return x.(driver.Execer)
}


func Shadow_InterfaceConvertTo2_ExecerContext(x interface{}) (y driver.ExecerContext, b bool) {
	y, b = x.(driver.ExecerContext)
	return
}

func Shadow_InterfaceConvertTo1_ExecerContext(x interface{}) driver.ExecerContext {
	return x.(driver.ExecerContext)
}


func Shadow_NewStruct_NamedValue(src *driver.NamedValue) *driver.NamedValue {
    if src == nil {
	   return &driver.NamedValue{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_NamedValueChecker(x interface{}) (y driver.NamedValueChecker, b bool) {
	y, b = x.(driver.NamedValueChecker)
	return
}

func Shadow_InterfaceConvertTo1_NamedValueChecker(x interface{}) driver.NamedValueChecker {
	return x.(driver.NamedValueChecker)
}


func Shadow_NewStruct_NotNull(src *driver.NotNull) *driver.NotNull {
    if src == nil {
	   return &driver.NotNull{}
    }
    a := *src
    return &a
}


func Shadow_NewStruct_Null(src *driver.Null) *driver.Null {
    if src == nil {
	   return &driver.Null{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Pinger(x interface{}) (y driver.Pinger, b bool) {
	y, b = x.(driver.Pinger)
	return
}

func Shadow_InterfaceConvertTo1_Pinger(x interface{}) driver.Pinger {
	return x.(driver.Pinger)
}


func Shadow_InterfaceConvertTo2_Queryer(x interface{}) (y driver.Queryer, b bool) {
	y, b = x.(driver.Queryer)
	return
}

func Shadow_InterfaceConvertTo1_Queryer(x interface{}) driver.Queryer {
	return x.(driver.Queryer)
}


func Shadow_InterfaceConvertTo2_QueryerContext(x interface{}) (y driver.QueryerContext, b bool) {
	y, b = x.(driver.QueryerContext)
	return
}

func Shadow_InterfaceConvertTo1_QueryerContext(x interface{}) driver.QueryerContext {
	return x.(driver.QueryerContext)
}


func Shadow_InterfaceConvertTo2_Result(x interface{}) (y driver.Result, b bool) {
	y, b = x.(driver.Result)
	return
}

func Shadow_InterfaceConvertTo1_Result(x interface{}) driver.Result {
	return x.(driver.Result)
}


func Shadow_InterfaceConvertTo2_Rows(x interface{}) (y driver.Rows, b bool) {
	y, b = x.(driver.Rows)
	return
}

func Shadow_InterfaceConvertTo1_Rows(x interface{}) driver.Rows {
	return x.(driver.Rows)
}


func Shadow_InterfaceConvertTo2_RowsColumnTypeDatabaseTypeName(x interface{}) (y driver.RowsColumnTypeDatabaseTypeName, b bool) {
	y, b = x.(driver.RowsColumnTypeDatabaseTypeName)
	return
}

func Shadow_InterfaceConvertTo1_RowsColumnTypeDatabaseTypeName(x interface{}) driver.RowsColumnTypeDatabaseTypeName {
	return x.(driver.RowsColumnTypeDatabaseTypeName)
}


func Shadow_InterfaceConvertTo2_RowsColumnTypeLength(x interface{}) (y driver.RowsColumnTypeLength, b bool) {
	y, b = x.(driver.RowsColumnTypeLength)
	return
}

func Shadow_InterfaceConvertTo1_RowsColumnTypeLength(x interface{}) driver.RowsColumnTypeLength {
	return x.(driver.RowsColumnTypeLength)
}


func Shadow_InterfaceConvertTo2_RowsColumnTypeNullable(x interface{}) (y driver.RowsColumnTypeNullable, b bool) {
	y, b = x.(driver.RowsColumnTypeNullable)
	return
}

func Shadow_InterfaceConvertTo1_RowsColumnTypeNullable(x interface{}) driver.RowsColumnTypeNullable {
	return x.(driver.RowsColumnTypeNullable)
}


func Shadow_InterfaceConvertTo2_RowsColumnTypePrecisionScale(x interface{}) (y driver.RowsColumnTypePrecisionScale, b bool) {
	y, b = x.(driver.RowsColumnTypePrecisionScale)
	return
}

func Shadow_InterfaceConvertTo1_RowsColumnTypePrecisionScale(x interface{}) driver.RowsColumnTypePrecisionScale {
	return x.(driver.RowsColumnTypePrecisionScale)
}


func Shadow_InterfaceConvertTo2_RowsColumnTypeScanType(x interface{}) (y driver.RowsColumnTypeScanType, b bool) {
	y, b = x.(driver.RowsColumnTypeScanType)
	return
}

func Shadow_InterfaceConvertTo1_RowsColumnTypeScanType(x interface{}) driver.RowsColumnTypeScanType {
	return x.(driver.RowsColumnTypeScanType)
}


func Shadow_InterfaceConvertTo2_RowsNextResultSet(x interface{}) (y driver.RowsNextResultSet, b bool) {
	y, b = x.(driver.RowsNextResultSet)
	return
}

func Shadow_InterfaceConvertTo1_RowsNextResultSet(x interface{}) driver.RowsNextResultSet {
	return x.(driver.RowsNextResultSet)
}


func Shadow_InterfaceConvertTo2_SessionResetter(x interface{}) (y driver.SessionResetter, b bool) {
	y, b = x.(driver.SessionResetter)
	return
}

func Shadow_InterfaceConvertTo1_SessionResetter(x interface{}) driver.SessionResetter {
	return x.(driver.SessionResetter)
}


func Shadow_InterfaceConvertTo2_Stmt(x interface{}) (y driver.Stmt, b bool) {
	y, b = x.(driver.Stmt)
	return
}

func Shadow_InterfaceConvertTo1_Stmt(x interface{}) driver.Stmt {
	return x.(driver.Stmt)
}


func Shadow_InterfaceConvertTo2_StmtExecContext(x interface{}) (y driver.StmtExecContext, b bool) {
	y, b = x.(driver.StmtExecContext)
	return
}

func Shadow_InterfaceConvertTo1_StmtExecContext(x interface{}) driver.StmtExecContext {
	return x.(driver.StmtExecContext)
}


func Shadow_InterfaceConvertTo2_StmtQueryContext(x interface{}) (y driver.StmtQueryContext, b bool) {
	y, b = x.(driver.StmtQueryContext)
	return
}

func Shadow_InterfaceConvertTo1_StmtQueryContext(x interface{}) driver.StmtQueryContext {
	return x.(driver.StmtQueryContext)
}


func Shadow_InterfaceConvertTo2_Tx(x interface{}) (y driver.Tx, b bool) {
	y, b = x.(driver.Tx)
	return
}

func Shadow_InterfaceConvertTo1_Tx(x interface{}) driver.Tx {
	return x.(driver.Tx)
}


func Shadow_NewStruct_TxOptions(src *driver.TxOptions) *driver.TxOptions {
    if src == nil {
	   return &driver.TxOptions{}
    }
    a := *src
    return &a
}


func Shadow_InterfaceConvertTo2_Value(x interface{}) (y driver.Value, b bool) {
	y, b = x.(driver.Value)
	return
}

func Shadow_InterfaceConvertTo1_Value(x interface{}) driver.Value {
	return x.(driver.Value)
}


func Shadow_InterfaceConvertTo2_ValueConverter(x interface{}) (y driver.ValueConverter, b bool) {
	y, b = x.(driver.ValueConverter)
	return
}

func Shadow_InterfaceConvertTo1_ValueConverter(x interface{}) driver.ValueConverter {
	return x.(driver.ValueConverter)
}


func Shadow_InterfaceConvertTo2_Valuer(x interface{}) (y driver.Valuer, b bool) {
	y, b = x.(driver.Valuer)
	return
}

func Shadow_InterfaceConvertTo1_Valuer(x interface{}) driver.Valuer {
	return x.(driver.Valuer)
}

