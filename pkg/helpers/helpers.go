/**
  @author:panliang
  @data:2021/8/4
  @note
**/
package helpers

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

