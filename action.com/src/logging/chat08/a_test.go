package chat08

//一个简单的单元测试

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const balloX = "\u2717"

func TestDownload(t *testing.T) {

	url := "http://www.sina.cn"
	statusCode := 200

	t.Log("Given need to test downloading content.")
	{
		t.Logf("\tWhen Checking \"%s\" for status code \"%d\"", url, statusCode)
		{

			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", balloX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v ", statusCode, balloX, resp.StatusCode)
			}

		}
	}

}
