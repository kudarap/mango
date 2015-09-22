package user

import "github.com/javinc/playgo/goryo/resources/user"

func Find(o *user.Options) []user.Model {
    rows, _ := user.Find(*o)

    return rows
}
