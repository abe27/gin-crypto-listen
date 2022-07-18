package services

var HelloApi = "Hello, Welcome to Crypto API!"
var SystemErrorMessage = "เกิดข้อผิดพลาดกรุณาติดต่อผู้ดูแลระบบเพื่อทำการแก้ไข"
var RegisterSuccessMessage = "ลงทะเบียนเรียบร้อยแล้ว"
var CheckInputRequiredMessage = "กรุณาตรวจสอบสอบข้อมูลก่อนทำการบันทึกด้วย"
var DataIsDuplicateMessage = "ข้อมูลนี้ถูกลงทะเบียนไปแล้ว"
var SigInSuccessMessage = "เข้าสู่ระบบเรียบร้อยแล้ว"
var PasswordIsNotMatchMessage = "ระบุรหัสผ่านไม่ถูกต้อง"
var NotFoundUserMessage = "ไม่พบข้อมูลผู้ใช้งาน"
var AuthenticateRequiredTokenMessage = "กรุณาระบุ Token ของท่านด้วย"
var UserLeaveMessage = "ออกจากระบบเรียบร้อยแล้ว"
var RequiredAuthenticationMessage = "กรุณาระบุ Authorization ด้วย"
var NotFoundTokenMessage = "ไม่พบข้อมูล Authorization Token!"
var TokenExpiredMessage = "Token is expire!"

var ShowAllDataMessage = func(name string) string {
	return "แสดงข้อมูล " + name + " ทั้งหมด"
}

var FoundDataMessage = func(name string) string {
	return "แสดงข้อมูล " + name + ""
}

var NotFoundDataMessage = func(name string) string {
	return "ไม่พบข้อมูล " + name + "!"
}

var CreateDataSuccessMessage = func(name string) string {
	return "บันทึกข้อมูล " + name + " เรียบร้อยแล้ว"
}

var UpdateDataMessage = func(name string) string {
	return "อัพเดทข้อมูล " + name + " เรียบร้อยแล้ว"
}

var DeleteDataMessage = func(name string) string {
	return "ลบข้อมูล " + name + " เรียบร้อยแล้ว"
}
