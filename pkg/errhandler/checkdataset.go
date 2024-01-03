package errhandler

import(
	"io/fs"
)

func CheckArgs(args []string) bool{
	path := "/math-skills/cmd/" + args[1]
	if len(args) != 2 {
		return false
	} else if !fs.ValidPath(path){
		return false
	}
	return true
}

func CheckDataSet(data []string) bool{
	return true
}