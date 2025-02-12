package main

import (
	"bytes"
	"time"
)

// ボーレート
const BAUDRATE int = 9600

// シリアルポート名 ラズパイ4の場合、"/dev/serial0"
const SERIAL_PORT_NAME string = "/dev/serial0"

// 高速回転時の自動停止しきい値
const IMU_TOOFAST_THRESHOULD float64 = 35.0

// ボールセンサー故障検知しきい値。L: OPEN故障、H: CLOSE故障
const BALLSENS_LBREAK_THRESHOULD int = 100
const BALLSENS_HBREAK_THRESHOULD int = 100

// ボールセンサーの反応しきい値。
const BALLSENS_LOW_THRESHOULD int = 100

// バッテリーの低下しきい値。 150 = 15.0V
const BATTERY_LOW_THRESHOULD int = 150
const BATTERY_CRITICAL_THRESHOULD int = 145

var sendarray bytes.Buffer //送信用バッファ

// 受信時の構造体
type RecvStruct struct {
	Volt        uint8
	PhotoSensor uint16
	IsHoldBall  bool
	ImuDir      int16
}

// 送信時の構造体
type SendStruct struct {
	preamble     byte
	motor        [4]uint8
	dribblePower uint8
	kickPower    uint8
	chipPower    uint8
	imuDir       uint8
	imuFlg       uint8
	emg          bool
}

// 受信データ構造体
var recvdata RecvStruct

// imu角度
var imudegree int16

// imu速度超過時のフラグ
var imuError bool = false

var last_recv_time time.Time = time.Now()

// ポート8080番で待ち受ける。
const PORT string = ":9191"

var isRobotError = false

var RobotErrorCode = 0
var RobotErrorMessage = ""

var ballSensLowCount = 0

var doBuzzer = false
var buzzerTone = 0
var buzzerTime time.Duration = 0 * time.Millisecond

var alarmIgnore = false

var imuReset bool = false

var kicker_enable bool = false //キッカーの入力のON OFFを定義する
var kicker_val uint8 = 0       //キッカーの値
var chip_enable bool = false   //チップキックの入力のON OFFを定義する
var chip_val uint8 = 0         //チップキックの値

// IMU Resetを確実に行うためのフラグ
// 待機モードにうつり、これがセットされているときはAIから受け取らない
var imuResetPending bool = false
