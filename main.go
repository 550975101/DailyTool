package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/godror/godror"
	"gopkg.in/gomail.v2"
	"mime"
	"os"
	"strconv"
	"time"
)

func getConn() *sql.DB {
	db, err := sql.Open("godror", `user="xbk_keeper" password="xbk_keeper" connectString="10.6.3.52:1521/rptall"`)
	if err != nil {
		fmt.Println("数据库连接异常", err)
	}
	return db
}

func SendMail(subject, body string, filePath, fileName string) error {
	//定义收件人
	mailTo := []string{
		"huangdu@zihexin.com",
		"wangjinrong@zihexin.com",
	}
	mailConn := map[string]string{
		"user": "system_monitor@zihexin.com",
		"pass": "zhxtech01!SM2021",
		"host": "10.6.4.78",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], mailConn["user"]))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if len(filePath) > 0 {
		m.Attach(filePath,
			gomail.Rename(fileName),
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": {
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", fileName)),
				},
			}),
		)
	}
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}

func genCsv(data []map[string]string, fileName string) {
	f, _ := os.Create(fileName)
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入utf-8 编码
	w := csv.NewWriter(f)
	w.Write([]string{"门店名", "消费笔数"})
	for _, datum := range data {
		w.Write([]string{datum["storeName"], datum["num"]})
	}
	w.Flush()
}

func main() {
	nTime := time.Date(2022, 12, 16, 01, 00, 00, 00, time.Local)
	if nTime.Weekday().String() == "Friday" {
		yesTime := nTime.AddDate(0, 0, -1)
		oldTime := nTime.AddDate(0, 0, -7)
		yesTimeStr := yesTime.Format("20060102")
		oldTimeStr := oldTime.Format("20060102")
		//logDay := yesTime.Format("20060102")
		fmt.Println(oldTimeStr, yesTimeStr)
		db := getConn()
		sqlStr := "SELECT s.store_name AS storeName, COUNT(1) as num FROM sbux_card_trans s WHERE s.cardname = '星礼卡' AND s.swap_date BETWEEN to_date('%v 00:00:00', 'yyyymmdd hh24:mi:ss') AND to_date('%v 23:59:59', 'yyyymmdd hh24:mi:ss') AND s.trantype = '消费' GROUP BY s.store_name"
		sqlStr = fmt.Sprintf(sqlStr, oldTimeStr, yesTimeStr)
		rows, err := db.Query(sqlStr)
		if err != nil {
			fmt.Println("查询异常", err)
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				fmt.Println("关闭数据库失败:", err)
			}
		}(rows)
		var total int
		var strList []map[string]string
		for rows.Next() {
			var storeName string
			var num int
			err := rows.Scan(&storeName, &num)
			if err != nil {
				fmt.Println(err)
			}
			m := map[string]string{"storeName": storeName, "num": fmt.Sprintf("%d", num)}
			strList = append(strList, m)
			total = total + 1
		}
		//处理结尾处的异常
		err = rows.Err()
		if err != nil {
			fmt.Println("查询数据失败", err)
		}
		pwd, err := os.Getwd()
		filePath := pwd + string(os.PathSeparator) + "files" + string(os.PathSeparator) + "csvNew" + string(os.PathSeparator) + oldTimeStr + "-" + yesTimeStr + ".csv"
		err = os.MkdirAll(pwd+string(os.PathSeparator)+"files"+string(os.PathSeparator)+"csvNew", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		genCsv(strList, filePath)
		err = SendMail(oldTimeStr+"-"+yesTimeStr+"消费门店统计", "您好: </br> &nbsp&nbsp&nbsp&nbsp&nbsp&nbsp共计"+fmt.Sprintf("%d", total)+"个门店有消费记录", filePath, oldTimeStr+"-"+yesTimeStr+".csv")
		if err != nil {
			fmt.Println(err)
		}
	}
}
