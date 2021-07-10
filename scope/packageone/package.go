package packageone

import "fmt"

var PackageVar = "Exported Package Variable: PackageVar"

func PrintMe(s1, s2, s3 string) {
	fmt.Println(s1, "//", s2, "//", s3)

}
