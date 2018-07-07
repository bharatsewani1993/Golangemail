package main

import(
  "net/smtp"
  "fmt"
  "log"
)

func main(){
   //authentication information
   auth := smtp.PlainAuth("","username","password","smtpserver")

   //connect to server, authenticate and send email
    to := []string{"emailid"}
    msg := []byte("Hello World")

    err := smtp.SendMail("smtpserver:port",auth,"from",to,msg)
    if err != nil {
      log.Fatal(err)
    } else {
      fmt.Println("Mail Sent")
    }
}
