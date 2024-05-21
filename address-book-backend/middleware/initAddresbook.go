package middleware

import (
	"bufio"
	"embed"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/JJDoneAway/addressbook/models"
)

//go:embed DemoUser.txt
var demoUsers embed.FS

func AddDummies() {
	fs, err := demoUsers.Open("DemoUser.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fs)
	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		name := strings.Split(fileScanner.Text(), ", ")
		user := models.Address{FirstName: name[1], LastName: name[0], Email: fmt.Sprintf("%v@%v.de", name[1], name[0]), Phone: nextPhoneNumber()}
		if err := user.InsertAddress(); err != nil {
			slog.Error("DUMMY", "Message", "Not able to insert dummy", "Error", err)
		}
		count++
	}
	fs.Close()
	slog.Info("Inserted dummy users", "amount", count)
}

func nextPhoneNumber() string {
	min := 10000000000
	max := 99999999999
	return fmt.Sprint("+49", rand.Intn(max-min)+min)
}
