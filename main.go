package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// キッカーパワーが入力された時に、値を固定する関数
// 並列での処理が行われる
func kickCheck(chkicker chan bool) {
	for {
		//ストレートキックが入力されたとき
		if kicker_enable {
			//500ミリ秒待つ
			time.Sleep(500 * time.Millisecond)
			//ストレートキックをオフにし、値を0にする
			kicker_enable = false
			kicker_val = 0
		}
		//チップキックが入力されたとき
		if chip_enable {
			//500ミリ秒待つ
			time.Sleep(500 * time.Millisecond)
			//チップキックをオフにし、値を0にする
			chip_enable = false
			chip_val = 0
		}
		if imuError {
			time.Sleep(500 * time.Millisecond)
			imuError = false
		}
		//ループを行うため、少し待機する
		time.Sleep(16 * time.Millisecond)
	}
}

func main() {

	//GPIOの初期化
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	//GPIO 6, 25, 4, 5 を DIP 1, 2, 3, 4 に設定。入力
	dip1 := rpio.Pin(6)
	dip1.Input()
	dip1.PullUp()
	dip2 := rpio.Pin(25)
	dip2.Input()
	dip2.PullUp()
	dip3 := rpio.Pin(4)
	dip3.Input()
	dip3.PullUp()
	dip4 := rpio.Pin(5)
	dip4.Input()
	dip4.PullUp()

	// HEXの値を表示
	hex := dip1.Read() ^ 1 + (dip2.Read()^1)*2 + (dip3.Read()^1)*4 + (dip4.Read()^1)*8
	fmt.Println("GOT ID FROM DIP SW:", int(hex))
	diptoid := int(hex)

	netInterfaceAddresses, _ := net.InterfaceAddrs()

	ip := "0.0.0.0"
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip = networkIp.IP.String()
		}
	}

	//Ctrl+Cを押したときに終了するようにする
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			//終了時にGPIOを解放する
			rpio.Close()
			log.Println("Bye")
			os.Exit(0)
		}
	}()

	//Hostnameを取得する
	cmd := exec.Command("hostname")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	//もし初期値のraspberrypiだったら
	if string(out) == "raspberrypi\n" {
		//UNIX時間の下5桁を取得する
		unixtime := time.Now().UnixNano()
		log.Println("Unixtime is " + fmt.Sprintf("%d", unixtime))
		unixtime = unixtime % 100000

		//Hostnameをracoon-XXXXXに変更する
		hostname := "racoon-" + fmt.Sprintf("%05d", unixtime)

		log.Println("Change Hostname To " + hostname)
		//Change Hostname
		//hostnamectl set-hostname raspberrypi コマンド実行
		cmd = exec.Command("hostnamectl", "set-hostname", hostname)
		cmd.Run()

		//再起動
		log.Println("=====Reboot=====")

		buzzer := rpio.Pin(12)
		buzzer.Mode(rpio.Pwm)
		buzzer.Freq(1175 * 64)
		buzzer.DutyCycle(16, 32)
		time.Sleep(500 * time.Millisecond)
		buzzer.DutyCycle(0, 32)
		buzzer.Freq(1396 * 64)
		buzzer.DutyCycle(16, 32)
		time.Sleep(500 * time.Millisecond)
		buzzer.DutyCycle(0, 32)
		buzzer.Freq(1760 * 64)
		buzzer.DutyCycle(16, 32)
		time.Sleep(500 * time.Millisecond)
		buzzer.DutyCycle(0, 32)

		//reboot コマンド実行
		cmd = exec.Command("reboot")
		cmd.Run()

	}

	//MyIDで指定したロボットIDを取得
	var MyID uint32 = uint32(diptoid)

	chclient := make(chan bool)
	chserver := make(chan bool)
	chserial := make(chan bool)
	chkick := make(chan bool)
	chgpio := make(chan bool)
	chapi := make(chan bool)

	//各並列処理部分
	go RunClient(chclient, MyID, ip)
	go RunServer(chserver, MyID)
	go RunSerial(chserial, MyID)
	go kickCheck(chkick)
	go RunGPIO(chgpio)
	go RunApi(chapi, MyID)

	<-chclient
	<-chserver
	<-chserial
	<-chkick
	<-chgpio
	<-chapi
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
